// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: v1/auth/auth.proto

package authv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	Role_ROLE_CLIENT      Role = 1
	Role_ROLE_ADMIN       Role = 2
	Role_ROLE_STAFF       Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_CLIENT",
		2: "ROLE_ADMIN",
		3: "ROLE_STAFF",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_CLIENT":      1,
		"ROLE_ADMIN":       2,
		"ROLE_STAFF":       3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_auth_auth_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_v1_auth_auth_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{0}
}

type UpdateStaffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CurrentEmail  string                 `protobuf:"bytes,1,opt,name=current_email,json=currentEmail,proto3" json:"current_email,omitempty"`
	NewWorkEmail  *string                `protobuf:"bytes,2,opt,name=new_work_email,json=newWorkEmail,proto3,oneof" json:"new_work_email,omitempty"`
	NewPosition   *string                `protobuf:"bytes,3,opt,name=new_position,json=newPosition,proto3,oneof" json:"new_position,omitempty"`
	NewPassword   *string                `protobuf:"bytes,4,opt,name=new_password,json=newPassword,proto3,oneof" json:"new_password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStaffRequest) Reset() {
	*x = UpdateStaffRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffRequest) ProtoMessage() {}

func (x *UpdateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaffRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStaffRequest) GetCurrentEmail() string {
	if x != nil {
		return x.CurrentEmail
	}
	return ""
}

func (x *UpdateStaffRequest) GetNewWorkEmail() string {
	if x != nil && x.NewWorkEmail != nil {
		return *x.NewWorkEmail
	}
	return ""
}

func (x *UpdateStaffRequest) GetNewPosition() string {
	if x != nil && x.NewPosition != nil {
		return *x.NewPosition
	}
	return ""
}

func (x *UpdateStaffRequest) GetNewPassword() string {
	if x != nil && x.NewPassword != nil {
		return *x.NewPassword
	}
	return ""
}

type UpdateStaffResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Staff         *Staff                 `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStaffResponse) Reset() {
	*x = UpdateStaffResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffResponse) ProtoMessage() {}

func (x *UpdateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffResponse.ProtoReflect.Descriptor instead.
func (*UpdateStaffResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStaffResponse) GetStaff() *Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

type RegisterStaffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Staff         *Staff                 `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterStaffRequest) Reset() {
	*x = RegisterStaffRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterStaffRequest) ProtoMessage() {}

func (x *RegisterStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterStaffRequest.ProtoReflect.Descriptor instead.
func (*RegisterStaffRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterStaffRequest) GetStaff() *Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

type RegisterStaffResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Staff         *Staff                 `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterStaffResponse) Reset() {
	*x = RegisterStaffResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterStaffResponse) ProtoMessage() {}

func (x *RegisterStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterStaffResponse.ProtoReflect.Descriptor instead.
func (*RegisterStaffResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterStaffResponse) GetStaff() *Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

type DeactivateStaffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkEmail     string                 `protobuf:"bytes,1,opt,name=work_email,json=workEmail,proto3" json:"work_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeactivateStaffRequest) Reset() {
	*x = DeactivateStaffRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateStaffRequest) ProtoMessage() {}

func (x *DeactivateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateStaffRequest.ProtoReflect.Descriptor instead.
func (*DeactivateStaffRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *DeactivateStaffRequest) GetWorkEmail() string {
	if x != nil {
		return x.WorkEmail
	}
	return ""
}

type DeactivateStaffResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkEmail     string                 `protobuf:"bytes,1,opt,name=work_email,json=workEmail,proto3" json:"work_email,omitempty"`
	Position      string                 `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeactivateStaffResponse) Reset() {
	*x = DeactivateStaffResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateStaffResponse) ProtoMessage() {}

func (x *DeactivateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateStaffResponse.ProtoReflect.Descriptor instead.
func (*DeactivateStaffResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *DeactivateStaffResponse) GetWorkEmail() string {
	if x != nil {
		return x.WorkEmail
	}
	return ""
}

func (x *DeactivateStaffResponse) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *DeactivateStaffResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type LoginClientInitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginClientInitRequest) Reset() {
	*x = LoginClientInitRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginClientInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginClientInitRequest) ProtoMessage() {}

