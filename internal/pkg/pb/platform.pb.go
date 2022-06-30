// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: platform.proto

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

type PlatformPlayer struct {
	state         protoimpl.MessageState
	Platform      string `protobuf:"bytes,12,opt,name=platform,proto3" json:"platform,omitempty"`
	DisplayName   string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken   string `protobuf:"bytes,10,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Language      string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	IdAtPlatform  string `protobuf:"bytes,9,opt,name=id_at_platform,json=idAtPlatform,proto3" json:"id_at_platform,omitempty"`
	Icon          string `protobuf:"bytes,13,opt,name=icon,proto3" json:"icon,omitempty"`
	unknownFields protoimpl.UnknownFields
	Balance       float64 `protobuf:"fixed64,7,opt,name=balance,proto3" json:"balance,omitempty"`
	LockBalance   float64 `protobuf:"fixed64,8,opt,name=lock_balance,json=lockBalance,proto3" json:"lock_balance,omitempty"`
	Credit        int64   `protobuf:"varint,5,opt,name=credit,proto3" json:"credit,omitempty"`
	Sn            int64   `protobuf:"varint,2,opt,name=sn,proto3" json:"sn,omitempty"`
	LockCredit    int64   `protobuf:"varint,6,opt,name=lock_credit,json=lockCredit,proto3" json:"lock_credit,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *PlatformPlayer) Reset() {
	*x = PlatformPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformPlayer) ProtoMessage() {}

func (x *PlatformPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformPlayer.ProtoReflect.Descriptor instead.
func (*PlatformPlayer) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformPlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformPlayer) GetSn() int64 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *PlatformPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatformPlayer) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *PlatformPlayer) GetCredit() int64 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *PlatformPlayer) GetLockCredit() int64 {
	if x != nil {
		return x.LockCredit
	}
	return 0
}

func (x *PlatformPlayer) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *PlatformPlayer) GetLockBalance() float64 {
	if x != nil {
		return x.LockBalance
	}
	return 0
}

func (x *PlatformPlayer) GetIdAtPlatform() string {
	if x != nil {
		return x.IdAtPlatform
	}
	return ""
}

func (x *PlatformPlayer) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PlatformPlayer) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *PlatformPlayer) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *PlatformPlayer) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type PlatformProvider struct {
	state           protoimpl.MessageState
	Auth            *Auth  `protobuf:"bytes,8,opt,name=auth,proto3" json:"auth,omitempty"`
	ApiUrlBase      string `protobuf:"bytes,5,opt,name=api_url_base,json=apiUrlBase,proto3" json:"api_url_base,omitempty"`
	PublicIpAddress string `protobuf:"bytes,6,opt,name=public_ip_address,json=publicIpAddress,proto3" json:"public_ip_address,omitempty"`
	FactoryName     string `protobuf:"bytes,1,opt,name=factory_name,json=factoryName,proto3" json:"factory_name,omitempty"`
	AesKey          string `protobuf:"bytes,3,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	AesIv           string `protobuf:"bytes,4,opt,name=aes_iv,json=aesIv,proto3" json:"aes_iv,omitempty"`
	RunMode         string `protobuf:"bytes,7,opt,name=run_mode,json=runMode,proto3" json:"run_mode,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PlatformProvider) Reset() {
	*x = PlatformProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformProvider) ProtoMessage() {}

func (x *PlatformProvider) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformProvider.ProtoReflect.Descriptor instead.
func (*PlatformProvider) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{1}
}

func (x *PlatformProvider) GetFactoryName() string {
	if x != nil {
		return x.FactoryName
	}
	return ""
}

func (x *PlatformProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlatformProvider) GetAesKey() string {
	if x != nil {
		return x.AesKey
	}
	return ""
}

func (x *PlatformProvider) GetAesIv() string {
	if x != nil {
		return x.AesIv
	}
	return ""
}

func (x *PlatformProvider) GetApiUrlBase() string {
	if x != nil {
		return x.ApiUrlBase
	}
	return ""
}

func (x *PlatformProvider) GetPublicIpAddress() string {
	if x != nil {
		return x.PublicIpAddress
	}
	return ""
}

func (x *PlatformProvider) GetRunMode() string {
	if x != nil {
		return x.RunMode
	}
	return ""
}

