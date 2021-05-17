// Copyright (c) 2021 Yandex LLC. All rights reserved.
// Author: Martynov Pavel <covariance@yandex-team.ru>

package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sakey "k8s-connectors/connectors/sakey/api/v1"
	connectorsv1 "k8s-connectors/connectors/yos/api/v1"
	"k8s-connectors/connectors/yos/pkg/config"
)

type AwsSdkProvider = func(ctx context.Context, key string, secret string) (*s3.S3, error)

func NewStaticProvider() AwsSdkProvider {
	return func(_ context.Context, key string, secret string) (*s3.S3, error) {
		ses, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
			Endpoint:    aws.String(config.Endpoint),
			EndpointResolver: endpoints.ResolverFunc(
				func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
					return endpoints.ResolvedEndpoint{URL: config.Endpoint}, nil
				}),
			Region:           aws.String(config.AwsRegion),
			S3ForcePathStyle: aws.Bool(true),
		})

		if err != nil {
			return nil, fmt.Errorf("unable to get yos sdk: %v", err)
		}

		return s3.New(ses), nil
	}
}

func KeyAndSecretFromStaticAccessKey(ctx context.Context, bucket *connectorsv1.YandexObjectStorage, cl client.Client) (keyData, secretData string, err error) {
	var key sakey.StaticAccessKey
	if err = cl.Get(ctx, types.NamespacedName{
		Namespace: bucket.Namespace,
		Name:      bucket.Spec.SAKeyName,
	}, &key); err != nil {
		err = fmt.Errorf("unable to retrieve corresponding SAKey: %v", err)
		return
	}

	var secret v1.Secret
	if err = cl.Get(ctx, types.NamespacedName{
		Namespace: "default",
		Name:      key.Status.SecretName,
	}, &secret); err != nil {
		err = fmt.Errorf("unable to retrieve corresponding secret: %v", err)
		return
	}

	keyData = string(secret.Data["key"])
	secretData = string(secret.Data["secret"])

	return
}
