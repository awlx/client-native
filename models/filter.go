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

// Filter Filter
//
// HAProxy filters
// Example: {"trace_name":"name","trace_rnd_parsing":true,"type":"trace"}
//
// swagger:model filter
type Filter struct {
	// Name of the fcgi-app section this filter will use.
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	AppName string `json:"app_name,omitempty"`

	// Filter name that will be used by 'set-bandwidth-limit' actions to reference a specific bandwidth limitation filter
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	BandwidthLimitName string `json:"bandwidth_limit_name,omitempty"`

	// cache name
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	CacheName string `json:"cache_name,omitempty"`

	// The max number of bytes that can be forwarded over the period.
	// The value must be specified for per-stream and shared bandwidth limitation filters.
	// It follows the HAProxy size format and is expressed in bytes.
	DefaultLimit int64 `json:"default_limit,omitempty"`

	// The default time period used to evaluate the bandwidth limitation rate.
	// It can be specified for per-stream bandwidth limitation filters only.
	// It follows the HAProxy time format and is expressed in milliseconds.
	DefaultPeriod int64 `json:"default_period,omitempty"`

	// A sample expression rule.
	// It describes what elements will be analyzed, extracted, combined, and used to select which table entry to update the counters.
	// It must be specified for shared bandwidth limitation filters only.
	Key string `json:"key,omitempty"`

	// The max number of bytes that can be forwarded over the period.
	// The value must be specified for per-stream and shared bandwidth limitation filters.
	// It follows the HAProxy size format and is expressed in bytes.
	Limit int64 `json:"limit,omitempty"`

	// The optional minimum number of bytes forwarded at a time by a stream excluding the last packet that may be smaller.
	// This value can be specified for per-stream and shared bandwidth limitation filters.
	// It follows the HAProxy size format and is expressed in bytes.
	MinSize int64 `json:"min_size,omitempty"`

	// spoe config
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	SpoeConfig string `json:"spoe_config,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// An optional table to be used instead of the default one, which is the stick-table declared in the current proxy.
	// It can be specified for shared bandwidth limitation filters only.
	Table string `json:"table,omitempty"`

	// trace hexdump
	TraceHexdump bool `json:"trace_hexdump,omitempty"`

	// trace name
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	TraceName string `json:"trace_name,omitempty"`

	// trace rnd forwarding
	TraceRndForwarding bool `json:"trace_rnd_forwarding,omitempty"`

	// trace rnd parsing
	TraceRndParsing bool `json:"trace_rnd_parsing,omitempty"`

	// type
	// Required: true
	// Enum: [bwlim-in bwlim-out cache compression fcgi-app spoe trace]
	// +kubebuilder:validation:Enum=bwlim-in;bwlim-out;cache;compression;fcgi-app;spoe;trace;
	Type string `json:"type"`
}

// Validate validates this filter
func (m *Filter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthLimitName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceName(formats); err != nil {
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

func (m *Filter) validateAppName(formats strfmt.Registry) error {
	if swag.IsZero(m.AppName) { // not required
		return nil
	}

	if err := validate.Pattern("app_name", "body", m.AppName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateBandwidthLimitName(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthLimitName) { // not required
		return nil
	}

	if err := validate.Pattern("bandwidth_limit_name", "body", m.BandwidthLimitName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateCacheName(formats strfmt.Registry) error {
	if swag.IsZero(m.CacheName) { // not required
		return nil
	}

	if err := validate.Pattern("cache_name", "body", m.CacheName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateSpoeConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SpoeConfig) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_config", "body", m.SpoeConfig, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateSpoeEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", m.SpoeEngine, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateTraceName(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceName) { // not required
		return nil
	}

	if err := validate.Pattern("trace_name", "body", m.TraceName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var filterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bwlim-in","bwlim-out","cache","compression","fcgi-app","spoe","trace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filterTypeTypePropEnum = append(filterTypeTypePropEnum, v)
	}
}

const (

	// FilterTypeBwlimDashIn captures enum value "bwlim-in"
	FilterTypeBwlimDashIn string = "bwlim-in"

	// FilterTypeBwlimDashOut captures enum value "bwlim-out"
	FilterTypeBwlimDashOut string = "bwlim-out"

	// FilterTypeCache captures enum value "cache"
	FilterTypeCache string = "cache"

	// FilterTypeCompression captures enum value "compression"
	FilterTypeCompression string = "compression"

	// FilterTypeFCGIDashApp captures enum value "fcgi-app"
	FilterTypeFCGIDashApp string = "fcgi-app"

	// FilterTypeSpoe captures enum value "spoe"
	FilterTypeSpoe string = "spoe"

	// FilterTypeTrace captures enum value "trace"
	FilterTypeTrace string = "trace"
)

// prop value enum
func (m *Filter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Filter) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this filter based on context it is used
func (m *Filter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Filter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Filter) UnmarshalBinary(b []byte) error {
	var res Filter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
