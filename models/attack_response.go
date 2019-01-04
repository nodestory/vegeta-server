// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AttackResponse attack response
// swagger:model AttackResponse
type AttackResponse struct {

	// The attack ID used to fetch status and reports.
	ID string `json:"id,omitempty"`

	// The attack status
	// Enum: [scheduled running canceled completed failed]
	Status string `json:"status,omitempty"`
}

// Validate validates this attack response
func (m *AttackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attackResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scheduled","running","canceled","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attackResponseTypeStatusPropEnum = append(attackResponseTypeStatusPropEnum, v)
	}
}

const (

	// AttackResponseStatusScheduled captures enum value "scheduled"
	AttackResponseStatusScheduled string = "scheduled"

	// AttackResponseStatusRunning captures enum value "running"
	AttackResponseStatusRunning string = "running"

	// AttackResponseStatusCanceled captures enum value "canceled"
	AttackResponseStatusCanceled string = "canceled"

	// AttackResponseStatusCompleted captures enum value "completed"
	AttackResponseStatusCompleted string = "completed"

	// AttackResponseStatusFailed captures enum value "failed"
	AttackResponseStatusFailed string = "failed"
)

// prop value enum
func (m *AttackResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, attackResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AttackResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttackResponse) UnmarshalBinary(b []byte) error {
	var res AttackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
