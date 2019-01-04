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

// Target target
// swagger:model Target
type Target struct {

	// Method
	// Enum: [GET POST PUT HEAD DELETE]
	Method *string `json:"method,omitempty"`

	// scheme
	// Enum: [http https]
	Scheme *string `json:"scheme,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this target
func (m *Target) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var targetTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","POST","PUT","HEAD","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		targetTypeMethodPropEnum = append(targetTypeMethodPropEnum, v)
	}
}

const (

	// TargetMethodGET captures enum value "GET"
	TargetMethodGET string = "GET"

	// TargetMethodPOST captures enum value "POST"
	TargetMethodPOST string = "POST"

	// TargetMethodPUT captures enum value "PUT"
	TargetMethodPUT string = "PUT"

	// TargetMethodHEAD captures enum value "HEAD"
	TargetMethodHEAD string = "HEAD"

	// TargetMethodDELETE captures enum value "DELETE"
	TargetMethodDELETE string = "DELETE"
)

// prop value enum
func (m *Target) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, targetTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Target) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", *m.Method); err != nil {
		return err
	}

	return nil
}

var targetTypeSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		targetTypeSchemePropEnum = append(targetTypeSchemePropEnum, v)
	}
}

const (

	// TargetSchemeHTTP captures enum value "http"
	TargetSchemeHTTP string = "http"

	// TargetSchemeHTTPS captures enum value "https"
	TargetSchemeHTTPS string = "https"
)

// prop value enum
func (m *Target) validateSchemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, targetTypeSchemePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Target) validateScheme(formats strfmt.Registry) error {

	if swag.IsZero(m.Scheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateSchemeEnum("scheme", "body", *m.Scheme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Target) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Target) UnmarshalBinary(b []byte) error {
	var res Target
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
