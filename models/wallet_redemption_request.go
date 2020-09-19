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

// WalletRedemptionRequest wallet redemption request
//
// swagger:model WalletRedemptionRequest
type WalletRedemptionRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account,omitempty"`

	//  Max length = 4, Action Code: Four character code that identifies  the action to be taken.
	// Max Length: 4
	// Min Length: 0
	Action *string `json:"action,omitempty"`

	//  Max length = 11, Case number: number assigned to action item.
	// Max Length: 11
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CaseNbr *string `json:"caseNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Redemption Confirm flag/Indicator. Valid values are: Y - Redemption request confirmed N - Redemption request not confirmed Space - Redemption inquiry
	// Max Length: 1
	// Min Length: 0
	ConfirmInd *string `json:"confirmInd,omitempty"`

	//  Max length = 1, Flag  to indicate whether the redemption amount will be credited to direct credit account defined on the account base segment. Action should allow direct credit and the corresponding transaction code to mapped to the action code.
	// Max Length: 1
	// Min Length: 0
	DirCredit *string `json:"dirCredit,omitempty"`

	// Format: YYYYMMDD. Effective date.
	EffDate string `json:"effDate,omitempty"`

	//  Max length = 9, Fee amount in base currency. Should be numeric.
	// Pattern: ^(-)?[0-9]{1,9}$
	FeeAmtBase string `json:"feeAmtBase,omitempty"`

	//  Max length = 9, FEE AMT WALT
	// Pattern: ^(-)?[0-9]{1,9}$
	FeeAmtWalt string `json:"feeAmtWalt,omitempty"`

	//  Max length = 60, Free text that is stored against  the account, this can be used to identify to operators why the Action code has been used.
	// Max Length: 60
	// Min Length: 0
	LineData1 *string `json:"lineData1,omitempty"`

	//  Max length = 60, Memo Line 2
	// Max Length: 60
	// Min Length: 0
	LineData2 *string `json:"lineData2,omitempty"`

	//  Max length = 60, Memo Line 3
	// Max Length: 60
	// Min Length: 0
	LineData3 *string `json:"lineData3,omitempty"`

	//  Max length = 60, Memo Line 4
	// Max Length: 60
	// Min Length: 0
	LineData4 *string `json:"lineData4,omitempty"`

	//  Max length = 60, Memo Line 5
	// Max Length: 60
	// Min Length: 0
	LineData5 *string `json:"lineData5,omitempty"`

	//  Max length = 9, Net redemption Amount in base currency. Should be numeric.
	// Pattern: ^(-)?[0-9]{1,9}$
	NetRdmAmt string `json:"netRdmAmt,omitempty"`

	//  Max length = 9, NET RDM AMT BSE
	// Pattern: ^(-)?[0-9]{1,9}$
	NetRdmAmtBse string `json:"netRdmAmtBse,omitempty"`

	//  Max length = 3, Organization Number: Three digit Identification number of the organization.  Valid values are 001 - 998. Organization number must be on file.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Org *string `json:"org,omitempty"`

	//  Max length = 8, RATE
	// Max Length: 8
	// Min Length: 0
	Rate *string `json:"rate,omitempty"`

	//  Max length = 9, Redemption Amount. Should be numeric and greater than zeroes.
	// Pattern: ^(-)?[0-9]{1,9}$
	RdmAmount string `json:"rdmAmount,omitempty"`

	//  Max length = 9, Redemption Amount. Should be numeric. It is just checked for numeric value in service on redemption confirmation.
	// Pattern: ^(-)?[0-9]{1,9}$
	RdmAmt string `json:"rdmAmt,omitempty"`

	//  Max length = 9, RDM AMT BSE
	// Pattern: ^(-)?[0-9]{1,9}$
	RdmAmtBse string `json:"rdmAmtBse,omitempty"`

	//  Max length = 1, Redemption Indicator. Valid values are: F - Full Redemption P - Partial Redemption
	// Max Length: 1
	// Min Length: 0
	RdmInd *string `json:"rdmInd,omitempty"`

	//  Max length = 1, Redemption Level. Valid values are: W - Wallet level Redemption A - Account level Redemption
	// Max Length: 1
	// Min Length: 0
	RdmLevel *string `json:"rdmLevel,omitempty"`

	//  Max length = 3, Redeem Wallet: Field which contains the wallet/wallet currency from which the amount should be redeemed. Should be numeric.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	RdmWalt *string `json:"rdmWalt,omitempty"`

	//  Max length = 3, Account Service Management(ASM) user name. Must be a valid ASM Rep.
	// Max Length: 3
	// Min Length: 0
	Rep *string `json:"rep,omitempty"`

	//  Max length = 1, Input source. Values are: 0 = ASM online 1 = CMS load/reload 2 = External source
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SourceFlag *string `json:"sourceFlag,omitempty"`

	//  Max length = 9, Store: Store number associated with the transaction.
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Store *string `json:"store,omitempty"`
}

// Validate validates this wallet redemption request
func (m *WalletRedemptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaseNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfirmInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirCredit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeAmtBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeAmtWalt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineData1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineData2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineData3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineData4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineData5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetRdmAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetRdmAmtBse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmAmtBse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdmWalt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WalletRedemptionRequest) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if err := validate.MinLength("account", "body", string(*m.Account), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("account", "body", string(*m.Account), 19); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.MinLength("action", "body", string(*m.Action), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("action", "body", string(*m.Action), 4); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateCaseNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CaseNbr) { // not required
		return nil
	}

	if err := validate.MinLength("caseNbr", "body", string(*m.CaseNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("caseNbr", "body", string(*m.CaseNbr), 11); err != nil {
		return err
	}

	if err := validate.Pattern("caseNbr", "body", string(*m.CaseNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *WalletRedemptionRequest) validateConfirmInd(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfirmInd) { // not required
		return nil
	}

	if err := validate.MinLength("confirmInd", "body", string(*m.ConfirmInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("confirmInd", "body", string(*m.ConfirmInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateDirCredit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirCredit) { // not required
		return nil
	}

	if err := validate.MinLength("dirCredit", "body", string(*m.DirCredit), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("dirCredit", "body", string(*m.DirCredit), 1); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateFeeAmtBase(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeAmtBase) { // not required
		return nil
	}

	if err := validate.Pattern("feeAmtBase", "body", string(m.FeeAmtBase), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateFeeAmtWalt(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeAmtWalt) { // not required
		return nil
	}

	if err := validate.Pattern("feeAmtWalt", "body", string(m.FeeAmtWalt), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateLineData1(formats strfmt.Registry) error {

	if swag.IsZero(m.LineData1) { // not required
		return nil
	}

	if err := validate.MinLength("lineData1", "body", string(*m.LineData1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lineData1", "body", string(*m.LineData1), 60); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateLineData2(formats strfmt.Registry) error {

	if swag.IsZero(m.LineData2) { // not required
		return nil
	}

	if err := validate.MinLength("lineData2", "body", string(*m.LineData2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lineData2", "body", string(*m.LineData2), 60); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateLineData3(formats strfmt.Registry) error {

	if swag.IsZero(m.LineData3) { // not required
		return nil
	}

	if err := validate.MinLength("lineData3", "body", string(*m.LineData3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lineData3", "body", string(*m.LineData3), 60); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateLineData4(formats strfmt.Registry) error {

	if swag.IsZero(m.LineData4) { // not required
		return nil
	}

	if err := validate.MinLength("lineData4", "body", string(*m.LineData4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lineData4", "body", string(*m.LineData4), 60); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateLineData5(formats strfmt.Registry) error {

	if swag.IsZero(m.LineData5) { // not required
		return nil
	}

	if err := validate.MinLength("lineData5", "body", string(*m.LineData5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lineData5", "body", string(*m.LineData5), 60); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateNetRdmAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.NetRdmAmt) { // not required
		return nil
	}

	if err := validate.Pattern("netRdmAmt", "body", string(m.NetRdmAmt), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateNetRdmAmtBse(formats strfmt.Registry) error {

	if swag.IsZero(m.NetRdmAmtBse) { // not required
		return nil
	}

	if err := validate.Pattern("netRdmAmtBse", "body", string(m.NetRdmAmtBse), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.Org) { // not required
		return nil
	}

	if err := validate.MinLength("org", "body", string(*m.Org), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("org", "body", string(*m.Org), 3); err != nil {
		return err
	}

	if err := validate.Pattern("org", "body", string(*m.Org), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := validate.MinLength("rate", "body", string(*m.Rate), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rate", "body", string(*m.Rate), 8); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmAmount) { // not required
		return nil
	}

	if err := validate.Pattern("rdmAmount", "body", string(m.RdmAmount), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmAmt) { // not required
		return nil
	}

	if err := validate.Pattern("rdmAmt", "body", string(m.RdmAmt), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmAmtBse(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmAmtBse) { // not required
		return nil
	}

	if err := validate.Pattern("rdmAmtBse", "body", string(m.RdmAmtBse), `^(-)?[0-9]{1,9}$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmInd(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmInd) { // not required
		return nil
	}

	if err := validate.MinLength("rdmInd", "body", string(*m.RdmInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rdmInd", "body", string(*m.RdmInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmLevel) { // not required
		return nil
	}

	if err := validate.MinLength("rdmLevel", "body", string(*m.RdmLevel), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rdmLevel", "body", string(*m.RdmLevel), 1); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRdmWalt(formats strfmt.Registry) error {

	if swag.IsZero(m.RdmWalt) { // not required
		return nil
	}

	if err := validate.MinLength("rdmWalt", "body", string(*m.RdmWalt), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rdmWalt", "body", string(*m.RdmWalt), 3); err != nil {
		return err
	}

	if err := validate.Pattern("rdmWalt", "body", string(*m.RdmWalt), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateRep(formats strfmt.Registry) error {

	if swag.IsZero(m.Rep) { // not required
		return nil
	}

	if err := validate.MinLength("rep", "body", string(*m.Rep), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rep", "body", string(*m.Rep), 3); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateSourceFlag(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceFlag) { // not required
		return nil
	}

	if err := validate.MinLength("sourceFlag", "body", string(*m.SourceFlag), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("sourceFlag", "body", string(*m.SourceFlag), 1); err != nil {
		return err
	}

	if err := validate.Pattern("sourceFlag", "body", string(*m.SourceFlag), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *WalletRedemptionRequest) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if err := validate.MinLength("store", "body", string(*m.Store), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("store", "body", string(*m.Store), 9); err != nil {
		return err
	}

	if err := validate.Pattern("store", "body", string(*m.Store), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WalletRedemptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WalletRedemptionRequest) UnmarshalBinary(b []byte) error {
	var res WalletRedemptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}