package main

import (
	"aleksrosz/simple-distributed-file-system/datanode"
	"fmt"
)

func main() {
	create, err := datanode.Create(datanode.Config{
		DataDir: "./test_directory/dataNode01",
		Debug:   true,
		Port:    "8080",
		Addres:  "127.0.0.1",
	})

	if err != nil {
		return
	}
	fmt.Println(create)

	datanode.ListenForCommands()

}
