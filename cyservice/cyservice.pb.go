// Code generated by protoc-gen-go.
// source: cyservice.proto
// DO NOT EDIT!

/*
Package cyservice is a generated protocol buffer package.

It is generated from these files:
	cyservice.proto

It has these top-level messages:
	Fragment
	MetaData
	KeyValue
	Node
	Edge
	NodeAttribute
	EdgeAttribute
	NetworkAttribute
	AnonymousAspect
*/
package cyservice

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

type Fragment struct {
	// Types that are valid to be assigned to Element:
	//	*Fragment_Metadata
	//	*Fragment_Node
	//	*Fragment_Edge
	//	*Fragment_NodeAttribute
	//	*Fragment_EdgeAttribute
	//	*Fragment_NetworkAttribute
	//	*Fragment_Aspect
	Element isFragment_Element `protobuf_oneof:"element"`
}

func (m *Fragment) Reset()                    { *m = Fragment{} }
func (m *Fragment) String() string            { return proto.CompactTextString(m) }
func (*Fragment) ProtoMessage()               {}
func (*Fragment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isFragment_Element interface {
	isFragment_Element()
}

type Fragment_Metadata struct {
	Metadata *MetaData `protobuf:"bytes,1,opt,name=metadata,oneof"`
}
type Fragment_Node struct {
	Node *Node `protobuf:"bytes,2,opt,name=node,oneof"`
}
type Fragment_Edge struct {
	Edge *Edge `protobuf:"bytes,3,opt,name=edge,oneof"`
}
type Fragment_NodeAttribute struct {
	NodeAttribute *NodeAttribute `protobuf:"bytes,4,opt,name=nodeAttribute,oneof"`
}
type Fragment_EdgeAttribute struct {
	EdgeAttribute *EdgeAttribute `protobuf:"bytes,5,opt,name=edgeAttribute,oneof"`
}
type Fragment_NetworkAttribute struct {
	NetworkAttribute *NetworkAttribute `protobuf:"bytes,6,opt,name=networkAttribute,oneof"`
}
type Fragment_Aspect struct {
	Aspect *AnonymousAspect `protobuf:"bytes,7,opt,name=aspect,oneof"`
}

func (*Fragment_Metadata) isFragment_Element()         {}
func (*Fragment_Node) isFragment_Element()             {}
func (*Fragment_Edge) isFragment_Element()             {}
func (*Fragment_NodeAttribute) isFragment_Element()    {}
func (*Fragment_EdgeAttribute) isFragment_Element()    {}
func (*Fragment_NetworkAttribute) isFragment_Element() {}
func (*Fragment_Aspect) isFragment_Element()           {}

func (m *Fragment) GetElement() isFragment_Element {
	if m != nil {
		return m.Element
	}
	return nil
}

func (m *Fragment) GetMetadata() *MetaData {
	if x, ok := m.GetElement().(*Fragment_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Fragment) GetNode() *Node {
	if x, ok := m.GetElement().(*Fragment_Node); ok {
		return x.Node
	}
	return nil
}

func (m *Fragment) GetEdge() *Edge {
	if x, ok := m.GetElement().(*Fragment_Edge); ok {
		return x.Edge
	}
	return nil
}

func (m *Fragment) GetNodeAttribute() *NodeAttribute {
	if x, ok := m.GetElement().(*Fragment_NodeAttribute); ok {
		return x.NodeAttribute
	}
	return nil
}

func (m *Fragment) GetEdgeAttribute() *EdgeAttribute {
	if x, ok := m.GetElement().(*Fragment_EdgeAttribute); ok {
		return x.EdgeAttribute
	}
	return nil
}

func (m *Fragment) GetNetworkAttribute() *NetworkAttribute {
	if x, ok := m.GetElement().(*Fragment_NetworkAttribute); ok {
		return x.NetworkAttribute
	}
	return nil
}

func (m *Fragment) GetAspect() *AnonymousAspect {
	if x, ok := m.GetElement().(*Fragment_Aspect); ok {
		return x.Aspect
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Fragment) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Fragment_OneofMarshaler, _Fragment_OneofUnmarshaler, _Fragment_OneofSizer, []interface{}{
		(*Fragment_Metadata)(nil),
		(*Fragment_Node)(nil),
		(*Fragment_Edge)(nil),
		(*Fragment_NodeAttribute)(nil),
		(*Fragment_EdgeAttribute)(nil),
		(*Fragment_NetworkAttribute)(nil),
		(*Fragment_Aspect)(nil),
	}
}

func _Fragment_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Fragment)
	// element
	switch x := m.Element.(type) {
	case *Fragment_Metadata:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Metadata); err != nil {
			return err
		}
	case *Fragment_Node:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case *Fragment_Edge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Edge); err != nil {
			return err
		}
	case *Fragment_NodeAttribute:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeAttribute); err != nil {
			return err
		}
	case *Fragment_EdgeAttribute:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EdgeAttribute); err != nil {
			return err
		}
	case *Fragment_NetworkAttribute:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetworkAttribute); err != nil {
			return err
		}
	case *Fragment_Aspect:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aspect); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Fragment.Element has unexpected type %T", x)
	}
	return nil
}

