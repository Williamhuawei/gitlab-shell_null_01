// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SSHUploadPackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Parameters to use with git -c (key=value pairs)
	GitConfigOptions []string `protobuf:"bytes,4,rep,name=git_config_options,json=gitConfigOptions" json:"git_config_options,omitempty"`
}

func (m *SSHUploadPackRequest) Reset()                    { *m = SSHUploadPackRequest{} }
func (m *SSHUploadPackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackRequest) ProtoMessage()               {}
func (*SSHUploadPackRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SSHUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHUploadPackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHUploadPackRequest) GetGitConfigOptions() []string {
	if m != nil {
		return m.GitConfigOptions
	}
	return nil
}

type SSHUploadPackResponse struct {
	// A chunk of raw data from 'git upload-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git upload-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHUploadPackResponse) Reset()                    { *m = SSHUploadPackResponse{} }
func (m *SSHUploadPackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHUploadPackResponse) ProtoMessage()               {}
func (*SSHUploadPackResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SSHUploadPackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHUploadPackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHUploadPackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

type SSHReceivePackRequest struct {
	// 'repository' must be present in the first message.
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// A chunk of raw data to be copied to 'git upload-pack' standard input
	Stdin []byte `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Contents of GL_ID, GL_REPOSITORY, and GL_USERNAME environment variables
	// for 'git receive-pack'
	GlId         string `protobuf:"bytes,3,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	GlRepository string `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository" json:"gl_repository,omitempty"`
	GlUsername   string `protobuf:"bytes,5,opt,name=gl_username,json=glUsername" json:"gl_username,omitempty"`
}

func (m *SSHReceivePackRequest) Reset()                    { *m = SSHReceivePackRequest{} }
func (m *SSHReceivePackRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackRequest) ProtoMessage()               {}
func (*SSHReceivePackRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *SSHReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *SSHReceivePackRequest) GetStdin() []byte {
	if m != nil {
		return m.Stdin
	}
	return nil
}

func (m *SSHReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *SSHReceivePackRequest) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

type SSHReceivePackResponse struct {
	// A chunk of raw data from 'git receive-pack' standard output
	Stdout []byte `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// A chunk of raw data from 'git receive-pack' standard error
	Stderr []byte `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// This field may be nil. This is intentional: only when the remote
	// command has finished can we return its exit status.
	ExitStatus *ExitStatus `protobuf:"bytes,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *SSHReceivePackResponse) Reset()                    { *m = SSHReceivePackResponse{} }
func (m *SSHReceivePackResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHReceivePackResponse) ProtoMessage()               {}
func (*SSHReceivePackResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *SSHReceivePackResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *SSHReceivePackResponse) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *SSHReceivePackResponse) GetExitStatus() *ExitStatus {
	if m != nil {
		return m.ExitStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SSHUploadPackRequest)(nil), "gitaly.SSHUploadPackRequest")
	proto.RegisterType((*SSHUploadPackResponse)(nil), "gitaly.SSHUploadPackResponse")
	proto.RegisterType((*SSHReceivePackRequest)(nil), "gitaly.SSHReceivePackRequest")
	proto.RegisterType((*SSHReceivePackResponse)(nil), "gitaly.SSHReceivePackResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSHService service

type SSHServiceClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error)
}

type sSHServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHServiceClient(cc *grpc.ClientConn) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[0], c.cc, "/gitaly.SSHService/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHUploadPackClient{stream}
	return x, nil
}

type SSHService_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHServiceClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHService_SSHReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSHService_serviceDesc.Streams[1], c.cc, "/gitaly.SSHService/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceSSHReceivePackClient{stream}
	return x, nil
}

type SSHService_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHServiceSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHServiceSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SSHService service

type SSHServiceServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSHService_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSHService_SSHReceivePackServer) error
}

func RegisterSSHServiceServer(s *grpc.Server, srv SSHServiceServer) {
	s.RegisterService(&_SSHService_serviceDesc, srv)
}

func _SSHService_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHUploadPack(&sSHServiceSSHUploadPackServer{stream})
}

type SSHService_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHService_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).SSHReceivePack(&sSHServiceSSHReceivePackServer{stream})
}

