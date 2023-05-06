// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: hello/api/hello/v1/hello.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetHelloRequest) Reset() {
	*x = GetHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_api_hello_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloRequest) ProtoMessage() {}

func (x *GetHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_api_hello_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloRequest.ProtoReflect.Descriptor instead.
func (*GetHelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_api_hello_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *GetHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetHelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetHelloReply) Reset() {
	*x = GetHelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_api_hello_v1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloReply) ProtoMessage() {}

func (x *GetHelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_api_hello_v1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloReply.ProtoReflect.Descriptor instead.
func (*GetHelloReply) Descriptor() ([]byte, []int) {
	return file_hello_api_hello_v1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *GetHelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PostHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PostHelloRequest) Reset() {
	*x = PostHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_api_hello_v1_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHelloRequest) ProtoMessage() {}

func (x *PostHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_api_hello_v1_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHelloRequest.ProtoReflect.Descriptor instead.
func (*PostHelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_api_hello_v1_hello_proto_rawDescGZIP(), []int{2}
}

func (x *PostHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PostHelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PostHelloReply) Reset() {
	*x = PostHelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_api_hello_v1_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostHelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHelloReply) ProtoMessage() {}

func (x *PostHelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_api_hello_v1_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHelloReply.ProtoReflect.Descriptor instead.
func (*PostHelloReply) Descriptor() ([]byte, []int) {
	return file_hello_api_hello_v1_hello_proto_rawDescGZIP(), []int{3}
}

func (x *PostHelloReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_hello_api_hello_v1_hello_proto protoreflect.FileDescriptor

var file_hello_api_hello_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x66, 0x7a, 0x70, 0x61, 0x6e, 0x78, 0x69, 0x2e, 0x67, 0x67, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a,
	0x10, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2c, 0x2e, 0x66, 0x7a, 0x70, 0x61,
	0x6e, 0x78, 0x69, 0x2e, 0x67, 0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x7a, 0x70, 0x61, 0x6e, 0x78,
	0x69, 0x2e, 0x67, 0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x18, 0x12,
	0x16, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7d, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2d, 0x2e, 0x66, 0x7a, 0x70, 0x61, 0x6e, 0x78, 0x69, 0x2e, 0x67,
	0x67, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x66, 0x7a, 0x70, 0x61, 0x6e, 0x78, 0x69, 0x2e, 0x67, 0x67,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x7a, 0x70, 0x61, 0x6e, 0x78, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x69, 0x6e, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hello_api_hello_v1_hello_proto_rawDescOnce sync.Once
	file_hello_api_hello_v1_hello_proto_rawDescData = file_hello_api_hello_v1_hello_proto_rawDesc
)

func file_hello_api_hello_v1_hello_proto_rawDescGZIP() []byte {
	file_hello_api_hello_v1_hello_proto_rawDescOnce.Do(func() {
		file_hello_api_hello_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_api_hello_v1_hello_proto_rawDescData)
	})
	return file_hello_api_hello_v1_hello_proto_rawDescData
}

var file_hello_api_hello_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hello_api_hello_v1_hello_proto_goTypes = []interface{}{
	(*GetHelloRequest)(nil),  // 0: fzpanxi.gg.example.hello.v1.GetHelloRequest
	(*GetHelloReply)(nil),    // 1: fzpanxi.gg.example.hello.v1.GetHelloReply
	(*PostHelloRequest)(nil), // 2: fzpanxi.gg.example.hello.v1.PostHelloRequest
	(*PostHelloReply)(nil),   // 3: fzpanxi.gg.example.hello.v1.PostHelloReply
}
var file_hello_api_hello_v1_hello_proto_depIdxs = []int32{
	0, // 0: fzpanxi.gg.example.hello.v1.HelloService.GetHello:input_type -> fzpanxi.gg.example.hello.v1.GetHelloRequest
	2, // 1: fzpanxi.gg.example.hello.v1.HelloService.PostHello:input_type -> fzpanxi.gg.example.hello.v1.PostHelloRequest
	1, // 2: fzpanxi.gg.example.hello.v1.HelloService.GetHello:output_type -> fzpanxi.gg.example.hello.v1.GetHelloReply
	3, // 3: fzpanxi.gg.example.hello.v1.HelloService.PostHello:output_type -> fzpanxi.gg.example.hello.v1.PostHelloReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_api_hello_v1_hello_proto_init() }
func file_hello_api_hello_v1_hello_proto_init() {
	if File_hello_api_hello_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_api_hello_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHelloRequest); i {
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
		file_hello_api_hello_v1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHelloReply); i {
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
		file_hello_api_hello_v1_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostHelloRequest); i {
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
		file_hello_api_hello_v1_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostHelloReply); i {
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
			RawDescriptor: file_hello_api_hello_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_api_hello_v1_hello_proto_goTypes,
		DependencyIndexes: file_hello_api_hello_v1_hello_proto_depIdxs,
		MessageInfos:      file_hello_api_hello_v1_hello_proto_msgTypes,
	}.Build()
	File_hello_api_hello_v1_hello_proto = out.File
	file_hello_api_hello_v1_hello_proto_rawDesc = nil
	file_hello_api_hello_v1_hello_proto_goTypes = nil
	file_hello_api_hello_v1_hello_proto_depIdxs = nil
}
