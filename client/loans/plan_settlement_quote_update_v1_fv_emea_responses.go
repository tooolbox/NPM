// Code generated by go-swagger; DO NOT EDIT.

package loans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// PlanSettlementQuoteUpdateV1FvEmeaReader is a Reader for the PlanSettlementQuoteUpdateV1FvEmea structure.
type PlanSettlementQuoteUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlanSettlementQuoteUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewPlanSettlementQuoteUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPlanSettlementQuoteUpdateV1FvEmeaOK creates a PlanSettlementQuoteUpdateV1FvEmeaOK with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaOK() *PlanSettlementQuoteUpdateV1FvEmeaOK {
	return &PlanSettlementQuoteUpdateV1FvEmeaOK{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type PlanSettlementQuoteUpdateV1FvEmeaOK struct {
	Payload *models.PlanSettlementQuoteUpdateResponse
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaOK) GetPayload() *models.PlanSettlementQuoteUpdateResponse {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanSettlementQuoteUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaBadRequest creates a PlanSettlementQuoteUpdateV1FvEmeaBadRequest with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaBadRequest() *PlanSettlementQuoteUpdateV1FvEmeaBadRequest {
	return &PlanSettlementQuoteUpdateV1FvEmeaBadRequest{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type PlanSettlementQuoteUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaUnauthorized creates a PlanSettlementQuoteUpdateV1FvEmeaUnauthorized with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaUnauthorized() *PlanSettlementQuoteUpdateV1FvEmeaUnauthorized {
	return &PlanSettlementQuoteUpdateV1FvEmeaUnauthorized{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type PlanSettlementQuoteUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaForbidden creates a PlanSettlementQuoteUpdateV1FvEmeaForbidden with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaForbidden() *PlanSettlementQuoteUpdateV1FvEmeaForbidden {
	return &PlanSettlementQuoteUpdateV1FvEmeaForbidden{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type PlanSettlementQuoteUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaNotFound creates a PlanSettlementQuoteUpdateV1FvEmeaNotFound with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaNotFound() *PlanSettlementQuoteUpdateV1FvEmeaNotFound {
	return &PlanSettlementQuoteUpdateV1FvEmeaNotFound{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type PlanSettlementQuoteUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaTooManyRequests creates a PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaTooManyRequests() *PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests {
	return &PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaStatus452 creates a PlanSettlementQuoteUpdateV1FvEmeaStatus452 with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaStatus452() *PlanSettlementQuoteUpdateV1FvEmeaStatus452 {
	return &PlanSettlementQuoteUpdateV1FvEmeaStatus452{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type PlanSettlementQuoteUpdateV1FvEmeaStatus452 struct {
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaStatus452 ", 452)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaStatus453 creates a PlanSettlementQuoteUpdateV1FvEmeaStatus453 with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaStatus453() *PlanSettlementQuoteUpdateV1FvEmeaStatus453 {
	return &PlanSettlementQuoteUpdateV1FvEmeaStatus453{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type PlanSettlementQuoteUpdateV1FvEmeaStatus453 struct {
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaStatus453 ", 453)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaStatus455 creates a PlanSettlementQuoteUpdateV1FvEmeaStatus455 with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaStatus455() *PlanSettlementQuoteUpdateV1FvEmeaStatus455 {
	return &PlanSettlementQuoteUpdateV1FvEmeaStatus455{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type PlanSettlementQuoteUpdateV1FvEmeaStatus455 struct {
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaStatus455 ", 455)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPlanSettlementQuoteUpdateV1FvEmeaStatus465 creates a PlanSettlementQuoteUpdateV1FvEmeaStatus465 with default headers values
func NewPlanSettlementQuoteUpdateV1FvEmeaStatus465() *PlanSettlementQuoteUpdateV1FvEmeaStatus465 {
	return &PlanSettlementQuoteUpdateV1FvEmeaStatus465{}
}

/*PlanSettlementQuoteUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPE5SPQZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPE5SPQZ2S : SERVICE INPUT FOR SERVICE IS AN INCORRECT LENGTH<BR/>VPE5SPQZ3S : FILE TABLE IS NOT FOUND<BR/>VPE5SPQZ4S : COULD NOT DETERMINE ASSOCIATED ORG<BR/>VPE5SPQZ5S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPE5SPQZ6S : FILE TABLE IS NOT FOUND<BR/>VPE5SPQ02S000 : RESOURCE IS UNAVAILABLE - CONTACT YOUR ADMINISTRATOR<BR/>VPE5SPQ04E000 : AMORT TABLE NOT FOUND<BR/>VPE5SPQ05E000 : AMORT MISSING METHOD<BR/>VPE5SPQ07S000 : SYSTEM CONTROL AMCR NOT FOUND<BR/>VPE5SPQ08S000 : CREDIT PLAN RECORD IS NOT FOUND OR ADD PENDING<BR/>VPE5SPQ09S000 : RESOURCE IS UNAVAILABLE - CONTACT YOUR ADMINISTRATOR<BR/>VPE5SPQ10S000 : INVALID ORG<BR/>VPE5SPQ11S000 : INVALID ACCOUNT OR CARD NUMBER<BR/>VPE5SPQ12E000 : INVALID FOREIGN USE FLAG<BR/>VPE5SPQ12S000 : INVALID FOREIGN USE FLAG<BR/>VPE5SPQ13S000 : INVALID PLAN NUMBER<BR/>VPE5SPQ14S000 : INVALID PLAN SEQUENCE<BR/>VPE5SPQ15S000 : INVALID SETTLEMENT TYPE<BR/>VPE5SPQ16S000 : CUSTOMER SELECT QUOTE DATE NOT ALLOWED<BR/>VPE5SPQ17S000 : INPUT PAYOFF DATE IS INVALID<BR/>VPE5SPQ18S000 : ORG NOT FOUND OR IS IN ADD PENDING CONDITION<BR/>VPE5SPQ19S000 : FOREIGN ORG NOT FOUND OR IS IN ADD PENDING CONDITION<BR/>VPE5SPQ20S000 : ACCOUNT IS NOT ACTIVE OR NOT FOUND<BR/>VPE5SPQ21S000 : LOGO NOT FOUND OR IS IN ADD PENDING CONDITION<BR/>VPE5SPQ22S000 : INVALID ACCOUNT CONTROL TABLE LOOKUP<BR/>VPE5SPQ23S000 : INVALID INTEREST RECORD LOOKUP<BR/>VPE5SPQ24E000 : LOAN CONTAINS A DISPUTED AMOUNT<BR/>VPE5SPQ25I000 : ADJUSTED TO FLOOR RATE 1<BR/>VPE5SPQ26I000 : ADJUSTED TO FLOOR RATE 2<BR/>VPE5SPQ27I000 : ADJUSTED TO CEILING RATE 1<BR/>VPE5SPQ28I000 : ADJUSTED TO CEILING RATE 2<BR/>VPE5SPQ29I : ADJUSTED TO ALLOWABLE CAP DECREASE<BR/>VPE5SPQ30I : ADJUSTED TO ALLOWABLE CAP INCREASE<BR/>VPE5SPQ31I : MAKING A USURY ADJUSTMENT<BR/>VPE5SPQ32I : USED PCT LEVEL FROM ACCOUNT<BR/>VPE5SPQ33E : SYSTEM CONTROL PCT OVERRIDE LOOKUP INVALID<BR/>VPE5SPQ35E : ORG CONTROL PCT OVERRIDE LOOKUP INVALID<BR/>VPE5SPQ38E : LOGO CONTROL PCT OVERRIDE LOOKUP INVALID<BR/>VPE5SPQ41E : PCT OVERRIDE NOT FOUND<BR/>VPE5SPQ42E : SYSTEM CONTROL PCT LEVEL NOT FOUND<BR/>VPE5SPQ43E : ORG CONTROL PCT LEVEL NOT FOUND<BR/>VPE5SPQ44E : LOGO CONTROL PCT LEVEL NOT FOUND<BR/>VPE5SPQ45E : SYSTEM CONTROL INTEREST STATE NOT FOUND<BR/>VPE5SPQ46E : ORG CONTROL INTEREST STATE NOT FOUND<BR/>VPE5SPQ47E : LOGO CONTROL INTEREST STATE NOT FOUND<BR/>VPE5SPQ48E : SYSTEM CONTROL DEFAULT INTEREST STATE NOT FOUND<BR/>VPE5SPQ49E : ORG CONTROL DEFAULT INTEREST STATE NOT FOUND<BR/>VPE5SPQ50E : LOGO CONTROL DEFAULT INTEREST STATE NOT FOUND<BR/>VPE5SPQ54E : INTEREST RATE TABLE NOT FOUND<BR/>VPE5SPQ55E : USURY TABLE NOT FOUND<BR/>VPE5SPQ56E : DEFAULT USURY STATE OF ISSUE NOT FOUND ON TABLE<BR/>VPE5SPQ57E : DEFAULT USURY STATE OF RESIDENCE NOT FOUND ON TABLE<BR/>VPE5SPQ58E : RATES ARE ZERO ON THE STATE USURY TABLE<BR/>VPE5SPQ59E : NO STATE OF ISSUE ON THE ACCOUNT<BR/>VPE5SPQ60E : NO STATE OF RESIDENCE ON THE ACCOUNT<BR/>VPE5SPQ61E : DEFAULT LETTER VALID VALUES ARE 01 AND 2<BR/>VPE5SPQ62E : LETTER ID REQUIRED IN CASE OF NO DEFAULT LETTER<BR/>VPE5SPQ63E : SYSTEM CONTROL AMCR NOT FOUND<BR/>VPE5SPQ64I : MEMO NOTES NOT UPDATED<BR/>VPE5SPQ65I : LETTER SEND REQUEST FAILED<BR/>VPE5SPQ85E : DEFAULT LTTR  VALUE MUST BE 1 INCASE OF DEFAULT LTTR TO BE SEND<BR/>VPE5SPQ86S : PARTIAL EARLY SETTLEMENT NOT ALLOWED FOR THIS LOGO<BR/>VPE5SPQ87S : PARTIAL EARLY SETTLEMENT AMOUNT MUST BE NUMERIC<BR/>VPE5SPQ88S : SETTLEMENT TYPE MUST BE  W  FOR PARTIAL EARLY SETL QUOTE<BR/>VPE5SPQ89S : PARTIAL EARLY SETL QUOTE CANNOT BE CREATED FOR ACCT WITH CD > 1<BR/>VPE5SPQ90S : PARTIAL EARLY SETL AMT CANNOT BE GREATER THAN NEXT CYCLE PRIN BAL<BR/>VPE5SPQ91S : INVALID PES RC VALID VALUES ARE 0 AND 1<BR/>VPE5SPQ92S : PES AMT CANNOT BE ZERO WHEN PES RC IS
*/
type PlanSettlementQuoteUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/planSettlementQuoteUpdate][%d] planSettlementQuoteUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *PlanSettlementQuoteUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}