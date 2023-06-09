// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: idl/api/dash/comment.proto

package dash

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WebsiteRefer uint64          `protobuf:"varint,2,opt,name=website_refer,json=websiteRefer,proto3" json:"website_refer,omitempty"`
	Name         string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email        string          `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Url          string          `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Content      string          `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt    *base.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *base.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_dash_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_dash_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_idl_api_dash_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetWebsiteRefer() uint64 {
	if x != nil {
		return x.WebsiteRefer
	}
	return 0
}

func (x *Comment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Comment) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Comment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() *base.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *base.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WebsiteRefer uint64           `protobuf:"varint,2,opt,name=website_refer,json=websiteRefer,proto3" json:"website_refer,omitempty"`
	Pagination   *base.Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListCommentsRequest) Reset() {
	*x = ListCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_dash_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsRequest) ProtoMessage() {}

func (x *ListCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_dash_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListCommentsRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_dash_comment_proto_rawDescGZIP(), []int{1}
}

func (x *ListCommentsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListCommentsRequest) GetWebsiteRefer() uint64 {
	if x != nil {
		return x.WebsiteRefer
	}
	return 0
}

func (x *ListCommentsRequest) GetPagination() *base.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64      `protobuf:"zigzag64,1,opt,name=total,proto3" json:"total,omitempty"`
	Comments []*Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *ListCommentsResponse) Reset() {
	*x = ListCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_dash_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsResponse) ProtoMessage() {}

func (x *ListCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_dash_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListCommentsResponse) Descriptor() ([]byte, []int) {
	return file_idl_api_dash_comment_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_dash_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_dash_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_dash_comment_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCommentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_api_dash_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_api_dash_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_idl_api_dash_comment_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCommentRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteCommentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_idl_api_dash_comment_proto protoreflect.FileDescriptor

var file_idl_api_dash_comment_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x1a, 0x14, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x64,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdb, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2d,
	0x62, 0x6f, 0x78, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_api_dash_comment_proto_rawDescOnce sync.Once
	file_idl_api_dash_comment_proto_rawDescData = file_idl_api_dash_comment_proto_rawDesc
)

func file_idl_api_dash_comment_proto_rawDescGZIP() []byte {
	file_idl_api_dash_comment_proto_rawDescOnce.Do(func() {
		file_idl_api_dash_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_api_dash_comment_proto_rawDescData)
	})
	return file_idl_api_dash_comment_proto_rawDescData
}

var file_idl_api_dash_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_api_dash_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),              // 0: api.dash.Comment
	(*ListCommentsRequest)(nil),  // 1: api.dash.ListCommentsRequest
	(*ListCommentsResponse)(nil), // 2: api.dash.ListCommentsResponse
	(*GetCommentRequest)(nil),    // 3: api.dash.GetCommentRequest
	(*DeleteCommentRequest)(nil), // 4: api.dash.DeleteCommentRequest
	(*base.Timestamp)(nil),       // 5: base.Timestamp
	(*base.Pagination)(nil),      // 6: base.Pagination
	(*base.Empty)(nil),           // 7: base.Empty
}
var file_idl_api_dash_comment_proto_depIdxs = []int32{
	5, // 0: api.dash.Comment.created_at:type_name -> base.Timestamp
	5, // 1: api.dash.Comment.updated_at:type_name -> base.Timestamp
	6, // 2: api.dash.ListCommentsRequest.pagination:type_name -> base.Pagination
	0, // 3: api.dash.ListCommentsResponse.comments:type_name -> api.dash.Comment
	3, // 4: api.dash.CommentService.GetComment:input_type -> api.dash.GetCommentRequest
	1, // 5: api.dash.CommentService.ListComments:input_type -> api.dash.ListCommentsRequest
	4, // 6: api.dash.CommentService.DeleteComment:input_type -> api.dash.DeleteCommentRequest
	0, // 7: api.dash.CommentService.GetComment:output_type -> api.dash.Comment
	2, // 8: api.dash.CommentService.ListComments:output_type -> api.dash.ListCommentsResponse
	7, // 9: api.dash.CommentService.DeleteComment:output_type -> base.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_idl_api_dash_comment_proto_init() }
func file_idl_api_dash_comment_proto_init() {
	if File_idl_api_dash_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_api_dash_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_idl_api_dash_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsRequest); i {
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
		file_idl_api_dash_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsResponse); i {
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
		file_idl_api_dash_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_idl_api_dash_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
			RawDescriptor: file_idl_api_dash_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_api_dash_comment_proto_goTypes,
		DependencyIndexes: file_idl_api_dash_comment_proto_depIdxs,
		MessageInfos:      file_idl_api_dash_comment_proto_msgTypes,
	}.Build()
	File_idl_api_dash_comment_proto = out.File
	file_idl_api_dash_comment_proto_rawDesc = nil
	file_idl_api_dash_comment_proto_goTypes = nil
	file_idl_api_dash_comment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.5.1. DO NOT EDIT.

type CommentService interface {
	GetComment(ctx context.Context, req *GetCommentRequest) (res *Comment, err error)
	ListComments(ctx context.Context, req *ListCommentsRequest) (res *ListCommentsResponse, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest) (res *base.Empty, err error)
}
