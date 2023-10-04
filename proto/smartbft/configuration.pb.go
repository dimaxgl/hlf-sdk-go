// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: smartbft/configuration.proto

package smartbft

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

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "smartbft".
type ConfigMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consenters []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options    *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ConfigMetadata) Reset() {
	*x = ConfigMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartbft_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMetadata) ProtoMessage() {}

func (x *ConfigMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_smartbft_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMetadata.ProtoReflect.Descriptor instead.
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return file_smartbft_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigMetadata) GetConsenters() []*Consenter {
	if x != nil {
		return x.Consenters
	}
	return nil
}

func (x *ConfigMetadata) GetOptions() *Options {
	if x != nil {
		return x.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsenterId   uint64 `protobuf:"varint,1,opt,name=consenter_id,json=consenterId,proto3" json:"consenter_id,omitempty"`
	Host          string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	MspId         string `protobuf:"bytes,4,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	Identity      []byte `protobuf:"bytes,5,opt,name=identity,proto3" json:"identity,omitempty"`
	ClientTlsCert []byte `protobuf:"bytes,6,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert []byte `protobuf:"bytes,7,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
}

func (x *Consenter) Reset() {
	*x = Consenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartbft_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consenter) ProtoMessage() {}

func (x *Consenter) ProtoReflect() protoreflect.Message {
	mi := &file_smartbft_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consenter.ProtoReflect.Descriptor instead.
func (*Consenter) Descriptor() ([]byte, []int) {
	return file_smartbft_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *Consenter) GetConsenterId() uint64 {
	if x != nil {
		return x.ConsenterId
	}
	return 0
}

func (x *Consenter) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Consenter) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Consenter) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *Consenter) GetIdentity() []byte {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Consenter) GetClientTlsCert() []byte {
	if x != nil {
		return x.ClientTlsCert
	}
	return nil
}

