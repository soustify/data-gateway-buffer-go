// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: context_domain_transaction.proto

package contextdomain_transaction

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

type ContextDomainTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserPool string `protobuf:"bytes,3,opt,name=user_pool,json=userPool,proto3" json:"user_pool,omitempty"`
}

func (x *ContextDomainTransactionRequest) Reset() {
	*x = ContextDomainTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_context_domain_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextDomainTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextDomainTransactionRequest) ProtoMessage() {}

func (x *ContextDomainTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_context_domain_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextDomainTransactionRequest.ProtoReflect.Descriptor instead.
func (*ContextDomainTransactionRequest) Descriptor() ([]byte, []int) {
	return file_context_domain_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ContextDomainTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContextDomainTransactionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ContextDomainTransactionRequest) GetUserPool() string {
	if x != nil {
		return x.UserPool
	}
	return ""
}

var File_context_domain_transaction_proto protoreflect.FileDescriptor

var file_context_domain_transaction_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x0a, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x32, 0x8a,
	0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x8c, 0x01, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x2d,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_context_domain_transaction_proto_rawDescOnce sync.Once
	file_context_domain_transaction_proto_rawDescData = file_context_domain_transaction_proto_rawDesc
)

func file_context_domain_transaction_proto_rawDescGZIP() []byte {
	file_context_domain_transaction_proto_rawDescOnce.Do(func() {
		file_context_domain_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_context_domain_transaction_proto_rawDescData)
	})
	return file_context_domain_transaction_proto_rawDescData
}

var file_context_domain_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_context_domain_transaction_proto_goTypes = []interface{}{
	(*ContextDomainTransactionRequest)(nil), // 0: contextdomain_transaction.ContextDomainTransactionRequest
	(*output.PersistenceDataResponse)(nil),  // 1: output.PersistenceDataResponse
}
var file_context_domain_transaction_proto_depIdxs = []int32{
	0, // 0: contextdomain_transaction.ContextDomainTransactionService.Create:input_type -> contextdomain_transaction.ContextDomainTransactionRequest
	1, // 1: contextdomain_transaction.ContextDomainTransactionService.Create:output_type -> output.PersistenceDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_context_domain_transaction_proto_init() }
func file_context_domain_transaction_proto_init() {
	if File_context_domain_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_context_domain_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextDomainTransactionRequest); i {
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
			RawDescriptor: file_context_domain_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_context_domain_transaction_proto_goTypes,
		DependencyIndexes: file_context_domain_transaction_proto_depIdxs,
		MessageInfos:      file_context_domain_transaction_proto_msgTypes,
	}.Build()
	File_context_domain_transaction_proto = out.File
	file_context_domain_transaction_proto_rawDesc = nil
	file_context_domain_transaction_proto_goTypes = nil
	file_context_domain_transaction_proto_depIdxs = nil
}
