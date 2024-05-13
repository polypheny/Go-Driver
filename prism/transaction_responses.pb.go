//
//Messages concerning the handling of transactions.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.0
// source: org/polypheny/prism/transaction_responses.proto

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

//
//Represents the response after attempting to commit a transaction.
//This message serves as an acknowledgment that the commit request was received and processed,
//finalizing all changes made during the transaction.
//This message does not contain any fields. It simply serves as an acknowledgment of the commit.
type CommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitResponse) Reset() {
	*x = CommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_polypheny_prism_transaction_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResponse) ProtoMessage() {}

func (x *CommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_polypheny_prism_transaction_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResponse.ProtoReflect.Descriptor instead.
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return file_org_polypheny_prism_transaction_responses_proto_rawDescGZIP(), []int{0}
}

//
//Represents the response after attempting to rollback a transaction.
//This message serves as an acknowledgment that the rollback request was received and processed,
//undoing all changes made during the transaction.
//This message does not contain any fields. It simply serves as an acknowledgment of the rollback.
type RollbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RollbackResponse) Reset() {
	*x = RollbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_polypheny_prism_transaction_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackResponse) ProtoMessage() {}

func (x *RollbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_polypheny_prism_transaction_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackResponse.ProtoReflect.Descriptor instead.
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return file_org_polypheny_prism_transaction_responses_proto_rawDescGZIP(), []int{1}
}

var File_org_polypheny_prism_transaction_responses_proto protoreflect.FileDescriptor

var file_org_polypheny_prism_transaction_responses_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f,
	0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5a, 0x0a, 0x13,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x69, 0x73, 0x6d, 0x42, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x19, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d,
	0x3b, 0x70, 0x72, 0x69, 0x73, 0x6d, 0xaa, 0x02, 0x0f, 0x50, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65,
	0x6e, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x73, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_polypheny_prism_transaction_responses_proto_rawDescOnce sync.Once
	file_org_polypheny_prism_transaction_responses_proto_rawDescData = file_org_polypheny_prism_transaction_responses_proto_rawDesc
)

func file_org_polypheny_prism_transaction_responses_proto_rawDescGZIP() []byte {
	file_org_polypheny_prism_transaction_responses_proto_rawDescOnce.Do(func() {
		file_org_polypheny_prism_transaction_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_polypheny_prism_transaction_responses_proto_rawDescData)
	})
	return file_org_polypheny_prism_transaction_responses_proto_rawDescData
}

var file_org_polypheny_prism_transaction_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_org_polypheny_prism_transaction_responses_proto_goTypes = []interface{}{
	(*CommitResponse)(nil),   // 0: org.polypheny.prism.CommitResponse
	(*RollbackResponse)(nil), // 1: org.polypheny.prism.RollbackResponse
}
var file_org_polypheny_prism_transaction_responses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_org_polypheny_prism_transaction_responses_proto_init() }
func file_org_polypheny_prism_transaction_responses_proto_init() {
	if File_org_polypheny_prism_transaction_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_polypheny_prism_transaction_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResponse); i {
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
		file_org_polypheny_prism_transaction_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackResponse); i {
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
			RawDescriptor: file_org_polypheny_prism_transaction_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_polypheny_prism_transaction_responses_proto_goTypes,
		DependencyIndexes: file_org_polypheny_prism_transaction_responses_proto_depIdxs,
		MessageInfos:      file_org_polypheny_prism_transaction_responses_proto_msgTypes,
	}.Build()
	File_org_polypheny_prism_transaction_responses_proto = out.File
	file_org_polypheny_prism_transaction_responses_proto_rawDesc = nil
	file_org_polypheny_prism_transaction_responses_proto_goTypes = nil
	file_org_polypheny_prism_transaction_responses_proto_depIdxs = nil
}
