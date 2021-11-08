// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package util

import (
	"context"
	"fmt"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"

	"github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/controller/adapter"
	ycrconfig "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/pkg/config"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/config"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/errorhandling"
)

func checkClusterMatchWithYmr(ymr *redis.Cluster, redisName, clusterName string) bool {
	cluster, ok1 := ymr.Labels[config.CloudClusterLabel]
	name, ok2 := ymr.Labels[config.CloudNameLabel]
	return ok1 && ok2 && cluster == clusterName && name == redisName
}

func GetRedisCluster(
	ctx context.Context, clusterID, folderID, metaName, clusterName string,
	ad adapter.YandexManagedRedisAdapter,
) (*redis.Cluster, error) {
	// If id is written in the status, we need to check
	// whether it exists in the cloud
	if clusterID != "" {
		ymr, err := ad.Read(ctx, clusterID)
		if err != nil {
			// If redis cluster was not found then it does not exist,
			// but this error is not fatal, just a mismatch between
			// out status and real world state.
			if !errorhandling.CheckRPCErrorNotFound(err) {
				return nil, fmt.Errorf("cannot get redis cluster from cloud: %w", err)
			}
		} else if checkClusterMatchWithYmr(ymr, metaName, clusterName) {
			// If labels do match with our object, then we have found it
			return ymr, nil
		}
		// Otherwise registry is not found, but that is ok:
		// we will try to list resources and find the one we need.
	}

	// TODO (covariance) pagination
	list, err := ad.List(ctx, folderID)
	if err != nil {
		// This error is fatal
		return nil, fmt.Errorf("cannot list redis clusters in folder: %w", err)
	}

	for _, res := range list {
		// If labels do match with our object, then we have found it
		if checkClusterMatchWithYmr(res, metaName, clusterName) {
			return res, nil
		}
	}

	// Nothing found, no such registry
	return nil, errorhandling.New("unable to find resource in the cloud", ycrconfig.ErrCodeYMRNotFound, nil)
}

func GbToBytes(gb int) (bytes int64) {
	return int64(gb * 1024 * 1024 * 1024) //nolint
}
