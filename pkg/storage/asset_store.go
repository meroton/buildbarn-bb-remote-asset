package storage

import (
	"context"

	"github.com/buildbarn/bb-remote-asset/pkg/proto/asset"
	"github.com/buildbarn/bb-storage/pkg/digest"
)

// AssetStore is a wrapper around a BlobAccess to inteface well with
// AssetReference messages
type AssetStore interface {
	Get(ctx context.Context, ref *asset.AssetReference, instance digest.InstanceName) (*asset.Asset, error)
	Put(ctx context.Context, ref *asset.AssetReference, data *asset.Asset, instance digest.InstanceName) error
}
