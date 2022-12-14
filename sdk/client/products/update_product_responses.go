// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"micro-services/sdk/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 433:
		result := NewUpdateProductStatus433()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductOK creates a UpdateProductOK with default headers values
func NewUpdateProductOK() *UpdateProductOK {
	return &UpdateProductOK{}
}

/* UpdateProductOK describes a response with status code 200, with default header values.

A list of products
*/
type UpdateProductOK struct {
	Payload []*models.Product
}

// IsSuccess returns true when this update product o k response has a 2xx status code
func (o *UpdateProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update product o k response has a 3xx status code
func (o *UpdateProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product o k response has a 4xx status code
func (o *UpdateProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update product o k response has a 5xx status code
func (o *UpdateProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update product o k response a status code equal to that given
func (o *UpdateProductOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProductOK) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductOK  %+v", 200, o.Payload)
}

func (o *UpdateProductOK) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductOK  %+v", 200, o.Payload)
}

func (o *UpdateProductOK) GetPayload() []*models.Product {
	return o.Payload
}

func (o *UpdateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/* UpdateProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type UpdateProductNotFound struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this update product not found response has a 2xx status code
func (o *UpdateProductNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update product not found response has a 3xx status code
func (o *UpdateProductNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product not found response has a 4xx status code
func (o *UpdateProductNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update product not found response has a 5xx status code
func (o *UpdateProductNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update product not found response a status code equal to that given
func (o *UpdateProductNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductStatus433 creates a UpdateProductStatus433 with default headers values
func NewUpdateProductStatus433() *UpdateProductStatus433 {
	return &UpdateProductStatus433{}
}

/* UpdateProductStatus433 describes a response with status code 433, with default header values.

Validation errors defined as an array of strings
*/
type UpdateProductStatus433 struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this update product status433 response has a 2xx status code
func (o *UpdateProductStatus433) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update product status433 response has a 3xx status code
func (o *UpdateProductStatus433) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product status433 response has a 4xx status code
func (o *UpdateProductStatus433) IsClientError() bool {
	return true
}

// IsServerError returns true when this update product status433 response has a 5xx status code
func (o *UpdateProductStatus433) IsServerError() bool {
	return false
}

// IsCode returns true when this update product status433 response a status code equal to that given
func (o *UpdateProductStatus433) IsCode(code int) bool {
	return code == 433
}

func (o *UpdateProductStatus433) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductStatus433  %+v", 433, o.Payload)
}

func (o *UpdateProductStatus433) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductStatus433  %+v", 433, o.Payload)
}

func (o *UpdateProductStatus433) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateProductStatus433) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
