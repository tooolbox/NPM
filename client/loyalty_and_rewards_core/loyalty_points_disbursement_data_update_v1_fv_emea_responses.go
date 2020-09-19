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

// LoyaltyPointsDisbursementDataUpdateV1FvEmeaReader is a Reader for the LoyaltyPointsDisbursementDataUpdateV1FvEmea structure.
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaOK creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaOK() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK struct {
	Payload *models.LoyaltyPointsDisbursementDataUpdateResponse
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK) GetPayload() *models.LoyaltyPointsDisbursementDataUpdateResponse {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltyPointsDisbursementDataUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 struct {
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452 ", 452)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 struct {
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453 ", 453)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 creates a LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 with default headers values
func NewLoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455() *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 {
	return &LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455{}
}

/*LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 struct {
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsDisbursementDataUpdate][%d] loyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455 ", 455)
}

func (o *LoyaltyPointsDisbursementDataUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}