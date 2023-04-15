// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: idl/api/website.proto

package api

import (
	context "context"
	base "github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
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

type CreateWebsiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AvatarIcon     string `protobuf:"bytes,4,opt,name=avatar_icon,json=avatarIcon,proto3" json:"avatar_icon,omitempty"`
	Background     string `protobuf:"bytes,5,opt,name=background,proto3" json:"background,omitempty"`
	Language       string `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	IsPublic       bool   `protobuf:"varint,7,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	AllowAnonymous bool   `protobuf:"varint,8,opt,name=allow_anonymous,json=allowAnonymous,proto3" json:"allow_anonymous,omitempty"`
}

func (x *CreateWebsiteRequest) Reset() {
	*x = CreateWebsiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebsiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebsiteRequest) ProtoMessage() {}

func (x *CreateWebsiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebsiteRequest.ProtoReflect.Descriptor instead.
func (*CreateWebsiteRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWebsiteRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateWebsiteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWebsiteRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWebsiteRequest) GetAvatarIcon() string {
	if x != nil {
		return x.AvatarIcon
	}
	return ""
}

func (x *CreateWebsiteRequest) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *CreateWebsiteRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CreateWebsiteRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *CreateWebsiteRequest) GetAllowAnonymous() bool {
	if x != nil {
		return x.AllowAnonymous
	}
	return false
}

type CreateWebsiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateWebsiteResponse) Reset() {
	*x = CreateWebsiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebsiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebsiteResponse) ProtoMessage() {}

