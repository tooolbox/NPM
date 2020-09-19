// Code generated by go-swagger; DO NOT EDIT.

package loyalty_and_rewards_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// LoyaltyDemographicStatusUpdateV1FvEmeaReader is a Reader for the LoyaltyDemographicStatusUpdateV1FvEmea structure.
type LoyaltyDemographicStatusUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltyDemographicStatusUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaOK creates a LoyaltyDemographicStatusUpdateV1FvEmeaOK with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaOK() *LoyaltyDemographicStatusUpdateV1FvEmeaOK {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaOK{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaOK struct {
	Payload *models.LoyaltyDemographicStatusUpdateResponse
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaOK) GetPayload() *models.LoyaltyDemographicStatusUpdateResponse {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltyDemographicStatusUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaBadRequest creates a LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaBadRequest() *LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized creates a LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized() *LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaForbidden creates a LoyaltyDemographicStatusUpdateV1FvEmeaForbidden with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaForbidden() *LoyaltyDemographicStatusUpdateV1FvEmeaForbidden {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaForbidden{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaNotFound creates a LoyaltyDemographicStatusUpdateV1FvEmeaNotFound with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaNotFound() *LoyaltyDemographicStatusUpdateV1FvEmeaNotFound {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaNotFound{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests creates a LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests() *LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus452 creates a LoyaltyDemographicStatusUpdateV1FvEmeaStatus452 with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus452() *LoyaltyDemographicStatusUpdateV1FvEmeaStatus452 {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaStatus452{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaStatus452 struct {
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaStatus452 ", 452)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus453 creates a LoyaltyDemographicStatusUpdateV1FvEmeaStatus453 with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus453() *LoyaltyDemographicStatusUpdateV1FvEmeaStatus453 {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaStatus453{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaStatus453 struct {
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaStatus453 ", 453)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus455 creates a LoyaltyDemographicStatusUpdateV1FvEmeaStatus455 with default headers values
func NewLoyaltyDemographicStatusUpdateV1FvEmeaStatus455() *LoyaltyDemographicStatusUpdateV1FvEmeaStatus455 {
	return &LoyaltyDemographicStatusUpdateV1FvEmeaStatus455{}
}

/*LoyaltyDemographicStatusUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltyDemographicStatusUpdateV1FvEmeaStatus455 struct {
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyDemographicStatusUpdate][%d] loyaltyDemographicStatusUpdateV1FvEmeaStatus455 ", 455)
}

func (o *LoyaltyDemographicStatusUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}