// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pkg/experience/pb/experience.proto

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

type CreateExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Images      []string `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	AltImage    string   `protobuf:"bytes,4,opt,name=altImage,proto3" json:"altImage,omitempty"`
	Price       float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int32    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Location    string   `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Category    string   `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Tags        []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"` // Add any other relevant fields as needed
}

func (x *CreateExperienceRequest) Reset() {
	*x = CreateExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperienceRequest) ProtoMessage() {}

func (x *CreateExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperienceRequest.ProtoReflect.Descriptor instead.
func (*CreateExperienceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExperienceRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateExperienceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateExperienceRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *CreateExperienceRequest) GetAltImage() string {
	if x != nil {
		return x.AltImage
	}
	return ""
}

func (x *CreateExperienceRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateExperienceRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateExperienceRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateExperienceRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateExperienceRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Experience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Images      []string `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`
	AltImage    string   `protobuf:"bytes,5,opt,name=altImage,proto3" json:"altImage,omitempty"`
	Price       float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int32    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Location    string   `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	Category    string   `protobuf:"bytes,9,opt,name=category,proto3" json:"category,omitempty"`
	Tags        []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"` // Add any other relevant fields as needed
}

func (x *Experience) Reset() {
	*x = Experience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experience) ProtoMessage() {}

func (x *Experience) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experience.ProtoReflect.Descriptor instead.
func (*Experience) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{1}
}

func (x *Experience) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Experience) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Experience) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Experience) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Experience) GetAltImage() string {
	if x != nil {
		return x.AltImage
	}
	return ""
}

func (x *Experience) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Experience) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Experience) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Experience) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Experience) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreateExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Add any additional fields as needed
}

func (x *CreateExperienceResponse) Reset() {
	*x = CreateExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperienceResponse) ProtoMessage() {}

func (x *CreateExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperienceResponse.ProtoReflect.Descriptor instead.
func (*CreateExperienceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{2}
}

func (x *CreateExperienceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateExperienceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Categories  []string `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
	OrganizerId string   `protobuf:"bytes,5,opt,name=organizerId,proto3" json:"organizerId,omitempty"`
	Images      []string `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	AltImage    string   `protobuf:"bytes,7,opt,name=altImage,proto3" json:"altImage,omitempty"`
	Price       float64  `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int32    `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	Location    string   `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	Category    string   `protobuf:"bytes,11,opt,name=category,proto3" json:"category,omitempty"`
	Tags        []string `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"` // Add any additional fields as needed
}

func (x *UpdateExperienceRequest) Reset() {
	*x = UpdateExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExperienceRequest) ProtoMessage() {}

func (x *UpdateExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExperienceRequest.ProtoReflect.Descriptor instead.
func (*UpdateExperienceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateExperienceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateExperienceRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateExperienceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateExperienceRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *UpdateExperienceRequest) GetOrganizerId() string {
	if x != nil {
		return x.OrganizerId
	}
	return ""
}

func (x *UpdateExperienceRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *UpdateExperienceRequest) GetAltImage() string {
	if x != nil {
		return x.AltImage
	}
	return ""
}

func (x *UpdateExperienceRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateExperienceRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UpdateExperienceRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateExperienceRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateExperienceRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Add any additional fields as needed
}

func (x *UpdateExperienceResponse) Reset() {
	*x = UpdateExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExperienceResponse) ProtoMessage() {}

func (x *UpdateExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExperienceResponse.ProtoReflect.Descriptor instead.
func (*UpdateExperienceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateExperienceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateExperienceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Add any additional fields as needed
}

func (x *DeleteExperienceRequest) Reset() {
	*x = DeleteExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExperienceRequest) ProtoMessage() {}

func (x *DeleteExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExperienceRequest.ProtoReflect.Descriptor instead.
func (*DeleteExperienceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteExperienceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Add any additional fields as needed
}

func (x *DeleteExperienceResponse) Reset() {
	*x = DeleteExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_experience_pb_experience_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExperienceResponse) ProtoMessage() {}

func (x *DeleteExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_experience_pb_experience_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExperienceResponse.ProtoReflect.Descriptor instead.
func (*DeleteExperienceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_experience_pb_experience_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteExperienceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteExperienceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_experience_pb_experience_proto protoreflect.FileDescriptor

var file_pkg_experience_pb_experience_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x83, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x4e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xd5, 0x02, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x6c, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x4e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xb6, 0x02, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_experience_pb_experience_proto_rawDescOnce sync.Once
	file_pkg_experience_pb_experience_proto_rawDescData = file_pkg_experience_pb_experience_proto_rawDesc
)

func file_pkg_experience_pb_experience_proto_rawDescGZIP() []byte {
	file_pkg_experience_pb_experience_proto_rawDescOnce.Do(func() {
		file_pkg_experience_pb_experience_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_experience_pb_experience_proto_rawDescData)
	})
	return file_pkg_experience_pb_experience_proto_rawDescData
}

var file_pkg_experience_pb_experience_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_experience_pb_experience_proto_goTypes = []interface{}{
	(*CreateExperienceRequest)(nil),  // 0: experience.CreateExperienceRequest
	(*Experience)(nil),               // 1: experience.Experience
	(*CreateExperienceResponse)(nil), // 2: experience.CreateExperienceResponse
	(*UpdateExperienceRequest)(nil),  // 3: experience.UpdateExperienceRequest
	(*UpdateExperienceResponse)(nil), // 4: experience.UpdateExperienceResponse
	(*DeleteExperienceRequest)(nil),  // 5: experience.DeleteExperienceRequest
	(*DeleteExperienceResponse)(nil), // 6: experience.DeleteExperienceResponse
}
var file_pkg_experience_pb_experience_proto_depIdxs = []int32{
	0, // 0: experience.ExperienceService.CreateExperience:input_type -> experience.CreateExperienceRequest
	3, // 1: experience.ExperienceService.UpdateExperience:input_type -> experience.UpdateExperienceRequest
	5, // 2: experience.ExperienceService.DeleteExperience:input_type -> experience.DeleteExperienceRequest
	2, // 3: experience.ExperienceService.CreateExperience:output_type -> experience.CreateExperienceResponse
	4, // 4: experience.ExperienceService.UpdateExperience:output_type -> experience.UpdateExperienceResponse
	6, // 5: experience.ExperienceService.DeleteExperience:output_type -> experience.DeleteExperienceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_experience_pb_experience_proto_init() }
func file_pkg_experience_pb_experience_proto_init() {
	if File_pkg_experience_pb_experience_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_experience_pb_experience_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperienceRequest); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experience); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperienceResponse); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExperienceRequest); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExperienceResponse); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExperienceRequest); i {
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
		file_pkg_experience_pb_experience_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExperienceResponse); i {
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
			RawDescriptor: file_pkg_experience_pb_experience_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_experience_pb_experience_proto_goTypes,
		DependencyIndexes: file_pkg_experience_pb_experience_proto_depIdxs,
		MessageInfos:      file_pkg_experience_pb_experience_proto_msgTypes,
	}.Build()
	File_pkg_experience_pb_experience_proto = out.File
	file_pkg_experience_pb_experience_proto_rawDesc = nil
	file_pkg_experience_pb_experience_proto_goTypes = nil
	file_pkg_experience_pb_experience_proto_depIdxs = nil
}