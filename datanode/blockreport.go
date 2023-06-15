package datanode

import (
	"aleksrosz/simple-distributed-file-system/proto/block_report"
	"context"
	"log"
)

// A proto contains a list of all blocks on a DataNode.
type BlockReportItem struct {
	FileName       string
	BlockID        string
	DataNodeNumber string
}

func sendBlockReport(c block_report.BlockReportServiceClient) {
	log.Println("---createBlockReport---")

	report := &block_report.BlockReport{
		FileName:       "test",
		BlockID:        1,
		DataNodeNumber: 1,
	}

	res, err := c.SendBlockReport(context.Background(), report)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Response from server: %v", res)
}
