// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: tasks.proto

package api

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

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_tasks_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllTasksRequest) Reset() {
	*x = GetAllTasksRequest{}
	mi := &file_tasks_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksRequest) ProtoMessage() {}

func (x *GetAllTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksRequest.ProtoReflect.Descriptor instead.
func (*GetAllTasksRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{1}
}

type GetAllTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllTasksResponse) Reset() {
	*x = GetAllTasksResponse{}
	mi := &file_tasks_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksResponse) ProtoMessage() {}

func (x *GetAllTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksResponse.ProtoReflect.Descriptor instead.
func (*GetAllTasksResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_tasks_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	mi := &file_tasks_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type AddTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	mi := &file_tasks_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *AddTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type AddTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTaskResponse) Reset() {
	*x = AddTaskResponse{}
	mi := &file_tasks_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponse) ProtoMessage() {}

func (x *AddTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponse.ProtoReflect.Descriptor instead.
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{6}
}

func (x *AddTaskResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Task          *Task                  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_tasks_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	mi := &file_tasks_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_tasks_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_tasks_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_tasks_proto protoreflect.FileDescriptor

const file_tasks_proto_rawDesc = "" +
	"\n" +
	"\vtasks.proto\x12\vtasks_proto\"Q\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x04R\x06userId\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"\x14\n" +
	"\x12GetAllTasksRequest\">\n" +
	"\x13GetAllTasksResponse\x12'\n" +
	"\x05tasks\x18\x01 \x03(\v2\x11.tasks_proto.TaskR\x05tasks\" \n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"8\n" +
	"\x0fGetTaskResponse\x12%\n" +
	"\x04task\x18\x01 \x01(\v2\x11.tasks_proto.TaskR\x04task\"7\n" +
	"\x0eAddTaskRequest\x12%\n" +
	"\x04task\x18\x01 \x01(\v2\x11.tasks_proto.TaskR\x04task\"!\n" +
	"\x0fAddTaskResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"J\n" +
	"\x11UpdateTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12%\n" +
	"\x04task\x18\x02 \x01(\v2\x11.tasks_proto.TaskR\x04task\";\n" +
	"\x12UpdateTaskResponse\x12%\n" +
	"\x04task\x18\x01 \x01(\v2\x11.tasks_proto.TaskR\x04task\"#\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\".\n" +
	"\x12DeleteTaskResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xdb\x03\n" +
	"\vTaskService\x12P\n" +
	"\vGetAllTasks\x12\x1f.tasks_proto.GetAllTasksRequest\x1a .tasks_proto.GetAllTasksResponse\x12P\n" +
	"\x0fGetUserAllTasks\x12\x1b.tasks_proto.GetTaskRequest\x1a .tasks_proto.GetAllTasksResponse\x12D\n" +
	"\aGetTask\x12\x1b.tasks_proto.GetTaskRequest\x1a\x1c.tasks_proto.GetTaskResponse\x12D\n" +
	"\aAddTask\x12\x1b.tasks_proto.AddTaskRequest\x1a\x1c.tasks_proto.AddTaskResponse\x12M\n" +
	"\n" +
	"UpdateTask\x12\x1e.tasks_proto.UpdateTaskRequest\x1a\x1f.tasks_proto.UpdateTaskResponse\x12M\n" +
	"\n" +
	"DeleteTask\x12\x1e.tasks_proto.DeleteTaskRequest\x1a\x1f.tasks_proto.DeleteTaskResponseB(Z&github.com/literaen/simple_project/apib\x06proto3"

var (
	file_tasks_proto_rawDescOnce sync.Once
	file_tasks_proto_rawDescData []byte
)

func file_tasks_proto_rawDescGZIP() []byte {
	file_tasks_proto_rawDescOnce.Do(func() {
		file_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tasks_proto_rawDesc), len(file_tasks_proto_rawDesc)))
	})
	return file_tasks_proto_rawDescData
}

var file_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tasks_proto_goTypes = []any{
	(*Task)(nil),                // 0: tasks_proto.Task
	(*GetAllTasksRequest)(nil),  // 1: tasks_proto.GetAllTasksRequest
	(*GetAllTasksResponse)(nil), // 2: tasks_proto.GetAllTasksResponse
	(*GetTaskRequest)(nil),      // 3: tasks_proto.GetTaskRequest
	(*GetTaskResponse)(nil),     // 4: tasks_proto.GetTaskResponse
	(*AddTaskRequest)(nil),      // 5: tasks_proto.AddTaskRequest
	(*AddTaskResponse)(nil),     // 6: tasks_proto.AddTaskResponse
	(*UpdateTaskRequest)(nil),   // 7: tasks_proto.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),  // 8: tasks_proto.UpdateTaskResponse
	(*DeleteTaskRequest)(nil),   // 9: tasks_proto.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),  // 10: tasks_proto.DeleteTaskResponse
}
var file_tasks_proto_depIdxs = []int32{
	0,  // 0: tasks_proto.GetAllTasksResponse.tasks:type_name -> tasks_proto.Task
	0,  // 1: tasks_proto.GetTaskResponse.task:type_name -> tasks_proto.Task
	0,  // 2: tasks_proto.AddTaskRequest.task:type_name -> tasks_proto.Task
	0,  // 3: tasks_proto.UpdateTaskRequest.task:type_name -> tasks_proto.Task
	0,  // 4: tasks_proto.UpdateTaskResponse.task:type_name -> tasks_proto.Task
	1,  // 5: tasks_proto.TaskService.GetAllTasks:input_type -> tasks_proto.GetAllTasksRequest
	3,  // 6: tasks_proto.TaskService.GetUserAllTasks:input_type -> tasks_proto.GetTaskRequest
	3,  // 7: tasks_proto.TaskService.GetTask:input_type -> tasks_proto.GetTaskRequest
	5,  // 8: tasks_proto.TaskService.AddTask:input_type -> tasks_proto.AddTaskRequest
	7,  // 9: tasks_proto.TaskService.UpdateTask:input_type -> tasks_proto.UpdateTaskRequest
	9,  // 10: tasks_proto.TaskService.DeleteTask:input_type -> tasks_proto.DeleteTaskRequest
	2,  // 11: tasks_proto.TaskService.GetAllTasks:output_type -> tasks_proto.GetAllTasksResponse
	2,  // 12: tasks_proto.TaskService.GetUserAllTasks:output_type -> tasks_proto.GetAllTasksResponse
	4,  // 13: tasks_proto.TaskService.GetTask:output_type -> tasks_proto.GetTaskResponse
	6,  // 14: tasks_proto.TaskService.AddTask:output_type -> tasks_proto.AddTaskResponse
	8,  // 15: tasks_proto.TaskService.UpdateTask:output_type -> tasks_proto.UpdateTaskResponse
	10, // 16: tasks_proto.TaskService.DeleteTask:output_type -> tasks_proto.DeleteTaskResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_tasks_proto_init() }
func file_tasks_proto_init() {
	if File_tasks_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tasks_proto_rawDesc), len(file_tasks_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasks_proto_goTypes,
		DependencyIndexes: file_tasks_proto_depIdxs,
		MessageInfos:      file_tasks_proto_msgTypes,
	}.Build()
	File_tasks_proto = out.File
	file_tasks_proto_goTypes = nil
	file_tasks_proto_depIdxs = nil
}
