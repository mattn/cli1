// Code generated by go-swagger; DO NOT EDIT.

package triggers

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
	"github.com/go-openapi/swag"
)

// NewListTriggersParams creates a new ListTriggersParams object
// with the default values initialized.
func NewListTriggersParams() *ListTriggersParams {
	var ()
	return &ListTriggersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTriggersParamsWithTimeout creates a new ListTriggersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTriggersParamsWithTimeout(timeout time.Duration) *ListTriggersParams {
	var ()
	return &ListTriggersParams{

		timeout: timeout,
	}
}

// NewListTriggersParamsWithContext creates a new ListTriggersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTriggersParamsWithContext(ctx context.Context) *ListTriggersParams {
	var ()
	return &ListTriggersParams{

		Context: ctx,
	}
}

// NewListTriggersParamsWithHTTPClient creates a new ListTriggersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTriggersParamsWithHTTPClient(client *http.Client) *ListTriggersParams {
	var ()
	return &ListTriggersParams{
		HTTPClient: client,
	}
}

/*ListTriggersParams contains all the parameters to send to the API endpoint
for the list triggers operation typically these are written to a http.Request
*/
type ListTriggersParams struct {

	/*AppID
	  Application ID.

	*/
	AppID *string
	/*Cursor
	  Cursor from previous response.next_cursor to begin results after, if any.

	*/
	Cursor *string
	/*FnID
	  Function ID.

	*/
	FnID *string
	/*Name
	  A Trigger name to filter by.

	*/
	Name *string
	/*PerPage
	  Number of results to return, defaults to 30. Max of 100.

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list triggers params
func (o *ListTriggersParams) WithTimeout(timeout time.Duration) *ListTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list triggers params
func (o *ListTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list triggers params
func (o *ListTriggersParams) WithContext(ctx context.Context) *ListTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list triggers params
func (o *ListTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list triggers params
func (o *ListTriggersParams) WithHTTPClient(client *http.Client) *ListTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list triggers params
func (o *ListTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the list triggers params
func (o *ListTriggersParams) WithAppID(appID *string) *ListTriggersParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the list triggers params
func (o *ListTriggersParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithCursor adds the cursor to the list triggers params
func (o *ListTriggersParams) WithCursor(cursor *string) *ListTriggersParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list triggers params
func (o *ListTriggersParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithFnID adds the fnID to the list triggers params
func (o *ListTriggersParams) WithFnID(fnID *string) *ListTriggersParams {
	o.SetFnID(fnID)
	return o
}

// SetFnID adds the fnId to the list triggers params
func (o *ListTriggersParams) SetFnID(fnID *string) {
	o.FnID = fnID
}

// WithName adds the name to the list triggers params
func (o *ListTriggersParams) WithName(name *string) *ListTriggersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list triggers params
func (o *ListTriggersParams) SetName(name *string) {
	o.Name = name
}

// WithPerPage adds the perPage to the list triggers params
func (o *ListTriggersParams) WithPerPage(perPage *int64) *ListTriggersParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list triggers params
func (o *ListTriggersParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *ListTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppID != nil {

		// query param app_id
		var qrAppID string
		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {
			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}

	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.FnID != nil {

		// query param fn_id
		var qrFnID string
		if o.FnID != nil {
			qrFnID = *o.FnID
		}
		qFnID := qrFnID
		if qFnID != "" {
			if err := r.SetQueryParam("fn_id", qFnID); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
