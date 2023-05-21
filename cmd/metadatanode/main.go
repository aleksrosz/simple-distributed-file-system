package main

import (
	"aleksrosz/simple-distributed-file-system/metadatanode"
	"fmt"
)

func main() {
	metadatanode1, err := metadatanode.Create(metadatanode.Config{
		DataDir: "./test_directory/metadatanode",
		Debug:   true,
		Port:    "8080",
		Addres:  "0.0.0.0",
	})

	if err != nil {
		return
	}
	fmt.Println(metadatanode1)

	go metadatanode.ListenBlockReportServiceServer("0.0.0.0:8080")
	metadatanode.QueryHealthCheck("0.0.0.0:8081", metadatanode1.HeartbeatInterval)

}
