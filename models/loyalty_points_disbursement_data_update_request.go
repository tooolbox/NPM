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

// LoyaltyPointsDisbursementDataUpdateRequest loyalty points disbursement data update request
//
// swagger:model LoyaltyPointsDisbursementDataUpdateRequest
type LoyaltyPointsDisbursementDataUpdateRequest struct {

	//  Max length = 19, ACCT NBR
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	//  Max length = 25, AUTO DISB ACCT
	// Max Length: 25
	// Min Length: 0
	AutoDisbAcct *string `json:"autoDisbAcct,omitempty"`

	//  Max length = 43, AUTO DISB DBA NAME
	// Max Length: 43
	// Min Length: 0
	AutoDisbDbaName *string `json:"autoDisbDbaName,omitempty"`

	//  Max length = 2, AUTO DISB DOM
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AutoDisbDom *string `json:"autoDisbDom,omitempty"`

	//  Max length = 1, AUTO DISB FLG
	// Max Length: 1
	// Min Length: 0
	AutoDisbFlg *string `json:"autoDisbFlg,omitempty"`

	//  Max length = 2, AUTO DISB FREQ
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AutoDisbFreq *string `json:"autoDisbFreq,omitempty"`

	//  Max length = 1, AUTO DISB IND
	// Max Length: 1
	// Min Length: 0
	AutoDisbInd *string `json:"autoDisbInd,omitempty"`

	//  Max length = 17, AUTO DISB THRHLD
	// Max Length: 17
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AutoDisbThrhld *string `json:"autoDisbThrhld,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, DISB PKT PTS
	// Max Length: 17
	// Min Length: 0
	// Pattern: ^[0-9]*$
	DisbPktPts *string `json:"disbPktPts,omitempty"`

	//  Max length = 6, RWD CAT
	// Max Length: 6
	// Min Length: 0
	RwdCat *string `json:"rwdCat,omitempty"`

	//  Max length = 5, SCHM
	// Max Length: 5
	// Min Length: 0
	Schm *string `json:"schm,omitempty"`
}

// Validate validates this loyalty points disbursement data update request
func (m *LoyaltyPointsDisbursementDataUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbAcct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbDbaName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbDom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbFlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbFreq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoDisbThrhld(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisbPktPts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRwdCat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbAcct(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbAcct) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbAcct", "body", string(*m.AutoDisbAcct), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbAcct", "body", string(*m.AutoDisbAcct), 25); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbDbaName(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbDbaName) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbDbaName", "body", string(*m.AutoDisbDbaName), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbDbaName", "body", string(*m.AutoDisbDbaName), 43); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbDom(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbDom) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbDom", "body", string(*m.AutoDisbDom), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbDom", "body", string(*m.AutoDisbDom), 2); err != nil {
		return err
	}

	if err := validate.Pattern("autoDisbDom", "body", string(*m.AutoDisbDom), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbFlg(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbFlg) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbFlg", "body", string(*m.AutoDisbFlg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbFlg", "body", string(*m.AutoDisbFlg), 1); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbFreq(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbFreq) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbFreq", "body", string(*m.AutoDisbFreq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbFreq", "body", string(*m.AutoDisbFreq), 2); err != nil {
		return err
	}

	if err := validate.Pattern("autoDisbFreq", "body", string(*m.AutoDisbFreq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbInd(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbInd) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbInd", "body", string(*m.AutoDisbInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbInd", "body", string(*m.AutoDisbInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateAutoDisbThrhld(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDisbThrhld) { // not required
		return nil
	}

	if err := validate.MinLength("autoDisbThrhld", "body", string(*m.AutoDisbThrhld), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("autoDisbThrhld", "body", string(*m.AutoDisbThrhld), 17); err != nil {
		return err
	}

	if err := validate.Pattern("autoDisbThrhld", "body", string(*m.AutoDisbThrhld), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateDisbPktPts(formats strfmt.Registry) error {

	if swag.IsZero(m.DisbPktPts) { // not required
		return nil
	}

	if err := validate.MinLength("disbPktPts", "body", string(*m.DisbPktPts), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("disbPktPts", "body", string(*m.DisbPktPts), 17); err != nil {
		return err
	}

	if err := validate.Pattern("disbPktPts", "body", string(*m.DisbPktPts), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateRwdCat(formats strfmt.Registry) error {

	if swag.IsZero(m.RwdCat) { // not required
		return nil
	}

	if err := validate.MinLength("rwdCat", "body", string(*m.RwdCat), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rwdCat", "body", string(*m.RwdCat), 6); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyPointsDisbursementDataUpdateRequest) validateSchm(formats strfmt.Registry) error {

	if swag.IsZero(m.Schm) { // not required
		return nil
	}

	if err := validate.MinLength("schm", "body", string(*m.Schm), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("schm", "body", string(*m.Schm), 5); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyPointsDisbursementDataUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyPointsDisbursementDataUpdateRequest) UnmarshalBinary(b []byte) error {
	var res LoyaltyPointsDisbursementDataUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}