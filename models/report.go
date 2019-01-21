// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Report report
// swagger:model Report
type Report struct {

	// bytes in
	BytesIn *ReportBytesIn `json:"bytes_in,omitempty"`

	// bytes out
	BytesOut *ReportBytesOut `json:"bytes_out,omitempty"`

	// duration
	Duration int64 `json:"duration,omitempty"`

	// earliest
	// Format: date-time
	Earliest strfmt.DateTime `json:"earliest,omitempty"`

	// end
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// errors
	Errors []string `json:"errors"`

	// latencies
	Latencies *ReportLatencies `json:"latencies,omitempty"`

	// latest
	// Format: date-time
	Latest strfmt.DateTime `json:"latest,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// requests
	Requests int64 `json:"requests,omitempty"`

	// success
	Success int64 `json:"success,omitempty"`

	// wait
	Wait int64 `json:"wait,omitempty"`
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytesOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEarliest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) validateBytesIn(formats strfmt.Registry) error {

	if swag.IsZero(m.BytesIn) { // not required
		return nil
	}

	if m.BytesIn != nil {
		if err := m.BytesIn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bytes_in")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateBytesOut(formats strfmt.Registry) error {

	if swag.IsZero(m.BytesOut) { // not required
		return nil
	}

	if m.BytesOut != nil {
		if err := m.BytesOut.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bytes_out")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateEarliest(formats strfmt.Registry) error {

	if swag.IsZero(m.Earliest) { // not required
		return nil
	}

	if err := validate.FormatOf("earliest", "body", "date-time", m.Earliest.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateLatencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Latencies) { // not required
		return nil
	}

	if m.Latencies != nil {
		if err := m.Latencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latencies")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateLatest(formats strfmt.Registry) error {

	if swag.IsZero(m.Latest) { // not required
		return nil
	}

	if err := validate.FormatOf("latest", "body", "date-time", m.Latest.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReportBytesIn report bytes in
// swagger:model ReportBytesIn
type ReportBytesIn struct {

	// mean
	Mean int64 `json:"mean,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this report bytes in
func (m *ReportBytesIn) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportBytesIn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportBytesIn) UnmarshalBinary(b []byte) error {
	var res ReportBytesIn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReportBytesOut report bytes out
// swagger:model ReportBytesOut
type ReportBytesOut struct {

	// mean
	Mean int64 `json:"mean,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this report bytes out
func (m *ReportBytesOut) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportBytesOut) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportBytesOut) UnmarshalBinary(b []byte) error {
	var res ReportBytesOut
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReportLatencies report latencies
// swagger:model ReportLatencies
type ReportLatencies struct {

	// 50th
	Nr50th uint64 `json:"50th,omitempty"`

	// 95th
	Nr95th uint64 `json:"95th,omitempty"`

	// 99th
	Nr99th uint64 `json:"99th,omitempty"`

	// max
	Max int64 `json:"max,omitempty"`

	// mean
	Mean uint64 `json:"mean,omitempty"`

	// total
	Total uint64 `json:"total,omitempty"`
}

// Validate validates this report latencies
func (m *ReportLatencies) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportLatencies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportLatencies) UnmarshalBinary(b []byte) error {
	var res ReportLatencies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
