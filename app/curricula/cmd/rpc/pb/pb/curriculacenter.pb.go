// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.11
// source: curriculacenter.proto

package pb

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

type AddCurriculaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurriculaId   uint32 `protobuf:"varint,1,opt,name=CurriculaId,proto3" json:"CurriculaId,omitempty"`
	CurriculaName string `protobuf:"bytes,2,opt,name=CurriculaName,proto3" json:"CurriculaName,omitempty"`
	Teacher       string `protobuf:"bytes,3,opt,name=Teacher,proto3" json:"Teacher,omitempty"`
	Type          uint32 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *AddCurriculaRequest) Reset() {
	*x = AddCurriculaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curriculacenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCurriculaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCurriculaRequest) ProtoMessage() {}

func (x *AddCurriculaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_curriculacenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCurriculaRequest.ProtoReflect.Descriptor instead.
func (*AddCurriculaRequest) Descriptor() ([]byte, []int) {
	return file_curriculacenter_proto_rawDescGZIP(), []int{0}
}

func (x *AddCurriculaRequest) GetCurriculaId() uint32 {
	if x != nil {
		return x.CurriculaId
	}
	return 0
}

func (x *AddCurriculaRequest) GetCurriculaName() string {
	if x != nil {
		return x.CurriculaName
	}
	return ""
}

func (x *AddCurriculaRequest) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *AddCurriculaRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type AddCurriculaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCurriculaResponse) Reset() {
	*x = AddCurriculaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curriculacenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCurriculaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCurriculaResponse) ProtoMessage() {}

func (x *AddCurriculaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_curriculacenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCurriculaResponse.ProtoReflect.Descriptor instead.
func (*AddCurriculaResponse) Descriptor() ([]byte, []int) {
	return file_curriculacenter_proto_rawDescGZIP(), []int{1}
}

var File_curriculacenter_proto protoreflect.FileDescriptor

var file_curriculacenter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x8b, 0x01, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x43, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x69, 0x63,
	0x75, 0x6c, 0x61, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75,
	0x6c, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x75,
	0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x43, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x54, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x43, 0x75, 0x72, 0x72, 0x69,
	0x63, 0x75, 0x6c, 0x61, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x72,
	0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_curriculacenter_proto_rawDescOnce sync.Once
	file_curriculacenter_proto_rawDescData = file_curriculacenter_proto_rawDesc
)

func file_curriculacenter_proto_rawDescGZIP() []byte {
	file_curriculacenter_proto_rawDescOnce.Do(func() {
		file_curriculacenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_curriculacenter_proto_rawDescData)
	})
	return file_curriculacenter_proto_rawDescData
}

var file_curriculacenter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_curriculacenter_proto_goTypes = []interface{}{
	(*AddCurriculaRequest)(nil),  // 0: pb.AddCurriculaRequest
	(*AddCurriculaResponse)(nil), // 1: pb.AddCurriculaResponse
}
var file_curriculacenter_proto_depIdxs = []int32{
	0, // 0: pb.curriculacenter.addCurricula:input_type -> pb.AddCurriculaRequest
	1, // 1: pb.curriculacenter.addCurricula:output_type -> pb.AddCurriculaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_curriculacenter_proto_init() }
func file_curriculacenter_proto_init() {
	if File_curriculacenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_curriculacenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCurriculaRequest); i {
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
		file_curriculacenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCurriculaResponse); i {
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
			RawDescriptor: file_curriculacenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_curriculacenter_proto_goTypes,
		DependencyIndexes: file_curriculacenter_proto_depIdxs,
		MessageInfos:      file_curriculacenter_proto_msgTypes,
	}.Build()
	File_curriculacenter_proto = out.File
	file_curriculacenter_proto_rawDesc = nil
	file_curriculacenter_proto_goTypes = nil
	file_curriculacenter_proto_depIdxs = nil
}