func (x *PlatformProvider) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type PlatformProviderAPIPath struct {
	state                   protoimpl.MessageState
	PlayerBetPlace          string `protobuf:"bytes,4,opt,name=player_bet_place,json=playerBetPlace,proto3" json:"player_bet_place,omitempty"`
	PlayerTransactionUnlock string `protobuf:"bytes,8,opt,name=player_transaction_unlock,json=playerTransactionUnlock,proto3" json:"player_transaction_unlock,omitempty"`
	OauthAccessToken        string `protobuf:"bytes,1,opt,name=oauth_access_token,json=oauthAccessToken,proto3" json:"oauth_access_token,omitempty"`
	PlayerTokenValidate     string `protobuf:"bytes,2,opt,name=player_token_validate,json=playerTokenValidate,proto3" json:"player_token_validate,omitempty"`
	PlayerBalance           string `protobuf:"bytes,3,opt,name=player_balance,json=playerBalance,proto3" json:"player_balance,omitempty"`
	PlayerTransactionCancel string `protobuf:"bytes,9,opt,name=player_transaction_cancel,json=playerTransactionCancel,proto3" json:"player_transaction_cancel,omitempty"`
	PlayerBetCancel         string `protobuf:"bytes,5,opt,name=player_bet_cancel,json=playerBetCancel,proto3" json:"player_bet_cancel,omitempty"`
	PlayerBetSettle         string `protobuf:"bytes,6,opt,name=player_bet_settle,json=playerBetSettle,proto3" json:"player_bet_settle,omitempty"`
	PlayerTransactionLock   string `protobuf:"bytes,7,opt,name=player_transaction_lock,json=playerTransactionLock,proto3" json:"player_transaction_lock,omitempty"`
	PlayerTransactionStatus string `protobuf:"bytes,10,opt,name=player_transaction_status,json=playerTransactionStatus,proto3" json:"player_transaction_status,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *PlatformProviderAPIPath) Reset() {
	*x = PlatformProviderAPIPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformProviderAPIPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformProviderAPIPath) ProtoMessage() {}

func (x *PlatformProviderAPIPath) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformProviderAPIPath.ProtoReflect.Descriptor instead.
func (*PlatformProviderAPIPath) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{2}
}

func (x *PlatformProviderAPIPath) GetOauthAccessToken() string {
	if x != nil {
		return x.OauthAccessToken
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerTokenValidate() string {
	if x != nil {
		return x.PlayerTokenValidate
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerBalance() string {
	if x != nil {
		return x.PlayerBalance
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerBetPlace() string {
	if x != nil {
		return x.PlayerBetPlace
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerBetCancel() string {
	if x != nil {
		return x.PlayerBetCancel
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerBetSettle() string {
	if x != nil {
		return x.PlayerBetSettle
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerTransactionLock() string {
	if x != nil {
		return x.PlayerTransactionLock
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerTransactionUnlock() string {
	if x != nil {
		return x.PlayerTransactionUnlock
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerTransactionCancel() string {
	if x != nil {
		return x.PlayerTransactionCancel
	}
	return ""
}

func (x *PlatformProviderAPIPath) GetPlayerTransactionStatus() string {
	if x != nil {
		return x.PlayerTransactionStatus
	}
	return ""
}

var File_platform_proto protoreflect.FileDescriptor

var file_platform_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x02, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x64, 0x5f,
	0x61, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x64, 0x41, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x80, 0x02, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x65,
	0x73, 0x5f, 0x69, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x65, 0x73, 0x49,
	0x76, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x90, 0x04, 0x0a, 0x17, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x62, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x65, 0x74, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x12,
	0x36, 0x0a, 0x17, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12,
	0x3a, 0x0a, 0x19, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x6f,
	0x68, 0x64, 0x61, 0x64, 0x61, 0x2f, 0x67, 0x32, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_proto_rawDescOnce sync.Once
	file_platform_proto_rawDescData = file_platform_proto_rawDesc
)

func file_platform_proto_rawDescGZIP() []byte {
	file_platform_proto_rawDescOnce.Do(func() {
		file_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_proto_rawDescData)
	})
	return file_platform_proto_rawDescData
}

var file_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_platform_proto_goTypes = []interface{}{
	(*PlatformPlayer)(nil),          // 0: pb.PlatformPlayer
	(*PlatformProvider)(nil),        // 1: pb.PlatformProvider
	(*PlatformProviderAPIPath)(nil), // 2: pb.PlatformProviderAPIPath
	(*Auth)(nil),                    // 3: pb.Auth
}
var file_platform_proto_depIdxs = []int32{
	3, // 0: pb.PlatformProvider.auth:type_name -> pb.Auth
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_platform_proto_init() }
func file_platform_proto_init() {
	if File_platform_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformPlayer); i {
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
		file_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformProvider); i {
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
		file_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformProviderAPIPath); i {
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
			RawDescriptor: file_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_proto_goTypes,
		DependencyIndexes: file_platform_proto_depIdxs,
		MessageInfos:      file_platform_proto_msgTypes,
	}.Build()
	File_platform_proto = out.File
	file_platform_proto_rawDesc = nil
	file_platform_proto_goTypes = nil
	file_platform_proto_depIdxs = nil
}