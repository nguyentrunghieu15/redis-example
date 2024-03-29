// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: example/udemy/chatapplication/chat_grpc/proto_file.proto

package chat_proto

import (
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

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescGZIP(), []int{1}
}

func (x *ClientMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_example_udemy_chatapplication_chat_grpc_proto_file_proto protoreflect.FileDescriptor

var file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x75, 0x64, 0x65, 0x6d, 0x79, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x80, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x07,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x07,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x1a, 0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6e, 0x67, 0x68, 0x69, 0x65, 0x75,
	0x31, 0x35, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescOnce sync.Once
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescData = file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDesc
)

func file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescGZIP() []byte {
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescOnce.Do(func() {
		file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescData)
	})
	return file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDescData
}

var file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_udemy_chatapplication_chat_grpc_proto_file_proto_goTypes = []interface{}{
	(*Client)(nil),        // 0: Client
	(*ClientMessage)(nil), // 1: ClientMessage
}
var file_example_udemy_chatapplication_chat_grpc_proto_file_proto_depIdxs = []int32{
	0, // 0: ChatService.ReqJoinChat:input_type -> Client
	1, // 1: ChatService.ClientChat:input_type -> ClientMessage
	0, // 2: ChatService.ServerChat:input_type -> Client
	0, // 3: ChatService.ReqJoinChat:output_type -> Client
	0, // 4: ChatService.ClientChat:output_type -> Client
	1, // 5: ChatService.ServerChat:output_type -> ClientMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_udemy_chatapplication_chat_grpc_proto_file_proto_init() }
func file_example_udemy_chatapplication_chat_grpc_proto_file_proto_init() {
	if File_example_udemy_chatapplication_chat_grpc_proto_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
			RawDescriptor: file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_udemy_chatapplication_chat_grpc_proto_file_proto_goTypes,
		DependencyIndexes: file_example_udemy_chatapplication_chat_grpc_proto_file_proto_depIdxs,
		MessageInfos:      file_example_udemy_chatapplication_chat_grpc_proto_file_proto_msgTypes,
	}.Build()
	File_example_udemy_chatapplication_chat_grpc_proto_file_proto = out.File
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_rawDesc = nil
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_goTypes = nil
	file_example_udemy_chatapplication_chat_grpc_proto_file_proto_depIdxs = nil
}
