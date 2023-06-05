package main

import (
	"aleksrosz/simple-distributed-file-system/datanode"
	"fmt"
)

func main() {
	create, err := datanode.Create(datanode.Config{

		DataDir:       "./test_directory/dataNode01",
		Debug:         true,
		Port:          "8081",
		Addres:        "0.0.0.0",
		LeaderAddress: "0.0.0.0",
		LeaderPort:    "8080",
	})

	if err != nil {
		return
	}
	fmt.Println(create)

	go datanode.ListenHealthCheckServer("0.0.0.0:8081")
	go datanode.ListenFileRequestServiceServer("0.0.0.0:8085")
	datanode.SendBlockReport("0.0.0.0:8080")
}
