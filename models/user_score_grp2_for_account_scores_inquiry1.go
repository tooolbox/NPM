// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserScoreGrp2ForAccountScoresInquiry1 user score grp2 for account scores inquiry1
//
// swagger:model UserScoreGrp2ForAccountScoresInquiry1
type UserScoreGrp2ForAccountScoresInquiry1 struct {

	//  Max length = 2, User Score 11 to 20. User score to be used as a decision key in TRIAD batch and online
	ScrGrp2Fld string `json:"scrGrp2Fld,omitempty"`
}

// Validate validates this user score grp2 for account scores inquiry1
func (m *UserScoreGrp2ForAccountScoresInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserScoreGrp2ForAccountScoresInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserScoreGrp2ForAccountScoresInquiry1) UnmarshalBinary(b []byte) error {
	var res UserScoreGrp2ForAccountScoresInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}