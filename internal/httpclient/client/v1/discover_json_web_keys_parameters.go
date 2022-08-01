// Code generated by go-swagger; DO NOT EDIT.

package v1

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

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// NewDiscoverJSONWebKeysParams creates a new DiscoverJSONWebKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscoverJSONWebKeysParams() *DiscoverJSONWebKeysParams {
	return &DiscoverJSONWebKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscoverJSONWebKeysParamsWithTimeout creates a new DiscoverJSONWebKeysParams object
// with the ability to set a timeout on a request.
func NewDiscoverJSONWebKeysParamsWithTimeout(timeout time.Duration) *DiscoverJSONWebKeysParams {
	return &DiscoverJSONWebKeysParams{
		timeout: timeout,
	}
}

// NewDiscoverJSONWebKeysParamsWithContext creates a new DiscoverJSONWebKeysParams object
// with the ability to set a context for a request.
func NewDiscoverJSONWebKeysParamsWithContext(ctx context.Context) *DiscoverJSONWebKeysParams {
	return &DiscoverJSONWebKeysParams{
		Context: ctx,
	}
}

// NewDiscoverJSONWebKeysParamsWithHTTPClient creates a new DiscoverJSONWebKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscoverJSONWebKeysParamsWithHTTPClient(client *http.Client) *DiscoverJSONWebKeysParams {
	return &DiscoverJSONWebKeysParams{
		HTTPClient: client,
	}
}

/* DiscoverJSONWebKeysParams contains all the parameters to send to the API endpoint
   for the discover Json web keys operation.

   Typically these are written to a http.Request.
*/
type DiscoverJSONWebKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the discover Json web keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscoverJSONWebKeysParams) WithDefaults() *DiscoverJSONWebKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discover Json web keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscoverJSONWebKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) WithTimeout(timeout time.Duration) *DiscoverJSONWebKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) WithContext(ctx context.Context) *DiscoverJSONWebKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) WithHTTPClient(client *http.Client) *DiscoverJSONWebKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discover Json web keys params
func (o *DiscoverJSONWebKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DiscoverJSONWebKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
