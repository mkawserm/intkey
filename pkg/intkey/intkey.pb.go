// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: intkey/intkey.proto

package intkey

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

type IntKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IntKey) Reset() {
	*x = IntKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_intkey_intkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntKey) ProtoMessage() {}

func (x *IntKey) ProtoReflect() protoreflect.Message {
	mi := &file_intkey_intkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntKey.ProtoReflect.Descriptor instead.
func (*IntKey) Descriptor() ([]byte, []int) {
	return file_intkey_intkey_proto_rawDescGZIP(), []int{0}
}

func (x *IntKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IntKey) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_intkey_intkey_proto protoreflect.FileDescriptor

var file_intkey_intkey_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x22, 0x30, 0x0a,
	0x06, 0x49, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0xc1, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x50, 0x43, 0x12, 0x2a, 0x0a,
	0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79,
	0x2e, 0x49, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79,
	0x2e, 0x49, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x1a, 0x0e, 0x2e, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x6b, 0x61, 0x77, 0x73, 0x65, 0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x6b, 0x65,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x6b, 0x65, 0x79, 0x3b, 0x69, 0x6e, 0x74,
	0x6b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_intkey_intkey_proto_rawDescOnce sync.Once
	file_intkey_intkey_proto_rawDescData = file_intkey_intkey_proto_rawDesc
)

func file_intkey_intkey_proto_rawDescGZIP() []byte {
	file_intkey_intkey_proto_rawDescOnce.Do(func() {
		file_intkey_intkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_intkey_intkey_proto_rawDescData)
	})
	return file_intkey_intkey_proto_rawDescData
}

var file_intkey_intkey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_intkey_intkey_proto_goTypes = []interface{}{
	(*IntKey)(nil), // 0: intkey.IntKey
}
var file_intkey_intkey_proto_depIdxs = []int32{
	0, // 0: intkey.IntKeyRPC.Insert:input_type -> intkey.IntKey
	0, // 1: intkey.IntKeyRPC.Delete:input_type -> intkey.IntKey
	0, // 2: intkey.IntKeyRPC.Increment:input_type -> intkey.IntKey
	0, // 3: intkey.IntKeyRPC.Decrement:input_type -> intkey.IntKey
	0, // 4: intkey.IntKeyRPC.Insert:output_type -> intkey.IntKey
	0, // 5: intkey.IntKeyRPC.Delete:output_type -> intkey.IntKey
	0, // 6: intkey.IntKeyRPC.Increment:output_type -> intkey.IntKey
	0, // 7: intkey.IntKeyRPC.Decrement:output_type -> intkey.IntKey
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_intkey_intkey_proto_init() }
func file_intkey_intkey_proto_init() {
	if File_intkey_intkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_intkey_intkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntKey); i {
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
			RawDescriptor: file_intkey_intkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_intkey_intkey_proto_goTypes,
		DependencyIndexes: file_intkey_intkey_proto_depIdxs,
		MessageInfos:      file_intkey_intkey_proto_msgTypes,
	}.Build()
	File_intkey_intkey_proto = out.File
	file_intkey_intkey_proto_rawDesc = nil
	file_intkey_intkey_proto_goTypes = nil
	file_intkey_intkey_proto_depIdxs = nil
}
