// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server-svc/image.proto

/*
Package image_proto is a generated protocol buffer package.

It is generated from these files:

	server-svc/image.proto

It has these top-level messages:

	Image
	ImageResponse
*/
package image_proto

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Image struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ImageResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ImageResponse) Reset()                    { *m = ImageResponse{} }
func (m *ImageResponse) String() string            { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()               {}
func (*ImageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "image_proto.Image")
	proto.RegisterType((*ImageResponse)(nil), "image_proto.ImageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageService service

type ImageServiceClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadImageClient, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadImageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ImageService_serviceDesc.Streams[0], c.cc, "/image_proto.ImageService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceUploadImageClient{stream}
	return x, nil
}

type ImageService_UploadImageClient interface {
	Send(*Image) error
	CloseAndRecv() (*ImageResponse, error)
	grpc.ClientStream
}

type imageServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *imageServiceUploadImageClient) Send(m *Image) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceUploadImageClient) CloseAndRecv() (*ImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ImageService service

type ImageServiceServer interface {
	UploadImage(ImageService_UploadImageServer) error
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).UploadImage(&imageServiceUploadImageServer{stream})
}

type ImageService_UploadImageServer interface {
	SendAndClose(*ImageResponse) error
	Recv() (*Image, error)
	grpc.ServerStream
}

type imageServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *imageServiceUploadImageServer) SendAndClose(m *ImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceUploadImageServer) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image_proto.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _ImageService_UploadImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "server-svc/image.proto",
}

func init() { proto.RegisterFile("server-svc/image.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2d, 0x2e, 0x4b, 0xd6, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x73, 0xe2, 0xc1, 0x1c, 0x25, 0x7d, 0x2e, 0x56, 0x4f, 0x10, 0x57,
	0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc,
	0x06, 0x89, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0x9a, 0x5c, 0xbc, 0x60, 0x0d, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x60, 0xbd, 0x9c, 0x41, 0x30, 0xae, 0x91, 0x3f,
	0x17, 0x0f, 0x58, 0x69, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3d, 0x17, 0x77, 0x68,
	0x41, 0x4e, 0x7e, 0x62, 0x0a, 0xd4, 0x46, 0x3d, 0x24, 0x87, 0xe8, 0x81, 0xc5, 0xa4, 0xa4, 0x30,
	0xc5, 0x60, 0x16, 0x69, 0x30, 0x26, 0xb1, 0x81, 0x85, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xd1, 0xaf, 0xf7, 0xda, 0x00, 0x00, 0x00,
}