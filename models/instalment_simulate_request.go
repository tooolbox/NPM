// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstalmentSimulateRequest instalment simulate request
//
// swagger:model InstalmentSimulateRequest
type InstalmentSimulateRequest struct {

	//  Max length = 19, ACCT NBR
	// Required: true
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 13, FIXED PMT AMT
	// Required: true
	// Pattern: ^(-)?[0-9]{1,13}$
	FixedPmtAmt *string `json:"fixedPmtAmt"`

	//  Max length = 1, INST COMP MODE
	// Required: true
	// Max Length: 1
	// Min Length: 0
	InstCompMode *string `json:"instCompMode"`

	//  Max length = 5, INST PLAN NBR
	// Required: true
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	InstPlanNbr *string `json:"instPlanNbr"`

	//  Max length = 13, PRIN AMT
	// Required: true
	// Pattern: ^(-)?[0-9]{1,13}$
	PrinAmt *string `json:"prinAmt"`

	//  Max length = 2, REQ TERM
	// Required: true
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ReqTerm *string `json:"reqTerm"`
}

// Validate validates this instalment simulate request
func (m *InstalmentSimulateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedPmtAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstCompMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstPlanNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrinAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstalmentSimulateRequest) validateAcctNbr(formats strfmt.Registry) error {

	if err := validate.Required("acctNbr", "body", m.AcctNbr); err != nil {
		return err
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *InstalmentSimulateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *InstalmentSimulateRequest) validateFixedPmtAmt(formats strfmt.Registry) error {

	if err := validate.Required("fixedPmtAmt", "body", m.FixedPmtAmt); err != nil {
		return err
	}

	if err := validate.Pattern("fixedPmtAmt", "body", string(*m.FixedPmtAmt), `^(-)?[0-9]{1,13}$`); err != nil {
		return err
	}

	return nil
}

func (m *InstalmentSimulateRequest) validateInstCompMode(formats strfmt.Registry) error {

	if err := validate.Required("instCompMode", "body", m.InstCompMode); err != nil {
		return err
	}

	if err := validate.MinLength("instCompMode", "body", string(*m.InstCompMode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("instCompMode", "body", string(*m.InstCompMode), 1); err != nil {
		return err
	}

	return nil
}

func (m *InstalmentSimulateRequest) validateInstPlanNbr(formats strfmt.Registry) error {

	if err := validate.Required("instPlanNbr", "body", m.InstPlanNbr); err != nil {
		return err
	}

	if err := validate.MinLength("instPlanNbr", "body", string(*m.InstPlanNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("instPlanNbr", "body", string(*m.InstPlanNbr), 5); err != nil {
		return err
	}

	if err := validate.Pattern("instPlanNbr", "body", string(*m.InstPlanNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *InstalmentSimulateRequest) validatePrinAmt(formats strfmt.Registry) error {

	if err := validate.Required("prinAmt", "body", m.PrinAmt); err != nil {
		return err
	}

	if err := validate.Pattern("prinAmt", "body", string(*m.PrinAmt), `^(-)?[0-9]{1,13}$`); err != nil {
		return err
	}

	return nil
}

func (m *InstalmentSimulateRequest) validateReqTerm(formats strfmt.Registry) error {

	if err := validate.Required("reqTerm", "body", m.ReqTerm); err != nil {
		return err
	}

	if err := validate.MinLength("reqTerm", "body", string(*m.ReqTerm), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("reqTerm", "body", string(*m.ReqTerm), 2); err != nil {
		return err
	}

	if err := validate.Pattern("reqTerm", "body", string(*m.ReqTerm), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstalmentSimulateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstalmentSimulateRequest) UnmarshalBinary(b []byte) error {
	var res InstalmentSimulateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}