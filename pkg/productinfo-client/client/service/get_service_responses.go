// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/banzaicloud/productinfo/pkg/productinfo-client/models"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 503:
		result := NewGetServiceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

ServiceResponse
*/
type GetServiceOK struct {
	Payload *models.ServiceResponse
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /providers/{provider}/services/{service}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceServiceUnavailable creates a GetServiceServiceUnavailable with default headers values
func NewGetServiceServiceUnavailable() *GetServiceServiceUnavailable {
	return &GetServiceServiceUnavailable{}
}

/*GetServiceServiceUnavailable handles this case with default header values.

ErrorResponse
*/
type GetServiceServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *GetServiceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /providers/{provider}/services/{service}][%d] getServiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetServiceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
