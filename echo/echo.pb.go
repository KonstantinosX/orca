// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Time
	Location
	Device
	Request
	Reply
*/
package echo

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
const _ = proto.ProtoPackageIsVersion1

// Time preserves nanosecond latency measurements by using a custom time
// struct which should include either seconds or nanoseconds since the Unix
// epoch as unsigned int64. In Go, use time.Unix to parse this field.
type Time struct {
	Seconds     int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

// Reset the message
func (m *Time) Reset() { *m = Time{} }

// String returns a string representation of the message
func (m *Time) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Time) ProtoMessage() {}

// Descriptor is a generated method
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Location describes the results of an GeoIP query
type Location struct {
	Ipaddr       string  `protobuf:"bytes,1,opt,name=ipaddr" json:"ipaddr,omitempty"`
	Latitude     float64 `protobuf:"fixed64,2,opt,name=latitude" json:"latitude,omitempty"`
	Longitude    float64 `protobuf:"fixed64,3,opt,name=longitude" json:"longitude,omitempty"`
	City         string  `protobuf:"bytes,4,opt,name=city" json:"city,omitempty"`
	Postal       string  `protobuf:"bytes,5,opt,name=postal" json:"postal,omitempty"`
	Country      string  `protobuf:"bytes,6,opt,name=country" json:"country,omitempty"`
	Organization string  `protobuf:"bytes,7,opt,name=organization" json:"organization,omitempty"`
	Domain       string  `protobuf:"bytes,8,opt,name=domain" json:"domain,omitempty"`
}

// Reset the message
func (m *Location) Reset() { *m = Location{} }

// String returns a string representation of the message
func (m *Location) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Location) ProtoMessage() {}

// Descriptor is a generated method
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Device describes an sender or a receiver on the network.
type Device struct {
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IPAddr   string    `protobuf:"bytes,2,opt,name=ipaddr" json:"ipaddr,omitempty"`
	Domain   string    `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Location *Location `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

// Reset the message
func (m *Device) Reset() { *m = Device{} }

// String returns a string representation of the message
func (m *Device) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Device) ProtoMessage() {}

// Descriptor is a generated method
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// GetLocation returns the location message if it exists
func (m *Device) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

