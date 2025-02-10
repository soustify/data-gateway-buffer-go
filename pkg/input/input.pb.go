// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: input.proto

package input

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaginationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Filter        []*FilteredRequest     `protobuf:"bytes,3,rep,name=filter,proto3" json:"filter,omitempty"`
	OrderBy       *OrderByRequest        `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_input_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationRequest) GetFilter() []*FilteredRequest {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *PaginationRequest) GetOrderBy() *OrderByRequest {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

type FilteredRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FilterContent  string                 `protobuf:"bytes,1,opt,name=filter_content,json=filterContent,proto3" json:"filter_content,omitempty"`
	FilteredFields []*FilterFieldsRequest `protobuf:"bytes,2,rep,name=filtered_fields,json=filteredFields,proto3" json:"filtered_fields,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FilteredRequest) Reset() {
	*x = FilteredRequest{}
	mi := &file_input_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilteredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteredRequest) ProtoMessage() {}

func (x *FilteredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteredRequest.ProtoReflect.Descriptor instead.
func (*FilteredRequest) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{1}
}

func (x *FilteredRequest) GetFilterContent() string {
	if x != nil {
		return x.FilterContent
	}
	return ""
}

func (x *FilteredRequest) GetFilteredFields() []*FilterFieldsRequest {
	if x != nil {
		return x.FilteredFields
	}
	return nil
}

type FilterFieldsRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Field          string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Operator       string                 `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	ComparisonKind string                 `protobuf:"bytes,3,opt,name=comparison_kind,json=comparisonKind,proto3" json:"comparison_kind,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FilterFieldsRequest) Reset() {
	*x = FilterFieldsRequest{}
	mi := &file_input_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterFieldsRequest) ProtoMessage() {}

func (x *FilterFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterFieldsRequest.ProtoReflect.Descriptor instead.
func (*FilterFieldsRequest) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{2}
}

func (x *FilterFieldsRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FilterFieldsRequest) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *FilterFieldsRequest) GetComparisonKind() string {
	if x != nil {
		return x.ComparisonKind
	}
	return ""
}

type OrderByRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order         string                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderByRequest) Reset() {
	*x = OrderByRequest{}
	mi := &file_input_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderByRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByRequest) ProtoMessage() {}

func (x *OrderByRequest) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByRequest.ProtoReflect.Descriptor instead.
func (*OrderByRequest) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{3}
}

func (x *OrderByRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *OrderByRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

var File_input_proto protoreflect.FileDescriptor

var file_input_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x7d, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x43, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x70, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69,
	0x73, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x50, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x75,
	0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x66, 0x79,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x3b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_input_proto_rawDescOnce sync.Once
	file_input_proto_rawDescData []byte
)

func file_input_proto_rawDescGZIP() []byte {
	file_input_proto_rawDescOnce.Do(func() {
		file_input_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_input_proto_rawDesc), len(file_input_proto_rawDesc)))
	})
	return file_input_proto_rawDescData
}

var file_input_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_input_proto_goTypes = []any{
	(*PaginationRequest)(nil),   // 0: input.PaginationRequest
	(*FilteredRequest)(nil),     // 1: input.FilteredRequest
	(*FilterFieldsRequest)(nil), // 2: input.FilterFieldsRequest
	(*OrderByRequest)(nil),      // 3: input.OrderByRequest
}
var file_input_proto_depIdxs = []int32{
	1, // 0: input.PaginationRequest.filter:type_name -> input.FilteredRequest
	3, // 1: input.PaginationRequest.order_by:type_name -> input.OrderByRequest
	2, // 2: input.FilteredRequest.filtered_fields:type_name -> input.FilterFieldsRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_input_proto_init() }
func file_input_proto_init() {
	if File_input_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_input_proto_rawDesc), len(file_input_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_input_proto_goTypes,
		DependencyIndexes: file_input_proto_depIdxs,
		MessageInfos:      file_input_proto_msgTypes,
	}.Build()
	File_input_proto = out.File
	file_input_proto_goTypes = nil
	file_input_proto_depIdxs = nil
}
