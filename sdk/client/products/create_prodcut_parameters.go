// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateProdcutParams creates a new CreateProdcutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProdcutParams() *CreateProdcutParams {
	return &CreateProdcutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProdcutParamsWithTimeout creates a new CreateProdcutParams object
// with the ability to set a timeout on a request.
func NewCreateProdcutParamsWithTimeout(timeout time.Duration) *CreateProdcutParams {
	return &CreateProdcutParams{
		timeout: timeout,
	}
}

// NewCreateProdcutParamsWithContext creates a new CreateProdcutParams object
// with the ability to set a context for a request.
func NewCreateProdcutParamsWithContext(ctx context.Context) *CreateProdcutParams {
	return &CreateProdcutParams{
		Context: ctx,
	}
}

// NewCreateProdcutParamsWithHTTPClient creates a new CreateProdcutParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProdcutParamsWithHTTPClient(client *http.Client) *CreateProdcutParams {
	return &CreateProdcutParams{
		HTTPClient: client,
	}
}

/* CreateProdcutParams contains all the parameters to send to the API endpoint
   for the create prodcut operation.

   Typically these are written to a http.Request.
*/
type CreateProdcutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create prodcut params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProdcutParams) WithDefaults() *CreateProdcutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create prodcut params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProdcutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create prodcut params
func (o *CreateProdcutParams) WithTimeout(timeout time.Duration) *CreateProdcutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create prodcut params
func (o *CreateProdcutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create prodcut params
func (o *CreateProdcutParams) WithContext(ctx context.Context) *CreateProdcutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create prodcut params
func (o *CreateProdcutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create prodcut params
func (o *CreateProdcutParams) WithHTTPClient(client *http.Client) *CreateProdcutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create prodcut params
func (o *CreateProdcutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProdcutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
