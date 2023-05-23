package metadatanode

import (
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"context"
	"log"
)

func queryHealthCheck(c pb2.HealthClient) {
	log.Println("---query_healthcheck---")

	res, err := c.Check(context.Background(), &pb2.HealthCheckRequest{})
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Response from server: %v", res)
}
