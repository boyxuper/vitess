//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file defines the replication related structures we use.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: replicationdata.proto

package replicationdata

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

// Status is the replication status for MySQL (returned by 'show slave status'
// and parsed into a Position and fields).
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position            string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	SlaveIoRunning      bool   `protobuf:"varint,2,opt,name=slave_io_running,json=slaveIoRunning,proto3" json:"slave_io_running,omitempty"`
	SlaveSqlRunning     bool   `protobuf:"varint,3,opt,name=slave_sql_running,json=slaveSqlRunning,proto3" json:"slave_sql_running,omitempty"`
	SecondsBehindMaster uint32 `protobuf:"varint,4,opt,name=seconds_behind_master,json=secondsBehindMaster,proto3" json:"seconds_behind_master,omitempty"`
	MasterHost          string `protobuf:"bytes,5,opt,name=master_host,json=masterHost,proto3" json:"master_host,omitempty"`
	MasterPort          int32  `protobuf:"varint,6,opt,name=master_port,json=masterPort,proto3" json:"master_port,omitempty"`
	MasterConnectRetry  int32  `protobuf:"varint,7,opt,name=master_connect_retry,json=masterConnectRetry,proto3" json:"master_connect_retry,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_replicationdata_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Status) GetSlaveIoRunning() bool {
	if x != nil {
		return x.SlaveIoRunning
	}
	return false
}

func (x *Status) GetSlaveSqlRunning() bool {
	if x != nil {
		return x.SlaveSqlRunning
	}
	return false
}

func (x *Status) GetSecondsBehindMaster() uint32 {
	if x != nil {
		return x.SecondsBehindMaster
	}
	return 0
}

func (x *Status) GetMasterHost() string {
	if x != nil {
		return x.MasterHost
	}
	return ""
}

func (x *Status) GetMasterPort() int32 {
	if x != nil {
		return x.MasterPort
	}
	return 0
}

func (x *Status) GetMasterConnectRetry() int32 {
	if x != nil {
		return x.MasterConnectRetry
	}
	return 0
}

var File_replicationdata_proto protoreflect.FileDescriptor

var file_replicationdata_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x6f, 0x5f, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6c, 0x61, 0x76, 0x65,
	0x49, 0x6f, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6c, 0x61,
	0x76, 0x65, 0x5f, 0x73, 0x71, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x71, 0x6c, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x5f, 0x62, 0x65, 0x68, 0x69, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x65, 0x68,
	0x69, 0x6e, 0x64, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6c, 0x74,
	0x68, 0x75, 0x62, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replicationdata_proto_rawDescOnce sync.Once
	file_replicationdata_proto_rawDescData = file_replicationdata_proto_rawDesc
)

func file_replicationdata_proto_rawDescGZIP() []byte {
	file_replicationdata_proto_rawDescOnce.Do(func() {
		file_replicationdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_replicationdata_proto_rawDescData)
	})
	return file_replicationdata_proto_rawDescData
}

var file_replicationdata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_replicationdata_proto_goTypes = []interface{}{
	(*Status)(nil), // 0: replicationdata.Status
}
var file_replicationdata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_replicationdata_proto_init() }
func file_replicationdata_proto_init() {
	if File_replicationdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replicationdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_replicationdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_replicationdata_proto_goTypes,
		DependencyIndexes: file_replicationdata_proto_depIdxs,
		MessageInfos:      file_replicationdata_proto_msgTypes,
	}.Build()
	File_replicationdata_proto = out.File
	file_replicationdata_proto_rawDesc = nil
	file_replicationdata_proto_goTypes = nil
	file_replicationdata_proto_depIdxs = nil
}
