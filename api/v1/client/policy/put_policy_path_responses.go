package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PutPolicyPathReader is a Reader for the PutPolicyPath structure.
type PutPolicyPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPolicyPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutPolicyPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutPolicyPathInvalidPolicy()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 460:
		result := NewPutPolicyPathInvalidPath()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutPolicyPathFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPolicyPathOK creates a PutPolicyPathOK with default headers values
func NewPutPolicyPathOK() *PutPolicyPathOK {
	return &PutPolicyPathOK{}
}

/*PutPolicyPathOK handles this case with default header values.

Success
*/
type PutPolicyPathOK struct {
	Payload models.PolicyTree
}

func (o *PutPolicyPathOK) Error() string {
	return fmt.Sprintf("[PUT /policy/{path}][%d] putPolicyPathOK  %+v", 200, o.Payload)
}

func (o *PutPolicyPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyPathInvalidPolicy creates a PutPolicyPathInvalidPolicy with default headers values
func NewPutPolicyPathInvalidPolicy() *PutPolicyPathInvalidPolicy {
	return &PutPolicyPathInvalidPolicy{}
}

/*PutPolicyPathInvalidPolicy handles this case with default header values.

Invalid policy
*/
type PutPolicyPathInvalidPolicy struct {
	Payload models.Error
}

func (o *PutPolicyPathInvalidPolicy) Error() string {
	return fmt.Sprintf("[PUT /policy/{path}][%d] putPolicyPathInvalidPolicy  %+v", 400, o.Payload)
}

func (o *PutPolicyPathInvalidPolicy) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyPathInvalidPath creates a PutPolicyPathInvalidPath with default headers values
func NewPutPolicyPathInvalidPath() *PutPolicyPathInvalidPath {
	return &PutPolicyPathInvalidPath{}
}

/*PutPolicyPathInvalidPath handles this case with default header values.

Invalid path
*/
type PutPolicyPathInvalidPath struct {
	Payload models.Error
}

func (o *PutPolicyPathInvalidPath) Error() string {
	return fmt.Sprintf("[PUT /policy/{path}][%d] putPolicyPathInvalidPath  %+v", 460, o.Payload)
}

func (o *PutPolicyPathInvalidPath) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyPathFailure creates a PutPolicyPathFailure with default headers values
func NewPutPolicyPathFailure() *PutPolicyPathFailure {
	return &PutPolicyPathFailure{}
}

/*PutPolicyPathFailure handles this case with default header values.

Policy import failed
*/
type PutPolicyPathFailure struct {
	Payload models.Error
}

func (o *PutPolicyPathFailure) Error() string {
	return fmt.Sprintf("[PUT /policy/{path}][%d] putPolicyPathFailure  %+v", 500, o.Payload)
}

func (o *PutPolicyPathFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
