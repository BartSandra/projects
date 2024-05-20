package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm/clause"
	"main/db"
	"main/storage"
	"strconv"
	"time"
)

type storageServer struct {
	storage.UnimplementedStorageServiceServer
}

func (st *storageServer) Get(ctx context.Context, req *storage.GetRequest) (*storage.GetResponse, error) {

	record := db.Record{Uuid: req.Uuid}

	result := _db.First(&record)

	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &storage.GetResponse{
		Uuid: record.Uuid,
		Name: record.Name,
		Time: timestamppb.New(record.UpdatedAt),
	}, nil
}
func (st *storageServer) GetAll(ctx context.Context, req *storage.GetAllRequest) (*storage.GetAllResponse, error) {

	var records []db.Record
	_db.Unscoped().Find(&records)

	var result []*storage.Record

	for _, val := range records {

		newRecord := &storage.Record{
			Uuid:       val.Uuid,
			Name:       val.Name,
			TimeUpdate: timestamppb.New(val.UpdatedAt),
		}

		if val.DeletedAt.Valid {
			newRecord.TimeDelete = timestamppb.New(val.DeletedAt.Time)
		}
		result = append(result, newRecord)
	}

	return &storage.GetAllResponse{
		Record: result,
	}, nil
}
func (st *storageServer) GetLastTime(ctx context.Context, req *storage.GetLastTimeRequest) (*storage.GetLastTimeResponse, error) {
	return &storage.GetLastTimeResponse{Time: timestamppb.New(GetLastTime())}, nil
}
func (st *storageServer) Set(ctx context.Context, req *storage.SetRequest) (*storage.SetResponse, error) {

	_db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&db.Record{
		Uuid: req.Uuid,
		Name: req.Name,
	})

	if req.IsNeedReplication {
		for _, port := range _clusterPorts {
			if port == _thisPort {
				continue
			}
			_ = replicateSetToPod(port, req)
		}
	}

	return &storage.SetResponse{}, nil
}
func (st *storageServer) Delete(ctx context.Context, req *storage.DeleteRequest) (*storage.DeleteResponse, error) {

	record := db.Record{Uuid: req.Uuid}

	_db.Delete(&record)

	if req.IsNeedReplication {

		for _, port := range _clusterPorts {
			if port == _thisPort {
				continue
			}
			_ = replicateDeleteToPod(port, req)
		}
	}

	return nil, nil
}

func replicateSetToPod(port int32, request *storage.SetRequest) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return errors.New("не удалось подключиться к серверу кластера для репликации set")
	}
	defer conn.Close()
	client := storage.NewStorageServiceClient(conn)

	request.IsNeedReplication = false

	_, err = client.Set(context.Background(), request)
	if err != nil {
		return errors.New("ошибка при репликации set")
	}

	return nil
}

func replicateDeleteToPod(port int32, request *storage.DeleteRequest) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(port)), opts...)
	if err != nil {
		return errors.New("не удалось подключиться к серверу кластера для репликации set")
	}
	defer conn.Close()
	client := storage.NewStorageServiceClient(conn)

	request.IsNeedReplication = false

	_, err = client.Delete(context.Background(), request)
	if err != nil {
		return errors.New("ошибка при репликации set")
	}

	return nil
}

func GetLastTime() time.Time {
	var lastTime time.Time

	_db.Raw("SELECT max(time) FROM (SELECT max(updated_at) AS time FROM records UNION SELECT max(deleted_at) AS time FROM records) AS un").Scan(&lastTime)

	return lastTime
}
