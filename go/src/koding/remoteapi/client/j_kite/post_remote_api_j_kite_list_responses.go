package j_kite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJKiteListReader is a Reader for the PostRemoteAPIJKiteList structure.
type PostRemoteAPIJKiteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJKiteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJKiteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJKiteListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJKiteListOK creates a PostRemoteAPIJKiteListOK with default headers values
func NewPostRemoteAPIJKiteListOK() *PostRemoteAPIJKiteListOK {
	return &PostRemoteAPIJKiteListOK{}
}

/*PostRemoteAPIJKiteListOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJKiteListOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJKiteListOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JKite.list][%d] postRemoteApiJKiteListOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJKiteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJKiteListUnauthorized creates a PostRemoteAPIJKiteListUnauthorized with default headers values
func NewPostRemoteAPIJKiteListUnauthorized() *PostRemoteAPIJKiteListUnauthorized {
	return &PostRemoteAPIJKiteListUnauthorized{}
}

/*PostRemoteAPIJKiteListUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJKiteListUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJKiteListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JKite.list][%d] postRemoteApiJKiteListUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJKiteListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
