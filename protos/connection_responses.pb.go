//
//Messages related to establishing and maintaining a connection.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: polyprism/connection_responses.proto

package protos

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

// The ConnectionResponse message is sent by the server in response to a ConnectionRequest from the client.
// It provides feedback regarding the compatibility of the client’s API version with the server, as well as other important information about the connection.
type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the client’s API version is compatible with the server.
	// If true, the connection can proceed; if false, the client should consider updating or downgrading its API version.
	IsCompatible bool `protobuf:"varint,1,opt,name=is_compatible,json=isCompatible,proto3" json:"is_compatible,omitempty"`
	// The major version number of the server’s API. Helps the client ascertain the API level of the server.
	MajorApiVersion int32 `protobuf:"varint,2,opt,name=major_api_version,json=majorApiVersion,proto3" json:"major_api_version,omitempty"`
	// The minor version number of the server’s API. Provides finer granularity about the server’s capabilities.
	MinorApiVersion int32 `protobuf:"varint,3,opt,name=minor_api_version,json=minorApiVersion,proto3" json:"minor_api_version,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyprism_connection_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyprism_connection_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_polyprism_connection_responses_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionResponse) GetIsCompatible() bool {
	if x != nil {
		return x.IsCompatible
	}
	return false
}

func (x *ConnectionResponse) GetMajorApiVersion() int32 {
	if x != nil {
		return x.MajorApiVersion
	}
	return 0
}

func (x *ConnectionResponse) GetMinorApiVersion() int32 {
	if x != nil {
		return x.MinorApiVersion
	}
	return 0
}

// The DisconnectionResponse message is sent by the server in response to a DisconnectRequest from the client.
// It provides an acknowledgment that the disconnection process has been acknowledged and processed by the server.
// This message does not contain any fields. It simply serves as an acknowledgment of the disconnection process.
type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyprism_connection_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyprism_connection_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_polyprism_connection_responses_proto_rawDescGZIP(), []int{1}
}

// The ConnectionCheckResponse message is sent by the server in response to a ConnectionCheckRequest from the client.
// It acts as an affirmation that the server is responsive, and the connection is still valid.
// This message does not contain any fields.
// Its receipt by the client is an indicator of the server’s responsiveness and the validity of the connection.
type ConnectionCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionCheckResponse) Reset() {
	*x = ConnectionCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyprism_connection_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionCheckResponse) ProtoMessage() {}

func (x *ConnectionCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyprism_connection_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionCheckResponse.ProtoReflect.Descriptor instead.
func (*ConnectionCheckResponse) Descriptor() ([]byte, []int) {
	return file_polyprism_connection_responses_proto_rawDescGZIP(), []int{2}
}

// The ConnectionPropertiesUpdateResponse message is sent by the server in response to a ConnectionPropertiesUpdateRequest from the client.
// It acknowledges the success of the requested connection properties update.
// This message does not contain any fields. It serves as an acknowledgment of the update process for the connection properties.
type ConnectionPropertiesUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionPropertiesUpdateResponse) Reset() {
	*x = ConnectionPropertiesUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyprism_connection_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionPropertiesUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionPropertiesUpdateResponse) ProtoMessage() {}

func (x *ConnectionPropertiesUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyprism_connection_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionPropertiesUpdateResponse.ProtoReflect.Descriptor instead.
func (*ConnectionPropertiesUpdateResponse) Descriptor() ([]byte, []int) {
	return file_polyprism_connection_responses_proto_rawDescGZIP(), []int{3}
}

var File_polyprism_connection_responses_proto protoreflect.FileDescriptor

var file_polyprism_connection_responses_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x22, 0x91, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x41, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x22, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x0a, 0x25, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x79, 0x2e, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_polyprism_connection_responses_proto_rawDescOnce sync.Once
	file_polyprism_connection_responses_proto_rawDescData = file_polyprism_connection_responses_proto_rawDesc
)

func file_polyprism_connection_responses_proto_rawDescGZIP() []byte {
	file_polyprism_connection_responses_proto_rawDescOnce.Do(func() {
		file_polyprism_connection_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_polyprism_connection_responses_proto_rawDescData)
	})
	return file_polyprism_connection_responses_proto_rawDescData
}

var file_polyprism_connection_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_polyprism_connection_responses_proto_goTypes = []interface{}{
	(*ConnectionResponse)(nil),                 // 0: polypheny.protointerface.ConnectionResponse
	(*DisconnectResponse)(nil),                 // 1: polypheny.protointerface.DisconnectResponse
	(*ConnectionCheckResponse)(nil),            // 2: polypheny.protointerface.ConnectionCheckResponse
	(*ConnectionPropertiesUpdateResponse)(nil), // 3: polypheny.protointerface.ConnectionPropertiesUpdateResponse
}
var file_polyprism_connection_responses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_polyprism_connection_responses_proto_init() }
func file_polyprism_connection_responses_proto_init() {
	if File_polyprism_connection_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_polyprism_connection_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
		file_polyprism_connection_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
		file_polyprism_connection_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionCheckResponse); i {
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
		file_polyprism_connection_responses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionPropertiesUpdateResponse); i {
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
			RawDescriptor: file_polyprism_connection_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_polyprism_connection_responses_proto_goTypes,
		DependencyIndexes: file_polyprism_connection_responses_proto_depIdxs,
		MessageInfos:      file_polyprism_connection_responses_proto_msgTypes,
	}.Build()
	File_polyprism_connection_responses_proto = out.File
	file_polyprism_connection_responses_proto_rawDesc = nil
	file_polyprism_connection_responses_proto_goTypes = nil
	file_polyprism_connection_responses_proto_depIdxs = nil
}
