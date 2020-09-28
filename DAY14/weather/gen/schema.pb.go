// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package weather

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

type QueryResponse_Weather int32

const (
	QueryResponse_SUNNY  QueryResponse_Weather = 0
	QueryResponse_CLOUDY QueryResponse_Weather = 1
)

var QueryResponse_Weather_name = map[int32]string{
	0: "SUNNY",
	1: "CLOUDY",
}

var QueryResponse_Weather_value = map[string]int32{
	"SUNNY":  0,
	"CLOUDY": 1,
}

func (x QueryResponse_Weather) String() string {
	return proto.EnumName(QueryResponse_Weather_name, int32(x))
}

func (QueryResponse_Weather) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1, 0}
}

type QueryRequest struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type QueryResponse struct {
	Location             string                `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Weather              QueryResponse_Weather `protobuf:"varint,2,opt,name=weather,proto3,enum=weather.QueryResponse_Weather" json:"weather,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *QueryResponse) GetWeather() QueryResponse_Weather {
	if m != nil {
		return m.Weather
	}
	return QueryResponse_SUNNY
}

func init() {
	proto.RegisterEnum("weather.QueryResponse_Weather", QueryResponse_Weather_name, QueryResponse_Weather_value)
	proto.RegisterType((*QueryRequest)(nil), "weather.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "weather.QueryResponse")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4f, 0x4d, 0x2c, 0xc9, 0x48,
	0x2d, 0x52, 0xd2, 0xe2, 0xe2, 0x09, 0x2c, 0x4d, 0x2d, 0xaa, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0xc9, 0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0xda, 0x19, 0xb9, 0x78, 0xa1, 0x8a, 0x8b, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0xf1, 0xa9, 0x16, 0xb2, 0xe0, 0x82, 0x59, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x67, 0x24, 0xa7, 0x07, 0xe5, 0xeb, 0xa1, 0x18, 0xa2, 0x17, 0x0e, 0x11, 0x0d, 0x82, 0xbb, 0x49,
	0x81, 0x8b, 0x1d, 0x2a, 0x26, 0xc4, 0xc9, 0xc5, 0x1a, 0x1c, 0xea, 0xe7, 0x17, 0x29, 0xc0, 0x20,
	0xc4, 0xc5, 0xc5, 0xe6, 0xec, 0xe3, 0x1f, 0xea, 0x12, 0x29, 0xc0, 0x68, 0xe4, 0x8e, 0x50, 0x61,
	0xc3, 0xc5, 0x0a, 0x36, 0x4e, 0x48, 0x14, 0xdd, 0x78, 0xb0, 0x87, 0xa4, 0xc4, 0xb0, 0xdb, 0xaa,
	0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x0e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xec, 0xeb, 0xfa, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherClient interface {
	Query(ctx context.Context, opts ...grpc.CallOption) (Weather_QueryClient, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) Query(ctx context.Context, opts ...grpc.CallOption) (Weather_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Weather_serviceDesc.Streams[0], "/weather.Weather/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherQueryClient{stream}
	return x, nil
}

type Weather_QueryClient interface {
	Send(*QueryRequest) error
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type weatherQueryClient struct {
	grpc.ClientStream
}

func (x *weatherQueryClient) Send(m *QueryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *weatherQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServer is the server API for Weather service.
type WeatherServer interface {
	Query(Weather_QueryServer) error
}

// UnimplementedWeatherServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (*UnimplementedWeatherServer) Query(srv Weather_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterWeatherServer(s *grpc.Server, srv WeatherServer) {
	s.RegisterService(&_Weather_serviceDesc, srv)
}

func _Weather_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WeatherServer).Query(&weatherQueryServer{stream})
}

type Weather_QueryServer interface {
	Send(*QueryResponse) error
	Recv() (*QueryRequest, error)
	grpc.ServerStream
}

type weatherQueryServer struct {
	grpc.ServerStream
}

func (x *weatherQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *weatherQueryServer) Recv() (*QueryRequest, error) {
	m := new(QueryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Weather_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weather.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _Weather_Query_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "schema.proto",
}
