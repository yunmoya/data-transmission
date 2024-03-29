// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: gocosi.proto

package gocosi

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

type RegisterNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes string `protobuf:"bytes,1,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *RegisterNodeReq) Reset() {
	*x = RegisterNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeReq) ProtoMessage() {}

func (x *RegisterNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeReq.ProtoReflect.Descriptor instead.
func (*RegisterNodeReq) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterNodeReq) GetNodes() string {
	if x != nil {
		return x.Nodes
	}
	return ""
}

type NewMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NewMsgReq) Reset() {
	*x = NewMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMsgReq) ProtoMessage() {}

func (x *NewMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMsgReq.ProtoReflect.Descriptor instead.
func (*NewMsgReq) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{1}
}

func (x *NewMsgReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Gossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subs   []string `protobuf:"bytes,1,rep,name=subs,proto3" json:"subs,omitempty"`
	Unsubs []string `protobuf:"bytes,2,rep,name=unsubs,proto3" json:"unsubs,omitempty"`
	// repeated string events_id = 3;
	EventList map[string]string `protobuf:"bytes,3,rep,name=event_list,json=eventList,proto3" json:"event_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Gossip) Reset() {
	*x = Gossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gossip) ProtoMessage() {}

func (x *Gossip) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gossip.ProtoReflect.Descriptor instead.
func (*Gossip) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{2}
}

func (x *Gossip) GetSubs() []string {
	if x != nil {
		return x.Subs
	}
	return nil
}

func (x *Gossip) GetUnsubs() []string {
	if x != nil {
		return x.Unsubs
	}
	return nil
}

func (x *Gossip) GetEventList() map[string]string {
	if x != nil {
		return x.EventList
	}
	return nil
}

type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subs       []string  `protobuf:"bytes,1,rep,name=subs,proto3" json:"subs,omitempty"`
	Unsubs     []string  `protobuf:"bytes,2,rep,name=unsubs,proto3" json:"unsubs,omitempty"`
	Pubkey     []string  `protobuf:"bytes,3,rep,name=pubkey,proto3" json:"pubkey,omitempty"`
	Latency    []float64 `protobuf:"fixed64,4,rep,packed,name=latency,proto3" json:"latency,omitempty"`
	RoundTime  string    `protobuf:"bytes,5,opt,name=round_time,json=roundTime,proto3" json:"round_time,omitempty"`
	Neighbours int32     `protobuf:"varint,6,opt,name=neighbours,proto3" json:"neighbours,omitempty"`
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResp) GetSubs() []string {
	if x != nil {
		return x.Subs
	}
	return nil
}

func (x *InfoResp) GetUnsubs() []string {
	if x != nil {
		return x.Unsubs
	}
	return nil
}

func (x *InfoResp) GetPubkey() []string {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *InfoResp) GetLatency() []float64 {
	if x != nil {
		return x.Latency
	}
	return nil
}

func (x *InfoResp) GetRoundTime() string {
	if x != nil {
		return x.RoundTime
	}
	return ""
}

func (x *InfoResp) GetNeighbours() int32 {
	if x != nil {
		return x.Neighbours
	}
	return 0
}

type GetPubkeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publickeys map[string]string `protobuf:"bytes,1,rep,name=publickeys,proto3" json:"publickeys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPubkeyResp) Reset() {
	*x = GetPubkeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPubkeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPubkeyResp) ProtoMessage() {}

func (x *GetPubkeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPubkeyResp.ProtoReflect.Descriptor instead.
func (*GetPubkeyResp) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{4}
}

func (x *GetPubkeyResp) GetPublickeys() map[string]string {
	if x != nil {
		return x.Publickeys
	}
	return nil
}

type NewMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string  `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Signers   []int32 `protobuf:"varint,2,rep,packed,name=signers,proto3" json:"signers,omitempty"`
}

func (x *NewMsgResp) Reset() {
	*x = NewMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMsgResp) ProtoMessage() {}

func (x *NewMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMsgResp.ProtoReflect.Descriptor instead.
func (*NewMsgResp) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{5}
}

func (x *NewMsgResp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *NewMsgResp) GetSigners() []int32 {
	if x != nil {
		return x.Signers
	}
	return nil
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{6}
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocosi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gocosi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gocosi_proto_rawDescGZIP(), []int{7}
}

