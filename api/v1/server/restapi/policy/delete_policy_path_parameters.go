package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePolicyPathParams creates a new DeletePolicyPathParams object
// with the default values initialized.
func NewDeletePolicyPathParams() DeletePolicyPathParams {
	var ()
	return DeletePolicyPathParams{}
}

// DeletePolicyPathParams contains all the bound params for the delete policy path operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeletePolicyPath
type DeletePolicyPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Path to policy node
	  Required: true
	  In: path
	*/
	Path string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeletePolicyPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePolicyPathParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Path = raw

	return nil
}
