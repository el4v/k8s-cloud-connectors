// Code generated by sdkgen. DO NOT EDIT.

//nolint
package apigateway

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	apigateway "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
)

//revive:disable

// ApiGatewayServiceClient is a apigateway.ApiGatewayServiceClient with
// lazy GRPC connection initialization.
type ApiGatewayServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) Create(ctx context.Context, in *apigateway.CreateApiGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) Delete(ctx context.Context, in *apigateway.DeleteApiGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) Get(ctx context.Context, in *apigateway.GetApiGatewayRequest, opts ...grpc.CallOption) (*apigateway.ApiGateway, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).Get(ctx, in, opts...)
}

// GetOpenapiSpec implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) GetOpenapiSpec(ctx context.Context, in *apigateway.GetOpenapiSpecRequest, opts ...grpc.CallOption) (*apigateway.GetOpenapiSpecResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).GetOpenapiSpec(ctx, in, opts...)
}

// List implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) List(ctx context.Context, in *apigateway.ListApiGatewayRequest, opts ...grpc.CallOption) (*apigateway.ListApiGatewayResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).List(ctx, in, opts...)
}

type ApiGatewayIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ApiGatewayServiceClient
	request *apigateway.ListApiGatewayRequest

	items []*apigateway.ApiGateway
}

func (c *ApiGatewayServiceClient) ApiGatewayIterator(ctx context.Context, req *apigateway.ListApiGatewayRequest, opts ...grpc.CallOption) *ApiGatewayIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ApiGatewayIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ApiGatewayIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.ApiGateways
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ApiGatewayIterator) Take(size int64) ([]*apigateway.ApiGateway, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*apigateway.ApiGateway

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ApiGatewayIterator) TakeAll() ([]*apigateway.ApiGateway, error) {
	return it.Take(0)
}

func (it *ApiGatewayIterator) Value() *apigateway.ApiGateway {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ApiGatewayIterator) Error() error {
	return it.err
}

// ListAccessBindings implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type ApiGatewayAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ApiGatewayServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *ApiGatewayServiceClient) ApiGatewayAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *ApiGatewayAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ApiGatewayAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ApiGatewayAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ApiGatewayAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ApiGatewayAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *ApiGatewayAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ApiGatewayAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) ListOperations(ctx context.Context, in *apigateway.ListOperationsRequest, opts ...grpc.CallOption) (*apigateway.ListOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ApiGatewayOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ApiGatewayServiceClient
	request *apigateway.ListOperationsRequest

	items []*operation.Operation
}

func (c *ApiGatewayServiceClient) ApiGatewayOperationsIterator(ctx context.Context, req *apigateway.ListOperationsRequest, opts ...grpc.CallOption) *ApiGatewayOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ApiGatewayOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ApiGatewayOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ApiGatewayOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ApiGatewayOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *ApiGatewayOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ApiGatewayOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) Update(ctx context.Context, in *apigateway.UpdateApiGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements apigateway.ApiGatewayServiceClient
func (c *ApiGatewayServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return apigateway.NewApiGatewayServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
