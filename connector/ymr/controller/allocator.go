// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"

	ymrconfig "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/pkg/config"

	connectorsv1 "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/api/v1"

	ymrutils "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/pkg/util"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/config"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/errorhandling"
)

func (r *yandexManagedRedisReconciler) allocateResource(
	ctx context.Context, log logr.Logger, object *connectorsv1.YandexManagedRedis,
) (*redis.Cluster, error) {
	log.V(1).Info("started")

	res, err := ymrutils.GetRedisCluster(
		ctx, object.Status.ID, object.Spec.FolderID, object.ObjectMeta.Name, r.clusterID, r.adapter,
	)
	if err == nil {
		return res, nil
	}

	if !errorhandling.CheckConnectorErrorCode(err, ymrconfig.ErrCodeYMRNotFound) {
		return nil, fmt.Errorf("unable to get resource: %w", err)
	}

	clusterEnvironment := redis.Cluster_ENVIRONMENT_UNSPECIFIED

	if object.Spec.Environment == "production" {
		clusterEnvironment = redis.Cluster_PRODUCTION
	}

	if object.Spec.Environment == "prestable" {
		clusterEnvironment = redis.Cluster_PRESTABLE
	}

	hostsSpecs := []*redis.HostSpec{}

	for _, host := range object.Spec.HostSpecs {
		hostsSpecs = append(hostsSpecs, &redis.HostSpec{
			ZoneId:    string(host.ZoneID),
			SubnetId:  string(host.SubnetID),
			ShardName: host.ShardName,
		})
	}

	diskSize := ymrutils.GbToBytes(int(object.Spec.DiskSize))
	clusterRequest := &redis.CreateClusterRequest{
		FolderId:    object.Spec.FolderID,
		Name:        object.Spec.Name,
		Description: object.Spec.Description,
		Labels: map[string]string{
			config.CloudClusterLabel: r.clusterID,
			config.CloudNameLabel:    object.Name,
		},
		Environment: clusterEnvironment,
		ConfigSpec: &redis.ConfigSpec{
			Version:   object.Spec.Version,
			RedisSpec: &redis.ConfigSpec_RedisConfig_6_0{},
			Resources: &redis.Resources{
				ResourcePresetId: string(object.Spec.HostClass),
				DiskSize:         diskSize,
				DiskTypeId:       string(object.Spec.DiskType),
			},
		},
		Sharded:   object.Spec.Sharded,
		HostSpecs: hostsSpecs,
	}
	resp, err := r.adapter.Create(ctx, clusterRequest)
	if err != nil {
		return nil, fmt.Errorf("unable to create resource: %w", err)
	}
	log.Info("successful")
	return resp, nil
}

func (r *yandexManagedRedisReconciler) deallocateResource(
	ctx context.Context, log logr.Logger, object *connectorsv1.YandexManagedRedis,
) error {
	log.V(1).Info("started")

	ycr, err := ymrutils.GetRedisCluster(
		ctx, object.Status.ID, object.Spec.FolderID, object.ObjectMeta.Name, r.clusterID, r.adapter,
	)
	if err != nil {
		if errorhandling.CheckConnectorErrorCode(err, ymrconfig.ErrCodeYMRNotFound) {
			log.Info("already deleted")
			return nil
		}
		return fmt.Errorf("unable to get resource: %w", err)
	}

	if err := r.adapter.Delete(ctx, ycr.Id); err != nil {
		return fmt.Errorf("unable to delete resource: %w", err)
	}
	log.Info("successful")
	return nil
}