func _Fragment_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Fragment)
	switch tag {
	case 1: // element.metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MetaData)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_Metadata{msg}
		return true, err
	case 2: // element.node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_Node{msg}
		return true, err
	case 3: // element.edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Edge)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_Edge{msg}
		return true, err
	case 4: // element.nodeAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NodeAttribute)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_NodeAttribute{msg}
		return true, err
	case 5: // element.edgeAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EdgeAttribute)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_EdgeAttribute{msg}
		return true, err
	case 6: // element.networkAttribute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NetworkAttribute)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_NetworkAttribute{msg}
		return true, err
	case 7: // element.aspect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AnonymousAspect)
		err := b.DecodeMessage(msg)
		m.Element = &Fragment_Aspect{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Fragment_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Fragment)
	// element
	switch x := m.Element.(type) {
	case *Fragment_Metadata:
		s := proto.Size(x.Metadata)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_Node:
		s := proto.Size(x.Node)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_Edge:
		s := proto.Size(x.Edge)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_NodeAttribute:
		s := proto.Size(x.NodeAttribute)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_EdgeAttribute:
		s := proto.Size(x.EdgeAttribute)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_NetworkAttribute:
		s := proto.Size(x.NetworkAttribute)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Fragment_Aspect:
		s := proto.Size(x.Aspect)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MetaData struct {
	Name             string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version          string      `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	IdCounter        int64       `protobuf:"varint,3,opt,name=idCounter" json:"idCounter,omitempty"`
	ElementCount     int64       `protobuf:"varint,4,opt,name=elementCount" json:"elementCount,omitempty"`
	ConsistencyGroup int64       `protobuf:"varint,5,opt,name=consistencyGroup" json:"consistencyGroup,omitempty"`
	Checksum         int64       `protobuf:"varint,6,opt,name=checksum" json:"checksum,omitempty"`
	Properties       []*KeyValue `protobuf:"bytes,7,rep,name=properties" json:"properties,omitempty"`
}

func (m *MetaData) Reset()                    { *m = MetaData{} }
func (m *MetaData) String() string            { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()               {}
func (*MetaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetaData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetaData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MetaData) GetIdCounter() int64 {
	if m != nil {
		return m.IdCounter
	}
	return 0
}

func (m *MetaData) GetElementCount() int64 {
	if m != nil {
		return m.ElementCount
	}
	return 0
}

func (m *MetaData) GetConsistencyGroup() int64 {
	if m != nil {
		return m.ConsistencyGroup
	}
	return 0
}

func (m *MetaData) GetChecksum() int64 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

func (m *MetaData) GetProperties() []*KeyValue {
	if m != nil {
		return m.Properties
	}
	return nil
}

type KeyValue struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Node struct {
	Id         int64  `protobuf:"varint,1,opt,name=id,json=@id" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Represents string `protobuf:"bytes,3,opt,name=represents,json=r" json:"represents,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetRepresents() string {
	if m != nil {
		return m.Represents
	}
	return ""
}

type Edge struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,json=@id" json:"id,omitempty"`
	SourceId    int64  `protobuf:"varint,2,opt,name=sourceId,json=s" json:"sourceId,omitempty"`
	TargetId    int64  `protobuf:"varint,3,opt,name=targetId,json=t" json:"targetId,omitempty"`
	Interaction string `protobuf:"bytes,4,opt,name=interaction,json=i" json:"interaction,omitempty"`
}

func (m *Edge) Reset()                    { *m = Edge{} }
func (m *Edge) String() string            { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()               {}
func (*Edge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Edge) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Edge) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Edge) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Edge) GetInteraction() string {
	if m != nil {
		return m.Interaction
	}
	return ""
}

