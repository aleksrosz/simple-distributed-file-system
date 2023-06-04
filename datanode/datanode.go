package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"bytes"
	"errors"

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

type handleFileRequestServiceServer struct {
	pb.HandleFileRequestsServiceServer
}

type FileCommand struct {
	fileCommand int32
	fileName    string
	fileSize    int
	fileData    bytes.Buffer
}

type FileResponse struct {
	message  string
	fileName string
	fileSize int
	fileData bytes.Buffer
}

func (s *handleFileRequestServiceServer) HandleFileService(ctx interface{}, fc *FileCommand) (FileResponse, error) {
	// 0 = odczyt  1 = zapis  -1 = usun
	switch fc.fileCommand {
	case 0:
		{
			fileData, _ := asembleFile(fc.fileName)
			return FileResponse{
				message:  "file retrieved",
				fileName: fc.fileName,
				fileSize: fileData.Len(),
				fileData: fileData,
			}, nil
		}
	case 1:
		{
			splitFile(fc.fileName, fc.fileData, fc.fileSize)
			return FileResponse{
				message:  "file saved",
				fileName: fc.fileName,
				fileSize: 0,
			}, nil
		}
	case -1:
		{
			deleteChunks(fc.fileName)
			return FileResponse{
				message:  "file deleted",
				fileName: fc.fileName,
				fileSize: 0,
			}, nil
		}

	}
	unknownCommandErr := errors.New("unknown command")
	return FileResponse{}, unknownCommandErr
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

func asembleFile(fileName string) (bytes.Buffer, error) {
	fileData := bytes.Buffer{}

	chunkNum := 0
	chunkSize := 128
	for {
		chunkName := fmt.Sprintf("%s.%03d", fileName, chunkNum)
		path := filepath.Join(dataDir, chunkName)
		chunkFile, err := os.Open(path)
		if err != nil {
			break
		}
		buffer := make([]byte, chunkSize)
		chunkFile.Read(buffer)
		fileData.Write(buffer)
		chunkNum++
	}
	return fileData, nil
}

func splitFile(fileName string, fileData bytes.Buffer, fileSize int) {
	chunkNum := 0
	chunkSize := 128
	var chunkCount = fileSize / chunkSize
	var chunkPadding = chunkSize - (fileSize % chunkSize)
	fmt.Println("padding: ", chunkPadding)
	for {
		chunkName := fmt.Sprintf("%s.%03d", fileName, chunkNum)
		path := filepath.Join(dataDir, chunkName)
		chunkFile, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer chunkFile.Close()
		_, err = chunkFile.Write(fileData.Bytes()[chunkNum*chunkSize : (chunkNum+1)*chunkSize])
		if err != nil {
			fmt.Println(err)
		}
		chunkNum++
		if chunkNum >= chunkCount {
			break
		}
	}
}

func deleteChunks(fileName string) {
	fmt.Println("delete: ", fileName)
	chunkNum := 0
	for {
		chunkName := fmt.Sprintf("%s.%03d", fileName, chunkNum)
		path := filepath.Join(dataDir, chunkName)
		e := os.Remove(path)
		if e != nil {
			return
		}
	}
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
