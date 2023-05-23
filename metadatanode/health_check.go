package metadatanode

import (
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"context"
	"log"
	"time"
)

func queryHealthCheck(c pb2.HealthClient) *pb2.HealthCheckResponse {
	log.Println("---query_healthcheck---")

	res, err := c.Check(context.Background(), &pb2.HealthCheckRequest{})
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Response from server: %v", res)

	return res
}

// TODO because I assumed that ID in database it is this same as DataNodeNumber
func addHealthCheckValuesToDatabase(dataNodeNumer int32, results *pb2.HealthCheckResponse) bool {
	data, err := DatanodeDatabase.Get(int(dataNodeNumer))
	if err != true {
		// TODO imho it should return error not only bool
		return false
	}

	res2 := DatanodeItem{
		Status:         data.Status,
		DataNodeNumber: results.DataNodeNumber,
		IpAddr:         results.IpAddress,
		LastContact:    time.Now().Unix(),
	}

	DatanodeDatabase.Add(res2)

	return true
}