func (x *LoginClientInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginClientInitRequest.ProtoReflect.Descriptor instead.
func (*LoginClientInitRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *LoginClientInitRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type LoginClientConfirmRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginClientConfirmRequest) Reset() {
	*x = LoginClientConfirmRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginClientConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginClientConfirmRequest) ProtoMessage() {}

func (x *LoginClientConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginClientConfirmRequest.ProtoReflect.Descriptor instead.
func (*LoginClientConfirmRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *LoginClientConfirmRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginClientConfirmRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginStaffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Staff         *Staff                 `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginStaffRequest) Reset() {
	*x = LoginStaffRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginStaffRequest) ProtoMessage() {}

func (x *LoginStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginStaffRequest.ProtoReflect.Descriptor instead.
func (*LoginStaffRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *LoginStaffRequest) GetStaff() *Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

func (x *LoginStaffRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type OAuthYandexRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri   string                 `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuthYandexRequest) Reset() {
	*x = OAuthYandexRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthYandexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthYandexRequest) ProtoMessage() {}

func (x *OAuthYandexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthYandexRequest.ProtoReflect.Descriptor instead.
func (*OAuthYandexRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{9}
}

func (x *OAuthYandexRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OAuthYandexRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type LoginInitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginInitResponse) Reset() {
	*x = LoginInitResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInitResponse) ProtoMessage() {}

func (x *LoginInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInitResponse.ProtoReflect.Descriptor instead.
func (*LoginInitResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{10}
}

func (x *LoginInitResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LoginInitResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	AccessToken           string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn             int64                  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int64                  `protobuf:"varint,4,opt,name=refresh_token_expires_in,json=refreshTokenExpiresIn,proto3" json:"refresh_token_expires_in,omitempty"`
	User                  *User                  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{11}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshTokenExpiresIn() int64 {
	if x != nil {
		return x.RefreshTokenExpiresIn
	}
	return 0
}

func (x *LoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ValidateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{12}
}

func (x *ValidateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Valid         bool                   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	mi := &file_v1_auth_auth_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{13}
}

func (x *ValidateResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type RefreshRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	mi := &file_v1_auth_auth_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{14}
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to UserType:
	//
	//	*User_Client
	//	*User_Staff
	UserType      isUser_UserType `protobuf_oneof:"user_type"`
	Roles         []Role          `protobuf:"varint,4,rep,packed,name=roles,proto3,enum=auth.Role" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_v1_auth_auth_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[15]
	if x != nil {
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
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{15}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUserType() isUser_UserType {
	if x != nil {
		return x.UserType
	}
	return nil
}

func (x *User) GetClient() *Client {
	if x != nil {
		if x, ok := x.UserType.(*User_Client); ok {
			return x.Client
		}
	}
	return nil
}

func (x *User) GetStaff() *Staff {
	if x != nil {
		if x, ok := x.UserType.(*User_Staff); ok {
			return x.Staff
		}
	}
	return nil
}

func (x *User) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type isUser_UserType interface {
	isUser_UserType()
}

type User_Client struct {
	Client *Client `protobuf:"bytes,2,opt,name=client,proto3,oneof"`
}

type User_Staff struct {
	Staff *Staff `protobuf:"bytes,3,opt,name=staff,proto3,oneof"`
}

func (*User_Client) isUser_UserType() {}

func (*User_Staff) isUser_UserType() {}

type Client struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Client) Reset() {
	*x = Client{}
	mi := &file_v1_auth_auth_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{16}
}

func (x *Client) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Client) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Staff struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkEmail     string                 `protobuf:"bytes,1,opt,name=work_email,json=workEmail,proto3" json:"work_email,omitempty"`
	Position      string                 `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Staff) Reset() {
	*x = Staff{}
	mi := &file_v1_auth_auth_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Staff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Staff) ProtoMessage() {}

func (x *Staff) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_auth_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Staff.ProtoReflect.Descriptor instead.
func (*Staff) Descriptor() ([]byte, []int) {
	return file_v1_auth_auth_proto_rawDescGZIP(), []int{17}
}

func (x *Staff) GetWorkEmail() string {
	if x != nil {
		return x.WorkEmail
	}
	return ""
}

func (x *Staff) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

var File_v1_auth_auth_proto protoreflect.FileDescriptor

const file_v1_auth_auth_proto_rawDesc = "" +
	"\n" +
	"\x12v1/auth/auth.proto\x12\x04auth\x1a\x1cgoogle/api/annotations.proto\"\xe9\x01\n" +
	"\x12UpdateStaffRequest\x12#\n" +
	"\rcurrent_email\x18\x01 \x01(\tR\fcurrentEmail\x12)\n" +
	"\x0enew_work_email\x18\x02 \x01(\tH\x00R\fnewWorkEmail\x88\x01\x01\x12&\n" +
	"\fnew_position\x18\x03 \x01(\tH\x01R\vnewPosition\x88\x01\x01\x12&\n" +
	"\fnew_password\x18\x04 \x01(\tH\x02R\vnewPassword\x88\x01\x01B\x11\n" +
	"\x0f_new_work_emailB\x0f\n" +
	"\r_new_positionB\x0f\n" +
	"\r_new_password\"8\n" +
	"\x13UpdateStaffResponse\x12!\n" +
	"\x05staff\x18\x01 \x01(\v2\v.auth.StaffR\x05staff\"9\n" +
	"\x14RegisterStaffRequest\x12!\n" +
	"\x05staff\x18\x01 \x01(\v2\v.auth.StaffR\x05staff\":\n" +
	"\x15RegisterStaffResponse\x12!\n" +
	"\x05staff\x18\x01 \x01(\v2\v.auth.StaffR\x05staff\"7\n" +
	"\x16DeactivateStaffRequest\x12\x1d\n" +
	"\n" +
	"work_email\x18\x01 \x01(\tR\tworkEmail\"l\n" +
	"\x17DeactivateStaffResponse\x12\x1d\n" +
	"\n" +
	"work_email\x18\x01 \x01(\tR\tworkEmail\x12\x1a\n" +
	"\bposition\x18\x02 \x01(\tR\bposition\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\".\n" +
	"\x16LoginClientInitRequest\x12\x14\n" +
	"\x05phone\x18\x01 \x01(\tR\x05phone\"E\n" +
	"\x19LoginClientConfirmRequest\x12\x14\n" +
	"\x05phone\x18\x01 \x01(\tR\x05phone\x12\x12\n" +
	"\x04code\x18\x02 \x01(\tR\x04code\"R\n" +
	"\x11LoginStaffRequest\x12!\n" +
	"\x05staff\x18\x01 \x01(\v2\v.auth.StaffR\x05staff\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"K\n" +
	"\x12OAuthYandexRequest\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12!\n" +
	"\fredirect_uri\x18\x02 \x01(\tR\vredirectUri\"A\n" +
	"\x11LoginInitResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error\"\xcf\x01\n" +
	"\rLoginResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12\x1d\n" +
	"\n" +
	"expires_in\x18\x02 \x01(\x03R\texpiresIn\x12#\n" +
	"\rrefresh_token\x18\x03 \x01(\tR\frefreshToken\x127\n" +
	"\x18refresh_token_expires_in\x18\x04 \x01(\x03R\x15refreshTokenExpiresIn\x12\x1e\n" +
	"\x04user\x18\x05 \x01(\v2\n" +
	".auth.UserR\x04user\"'\n" +
	"\x0fValidateRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"H\n" +
	"\x10ValidateResponse\x12\x14\n" +
	"\x05valid\x18\x01 \x01(\bR\x05valid\x12\x1e\n" +
	"\x04user\x18\x02 \x01(\v2\n" +
	".auth.UserR\x04user\"5\n" +
	"\x0eRefreshRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"\x92\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12&\n" +
	"\x06client\x18\x02 \x01(\v2\f.auth.ClientH\x00R\x06client\x12#\n" +
	"\x05staff\x18\x03 \x01(\v2\v.auth.StaffH\x00R\x05staff\x12 \n" +
	"\x05roles\x18\x04 \x03(\x0e2\n" +
	".auth.RoleR\x05rolesB\v\n" +
	"\tuser_type\"4\n" +
	"\x06Client\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x02 \x01(\tR\x05phone\"B\n" +
	"\x05Staff\x12\x1d\n" +
	"\n" +
	"work_email\x18\x01 \x01(\tR\tworkEmail\x12\x1a\n" +
	"\bposition\x18\x02 \x01(\tR\bposition*M\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vROLE_CLIENT\x10\x01\x12\x0e\n" +
	"\n" +
	"ROLE_ADMIN\x10\x02\x12\x0e\n" +
	"\n" +
	"ROLE_STAFF\x10\x032\xaa\a\n" +
	"\x04Auth\x12o\n" +
	"\x0fLoginClientInit\x12\x1c.auth.LoginClientInitRequest\x1a\x17.auth.LoginInitResponse\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/v1/auth/client/login/init\x12t\n" +
	"\x12LoginClientConfirm\x12\x1f.auth.LoginClientConfirmRequest\x1a\x13.auth.LoginResponse\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/v1/auth/client/login/confirm\x12l\n" +
	"\rRegisterStaff\x12\x1a.auth.RegisterStaffRequest\x1a\x1b.auth.RegisterStaffResponse\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/v1/auth/staff/register\x12d\n" +
	"\vUpdateStaff\x12\x18.auth.UpdateStaffRequest\x1a\x19.auth.UpdateStaffResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\x1a\x15/v1/auth/staff/update\x12[\n" +
	"\n" +
	"LoginStaff\x12\x17.auth.LoginStaffRequest\x1a\x13.auth.LoginResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v1/auth/staff/login\x12~\n" +
	"\x0fDeactivateStaff\x12\x1c.auth.DeactivateStaffRequest\x1a\x1d.auth.DeactivateStaffResponse\".\x82\xd3\xe4\x93\x02(*&/v1/auth/staff/deactivate/{work_email}\x12^\n" +
	"\vLoginYandex\x12\x18.auth.OAuthYandexRequest\x1a\x13.auth.LoginResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/v1/auth/yandex/login\x12W\n" +
	"\bValidate\x12\x15.auth.ValidateRequest\x1a\x16.auth.ValidateResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v1/auth/validate\x12Q\n" +
	"\aRefresh\x12\x14.auth.RefreshRequest\x1a\x13.auth.LoginResponse\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/v1/auth/refreshB\x1cZ\x1anetscrawler.auth.v1;authv1b\x06proto3"

var (
	file_v1_auth_auth_proto_rawDescOnce sync.Once
	file_v1_auth_auth_proto_rawDescData []byte
)

func file_v1_auth_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_auth_auth_proto_rawDesc), len(file_v1_auth_auth_proto_rawDesc)))
	})
	return file_v1_auth_auth_proto_rawDescData
}

