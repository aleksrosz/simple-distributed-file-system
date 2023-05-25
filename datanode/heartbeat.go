package datanode

import (
	"time"
)

func (self *DataNodeState) heartbeat() {
	for {
		ping(self)
		time.Sleep(self.heartbeatInterval)
	}
}

func ping(dn *DataNodeState) {

}

/*
// Jeden heartbeat health chekc tylko
// drugi z raportem bloków

func healthCheck(self *DataNodeState) {
	for {

		time.Sleep(self.healthCheckInterval)
	}

}

func sendHealthCheck(c pb.Health_WatchClient) {
	log.Println("---sendHealthCheck---")

	healthChekck := &pb.HealthCheck{
		HealthCheck: true,
	}
}

*/
