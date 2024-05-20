package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	heart "main/heartBeat"
	"main/storage"
	"net"
	"slices"
	"strconv"
	"sync"
	"time"
)

var _clusterPorts []int32
var _masterPort int32
var _mutexPorts sync.Mutex

var _replica int
var _thisPort int32

func main() {

	prepareServer()

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case _ = <-ticker.C:
				if len(_clusterPorts) > 0 {
					log.Print("Опрос мастера")
					_mutexPorts.Lock()
					tickPulse()
					_mutexPorts.Unlock()
				} else {
					log.Fatal("Кластер вышел из строя")
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
	response, err := pulseCluster(_masterPort)

	if err != nil {
		log.Println("Мастер-сервер на порту ", _masterPort, " вышел из строя")
		_clusterPorts = RemoveIndex(_clusterPorts, 0)
		if len(_clusterPorts) > 0 {
			_masterPort = _clusterPorts[0]
			log.Println("Смена мастер-сервера на", _masterPort)
		}
	} else {
		log.Println("Мастер-сервер на порту ", _masterPort, " функционирует")
		_clusterPorts = append(_clusterPorts, response.ClusterPorts...)
		slices.Sort(_clusterPorts)
		_clusterPorts = slices.Compact(_clusterPorts)

		if _masterPort != _clusterPorts[0] {
			_masterPort = _clusterPorts[0]
			log.Println("Смена мастер сервера на", _masterPort)
		}
	}

	log.Printf("Кластер: %+v\n", _clusterPorts)
}

func pulseCluster(port int32) (*heart.PulseResponse, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return nil, errors.New("не удалось подключиться к серверу кластера")
	}
	defer conn.Close()
	client := heart.NewHeartBeatServiceClient(conn)

	pulse, err := client.Pulse(context.Background(), &heart.PulseRequest{
		IsClient: true,
	})
	if err != nil {
		return nil, errors.New("не удалось подключиться к серверу кластера")
	}

	return pulse, nil
}

func prepareServer() {
	targetPortPtr := flag.Int("pod_port", 0, "Порт уже существующего пода в кластере")
	listenPortPtr := flag.Int("listen_port", 6000, "Порт для прослушивания")

	flag.Parse()

	if *targetPortPtr < 0 {
		log.Fatal("порт не может быть отрицательным")
	}

	if *targetPortPtr == 0 {
		log.Fatal("Порт пода предоставляющего кластер не задан. Для подключения к кластеру необходимо указать порт")
	} else {
		pulse, err := pulseCluster(int32(*targetPortPtr))

		if err != nil {
			log.Fatal(err)
		}

		_replica = int(pulse.ReplicationCoefficient)
		_clusterPorts = pulse.ClusterPorts
		_masterPort = int32(*targetPortPtr)
	}

	if *listenPortPtr == 0 {
		_thisPort = 3333
	} else {
		_thisPort = int32(*listenPortPtr)
	}
}

func startServer() {
	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(int(_thisPort)))
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer) // чтобы у сервера можно было вызвать reflection
	storage.RegisterStorageServiceServer(grpcServer, &storageServer{})

	fmt.Println("Начало прослушивания localhost:" + strconv.Itoa(int(_thisPort)))
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
