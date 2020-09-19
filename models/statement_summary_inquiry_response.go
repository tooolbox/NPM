// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatementSummaryInquiryResponse statement summary inquiry response
//
// swagger:model StatementSummaryInquiryResponse
type StatementSummaryInquiryResponse struct {

	//  Max length = 17, The value for available credit is calculated during the service request. The value is being calculated using the VisionPlus available credit calculation routine.
	AvailCredit string `json:"availCredit,omitempty"`

	//  Max length = 17, Cash Credit Limit at account level
	CashCrlim string `json:"cashCrlim,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Total credit limit established for this account
	CreditLimit string `json:"creditLimit,omitempty"`

	// Format: YYYYDDD. Date of the Payment. Date format is CCYY-MM-DD.
	DdPaymentDate string `json:"ddPaymentDate,omitempty"`

	//  Max length = 17, Total amount of interest estimated for this account. This amount is the sum of the estimated interest calculated for all applicable plans on the account.
	EstimatedInterest string `json:"estimatedInterest,omitempty"`

	//  Max length = 17, Accumulation of statement begin balance  for all plans associated with this account. Balance for this account at the beginning of this cycle.
	LastStmtBal string `json:"lastStmtBal,omitempty"`

	//  Max length = 17, Accumulation of payment due amounts for all plans associated with this account base segment record.
	PaymentDueAmt string `json:"paymentDueAmt,omitempty"`

	// Format: YYYYMMDD. Due date of the paymentbilled on this statement. Date format is CCYY-MM-DD.
	PaymentDueDate string `json:"paymentDueDate,omitempty"`

	//  Max length = 17, Ending balance of the account on this statement.
	StatementBal string `json:"statementBal,omitempty"`

	// Format: YYYYMMDD. Effective date of this statement. Date format is CCYY-MM-DD.
	StatementDate string `json:"statementDate,omitempty"`

	//  Max length = 1, The type of statement produced for the account. The values are: Blank = Normal statement C = Cycle statement H = Hold statement I = Interim statement R = Returned statement S = Sign-out statement O = Online only
	StmtFlag string `json:"stmtFlag,omitempty"`

	//  Max length = 3, Total number of debits taken from all plans associated with this account base segment record.
	TotNbrDb string `json:"totNbrDb,omitempty"`

	//  Max length = 17, Total amount of cycle-to-date credits. Accumulation of the AMPS-AMT-CR for all plans associated with this account.
	TotalCreditAmount string `json:"totalCreditAmount,omitempty"`

	//  Max length = 17, Total amount of cycle-to-date debits. Accumulation of the AMPS-AMT-DB for all plans associated with this account.
	TotalDebitAmount string `json:"totalDebitAmount,omitempty"`

	//  Max length = 17, Interest charged for all plans associated with this account. Total amount of interest charged for this cycle.
	TotalInterestCharged string `json:"totalInterestCharged,omitempty"`
}

// Validate validates this statement summary inquiry response
func (m *StatementSummaryInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatementSummaryInquiryResponse) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.Common) { // not required
		return nil
	}

	if m.Common != nil {
		if err := m.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatementSummaryInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatementSummaryInquiryResponse) UnmarshalBinary(b []byte) error {
	var res StatementSummaryInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}