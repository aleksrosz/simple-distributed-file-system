package metadatanode

import (
	pb "aleksrosz/simple-distributed-file-system/proto/block_report"
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
	BlockReportDatabase.Add(*data)

	return &pb.BlockReport{
		FileName: data.FileName,
	}, nil
}
