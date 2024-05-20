package main

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"main/storage"
	"strconv"
)

type storageServer struct {
	storage.UnsafeStorageServiceServer
}

func (st *storageServer) Get(ctx context.Context, req *storage.GetRequest) (*storage.GetResponse, error) {

	log.Printf("GET обьект: %+v\n", req)

	response, err := GetToPod(_masterPort, req)
	if err != nil {

		_mutexPorts.Lock()
		log.Printf("Мастер-сервер на порту %d вышел из строя. При попытке провести GET.", _masterPort)

		for idx, port := range _clusterPorts {

			if port == _masterPort {
				RemoveIndex(_clusterPorts, idx)
				continue
			}

			response, err = GetToPod(_masterPort, req)

			if err == nil {
				_masterPort = port
				log.Printf("Мастер сервером назначен сервер на порту %d\n", _masterPort)
				break
			}

			RemoveIndex(_clusterPorts, idx)
			log.Printf("Cервер на порту %d вышел из строя. При попытке провести Get\n", _masterPort)

		}

		if err != nil {
			log.Fatal("Кластер вышел из строя. При попытке провести GET")
		}

		_mutexPorts.Unlock()
	}

	return response, nil
}
func (st *storageServer) Set(ctx context.Context, req *storage.SetRequest) (*storage.SetResponse, error) {

	log.Printf("SET обьект: %+v\n", req)

	if uuid.FromStringOrNil(req.Uuid) == uuid.Nil {

		example, _ := uuid.NewV7()

		log.Printf("Ошибка SET невалидный uuid. Обьект: %+v\n", req)

		return nil, status.Errorf(codes.InvalidArgument, "Uuid запроса должен соответстствовать формату uuid, например %s", example.String())
	}

	err := SetToPod(_masterPort, req)
	if err != nil {

		_mutexPorts.Lock()
		log.Printf("Мастер-сервер на порту %d вышел из строя. При попытке провести SET.", _masterPort)

		for idx, port := range _clusterPorts {

			if port == _masterPort {
				RemoveIndex(_clusterPorts, idx)
				continue
			}

			err = SetToPod(_masterPort, req)

			if err == nil {
				_masterPort = port
				log.Printf("Мастер сервером назначен сервер на порту %d\n", _masterPort)
				break
			}

			RemoveIndex(_clusterPorts, idx)
			log.Printf("Сервер на порту %d вышел из строя. При попытке провести SET\n", port)

		}

		if err != nil {
			log.Fatal("Кластер вышел из строя. При попытке провести SET")
		}

		_mutexPorts.Unlock()
	}

	return &storage.SetResponse{}, nil
}
func (st *storageServer) Delete(ctx context.Context, req *storage.DeleteRequest) (*storage.DeleteResponse, error) {

	log.Printf("DELETE обьект: %+v\n", req)
	err := DeleteToPod(_masterPort, req)
	if err != nil {

		_mutexPorts.Lock()
		log.Printf("Мастер-сервер на порту %d вышел из строя. При попытке провести DELETE.", _masterPort)

		for idx, port := range _clusterPorts {

			if port == _masterPort {
				RemoveIndex(_clusterPorts, idx)
				continue
			}

			err = DeleteToPod(_masterPort, req)

			if err == nil {
				_masterPort = port
				log.Printf("Мастер сервером назначен сервер на порту %d\n", _masterPort)
				break
			}

			RemoveIndex(_clusterPorts, idx)
			log.Printf("Сервер на порту %d вышел из строя. При попытке провести DELETE\n", port)

		}

		if err != nil {
			log.Fatal("Кластер вышел из строя. При попытке провести DELETE")
		}

		_mutexPorts.Unlock()
	}

	return &storage.DeleteResponse{}, nil
}
func (st *storageServer) GetAll(context.Context, *storage.GetAllRequest) (*storage.GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Нельзя вызвать серверный метод у клиента")
}
func (st *storageServer) GetLastTime(ctx context.Context, req *storage.GetLastTimeRequest) (*storage.GetLastTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Нельзя вызвать серверный метод у клиента")
}

func GetToPod(port int32, request *storage.GetRequest) (*storage.GetResponse, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return nil, errors.New("не удалось подключиться к серверу кластера для get")
	}
	defer conn.Close()
	client := storage.NewStorageServiceClient(conn)
	res, err := client.Get(context.Background(), request)
	if err != nil {
		return nil, errors.New("ошибка при get")
	}

	return res, nil
}

func SetToPod(port int32, request *storage.SetRequest) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return errors.New("не удалось подключиться к серверу кластера для set")
	}
	defer conn.Close()
	client := storage.NewStorageServiceClient(conn)

	request.IsNeedReplication = true

	_, err = client.Set(context.Background(), request)
	if err != nil {
		return errors.New("ошибка при set")
	}

	return nil
}

func DeleteToPod(port int32, request *storage.DeleteRequest) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return errors.New("не удалось подключиться к серверу кластера для delete")
	}
	defer conn.Close()
	client := storage.NewStorageServiceClient(conn)

	request.IsNeedReplication = true

	_, err = client.Delete(context.Background(), request)
	if err != nil {
		return errors.New("ошибка при delete")
	}

	return nil
}
