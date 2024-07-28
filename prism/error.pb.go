//
//Messages used to send errors to the client.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: org/polypheny/prism/error.proto

package prism

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

// The ErrorDetails message conveys specific information about an error encountered during the processing of a request.
// This detailed feedback aids clients in understanding the nature and cause of an issue.
type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A numeric code representing the specific error encountered. This code can be used programmatically to handle specific error cases.
	ErrorCode *int32 `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,oneof" json:"error_code,omitempty"`
	// An optional state description in the form of an identifier that provides further context about the error.
	State *string `protobuf:"bytes,2,opt,name=state,proto3,oneof" json:"state,omitempty"`
	// A human-readable message describing the error in detail. This message offers clients a clear understanding of the issue.
	Message *string `protobuf:"bytes,3,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_polypheny_prism_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_org_polypheny_prism_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_org_polypheny_prism_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetErrorCode() int32 {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return 0
}

func (x *ErrorDetails) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

func (x *ErrorDetails) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_org_polypheny_prism_error_proto protoreflect.FileDescriptor

var file_org_polypheny_prism_error_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f,
	0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4b, 0x0a, 0x13, 0x6f, 0x72,
	0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x69, 0x73,
	0x6d, 0x42, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x01, 0x5a, 0x19, 0x6f, 0x72, 0x67, 0x2f,
	0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x3b,
	0x70, 0x72, 0x69, 0x73, 0x6d, 0xaa, 0x02, 0x0f, 0x50, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e,
	0x79, 0x2e, 0x50, 0x72, 0x69, 0x73, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_polypheny_prism_error_proto_rawDescOnce sync.Once
	file_org_polypheny_prism_error_proto_rawDescData = file_org_polypheny_prism_error_proto_rawDesc
)

func file_org_polypheny_prism_error_proto_rawDescGZIP() []byte {
	file_org_polypheny_prism_error_proto_rawDescOnce.Do(func() {
		file_org_polypheny_prism_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_polypheny_prism_error_proto_rawDescData)
	})
	return file_org_polypheny_prism_error_proto_rawDescData
}

var file_org_polypheny_prism_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_org_polypheny_prism_error_proto_goTypes = []any{
	(*ErrorDetails)(nil), // 0: org.polypheny.prism.ErrorDetails
}
var file_org_polypheny_prism_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_org_polypheny_prism_error_proto_init() }
func file_org_polypheny_prism_error_proto_init() {
	if File_org_polypheny_prism_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_polypheny_prism_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorDetails); i {
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
	file_org_polypheny_prism_error_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_org_polypheny_prism_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_polypheny_prism_error_proto_goTypes,
		DependencyIndexes: file_org_polypheny_prism_error_proto_depIdxs,
		MessageInfos:      file_org_polypheny_prism_error_proto_msgTypes,
	}.Build()
	File_org_polypheny_prism_error_proto = out.File
	file_org_polypheny_prism_error_proto_rawDesc = nil
	file_org_polypheny_prism_error_proto_goTypes = nil
	file_org_polypheny_prism_error_proto_depIdxs = nil
}
