// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/trustbloc/edge-service/pkg/client/comparator/models"
)

// NewPostCompareParams creates a new PostCompareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCompareParams() *PostCompareParams {
	return &PostCompareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCompareParamsWithTimeout creates a new PostCompareParams object
// with the ability to set a timeout on a request.
func NewPostCompareParamsWithTimeout(timeout time.Duration) *PostCompareParams {
	return &PostCompareParams{
		timeout: timeout,
	}
}

// NewPostCompareParamsWithContext creates a new PostCompareParams object
// with the ability to set a context for a request.
func NewPostCompareParamsWithContext(ctx context.Context) *PostCompareParams {
	return &PostCompareParams{
		Context: ctx,
	}
}

// NewPostCompareParamsWithHTTPClient creates a new PostCompareParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCompareParamsWithHTTPClient(client *http.Client) *PostCompareParams {
	return &PostCompareParams{
		HTTPClient: client,
	}
}

/* PostCompareParams contains all the parameters to send to the API endpoint
   for the post compare operation.

   Typically these are written to a http.Request.
*/
type PostCompareParams struct {

	// Comparison.
	Comparison *models.Comparison

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post compare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCompareParams) WithDefaults() *PostCompareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post compare params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCompareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post compare params
func (o *PostCompareParams) WithTimeout(timeout time.Duration) *PostCompareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post compare params
func (o *PostCompareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post compare params
func (o *PostCompareParams) WithContext(ctx context.Context) *PostCompareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post compare params
func (o *PostCompareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post compare params
func (o *PostCompareParams) WithHTTPClient(client *http.Client) *PostCompareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post compare params
func (o *PostCompareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComparison adds the comparison to the post compare params
func (o *PostCompareParams) WithComparison(comparison *models.Comparison) *PostCompareParams {
	o.SetComparison(comparison)
	return o
}

// SetComparison adds the comparison to the post compare params
func (o *PostCompareParams) SetComparison(comparison *models.Comparison) {
	o.Comparison = comparison
}

// WriteToRequest writes these params to a swagger request
func (o *PostCompareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Comparison != nil {
		if err := r.SetBodyParam(o.Comparison); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
