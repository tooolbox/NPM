// Code generated by go-swagger; DO NOT EDIT.

package card_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// AcsCardInquiryV1FvEmeaReader is a Reader for the AcsCardInquiryV1FvEmea structure.
type AcsCardInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcsCardInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcsCardInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAcsCardInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAcsCardInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcsCardInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcsCardInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAcsCardInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAcsCardInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAcsCardInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAcsCardInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewAcsCardInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcsCardInquiryV1FvEmeaOK creates a AcsCardInquiryV1FvEmeaOK with default headers values
func NewAcsCardInquiryV1FvEmeaOK() *AcsCardInquiryV1FvEmeaOK {
	return &AcsCardInquiryV1FvEmeaOK{}
}

/*AcsCardInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type AcsCardInquiryV1FvEmeaOK struct {
	Payload *models.AcsCardInquiryResponse
}

func (o *AcsCardInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaOK) GetPayload() *models.AcsCardInquiryResponse {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AcsCardInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaBadRequest creates a AcsCardInquiryV1FvEmeaBadRequest with default headers values
func NewAcsCardInquiryV1FvEmeaBadRequest() *AcsCardInquiryV1FvEmeaBadRequest {
	return &AcsCardInquiryV1FvEmeaBadRequest{}
}

/*AcsCardInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AcsCardInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaUnauthorized creates a AcsCardInquiryV1FvEmeaUnauthorized with default headers values
func NewAcsCardInquiryV1FvEmeaUnauthorized() *AcsCardInquiryV1FvEmeaUnauthorized {
	return &AcsCardInquiryV1FvEmeaUnauthorized{}
}

/*AcsCardInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AcsCardInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaForbidden creates a AcsCardInquiryV1FvEmeaForbidden with default headers values
func NewAcsCardInquiryV1FvEmeaForbidden() *AcsCardInquiryV1FvEmeaForbidden {
	return &AcsCardInquiryV1FvEmeaForbidden{}
}

/*AcsCardInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AcsCardInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaNotFound creates a AcsCardInquiryV1FvEmeaNotFound with default headers values
func NewAcsCardInquiryV1FvEmeaNotFound() *AcsCardInquiryV1FvEmeaNotFound {
	return &AcsCardInquiryV1FvEmeaNotFound{}
}

/*AcsCardInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AcsCardInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaTooManyRequests creates a AcsCardInquiryV1FvEmeaTooManyRequests with default headers values
func NewAcsCardInquiryV1FvEmeaTooManyRequests() *AcsCardInquiryV1FvEmeaTooManyRequests {
	return &AcsCardInquiryV1FvEmeaTooManyRequests{}
}

/*AcsCardInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AcsCardInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcsCardInquiryV1FvEmeaStatus452 creates a AcsCardInquiryV1FvEmeaStatus452 with default headers values
func NewAcsCardInquiryV1FvEmeaStatus452() *AcsCardInquiryV1FvEmeaStatus452 {
	return &AcsCardInquiryV1FvEmeaStatus452{}
}

/*AcsCardInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AcsCardInquiryV1FvEmeaStatus452 struct {
}

func (o *AcsCardInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaStatus452 ", 452)
}

func (o *AcsCardInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAcsCardInquiryV1FvEmeaStatus453 creates a AcsCardInquiryV1FvEmeaStatus453 with default headers values
func NewAcsCardInquiryV1FvEmeaStatus453() *AcsCardInquiryV1FvEmeaStatus453 {
	return &AcsCardInquiryV1FvEmeaStatus453{}
}

/*AcsCardInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AcsCardInquiryV1FvEmeaStatus453 struct {
}

func (o *AcsCardInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaStatus453 ", 453)
}

func (o *AcsCardInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAcsCardInquiryV1FvEmeaStatus455 creates a AcsCardInquiryV1FvEmeaStatus455 with default headers values
func NewAcsCardInquiryV1FvEmeaStatus455() *AcsCardInquiryV1FvEmeaStatus455 {
	return &AcsCardInquiryV1FvEmeaStatus455{}
}

/*AcsCardInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AcsCardInquiryV1FvEmeaStatus455 struct {
}

func (o *AcsCardInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaStatus455 ", 455)
}

func (o *AcsCardInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAcsCardInquiryV1FvEmeaStatus465 creates a AcsCardInquiryV1FvEmeaStatus465 with default headers values
func NewAcsCardInquiryV1FvEmeaStatus465() *AcsCardInquiryV1FvEmeaStatus465 {
	return &AcsCardInquiryV1FvEmeaStatus465{}
}

/*AcsCardInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SC302S : FILE TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SC303S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SC304S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SC305S : SERVICE INPUT TO CARD INQUIRY SERVICE IS OF INCORRECT LENGTH<BR/>VPL5SC306S : CARD NUMBER MUST BE NUMERIC<BR/>VPL5SC307S : ORG IS NOT PRESENT IN SYSTEM<BR/>VPL5SC308S : ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SC309S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SC310S : FILE TABLE USED BY THIS SERVICE IS MISSING<BR/>VPL5SC311S : ORG RECORD IS NOT FOUND OR ADD PENDING IN SYSTEM<BR/>VPL5SC312S : EMBOSSER RECORD IS NOT FOUND OR ADD PENDING IN SYSTEM<BR/>VPL5SC313S : ACCOUNT BASE RECORD IS NOT FOUND OR ADD PENDING IN SYSTEM<BR/>VPL5SC314S : CUSTOMER NUMBER IS NOT FOUND OR ADD PENDING IN SYSTEM<BR/>VPL5SC315S : MATCH NOT FOUND IN ACCOUNTorCARD CROSS REFERENCE FILE<BR/>VPL5SC316S : ACCTorCARD NOT FOUND IN AMED AMBS AND FMIC
*/
type AcsCardInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *AcsCardInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/acsCardInquiry][%d] acsCardInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *AcsCardInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AcsCardInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}