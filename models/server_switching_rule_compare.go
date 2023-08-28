// Code generated with struct_equal_generator; DO NOT EDIT.

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

// Equal checks if two structs of type ServerSwitchingRule are equal
//
// By default empty maps and slices are equal to nil:
//
//	var a, b ServerSwitchingRule
//	equal := a.Equal(b)
//
// For more advanced use case you can configure these options (default values are shown):
//
//	var a, b ServerSwitchingRule
//	equal := a.Equal(b,Options{
//		SkipIndex: true,
//	})
func (s ServerSwitchingRule) Equal(t ServerSwitchingRule, opts ...Options) bool {
	opt := getOptions(opts...)

	if s.Cond != t.Cond {
		return false
	}

	if s.CondTest != t.CondTest {
		return false
	}

	if !opt.SkipIndex && !equalPointers(s.Index, t.Index) {
		return false
	}

	if s.TargetServer != t.TargetServer {
		return false
	}

	return true
}

// Diff checks if two structs of type ServerSwitchingRule are equal
//
//	var a, b ServerSwitchingRule
//	diff := a.Diff(b)
//
// For more advanced use case you can configure the options (default values are shown):
//
//	var a, b ServerSwitchingRule
//	equal := a.Diff(b,Options{
//		SkipIndex: true,
//	})
func (s ServerSwitchingRule) Diff(t ServerSwitchingRule, opts ...Options) map[string][]interface{} {
	opt := getOptions(opts...)

	diff := make(map[string][]interface{})
	if s.Cond != t.Cond {
		diff["Cond"] = []interface{}{s.Cond, t.Cond}
	}

	if s.CondTest != t.CondTest {
		diff["CondTest"] = []interface{}{s.CondTest, t.CondTest}
	}

	if !opt.SkipIndex && !equalPointers(s.Index, t.Index) {
		diff["Index"] = []interface{}{ValueOrNil(s.Index), ValueOrNil(t.Index)}
	}

	if s.TargetServer != t.TargetServer {
		diff["TargetServer"] = []interface{}{s.TargetServer, t.TargetServer}
	}

	return diff
}
