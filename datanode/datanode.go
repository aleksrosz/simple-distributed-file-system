package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	pb2 "aleksrosz/simple-distributed-file-system/proto/health_check"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"sync"
	"time"
)

var Debug bool //TODO debug

type DataNodeState struct {
	mutex             sync.Mutex
	NodeID            string
	heartbeatInterval time.Duration //TODO heartbeat
	Addr              string
	LeaderAddress     string
}

type HealthCheckServer struct {
	pb2.HealthServer
}

// Create a new datanode
func Create(conf Config) (*DataNodeState, error) {
	var dn DataNodeState

	dn.Addr = conf.Addres + ":" + conf.Port
	dn.LeaderAddress = conf.LeaderAddress + ":" + conf.LeaderPort

	lis, err := net.Listen("tcp", dn.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", dn.Addr)
	s := grpc.NewServer()
	pb2.RegisterHealthServer(s, &HealthCheckServer{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("test")
	//Connect for block report
	conn, err := grpc.Dial(dn.LeaderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("filed to connect", err)
	}
	defer conn.Close()
	c := pb.NewBlockReportServiceClient(conn)
	sendBlockReport(c)

	return &dn, nil
}
