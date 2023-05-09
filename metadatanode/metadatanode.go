package metadatanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
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

	dn.Addr = conf.Addres + ":" + conf.Port
	//dn.heartbeatInterval = conf.HeartbeatInterval //TODO heartbeat

	// TODO gRPC
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
	//go dn.grpcstart(conf.Listener) // Start the RPC server https://grpc.io/
	//go dn.Heartbeat() //TODO heartbeat

	// create in memory database for storing blockReport structs
	//blockReportStore := New()

	return &dn, nil
}