type NodeAttribute struct {
	NodeId   int64  `protobuf:"varint,1,opt,name=nodeId,json=po" json:"nodeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *NodeAttribute) Reset()                    { *m = NodeAttribute{} }
func (m *NodeAttribute) String() string            { return proto.CompactTextString(m) }
func (*NodeAttribute) ProtoMessage()               {}
func (*NodeAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NodeAttribute) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NodeAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NodeAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type EdgeAttribute struct {
	EdgeId   int64  `protobuf:"varint,1,opt,name=edgeId,json=po" json:"edgeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *EdgeAttribute) Reset()                    { *m = EdgeAttribute{} }
func (m *EdgeAttribute) String() string            { return proto.CompactTextString(m) }
func (*EdgeAttribute) ProtoMessage()               {}
func (*EdgeAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EdgeAttribute) GetEdgeId() int64 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *EdgeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EdgeAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *EdgeAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EdgeAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type NetworkAttribute struct {
	EdgeId   int64  `protobuf:"varint,1,opt,name=edgeId,json=po" json:"edgeId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,json=n" json:"name,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,json=v" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,json=d" json:"type,omitempty"`
	SubnetId int64  `protobuf:"varint,5,opt,name=subnetId,json=s" json:"subnetId,omitempty"`
}

func (m *NetworkAttribute) Reset()                    { *m = NetworkAttribute{} }
func (m *NetworkAttribute) String() string            { return proto.CompactTextString(m) }
func (*NetworkAttribute) ProtoMessage()               {}
func (*NetworkAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NetworkAttribute) GetEdgeId() int64 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *NetworkAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NetworkAttribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkAttribute) GetSubnetId() int64 {
	if m != nil {
		return m.SubnetId
	}
	return 0
}

