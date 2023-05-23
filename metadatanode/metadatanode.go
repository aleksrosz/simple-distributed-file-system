package metadatanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"sync"
	"time"
)

var Debug bool //TODO debug

type MetadataNodeState struct {
	mutex             sync.Mutex
	NodeID            string
	HeartbeatInterval time.Duration //TODO heartbeat
	Addr              string
	LeaderAddress     string
}

type Server struct {
	pb.BlockReportServiceServer
}

var BlockReportDatabase = NewDatabase()
var DatanodeDatabase = NewDatanodeDatabase()

func ListenBlockReportServiceServer(adres string) {
	lis, err := net.Listen("tcp", adres)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", adres)
	s := grpc.NewServer()
	pb.RegisterBlockReportServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func QueryHealthCheck(adres string, timeInSec time.Duration) {
	//Connect for health check
	for {
		conn, err := grpc.Dial(adres, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("failed to connect", err)
		}
		defer conn.Close()
		c := pb2.NewHealthClient(conn)
		queryHealthCheck(c)
		time.Sleep(timeInSec * time.Second)
	}
}

// Create a new MetadataNode
func Create(conf Config) (*MetadataNodeState, error) {
	var dn MetadataNodeState

	dn.Addr = conf.Addres + ":" + "8080"

	return &dn, nil
}
