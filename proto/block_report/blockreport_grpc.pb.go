// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: proto/blockreport.proto

package block_report

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BlockReportService_SendBlockReport_FullMethodName   = "/blockreport.BlockReportService/SendBlockReport"
	BlockReportService_UpdateBlockReport_FullMethodName = "/blockreport.BlockReportService/UpdateBlockReport"
)

// BlockReportServiceClient is the client API for BlockReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockReportServiceClient interface {
	SendBlockReport(ctx context.Context, in *BlockReport, opts ...grpc.CallOption) (*BlockReport, error)
	UpdateBlockReport(ctx context.Context, in *BlockReport, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type blockReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockReportServiceClient(cc grpc.ClientConnInterface) BlockReportServiceClient {
	return &blockReportServiceClient{cc}
}

func (c *blockReportServiceClient) SendBlockReport(ctx context.Context, in *BlockReport, opts ...grpc.CallOption) (*BlockReport, error) {
	out := new(BlockReport)
	err := c.cc.Invoke(ctx, BlockReportService_SendBlockReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockReportServiceClient) UpdateBlockReport(ctx context.Context, in *BlockReport, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BlockReportService_UpdateBlockReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockReportServiceServer is the server API for BlockReportService service.
// All implementations must embed UnimplementedBlockReportServiceServer
// for forward compatibility
type BlockReportServiceServer interface {
	SendBlockReport(context.Context, *BlockReport) (*BlockReport, error)
	UpdateBlockReport(context.Context, *BlockReport) (*emptypb.Empty, error)
	mustEmbedUnimplementedBlockReportServiceServer()
}

// UnimplementedBlockReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockReportServiceServer struct {
}

func (UnimplementedBlockReportServiceServer) SendBlockReport(context.Context, *BlockReport) (*BlockReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBlockReport not implemented")
}
func (UnimplementedBlockReportServiceServer) UpdateBlockReport(context.Context, *BlockReport) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlockReport not implemented")
}
func (UnimplementedBlockReportServiceServer) mustEmbedUnimplementedBlockReportServiceServer() {}

// UnsafeBlockReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockReportServiceServer will
// result in compilation errors.
type UnsafeBlockReportServiceServer interface {
	mustEmbedUnimplementedBlockReportServiceServer()
}

func RegisterBlockReportServiceServer(s grpc.ServiceRegistrar, srv BlockReportServiceServer) {
	s.RegisterService(&BlockReportService_ServiceDesc, srv)
}

func _BlockReportService_SendBlockReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockReportServiceServer).SendBlockReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockReportService_SendBlockReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockReportServiceServer).SendBlockReport(ctx, req.(*BlockReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockReportService_UpdateBlockReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockReportServiceServer).UpdateBlockReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlockReportService_UpdateBlockReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockReportServiceServer).UpdateBlockReport(ctx, req.(*BlockReport))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockReportService_ServiceDesc is the grpc.ServiceDesc for BlockReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockreport.BlockReportService",
	HandlerType: (*BlockReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBlockReport",
			Handler:    _BlockReportService_SendBlockReport_Handler,
		},
		{
			MethodName: "UpdateBlockReport",
			Handler:    _BlockReportService_UpdateBlockReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockreport.proto",
}