type AnonymousAspect struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Element []byte `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
}

func (m *AnonymousAspect) Reset()                    { *m = AnonymousAspect{} }
func (m *AnonymousAspect) String() string            { return proto.CompactTextString(m) }
func (*AnonymousAspect) ProtoMessage()               {}
func (*AnonymousAspect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AnonymousAspect) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnonymousAspect) GetElement() []byte {
	if m != nil {
		return m.Element
	}
	return nil
}

func init() {
	proto.RegisterType((*Fragment)(nil), "cyservice.Fragment")
	proto.RegisterType((*MetaData)(nil), "cyservice.MetaData")
	proto.RegisterType((*KeyValue)(nil), "cyservice.KeyValue")
	proto.RegisterType((*Node)(nil), "cyservice.Node")
	proto.RegisterType((*Edge)(nil), "cyservice.Edge")
	proto.RegisterType((*NodeAttribute)(nil), "cyservice.NodeAttribute")
	proto.RegisterType((*EdgeAttribute)(nil), "cyservice.EdgeAttribute")
	proto.RegisterType((*NetworkAttribute)(nil), "cyservice.NetworkAttribute")
	proto.RegisterType((*AnonymousAspect)(nil), "cyservice.AnonymousAspect")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CyService service

type CyServiceClient interface {
	StreamFragments(ctx context.Context, opts ...grpc.CallOption) (CyService_StreamFragmentsClient, error)
}

type cyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCyServiceClient(cc *grpc.ClientConn) CyServiceClient {
	return &cyServiceClient{cc}
}

func (c *cyServiceClient) StreamFragments(ctx context.Context, opts ...grpc.CallOption) (CyService_StreamFragmentsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CyService_serviceDesc.Streams[0], c.cc, "/cyservice.CyService/StreamFragments", opts...)
	if err != nil {
		return nil, err
	}
	x := &cyServiceStreamFragmentsClient{stream}
	return x, nil
}

type CyService_StreamFragmentsClient interface {
	Send(*Fragment) error
	Recv() (*Fragment, error)
	grpc.ClientStream
}

type cyServiceStreamFragmentsClient struct {
	grpc.ClientStream
}

func (x *cyServiceStreamFragmentsClient) Send(m *Fragment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cyServiceStreamFragmentsClient) Recv() (*Fragment, error) {
	m := new(Fragment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CyService service

type CyServiceServer interface {
	StreamFragments(CyService_StreamFragmentsServer) error
}

func RegisterCyServiceServer(s *grpc.Server, srv CyServiceServer) {
	s.RegisterService(&_CyService_serviceDesc, srv)
}

func _CyService_StreamFragments_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CyServiceServer).StreamFragments(&cyServiceStreamFragmentsServer{stream})
}

type CyService_StreamFragmentsServer interface {
	Send(*Fragment) error
	Recv() (*Fragment, error)
	grpc.ServerStream
}

type cyServiceStreamFragmentsServer struct {
	grpc.ServerStream
}

func (x *cyServiceStreamFragmentsServer) Send(m *Fragment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cyServiceStreamFragmentsServer) Recv() (*Fragment, error) {
	m := new(Fragment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cyservice.CyService",
	HandlerType: (*CyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamFragments",
			Handler:       _CyService_StreamFragments_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cyservice.proto",
}

func init() { proto.RegisterFile("cyservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x6b, 0x37, 0xb1, 0xa7, 0xad, 0x1c, 0x6d, 0x01, 0x59, 0x81, 0x43, 0x65, 0x09, 0xa9,
	0xe2, 0x50, 0x41, 0xdb, 0x7b, 0x1b, 0x4a, 0x21, 0x11, 0xa2, 0x87, 0xad, 0x84, 0xc4, 0x71, 0x63,
	0x8f, 0x82, 0xd5, 0x78, 0xd7, 0xda, 0x5d, 0x07, 0xf9, 0x5b, 0xf8, 0x50, 0xae, 0x68, 0x37, 0x76,
	0xe2, 0xc4, 0xbd, 0xf6, 0xe6, 0x7d, 0xf3, 0xe6, 0xcd, 0xec, 0xcc, 0x5b, 0x43, 0x94, 0xd6, 0x0a,
	0xe5, 0x2a, 0x4f, 0xf1, 0xa2, 0x94, 0x42, 0x0b, 0x12, 0x6e, 0x80, 0xe4, 0xaf, 0x07, 0xc1, 0x57,
	0xc9, 0x16, 0x05, 0x72, 0x4d, 0x3e, 0x41, 0x50, 0xa0, 0x66, 0x19, 0xd3, 0x2c, 0x76, 0xcf, 0xdc,
	0xf3, 0xa3, 0xcb, 0xd3, 0x8b, 0x6d, 0xee, 0x0f, 0xd4, 0xec, 0x0b, 0xd3, 0x6c, 0xea, 0xd0, 0x0d,
	0x8d, 0xbc, 0x07, 0x9f, 0x8b, 0x0c, 0xe3, 0x03, 0x4b, 0x8f, 0x3a, 0xf4, 0x07, 0x91, 0xe1, 0xd4,
	0xa1, 0x36, 0x6c, 0x68, 0x98, 0x2d, 0x30, 0xf6, 0x7a, 0xb4, 0xfb, 0x6c, 0x61, 0x69, 0x26, 0x4c,
	0x6e, 0xe1, 0xc4, 0xd0, 0x27, 0x5a, 0xcb, 0x7c, 0x5e, 0x69, 0x8c, 0x7d, 0xcb, 0x8f, 0xf7, 0x64,
	0x37, 0xf1, 0xa9, 0x43, 0x77, 0x13, 0x8c, 0x82, 0x51, 0xda, 0x2a, 0x1c, 0xf6, 0x14, 0xee, 0xbb,
	0x71, 0xa3, 0xb0, 0x93, 0x40, 0x66, 0x30, 0xe2, 0xa8, 0xff, 0x08, 0xf9, 0xb4, 0x15, 0x19, 0x58,
	0x91, 0xb7, 0xdd, 0x36, 0xf6, 0x28, 0x53, 0x87, 0xf6, 0xd2, 0xc8, 0x35, 0x0c, 0x98, 0x2a, 0x31,
	0xd5, 0xf1, 0xd0, 0x0a, 0x8c, 0x3b, 0x02, 0x13, 0x2e, 0x78, 0x5d, 0x88, 0x4a, 0x4d, 0x2c, 0x63,
	0xea, 0xd0, 0x86, 0xfb, 0x39, 0x84, 0x21, 0x2e, 0xd1, 0x2c, 0x24, 0xf9, 0xe7, 0x42, 0xd0, 0x8e,
	0x9d, 0x10, 0xf0, 0x39, 0x2b, 0xd0, 0x6e, 0x26, 0xa4, 0xf6, 0x9b, 0xc4, 0x30, 0x5c, 0xa1, 0x54,
	0xb9, 0xe0, 0x76, 0x03, 0x21, 0x6d, 0x8f, 0xe4, 0x1d, 0x84, 0x79, 0x76, 0x27, 0x2a, 0xae, 0x51,
	0xda, 0xb1, 0x7b, 0x74, 0x0b, 0x90, 0x04, 0x8e, 0x9b, 0x1a, 0x16, 0xb1, 0x73, 0xf6, 0xe8, 0x0e,
	0x46, 0x3e, 0xc0, 0x28, 0x15, 0x5c, 0xe5, 0x4a, 0x23, 0x4f, 0xeb, 0x6f, 0x52, 0x54, 0xa5, 0x9d,
	0xa6, 0x47, 0x7b, 0x38, 0x19, 0x43, 0x90, 0xfe, 0xc6, 0xf4, 0x49, 0x55, 0x85, 0x1d, 0x96, 0x47,
	0x37, 0x67, 0x72, 0x05, 0x50, 0x4a, 0x51, 0xa2, 0xd4, 0x39, 0xaa, 0x78, 0x78, 0xe6, 0xed, 0xf9,
	0xea, 0x3b, 0xd6, 0x3f, 0xd9, 0xb2, 0x42, 0xda, 0xa1, 0x25, 0xd7, 0x10, 0xb4, 0xf8, 0xb3, 0x17,
	0x7f, 0x05, 0x87, 0x2b, 0x13, 0x6c, 0xae, 0xbd, 0x3e, 0x24, 0x37, 0xe0, 0x1b, 0x7f, 0x90, 0x08,
	0x0e, 0xf2, 0xcc, 0xf2, 0x3d, 0xea, 0xdd, 0xe6, 0x19, 0x89, 0x1a, 0x89, 0x35, 0xdb, 0xe5, 0xe4,
	0x35, 0x80, 0xc4, 0x52, 0xa2, 0x42, 0xae, 0x95, 0x9d, 0x4f, 0x48, 0x5d, 0x99, 0xfc, 0x02, 0xdf,
	0xd8, 0xa3, 0x2f, 0x70, 0x0a, 0x81, 0x12, 0x95, 0x4c, 0x71, 0x96, 0x59, 0x11, 0x8f, 0xba, 0xca,
	0x80, 0x9a, 0xc9, 0x05, 0xea, 0x59, 0xd6, 0x8c, 0xd8, 0xd5, 0xe4, 0x0d, 0x1c, 0xe5, 0x66, 0xc6,
	0x2c, 0xd5, 0x66, 0x2d, 0xfe, 0x5a, 0x3a, 0x4f, 0x96, 0x70, 0xb2, 0xe3, 0x5d, 0x42, 0x60, 0x60,
	0xbc, 0x3b, 0x6b, 0xeb, 0x1c, 0x94, 0xa2, 0xdf, 0xe7, 0xa8, 0xbd, 0x67, 0xd3, 0xe2, 0xca, 0x50,
	0x74, 0x5d, 0x62, 0x2b, 0xbc, 0x6e, 0xad, 0x9a, 0x73, 0xdb, 0xc5, 0x61, 0xd3, 0x9a, 0xa9, 0xb6,
	0xe3, 0x73, 0x53, 0xcd, 0xf8, 0xfc, 0x65, 0xaa, 0x09, 0x18, 0xed, 0x3f, 0x88, 0x97, 0x2d, 0x78,
	0x03, 0xd1, 0xde, 0x03, 0x32, 0x2e, 0xb1, 0x89, 0x8d, 0x4b, 0xcc, 0xb7, 0x79, 0x1e, 0x8d, 0xa5,
	0x6d, 0xc9, 0x63, 0xda, 0x1e, 0x2f, 0x1f, 0x20, 0xbc, 0xab, 0x1f, 0xd7, 0x0e, 0x24, 0x13, 0x88,
	0x1e, 0xb5, 0x44, 0x56, 0xb4, 0x7f, 0x42, 0x45, 0xba, 0x06, 0x6d, 0xd1, 0xf1, 0x73, 0x60, 0xe2,
	0x9c, 0xbb, 0x1f, 0xdd, 0xf9, 0xc0, 0xfe, 0x59, 0xaf, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa6,
	0xf7, 0xbf, 0xc8, 0x6c, 0x05, 0x00, 0x00,
}