package main

import (
	"aleksrosz/simple-distributed-file-system/metadatanode"
	"fmt"
)

func main() {
	create, err := metadatanode.Create(metadatanode.Config{
		DataDir: "./test_directory/metadatanode",
		Debug:   true,
		Port:    "8080",
		Addres:  "0.0.0.0",
	})

	if err != nil {
		return
	}
	fmt.Println(create)

}
