// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"fmt"
	"math"
	"reflect"
	"unicode"
)

// Password config options.
type Password struct {
	Min                int
	Max                int
	RequiredRangeTable map[string][]*unicode.RangeTable
}

type passwordValidator struct {
	opts Password
}

// NewPasswordValidator constructor.
func NewPasswordValidator(opts Password) Validator {
	if opts.Max == 0 {
		opts.Max = math.MaxInt64
	}

	if len(opts.RequiredRangeTable) == 0 {
		opts.RequiredRangeTable = map[string][]*unicode.RangeTable{
			"upper case": {unicode.Upper, unicode.Title},
			"lower case": {unicode.Lower},
			"numeric":    {unicode.Number, unicode.Digit},
			"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
		}
	}

	return &passwordValidator{
		opts: opts,
	}
}

func (v passwordValidator) Validate(input interface{}) error {
	switch val := input.(type) {
	case string:
		size := len(val)

		if size < v.opts.Min {
			return fmt.Errorf("the password is less than %d characters long", v.opts.Min)
		}

		if size > v.opts.Max {
			return fmt.Errorf("the password is more than %d characters long", v.opts.Max)
		}

		return v.checkPassword(val)
	default:
		return fmt.Errorf("invalid \"%s\" type given. String expected", reflect.TypeOf(val))
	}
}

func (v passwordValidator) checkPassword(password string) error {
next:
	for name, classes := range v.opts.RequiredRangeTable {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}

		return fmt.Errorf("the password must have at least one %s character", name)
	}

	return nil
}
