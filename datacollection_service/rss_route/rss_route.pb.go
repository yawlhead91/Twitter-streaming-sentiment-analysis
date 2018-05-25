// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rss_route.proto

package rss_route

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

type ParamsRss struct {
	Maxcount             int32    `protobuf:"varint,1,opt,name=Maxcount" json:"Maxcount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamsRss) Reset()         { *m = ParamsRss{} }
func (m *ParamsRss) String() string { return proto.CompactTextString(m) }
func (*ParamsRss) ProtoMessage()    {}
func (*ParamsRss) Descriptor() ([]byte, []int) {
	return fileDescriptor_rss_route_2ca1595d58dd0767, []int{0}
}
func (m *ParamsRss) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamsRss.Unmarshal(m, b)
}
func (m *ParamsRss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamsRss.Marshal(b, m, deterministic)
}
func (dst *ParamsRss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRss.Merge(dst, src)
}
func (m *ParamsRss) XXX_Size() int {
	return xxx_messageInfo_ParamsRss.Size(m)
}
func (m *ParamsRss) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRss.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRss proto.InternalMessageInfo

func (m *ParamsRss) GetMaxcount() int32 {
	if m != nil {
		return m.Maxcount
	}
	return 0
}

type FeedItem struct {
	CreatedAt            string   `protobuf:"bytes,1,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=Text" json:"Text,omitempty"`
	Score                int32    `protobuf:"varint,4,opt,name=Score" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItem) Reset()         { *m = FeedItem{} }
func (m *FeedItem) String() string { return proto.CompactTextString(m) }
func (*FeedItem) ProtoMessage()    {}
func (*FeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_rss_route_2ca1595d58dd0767, []int{1}
}
func (m *FeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItem.Unmarshal(m, b)
}
func (m *FeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItem.Marshal(b, m, deterministic)
}
func (dst *FeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItem.Merge(dst, src)
}
func (m *FeedItem) XXX_Size() int {
	return xxx_messageInfo_FeedItem.Size(m)
}
func (m *FeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItem proto.InternalMessageInfo

func (m *FeedItem) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FeedItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *FeedItem) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *FeedItem) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*ParamsRss)(nil), "ParamsRss")
	proto.RegisterType((*FeedItem)(nil), "FeedItem")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RssRouteClient is the client API for RssRoute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RssRouteClient interface {
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	GetRss(ctx context.Context, in *ParamsRss, opts ...grpc.CallOption) (RssRoute_GetRssClient, error)
}

type rssRouteClient struct {
	cc *grpc.ClientConn
}

func NewRssRouteClient(cc *grpc.ClientConn) RssRouteClient {
	return &rssRouteClient{cc}
}

func (c *rssRouteClient) GetRss(ctx context.Context, in *ParamsRss, opts ...grpc.CallOption) (RssRoute_GetRssClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RssRoute_serviceDesc.Streams[0], "/RssRoute/GetRss", opts...)
	if err != nil {
		return nil, err
	}
	x := &rssRouteGetRssClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RssRoute_GetRssClient interface {
	Recv() (*FeedItem, error)
	grpc.ClientStream
}

type rssRouteGetRssClient struct {
	grpc.ClientStream
}

func (x *rssRouteGetRssClient) Recv() (*FeedItem, error) {
	m := new(FeedItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RssRouteServer is the server API for RssRoute service.
type RssRouteServer interface {
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	GetRss(*ParamsRss, RssRoute_GetRssServer) error
}

func RegisterRssRouteServer(s *grpc.Server, srv RssRouteServer) {
	s.RegisterService(&_RssRoute_serviceDesc, srv)
}

func _RssRoute_GetRss_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ParamsRss)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RssRouteServer).GetRss(m, &rssRouteGetRssServer{stream})
}

type RssRoute_GetRssServer interface {
	Send(*FeedItem) error
	grpc.ServerStream
}

type rssRouteGetRssServer struct {
	grpc.ServerStream
}

func (x *rssRouteGetRssServer) Send(m *FeedItem) error {
	return x.ServerStream.SendMsg(m)
}

var _RssRoute_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RssRoute",
	HandlerType: (*RssRouteServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRss",
			Handler:       _RssRoute_GetRss_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rss_route.proto",
}

func init() { proto.RegisterFile("rss_route.proto", fileDescriptor_rss_route_2ca1595d58dd0767) }

var fileDescriptor_rss_route_2ca1595d58dd0767 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2a, 0x2e, 0x8e,
	0x2f, 0xca, 0x2f, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe7, 0xe2, 0x0c,
	0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x0e, 0x2a, 0x2e, 0x16, 0x92, 0xe2, 0xe2, 0xf0, 0x4d, 0xac, 0x48,
	0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf3, 0x95, 0x32, 0xb8,
	0x38, 0xdc, 0x52, 0x53, 0x53, 0x3c, 0x4b, 0x52, 0x73, 0x85, 0x64, 0xb8, 0x38, 0x9d, 0x8b, 0x52,
	0x13, 0x4b, 0x52, 0x53, 0x1c, 0x21, 0x0a, 0x39, 0x83, 0x10, 0x02, 0x42, 0x22, 0x5c, 0xac, 0x21,
	0x99, 0x25, 0x39, 0xa9, 0x12, 0x4c, 0x60, 0x19, 0x08, 0x47, 0x48, 0x88, 0x8b, 0x25, 0x24, 0xb5,
	0xa2, 0x44, 0x82, 0x19, 0x2c, 0x08, 0x66, 0x83, 0x54, 0x06, 0x27, 0xe7, 0x17, 0xa5, 0x4a, 0xb0,
	0x80, 0x2d, 0x83, 0x70, 0x8c, 0xf4, 0xb9, 0x38, 0x82, 0x8a, 0x8b, 0x83, 0x40, 0x8e, 0x14, 0x52,
	0xe6, 0x62, 0x73, 0x4f, 0x2d, 0x01, 0xb9, 0x8d, 0x4b, 0x0f, 0xee, 0x4e, 0x29, 0x4e, 0x3d, 0x98,
	0x53, 0x94, 0x18, 0x0c, 0x18, 0x93, 0xd8, 0xc0, 0x5e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x7c, 0xc3, 0xe3, 0xce, 0xdd, 0x00, 0x00, 0x00,
}
