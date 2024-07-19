// Code generated by protoc-gen-go. DO NOT EDIT.

package protocol

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection       string            `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Filters          *Filters          `protobuf:"bytes,2,opt,name=filters,proto3" json:"filters,omitempty"`
	Verbose          bool              `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
	DryRun           bool              `protobuf:"varint,4,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	ConsistencyLevel *ConsistencyLevel `protobuf:"varint,5,opt,name=consistency_level,json=consistencyLevel,proto3,enum=weaviate.v1.ConsistencyLevel,oneof" json:"consistency_level,omitempty"`
	Tenant           *string           `protobuf:"bytes,6,opt,name=tenant,proto3,oneof" json:"tenant,omitempty"`
}

func (x *BatchDeleteRequest) Reset() {
	*x = BatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_batch_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteRequest) ProtoMessage() {}

func (x *BatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_batch_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*BatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_batch_delete_proto_rawDescGZIP(), []int{0}
}

func (x *BatchDeleteRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *BatchDeleteRequest) GetFilters() *Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *BatchDeleteRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

func (x *BatchDeleteRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *BatchDeleteRequest) GetConsistencyLevel() ConsistencyLevel {
	if x != nil && x.ConsistencyLevel != nil {
		return *x.ConsistencyLevel
	}
	return ConsistencyLevel_CONSISTENCY_LEVEL_UNSPECIFIED
}

func (x *BatchDeleteRequest) GetTenant() string {
	if x != nil && x.Tenant != nil {
		return *x.Tenant
	}
	return ""
}

type BatchDeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Took       float32              `protobuf:"fixed32,1,opt,name=took,proto3" json:"took,omitempty"`
	Failed     int64                `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	Matches    int64                `protobuf:"varint,3,opt,name=matches,proto3" json:"matches,omitempty"`
	Successful int64                `protobuf:"varint,4,opt,name=successful,proto3" json:"successful,omitempty"`
	Objects    []*BatchDeleteObject `protobuf:"bytes,5,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *BatchDeleteReply) Reset() {
	*x = BatchDeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_batch_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteReply) ProtoMessage() {}

func (x *BatchDeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_batch_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteReply.ProtoReflect.Descriptor instead.
func (*BatchDeleteReply) Descriptor() ([]byte, []int) {
	return file_v1_batch_delete_proto_rawDescGZIP(), []int{1}
}

func (x *BatchDeleteReply) GetTook() float32 {
	if x != nil {
		return x.Took
	}
	return 0
}

func (x *BatchDeleteReply) GetFailed() int64 {
	if x != nil {
		return x.Failed
	}
	return 0
}

func (x *BatchDeleteReply) GetMatches() int64 {
	if x != nil {
		return x.Matches
	}
	return 0
}

func (x *BatchDeleteReply) GetSuccessful() int64 {
	if x != nil {
		return x.Successful
	}
	return 0
}

func (x *BatchDeleteReply) GetObjects() []*BatchDeleteObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

type BatchDeleteObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       []byte  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Successful bool    `protobuf:"varint,2,opt,name=successful,proto3" json:"successful,omitempty"`
	Error      *string `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"` // empty string means no error
}

func (x *BatchDeleteObject) Reset() {
	*x = BatchDeleteObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_batch_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteObject) ProtoMessage() {}

func (x *BatchDeleteObject) ProtoReflect() protoreflect.Message {
	mi := &file_v1_batch_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteObject.ProtoReflect.Descriptor instead.
func (*BatchDeleteObject) Descriptor() ([]byte, []int) {
	return file_v1_batch_delete_proto_rawDescGZIP(), []int{2}
}

func (x *BatchDeleteObject) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *BatchDeleteObject) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

func (x *BatchDeleteObject) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_v1_batch_delete_proto protoreflect.FileDescriptor

var file_v1_batch_delete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x65,
	0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x4f, 0x0a,
	0x11, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0xb2, 0x01, 0x0a,
	0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x6c, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x75, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61,
	0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_batch_delete_proto_rawDescOnce sync.Once
	file_v1_batch_delete_proto_rawDescData = file_v1_batch_delete_proto_rawDesc
)

func file_v1_batch_delete_proto_rawDescGZIP() []byte {
	file_v1_batch_delete_proto_rawDescOnce.Do(func() {
		file_v1_batch_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_batch_delete_proto_rawDescData)
	})
	return file_v1_batch_delete_proto_rawDescData
}

var (
	file_v1_batch_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_v1_batch_delete_proto_goTypes  = []interface{}{
		(*BatchDeleteRequest)(nil), // 0: weaviate.v1.BatchDeleteRequest
		(*BatchDeleteReply)(nil),   // 1: weaviate.v1.BatchDeleteReply
		(*BatchDeleteObject)(nil),  // 2: weaviate.v1.BatchDeleteObject
		(*Filters)(nil),            // 3: weaviate.v1.Filters
		(ConsistencyLevel)(0),      // 4: weaviate.v1.ConsistencyLevel
	}
)
var file_v1_batch_delete_proto_depIdxs = []int32{
	3, // 0: weaviate.v1.BatchDeleteRequest.filters:type_name -> weaviate.v1.Filters
	4, // 1: weaviate.v1.BatchDeleteRequest.consistency_level:type_name -> weaviate.v1.ConsistencyLevel
	2, // 2: weaviate.v1.BatchDeleteReply.objects:type_name -> weaviate.v1.BatchDeleteObject
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_batch_delete_proto_init() }
func file_v1_batch_delete_proto_init() {
	if File_v1_batch_delete_proto != nil {
		return
	}
	file_v1_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_batch_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteRequest); i {
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
		file_v1_batch_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteReply); i {
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
		file_v1_batch_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteObject); i {
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
	file_v1_batch_delete_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_batch_delete_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_batch_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_batch_delete_proto_goTypes,
		DependencyIndexes: file_v1_batch_delete_proto_depIdxs,
		MessageInfos:      file_v1_batch_delete_proto_msgTypes,
	}.Build()
	File_v1_batch_delete_proto = out.File
	file_v1_batch_delete_proto_rawDesc = nil
	file_v1_batch_delete_proto_goTypes = nil
	file_v1_batch_delete_proto_depIdxs = nil
}
