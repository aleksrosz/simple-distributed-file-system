package metadatanode

import "net"

type Config struct {
	DataDir  string
	Debug    bool
	Port     string
	Addres   string
	Listener net.Listener
	//HeartbeatInterval time.Duration //TODO Heartbeat
}
