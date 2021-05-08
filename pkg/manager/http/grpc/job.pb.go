// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pkg/manager/http/grpc/job.proto

package grpc

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

type JobMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DoneAt    string `protobuf:"bytes,5,opt,name=doneAt,proto3" json:"doneAt,omitempty"`
}

func (x *JobMessage) Reset() {
	*x = JobMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMessage) ProtoMessage() {}

func (x *JobMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMessage.ProtoReflect.Descriptor instead.
func (*JobMessage) Descriptor() ([]byte, []int) {
	return file_pkg_manager_http_grpc_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *JobMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JobMessage) GetDoneAt() string {
	if x != nil {
		return x.DoneAt
	}
	return ""
}

type AddJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	DoneAt string `protobuf:"bytes,3,opt,name=doneAt,proto3" json:"doneAt,omitempty"`
}

func (x *AddJobRequest) Reset() {
	*x = AddJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddJobRequest) ProtoMessage() {}

func (x *AddJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddJobRequest.ProtoReflect.Descriptor instead.
func (*AddJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_manager_http_grpc_job_proto_rawDescGZIP(), []int{1}
}

func (x *AddJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddJobRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddJobRequest) GetDoneAt() string {
	if x != nil {
		return x.DoneAt
	}
	return ""
}

type AddJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddJobResponse) Reset() {
	*x = AddJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddJobResponse) ProtoMessage() {}

func (x *AddJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddJobResponse.ProtoReflect.Descriptor instead.
func (*AddJobResponse) Descriptor() ([]byte, []int) {
	return file_pkg_manager_http_grpc_job_proto_rawDescGZIP(), []int{2}
}

func (x *AddJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_manager_http_grpc_job_proto_rawDescGZIP(), []int{3}
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*JobMessage `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_manager_http_grpc_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_manager_http_grpc_job_proto_rawDescGZIP(), []int{4}
}

func (x *ListJobsResponse) GetJobs() []*JobMessage {
	if x != nil {
		return x.Jobs
	}
	return nil
}

var File_pkg_manager_http_grpc_job_proto protoreflect.FileDescriptor

var file_pkg_manager_http_grpc_job_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x70, 0x6b, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x7c, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x22, 0x51, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x32, 0xb9, 0x01, 0x0a, 0x03, 0x4a, 0x6f,
	0x62, 0x12, 0x55, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x4a, 0x6f, 0x62, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x6e, 0x67, 0x6a, 0x75, 0x6e, 0x79, 0x6f, 0x75, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_manager_http_grpc_job_proto_rawDescOnce sync.Once
	file_pkg_manager_http_grpc_job_proto_rawDescData = file_pkg_manager_http_grpc_job_proto_rawDesc
)

func file_pkg_manager_http_grpc_job_proto_rawDescGZIP() []byte {
	file_pkg_manager_http_grpc_job_proto_rawDescOnce.Do(func() {
		file_pkg_manager_http_grpc_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_manager_http_grpc_job_proto_rawDescData)
	})
	return file_pkg_manager_http_grpc_job_proto_rawDescData
}

var file_pkg_manager_http_grpc_job_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_manager_http_grpc_job_proto_goTypes = []interface{}{
	(*JobMessage)(nil),       // 0: pkg.manager.http.grpc.JobMessage
	(*AddJobRequest)(nil),    // 1: pkg.manager.http.grpc.AddJobRequest
	(*AddJobResponse)(nil),   // 2: pkg.manager.http.grpc.AddJobResponse
	(*ListJobsRequest)(nil),  // 3: pkg.manager.http.grpc.ListJobsRequest
	(*ListJobsResponse)(nil), // 4: pkg.manager.http.grpc.ListJobsResponse
}
var file_pkg_manager_http_grpc_job_proto_depIdxs = []int32{
	0, // 0: pkg.manager.http.grpc.ListJobsResponse.jobs:type_name -> pkg.manager.http.grpc.JobMessage
	1, // 1: pkg.manager.http.grpc.Job.AddJob:input_type -> pkg.manager.http.grpc.AddJobRequest
	3, // 2: pkg.manager.http.grpc.Job.ListJobs:input_type -> pkg.manager.http.grpc.ListJobsRequest
	2, // 3: pkg.manager.http.grpc.Job.AddJob:output_type -> pkg.manager.http.grpc.AddJobResponse
	4, // 4: pkg.manager.http.grpc.Job.ListJobs:output_type -> pkg.manager.http.grpc.ListJobsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_manager_http_grpc_job_proto_init() }
func file_pkg_manager_http_grpc_job_proto_init() {
	if File_pkg_manager_http_grpc_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_manager_http_grpc_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMessage); i {
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
		file_pkg_manager_http_grpc_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddJobRequest); i {
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
		file_pkg_manager_http_grpc_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddJobResponse); i {
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
		file_pkg_manager_http_grpc_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_pkg_manager_http_grpc_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
			RawDescriptor: file_pkg_manager_http_grpc_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_manager_http_grpc_job_proto_goTypes,
		DependencyIndexes: file_pkg_manager_http_grpc_job_proto_depIdxs,
		MessageInfos:      file_pkg_manager_http_grpc_job_proto_msgTypes,
	}.Build()
	File_pkg_manager_http_grpc_job_proto = out.File
	file_pkg_manager_http_grpc_job_proto_rawDesc = nil
	file_pkg_manager_http_grpc_job_proto_goTypes = nil
	file_pkg_manager_http_grpc_job_proto_depIdxs = nil
}