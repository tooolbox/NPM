// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DelinquencyAdjustmentUpdateResponse delinquency adjustment update response
//
// swagger:model DelinquencyAdjustmentUpdateResponse
type DelinquencyAdjustmentUpdateResponse struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	AcctNbr string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Amount of payment currently due (not including delinquent amounts).
	CurrDue string `json:"currDue,omitempty"`

	//  Max length = 17, Current balance: Total current outstanding balance of the account
	CurrentBalance string `json:"currentBalance,omitempty"`

	//  Max length = 17, Amount of payment that is 120 to 149 days past due.
	DaysDue120 string `json:"daysDue120,omitempty"`

	//  Max length = 17, Amount of payment that is 150 to 179 days past due.
	DaysDue150 string `json:"daysDue150,omitempty"`

	//  Max length = 17, Amount of payment that is 180 to 209 days past due.
	DaysDue180 string `json:"daysDue180,omitempty"`

	//  Max length = 17, Amount of payment that is 210+ days past due.
	DaysDue210 string `json:"daysDue210,omitempty"`

	//  Max length = 17, DAYS DUE 30
	DaysDue30 string `json:"daysDue30,omitempty"`

	//  Max length = 17, DAYS DUE 60
	DaysDue60 string `json:"daysDue60,omitempty"`

	//  Max length = 17, DAYS DUE 90
	DaysDue90 string `json:"daysDue90,omitempty"`

	//  Max length = 17, Difference in the total amount due on the account and on all credit plans.
	DiffInDue string `json:"diffInDue,omitempty"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	//  Max length = 2, Number of Items returned.
	NbrReturnedItems string `json:"nbrReturnedItems,omitempty"`

	//  Max length = 17, Current open-to-buy amount of the account.
	OpenToBuy string `json:"openToBuy,omitempty"`

	//  Max length = 17, Amount of payment past due.
	PastDue string `json:"pastDue,omitempty"`

	// Group. Plan Data. Number of returned items; occurs 99 times.
	PlanData []*PlanDataForDelinquencyAdjustmentUpdate1 `json:"planData"`

	//  Max length = 1, Reage code. Values are: X = Manual Reage 1-8 = Automatic Reage
	ReageCd string `json:"reageCd,omitempty"`

	//  Max length = 3, Reage Period: Period in months during which the account can be reaged a limited number of times.
	ReagePeriod string `json:"reagePeriod,omitempty"`

	//  Max length = 1, Reage Period Limit: Number of times the account can be reaged within a specified period.
	ReagePeriodLimit string `json:"reagePeriodLimit,omitempty"`

	//  Max length = 3, Reage Times: Number of times the account has been reaged.
	ReageTimes string `json:"reageTimes,omitempty"`

	//  Max length = 17, Sum of all the current and past due amounts on the account.
	TotDueBase string `json:"totDueBase,omitempty"`

	//  Max length = 17, Sum of all total amount due for all credit plan segments on the account.
	TotDuePlans string `json:"totDuePlans,omitempty"`
}

// Validate validates this delinquency adjustment update response
func (m *DelinquencyAdjustmentUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DelinquencyAdjustmentUpdateResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *DelinquencyAdjustmentUpdateResponse) validatePlanData(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanData) { // not required
		return nil
	}

	for i := 0; i < len(m.PlanData); i++ {
		if swag.IsZero(m.PlanData[i]) { // not required
			continue
		}

		if m.PlanData[i] != nil {
			if err := m.PlanData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("planData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DelinquencyAdjustmentUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DelinquencyAdjustmentUpdateResponse) UnmarshalBinary(b []byte) error {
	var res DelinquencyAdjustmentUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}