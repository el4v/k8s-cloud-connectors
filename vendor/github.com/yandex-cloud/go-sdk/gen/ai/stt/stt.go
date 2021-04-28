// Code generated by sdkgen. DO NOT EDIT.

//nolint
package stt

import (
	"context"

	"google.golang.org/grpc"

	stt "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// SttServiceClient is a stt.SttServiceClient with
// lazy GRPC connection initialization.
type SttServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// LongRunningRecognize implements stt.SttServiceClient
func (c *SttServiceClient) LongRunningRecognize(ctx context.Context, in *stt.LongRunningRecognitionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return stt.NewSttServiceClient(conn).LongRunningRecognize(ctx, in, opts...)
}

// StreamingRecognize implements stt.SttServiceClient
func (c *SttServiceClient) StreamingRecognize(ctx context.Context, opts ...grpc.CallOption) (stt.SttService_StreamingRecognizeClient, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return stt.NewSttServiceClient(conn).StreamingRecognize(ctx, opts...)
}
