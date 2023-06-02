package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Debug bool //TODO debug
var listener net.Listener
var dataDir string

type DataNodeState struct {
	mutex  sync.Mutex
	NodeID string
	//heartbeatInterval time.Duration //TODO heartbeat
	Addr          string
	LeaderAddress string
}

// Create a new datanode
func Create(conf Config) (*DataNodeState, error) {
	var dn DataNodeState
	dn.Addr = conf.Addres + ":" + conf.Port
	var conn, err = grpc.Dial(dn.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("filed to connect", err)
	}
	dataDir = conf.DataDir
	defer conn.Close()
	c := pb.NewBlockReportServiceClient(conn)
	//clientHandler := pb.NewHandleFileRequestsService(conn)
	createBlockReport(c)
	listener, _ = net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		create(dataDir)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started. Listening on port 8080...")
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
