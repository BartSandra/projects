package main

import (
	"context"
	"log"
	heart "main/heartBeat"
	"slices"
)

type heartBeatServer struct {
	heart.UnimplementedHeartBeatServiceServer
}

func (p *heartBeatServer) Pulse(ctx context.Context, req *heart.PulseRequest) (*heart.PulseResponse, error) {

	if req.IsClient {
		log.Print("Пульс от клиента балансировщика")
	} else if req.Port == 0 {
		log.Print("Пульс от стартущего пода")
	} else {
		log.Print("Пульс от ", req.Port)

		_mutexPorts.Lock()
		if !slices.Contains(_clusterPorts, req.Port) {
			_clusterPorts = append(_clusterPorts, req.Port)
			slices.Sort(_clusterPorts)
			_clusterPorts = slices.Compact(_clusterPorts)
		}
		_mutexPorts.Unlock()
	}

	resp := &heart.PulseResponse{
		Port:                   _thisPort,
		ClusterPorts:           _clusterPorts,
		ReplicationCoefficient: int32(_replica),
	}
	return resp, nil
}
