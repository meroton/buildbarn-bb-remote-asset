// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.3
// source: pkg/proto/configuration/bb_remote_asset/bb_remote_asset.proto

package bb_remote_asset

import (
	fetch "github.com/buildbarn/bb-remote-asset/pkg/proto/configuration/bb_remote_asset/fetch"
	auth "github.com/buildbarn/bb-storage/pkg/proto/configuration/auth"
	blobstore "github.com/buildbarn/bb-storage/pkg/proto/configuration/blobstore"
	global "github.com/buildbarn/bb-storage/pkg/proto/configuration/global"
	grpc "github.com/buildbarn/bb-storage/pkg/proto/configuration/grpc"
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

type ApplicationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpcServers               []*grpc.ServerConfiguration        `protobuf:"bytes,3,rep,name=grpc_servers,json=grpcServers,proto3" json:"grpc_servers,omitempty"`
	ContentAddressableStorage *blobstore.BlobAccessConfiguration `protobuf:"bytes,4,opt,name=content_addressable_storage,json=contentAddressableStorage,proto3" json:"content_addressable_storage,omitempty"`
	MaximumMessageSizeBytes   int64                              `protobuf:"varint,5,opt,name=maximum_message_size_bytes,json=maximumMessageSizeBytes,proto3" json:"maximum_message_size_bytes,omitempty"`
	Global                    *global.Configuration              `protobuf:"bytes,6,opt,name=global,proto3" json:"global,omitempty"`
	AllowUpdatesForInstances  []string                           `protobuf:"bytes,7,rep,name=allow_updates_for_instances,json=allowUpdatesForInstances,proto3" json:"allow_updates_for_instances,omitempty"`
	Fetcher                   *fetch.FetcherConfiguration        `protobuf:"bytes,8,opt,name=fetcher,proto3" json:"fetcher,omitempty"`
	AssetCache                *AssetCacheConfiguration           `protobuf:"bytes,9,opt,name=asset_cache,json=assetCache,proto3" json:"asset_cache,omitempty"`
	FetchAuthorizer           *auth.AuthorizerConfiguration      `protobuf:"bytes,10,opt,name=fetch_authorizer,json=fetchAuthorizer,proto3" json:"fetch_authorizer,omitempty"`
	PushAuthorizer            *auth.AuthorizerConfiguration      `protobuf:"bytes,11,opt,name=push_authorizer,json=pushAuthorizer,proto3" json:"push_authorizer,omitempty"`
}

func (x *ApplicationConfiguration) Reset() {
	*x = ApplicationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfiguration) ProtoMessage() {}

func (x *ApplicationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfiguration.ProtoReflect.Descriptor instead.
func (*ApplicationConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationConfiguration) GetGrpcServers() []*grpc.ServerConfiguration {
	if x != nil {
		return x.GrpcServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetContentAddressableStorage() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.ContentAddressableStorage
	}
	return nil
}

func (x *ApplicationConfiguration) GetMaximumMessageSizeBytes() int64 {
	if x != nil {
		return x.MaximumMessageSizeBytes
	}
	return 0
}

func (x *ApplicationConfiguration) GetGlobal() *global.Configuration {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *ApplicationConfiguration) GetAllowUpdatesForInstances() []string {
	if x != nil {
		return x.AllowUpdatesForInstances
	}
	return nil
}

func (x *ApplicationConfiguration) GetFetcher() *fetch.FetcherConfiguration {
	if x != nil {
		return x.Fetcher
	}
	return nil
}

func (x *ApplicationConfiguration) GetAssetCache() *AssetCacheConfiguration {
	if x != nil {
		return x.AssetCache
	}
	return nil
}

func (x *ApplicationConfiguration) GetFetchAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.FetchAuthorizer
	}
	return nil
}

func (x *ApplicationConfiguration) GetPushAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.PushAuthorizer
	}
	return nil
}

type AssetCacheConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Backend:
	//
	//	*AssetCacheConfiguration_BlobAccess
	//	*AssetCacheConfiguration_ActionCache
	Backend isAssetCacheConfiguration_Backend `protobuf_oneof:"backend"`
}

