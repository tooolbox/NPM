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

// CardAntiReferralUpdateRequest card anti referral update request
//
// swagger:model CardAntiReferralUpdateRequest
type CardAntiReferralUpdateRequest struct {

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNumber *string `json:"cardNumber"`

	//  Max length = 4, Card sequence number: record number assigned to the card (for card numbering schemes of 0, 1, and 3) and the sequence number assigned to the card (for card numbering schemes of 2). This number is part of the record key.
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Foreign use indicator: Value indicates whether the incoming action is applied to the local or foreign account. The values are: 0 = Local account (default) 1 = Foreign account SPACE = defaults to 0
	// Required: true
	// Max Length: 1
	// Min Length: 0
	ForeignUse *string `json:"foreignUse"`

	// Format: YYYYMMDD. Refer Anti-Refer Expiry Date: Indicates the date on which the refer/anti-refer flag is automatically removed from the card. Format: CCYYMMDD. This is a mandatory field if value of Referral/Anti-referral flag is '1' or '2'.
	// Required: true
	ReferAntiReferExpDt *string `json:"referAntiReferExpDt"`

	//  Max length = 1, Refer Anti-Refer Flag: Indicates whether the card is flagged to have all authorizations referred or not referred (anti-referred) to a fraud analyst for possible fraudulent activity.  Values are: '0' - No referral/anti-referral assigned  (default)  '1' - Referral  '2' - Anti-referral
	// Required: true
	// Max Length: 1
	// Min Length: 0
	ReferAntiReferFlag *string `json:"referAntiReferFlag"`
}

// Validate validates this card anti referral update request
func (m *CardAntiReferralUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferAntiReferExpDt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferAntiReferFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardAntiReferralUpdateRequest) validateCardNumber(formats strfmt.Registry) error {

	if err := validate.Required("cardNumber", "body", m.CardNumber); err != nil {
		return err
	}

	if err := validate.MinLength("cardNumber", "body", string(*m.CardNumber), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNumber", "body", string(*m.CardNumber), 19); err != nil {
		return err
	}

	return nil
}

func (m *CardAntiReferralUpdateRequest) validateCardSeq(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSeq) { // not required
		return nil
	}

	if err := validate.MinLength("cardSeq", "body", string(*m.CardSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSeq", "body", string(*m.CardSeq), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSeq", "body", string(*m.CardSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardAntiReferralUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *CardAntiReferralUpdateRequest) validateForeignUse(formats strfmt.Registry) error {

	if err := validate.Required("foreignUse", "body", m.ForeignUse); err != nil {
		return err
	}

	if err := validate.MinLength("foreignUse", "body", string(*m.ForeignUse), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("foreignUse", "body", string(*m.ForeignUse), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardAntiReferralUpdateRequest) validateReferAntiReferExpDt(formats strfmt.Registry) error {

	if err := validate.Required("referAntiReferExpDt", "body", m.ReferAntiReferExpDt); err != nil {
		return err
	}

	return nil
}

func (m *CardAntiReferralUpdateRequest) validateReferAntiReferFlag(formats strfmt.Registry) error {

	if err := validate.Required("referAntiReferFlag", "body", m.ReferAntiReferFlag); err != nil {
		return err
	}

	if err := validate.MinLength("referAntiReferFlag", "body", string(*m.ReferAntiReferFlag), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("referAntiReferFlag", "body", string(*m.ReferAntiReferFlag), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardAntiReferralUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardAntiReferralUpdateRequest) UnmarshalBinary(b []byte) error {
	var res CardAntiReferralUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}