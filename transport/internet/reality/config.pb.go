// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: transport/internet/reality/config.proto

package reality

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Show                  bool     `protobuf:"varint,1,opt,name=show,proto3" json:"show,omitempty"`
	Dest                  string   `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
	Type                  string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Xver                  uint64   `protobuf:"varint,4,opt,name=xver,proto3" json:"xver,omitempty"`
	ServerNames           []string `protobuf:"bytes,5,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	PrivateKey            []byte   `protobuf:"bytes,6,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	MinClientVer          []byte   `protobuf:"bytes,7,opt,name=min_client_ver,json=minClientVer,proto3" json:"min_client_ver,omitempty"`
	MaxClientVer          []byte   `protobuf:"bytes,8,opt,name=max_client_ver,json=maxClientVer,proto3" json:"max_client_ver,omitempty"`
	MaxTimeDiff           uint64   `protobuf:"varint,9,opt,name=max_time_diff,json=maxTimeDiff,proto3" json:"max_time_diff,omitempty"`
	ShortIds              [][]byte `protobuf:"bytes,10,rep,name=short_ids,json=shortIds,proto3" json:"short_ids,omitempty"`
	ServerRandPacket      string   `protobuf:"bytes,11,opt,name=server_rand_packet,json=serverRandPacket,proto3" json:"server_rand_packet,omitempty"`
	ClientRandPacket      string   `protobuf:"bytes,12,opt,name=client_rand_packet,json=clientRandPacket,proto3" json:"client_rand_packet,omitempty"`
	ServerRandPacketCount string   `protobuf:"bytes,13,opt,name=server_rand_packet_count,json=serverRandPacketCount,proto3" json:"server_rand_packet_count,omitempty"`
	ClientRandPacketCount string   `protobuf:"bytes,14,opt,name=client_rand_packet_count,json=clientRandPacketCount,proto3" json:"client_rand_packet_count,omitempty"`
	SplitPacket           string   `protobuf:"bytes,15,opt,name=split_packet,json=splitPacket,proto3" json:"split_packet,omitempty"`
	PaddingSize           uint32   `protobuf:"varint,16,opt,name=padding_size,json=paddingSize,proto3" json:"padding_size,omitempty"`
	SubchunkSize          uint32   `protobuf:"varint,17,opt,name=subchunk_size,json=subchunkSize,proto3" json:"subchunk_size,omitempty"`
	Fingerprint           string   `protobuf:"bytes,21,opt,name=Fingerprint,proto3" json:"Fingerprint,omitempty"`
	ServerName            string   `protobuf:"bytes,22,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	PublicKey             []byte   `protobuf:"bytes,23,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	ShortId               []byte   `protobuf:"bytes,24,opt,name=short_id,json=shortId,proto3" json:"short_id,omitempty"`
	SpiderX               string   `protobuf:"bytes,25,opt,name=spider_x,json=spiderX,proto3" json:"spider_x,omitempty"`
	SpiderY               []int64  `protobuf:"varint,26,rep,packed,name=spider_y,json=spiderY,proto3" json:"spider_y,omitempty"`
	MasterKeyLog          string   `protobuf:"bytes,27,opt,name=master_key_log,json=masterKeyLog,proto3" json:"master_key_log,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetShow() bool {
	if x != nil {
		return x.Show
	}
	return false
}

func (x *Config) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *Config) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Config) GetXver() uint64 {
	if x != nil {
		return x.Xver
	}
	return 0
}

func (x *Config) GetServerNames() []string {
	if x != nil {
		return x.ServerNames
	}
	return nil
}

func (x *Config) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *Config) GetMinClientVer() []byte {
	if x != nil {
		return x.MinClientVer
	}
	return nil
}

func (x *Config) GetMaxClientVer() []byte {
	if x != nil {
		return x.MaxClientVer
	}
	return nil
}

func (x *Config) GetMaxTimeDiff() uint64 {
	if x != nil {
		return x.MaxTimeDiff
	}
	return 0
}

func (x *Config) GetShortIds() [][]byte {
	if x != nil {
		return x.ShortIds
	}
	return nil
}

func (x *Config) GetServerRandPacket() string {
	if x != nil {
		return x.ServerRandPacket
	}
	return ""
}

func (x *Config) GetClientRandPacket() string {
	if x != nil {
		return x.ClientRandPacket
	}
	return ""
}

func (x *Config) GetServerRandPacketCount() string {
	if x != nil {
		return x.ServerRandPacketCount
	}
	return ""
}

func (x *Config) GetClientRandPacketCount() string {
	if x != nil {
		return x.ClientRandPacketCount
	}
	return ""
}

func (x *Config) GetSplitPacket() string {
	if x != nil {
		return x.SplitPacket
	}
	return ""
}

func (x *Config) GetPaddingSize() uint32 {
	if x != nil {
		return x.PaddingSize
	}
	return 0
}

func (x *Config) GetSubchunkSize() uint32 {
	if x != nil {
		return x.SubchunkSize
	}
	return 0
}

func (x *Config) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *Config) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Config) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Config) GetShortId() []byte {
	if x != nil {
		return x.ShortId
	}
	return nil
}

func (x *Config) GetSpiderX() string {
	if x != nil {
		return x.SpiderX
	}
	return ""
}

func (x *Config) GetSpiderY() []int64 {
	if x != nil {
		return x.SpiderY
	}
	return nil
}

func (x *Config) GetMasterKeyLog() string {
	if x != nil {
		return x.MasterKeyLog
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22,
	0xbb, 0x06, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x78, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x24,
	0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61,
	0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e,
	0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x37, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x78, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x58, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x79, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x59, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x42, 0x7f, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x74, 0x79, 0xaa, 0x02, 0x1f, 0x58,
	0x72, 0x61, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_proto_goTypes = []any{
	(*Config)(nil), // 0: xray.transport.internet.reality.Config
}
var file_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