func (x *Consenter) GetServerTlsCert() []byte {
	if x != nil {
		return x.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the smartbft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestBatchMaxCount      uint64 `protobuf:"varint,2,opt,name=request_batch_max_count,json=requestBatchMaxCount,proto3" json:"request_batch_max_count,omitempty"`
	RequestBatchMaxBytes      uint64 `protobuf:"varint,3,opt,name=request_batch_max_bytes,json=requestBatchMaxBytes,proto3" json:"request_batch_max_bytes,omitempty"`
	RequestBatchMaxInterval   string `protobuf:"bytes,4,opt,name=request_batch_max_interval,json=requestBatchMaxInterval,proto3" json:"request_batch_max_interval,omitempty"`
	IncomingMessageBufferSize uint64 `protobuf:"varint,5,opt,name=incoming_message_buffer_size,json=incomingMessageBufferSize,proto3" json:"incoming_message_buffer_size,omitempty"`
	RequestPoolSize           uint64 `protobuf:"varint,6,opt,name=request_pool_size,json=requestPoolSize,proto3" json:"request_pool_size,omitempty"`
	RequestForwardTimeout     string `protobuf:"bytes,7,opt,name=request_forward_timeout,json=requestForwardTimeout,proto3" json:"request_forward_timeout,omitempty"`
	RequestComplainTimeout    string `protobuf:"bytes,8,opt,name=request_complain_timeout,json=requestComplainTimeout,proto3" json:"request_complain_timeout,omitempty"`
	RequestAutoRemoveTimeout  string `protobuf:"bytes,9,opt,name=request_auto_remove_timeout,json=requestAutoRemoveTimeout,proto3" json:"request_auto_remove_timeout,omitempty"`
	ViewChangeResendInterval  string `protobuf:"bytes,10,opt,name=view_change_resend_interval,json=viewChangeResendInterval,proto3" json:"view_change_resend_interval,omitempty"`
	ViewChangeTimeout         string `protobuf:"bytes,11,opt,name=view_change_timeout,json=viewChangeTimeout,proto3" json:"view_change_timeout,omitempty"`
	LeaderHeartbeatTimeout    string `protobuf:"bytes,12,opt,name=leader_heartbeat_timeout,json=leaderHeartbeatTimeout,proto3" json:"leader_heartbeat_timeout,omitempty"`
	LeaderHeartbeatCount      uint64 `protobuf:"varint,13,opt,name=leader_heartbeat_count,json=leaderHeartbeatCount,proto3" json:"leader_heartbeat_count,omitempty"`
	CollectTimeout            string `protobuf:"bytes,14,opt,name=collect_timeout,json=collectTimeout,proto3" json:"collect_timeout,omitempty"`
	SyncOnStart               bool   `protobuf:"varint,15,opt,name=sync_on_start,json=syncOnStart,proto3" json:"sync_on_start,omitempty"`
	SpeedUpViewChange         bool   `protobuf:"varint,16,opt,name=speed_up_view_change,json=speedUpViewChange,proto3" json:"speed_up_view_change,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartbft_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_smartbft_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_smartbft_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetRequestBatchMaxCount() uint64 {
	if x != nil {
		return x.RequestBatchMaxCount
	}
	return 0
}

func (x *Options) GetRequestBatchMaxBytes() uint64 {
	if x != nil {
		return x.RequestBatchMaxBytes
	}
	return 0
}

func (x *Options) GetRequestBatchMaxInterval() string {
	if x != nil {
		return x.RequestBatchMaxInterval
	}
	return ""
}

func (x *Options) GetIncomingMessageBufferSize() uint64 {
	if x != nil {
		return x.IncomingMessageBufferSize
	}
	return 0
}

func (x *Options) GetRequestPoolSize() uint64 {
	if x != nil {
		return x.RequestPoolSize
	}
	return 0
}

func (x *Options) GetRequestForwardTimeout() string {
	if x != nil {
		return x.RequestForwardTimeout
	}
	return ""
}

func (x *Options) GetRequestComplainTimeout() string {
	if x != nil {
		return x.RequestComplainTimeout
	}
	return ""
}

func (x *Options) GetRequestAutoRemoveTimeout() string {
	if x != nil {
		return x.RequestAutoRemoveTimeout
	}
	return ""
}

func (x *Options) GetViewChangeResendInterval() string {
	if x != nil {
		return x.ViewChangeResendInterval
	}
	return ""
}

func (x *Options) GetViewChangeTimeout() string {
	if x != nil {
		return x.ViewChangeTimeout
	}
	return ""
}

func (x *Options) GetLeaderHeartbeatTimeout() string {
	if x != nil {
		return x.LeaderHeartbeatTimeout
	}
	return ""
}

func (x *Options) GetLeaderHeartbeatCount() uint64 {
	if x != nil {
		return x.LeaderHeartbeatCount
	}
	return 0
}

func (x *Options) GetCollectTimeout() string {
	if x != nil {
		return x.CollectTimeout
	}
	return ""
}

func (x *Options) GetSyncOnStart() bool {
	if x != nil {
		return x.SyncOnStart
	}
	return false
}

func (x *Options) GetSpeedUpViewChange() bool {
	if x != nil {
		return x.SpeedUpViewChange
	}
	return false
}

var File_smartbft_configuration_proto protoreflect.FileDescriptor

var file_smartbft_configuration_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x22, 0x72, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd9, 0x01, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x22, 0xbd, 0x06, 0x0a, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x19, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x17,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d,
	0x0a, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a,
	0x1b, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x18, 0x76, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x18,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x79,
	0x6e, 0x63, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x5f, 0x75, 0x70, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x70, 0x65, 0x65, 0x64, 0x55, 0x70,
	0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62,
	0x2f, 0x68, 0x6c, 0x66, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_smartbft_configuration_proto_rawDescOnce sync.Once
	file_smartbft_configuration_proto_rawDescData = file_smartbft_configuration_proto_rawDesc
)

func file_smartbft_configuration_proto_rawDescGZIP() []byte {
	file_smartbft_configuration_proto_rawDescOnce.Do(func() {
		file_smartbft_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_smartbft_configuration_proto_rawDescData)
	})
	return file_smartbft_configuration_proto_rawDescData
}

var file_smartbft_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_smartbft_configuration_proto_goTypes = []interface{}{
	(*ConfigMetadata)(nil), // 0: smartbft.ConfigMetadata
	(*Consenter)(nil),      // 1: smartbft.Consenter
	(*Options)(nil),        // 2: smartbft.Options
}
var file_smartbft_configuration_proto_depIdxs = []int32{
	1, // 0: smartbft.ConfigMetadata.consenters:type_name -> smartbft.Consenter
	2, // 1: smartbft.ConfigMetadata.options:type_name -> smartbft.Options
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_smartbft_configuration_proto_init() }
func file_smartbft_configuration_proto_init() {
	if File_smartbft_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smartbft_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMetadata); i {
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
		file_smartbft_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consenter); i {
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
		file_smartbft_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_smartbft_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smartbft_configuration_proto_goTypes,
		DependencyIndexes: file_smartbft_configuration_proto_depIdxs,
		MessageInfos:      file_smartbft_configuration_proto_msgTypes,
	}.Build()
	File_smartbft_configuration_proto = out.File
	file_smartbft_configuration_proto_rawDesc = nil
	file_smartbft_configuration_proto_goTypes = nil
	file_smartbft_configuration_proto_depIdxs = nil
}