package datanode

import (
	"net"
)

type Config struct {
	DataDir       string
	Debug         bool
	Port          string
	Addres        string
	Listener      net.Listener
	LeaderAddress string
	LeaderPort    string
}