func (x *CreateWebsiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebsiteResponse.ProtoReflect.Descriptor instead.
func (*CreateWebsiteResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWebsiteResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateWebsiteResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetWebsiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWebsiteRequest) Reset() {
	*x = GetWebsiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebsiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebsiteRequest) ProtoMessage() {}

func (x *GetWebsiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebsiteRequest.ProtoReflect.Descriptor instead.
func (*GetWebsiteRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{2}
}

func (x *GetWebsiteRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetWebsiteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWebsiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key            string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description    string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	AvatarIcon     string `protobuf:"bytes,5,opt,name=avatar_icon,json=avatarIcon,proto3" json:"avatar_icon,omitempty"`
	Background     string `protobuf:"bytes,6,opt,name=background,proto3" json:"background,omitempty"`
	Language       string `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	IsPublic       bool   `protobuf:"varint,8,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	AllowAnonymous bool   `protobuf:"varint,9,opt,name=allow_anonymous,json=allowAnonymous,proto3" json:"allow_anonymous,omitempty"`
}

func (x *GetWebsiteResponse) Reset() {
	*x = GetWebsiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebsiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebsiteResponse) ProtoMessage() {}

func (x *GetWebsiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebsiteResponse.ProtoReflect.Descriptor instead.
func (*GetWebsiteResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{3}
}

func (x *GetWebsiteResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetWebsiteResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetWebsiteResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetWebsiteResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetWebsiteResponse) GetAvatarIcon() string {
	if x != nil {
		return x.AvatarIcon
	}
	return ""
}

func (x *GetWebsiteResponse) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *GetWebsiteResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetWebsiteResponse) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *GetWebsiteResponse) GetAllowAnonymous() bool {
	if x != nil {
		return x.AllowAnonymous
	}
	return false
}

type UpdateWebsiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id             uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Key            string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description    string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	AvatarIcon     string `protobuf:"bytes,6,opt,name=avatar_icon,json=avatarIcon,proto3" json:"avatar_icon,omitempty"`
	Background     string `protobuf:"bytes,7,opt,name=background,proto3" json:"background,omitempty"`
	Language       string `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`
	IsPublic       bool   `protobuf:"varint,9,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	AllowAnonymous bool   `protobuf:"varint,10,opt,name=allow_anonymous,json=allowAnonymous,proto3" json:"allow_anonymous,omitempty"`
}

func (x *UpdateWebsiteRequest) Reset() {
	*x = UpdateWebsiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWebsiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebsiteRequest) ProtoMessage() {}

func (x *UpdateWebsiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebsiteRequest.ProtoReflect.Descriptor instead.
func (*UpdateWebsiteRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateWebsiteRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateWebsiteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateWebsiteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetAvatarIcon() string {
	if x != nil {
		return x.AvatarIcon
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UpdateWebsiteRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *UpdateWebsiteRequest) GetAllowAnonymous() bool {
	if x != nil {
		return x.AllowAnonymous
	}
	return false
}

type ListWebsitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Pagination *base.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListWebsitesRequest) Reset() {
	*x = ListWebsitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebsitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebsitesRequest) ProtoMessage() {}

func (x *ListWebsitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebsitesRequest.ProtoReflect.Descriptor instead.
func (*ListWebsitesRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{5}
}

func (x *ListWebsitesRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListWebsitesRequest) GetPagination() *base.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListWebsitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64                 `protobuf:"zigzag64,1,opt,name=total,proto3" json:"total,omitempty"`
	Websites []*GetWebsiteResponse `protobuf:"bytes,2,rep,name=websites,proto3" json:"websites,omitempty"`
}

func (x *ListWebsitesResponse) Reset() {
	*x = ListWebsitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_website_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebsitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebsitesResponse) ProtoMessage() {}

func (x *ListWebsitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_website_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebsitesResponse.ProtoReflect.Descriptor instead.
func (*ListWebsitesResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_website_proto_rawDescGZIP(), []int{6}
}

func (x *ListWebsitesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListWebsitesResponse) GetWebsites() []*GetWebsiteResponse {
	if x != nil {
		return x.Websites
	}
	return nil
}

var File_idl_api_website_proto protoreflect.FileDescriptor

var file_idl_api_website_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x14, 0x69, 0x64,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x22, 0x39, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x8f, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x63, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x22, 0xaa, 0x02, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x22, 0x60, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x33, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x73, 0x32, 0x95, 0x02, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x72,
	0x2d, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x2d, 0x62, 0x6f, 0x78, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_idl_api_website_proto_rawDescOnce sync.Once
	file_idl_api_website_proto_rawDescData = file_idl_api_website_proto_rawDesc
)

func file_idl_api_website_proto_rawDescGZIP() []byte {
	file_idl_api_website_proto_rawDescOnce.Do(func() {
		file_idl_api_website_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_api_website_proto_rawDescData)
	})
	return file_idl_api_website_proto_rawDescData
}

var file_idl_api_website_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_idl_api_website_proto_goTypes = []interface{}{
	(*CreateWebsiteRequest)(nil),  // 0: api.CreateWebsiteRequest
	(*CreateWebsiteResponse)(nil), // 1: api.CreateWebsiteResponse
	(*GetWebsiteRequest)(nil),     // 2: api.GetWebsiteRequest
	(*GetWebsiteResponse)(nil),    // 3: api.GetWebsiteResponse
	(*UpdateWebsiteRequest)(nil),  // 4: api.UpdateWebsiteRequest
	(*ListWebsitesRequest)(nil),   // 5: api.ListWebsitesRequest
	(*ListWebsitesResponse)(nil),  // 6: api.ListWebsitesResponse
	(*base.Pagination)(nil),       // 7: base.Pagination
	(*base.Empty)(nil),            // 8: base.Empty
}
var file_idl_api_website_proto_depIdxs = []int32{
	7, // 0: api.ListWebsitesRequest.pagination:type_name -> base.Pagination
	3, // 1: api.ListWebsitesResponse.websites:type_name -> api.GetWebsiteResponse
	0, // 2: api.WebsiteService.CreateWebsite:input_type -> api.CreateWebsiteRequest
	2, // 3: api.WebsiteService.GetWebsite:input_type -> api.GetWebsiteRequest
	4, // 4: api.WebsiteService.UpdateWebsite:input_type -> api.UpdateWebsiteRequest
	5, // 5: api.WebsiteService.ListWebsites:input_type -> api.ListWebsitesRequest
	1, // 6: api.WebsiteService.CreateWebsite:output_type -> api.CreateWebsiteResponse
	3, // 7: api.WebsiteService.GetWebsite:output_type -> api.GetWebsiteResponse
	8, // 8: api.WebsiteService.UpdateWebsite:output_type -> base.Empty
	6, // 9: api.WebsiteService.ListWebsites:output_type -> api.ListWebsitesResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_idl_api_website_proto_init() }
func file_idl_api_website_proto_init() {
	if File_idl_api_website_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_api_website_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebsiteRequest); i {
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
		file_idl_api_website_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebsiteResponse); i {
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
		file_idl_api_website_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebsiteRequest); i {
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
		file_idl_api_website_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebsiteResponse); i {
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
		file_idl_api_website_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWebsiteRequest); i {
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
		file_idl_api_website_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebsitesRequest); i {
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
		file_idl_api_website_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebsitesResponse); i {
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
			RawDescriptor: file_idl_api_website_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_api_website_proto_goTypes,
		DependencyIndexes: file_idl_api_website_proto_depIdxs,
		MessageInfos:      file_idl_api_website_proto_msgTypes,
	}.Build()
	File_idl_api_website_proto = out.File
	file_idl_api_website_proto_rawDesc = nil
	file_idl_api_website_proto_goTypes = nil
	file_idl_api_website_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.5.1. DO NOT EDIT.

type WebsiteService interface {
	CreateWebsite(ctx context.Context, req *CreateWebsiteRequest) (res *CreateWebsiteResponse, err error)
	GetWebsite(ctx context.Context, req *GetWebsiteRequest) (res *GetWebsiteResponse, err error)
	UpdateWebsite(ctx context.Context, req *UpdateWebsiteRequest) (res *base.Empty, err error)
	ListWebsites(ctx context.Context, req *ListWebsitesRequest) (res *ListWebsitesResponse, err error)
}
