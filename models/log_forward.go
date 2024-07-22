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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogForward LogForward with all it's children resources
//
// swagger:model LogForward
type LogForward struct {
	LogForwardBase `json:",inline"`

	// log target list
	LogTargetList LogTargets `json:"log_target_list,omitempty"`

	// binds
	Binds map[string]Bind `json:"binds,omitempty"`

	// dgram binds
	DgramBinds map[string]DgramBind `json:"dgram_binds,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LogForward) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 LogForwardBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.LogForwardBase = aO0

	// AO1
	var dataAO1 struct {
		LogTargetList LogTargets `json:"log_target_list,omitempty"`

		Binds map[string]Bind `json:"binds,omitempty"`

		DgramBinds map[string]DgramBind `json:"dgram_binds,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LogTargetList = dataAO1.LogTargetList

	m.Binds = dataAO1.Binds

	m.DgramBinds = dataAO1.DgramBinds

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LogForward) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.LogForwardBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		LogTargetList LogTargets `json:"log_target_list,omitempty"`

		Binds map[string]Bind `json:"binds,omitempty"`

		DgramBinds map[string]DgramBind `json:"dgram_binds,omitempty"`
	}

	dataAO1.LogTargetList = m.LogTargetList

	dataAO1.Binds = m.Binds

	dataAO1.DgramBinds = m.DgramBinds

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this log forward
func (m *LogForward) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogForwardBase
	if err := m.LogForwardBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTargetList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDgramBinds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogForward) validateLogTargetList(formats strfmt.Registry) error {

	if swag.IsZero(m.LogTargetList) { // not required
		return nil
	}

	if err := m.LogTargetList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("log_target_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("log_target_list")
		}
		return err
	}

	return nil
}

func (m *LogForward) validateBinds(formats strfmt.Registry) error {

	if swag.IsZero(m.Binds) { // not required
		return nil
	}

	for k := range m.Binds {

		if err := validate.Required("binds"+"."+k, "body", m.Binds[k]); err != nil {
			return err
		}
		if val, ok := m.Binds[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("binds" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("binds" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogForward) validateDgramBinds(formats strfmt.Registry) error {

	if swag.IsZero(m.DgramBinds) { // not required
		return nil
	}

	for k := range m.DgramBinds {

		if err := validate.Required("dgram_binds"+"."+k, "body", m.DgramBinds[k]); err != nil {
			return err
		}
		if val, ok := m.DgramBinds[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dgram_binds" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dgram_binds" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this log forward based on the context it is used
func (m *LogForward) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with LogForwardBase
	if err := m.LogForwardBase.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogTargetList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDgramBinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogForward) contextValidateLogTargetList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LogTargetList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("log_target_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("log_target_list")
		}
		return err
	}

	return nil
}

func (m *LogForward) contextValidateBinds(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Binds {

		if val, ok := m.Binds[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *LogForward) contextValidateDgramBinds(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.DgramBinds {

		if val, ok := m.DgramBinds[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogForward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogForward) UnmarshalBinary(b []byte) error {
	var res LogForward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
