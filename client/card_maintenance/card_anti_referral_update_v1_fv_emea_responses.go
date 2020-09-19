// Code generated by go-swagger; DO NOT EDIT.

package card_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// CardAntiReferralUpdateV1FvEmeaReader is a Reader for the CardAntiReferralUpdateV1FvEmea structure.
type CardAntiReferralUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CardAntiReferralUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCardAntiReferralUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCardAntiReferralUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCardAntiReferralUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCardAntiReferralUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCardAntiReferralUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCardAntiReferralUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewCardAntiReferralUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewCardAntiReferralUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewCardAntiReferralUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewCardAntiReferralUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCardAntiReferralUpdateV1FvEmeaOK creates a CardAntiReferralUpdateV1FvEmeaOK with default headers values
func NewCardAntiReferralUpdateV1FvEmeaOK() *CardAntiReferralUpdateV1FvEmeaOK {
	return &CardAntiReferralUpdateV1FvEmeaOK{}
}

/*CardAntiReferralUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type CardAntiReferralUpdateV1FvEmeaOK struct {
	Payload *models.CardAntiReferralUpdateResponse
}

func (o *CardAntiReferralUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaOK) GetPayload() *models.CardAntiReferralUpdateResponse {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardAntiReferralUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaBadRequest creates a CardAntiReferralUpdateV1FvEmeaBadRequest with default headers values
func NewCardAntiReferralUpdateV1FvEmeaBadRequest() *CardAntiReferralUpdateV1FvEmeaBadRequest {
	return &CardAntiReferralUpdateV1FvEmeaBadRequest{}
}

/*CardAntiReferralUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type CardAntiReferralUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaUnauthorized creates a CardAntiReferralUpdateV1FvEmeaUnauthorized with default headers values
func NewCardAntiReferralUpdateV1FvEmeaUnauthorized() *CardAntiReferralUpdateV1FvEmeaUnauthorized {
	return &CardAntiReferralUpdateV1FvEmeaUnauthorized{}
}

/*CardAntiReferralUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type CardAntiReferralUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaForbidden creates a CardAntiReferralUpdateV1FvEmeaForbidden with default headers values
func NewCardAntiReferralUpdateV1FvEmeaForbidden() *CardAntiReferralUpdateV1FvEmeaForbidden {
	return &CardAntiReferralUpdateV1FvEmeaForbidden{}
}

/*CardAntiReferralUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type CardAntiReferralUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaNotFound creates a CardAntiReferralUpdateV1FvEmeaNotFound with default headers values
func NewCardAntiReferralUpdateV1FvEmeaNotFound() *CardAntiReferralUpdateV1FvEmeaNotFound {
	return &CardAntiReferralUpdateV1FvEmeaNotFound{}
}

/*CardAntiReferralUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type CardAntiReferralUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaTooManyRequests creates a CardAntiReferralUpdateV1FvEmeaTooManyRequests with default headers values
func NewCardAntiReferralUpdateV1FvEmeaTooManyRequests() *CardAntiReferralUpdateV1FvEmeaTooManyRequests {
	return &CardAntiReferralUpdateV1FvEmeaTooManyRequests{}
}

/*CardAntiReferralUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type CardAntiReferralUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaStatus452 creates a CardAntiReferralUpdateV1FvEmeaStatus452 with default headers values
func NewCardAntiReferralUpdateV1FvEmeaStatus452() *CardAntiReferralUpdateV1FvEmeaStatus452 {
	return &CardAntiReferralUpdateV1FvEmeaStatus452{}
}

/*CardAntiReferralUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type CardAntiReferralUpdateV1FvEmeaStatus452 struct {
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaStatus452 ", 452)
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaStatus453 creates a CardAntiReferralUpdateV1FvEmeaStatus453 with default headers values
func NewCardAntiReferralUpdateV1FvEmeaStatus453() *CardAntiReferralUpdateV1FvEmeaStatus453 {
	return &CardAntiReferralUpdateV1FvEmeaStatus453{}
}

/*CardAntiReferralUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type CardAntiReferralUpdateV1FvEmeaStatus453 struct {
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaStatus453 ", 453)
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaStatus455 creates a CardAntiReferralUpdateV1FvEmeaStatus455 with default headers values
func NewCardAntiReferralUpdateV1FvEmeaStatus455() *CardAntiReferralUpdateV1FvEmeaStatus455 {
	return &CardAntiReferralUpdateV1FvEmeaStatus455{}
}

/*CardAntiReferralUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type CardAntiReferralUpdateV1FvEmeaStatus455 struct {
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaStatus455 ", 455)
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardAntiReferralUpdateV1FvEmeaStatus465 creates a CardAntiReferralUpdateV1FvEmeaStatus465 with default headers values
func NewCardAntiReferralUpdateV1FvEmeaStatus465() *CardAntiReferralUpdateV1FvEmeaStatus465 {
	return &CardAntiReferralUpdateV1FvEmeaStatus465{}
}

/*CardAntiReferralUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SFR01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 000-998<BR/>VPL5SFR02E : CARD NUMBER IS NOT NUMERIC OR EQUALS ZEROES<BR/>VPL5SFR03E : CARD SEQUENCE NUMBER MUST BE NUMERIC VALID VALUES ARE 0001-9998<BR/>VPL5SFR04E : INVALID FOREIGN USE INDICATOR VALID VALUES ARE 0 OR 1<BR/>VPL5SFR05S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SFR06S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SFR07S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SFR08S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SFR09S : SYSTEM RECORD NOT FOUND<BR/>VPL5SFR10S : ORGANIZATION FILE UNAVAILABLE<BR/>VPL5SFR11S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SFR12I : NO FOREIGN ORG FOUND IN ORG RECORD<BR/>VPL5SFR14S : LOGO NUMBER NOT FOUND<BR/>VPL5SFR15S : EMBOSSER CARD FILE UNAVAILABLE<BR/>VPL5SFR16S : CARD NUMBER NOT FOUND<BR/>VPL5SFR17S : ACCOUNT BASE SEGMENT FILE UNAVAILABLE<BR/>VPL5SFR18S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SFR19S : CARD RELATIONSHIP FILE UNAVAILABLE<BR/>VPL5SFR20S : REFER ANTI REFER FLAG VALID VALUES ARE 01 AND 2<BR/>VPL5SFR21S : EXPIRY DATE FORMAT SHOULD BE YYYYMMDD<BR/>VPL5SFR22S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SFR23S : SERVICE INPUT TO CARD UPDATE SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SFR25S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SFR26S : FRAUD SWITCH IS INACTIVE AT SYSTEM OR ORG OR LOGO LEVEL<BR/>VPL5SFR27S : REFER ANTI REFER EXPIRY DATE MUST BE ENTERED<BR/>VPL5SFR28S : EXPIRY DATE MUST BE GREATER THAN OR EQUAL TO TODAYS DAT
*/
type CardAntiReferralUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardAntiReferralUpdate][%d] cardAntiReferralUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardAntiReferralUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}