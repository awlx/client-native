// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SpoeMessage SPOE message
//
// # SPOE message section configuration
//
// swagger:model spoe_message
type SpoeMessage struct {

	// acl
	ACL Acls `json:"acl,omitempty"`

	// args
	Args string `json:"args,omitempty"`

	// event
	Event *SpoeMessageEvent `json:"event,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this spoe message
func (m *SpoeMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpoeMessage) validateACL(formats strfmt.Registry) error {
	if swag.IsZero(m.ACL) { // not required
		return nil
	}

	if err := m.ACL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("acl")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("acl")
		}
		return err
	}

	return nil
}

func (m *SpoeMessage) validateEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *SpoeMessage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this spoe message based on the context it is used
func (m *SpoeMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpoeMessage) contextValidateACL(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ACL.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("acl")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("acl")
		}
		return err
	}

	return nil
}

func (m *SpoeMessage) contextValidateEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.Event != nil {
		if err := m.Event.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpoeMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpoeMessage) UnmarshalBinary(b []byte) error {
	var res SpoeMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SpoeMessageEvent spoe message event
//
// swagger:model SpoeMessageEvent
type SpoeMessageEvent struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// name
	// Required: true
	// Enum: [on-client-session on-server-session on-frontend-tcp-request on-backend-tcp-request on-tcp-response on-frontend-http-request on-backend-http-request on-http-response]
	Name *string `json:"name"`
}

// Validate validates this spoe message event
func (m *SpoeMessageEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var spoeMessageEventTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		spoeMessageEventTypeCondPropEnum = append(spoeMessageEventTypeCondPropEnum, v)
	}
}

const (

	// SpoeMessageEventCondIf captures enum value "if"
	SpoeMessageEventCondIf string = "if"

	// SpoeMessageEventCondUnless captures enum value "unless"
	SpoeMessageEventCondUnless string = "unless"
)

// prop value enum
func (m *SpoeMessageEvent) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, spoeMessageEventTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SpoeMessageEvent) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("event"+"."+"cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

var spoeMessageEventTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on-client-session","on-server-session","on-frontend-tcp-request","on-backend-tcp-request","on-tcp-response","on-frontend-http-request","on-backend-http-request","on-http-response"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		spoeMessageEventTypeNamePropEnum = append(spoeMessageEventTypeNamePropEnum, v)
	}
}

const (

	// SpoeMessageEventNameOnDashClientDashSession captures enum value "on-client-session"
	SpoeMessageEventNameOnDashClientDashSession string = "on-client-session"

	// SpoeMessageEventNameOnDashServerDashSession captures enum value "on-server-session"
	SpoeMessageEventNameOnDashServerDashSession string = "on-server-session"

	// SpoeMessageEventNameOnDashFrontendDashTCPDashRequest captures enum value "on-frontend-tcp-request"
	SpoeMessageEventNameOnDashFrontendDashTCPDashRequest string = "on-frontend-tcp-request"

	// SpoeMessageEventNameOnDashBackendDashTCPDashRequest captures enum value "on-backend-tcp-request"
	SpoeMessageEventNameOnDashBackendDashTCPDashRequest string = "on-backend-tcp-request"

	// SpoeMessageEventNameOnDashTCPDashResponse captures enum value "on-tcp-response"
	SpoeMessageEventNameOnDashTCPDashResponse string = "on-tcp-response"

	// SpoeMessageEventNameOnDashFrontendDashHTTPDashRequest captures enum value "on-frontend-http-request"
	SpoeMessageEventNameOnDashFrontendDashHTTPDashRequest string = "on-frontend-http-request"

	// SpoeMessageEventNameOnDashBackendDashHTTPDashRequest captures enum value "on-backend-http-request"
	SpoeMessageEventNameOnDashBackendDashHTTPDashRequest string = "on-backend-http-request"

	// SpoeMessageEventNameOnDashHTTPDashResponse captures enum value "on-http-response"
	SpoeMessageEventNameOnDashHTTPDashResponse string = "on-http-response"
)

// prop value enum
func (m *SpoeMessageEvent) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, spoeMessageEventTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SpoeMessageEvent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("event"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("event"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this spoe message event based on context it is used
func (m *SpoeMessageEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpoeMessageEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpoeMessageEvent) UnmarshalBinary(b []byte) error {
	var res SpoeMessageEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
