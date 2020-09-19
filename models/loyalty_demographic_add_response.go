// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoyaltyDemographicAddResponse loyalty demographic add response
//
// swagger:model LoyaltyDemographicAddResponse
type LoyaltyDemographicAddResponse struct {

	//  Max length = 19, ACCT
	Acct string `json:"acct,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`
}

// Validate validates this loyalty demographic add response
func (m *LoyaltyDemographicAddResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyDemographicAddResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *LoyaltyDemographicAddResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyDemographicAddResponse) UnmarshalBinary(b []byte) error {
	var res LoyaltyDemographicAddResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}