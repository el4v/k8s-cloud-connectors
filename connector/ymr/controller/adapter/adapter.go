package adapter

import (
	"context"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type YandexManagedRedisAdapterSDK struct {
	sdk *ycsdk.SDK
}

func NewYandexManagedRedisAdapterSDK(sdk *ycsdk.SDK) YandexManagedRedisAdapterSDK {
	return YandexManagedRedisAdapterSDK{
		sdk: sdk,
	}
}

func (r YandexManagedRedisAdapterSDK) Create(ctx context.Context,
	request *redis.CreateClusterRequest) (*redis.Cluster, error) {
	op, err := r.sdk.WrapOperation(r.sdk.MDB().Redis().Cluster().Create(ctx, request))
	if err != nil {
		return nil, err
	}

	if err := op.Wait(ctx); err != nil {
		return nil, err
	}

	res, err := op.Response()
	if err != nil {
		return nil, err
	}

	return res.(*redis.Cluster), nil
}

func (r YandexManagedRedisAdapterSDK) Read(ctx context.Context, clusterID string) (*redis.Cluster, error) {
	return r.sdk.MDB().Redis().Cluster().Get(
		ctx, &redis.GetClusterRequest{
			ClusterId: clusterID,
		},
	)
}

func (r YandexManagedRedisAdapterSDK) List(ctx context.Context, folderID string) ([]*redis.Cluster, error) {
	list, err := r.sdk.MDB().Redis().Cluster().List(
		ctx, &redis.ListClustersRequest{
			FolderId: folderID,
		},
	)
	if err != nil {
		return nil, err
	}
	return list.Clusters, nil
}

func (r YandexManagedRedisAdapterSDK) Update(ctx context.Context, request *redis.UpdateClusterRequest) error {
	op, err := r.sdk.WrapOperation(r.sdk.MDB().Redis().Cluster().Update(ctx, request))
	if err != nil {
		return err
	}
	if err := op.Wait(ctx); err != nil {
		return err
	}
	if _, err := op.Response(); err != nil {
		return err
	}

	return nil
}

func (r YandexManagedRedisAdapterSDK) Delete(ctx context.Context, clusterID string) error {
	op, err := r.sdk.WrapOperation(
		r.sdk.MDB().Redis().Cluster().Delete(
			ctx, &redis.DeleteClusterRequest{
				ClusterId: clusterID,
			},
		),
	)
	if err != nil {
		return err
	}
	if err := op.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func NewyandexmanagedredisAdapter() (YandexManagedRedisAdapter, error) {
	return YandexManagedRedisAdapterSDK{}, nil
}