func (x *AssetCacheConfiguration) Reset() {
	*x = AssetCacheConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetCacheConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetCacheConfiguration) ProtoMessage() {}

func (x *AssetCacheConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetCacheConfiguration.ProtoReflect.Descriptor instead.
func (*AssetCacheConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescGZIP(), []int{1}
}

func (m *AssetCacheConfiguration) GetBackend() isAssetCacheConfiguration_Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (x *AssetCacheConfiguration) GetBlobAccess() *blobstore.BlobAccessConfiguration {
	if x, ok := x.GetBackend().(*AssetCacheConfiguration_BlobAccess); ok {
		return x.BlobAccess
	}
	return nil
}

func (x *AssetCacheConfiguration) GetActionCache() *blobstore.BlobAccessConfiguration {
	if x, ok := x.GetBackend().(*AssetCacheConfiguration_ActionCache); ok {
		return x.ActionCache
	}
	return nil
}

type isAssetCacheConfiguration_Backend interface {
	isAssetCacheConfiguration_Backend()
}

type AssetCacheConfiguration_BlobAccess struct {
	BlobAccess *blobstore.BlobAccessConfiguration `protobuf:"bytes,1,opt,name=blob_access,json=blobAccess,proto3,oneof"`
}

type AssetCacheConfiguration_ActionCache struct {
	ActionCache *blobstore.BlobAccessConfiguration `protobuf:"bytes,2,opt,name=action_cache,json=actionCache,proto3,oneof"`
}

func (*AssetCacheConfiguration_BlobAccess) isAssetCacheConfiguration_Backend() {}

func (*AssetCacheConfiguration_ActionCache) isAssetCacheConfiguration_Backend() {}

var File_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x06, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x67,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x7a, 0x0a, 0x1b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x1b, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x46, 0x6f, 0x72,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x07, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x5e, 0x0a,
	0x0f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x70,
	0x75, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe4, 0x01, 0x0a, 0x17, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x62, 0x62, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescData = file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDesc
)

func file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescData)
	})
	return file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDescData
}

var file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_goTypes = []interface{}{
	(*ApplicationConfiguration)(nil),          // 0: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration
	(*AssetCacheConfiguration)(nil),           // 1: buildbarn.configuration.bb_remote_asset.AssetCacheConfiguration
	(*grpc.ServerConfiguration)(nil),          // 2: buildbarn.configuration.grpc.ServerConfiguration
	(*blobstore.BlobAccessConfiguration)(nil), // 3: buildbarn.configuration.blobstore.BlobAccessConfiguration
	(*global.Configuration)(nil),              // 4: buildbarn.configuration.global.Configuration
	(*fetch.FetcherConfiguration)(nil),        // 5: buildbarn.configuration.bb_remote_asset.fetch.FetcherConfiguration
	(*auth.AuthorizerConfiguration)(nil),      // 6: buildbarn.configuration.auth.AuthorizerConfiguration
}
var file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_depIdxs = []int32{
	2, // 0: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.grpc_servers:type_name -> buildbarn.configuration.grpc.ServerConfiguration
	3, // 1: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.content_addressable_storage:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	4, // 2: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.global:type_name -> buildbarn.configuration.global.Configuration
	5, // 3: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.fetcher:type_name -> buildbarn.configuration.bb_remote_asset.fetch.FetcherConfiguration
	1, // 4: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.asset_cache:type_name -> buildbarn.configuration.bb_remote_asset.AssetCacheConfiguration
	6, // 5: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.fetch_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	6, // 6: buildbarn.configuration.bb_remote_asset.ApplicationConfiguration.push_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	3, // 7: buildbarn.configuration.bb_remote_asset.AssetCacheConfiguration.blob_access:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	3, // 8: buildbarn.configuration.bb_remote_asset.AssetCacheConfiguration.action_cache:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_init() }
func file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_init() {
	if File_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationConfiguration); i {
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
		file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetCacheConfiguration); i {
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
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AssetCacheConfiguration_BlobAccess)(nil),
		(*AssetCacheConfiguration_ActionCache)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto = out.File
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_rawDesc = nil
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_goTypes = nil
	file_pkg_proto_configuration_bb_remote_asset_bb_remote_asset_proto_depIdxs = nil
}