var File_gocosi_proto protoreflect.FileDescriptor

var file_gocosi_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22,
	0x1d, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xb0,
	0x01, 0x0a, 0x06, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x62,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x75, 0x62, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x63, 0x6f,
	0x73, 0x69, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x75,
	0x62, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b,
	0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x6b, 0x65, 0x79, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65,
	0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x90, 0x02, 0x0a, 0x09, 0x47,
	0x6f, 0x63, 0x6f, 0x73, 0x69, 0x52, 0x50, 0x43, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73,
	0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x4d, 0x73,
	0x67, 0x12, 0x11, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x4e, 0x65,
	0x77, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69,
	0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x0d, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x63, 0x6f, 0x73, 0x69, 0x2f, 0x67, 0x6f, 0x63, 0x6f,
	0x73, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gocosi_proto_rawDescOnce sync.Once
	file_gocosi_proto_rawDescData = file_gocosi_proto_rawDesc
)

func file_gocosi_proto_rawDescGZIP() []byte {
	file_gocosi_proto_rawDescOnce.Do(func() {
		file_gocosi_proto_rawDescData = protoimpl.X.CompressGZIP(file_gocosi_proto_rawDescData)
	})
	return file_gocosi_proto_rawDescData
}

var file_gocosi_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gocosi_proto_goTypes = []interface{}{
	(*RegisterNodeReq)(nil), // 0: gocosi.RegisterNodeReq
	(*NewMsgReq)(nil),       // 1: gocosi.NewMsgReq
	(*Gossip)(nil),          // 2: gocosi.Gossip
	(*InfoResp)(nil),        // 3: gocosi.InfoResp
	(*GetPubkeyResp)(nil),   // 4: gocosi.GetPubkeyResp
	(*NewMsgResp)(nil),      // 5: gocosi.NewMsgResp
	(*CommonResp)(nil),      // 6: gocosi.CommonResp
	(*Empty)(nil),           // 7: gocosi.Empty
	nil,                     // 8: gocosi.Gossip.EventListEntry
	nil,                     // 9: gocosi.GetPubkeyResp.PublickeysEntry
}
var file_gocosi_proto_depIdxs = []int32{
	8, // 0: gocosi.Gossip.event_list:type_name -> gocosi.Gossip.EventListEntry
	9, // 1: gocosi.GetPubkeyResp.publickeys:type_name -> gocosi.GetPubkeyResp.PublickeysEntry
	0, // 2: gocosi.GocosiRPC.RegisterNode:input_type -> gocosi.RegisterNodeReq
	1, // 3: gocosi.GocosiRPC.NewMsg:input_type -> gocosi.NewMsgReq
	2, // 4: gocosi.GocosiRPC.GossipReq:input_type -> gocosi.Gossip
	7, // 5: gocosi.GocosiRPC.Info:input_type -> gocosi.Empty
	7, // 6: gocosi.GocosiRPC.GetPubkey:input_type -> gocosi.Empty
	6, // 7: gocosi.GocosiRPC.RegisterNode:output_type -> gocosi.CommonResp
	5, // 8: gocosi.GocosiRPC.NewMsg:output_type -> gocosi.NewMsgResp
	6, // 9: gocosi.GocosiRPC.GossipReq:output_type -> gocosi.CommonResp
	3, // 10: gocosi.GocosiRPC.Info:output_type -> gocosi.InfoResp
	4, // 11: gocosi.GocosiRPC.GetPubkey:output_type -> gocosi.GetPubkeyResp
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gocosi_proto_init() }
func file_gocosi_proto_init() {
	if File_gocosi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gocosi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeReq); i {
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
		file_gocosi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMsgReq); i {
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
		file_gocosi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gossip); i {
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
		file_gocosi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResp); i {
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
		file_gocosi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPubkeyResp); i {
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
		file_gocosi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMsgResp); i {
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
		file_gocosi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_gocosi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_gocosi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gocosi_proto_goTypes,
		DependencyIndexes: file_gocosi_proto_depIdxs,
		MessageInfos:      file_gocosi_proto_msgTypes,
	}.Build()
	File_gocosi_proto = out.File
	file_gocosi_proto_rawDesc = nil
	file_gocosi_proto_goTypes = nil
	file_gocosi_proto_depIdxs = nil
}
