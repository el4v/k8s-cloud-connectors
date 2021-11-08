package controller

import (
	"context"
	"fmt"

	ycsdk "github.com/yandex-cloud/go-sdk"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	connectorsv1 "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/api/v1"

	"github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/controller/adapter"
	ymrconfig "github.com/yandex-cloud/k8s-cloud-connectors/connector/ymr/pkg/config"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/config"
	"github.com/yandex-cloud/k8s-cloud-connectors/pkg/phase"
)

// yandexManagedRedisReconciler reconciles a yandexmanagedredis object
type yandexManagedRedisReconciler struct {
	client.Client
	adapter   adapter.YandexManagedRedisAdapter
	log       logr.Logger
	clusterID string
}

func NewYandexManagedRedisReconciler(
	log logr.Logger, cl client.Client,
	sdk *ycsdk.SDK, clusterID string,
) (*yandexManagedRedisReconciler, error) {
	return &yandexManagedRedisReconciler{
		Client:    cl,
		adapter:   adapter.NewYandexManagedRedisAdapterSDK(sdk),
		log:       log,
		clusterID: clusterID,
	}, nil
}

// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmanagedrediss,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmanagedrediss/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=connectors.cloud.yandex.com,resources=yandexmanagedrediss/finalizers,verbs=update

func (r *yandexManagedRedisReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.log.WithValues("name", req.NamespacedName)
	log.V(1).Info("started reconciliation")

	// Try to retrieve object from k8s
	var object connectorsv1.YandexManagedRedis
	if err := r.Get(ctx, req.NamespacedName, &object); err != nil {
		// This outcome signifies that we just cannot find object, that is OK,
		// we just never want to reconcile it again unless triggered externally.
		if apierrors.IsNotFound(err) {
			log.V(1).Info("object not found in k8s, reconciliation not possible")
			return config.GetNeverResult()
		}

		return config.GetErroredResult(fmt.Errorf("unable to get object from k8s: %w", err))
	}

	// If object must be currently finalized, do it and quit
	if phase.MustBeFinalized(&object.ObjectMeta, ymrconfig.FinalizerName) {
		if err := r.finalize(ctx, log.WithName("finalize"), &object); err != nil {
			return config.GetErroredResult(fmt.Errorf("unable to finalize object: %w", err))
		}
		return config.GetNormalResult()
	}

	if err := phase.RegisterFinalizer(
		ctx, r.Client, log.WithName("register-finalizer"), &object.ObjectMeta, &object, ymrconfig.FinalizerName,
	); err != nil {
		return config.GetErroredResult(fmt.Errorf("unable to register finalizer: %w", err))
	}

	res, err := r.allocateResource(ctx, log.WithName("allocate-resource"), &object)
	if err != nil {
		return config.GetErroredResult(fmt.Errorf("unable to allocate resource: %w", err))
	}

	//! TODO: update redis cluster
	// if err := r.matchSpec(ctx, log.WithName("match-spec"), &object, res); err != nil { //nolint
	// 	return config.GetErroredResult(fmt.Errorf("unable to match spec: %w", err)) //nolint
	// } //nolint

	if err := r.updateStatus(ctx, log.WithName("update-status"), &object, res); err != nil {
		return config.GetErroredResult(fmt.Errorf("unable to update status: %w", err))
	}

	if err := phase.ProvideConfigmap(
		ctx,
		r.Client,
		log.WithName("provide-configmap"),
		object.Name, ymrconfig.ShortName, object.Namespace,
		map[string]string{"ID": object.Status.ID},
	); err != nil {
		return config.GetErroredResult(fmt.Errorf("unable to provide configmap: %w", err))
	}

	log.V(1).Info("finished reconciliation")
	return config.GetNormalResult()
}

func (r *yandexManagedRedisReconciler) finalize(
	ctx context.Context, log logr.Logger, object *connectorsv1.YandexManagedRedis,
) error {
	log.V(1).Info("started")

	if err := phase.RemoveConfigmap(
		ctx,
		r.Client,
		log.WithName("remove-configmap"),
		object.Name, ymrconfig.ShortName, object.Namespace,
	); err != nil {
		return fmt.Errorf("unable to remove configmap: %w", err)
	}

	if err := r.deallocateResource(ctx, log.WithName("deallocate-resource"), object); err != nil {
		return fmt.Errorf("unable to deallocate resource: %w", err)
	}

	if err := phase.DeregisterFinalizer(
		ctx, r.Client, log.WithName("deregister-finalizer"), &object.ObjectMeta, object, ymrconfig.FinalizerName,
	); err != nil {
		return fmt.Errorf("unable to deregister finalizer: %w", err)
	}

	log.Info("successful")
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *yandexManagedRedisReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&connectorsv1.YandexManagedRedis{}).
		Complete(r)
}
