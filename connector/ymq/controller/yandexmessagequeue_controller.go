// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package controller

import (
	"context"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	connectorsv1 "k8s-connectors/connector/ymq/api/v1"
	"k8s-connectors/connector/ymq/controller/adapter"
	ymqconfig "k8s-connectors/connector/ymq/pkg/config"
	"k8s-connectors/pkg/config"
	"k8s-connectors/pkg/util"
)

// yandexMessageQueueReconciler reconciles a YandexContainerRegistry object
type yandexMessageQueueReconciler struct {
	client.Client
	adapter adapter.YandexMessageQueueAdapter
	log     logr.Logger
}

func NewYandexMessageQueueReconciler(
	cl client.Client, log logr.Logger,
) (*yandexMessageQueueReconciler, error) {
	impl, err := adapter.NewYandexMessageQueueAdapterSDK()
	if err != nil {
		return nil, err
	}
	return &yandexMessageQueueReconciler{
		Client:  cl,
		adapter: impl,
		log:     log,
	}, nil
}

// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmessagequeues,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmessagequeues/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmessagequeues/finalizers,verbs=update
// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=staticaccesskeys,verbs=get
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get

func (r *yandexMessageQueueReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.log.WithValues(ymqconfig.LongName, req.NamespacedName)

	// Try to retrieve object from k8s
	var object connectorsv1.YandexMessageQueue
	if err := r.Get(ctx, req.NamespacedName, &object); err != nil {
		// It still can be OK if we have not found it, and we do not need to reconcile it again

		// This outcome signifies that we just cannot find object, that is ok
		if apierrors.IsNotFound(err) {
			log.Info("Resource not found in k8s, reconciliation not possible")
			return config.GetNeverResult()
		}

		// Some unexpected error occurred, must throw
		return config.GetErroredResult(err)
	}

	// If object must be currently finalized, do it and quit
	mustBeFinalized, err := r.mustBeFinalized(&object)
	if err != nil {
		return config.GetErroredResult(err)
	}
	if mustBeFinalized {
		if err := r.finalize(ctx, log, &object); err != nil {
			return config.GetErroredResult(err)
		}
		return config.GetNormalResult()
	}

	if err := util.RegisterFinalizer(
		ctx, r.Client, log, &object.ObjectMeta, &object, ymqconfig.FinalizerName,
	); err != nil {
		return config.GetErroredResult(err)
	}

	if err := r.allocateResource(ctx, log, &object); err != nil {
		return config.GetErroredResult(err)
	}

	return config.GetNormalResult()
}

func (r *yandexMessageQueueReconciler) mustBeFinalized(object *connectorsv1.YandexMessageQueue) (bool, error) {
	return !object.DeletionTimestamp.IsZero() && util.ContainsString(
		object.Finalizers, ymqconfig.FinalizerName,
	), nil
}

func (r *yandexMessageQueueReconciler) finalize(
	ctx context.Context, log logr.Logger, object *connectorsv1.YandexMessageQueue,
) error {
	if err := r.deallocateResource(ctx, log, object); err != nil {
		return err
	}

	if err := util.DeregisterFinalizer(
		ctx, r.Client, log, &object.ObjectMeta, object, ymqconfig.FinalizerName,
	); err != nil {
		return err
	}

	log.Info("object finalized successfully")
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *yandexMessageQueueReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&connectorsv1.YandexMessageQueue{}).
		Complete(r)
}
