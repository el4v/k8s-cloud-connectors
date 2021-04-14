// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phases

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	ycrconfig "k8s-connectors/connectors/ycr/pkg/config"
	connectorsv1 "k8s-connectors/connectors/yos/api/v1"
	yosconfig "k8s-connectors/connectors/yos/pkg/config"
	"k8s-connectors/pkg/utils"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FinalizerRegistrar struct {
	Client *client.Client
}

func (r *FinalizerRegistrar) IsUpdated(_ context.Context, resource *connectorsv1.YandexObjectStorage) (bool, error) {
	return utils.ContainsString(resource.Finalizers, yosconfig.FinalizerName), nil
}

func (r *FinalizerRegistrar) Update(ctx context.Context, log logr.Logger, resource *connectorsv1.YandexObjectStorage) error {
	resource.Finalizers = append(resource.Finalizers, yosconfig.FinalizerName)
	if err := (*r.Client).Update(ctx, resource); err != nil {
		return fmt.Errorf("unable to update resource status: %v", err)
	}
	log.Info("finalizer registered successfully")
	return nil
}

func (r *FinalizerRegistrar) Cleanup(ctx context.Context, log logr.Logger, resource *connectorsv1.YandexObjectStorage) error {
	resource.Finalizers = utils.RemoveString(resource.Finalizers, ycrconfig.FinalizerName)
	if err := (*r.Client).Update(ctx, resource); err != nil {
		return fmt.Errorf("unable to remove finalizer: %v", err)
	}

	log.Info("finalizer removed successfully")
	return nil
}
