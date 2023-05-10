package client

// TODO It should be possible use client to create datanodes and metadatanode on another nodes, with parameters like
// replication
// number of datanodes
// blocksize ???
//
// Tell which file should be uploaded to DFS

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Parse command-line flags
	serverIP := flag.String("s", "", "Server IP address")
	port := flag.String("p", "8080", "Network port")
	filePath := flag.String("f", "", "Path of file to send")
	flag.Parse()

	// Check if required flags are set
	if *serverIP == "" || *filePath == "" {
		fmt.Println("Usage: go run client.go -s <server-ip> -f <file-path> [-p <port>]")
		return
	}

	// Open the input file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Connect to the server
	conn, err := net.Dial("tcp", *serverIP+":"+*port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send the file data to the server
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		conn.Write(buf[:n])
	}
	fmt.Printf("File '%s' sent to server at %s:%s\n", *filePath, *serverIP, *port)
}