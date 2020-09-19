// Code generated by go-swagger; DO NOT EDIT.

package debit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// DebitCardAvailableBalanceUpdateV1FvEmeaReader is a Reader for the DebitCardAvailableBalanceUpdateV1FvEmea structure.
type DebitCardAvailableBalanceUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DebitCardAvailableBalanceUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaOK creates a DebitCardAvailableBalanceUpdateV1FvEmeaOK with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaOK() *DebitCardAvailableBalanceUpdateV1FvEmeaOK {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaOK{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaOK struct {
	Payload *models.DebitCardAvailableBalanceUpdateResponse
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaOK) GetPayload() *models.DebitCardAvailableBalanceUpdateResponse {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebitCardAvailableBalanceUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaBadRequest creates a DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaBadRequest() *DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized creates a DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized() *DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaForbidden creates a DebitCardAvailableBalanceUpdateV1FvEmeaForbidden with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaForbidden() *DebitCardAvailableBalanceUpdateV1FvEmeaForbidden {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaForbidden{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaNotFound creates a DebitCardAvailableBalanceUpdateV1FvEmeaNotFound with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaNotFound() *DebitCardAvailableBalanceUpdateV1FvEmeaNotFound {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaNotFound{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests creates a DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests() *DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus452 creates a DebitCardAvailableBalanceUpdateV1FvEmeaStatus452 with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus452() *DebitCardAvailableBalanceUpdateV1FvEmeaStatus452 {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaStatus452{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaStatus452 struct {
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaStatus452 ", 452)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus453 creates a DebitCardAvailableBalanceUpdateV1FvEmeaStatus453 with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus453() *DebitCardAvailableBalanceUpdateV1FvEmeaStatus453 {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaStatus453{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaStatus453 struct {
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaStatus453 ", 453)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus455 creates a DebitCardAvailableBalanceUpdateV1FvEmeaStatus455 with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus455() *DebitCardAvailableBalanceUpdateV1FvEmeaStatus455 {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaStatus455{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaStatus455 struct {
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaStatus455 ", 455)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus465 creates a DebitCardAvailableBalanceUpdateV1FvEmeaStatus465 with default headers values
func NewDebitCardAvailableBalanceUpdateV1FvEmeaStatus465() *DebitCardAvailableBalanceUpdateV1FvEmeaStatus465 {
	return &DebitCardAvailableBalanceUpdateV1FvEmeaStatus465{}
}

/*DebitCardAvailableBalanceUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SAX01E : INVALID USER - FAILED SECURITY<BR/>VPL5SAX02E : FILE TABLE NOT FOUND<BR/>VPL5SAX03E : XFR REPL MUST BE YES NO OR BLANK<BR/>VPL5SAX04E : REISSUE-CRD-SCHM-3 MUST BE Y N OR SPACE<BR/>VPL5SAX05E : CARDLOGO-XFR NOT NUMERIC<BR/>VPL5SAX06E : CARDLOGO-ISSUE NOT NUMERIC<BR/>VPL5SAX07E : CARDLOGO-XFR LOGO REC NOT FOUND<BR/>VPL5SAX08E : CARDLOGO-ISSUE LOGO REC NOT FOUND<BR/>VPL5SAX09E : CARDLOGO-XFR NOT ON TO-LOGO<BR/>VPL5SAX10E : CARDLOGO-ISSUE NOT ON TO-LOGO<BR/>VPL5SAX11E : CARDLOGO-XFR NOT ON FROM-AMBS<BR/>VPL5SAX12E : MULTISCHEME ACCT-ONLY LOGO<BR/>VPL5SAX13E : CANNOT BE BOTH SCHEME-3 AND MULTI-SCHEME<BR/>VPL5SAX14E : TOO MANY AMED FOR ACCT - CANNOT HANDLE<BR/>VPL5SAX15E : CARDLOGO-ISSUE = CARDLOGO-XFR<BR/>VPL5SAX16E : XFR LOGO NOT A MULTISCHEME LOGO<BR/>VPL5SAX17E : ISSUE-LOGO - ACCT-ONLY MULTISCHEME LOGO<BR/>VPL5SAX18E : MULTISCHEME FLAG INVALID FOR LOGO<BR/>VPL5SAX19E : MULTISCHEME LOGO UNKNOWN VALUE<BR/>VPL5SAX20E : SERVICE FAILED<BR/>VPL5SAX21E : SECURITY ORG 1<BR/>VPL5SAX22E : SECURITY ORG 2<BR/>VPL5SAX23E : FUNCTION NOT VALID MUST BE C P R S OR T<BR/>VPL5SAX24E : XFR LOGO NOT MULTISCHEME LOGO<BR/>VPL5SAX25E : ORG COULD NOT BE CALCULATED FROM ACCT<BR/>VPL5SAX26E : ACCT ALREADY BEING XFRED BY SOMEONE ELSE<BR/>VPL5SAX27E : NO SERVICE NAME PROVIDED<BR/>VPL5SAX28E : FROM ACCT NOT FOUND<BR/>VPL5SAX29E : FROM ACCT NOT FOUND2<BR/>VPL5SAX30E : FUNCTION REVERSAL AND EOD SET<BR/>VPL5SAX31E : CANNOT FIND PRIMARY CARDSET EMBOSSER<BR/>VPL5SAX32E : NO CARD ADDED OR TRANSFERED FOR PRIMARY<BR/>VPL5SAX33E : SMART-XFR NOT ALLOWED FOR MAG-STRIPE ACCTS<BR/>VPL5SAX34E : REVERSAL NOT ALLOWED INT-STATUS<BR/>VPL5SAX35E : TO LOGO MULTISCHEME CARD ONLY LOGO<BR/>VPL5SAX36E : TO LOGO MULTISCHEME - UNKNOWN TYPE<BR/>VPL5SAX37E : NEW LOGO NOT NUMERIC<BR/>VPL5SAX38E : TO LOGO NOT FOUND<BR/>VPL5SAX39E : NEW CUSTOMER HAS WRONG ORG<BR/>VPL5SAX40E : NEW CUSTOMER HAS VALUE ZERO<BR/>VPL5SAX41E : NO XFR LOGO OR ISSUE LOGO ENTERED FOR MULTI-SCHEME ACCOUNT<BR/>VPL5SAX42E : VALID VALUES ARE 0 AND 1<BR/>VPL5SAX43E : VALID VALUES FOR REFUND ANNUAL FEE FLAG ARE 0 AND 1<BR/>VPL5SAX44E : VALID ONLY WHEN INS SELECT AT LOGO LEVEL IS 2<BR/>VPL5SAX45E : VALID VALUES FOR EXPIRE PROMOTIONS FLAG ARE 0 AND 1<BR/>VPL5SAX46E : DEFAULT CASH MUST BE NUMERIC<BR/>VPL5SAX47E : DEFAULT RETAIL MUST BE NUMERIC<BR/>VPL5SAX48E : DEFAULT PROMOTIONAL BT PLAN SHOULD BE NUMERIC<BR/>VPL5SAX49E : ACCOUNT CONTROL TBL OVERRIDE SHOULD BE NUMERIC<BR/>VPL5SAX50E : SERVIVE FEE TBL OVERRIDE SHOULD BE NUMERIC<BR/>VPL5SAX51E : INVALID DATE IN SVF FEE TBL EXPIRE FIELD<BR/>VPL5SAX52E : SVC FEE TBL EXPIRE DATE INPUT WITHOUT SVC FEE TBL<BR/>VPL5SAX53E : EMBLEM MUST BE NUMERIC<BR/>VPL5SAX54E : DEFAULT LOYALTY CODE SHOULD BE NUMERIC<BR/>VPL5SAX55E : SERVICE INPUT TO INS ADD SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SAX56E : VALID ONLY WHEN EXPIRE RETorBT PROM IS ACTIVE AT LOGO LEVE<BR/>VPL5SAX57E : VALID ONLY WHEN ACS IS ACTIVE AT BOTH SYS and ORG LEVEL<BR/>VPL5SAX47S : MOBILE PI NOT SUPPORTED ON ACCOUN
*/
type DebitCardAvailableBalanceUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/debitCardAvailableBalanceUpdate][%d] debitCardAvailableBalanceUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DebitCardAvailableBalanceUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}