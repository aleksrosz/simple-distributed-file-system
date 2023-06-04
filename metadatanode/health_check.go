package metadatanode

import (
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"context"
	"fmt"
	"log"
	"time"
)

func queryHealthCheck(c pb2.HealthClient) (*pb2.HealthCheckResponse, error) {
	log.Println("---query_healthcheck---")

	res, err := c.Check(context.Background(), &pb2.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}
	return res, err
}

// TODO because I assumed that ID in database it is this same as DataNodeNumber
func addHealthCheckValuesToDatabase(dataNodeNumer int32, results *pb2.HealthCheckResponse) bool {
	var status []pb2.HealthCheckResponse_ServingStatus
	status = append(status, results.Status) //todo jeżeli już trzy będą wypełnione to skasować numer 0

	res2 := DatanodeItem{
		Status:         status,
		DataNodeNumber: results.DataNodeNumber,
		IpAddr:         results.IpAddress,
		LastContact:    time.Now().Unix(),
	}

	DatanodeDatabase.Add(res2)
	data, err := DatanodeDatabase.Get(0)
	fmt.Println(err)
	fmt.Println("database get2 2", data.IpAddr, data.DataNodeNumber, data.Status)

	return true
}
