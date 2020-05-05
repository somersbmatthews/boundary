// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: controller/storage/iam/store/v1/role.proto

package store

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is used to access the Role via an API
	// @inject_tag: gorm:"primary_key"
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// FriendlyName is the optional friendly name used to
	// access the Role via an API
	// @inject_tag: `gorm:"default:null"`
	FriendlyName string `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty" gorm:"default:null"`
	Description  string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// primary scope id of the PrimaryScope
	// @inject_tag: `gorm:"default:null"`
	PrimaryScopeId string `protobuf:"bytes,7,opt,name=primary_scope_id,json=primaryScopeId,proto3" json:"primary_scope_id,omitempty" gorm:"default:null"`
	// @inject_tag: gorm:"foreignkey:ScopeId"
	PrimaryScope *Scope `protobuf:"bytes,8,opt,name=primary_scope,json=primaryScope,proto3" json:"primary_scope,omitempty" gorm:"foreignkey:ScopeId"`
	// disabled is by default false and allows a Role to be marked disabled.
	Disabled bool `protobuf:"varint,9,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Role) GetCreateTime() *Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Role) GetUpdateTime() *Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Role) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPrimaryScopeId() string {
	if x != nil {
		return x.PrimaryScopeId
	}
	return ""
}

func (x *Role) GetPrimaryScope() *Scope {
	if x != nil {
		return x.PrimaryScope
	}
	return nil
}

func (x *Role) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// public_id is used to access the UserRole via an API
	PublicId string `protobuf:"bytes,4,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	// FriendlyName is the optional friendly name used to
	// access the UserRole via an API
	// @inject_tag: `gorm:"default:null"`
	FriendlyName string `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty" gorm:"default:null"`
	// primary scope id of the PrimaryScope
	// @inject_tag: `gorm:"default:null"`
	PrimaryScopeId string `protobuf:"bytes,6,opt,name=primary_scope_id,json=primaryScopeId,proto3" json:"primary_scope_id,omitempty" gorm:"default:null"`
	// @inject_tag: gorm:"foreignkey:ScopeId"
	PrimaryScope *Scope `protobuf:"bytes,7,opt,name=primary_scope,json=primaryScope,proto3" json:"primary_scope,omitempty" gorm:"foreignkey:ScopeId"`
	// role_id is the role of this principal.
	RoleId string `protobuf:"bytes,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// Principal type (User)
	Type        string `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	PrincipalId string `protobuf:"bytes,11,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_proto_rawDescGZIP(), []int{1}
}

func (x *UserRole) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRole) GetCreateTime() *Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *UserRole) GetUpdateTime() *Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *UserRole) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *UserRole) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *UserRole) GetPrimaryScopeId() string {
	if x != nil {
		return x.PrimaryScopeId
	}
	return ""
}

func (x *UserRole) GetPrimaryScope() *Scope {
	if x != nil {
		return x.PrimaryScope
	}
	return nil
}

func (x *UserRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UserRole) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserRole) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

type GroupRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// public_id is used to access the GroupRole via an API
	PublicId string `protobuf:"bytes,4,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	// FriendlyName is the optional friendly name used to
	// access the GroupRole via an API
	// @inject_tag: `gorm:"default:null"`
	FriendlyName string `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty" gorm:"default:null"`
	// primary scope id of the PrimaryScope
	// @inject_tag: `gorm:"default:null"`
	PrimaryScopeId string `protobuf:"bytes,6,opt,name=primary_scope_id,json=primaryScopeId,proto3" json:"primary_scope_id,omitempty" gorm:"default:null"`
	// @inject_tag: gorm:"foreignkey:ScopeId"
	PrimaryScope *Scope `protobuf:"bytes,7,opt,name=primary_scope,json=primaryScope,proto3" json:"primary_scope,omitempty" gorm:"foreignkey:ScopeId"`
	// role_id is the role of this principal.
	RoleId string `protobuf:"bytes,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// Principal type (Group)
	Type        string `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	PrincipalId string `protobuf:"bytes,11,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *GroupRole) Reset() {
	*x = GroupRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRole) ProtoMessage() {}

func (x *GroupRole) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRole.ProtoReflect.Descriptor instead.
func (*GroupRole) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_proto_rawDescGZIP(), []int{2}
}

func (x *GroupRole) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupRole) GetCreateTime() *Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GroupRole) GetUpdateTime() *Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GroupRole) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *GroupRole) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *GroupRole) GetPrimaryScopeId() string {
	if x != nil {
		return x.PrimaryScopeId
	}
	return ""
}

func (x *GroupRole) GetPrimaryScope() *Scope {
	if x != nil {
		return x.PrimaryScope
	}
	return nil
}

func (x *GroupRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *GroupRole) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GroupRole) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

type AssignedRoleView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// update_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// public_id is used to access the PrincipalRole via an API
	PublicId string `protobuf:"bytes,4,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	// FriendlyName is the optional friendly name used to
	// access the PrincipalRole via an API
	// @inject_tag: `gorm:"default:null"`
	FriendlyName string `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty" gorm:"default:null"`
	// primary scope id of the PrimaryScope
	// @inject_tag: `gorm:"default:null"`
	PrimaryScopeId string `protobuf:"bytes,6,opt,name=primary_scope_id,json=primaryScopeId,proto3" json:"primary_scope_id,omitempty" gorm:"default:null"`
	// @inject_tag: gorm:"foreignkey:ScopeId"
	PrimaryScope *Scope `protobuf:"bytes,7,opt,name=primary_scope,json=primaryScope,proto3" json:"primary_scope,omitempty" gorm:"foreignkey:ScopeId"`
	// role_id is the role of this principal.
	RoleId string `protobuf:"bytes,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// Principal type (User or Group)
	Type        string `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	PrincipalId string `protobuf:"bytes,11,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *AssignedRoleView) Reset() {
	*x = AssignedRoleView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignedRoleView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignedRoleView) ProtoMessage() {}

func (x *AssignedRoleView) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignedRoleView.ProtoReflect.Descriptor instead.
func (*AssignedRoleView) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_proto_rawDescGZIP(), []int{3}
}

func (x *AssignedRoleView) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AssignedRoleView) GetCreateTime() *Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AssignedRoleView) GetUpdateTime() *Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *AssignedRoleView) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *AssignedRoleView) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *AssignedRoleView) GetPrimaryScopeId() string {
	if x != nil {
		return x.PrimaryScopeId
	}
	return ""
}

func (x *AssignedRoleView) GetPrimaryScope() *Scope {
	if x != nil {
		return x.PrimaryScope
	}
	return nil
}

func (x *AssignedRoleView) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AssignedRoleView) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AssignedRoleView) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

var File_controller_storage_iam_store_v1_role_proto protoreflect.FileDescriptor

var file_controller_storage_iam_store_v1_role_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0d,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xbd, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xbe, 0x03, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0d, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xc5, 0x03, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x4b, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74, 0x6f, 0x77, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_controller_storage_iam_store_v1_role_proto_rawDescOnce sync.Once
	file_controller_storage_iam_store_v1_role_proto_rawDescData = file_controller_storage_iam_store_v1_role_proto_rawDesc
)

func file_controller_storage_iam_store_v1_role_proto_rawDescGZIP() []byte {
	file_controller_storage_iam_store_v1_role_proto_rawDescOnce.Do(func() {
		file_controller_storage_iam_store_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_iam_store_v1_role_proto_rawDescData)
	})
	return file_controller_storage_iam_store_v1_role_proto_rawDescData
}

var file_controller_storage_iam_store_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_storage_iam_store_v1_role_proto_goTypes = []interface{}{
	(*Role)(nil),             // 0: controller.storage.iam.store.v1.Role
	(*UserRole)(nil),         // 1: controller.storage.iam.store.v1.UserRole
	(*GroupRole)(nil),        // 2: controller.storage.iam.store.v1.GroupRole
	(*AssignedRoleView)(nil), // 3: controller.storage.iam.store.v1.AssignedRoleView
	(*Timestamp)(nil),        // 4: controller.storage.iam.store.v1.Timestamp
	(*Scope)(nil),            // 5: controller.storage.iam.store.v1.Scope
}
var file_controller_storage_iam_store_v1_role_proto_depIdxs = []int32{
	4,  // 0: controller.storage.iam.store.v1.Role.create_time:type_name -> controller.storage.iam.store.v1.Timestamp
	4,  // 1: controller.storage.iam.store.v1.Role.update_time:type_name -> controller.storage.iam.store.v1.Timestamp
	5,  // 2: controller.storage.iam.store.v1.Role.primary_scope:type_name -> controller.storage.iam.store.v1.Scope
	4,  // 3: controller.storage.iam.store.v1.UserRole.create_time:type_name -> controller.storage.iam.store.v1.Timestamp
	4,  // 4: controller.storage.iam.store.v1.UserRole.update_time:type_name -> controller.storage.iam.store.v1.Timestamp
	5,  // 5: controller.storage.iam.store.v1.UserRole.primary_scope:type_name -> controller.storage.iam.store.v1.Scope
	4,  // 6: controller.storage.iam.store.v1.GroupRole.create_time:type_name -> controller.storage.iam.store.v1.Timestamp
	4,  // 7: controller.storage.iam.store.v1.GroupRole.update_time:type_name -> controller.storage.iam.store.v1.Timestamp
	5,  // 8: controller.storage.iam.store.v1.GroupRole.primary_scope:type_name -> controller.storage.iam.store.v1.Scope
	4,  // 9: controller.storage.iam.store.v1.AssignedRoleView.create_time:type_name -> controller.storage.iam.store.v1.Timestamp
	4,  // 10: controller.storage.iam.store.v1.AssignedRoleView.update_time:type_name -> controller.storage.iam.store.v1.Timestamp
	5,  // 11: controller.storage.iam.store.v1.AssignedRoleView.primary_scope:type_name -> controller.storage.iam.store.v1.Scope
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_controller_storage_iam_store_v1_role_proto_init() }
func file_controller_storage_iam_store_v1_role_proto_init() {
	if File_controller_storage_iam_store_v1_role_proto != nil {
		return
	}
	file_controller_storage_iam_store_v1_timestamp_proto_init()
	file_controller_storage_iam_store_v1_scope_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_iam_store_v1_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_controller_storage_iam_store_v1_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_controller_storage_iam_store_v1_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRole); i {
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
		file_controller_storage_iam_store_v1_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignedRoleView); i {
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
			RawDescriptor: file_controller_storage_iam_store_v1_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_iam_store_v1_role_proto_goTypes,
		DependencyIndexes: file_controller_storage_iam_store_v1_role_proto_depIdxs,
		MessageInfos:      file_controller_storage_iam_store_v1_role_proto_msgTypes,
	}.Build()
	File_controller_storage_iam_store_v1_role_proto = out.File
	file_controller_storage_iam_store_v1_role_proto_rawDesc = nil
	file_controller_storage_iam_store_v1_role_proto_goTypes = nil
	file_controller_storage_iam_store_v1_role_proto_depIdxs = nil
}
