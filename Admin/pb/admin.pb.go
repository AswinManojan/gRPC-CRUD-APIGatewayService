// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: Admin/pb/admin.proto

package pb

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

type ProviderDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *ProviderDataRequest) Reset() {
	*x = ProviderDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Admin_pb_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderDataRequest) ProtoMessage() {}

func (x *ProviderDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_pb_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderDataRequest.ProtoReflect.Descriptor instead.
func (*ProviderDataRequest) Descriptor() ([]byte, []int) {
	return file_Admin_pb_admin_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderDataRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ProviderDataRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ProviderDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Admin_pb_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_pb_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_Admin_pb_admin_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Admin_pb_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_pb_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_Admin_pb_admin_proto_rawDescGZIP(), []int{2}
}

func (x *IdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Admin_pb_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_pb_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_Admin_pb_admin_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Result) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *UserDataRequest) Reset() {
	*x = UserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Admin_pb_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataRequest) ProtoMessage() {}

func (x *UserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_pb_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataRequest.ProtoReflect.Descriptor instead.
func (*UserDataRequest) Descriptor() ([]byte, []int) {
	return file_Admin_pb_admin_proto_rawDescGZIP(), []int{4}
}

func (x *UserDataRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDataRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_Admin_pb_admin_proto protoreflect.FileDescriptor

var file_Admin_pb_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x63, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xba, 0x03, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29,
	0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d,
	0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Admin_pb_admin_proto_rawDescOnce sync.Once
	file_Admin_pb_admin_proto_rawDescData = file_Admin_pb_admin_proto_rawDesc
)

func file_Admin_pb_admin_proto_rawDescGZIP() []byte {
	file_Admin_pb_admin_proto_rawDescOnce.Do(func() {
		file_Admin_pb_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_Admin_pb_admin_proto_rawDescData)
	})
	return file_Admin_pb_admin_proto_rawDescData
}

var file_Admin_pb_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Admin_pb_admin_proto_goTypes = []interface{}{
	(*ProviderDataRequest)(nil), // 0: pb.ProviderDataRequest
	(*LoginRequest)(nil),        // 1: pb.LoginRequest
	(*IdRequest)(nil),           // 2: pb.IdRequest
	(*Result)(nil),              // 3: pb.result
	(*UserDataRequest)(nil),     // 4: pb.UserDataRequest
}
var file_Admin_pb_admin_proto_depIdxs = []int32{
	1, // 0: pb.AdminService.AdminLogin:input_type -> pb.LoginRequest
	4, // 1: pb.AdminService.EditUser:input_type -> pb.UserDataRequest
	2, // 2: pb.AdminService.DeleteUserById:input_type -> pb.IdRequest
	2, // 3: pb.AdminService.FindUserById:input_type -> pb.IdRequest
	4, // 4: pb.AdminService.CreateUser:input_type -> pb.UserDataRequest
	0, // 5: pb.AdminService.EditProvider:input_type -> pb.ProviderDataRequest
	2, // 6: pb.AdminService.DeleteProviderById:input_type -> pb.IdRequest
	2, // 7: pb.AdminService.FindProviderById:input_type -> pb.IdRequest
	0, // 8: pb.AdminService.CreateProvider:input_type -> pb.ProviderDataRequest
	3, // 9: pb.AdminService.AdminLogin:output_type -> pb.result
	3, // 10: pb.AdminService.EditUser:output_type -> pb.result
	3, // 11: pb.AdminService.DeleteUserById:output_type -> pb.result
	3, // 12: pb.AdminService.FindUserById:output_type -> pb.result
	3, // 13: pb.AdminService.CreateUser:output_type -> pb.result
	3, // 14: pb.AdminService.EditProvider:output_type -> pb.result
	3, // 15: pb.AdminService.DeleteProviderById:output_type -> pb.result
	3, // 16: pb.AdminService.FindProviderById:output_type -> pb.result
	3, // 17: pb.AdminService.CreateProvider:output_type -> pb.result
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Admin_pb_admin_proto_init() }
func file_Admin_pb_admin_proto_init() {
	if File_Admin_pb_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Admin_pb_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderDataRequest); i {
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
		file_Admin_pb_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_Admin_pb_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_Admin_pb_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_Admin_pb_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataRequest); i {
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
			RawDescriptor: file_Admin_pb_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Admin_pb_admin_proto_goTypes,
		DependencyIndexes: file_Admin_pb_admin_proto_depIdxs,
		MessageInfos:      file_Admin_pb_admin_proto_msgTypes,
	}.Build()
	File_Admin_pb_admin_proto = out.File
	file_Admin_pb_admin_proto_rawDesc = nil
	file_Admin_pb_admin_proto_goTypes = nil
	file_Admin_pb_admin_proto_depIdxs = nil
}
