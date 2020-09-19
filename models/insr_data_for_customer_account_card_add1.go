// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InsrDataForCustomerAccountCardAdd1 insr data for customer account card add1
//
// swagger:model InsrDataForCustomerAccountCardAdd1
type InsrDataForCustomerAccountCardAdd1 struct {

	//  Max length = 1, Insurance Add Result Flag. Values are: P - Passed, F - Fail
	InsuranceAddResult string `json:"insuranceAddResult,omitempty"`

	//  Max length = 2, INSURANCE PROD CODE
	InsuranceProdCode string `json:"insuranceProdCode,omitempty"`
}

// Validate validates this insr data for customer account card add1
func (m *InsrDataForCustomerAccountCardAdd1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InsrDataForCustomerAccountCardAdd1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InsrDataForCustomerAccountCardAdd1) UnmarshalBinary(b []byte) error {
	var res InsrDataForCustomerAccountCardAdd1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}