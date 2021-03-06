// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ozonmp/bss_office_facade/v1/bss_office_facade.proto

package bss_office_facade

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OfficePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OfficePayload) Reset() {
	*x = OfficePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficePayload) ProtoMessage() {}

func (x *OfficePayload) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficePayload.ProtoReflect.Descriptor instead.
func (*OfficePayload) Descriptor() ([]byte, []int) {
	return file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescGZIP(), []int{0}
}

func (x *OfficePayload) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OfficePayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfficePayload) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OfficeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OfficeId uint64                 `protobuf:"varint,2,opt,name=office_id,json=officeId,proto3" json:"office_id,omitempty"`
	Status   uint64                 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Type     string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Created  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	Payload  *OfficePayload         `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *OfficeEvent) Reset() {
	*x = OfficeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeEvent) ProtoMessage() {}

func (x *OfficeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeEvent.ProtoReflect.Descriptor instead.
func (*OfficeEvent) Descriptor() ([]byte, []int) {
	return file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescGZIP(), []int{1}
}

func (x *OfficeEvent) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OfficeEvent) GetOfficeId() uint64 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *OfficeEvent) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OfficeEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OfficeEvent) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *OfficeEvent) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *OfficeEvent) GetPayload() *OfficePayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_ozonmp_bss_office_facade_v1_bss_office_facade_proto protoreflect.FileDescriptor

var file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x62, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x73,
	0x73, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x62, 0x73,
	0x73, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb5, 0x02, 0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x09, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d,
	0x70, 0x2e, 0x62, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63,
	0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x4d, 0x5a,
	0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2f, 0x62, 0x73, 0x73, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2d, 0x66, 0x61,
	0x63, 0x61, 0x64, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x73, 0x73, 0x2d, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x2d, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x3b, 0x62, 0x73, 0x73, 0x5f, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescOnce sync.Once
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescData = file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDesc
)

func file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescGZIP() []byte {
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescOnce.Do(func() {
		file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescData)
	})
	return file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDescData
}

var file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_goTypes = []interface{}{
	(*OfficePayload)(nil),         // 0: ozonmp.bss_office_facade.v1.OfficePayload
	(*OfficeEvent)(nil),           // 1: ozonmp.bss_office_facade.v1.OfficeEvent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_depIdxs = []int32{
	2, // 0: ozonmp.bss_office_facade.v1.OfficeEvent.created:type_name -> google.protobuf.Timestamp
	2, // 1: ozonmp.bss_office_facade.v1.OfficeEvent.updated:type_name -> google.protobuf.Timestamp
	0, // 2: ozonmp.bss_office_facade.v1.OfficeEvent.payload:type_name -> ozonmp.bss_office_facade.v1.OfficePayload
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_init() }
func file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_init() {
	if File_ozonmp_bss_office_facade_v1_bss_office_facade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficePayload); i {
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
		file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeEvent); i {
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
			RawDescriptor: file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_goTypes,
		DependencyIndexes: file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_depIdxs,
		MessageInfos:      file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_msgTypes,
	}.Build()
	File_ozonmp_bss_office_facade_v1_bss_office_facade_proto = out.File
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_rawDesc = nil
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_goTypes = nil
	file_ozonmp_bss_office_facade_v1_bss_office_facade_proto_depIdxs = nil
}