// Request is used to measure latency and uptime.
type Request struct {
	Sequence int64   `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Sender   *Device `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Sent     *Time   `protobuf:"bytes,3,opt,name=sent" json:"sent,omitempty"`
	TTL      int64   `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	Ping     int64   `protobuf:"varint,5,opt,name=ping" json:"ping,omitempty"`
	Payload  []byte  `protobuf:"bytes,15,opt,name=payload,proto3" json:"payload,omitempty"`
}

// Reset the message
func (m *Request) Reset() { *m = Request{} }

// String returns a string representation of the message
func (m *Request) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Request) ProtoMessage() {}

// Descriptor is a generated method
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetSender returns the sender message if it exists
func (m *Request) GetSender() *Device {
	if m != nil {
		return m.Sender
	}
	return nil
}

// GetSent returns the sent time message if it exists
func (m *Request) GetSent() *Time {
	if m != nil {
		return m.Sent
	}
	return nil
}

// Reply is used to respond to EchoRequest messages.
type Reply struct {
	Sequence int64    `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Receiver *Device  `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	Received *Time    `protobuf:"bytes,3,opt,name=received" json:"received,omitempty"`
	Echo     *Request `protobuf:"bytes,4,opt,name=echo" json:"echo,omitempty"`
}

// Reset the message
func (m *Reply) Reset() { *m = Reply{} }

// String returns a string representation of the message
func (m *Reply) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Reply) ProtoMessage() {}

// Descriptor is a generated method
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetReceiver returns the receiver message if it exists
func (m *Reply) GetReceiver() *Device {
	if m != nil {
		return m.Receiver
	}
	return nil
}

// GetReceived returns the received time message if it exists
func (m *Reply) GetReceived() *Time {
	if m != nil {
		return m.Received
	}
	return nil
}

// GetEcho returns the echo request message if it exists
func (m *Reply) GetEcho() *Request {
	if m != nil {
		return m.Echo
	}
	return nil
}

func init() {
	proto.RegisterType((*Time)(nil), "echo.Time")
	proto.RegisterType((*Location)(nil), "echo.Location")
	proto.RegisterType((*Device)(nil), "echo.Device")
	proto.RegisterType((*Request)(nil), "echo.Request")
	proto.RegisterType((*Reply)(nil), "echo.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Orca service

// OrcaClient is a generated interface
type OrcaClient interface {
	// Reflect allows nodes to respond to echo requests with echo replies.
	Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type orcaClient struct {
	cc *grpc.ClientConn
}

// NewOrcaClient is a generated function
func NewOrcaClient(cc *grpc.ClientConn) OrcaClient {
	return &orcaClient{cc}
}

func (c *orcaClient) Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/echo.Orca/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orca service

// OrcaServer is a generated interface
type OrcaServer interface {
	// Reflect allows nodes to respond to echo requests with echo replies.
	Echo(context.Context, *Request) (*Reply, error)
}

// RegisterOrcaServer is a generated function
func RegisterOrcaServer(s *grpc.Server, srv OrcaServer) {
	s.RegisterService(&_OrcaServiceDesc, srv)
}

func _OrcaEchoHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrcaServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Orca/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrcaServer).Echo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrcaServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Orca",
	HandlerType: (*OrcaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _OrcaEchoHandler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0xca, 0xd3, 0x40,
	0x14, 0x35, 0x5f, 0xf2, 0xa5, 0xf9, 0x6e, 0xea, 0x0f, 0xf7, 0x41, 0x86, 0x0f, 0x91, 0x1a, 0x44,
	0x8a, 0x48, 0x1f, 0xe2, 0x0e, 0x44, 0xdf, 0x04, 0x61, 0x70, 0x03, 0xe3, 0xcc, 0x50, 0x07, 0xd2,
	0x99, 0x98, 0x4c, 0x8b, 0x71, 0x25, 0xee, 0xc1, 0x3d, 0xb9, 0x16, 0x33, 0x77, 0x92, 0xb4, 0x05,
	0xf5, 0xed, 0x9c, 0x7b, 0x2e, 0x27, 0xe7, 0x9e, 0x69, 0x01, 0xb4, 0xfc, 0xea, 0x76, 0x6d, 0xe7,
	0xbc, 0xc3, 0x2c, 0xe0, 0xea, 0x1d, 0x64, 0x9f, 0xcd, 0x41, 0x23, 0x83, 0x55, 0xaf, 0xa5, 0xb3,
	0xaa, 0x67, 0xc9, 0x26, 0xd9, 0xa6, 0x7c, 0xa6, 0xb8, 0x81, 0xd2, 0x0a, 0xeb, 0x66, 0xf5, 0x86,
	0xd4, 0xcb, 0x51, 0xf5, 0x3b, 0x81, 0xe2, 0xa3, 0x93, 0xc2, 0x1b, 0x67, 0xf1, 0x29, 0xe4, 0xa6,
	0x15, 0x4a, 0x75, 0xe4, 0x73, 0xc7, 0x27, 0x86, 0xf7, 0x50, 0x34, 0xe3, 0x86, 0x3f, 0x2a, 0x4d,
	0x1e, 0x09, 0x5f, 0x38, 0x3e, 0x83, 0xbb, 0xc6, 0xd9, 0x7d, 0x14, 0x53, 0x12, 0xcf, 0x03, 0x44,
	0xc8, 0xa4, 0xf1, 0x03, 0xcb, 0xc8, 0x8f, 0x70, 0xf8, 0x4a, 0xeb, 0x7a, 0x2f, 0x1a, 0x76, 0x1b,
	0xbf, 0x12, 0x59, 0x38, 0x43, 0xba, 0xa3, 0xf5, 0xdd, 0xc0, 0x72, 0x12, 0x66, 0x8a, 0x15, 0xac,
	0x5d, 0xb7, 0x17, 0xd6, 0xfc, 0xa0, 0x9c, 0x6c, 0x45, 0xf2, 0xd5, 0x2c, 0xb8, 0x2a, 0x77, 0x10,
	0xc6, 0xb2, 0x22, 0xba, 0x46, 0x56, 0x7d, 0x87, 0xfc, 0xbd, 0x3e, 0x19, 0x49, 0x59, 0xac, 0x38,
	0xe8, 0xe9, 0x36, 0xc2, 0x17, 0x17, 0xdf, 0x5c, 0x5d, 0x7c, 0x76, 0x4b, 0x2f, 0xdd, 0xf0, 0xf5,
	0xd8, 0xc4, 0xd4, 0x16, 0xdd, 0x54, 0xd6, 0x8f, 0x76, 0xf4, 0x2e, 0x73, 0x87, 0x7c, 0xd1, 0xab,
	0x5f, 0x09, 0xac, 0xb8, 0xfe, 0x76, 0xd4, 0xbd, 0x0f, 0x0d, 0xf6, 0x01, 0x5a, 0xa9, 0xa7, 0x37,
	0x5a, 0x38, 0xbe, 0x84, 0xbc, 0xd7, 0x56, 0xe9, 0x98, 0xa1, 0xac, 0xd7, 0xd1, 0x31, 0xa6, 0xe6,
	0x93, 0x86, 0xcf, 0x21, 0x1b, 0x91, 0xa7, 0x3c, 0x65, 0x0d, 0x71, 0x27, 0x3c, 0x3f, 0xa7, 0x39,
	0x3e, 0x81, 0xd4, 0xfb, 0x86, 0x42, 0xa5, 0x3c, 0xc0, 0x70, 0x6f, 0x6b, 0xec, 0x9e, 0x5a, 0x4e,
	0x39, 0xe1, 0xd0, 0x71, 0x2b, 0x86, 0xc6, 0x09, 0xc5, 0x1e, 0x8f, 0xe3, 0x35, 0x9f, 0x69, 0xf5,
	0x33, 0x81, 0x5b, 0xae, 0xdb, 0x66, 0xf8, 0x6f, 0xd6, 0x2d, 0x14, 0x9d, 0x96, 0xda, 0x9c, 0xfe,
	0x91, 0x76, 0x51, 0xf1, 0xd5, 0xb2, 0xa9, 0xfe, 0x92, 0x79, 0xd1, 0xf0, 0x05, 0xd0, 0x8f, 0x79,
	0x6a, 0xf3, 0x61, 0xdc, 0x99, 0x6a, 0xe3, 0x24, 0xd5, 0x6f, 0x20, 0xfb, 0xd4, 0x49, 0x31, 0x16,
	0x95, 0x7d, 0x18, 0x39, 0x5e, 0x2f, 0xdd, 0x97, 0x33, 0x1d, 0xc3, 0x57, 0x0f, 0xbe, 0xe4, 0xf4,
	0x17, 0x79, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x41, 0x21, 0xb0, 0x4f, 0x30, 0x03, 0x00, 0x00,
}
