// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/my_service.proto

package proto

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

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_my_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_my_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_proto_my_service_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_my_service_proto protoreflect.FileDescriptor

var file_proto_my_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x31, 0x30, 0x31, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x66, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x59, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x31, 0x30, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x31, 0x30, 0x31, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x09, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x12, 0x01, 0x2f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_my_service_proto_rawDescOnce sync.Once
	file_proto_my_service_proto_rawDescData = file_proto_my_service_proto_rawDesc
)

func file_proto_my_service_proto_rawDescGZIP() []byte {
	file_proto_my_service_proto_rawDescOnce.Do(func() {
		file_proto_my_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_my_service_proto_rawDescData)
	})
	return file_proto_my_service_proto_rawDescData
}

var file_proto_my_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_my_service_proto_goTypes = []interface{}{
	(*StringMessage)(nil), // 0: grpc_gateway_101.v1.StringMessage
}
var file_proto_my_service_proto_depIdxs = []int32{
	0, // 0: grpc_gateway_101.v1.MyService.Echo:input_type -> grpc_gateway_101.v1.StringMessage
	0, // 1: grpc_gateway_101.v1.MyService.Echo:output_type -> grpc_gateway_101.v1.StringMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_my_service_proto_init() }
func file_proto_my_service_proto_init() {
	if File_proto_my_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_my_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
			RawDescriptor: file_proto_my_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_my_service_proto_goTypes,
		DependencyIndexes: file_proto_my_service_proto_depIdxs,
		MessageInfos:      file_proto_my_service_proto_msgTypes,
	}.Build()
	File_proto_my_service_proto = out.File
	file_proto_my_service_proto_rawDesc = nil
	file_proto_my_service_proto_goTypes = nil
	file_proto_my_service_proto_depIdxs = nil
}
