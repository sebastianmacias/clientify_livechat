// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicHTTPError public Http error
//
// swagger:model publicHttpError
type PublicHTTPError struct {

	// HTTP status code returned for the error
	// Example: 403
	// Required: true
	// Maximum: 599
	// Minimum: 100
	Code *int64 `json:"status"`

	// More detailed, human-readable, optional explanation of the error
	// Example: User is lacking permission to access this resource
	Detail string `json:"detail,omitempty"`

	// Short, human-readable description of the error
	// Example: Forbidden
	// Required: true
	Title *string `json:"title"`

	// Type of error returned, should be used for client-side error handling
	// Example: generic
	// Required: true
	Type *string `json:"type"`

	// Error identifier
	// Example: MOD_117
	XRef string `json:"x_ref,omitempty"`
}

// Validate validates this public Http error
func (m *PublicHTTPError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicHTTPError) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MinimumInt("status", "body", *m.Code, 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", *m.Code, 599, false); err != nil {
		return err
	}

	return nil
}

func (m *PublicHTTPError) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PublicHTTPError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this public Http error based on context it is used
func (m *PublicHTTPError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicHTTPError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicHTTPError) UnmarshalBinary(b []byte) error {
	var res PublicHTTPError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
