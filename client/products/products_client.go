// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProdcut(params *CreateProdcutParams, opts ...ClientOption) (*CreateProdcutOK, error)

	DeleteProduct(params *DeleteProductParams, opts ...ClientOption) (*DeleteProductCreated, error)

	ListProducts(params *ListProductsParams, opts ...ClientOption) (*ListProductsOK, error)

	ListSingleProduct(params *ListSingleProductParams, opts ...ClientOption) (*ListSingleProductOK, error)

	UpdateProduct(params *UpdateProductParams, opts ...ClientOption) (*UpdateProductOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProdcut Create a new product
*/
func (a *Client) CreateProdcut(params *CreateProdcutParams, opts ...ClientOption) (*CreateProdcutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProdcutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createProdcut",
		Method:             "POST",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProdcutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProdcutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProdcut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProduct Update a products details
*/
func (a *Client) DeleteProduct(params *DeleteProductParams, opts ...ClientOption) (*DeleteProductCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProductParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteProduct",
		Method:             "DELETE",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProductCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListProducts Returns a list of products
*/
func (a *Client) ListProducts(params *ListProductsParams, opts ...ClientOption) (*ListProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProductsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListProductsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProductsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProducts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSingleProduct Returns a single product
*/
func (a *Client) ListSingleProduct(params *ListSingleProductParams, opts ...ClientOption) (*ListSingleProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSingleProductParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSingleProduct",
		Method:             "GET",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSingleProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSingleProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSingleProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProduct Create a new product
*/
func (a *Client) UpdateProduct(params *UpdateProductParams, opts ...ClientOption) (*UpdateProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProductParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateProduct",
		Method:             "PUT",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProductReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
