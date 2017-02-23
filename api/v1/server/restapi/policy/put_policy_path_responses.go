package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// HTTP code for type PutPolicyPathOK
const PutPolicyPathOKCode int = 200

/*PutPolicyPathOK Success

swagger:response putPolicyPathOK
*/
type PutPolicyPathOK struct {

	/*
	  In: Body
	*/
	Payload models.PolicyTree `json:"body,omitempty"`
}

// NewPutPolicyPathOK creates PutPolicyPathOK with default headers values
func NewPutPolicyPathOK() *PutPolicyPathOK {
	return &PutPolicyPathOK{}
}

// WithPayload adds the payload to the put policy path o k response
func (o *PutPolicyPathOK) WithPayload(payload models.PolicyTree) *PutPolicyPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy path o k response
func (o *PutPolicyPathOK) SetPayload(payload models.PolicyTree) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type PutPolicyPathInvalidPolicy
const PutPolicyPathInvalidPolicyCode int = 400

/*PutPolicyPathInvalidPolicy Invalid policy

swagger:response putPolicyPathInvalidPolicy
*/
type PutPolicyPathInvalidPolicy struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyPathInvalidPolicy creates PutPolicyPathInvalidPolicy with default headers values
func NewPutPolicyPathInvalidPolicy() *PutPolicyPathInvalidPolicy {
	return &PutPolicyPathInvalidPolicy{}
}

// WithPayload adds the payload to the put policy path invalid policy response
func (o *PutPolicyPathInvalidPolicy) WithPayload(payload models.Error) *PutPolicyPathInvalidPolicy {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy path invalid policy response
func (o *PutPolicyPathInvalidPolicy) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyPathInvalidPolicy) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type PutPolicyPathInvalidPath
const PutPolicyPathInvalidPathCode int = 460

/*PutPolicyPathInvalidPath Invalid path

swagger:response putPolicyPathInvalidPath
*/
type PutPolicyPathInvalidPath struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyPathInvalidPath creates PutPolicyPathInvalidPath with default headers values
func NewPutPolicyPathInvalidPath() *PutPolicyPathInvalidPath {
	return &PutPolicyPathInvalidPath{}
}

// WithPayload adds the payload to the put policy path invalid path response
func (o *PutPolicyPathInvalidPath) WithPayload(payload models.Error) *PutPolicyPathInvalidPath {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy path invalid path response
func (o *PutPolicyPathInvalidPath) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyPathInvalidPath) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(460)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type PutPolicyPathFailure
const PutPolicyPathFailureCode int = 500

/*PutPolicyPathFailure Policy import failed

swagger:response putPolicyPathFailure
*/
type PutPolicyPathFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyPathFailure creates PutPolicyPathFailure with default headers values
func NewPutPolicyPathFailure() *PutPolicyPathFailure {
	return &PutPolicyPathFailure{}
}

// WithPayload adds the payload to the put policy path failure response
func (o *PutPolicyPathFailure) WithPayload(payload models.Error) *PutPolicyPathFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy path failure response
func (o *PutPolicyPathFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyPathFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
