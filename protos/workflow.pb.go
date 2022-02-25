// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: workflow.proto

// Area Workflow API

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A workflow is a set of task (action and reaction) to execute
type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner string  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tasks []*Task `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *Workflow) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Workflow) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Workflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workflow) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type CreateWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tasks []*Task `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *CreateWorkflowRequest) Reset() {
	*x = CreateWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkflowRequest) ProtoMessage() {}

func (x *CreateWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkflowRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateWorkflowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWorkflowRequest) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ListWorkflowsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ListWorkflowsRequest) Reset() {
	*x = ListWorkflowsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowsRequest) ProtoMessage() {}

func (x *ListWorkflowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowsRequest.ProtoReflect.Descriptor instead.
func (*ListWorkflowsRequest) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *ListWorkflowsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type ListWorkflowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflows []*Workflow `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
}

func (x *ListWorkflowsResponse) Reset() {
	*x = ListWorkflowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkflowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkflowsResponse) ProtoMessage() {}

func (x *ListWorkflowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkflowsResponse.ProtoReflect.Descriptor instead.
func (*ListWorkflowsResponse) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *ListWorkflowsResponse) GetWorkflows() []*Workflow {
	if x != nil {
		return x.Workflows
	}
	return nil
}

type GetWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWorkflowRequest) Reset() {
	*x = GetWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowRequest) ProtoMessage() {}

func (x *GetWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *GetWorkflowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner *string `protobuf:"bytes,2,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	Name  *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *UpdateWorkflowRequest) Reset() {
	*x = UpdateWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowRequest) ProtoMessage() {}

func (x *UpdateWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateWorkflowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWorkflowRequest) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

func (x *UpdateWorkflowRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type DeleteWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWorkflowRequest) Reset() {
	*x = DeleteWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkflowRequest) ProtoMessage() {}

func (x *DeleteWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkflowRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWorkflowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_workflow_proto protoreflect.FileDescriptor

var file_workflow_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x72, 0x65, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x68, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22,
	0x2c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x4e, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x72, 0x65, 0x61,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xaa, 0x03, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x24, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x12, 0x5a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x2e, 0x61,
	0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x4f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x24, 0x2e, 0x61, 0x72, 0x65,
	0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x4e, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x24, 0x2e, 0x61, 0x72,
	0x65, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_proto_rawDescOnce sync.Once
	file_workflow_proto_rawDescData = file_workflow_proto_rawDesc
)

func file_workflow_proto_rawDescGZIP() []byte {
	file_workflow_proto_rawDescOnce.Do(func() {
		file_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_proto_rawDescData)
	})
	return file_workflow_proto_rawDescData
}

var file_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workflow_proto_goTypes = []interface{}{
	(*Workflow)(nil),              // 0: area.workflow.Workflow
	(*CreateWorkflowRequest)(nil), // 1: area.workflow.CreateWorkflowRequest
	(*ListWorkflowsRequest)(nil),  // 2: area.workflow.ListWorkflowsRequest
	(*ListWorkflowsResponse)(nil), // 3: area.workflow.ListWorkflowsResponse
	(*GetWorkflowRequest)(nil),    // 4: area.workflow.GetWorkflowRequest
	(*UpdateWorkflowRequest)(nil), // 5: area.workflow.UpdateWorkflowRequest
	(*DeleteWorkflowRequest)(nil), // 6: area.workflow.DeleteWorkflowRequest
	(*Task)(nil),                  // 7: area.task.Task
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_workflow_proto_depIdxs = []int32{
	7, // 0: area.workflow.Workflow.tasks:type_name -> area.task.Task
	7, // 1: area.workflow.CreateWorkflowRequest.tasks:type_name -> area.task.Task
	0, // 2: area.workflow.ListWorkflowsResponse.workflows:type_name -> area.workflow.Workflow
	1, // 3: area.workflow.WorkflowService.CreateWorkflow:input_type -> area.workflow.CreateWorkflowRequest
	2, // 4: area.workflow.WorkflowService.ListWorkflows:input_type -> area.workflow.ListWorkflowsRequest
	4, // 5: area.workflow.WorkflowService.GetWorkflow:input_type -> area.workflow.GetWorkflowRequest
	5, // 6: area.workflow.WorkflowService.UpdateWorkflow:input_type -> area.workflow.UpdateWorkflowRequest
	6, // 7: area.workflow.WorkflowService.DeleteWorkflow:input_type -> area.workflow.DeleteWorkflowRequest
	0, // 8: area.workflow.WorkflowService.CreateWorkflow:output_type -> area.workflow.Workflow
	3, // 9: area.workflow.WorkflowService.ListWorkflows:output_type -> area.workflow.ListWorkflowsResponse
	0, // 10: area.workflow.WorkflowService.GetWorkflow:output_type -> area.workflow.Workflow
	0, // 11: area.workflow.WorkflowService.UpdateWorkflow:output_type -> area.workflow.Workflow
	8, // 12: area.workflow.WorkflowService.DeleteWorkflow:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_workflow_proto_init() }
func file_workflow_proto_init() {
	if File_workflow_proto != nil {
		return
	}
	file_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkflowRequest); i {
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
		file_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowsRequest); i {
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
		file_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkflowsResponse); i {
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
		file_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowRequest); i {
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
		file_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkflowRequest); i {
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
		file_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkflowRequest); i {
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
	file_workflow_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_proto_goTypes,
		DependencyIndexes: file_workflow_proto_depIdxs,
		MessageInfos:      file_workflow_proto_msgTypes,
	}.Build()
	File_workflow_proto = out.File
	file_workflow_proto_rawDesc = nil
	file_workflow_proto_goTypes = nil
	file_workflow_proto_depIdxs = nil
}
