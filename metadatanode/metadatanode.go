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
	mutex  sync.Mutex
	NodeID string
	//heartbeatInterval time.Duration //TODO heartbeat
	Addr          string
	LeaderAddress string
}

type Server struct {
	pb.BlockReportServiceServer
}

var Database1 = NewDatabase()

// Create a new MetadataNode
func Create(conf Config) (*MetadataNodeState, error) {
	var dn MetadataNodeState

	dn.Addr = conf.Addres + ":" + "8080"

	lis, err := net.Listen("tcp", dn.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", dn.Addr)
	s := grpc.NewServer()
	pb.RegisterBlockReportServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//Connect for health check
	for {
		conn, err := grpc.Dial("0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("filed to connect", err)
		}
		defer conn.Close()
		c := pb2.NewHealthClient(conn)
		queryHealthCheck(c)
		time.Sleep(5 * time.Second)
	}

	// create in memory database for storing blockReport structs
	//blockReportStore := New()

	return &dn, nil
}
