// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: user.proto

package user

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 会员等级
type MemberLevel int32

const (
	MemberLevel_ordinaryMember MemberLevel = 0 // 普通用户 0-30
	MemberLevel_bronzeMember   MemberLevel = 1 // 青铜会员 31-60
	MemberLevel_silverMember   MemberLevel = 2 // 白银会员 61-90
	MemberLevel_goldMember     MemberLevel = 3 // 黄金会员 91-120
	MemberLevel_platinumMember MemberLevel = 4 // 铂金会员 121-150
	MemberLevel_diamondsMember MemberLevel = 5 // 钻石会员 151-180
)

// Enum value maps for MemberLevel.
var (
	MemberLevel_name = map[int32]string{
		0: "ordinaryMember",
		1: "bronzeMember",
		2: "silverMember",
		3: "goldMember",
		4: "platinumMember",
		5: "diamondsMember",
	}
	MemberLevel_value = map[string]int32{
		"ordinaryMember": 0,
		"bronzeMember":   1,
		"silverMember":   2,
		"goldMember":     3,
		"platinumMember": 4,
		"diamondsMember": 5,
	}
)

func (x MemberLevel) Enum() *MemberLevel {
	p := new(MemberLevel)
	*p = x
	return p
}

func (x MemberLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (MemberLevel) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x MemberLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberLevel.Descriptor instead.
func (MemberLevel) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

// user base info
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string      `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`                                        // 用户id uuid
	NiceName     string      `protobuf:"bytes,2,opt,name=niceName,proto3" json:"niceName,omitempty"`                                    // 用户昵称
	Sex          bool        `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`                                             // 用户性别0 女 1 男
	Avatar       string      `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                                        // 用户头像
	UserLevel    MemberLevel `protobuf:"varint,5,opt,name=userLevel,proto3,enum=micro.mJu.user.MemberLevel" json:"userLevel,omitempty"` // 用户等级 根据积分升级
	Integral     int32       `protobuf:"varint,6,opt,name=integral,proto3" json:"integral,omitempty"`                                   // 用户当前积分
	PhoneNumber  string      `protobuf:"bytes,7,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`                              // 用户手机号
	PlayerScript int32       `protobuf:"varint,8,opt,name=playerScript,proto3" json:"playerScript,omitempty"`                           // 用户打本数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetNiceName() string {
	if x != nil {
		return x.NiceName
	}
	return ""
}

func (x *User) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetUserLevel() MemberLevel {
	if x != nil {
		return x.UserLevel
	}
	return MemberLevel_ordinaryMember
}

func (x *User) GetIntegral() int32 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetPlayerScript() int32 {
	if x != nil {
		return x.PlayerScript
	}
	return 0
}

// getUserInfosReq
type GetUserInfosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId []string `protobuf:"bytes,1,rep,name=userId,proto3" json:"userId,omitempty"` // userId
}

func (x *GetUserInfosReq) Reset() {
	*x = GetUserInfosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfosReq) ProtoMessage() {}

func (x *GetUserInfosReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfosReq.ProtoReflect.Descriptor instead.
func (*GetUserInfosReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfosReq) GetUserId() []string {
	if x != nil {
		return x.UserId
	}
	return nil
}

// getUserInfosRsp
type GetUserInfosRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RspCode   int32   `protobuf:"varint,1,opt,name=rspCode,proto3" json:"rspCode,omitempty"`
	Message   string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserInfos []*User `protobuf:"bytes,3,rep,name=userInfos,proto3" json:"userInfos,omitempty"`
}

func (x *GetUserInfosRsp) Reset() {
	*x = GetUserInfosRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfosRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfosRsp) ProtoMessage() {}

func (x *GetUserInfosRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfosRsp.ProtoReflect.Descriptor instead.
func (*GetUserInfosRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfosRsp) GetRspCode() int32 {
	if x != nil {
		return x.RspCode
	}
	return 0
}

func (x *GetUserInfosRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUserInfosRsp) GetUserInfos() []*User {
	if x != nil {
		return x.UserInfos
	}
	return nil
}

// UpdateInfosRep
type UpdateInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`    // userId
	UserInfo *User `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"` // 修改用户信息 客户端修改用户信息需要先getUserInfo 传全量信息
}

