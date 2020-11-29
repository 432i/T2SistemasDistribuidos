// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: serverclidn.proto

package serverclidn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreLibro string `protobuf:"bytes,1,opt,name=nombreLibro,proto3" json:"nombreLibro,omitempty"`
	TotalPartes string `protobuf:"bytes,2,opt,name=totalPartes,proto3" json:"totalPartes,omitempty"`
	Parte       string `protobuf:"bytes,3,opt,name=parte,proto3" json:"parte,omitempty"`
	Datos       []byte `protobuf:"bytes,4,opt,name=datos,proto3" json:"datos,omitempty"`
	Algoritmo   string `protobuf:"bytes,5,opt,name=algoritmo,proto3" json:"algoritmo,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverclidn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_serverclidn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_serverclidn_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

func (x *Chunk) GetTotalPartes() string {
	if x != nil {
		return x.TotalPartes
	}
	return ""
}

func (x *Chunk) GetParte() string {
	if x != nil {
		return x.Parte
	}
	return ""
}

func (x *Chunk) GetDatos() []byte {
	if x != nil {
		return x.Datos
	}
	return nil
}

func (x *Chunk) GetAlgoritmo() string {
	if x != nil {
		return x.Algoritmo
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serverclidn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_serverclidn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_serverclidn_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_serverclidn_proto protoreflect.FileDescriptor

var file_serverclidn_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e,
	0x22, 0x95, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x72, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x64, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x6d, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x81, 0x04, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74,
	0x43, 0x6c, 0x69, 0x44, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x61, 0x44,
	0x4e, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c,
	0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x41, 0x0a, 0x11, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x63, 0x69,
	0x6f, 0x6e, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69,
	0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x72, 0x4c, 0x6f,
	0x67, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x44, 0x4e, 0x12,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64,
	0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x70,
	0x65, 0x64, 0x69, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x65, 0x64, 0x69, 0x72, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63,
	0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x65,
	0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63,
	0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64,
	0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x6c, 0x69, 0x64, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serverclidn_proto_rawDescOnce sync.Once
	file_serverclidn_proto_rawDescData = file_serverclidn_proto_rawDesc
)

func file_serverclidn_proto_rawDescGZIP() []byte {
	file_serverclidn_proto_rawDescOnce.Do(func() {
		file_serverclidn_proto_rawDescData = protoimpl.X.CompressGZIP(file_serverclidn_proto_rawDescData)
	})
	return file_serverclidn_proto_rawDescData
}

var file_serverclidn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serverclidn_proto_goTypes = []interface{}{
	(*Chunk)(nil),   // 0: serverclidn.Chunk
	(*Message)(nil), // 1: serverclidn.Message
}
var file_serverclidn_proto_depIdxs = []int32{
	0, // 0: serverclidn.ChatCliDn.ChunkaDN:input_type -> serverclidn.Chunk
	1, // 1: serverclidn.ChatCliDn.ChunksDirecciones:input_type -> serverclidn.Message
	1, // 2: serverclidn.ChatCliDn.escribirLog:input_type -> serverclidn.Message
	0, // 3: serverclidn.ChatCliDn.ChunkEntreDN:input_type -> serverclidn.Chunk
	1, // 4: serverclidn.ChatCliDn.pedirChunk:input_type -> serverclidn.Message
	1, // 5: serverclidn.ChatCliDn.pedirCatalogo:input_type -> serverclidn.Message
	1, // 6: serverclidn.ChatCliDn.EnviarPeticion:input_type -> serverclidn.Message
	1, // 7: serverclidn.ChatCliDn.propuestaCentralizado:input_type -> serverclidn.Message
	1, // 8: serverclidn.ChatCliDn.ChunkaDN:output_type -> serverclidn.Message
	1, // 9: serverclidn.ChatCliDn.ChunksDirecciones:output_type -> serverclidn.Message
	1, // 10: serverclidn.ChatCliDn.escribirLog:output_type -> serverclidn.Message
	1, // 11: serverclidn.ChatCliDn.ChunkEntreDN:output_type -> serverclidn.Message
	0, // 12: serverclidn.ChatCliDn.pedirChunk:output_type -> serverclidn.Chunk
	1, // 13: serverclidn.ChatCliDn.pedirCatalogo:output_type -> serverclidn.Message
	1, // 14: serverclidn.ChatCliDn.EnviarPeticion:output_type -> serverclidn.Message
	1, // 15: serverclidn.ChatCliDn.propuestaCentralizado:output_type -> serverclidn.Message
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serverclidn_proto_init() }
func file_serverclidn_proto_init() {
	if File_serverclidn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serverclidn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_serverclidn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_serverclidn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serverclidn_proto_goTypes,
		DependencyIndexes: file_serverclidn_proto_depIdxs,
		MessageInfos:      file_serverclidn_proto_msgTypes,
	}.Build()
	File_serverclidn_proto = out.File
	file_serverclidn_proto_rawDesc = nil
	file_serverclidn_proto_goTypes = nil
	file_serverclidn_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatCliDnClient is the client API for ChatCliDn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatCliDnClient interface {
	ChunkaDN(ctx context.Context, opts ...grpc.CallOption) (ChatCliDn_ChunkaDNClient, error)
	ChunksDirecciones(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EscribirLog(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ChunkEntreDN(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error)
	PedirChunk(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Chunk, error)
	PedirCatalogo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EnviarPeticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	PropuestaCentralizado(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatCliDnClient struct {
	cc grpc.ClientConnInterface
}

func NewChatCliDnClient(cc grpc.ClientConnInterface) ChatCliDnClient {
	return &chatCliDnClient{cc}
}

func (c *chatCliDnClient) ChunkaDN(ctx context.Context, opts ...grpc.CallOption) (ChatCliDn_ChunkaDNClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatCliDn_serviceDesc.Streams[0], "/serverclidn.ChatCliDn/ChunkaDN", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatCliDnChunkaDNClient{stream}
	return x, nil
}

type ChatCliDn_ChunkaDNClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type chatCliDnChunkaDNClient struct {
	grpc.ClientStream
}

func (x *chatCliDnChunkaDNClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatCliDnChunkaDNClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatCliDnClient) ChunksDirecciones(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/ChunksDirecciones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) EscribirLog(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/escribirLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) ChunkEntreDN(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/ChunkEntreDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) PedirChunk(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/pedirChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) PedirCatalogo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/pedirCatalogo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) EnviarPeticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/EnviarPeticion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCliDnClient) PropuestaCentralizado(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/serverclidn.ChatCliDn/propuestaCentralizado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatCliDnServer is the server API for ChatCliDn service.
type ChatCliDnServer interface {
	ChunkaDN(ChatCliDn_ChunkaDNServer) error
	ChunksDirecciones(context.Context, *Message) (*Message, error)
	EscribirLog(context.Context, *Message) (*Message, error)
	ChunkEntreDN(context.Context, *Chunk) (*Message, error)
	PedirChunk(context.Context, *Message) (*Chunk, error)
	PedirCatalogo(context.Context, *Message) (*Message, error)
	EnviarPeticion(context.Context, *Message) (*Message, error)
	PropuestaCentralizado(context.Context, *Message) (*Message, error)
}

// UnimplementedChatCliDnServer can be embedded to have forward compatible implementations.
type UnimplementedChatCliDnServer struct {
}

func (*UnimplementedChatCliDnServer) ChunkaDN(ChatCliDn_ChunkaDNServer) error {
	return status.Errorf(codes.Unimplemented, "method ChunkaDN not implemented")
}
func (*UnimplementedChatCliDnServer) ChunksDirecciones(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChunksDirecciones not implemented")
}
func (*UnimplementedChatCliDnServer) EscribirLog(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscribirLog not implemented")
}
func (*UnimplementedChatCliDnServer) ChunkEntreDN(context.Context, *Chunk) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChunkEntreDN not implemented")
}
func (*UnimplementedChatCliDnServer) PedirChunk(context.Context, *Message) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirChunk not implemented")
}
func (*UnimplementedChatCliDnServer) PedirCatalogo(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirCatalogo not implemented")
}
func (*UnimplementedChatCliDnServer) EnviarPeticion(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPeticion not implemented")
}
func (*UnimplementedChatCliDnServer) PropuestaCentralizado(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropuestaCentralizado not implemented")
}

func RegisterChatCliDnServer(s *grpc.Server, srv ChatCliDnServer) {
	s.RegisterService(&_ChatCliDn_serviceDesc, srv)
}

func _ChatCliDn_ChunkaDN_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatCliDnServer).ChunkaDN(&chatCliDnChunkaDNServer{stream})
}

type ChatCliDn_ChunkaDNServer interface {
	SendAndClose(*Message) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type chatCliDnChunkaDNServer struct {
	grpc.ServerStream
}

func (x *chatCliDnChunkaDNServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatCliDnChunkaDNServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatCliDn_ChunksDirecciones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).ChunksDirecciones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/ChunksDirecciones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).ChunksDirecciones(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_EscribirLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).EscribirLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/EscribirLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).EscribirLog(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_ChunkEntreDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).ChunkEntreDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/ChunkEntreDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).ChunkEntreDN(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_PedirChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).PedirChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/PedirChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).PedirChunk(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_PedirCatalogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).PedirCatalogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/PedirCatalogo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).PedirCatalogo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_EnviarPeticion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).EnviarPeticion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/EnviarPeticion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).EnviarPeticion(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCliDn_PropuestaCentralizado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCliDnServer).PropuestaCentralizado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverclidn.ChatCliDn/PropuestaCentralizado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCliDnServer).PropuestaCentralizado(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatCliDn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverclidn.ChatCliDn",
	HandlerType: (*ChatCliDnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChunksDirecciones",
			Handler:    _ChatCliDn_ChunksDirecciones_Handler,
		},
		{
			MethodName: "escribirLog",
			Handler:    _ChatCliDn_EscribirLog_Handler,
		},
		{
			MethodName: "ChunkEntreDN",
			Handler:    _ChatCliDn_ChunkEntreDN_Handler,
		},
		{
			MethodName: "pedirChunk",
			Handler:    _ChatCliDn_PedirChunk_Handler,
		},
		{
			MethodName: "pedirCatalogo",
			Handler:    _ChatCliDn_PedirCatalogo_Handler,
		},
		{
			MethodName: "EnviarPeticion",
			Handler:    _ChatCliDn_EnviarPeticion_Handler,
		},
		{
			MethodName: "propuestaCentralizado",
			Handler:    _ChatCliDn_PropuestaCentralizado_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChunkaDN",
			Handler:       _ChatCliDn_ChunkaDN_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "serverclidn.proto",
}
