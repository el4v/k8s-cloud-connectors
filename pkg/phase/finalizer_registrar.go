// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phase

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s-connectors/pkg/util"
)

func RegisterFinalizer(
	ctx context.Context, cl client.Client, log logr.Logger, meta *metav1.ObjectMeta, object client.Object,
	finalizer string,
) error {
	log.V(1).Info("started")
	if util.ContainsString(meta.Finalizers, finalizer) {
		return nil
	}
	meta.Finalizers = append(meta.Finalizers, finalizer)
	if err := cl.Update(ctx, object); err != nil {
		return fmt.Errorf("unable to register finalizer: %v", err)
	}
	log.Info("successful")
	return nil
}

func DeregisterFinalizer(
	ctx context.Context, cl client.Client, log logr.Logger, meta *metav1.ObjectMeta, object client.Object,
	finalizer string,
) error {
	log.V(1).Info("started")
	meta.Finalizers = util.RemoveString(meta.Finalizers, finalizer)
	if err := cl.Update(ctx, object); err != nil {
		return fmt.Errorf("unable to deregister finalizer: %v", err)
	}

	log.Info("successful")
	return nil
}
