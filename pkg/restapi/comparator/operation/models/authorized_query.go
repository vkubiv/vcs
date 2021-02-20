// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthorizedQuery AuthorizedQuery is a query that has been pre-authorized by another Comparator.
// The AuthorizedQuery's `authToken` is the authorization token handed back by the other Comparator authorizing
// the comparison on a document.
//
//
// swagger:model AuthorizedQuery
type AuthorizedQuery struct {

	// auth token
	// Required: true
	AuthToken *string `json:"authToken"`
}

// Type gets the type of this subtype
func (m *AuthorizedQuery) Type() string {
	return "AuthorizedQuery"
}

// SetType sets the type of this subtype
func (m *AuthorizedQuery) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AuthorizedQuery) UnmarshalJSON(raw []byte) error {
	var data struct {

		// auth token
		// Required: true
		AuthToken *string `json:"authToken"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AuthorizedQuery

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.AuthToken = data.AuthToken

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AuthorizedQuery) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// auth token
		// Required: true
		AuthToken *string `json:"authToken"`
	}{

		AuthToken: m.AuthToken,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this authorized query
func (m *AuthorizedQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizedQuery) validateAuthToken(formats strfmt.Registry) error {

	if err := validate.Required("authToken", "body", m.AuthToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this authorized query based on the context it is used
func (m *AuthorizedQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizedQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizedQuery) UnmarshalBinary(b []byte) error {
	var res AuthorizedQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
