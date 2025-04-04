// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: context_domain_table_detail.proto

package contextdomain_table

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/soustify/data-gateway-buffer-go/pkg/input"
	output "github.com/soustify/data-gateway-buffer-go/pkg/output"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	Operation_INSERT Operation = 0
	Operation_UPDATE Operation = 1
	Operation_DELETE Operation = 2
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "INSERT",
		1: "UPDATE",
		2: "DELETE",
	}
	Operation_value = map[string]int32{
		"INSERT": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_context_domain_table_detail_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_context_domain_table_detail_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_context_domain_table_detail_proto_rawDescGZIP(), []int{0}
}

type ContextDomainTableDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextId   string    `protobuf:"bytes,1,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	EntityId    string    `protobuf:"bytes,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityTable string    `protobuf:"bytes,3,opt,name=entity_table,json=entityTable,proto3" json:"entity_table,omitempty"`
	Operation   Operation `protobuf:"varint,4,opt,name=operation,proto3,enum=contextdomain_table.Operation" json:"operation,omitempty"`
}

func (x *ContextDomainTableDetailRequest) Reset() {
	*x = ContextDomainTableDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_domain_table_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextDomainTableDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextDomainTableDetailRequest) ProtoMessage() {}

func (x *ContextDomainTableDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_context_domain_table_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextDomainTableDetailRequest.ProtoReflect.Descriptor instead.
func (*ContextDomainTableDetailRequest) Descriptor() ([]byte, []int) {
	return file_context_domain_table_detail_proto_rawDescGZIP(), []int{0}
}

func (x *ContextDomainTableDetailRequest) GetContextId() string {
	if x != nil {
		return x.ContextId
	}
	return ""
}

func (x *ContextDomainTableDetailRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ContextDomainTableDetailRequest) GetEntityTable() string {
	if x != nil {
		return x.EntityTable
	}
	return ""
}

func (x *ContextDomainTableDetailRequest) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_INSERT
}

var File_context_domain_table_detail_proto protoreflect.FileDescriptor

var file_context_domain_table_detail_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x0a, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x2f, 0x0a, 0x09, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52,
	0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x32, 0x84, 0x01, 0x0a, 0x1f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x61, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x42, 0x7a, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x75, 0x73, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_context_domain_table_detail_proto_rawDescOnce sync.Once
	file_context_domain_table_detail_proto_rawDescData = file_context_domain_table_detail_proto_rawDesc
)

func file_context_domain_table_detail_proto_rawDescGZIP() []byte {
	file_context_domain_table_detail_proto_rawDescOnce.Do(func() {
		file_context_domain_table_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_context_domain_table_detail_proto_rawDescData)
	})
	return file_context_domain_table_detail_proto_rawDescData
}

var file_context_domain_table_detail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_context_domain_table_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_context_domain_table_detail_proto_goTypes = []interface{}{
	(Operation)(0),                          // 0: contextdomain_table.Operation
	(*ContextDomainTableDetailRequest)(nil), // 1: contextdomain_table.ContextDomainTableDetailRequest
	(*output.PersistenceDataResponse)(nil),  // 2: output.PersistenceDataResponse
}
var file_context_domain_table_detail_proto_depIdxs = []int32{
	0, // 0: contextdomain_table.ContextDomainTableDetailRequest.operation:type_name -> contextdomain_table.Operation
	1, // 1: contextdomain_table.ContextDomainTableDetailService.Create:input_type -> contextdomain_table.ContextDomainTableDetailRequest
	2, // 2: contextdomain_table.ContextDomainTableDetailService.Create:output_type -> output.PersistenceDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_context_domain_table_detail_proto_init() }
func file_context_domain_table_detail_proto_init() {
	if File_context_domain_table_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_context_domain_table_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextDomainTableDetailRequest); i {
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
			RawDescriptor: file_context_domain_table_detail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_context_domain_table_detail_proto_goTypes,
		DependencyIndexes: file_context_domain_table_detail_proto_depIdxs,
		EnumInfos:         file_context_domain_table_detail_proto_enumTypes,
		MessageInfos:      file_context_domain_table_detail_proto_msgTypes,
	}.Build()
	File_context_domain_table_detail_proto = out.File
	file_context_domain_table_detail_proto_rawDesc = nil
	file_context_domain_table_detail_proto_goTypes = nil
	file_context_domain_table_detail_proto_depIdxs = nil
}