func (x *UpdateInfoReq) Reset() {
	*x = UpdateInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoReq) ProtoMessage() {}

func (x *UpdateInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateInfoReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateInfoReq) GetUserInfo() *User {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

// UpdateInfosRsp
type UpdateInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResCode int32  `protobuf:"varint,1,opt,name=resCode,proto3" json:"resCode,omitempty"` // resCode != 0 error != nil  resCode == 0 error == nil
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateInfoRsp) Reset() {
	*x = UpdateInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoRsp) ProtoMessage() {}

func (x *UpdateInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoRsp.ProtoReflect.Descriptor instead.
func (*UpdateInfoRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInfoRsp) GetResCode() int32 {
	if x != nil {
		return x.ResCode
	}
	return 0
}

func (x *UpdateInfoRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// DelUserInfoReq
type DelUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DelUserInfoReq) Reset() {
	*x = DelUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserInfoReq) ProtoMessage() {}

func (x *DelUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserInfoReq.ProtoReflect.Descriptor instead.
func (*DelUserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *DelUserInfoReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// DelUserInfoRsp
type DelUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResCode int32  `protobuf:"varint,1,opt,name=resCode,proto3" json:"resCode,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DelUserInfoRsp) Reset() {
	*x = DelUserInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserInfoRsp) ProtoMessage() {}

func (x *DelUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserInfoRsp.ProtoReflect.Descriptor instead.
func (*DelUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *DelUserInfoRsp) GetResCode() int32 {
	if x != nil {
		return x.ResCode
	}
	return 0
}

func (x *DelUserInfoRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a, 0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x81, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d,
	0x4a, 0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a,
	0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a, 0x75, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x43, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x44, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x7d, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x62, 0x72, 0x6f,
	0x6e, 0x7a, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x73,
	0x69, 0x6c, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x67, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x70, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10,
	0x04, 0x12, 0x12, 0x0a, 0x0e, 0x64, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x73, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x10, 0x05, 0x32, 0xfa, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a,
	0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d,
	0x4a, 0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x52, 0x73, 0x70, 0x12, 0x4a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a,
	0x75, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a, 0x75,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a, 0x75, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x4a, 0x75, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_proto_goTypes = []interface{}{
	(MemberLevel)(0),        // 0: micro.mJu.user.memberLevel
	(*User)(nil),            // 1: micro.mJu.user.User
	(*GetUserInfosReq)(nil), // 2: micro.mJu.user.GetUserInfosReq
	(*GetUserInfosRsp)(nil), // 3: micro.mJu.user.GetUserInfosRsp
	(*UpdateInfoReq)(nil),   // 4: micro.mJu.user.UpdateInfoReq
	(*UpdateInfoRsp)(nil),   // 5: micro.mJu.user.UpdateInfoRsp
	(*DelUserInfoReq)(nil),  // 6: micro.mJu.user.DelUserInfoReq
	(*DelUserInfoRsp)(nil),  // 7: micro.mJu.user.DelUserInfoRsp
}
var file_user_proto_depIdxs = []int32{
	0, // 0: micro.mJu.user.User.userLevel:type_name -> micro.mJu.user.memberLevel
	1, // 1: micro.mJu.user.GetUserInfosRsp.userInfos:type_name -> micro.mJu.user.User
	1, // 2: micro.mJu.user.UpdateInfoReq.userInfo:type_name -> micro.mJu.user.User
	2, // 3: micro.mJu.user.UserService.GetUserInfos:input_type -> micro.mJu.user.GetUserInfosReq
	4, // 4: micro.mJu.user.UserService.UpdateInfo:input_type -> micro.mJu.user.UpdateInfoReq
	6, // 5: micro.mJu.user.UserService.DelUserInfo:input_type -> micro.mJu.user.DelUserInfoReq
	3, // 6: micro.mJu.user.UserService.GetUserInfos:output_type -> micro.mJu.user.GetUserInfosRsp
	5, // 7: micro.mJu.user.UserService.UpdateInfo:output_type -> micro.mJu.user.UpdateInfoRsp
	7, // 8: micro.mJu.user.UserService.DelUserInfo:output_type -> micro.mJu.user.DelUserInfoRsp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfosReq); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfosRsp); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoReq); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoRsp); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserInfoReq); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserInfoRsp); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
