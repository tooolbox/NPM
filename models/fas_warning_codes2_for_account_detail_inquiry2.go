// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FasWarningCodes2ForAccountDetailInquiry2 fas warning codes2 for account detail inquiry2
//
// swagger:model FasWarningCodes2ForAccountDetailInquiry2
type FasWarningCodes2ForAccountDetailInquiry2 struct {

	//  Max length = 1, Warning Code: User defined field used by FirstVision authorisations to perform qualifying credit tests.  Values are:  '0' - Normal authorisations (Default)  '1' - Decline  '2' - Decline and pick up  '3' - Decline with fraud code  '4' - Decline with referral  '8' - Decline with a user exit  '9' - VIP account; always approve.
	WarningCode string `json:"warningCode,omitempty"`
}

// Validate validates this fas warning codes2 for account detail inquiry2
func (m *FasWarningCodes2ForAccountDetailInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FasWarningCodes2ForAccountDetailInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FasWarningCodes2ForAccountDetailInquiry2) UnmarshalBinary(b []byte) error {
	var res FasWarningCodes2ForAccountDetailInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}