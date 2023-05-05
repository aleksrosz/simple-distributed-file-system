package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

var Debug bool //TODO debug

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
	// TODO Networking
	//dn.heartbeatInterval = conf.HeartbeatInterval
	//dn.LeaderAddress = conf.LeaderAddress

	// TODO gRPC
	//lis, err := net.Listen("tcp", dn.Addr)
	//	if err != nil {
	//		log.Fatalf("failed to listen: %v", err)
	//	}

	//go dn.grpcstart(conf.Listener) // Start the RPC server https://grpc.io/
	//go dn.Heartbeat() // Check what is the best way to do this.

	conn, err := grpc.Dial(dn.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("filed to connect", err)
	}
	defer conn.Close()

	c := pb.NewBlockReportServiceClient(conn)

	createBlockReport(c)

	return &dn, nil
}
