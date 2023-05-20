package datanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"context"
	"log"
)

// A proto contains a list of all blocks on a DataNode.
type BlockReportItem struct {
	FileName       string
	BlockID        string
	DataNodeNumber string
}

func sendBlockReport(c pb.BlockReportServiceClient) {
	log.Println("---createBlockReport---")

	report := &pb.BlockReport{
		FileName:       "test",
		BlockID:        1,
		DataNodeNumber: 1,
	}

	res, err := c.SendBlockReport(context.Background(), report)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from server: %v", res)
}
