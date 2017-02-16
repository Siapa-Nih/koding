package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIPaymentFetchGroupInvoicesReader is a Reader for the PostRemoteAPIPaymentFetchGroupInvoices structure.
type PostRemoteAPIPaymentFetchGroupInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIPaymentFetchGroupInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIPaymentFetchGroupInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIPaymentFetchGroupInvoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIPaymentFetchGroupInvoicesOK creates a PostRemoteAPIPaymentFetchGroupInvoicesOK with default headers values
func NewPostRemoteAPIPaymentFetchGroupInvoicesOK() *PostRemoteAPIPaymentFetchGroupInvoicesOK {
	return &PostRemoteAPIPaymentFetchGroupInvoicesOK{}
}

/*PostRemoteAPIPaymentFetchGroupInvoicesOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIPaymentFetchGroupInvoicesOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIPaymentFetchGroupInvoicesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Payment.fetchGroupInvoices][%d] postRemoteApiPaymentFetchGroupInvoicesOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIPaymentFetchGroupInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIPaymentFetchGroupInvoicesUnauthorized creates a PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized with default headers values
func NewPostRemoteAPIPaymentFetchGroupInvoicesUnauthorized() *PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized {
	return &PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized{}
}

/*PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Payment.fetchGroupInvoices][%d] postRemoteApiPaymentFetchGroupInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIPaymentFetchGroupInvoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
