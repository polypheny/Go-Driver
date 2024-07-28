//
//Messages pertaining to document frames as used to represent results from the document model.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: org/polypheny/prism/document_frame.proto

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

// The DocumentFrame message represents a frame containing one or more documents and is used to relay results or data entries structured as documents from the server to the client.
type DocumentFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of documents contained within this frame. Each entry in this list is an instance of ProtoDocument, representing an individual document.
	Documents []*ProtoDocument `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *DocumentFrame) Reset() {
	*x = DocumentFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_polypheny_prism_document_frame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentFrame) ProtoMessage() {}

func (x *DocumentFrame) ProtoReflect() protoreflect.Message {
	mi := &file_org_polypheny_prism_document_frame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentFrame.ProtoReflect.Descriptor instead.
func (*DocumentFrame) Descriptor() ([]byte, []int) {
	return file_org_polypheny_prism_document_frame_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentFrame) GetDocuments() []*ProtoDocument {
	if x != nil {
		return x.Documents
	}
	return nil
}

var File_org_polypheny_prism_document_frame_proto protoreflect.FileDescriptor

var file_org_polypheny_prism_document_frame_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f,
	0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x1a,
	0x1f, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2f, 0x70,
	0x72, 0x69, 0x73, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70,
	0x68, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x59, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70,
	0x68, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x42, 0x13, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x19, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79,
	0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x3b, 0x70, 0x72, 0x69, 0x73, 0x6d, 0xaa, 0x02, 0x0f, 0x50,
	0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x73, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_org_polypheny_prism_document_frame_proto_rawDescOnce sync.Once
	file_org_polypheny_prism_document_frame_proto_rawDescData = file_org_polypheny_prism_document_frame_proto_rawDesc
)

func file_org_polypheny_prism_document_frame_proto_rawDescGZIP() []byte {
	file_org_polypheny_prism_document_frame_proto_rawDescOnce.Do(func() {
		file_org_polypheny_prism_document_frame_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_polypheny_prism_document_frame_proto_rawDescData)
	})
	return file_org_polypheny_prism_document_frame_proto_rawDescData
}

var file_org_polypheny_prism_document_frame_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_org_polypheny_prism_document_frame_proto_goTypes = []any{
	(*DocumentFrame)(nil), // 0: org.polypheny.prism.DocumentFrame
	(*ProtoDocument)(nil), // 1: org.polypheny.prism.ProtoDocument
}
var file_org_polypheny_prism_document_frame_proto_depIdxs = []int32{
	1, // 0: org.polypheny.prism.DocumentFrame.documents:type_name -> org.polypheny.prism.ProtoDocument
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_org_polypheny_prism_document_frame_proto_init() }
func file_org_polypheny_prism_document_frame_proto_init() {
	if File_org_polypheny_prism_document_frame_proto != nil {
		return
	}
	file_org_polypheny_prism_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_org_polypheny_prism_document_frame_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DocumentFrame); i {
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
			RawDescriptor: file_org_polypheny_prism_document_frame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_polypheny_prism_document_frame_proto_goTypes,
		DependencyIndexes: file_org_polypheny_prism_document_frame_proto_depIdxs,
		MessageInfos:      file_org_polypheny_prism_document_frame_proto_msgTypes,
	}.Build()
	File_org_polypheny_prism_document_frame_proto = out.File
	file_org_polypheny_prism_document_frame_proto_rawDesc = nil
	file_org_polypheny_prism_document_frame_proto_goTypes = nil
	file_org_polypheny_prism_document_frame_proto_depIdxs = nil
}
