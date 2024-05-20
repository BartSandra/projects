package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"main/db"
	heart "main/heartBeat"
	"main/storage"
	"net"
	"slices"
	"strconv"
	"sync"
	"time"
)

var _replica int
var _clusterPorts []int32
var _mutexPorts sync.Mutex

var _db *gorm.DB

var _thisPort int32

func main() {
	prepareServer()
	prepareStorage()

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case _ = <-ticker.C:

				if len(_clusterPorts) > 1 {
					log.Print("Опрос кластера")
					tickPulse()
				} else {
					log.Print("Текущий под является единственным в кластере")
				}

				if len(_clusterPorts) < _replica {
					log.Println("Внимание: количество подов", len(_clusterPorts), "ниже необходимого коэфициэента репликации", _replica)
				}
			}
		}
	}()

	startServer()
}

func RemoveIndex(s []int32, index int) []int32 {
	return append(s[:index], s[index+1:]...)
}

func tickPulse() {

	for idx, port := range _clusterPorts {

		if port == _thisPort {
			continue
		}

		response, err := pulsePod(port)

		if err != nil {
			log.Println("Сервер на порту ", port, " вышел из строя")
			_mutexPorts.Lock()
			_clusterPorts = RemoveIndex(_clusterPorts, idx)
			_mutexPorts.Unlock()
			break
		}

		log.Println("Сервер на порту ", port, " функционирует")

		_mutexPorts.Lock()
		_clusterPorts = append(_clusterPorts, response.ClusterPorts...)
		slices.Sort(_clusterPorts)
		_clusterPorts = slices.Compact(_clusterPorts)
		_mutexPorts.Unlock()
	}
}

func prepareServer() {
	targetPortPtr := flag.Int("port", 0, "Порт уже существующего пода")
	dbPortPtr := flag.Int("db_port", 0, "Порт базы данных")
	replicaCPtr := flag.Int("replicaC", 3, "Коэфициэнт репликации")

	flag.Parse()

	if *targetPortPtr < 0 {
		log.Fatal("необходимо задать порт сервера на кластере")
	}

	if *dbPortPtr <= 0 {
		log.Fatal("порт базы данных не может быть отрицательным или равным нулю")
	}

	if *replicaCPtr < 1 {
		log.Fatal("коэфициэнт репликации должен быть больше 0")
	}

	conn, err := db.Connect(*dbPortPtr)

	if err != nil {
		log.Fatal("не удалось подключиться к бд")
	}

	_db = conn

	_replica = *replicaCPtr

	if *targetPortPtr == 0 {
		fmt.Println("Порт пода предоставляющего кластер не задан. Будет образован новый кластер")
		_thisPort = 3333
		_clusterPorts = append(_clusterPorts, _thisPort)
	} else {
		pulse, err := pulsePod(int32(*targetPortPtr))

		if err != nil {
			log.Fatal(err)
		}

		if pulse.ReplicationCoefficient != int32(_replica) {
			log.Fatal("коэфициэнт репликации заданного пода кластера не равен текущему")
		}

		if len(pulse.ClusterPorts) >= _replica {
			log.Fatal("добавление еще одного пода невозможно. Превышен коэфициэент репликации")
		}

		for i := 3333; i < 3400; i++ {
			if !slices.Contains(pulse.ClusterPorts, int32(i)) {
				_thisPort = int32(i)
			}
		}

		_thisPort = slices.Max(pulse.ClusterPorts) + 1
		_clusterPorts = append(pulse.ClusterPorts, _thisPort)

		_, err = pulsePod(int32(*targetPortPtr))

		if err != nil {
			log.Fatal(err)
		}
	}
}

func prepareStorage() {

	for _, port := range _clusterPorts {

		if port == _thisPort {
			continue
		}

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
		if err != nil {
			log.Fatal("Не удалось подключаться к серверу кластера для обновления данных")
		}
		defer conn.Close()
		client := storage.NewStorageServiceClient(conn)

		response, err := client.GetAll(context.Background(), &storage.GetAllRequest{})
		if err != nil {
			log.Fatal("Не удалось подключаться к серверу кластера для обновления данных")
		}

		for _, record := range response.Record {

			newRecord := &db.Record{
				Uuid:      record.Uuid,
				Name:      record.Name,
				UpdatedAt: record.TimeUpdate.AsTime(),
				DeletedAt: gorm.DeletedAt{Valid: false},
			}

			if record.TimeDelete != nil {
				newRecord.DeletedAt = gorm.DeletedAt{Time: record.TimeDelete.AsTime(), Valid: true}
			}

			_db.Unscoped().Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "uuid"}},
				DoUpdates: clause.AssignmentColumns([]string{"name", "updated_at", "deleted_at"}),
			}).Create(newRecord)
		}

		break
	}
}

func pulsePod(port int32) (*heart.PulseResponse, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return nil, errors.New("не удалось подключиться к серверу кластера")
	}
	defer conn.Close()
	client := heart.NewHeartBeatServiceClient(conn)

	pulse, err := client.Pulse(context.Background(), &heart.PulseRequest{
		IsClient: false,
		Port:     _thisPort,
	})
	if err != nil {
		return nil, errors.New("не удалось подключиться к серверу кластера")
	}

	return pulse, nil
}

func startServer() {
	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(int(_thisPort)))
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer) // чтобы у сервера можно было вызвать reflection
	heart.RegisterHeartBeatServiceServer(grpcServer, &heartBeatServer{})
	storage.RegisterStorageServiceServer(grpcServer, &storageServer{})

	fmt.Println("Начало прослушивания localhost:" + strconv.Itoa(int(_thisPort)))
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
