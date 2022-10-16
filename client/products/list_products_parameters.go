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

// NewListProductsParams creates a new ListProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProductsParams() *ListProductsParams {
	return &ListProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProductsParamsWithTimeout creates a new ListProductsParams object
// with the ability to set a timeout on a request.
func NewListProductsParamsWithTimeout(timeout time.Duration) *ListProductsParams {
	return &ListProductsParams{
		timeout: timeout,
	}
}

// NewListProductsParamsWithContext creates a new ListProductsParams object
// with the ability to set a context for a request.
func NewListProductsParamsWithContext(ctx context.Context) *ListProductsParams {
	return &ListProductsParams{
		Context: ctx,
	}
}

// NewListProductsParamsWithHTTPClient creates a new ListProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProductsParamsWithHTTPClient(client *http.Client) *ListProductsParams {
	return &ListProductsParams{
		HTTPClient: client,
	}
}

/* ListProductsParams contains all the parameters to send to the API endpoint
   for the list products operation.

   Typically these are written to a http.Request.
*/
type ListProductsParams struct {

	/* Currency.

	     Currency used when returning the price of the product,
	when not specified currency is returned in GBP.
	*/
	Currency *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductsParams) WithDefaults() *ListProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list products params
func (o *ListProductsParams) WithTimeout(timeout time.Duration) *ListProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list products params
func (o *ListProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list products params
func (o *ListProductsParams) WithContext(ctx context.Context) *ListProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list products params
func (o *ListProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list products params
func (o *ListProductsParams) WithHTTPClient(client *http.Client) *ListProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list products params
func (o *ListProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the list products params
func (o *ListProductsParams) WithCurrency(currency *string) *ListProductsParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the list products params
func (o *ListProductsParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WriteToRequest writes these params to a swagger request
func (o *ListProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Currency != nil {

		// query param Currency
		var qrCurrency string

		if o.Currency != nil {
			qrCurrency = *o.Currency
		}
		qCurrency := qrCurrency
		if qCurrency != "" {

			if err := r.SetQueryParam("Currency", qCurrency); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
