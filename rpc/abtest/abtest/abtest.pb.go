// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: abtest.proto

package abtest

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

type AbReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AbReq) Reset() {
	*x = AbReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbReq) ProtoMessage() {}

func (x *AbReq) ProtoReflect() protoreflect.Message {
	mi := &file_abtest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbReq.ProtoReflect.Descriptor instead.
func (*AbReq) Descriptor() ([]byte, []int) {
	return file_abtest_proto_rawDescGZIP(), []int{0}
}

type AbRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan     string `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	UserCase string `protobuf:"bytes,2,opt,name=userCase,proto3" json:"userCase,omitempty"`
}

func (x *AbRsp) Reset() {
	*x = AbRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbRsp) ProtoMessage() {}

func (x *AbRsp) ProtoReflect() protoreflect.Message {
	mi := &file_abtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbRsp.ProtoReflect.Descriptor instead.
func (*AbRsp) Descriptor() ([]byte, []int) {
	return file_abtest_proto_rawDescGZIP(), []int{1}
}

func (x *AbRsp) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *AbRsp) GetUserCase() string {
	if x != nil {
		return x.UserCase
	}
	return ""
}

var File_abtest_proto protoreflect.FileDescriptor

var file_abtest_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x61, 0x62, 0x52, 0x65, 0x71, 0x22,
	0x37, 0x0a, 0x05, 0x61, 0x62, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x61, 0x73, 0x65, 0x32, 0x32, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x63,
	0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x12, 0x0d,
	0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x62, 0x52, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abtest_proto_rawDescOnce sync.Once
	file_abtest_proto_rawDescData = file_abtest_proto_rawDesc
)

func file_abtest_proto_rawDescGZIP() []byte {
	file_abtest_proto_rawDescOnce.Do(func() {
		file_abtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_abtest_proto_rawDescData)
	})
	return file_abtest_proto_rawDescData
}

var file_abtest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_abtest_proto_goTypes = []interface{}{
	(*AbReq)(nil), // 0: abtest.abReq
	(*AbRsp)(nil), // 1: abtest.abRsp
}
var file_abtest_proto_depIdxs = []int32{
	0, // 0: abtest.usecase.usecase:input_type -> abtest.abReq
	1, // 1: abtest.usecase.usecase:output_type -> abtest.abRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_abtest_proto_init() }
func file_abtest_proto_init() {
	if File_abtest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_abtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbReq); i {
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
		file_abtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbRsp); i {
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
			RawDescriptor: file_abtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_abtest_proto_goTypes,
		DependencyIndexes: file_abtest_proto_depIdxs,
		MessageInfos:      file_abtest_proto_msgTypes,
	}.Build()
	File_abtest_proto = out.File
	file_abtest_proto_rawDesc = nil
	file_abtest_proto_goTypes = nil
	file_abtest_proto_depIdxs = nil
}
