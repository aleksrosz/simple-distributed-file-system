package metadatanode

import (
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

	//go dn.grpcstart(conf.Listener) // Start the RPC server https://grpc.io/
	//go dn.Heartbeat() //TODO heartbeat

	// create in memory database for storing blockReport structs
	//blockReportStore := New()

	return &dn, nil
}
