// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/dns/v1/dns_zone.proto

package dns

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type DnsZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId          string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name              string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description       string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels            map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Zone              string               `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone,omitempty"`
	PrivateVisibility *PrivateVisibility   `protobuf:"bytes,8,opt,name=private_visibility,json=privateVisibility,proto3" json:"private_visibility,omitempty"`
	PublicVisibility  *PublicVisibility    `protobuf:"bytes,9,opt,name=public_visibility,json=publicVisibility,proto3" json:"public_visibility,omitempty"`
}

func (x *DnsZone) Reset() {
	*x = DnsZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZone) ProtoMessage() {}

func (x *DnsZone) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZone.ProtoReflect.Descriptor instead.
func (*DnsZone) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{0}
}

func (x *DnsZone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DnsZone) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *DnsZone) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DnsZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsZone) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DnsZone) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DnsZone) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *DnsZone) GetPrivateVisibility() *PrivateVisibility {
	if x != nil {
		return x.PrivateVisibility
	}
	return nil
}

func (x *DnsZone) GetPublicVisibility() *PublicVisibility {
	if x != nil {
		return x.PublicVisibility
	}
	return nil
}

// name + type is a unique record set identifier
type RecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Ttl  int64    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Data []string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{1}
}

func (x *RecordSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecordSet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RecordSet) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *RecordSet) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

type PrivateVisibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkIds []string `protobuf:"bytes,1,rep,name=network_ids,json=networkIds,proto3" json:"network_ids,omitempty"`
}

func (x *PrivateVisibility) Reset() {
	*x = PrivateVisibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateVisibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateVisibility) ProtoMessage() {}

func (x *PrivateVisibility) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateVisibility.ProtoReflect.Descriptor instead.
func (*PrivateVisibility) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateVisibility) GetNetworkIds() []string {
	if x != nil {
		return x.NetworkIds
	}
	return nil
}

type PublicVisibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublicVisibility) Reset() {
	*x = PublicVisibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVisibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVisibility) ProtoMessage() {}

func (x *PublicVisibility) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVisibility.ProtoReflect.Descriptor instead.
func (*PublicVisibility) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{3}
}

var File_yandex_cloud_dns_v1_dns_zone_proto protoreflect.FileDescriptor

var file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc = []byte{
	0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x07, 0x44, 0x6e,
	0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e,
	0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x11, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x52, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x98, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31,
	0x05, 0x31, 0x2d, 0x32, 0x35, 0x34, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04,
	0x31, 0x2d, 0x32, 0x30, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x10, 0xfa, 0xc7, 0x31, 0x0c, 0x30, 0x2d, 0x32,
	0x31, 0x34, 0x37, 0x34, 0x38, 0x33, 0x36, 0x34, 0x37, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x16, 0x82, 0xc8,
	0x31, 0x05, 0x31, 0x2d, 0x31, 0x30, 0x30, 0x8a, 0xc8, 0x31, 0x05, 0x31, 0x2d, 0x32, 0x35, 0x35,
	0x90, 0xc8, 0x31, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x48, 0x0a, 0x11, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x33, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x12, 0x82, 0xc8, 0x31, 0x04, 0x30, 0x2d, 0x31, 0x30, 0x8a, 0xc8,
	0x31, 0x02, 0x32, 0x30, 0x90, 0xc8, 0x31, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescOnce sync.Once
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData = file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc
)

func file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP() []byte {
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData)
	})
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData
}

var file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_dns_v1_dns_zone_proto_goTypes = []interface{}{
	(*DnsZone)(nil),             // 0: yandex.cloud.dns.v1.DnsZone
	(*RecordSet)(nil),           // 1: yandex.cloud.dns.v1.RecordSet
	(*PrivateVisibility)(nil),   // 2: yandex.cloud.dns.v1.PrivateVisibility
	(*PublicVisibility)(nil),    // 3: yandex.cloud.dns.v1.PublicVisibility
	nil,                         // 4: yandex.cloud.dns.v1.DnsZone.LabelsEntry
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.dns.v1.DnsZone.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.dns.v1.DnsZone.labels:type_name -> yandex.cloud.dns.v1.DnsZone.LabelsEntry
	2, // 2: yandex.cloud.dns.v1.DnsZone.private_visibility:type_name -> yandex.cloud.dns.v1.PrivateVisibility
	3, // 3: yandex.cloud.dns.v1.DnsZone.public_visibility:type_name -> yandex.cloud.dns.v1.PublicVisibility
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_dns_v1_dns_zone_proto_init() }
func file_yandex_cloud_dns_v1_dns_zone_proto_init() {
	if File_yandex_cloud_dns_v1_dns_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZone); i {
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
		file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSet); i {
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
		file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateVisibility); i {
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
		file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVisibility); i {
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
			RawDescriptor: file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_dns_v1_dns_zone_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes,
	}.Build()
	File_yandex_cloud_dns_v1_dns_zone_proto = out.File
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc = nil
	file_yandex_cloud_dns_v1_dns_zone_proto_goTypes = nil
	file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs = nil
}