type SSHService_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHServiceSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHServiceSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSHService",
	HandlerType: (*SSHServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSHService_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSHService_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh.proto",
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0x75, 0xa4, 0x10, 0xb9, 0xf4, 0x33, 0x64, 0x04, 0xd2, 0x10, 0x7f, 0x48, 0xdd, 0x74, 0x61,
	0x88, 0x81, 0x47, 0x30, 0x26, 0xe8, 0x46, 0x33, 0x0d, 0xeb, 0x66, 0x6c, 0xaf, 0xc3, 0xc4, 0xa1,
	0x53, 0x67, 0xa6, 0x04, 0x12, 0x7d, 0x22, 0x1f, 0xc0, 0x8d, 0x0f, 0x67, 0x32, 0xad, 0x58, 0x50,
	0x96, 0xba, 0xeb, 0x3d, 0xe7, 0xfe, 0x9c, 0x73, 0x6f, 0x07, 0x86, 0xd6, 0xee, 0x96, 0x95, 0xd1,
	0x4e, 0xd3, 0x81, 0x90, 0x8e, 0xab, 0xd3, 0x3c, 0xb4, 0x3b, 0x6e, 0xb0, 0x68, 0xd0, 0xf8, 0x1b,
	0x81, 0x49, 0x9a, 0x6e, 0xb6, 0x95, 0xd2, 0xbc, 0x78, 0xcf, 0xf3, 0x4f, 0x0c, 0x3f, 0xd7, 0x68,
	0x1d, 0x5d, 0x01, 0x18, 0xac, 0xb4, 0x95, 0x4e, 0x9b, 0x53, 0x44, 0x16, 0x24, 0x19, 0xad, 0xe8,
	0xb2, 0xe9, 0xb1, 0x64, 0x67, 0x86, 0x75, 0xb2, 0xe8, 0x04, 0xfa, 0xd6, 0x15, 0xb2, 0x8c, 0xee,
	0x2f, 0x48, 0x12, 0xb2, 0x26, 0xa0, 0x2f, 0x80, 0x0a, 0xe9, 0xb2, 0x5c, 0x97, 0x1f, 0xa5, 0xc8,
	0x74, 0xe5, 0xa4, 0x2e, 0x6d, 0x14, 0x2c, 0x7a, 0xc9, 0x90, 0x8d, 0x85, 0x74, 0xaf, 0x3c, 0xf1,
	0xae, 0xc1, 0xdf, 0x06, 0x0f, 0x7a, 0xe3, 0x80, 0x4d, 0x3b, 0x15, 0x15, 0x37, 0x7c, 0x8f, 0x0e,
	0x8d, 0x8d, 0xbf, 0xc0, 0xf4, 0x4a, 0xac, 0xad, 0x74, 0x69, 0x91, 0xce, 0x60, 0x60, 0x5d, 0xa1,
	0x6b, 0xe7, 0x95, 0x86, 0xac, 0x8d, 0x5a, 0x1c, 0x8d, 0x69, 0x25, 0xb5, 0x11, 0x5d, 0xc3, 0x08,
	0x8f, 0xd2, 0x65, 0xd6, 0x71, 0x57, 0xdb, 0xa8, 0x77, 0x69, 0xef, 0xf5, 0x51, 0xba, 0xd4, 0x33,
	0x0c, 0xf0, 0xfc, 0x1d, 0xff, 0x20, 0x7e, 0x3c, 0xc3, 0x1c, 0xe5, 0x01, 0xff, 0xcd, 0xb2, 0x1e,
	0x41, 0x5f, 0xa8, 0x4c, 0x16, 0x5e, 0xd2, 0x90, 0x05, 0x42, 0xbd, 0x29, 0xe8, 0x73, 0xb8, 0x13,
	0x2a, 0xeb, 0x4c, 0x08, 0x3c, 0x19, 0x0a, 0xf5, 0xbb, 0x37, 0x7d, 0x06, 0x23, 0xa1, 0xb2, 0xda,
	0xa2, 0x29, 0xf9, 0x1e, 0xa3, 0xbe, 0x4f, 0x01, 0xa1, 0xb6, 0x2d, 0x12, 0x7f, 0x85, 0xd9, 0xb5,
	0xfa, 0xff, 0xb8, 0xbd, 0xd5, 0x77, 0x02, 0x90, 0xa6, 0x9b, 0x14, 0xcd, 0x41, 0xe6, 0x48, 0x19,
	0xdc, 0x5d, 0x9c, 0x92, 0x3e, 0xfe, 0x55, 0xff, 0xb7, 0xdf, 0x71, 0xfe, 0xe4, 0x06, 0xdb, 0x38,
	0x88, 0xef, 0x25, 0xe4, 0x25, 0xa1, 0x5b, 0x78, 0x78, 0xe9, 0x90, 0x76, 0xcb, 0xfe, 0xbc, 0xdb,
	0xfc, 0xe9, 0x2d, 0xba, 0xdb, 0xf6, 0xc3, 0xc0, 0x3f, 0x95, 0xf5, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0x65, 0x3d, 0xab, 0x4d, 0x03, 0x00, 0x00,
}