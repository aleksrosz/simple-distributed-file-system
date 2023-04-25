package metadatanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto"
	"context"
	"log"
)

func (s *Server) SendBlockReport(ctx context.Context, in *pb.BlockReport) (*pb.BlockReport, error) {
	log.Printf("function SendBlockReport was invoked with %v", in)

	data := &BlockReportItem{
		FileName:       in.FileName,
		BlockID:        in.BlockID,
		DataNodeNumber: in.DataNodeNumber,
	}

	return &pb.BlockReport{
		FileName: data.FileName,
	}, nil
}

// A proto contains a list of all blocks on a DataNode.
type BlockReportItem struct {
	FileName       string
	BlockID        int32
	DataNodeNumber int32
}

// TODO implement this correctly
func saveToDB(data *BlockReportItem) *pb.BlockReport {
	return &pb.BlockReport{
		FileName:       data.FileName,
		BlockID:        data.BlockID,
		DataNodeNumber: data.DataNodeNumber,
	}
}
