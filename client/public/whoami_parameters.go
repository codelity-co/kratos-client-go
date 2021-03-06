// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewWhoamiParams creates a new WhoamiParams object
// with the default values initialized.
func NewWhoamiParams() *WhoamiParams {
	var ()
	return &WhoamiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWhoamiParamsWithTimeout creates a new WhoamiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWhoamiParamsWithTimeout(timeout time.Duration) *WhoamiParams {
	var ()
	return &WhoamiParams{

		timeout: timeout,
	}
}

// NewWhoamiParamsWithContext creates a new WhoamiParams object
// with the default values initialized, and the ability to set a context for a request
func NewWhoamiParamsWithContext(ctx context.Context) *WhoamiParams {
	var ()
	return &WhoamiParams{

		Context: ctx,
	}
}

// NewWhoamiParamsWithHTTPClient creates a new WhoamiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWhoamiParamsWithHTTPClient(client *http.Client) *WhoamiParams {
	var ()
	return &WhoamiParams{
		HTTPClient: client,
	}
}

/*WhoamiParams contains all the parameters to send to the API endpoint
for the whoami operation typically these are written to a http.Request
*/
type WhoamiParams struct {

	/*Authorization
	  in: authorization

	*/
	Authorization *string
	/*Cookie*/
	Cookie *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the whoami params
func (o *WhoamiParams) WithTimeout(timeout time.Duration) *WhoamiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the whoami params
func (o *WhoamiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the whoami params
func (o *WhoamiParams) WithContext(ctx context.Context) *WhoamiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the whoami params
func (o *WhoamiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the whoami params
func (o *WhoamiParams) WithHTTPClient(client *http.Client) *WhoamiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the whoami params
func (o *WhoamiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the whoami params
func (o *WhoamiParams) WithAuthorization(authorization *string) *WhoamiParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the whoami params
func (o *WhoamiParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithCookie adds the cookie to the whoami params
func (o *WhoamiParams) WithCookie(cookie *string) *WhoamiParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the whoami params
func (o *WhoamiParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WriteToRequest writes these params to a swagger request
func (o *WhoamiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// query param Authorization
		var qrAuthorization string
		if o.Authorization != nil {
			qrAuthorization = *o.Authorization
		}
		qAuthorization := qrAuthorization
		if qAuthorization != "" {
			if err := r.SetQueryParam("Authorization", qAuthorization); err != nil {
				return err
			}
		}

	}

	if o.Cookie != nil {

		// header param Cookie
		if err := r.SetHeaderParam("Cookie", *o.Cookie); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
