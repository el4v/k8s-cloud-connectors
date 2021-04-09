// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package phases

import (
	"context"
	"github.com/go-logr/logr"
	connectorsv1 "k8s-connectors/connectors/awskey/api/v1"
)

type AWSAccessKeyPhase interface {
	IsUpdated(context.Context, *connectorsv1.AWSAccessKey) (bool, error)
	Update(context.Context, logr.Logger, *connectorsv1.AWSAccessKey) error
	Cleanup(context.Context, logr.Logger, *connectorsv1.AWSAccessKey) error
}
