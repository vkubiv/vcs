// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/trustbloc/edge-service/pkg/client/comparator/models"
)

// PostCompareReader is a Reader for the PostCompare structure.
type PostCompareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCompareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCompareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostCompareInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCompareOK creates a PostCompareOK with default headers values
func NewPostCompareOK() *PostCompareOK {
	return &PostCompareOK{}
}

/* PostCompareOK describes a response with status code 200, with default header values.

Result of comparison.
*/
type PostCompareOK struct {
	Payload *models.ComparisonResult
}

func (o *PostCompareOK) Error() string {
	return fmt.Sprintf("[POST /compare][%d] postCompareOK  %+v", 200, o.Payload)
}
func (o *PostCompareOK) GetPayload() *models.ComparisonResult {
	return o.Payload
}

func (o *PostCompareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComparisonResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCompareInternalServerError creates a PostCompareInternalServerError with default headers values
func NewPostCompareInternalServerError() *PostCompareInternalServerError {
	return &PostCompareInternalServerError{}
}

/* PostCompareInternalServerError describes a response with status code 500, with default header values.

Generic Error
*/
type PostCompareInternalServerError struct {
	Payload *models.Error
}

func (o *PostCompareInternalServerError) Error() string {
	return fmt.Sprintf("[POST /compare][%d] postCompareInternalServerError  %+v", 500, o.Payload)
}
func (o *PostCompareInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCompareInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
