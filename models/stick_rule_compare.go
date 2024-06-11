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

// Equal checks if two structs of type StickRule are equal
//
//	var a, b StickRule
//	equal := a.Equal(b)
//
// opts ...Options are ignored in this method
func (s StickRule) Equal(t StickRule, opts ...Options) bool {
	if s.Cond != t.Cond {
		return false
	}

	if s.CondTest != t.CondTest {
		return false
	}

	if s.Pattern != t.Pattern {
		return false
	}

	if s.Table != t.Table {
		return false
	}

	if s.Type != t.Type {
		return false
	}

	return true
}

// Diff checks if two structs of type StickRule are equal
//
//	var a, b StickRule
//	diff := a.Diff(b)
//
// opts ...Options are ignored in this method
func (s StickRule) Diff(t StickRule, opts ...Options) map[string][]interface{} {
	diff := make(map[string][]interface{})
	if s.Cond != t.Cond {
		diff["Cond"] = []interface{}{s.Cond, t.Cond}
	}

	if s.CondTest != t.CondTest {
		diff["CondTest"] = []interface{}{s.CondTest, t.CondTest}
	}

	if s.Pattern != t.Pattern {
		diff["Pattern"] = []interface{}{s.Pattern, t.Pattern}
	}

	if s.Table != t.Table {
		diff["Table"] = []interface{}{s.Table, t.Table}
	}

	if s.Type != t.Type {
		diff["Type"] = []interface{}{s.Type, t.Type}
	}

	return diff
}