var file_v1_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_v1_auth_auth_proto_goTypes = []any{
	(Role)(0),                         // 0: auth.Role
	(*UpdateStaffRequest)(nil),        // 1: auth.UpdateStaffRequest
	(*UpdateStaffResponse)(nil),       // 2: auth.UpdateStaffResponse
	(*RegisterStaffRequest)(nil),      // 3: auth.RegisterStaffRequest
	(*RegisterStaffResponse)(nil),     // 4: auth.RegisterStaffResponse
	(*DeactivateStaffRequest)(nil),    // 5: auth.DeactivateStaffRequest
	(*DeactivateStaffResponse)(nil),   // 6: auth.DeactivateStaffResponse
	(*LoginClientInitRequest)(nil),    // 7: auth.LoginClientInitRequest
	(*LoginClientConfirmRequest)(nil), // 8: auth.LoginClientConfirmRequest
	(*LoginStaffRequest)(nil),         // 9: auth.LoginStaffRequest
	(*OAuthYandexRequest)(nil),        // 10: auth.OAuthYandexRequest
	(*LoginInitResponse)(nil),         // 11: auth.LoginInitResponse
	(*LoginResponse)(nil),             // 12: auth.LoginResponse
	(*ValidateRequest)(nil),           // 13: auth.ValidateRequest
	(*ValidateResponse)(nil),          // 14: auth.ValidateResponse
	(*RefreshRequest)(nil),            // 15: auth.RefreshRequest
	(*User)(nil),                      // 16: auth.User
	(*Client)(nil),                    // 17: auth.Client
	(*Staff)(nil),                     // 18: auth.Staff
}
var file_v1_auth_auth_proto_depIdxs = []int32{
	18, // 0: auth.UpdateStaffResponse.staff:type_name -> auth.Staff
	18, // 1: auth.RegisterStaffRequest.staff:type_name -> auth.Staff
	18, // 2: auth.RegisterStaffResponse.staff:type_name -> auth.Staff
	18, // 3: auth.LoginStaffRequest.staff:type_name -> auth.Staff
	16, // 4: auth.LoginResponse.user:type_name -> auth.User
	16, // 5: auth.ValidateResponse.user:type_name -> auth.User
	17, // 6: auth.User.client:type_name -> auth.Client
	18, // 7: auth.User.staff:type_name -> auth.Staff
	0,  // 8: auth.User.roles:type_name -> auth.Role
	7,  // 9: auth.Auth.LoginClientInit:input_type -> auth.LoginClientInitRequest
	8,  // 10: auth.Auth.LoginClientConfirm:input_type -> auth.LoginClientConfirmRequest
	3,  // 11: auth.Auth.RegisterStaff:input_type -> auth.RegisterStaffRequest
	1,  // 12: auth.Auth.UpdateStaff:input_type -> auth.UpdateStaffRequest
	9,  // 13: auth.Auth.LoginStaff:input_type -> auth.LoginStaffRequest
	5,  // 14: auth.Auth.DeactivateStaff:input_type -> auth.DeactivateStaffRequest
	10, // 15: auth.Auth.LoginYandex:input_type -> auth.OAuthYandexRequest
	13, // 16: auth.Auth.Validate:input_type -> auth.ValidateRequest
	15, // 17: auth.Auth.Refresh:input_type -> auth.RefreshRequest
	11, // 18: auth.Auth.LoginClientInit:output_type -> auth.LoginInitResponse
	12, // 19: auth.Auth.LoginClientConfirm:output_type -> auth.LoginResponse
	4,  // 20: auth.Auth.RegisterStaff:output_type -> auth.RegisterStaffResponse
	2,  // 21: auth.Auth.UpdateStaff:output_type -> auth.UpdateStaffResponse
	12, // 22: auth.Auth.LoginStaff:output_type -> auth.LoginResponse
	6,  // 23: auth.Auth.DeactivateStaff:output_type -> auth.DeactivateStaffResponse
	12, // 24: auth.Auth.LoginYandex:output_type -> auth.LoginResponse
	14, // 25: auth.Auth.Validate:output_type -> auth.ValidateResponse
	12, // 26: auth.Auth.Refresh:output_type -> auth.LoginResponse
	18, // [18:27] is the sub-list for method output_type
	9,  // [9:18] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_v1_auth_auth_proto_init() }
func file_v1_auth_auth_proto_init() {
	if File_v1_auth_auth_proto != nil {
		return
	}
	file_v1_auth_auth_proto_msgTypes[0].OneofWrappers = []any{}
	file_v1_auth_auth_proto_msgTypes[15].OneofWrappers = []any{
		(*User_Client)(nil),
		(*User_Staff)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_auth_auth_proto_rawDesc), len(file_v1_auth_auth_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_auth_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_auth_proto_depIdxs,
		EnumInfos:         file_v1_auth_auth_proto_enumTypes,
		MessageInfos:      file_v1_auth_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_auth_proto = out.File
	file_v1_auth_auth_proto_goTypes = nil
	file_v1_auth_auth_proto_depIdxs = nil
}
