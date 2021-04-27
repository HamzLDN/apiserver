/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package labels

import (
	"bytes"
	"fmt"
	"strings"
)

// Requirements is AND of all requirements.
type Requirements []Requirement

// Selector represents a label selector.
type Selector interface {
	// Matches returns true if this selector matches the given set of labels.
	//Matches(Labels) bool

	// Empty returns true if this selector does not restrict the selection space.
	Empty() bool

	// String returns a human readable string that represents this selector.
	String() string

	// Add adds requirements to the Selector
	Add(r ...Requirement) Selector

	// Requirements converts this interface into Requirements to expose
	// more detailed selection information.
	// If there are querying parameters, it will return converted requirements and selectable=true.
	// If this selector doesn't want to select anything, it will return selectable=false.
	Requirements() (requirements Requirements, selectable bool)

	// Make a deep copy of the selector.
	DeepCopySelector() Selector

	// RequiresExactMatch allows a caller to introspect whether a given selector
	// requires a single specific label to be set, and if so returns the value it
	// requires.
	//RequiresExactMatch(label string) (value string, found bool)
}

// Everything returns a selector that matches all labels.
func Everything() Selector {
	return internalSelector{}
}

type nothingSelector struct{}

//func (n nothingSelector) Matches(_ Labels) bool              { return false }
func (n nothingSelector) Empty() bool                        { return false }
func (n nothingSelector) String() string                     { return "" }
func (n nothingSelector) Add(_ ...Requirement) Selector      { return n }
func (n nothingSelector) Requirements() (Requirements, bool) { return nil, false }
func (n nothingSelector) DeepCopySelector() Selector         { return n }

//func (n nothingSelector) RequiresExactMatch(label string) (value string, found bool) {
//	return "", false
//}

// Nothing returns a selector that matches no labels
func Nothing() Selector {
	return nothingSelector{}
}

// NewSelector returns a nil selector
func NewSelector() Selector {
	return internalSelector(nil)
}

type internalSelector []Requirement

func (s internalSelector) DeepCopy() internalSelector {
	if s == nil {
		return nil
	}
	result := make([]Requirement, len(s))
	for i := range s {
		s[i].DeepCopyInto(&result[i])
	}
	return result
}

func (s internalSelector) DeepCopySelector() Selector {
	return s.DeepCopy()
}

type Requirement struct {
	query string
	// In huge majority of cases we have at most one value here.
	// It is generally faster to operate on a single-element slice
	// than on a single-element map, so we have a slice here.
	args []interface{}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requirement) DeepCopyInto(out *Requirement) {
	*out = *in
	if in.args != nil {
		in, out := &in.args, &out.args
		*out = make([]interface{}, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requirement.
func (in *Requirement) DeepCopy() *Requirement {
	if in == nil {
		return nil
	}
	out := new(Requirement)
	in.DeepCopyInto(out)
	return out
}

// NewRequirement is the constructor for a Requirement.
// If any of these rules is violated, an error is returned:
// (1) The operator can only be In, NotIn, Equals, DoubleEquals, NotEquals, Exists, or DoesNotExist.
// (2) If the operator is In or NotIn, the values set must be non-empty.
// (3) If the operator is Equals, DoubleEquals, or NotEquals, the values set must contain one value.
// (4) If the operator is Exists or DoesNotExist, the value set must be empty.
// (5) If the operator is Gt or Lt, the values set must contain only one value, which will be interpreted as an integer.
// (6) The key is invalid due to its length, or sequence
//     of characters. See validateLabelKey for more details.
//
// The empty string is a valid value in the input values set.
// Returned error, if not nil, is guaranteed to be an aggregated field.ErrorList
func NewRequirement(query string, args []interface{}) *Requirement {
	return &Requirement{query: query, args: args}
}

//func (r *Requirement) hasValue(value string) bool {
//	for i := range r.strValues {
//		if r.strValues[i] == value {
//			return true
//		}
//	}
//	return false
//}

//func (r *Requirement) Matches(ls Labels) bool {
//}

// Key returns requirement key
//func (r *Requirement) Key() string {
//	return r.key
//}

// Operator returns requirement operator
//func (r *Requirement) Operator() selection.Operator {
//	return r.operator
//}

// Values returns requirement values
//func (r *Requirement) Values() sets.String {
//	ret := sets.String{}
//	for i := range r.strValues {
//		ret.Insert(r.strValues[i])
//	}
//	return ret
//}

// Equal checks the equality of requirement.
//func (r Requirement) Equal(x Requirement) bool {
//	if r.key != x.key {
//		return false
//	}
//	if r.operator != x.operator {
//		return false
//	}
//	return cmp.Equal(r.strValues, x.strValues)
//}

// Empty returns true if the internalSelector doesn't restrict selection space
func (s internalSelector) Empty() bool {
	if s == nil {
		return true
	}
	return len(s) == 0
}

// String returns a human-readable string that represents this
// Requirement. If called on an invalid Requirement, an error is
// returned. See NewRequirement for creating a valid Requirement.
func (r *Requirement) String() string {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, r.query)
	fmt.Fprint(buf, r.args...)

	return buf.String()
}

// Add adds requirements to the selector. It copies the current selector returning a new one
func (s internalSelector) Add(reqs ...Requirement) Selector {
	var ret internalSelector
	for ix := range s {
		ret = append(ret, s[ix])
	}
	for _, r := range reqs {
		ret = append(ret, r)
	}
	return ret
}

// Matches for a internalSelector returns true if all
// its Requirements match the input Labels. If any
// Requirement does not match, false is returned.
//func (s internalSelector) Matches(l Labels) bool {
//	for ix := range s {
//		if matches := s[ix].Matches(l); !matches {
//			return false
//		}
//	}
//	return true
//}

func (s internalSelector) Requirements() (Requirements, bool) { return Requirements(s), true }

// String returns a comma-separated string of all
// the internalSelector Requirements' human-readable strings.
func (s internalSelector) String() string {
	var reqs []string
	for ix := range s {
		reqs = append(reqs, s[ix].String())
	}
	return strings.Join(reqs, ",")
}

// RequiresExactMatch introspects whether a given selector requires a single specific field
// to be set, and if so returns the value it requires.
//func (s internalSelector) RequiresExactMatch(label string) (value string, found bool) {
//	for ix := range s {
//		if s[ix].key == label {
//			switch s[ix].operator {
//			case selection.Equals, selection.DoubleEquals, selection.In:
//				if len(s[ix].strValues) == 1 {
//					return s[ix].strValues[0], true
//				}
//			}
//			return "", false
//		}
//	}
//	return "", false
//}
//

func SelectorFromSet(ls Set) Selector {
	return SelectorFromValidatedSet(ls)
}
func ValidatedSelectorFromSet(ls Set) (Selector, error) {
	if ls == nil || len(ls) == 0 {
		return internalSelector{}, nil
	}
	requirements := make([]Requirement, 0, len(ls))
	for label, value := range ls {
		r := NewRequirement(label+"=?", []interface{}{value})
		requirements = append(requirements, *r)
	}
	return internalSelector(requirements), nil
}
func SelectorFromValidatedSet(ls Set) Selector {
	if ls == nil || len(ls) == 0 {
		return internalSelector{}
	}
	requirements := make([]Requirement, 0, len(ls))
	for label, value := range ls {
		requirements = append(requirements, Requirement{query: label + "=?", args: []interface{}{value}})
	}
	// sort to have deterministic string representation
	return internalSelector(requirements)
}
