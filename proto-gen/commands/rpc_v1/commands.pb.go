// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/rpc_v1/commands.proto

package rpc_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("commands/rpc_v1/commands.proto", fileDescriptor_cf7880bd07f0f11b) }

var fileDescriptor_cf7880bd07f0f11b = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x33, 0xd4, 0x87, 0xf1, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20, 0xc2, 0x52, 0xd2, 0xe8, 0xea, 0x4a, 0x2a, 0x0b, 0x52,
	0xa1, 0x8a, 0xa4, 0x64, 0xb1, 0x18, 0x52, 0x90, 0x99, 0x93, 0x0a, 0x91, 0x36, 0x72, 0xe1, 0xe2,
	0x77, 0x2c, 0x4a, 0x29, 0xcd, 0xcc, 0xcb, 0x77, 0x86, 0xaa, 0x13, 0x32, 0xe4, 0x62, 0x77, 0x86,
	0xa8, 0x11, 0x12, 0xd2, 0x83, 0x68, 0xd2, 0x83, 0x0a, 0x04, 0xa5, 0x16, 0x4a, 0x09, 0xc0, 0xc4,
	0x5c, 0x2b, 0x52, 0x93, 0x83, 0x52, 0x8b, 0x0b, 0x0c, 0x18, 0x9d, 0x0c, 0xa3, 0xf4, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x13, 0x21, 0x06, 0xc2, 0x68, 0xdd, 0xe4,
	0x9c, 0x4c, 0x7d, 0xb0, 0x6d, 0xba, 0xe9, 0xa9, 0x79, 0x50, 0x67, 0x24, 0xb1, 0x81, 0x45, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x9b, 0x41, 0x68, 0xe5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArduinoCommandsClient is the client API for ArduinoCommands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArduinoCommandsClient interface {
	Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (ArduinoCommands_CompileClient, error)
}

type arduinoCommandsClient struct {
	cc *grpc.ClientConn
}

func NewArduinoCommandsClient(cc *grpc.ClientConn) ArduinoCommandsClient {
	return &arduinoCommandsClient{cc}
}

func (c *arduinoCommandsClient) Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (ArduinoCommands_CompileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCommands_serviceDesc.Streams[0], "/rpc_v1.ArduinoCommands/Compile", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCommandsCompileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCommands_CompileClient interface {
	Recv() (*ExecResp, error)
	grpc.ClientStream
}

type arduinoCommandsCompileClient struct {
	grpc.ClientStream
}

func (x *arduinoCommandsCompileClient) Recv() (*ExecResp, error) {
	m := new(ExecResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArduinoCommandsServer is the server API for ArduinoCommands service.
type ArduinoCommandsServer interface {
	Compile(*CompileReq, ArduinoCommands_CompileServer) error
}

// UnimplementedArduinoCommandsServer can be embedded to have forward compatible implementations.
type UnimplementedArduinoCommandsServer struct {
}

func (*UnimplementedArduinoCommandsServer) Compile(req *CompileReq, srv ArduinoCommands_CompileServer) error {
	return status.Errorf(codes.Unimplemented, "method Compile not implemented")
}

func RegisterArduinoCommandsServer(s *grpc.Server, srv ArduinoCommandsServer) {
	s.RegisterService(&_ArduinoCommands_serviceDesc, srv)
}

func _ArduinoCommands_Compile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompileReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCommandsServer).Compile(m, &arduinoCommandsCompileServer{stream})
}

type ArduinoCommands_CompileServer interface {
	Send(*ExecResp) error
	grpc.ServerStream
}

type arduinoCommandsCompileServer struct {
	grpc.ServerStream
}

func (x *arduinoCommandsCompileServer) Send(m *ExecResp) error {
	return x.ServerStream.SendMsg(m)
}

var _ArduinoCommands_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_v1.ArduinoCommands",
	HandlerType: (*ArduinoCommandsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Compile",
			Handler:       _ArduinoCommands_Compile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "commands/rpc_v1/commands.proto",
}
