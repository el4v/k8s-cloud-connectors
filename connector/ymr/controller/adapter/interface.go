package adapter

import (
	"context"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
)

type YandexManagedRedisAdapter interface {
	Create(ctx context.Context, request *redis.CreateClusterRequest) (*redis.Cluster, error)
	Read(ctx context.Context, clusterID string) (*redis.Cluster, error)
	List(ctx context.Context, folderID string) ([]*redis.Cluster, error)
	Update(ctx context.Context, request *redis.UpdateClusterRequest) error
	Delete(ctx context.Context, clusterID string) error
}
