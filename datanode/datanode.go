package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"

	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var Debug bool //TODO debug
var listener net.Listener
var dataDir string

type DataNodeState struct {
	mutex             sync.Mutex
	NodeID            string
	heartbeatInterval time.Duration //TODO heartbeat
	Addr              string
	LeaderAddress     string
}

type healthCheckServer struct {
	pb2.HealthServer
}

func ListenHealthCheckServer(adres string) {
	lis, err := net.Listen("tcp", adres)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", adres)
	s := grpc.NewServer()
	pb2.RegisterHealthServer(s, &healthCheckServer{})
	fmt.Println("test1")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("test2")
}

func SendBlockReport(adres string) {
	for {
		//Connect for block report
		conn, err := grpc.Dial(adres, grpc.WithTransportCredentials(insecure.NewCredentials()))
		fmt.Println("test3")
		if err != nil {
			log.Fatal("failed to connect", err)
		}
		defer conn.Close()
		c := pb.NewBlockReportServiceClient(conn)
		sendBlockReport(c)
		time.Sleep(5 * time.Second)
	}
}

// Create a new datanode
func Create(conf Config) (*DataNodeState, error) {
	var dn DataNodeState
	dn.Addr = conf.Addres + ":" + conf.Port
	dn.LeaderAddress = conf.LeaderAddress + ":" + conf.LeaderPort
	return &dn, nil
}

func ListenForCommands() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		var fileSize = 1024
		buffer := make([]byte, fileSize)
		conn.Read(buffer)
		chunkNum := 0
		chunkSize := 128
		var chunkCount = fileSize / chunkSize
		var chunkPadding = chunkSize - (fileSize % chunkSize)
		fmt.Println("padding: ", chunkPadding)
		for {
			chunkName := fmt.Sprintf("%s.%03d", "x.txt", chunkNum)
			path := filepath.Join(dataDir, chunkName)
			chunkFile, err := os.Create(path)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer chunkFile.Close()
			_, err = chunkFile.Write(buffer[chunkNum*chunkSize : (chunkNum+1)*chunkSize])
			if err != nil {
				fmt.Println(err)
			}
			chunkNum++
			if chunkNum >= chunkCount {
				break
			}
		}
	}
}
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
