package datanode

import (
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"context"
	"log"
)

func (s *healthCheckServer) Check(ctx context.Context, in *pb2.HealthCheckRequest) (*pb2.HealthCheckResponse, error) {
	log.Printf("health chec was invoked %v", in)

	return &pb2.HealthCheckResponse{
		Status:         1,         //TODO
		DataNodeNumber: 3,         //TODO
		IpAddress:      "9.9.9.9", //TODO
	}, nil
}
