// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"fmt"
	"math"
	"reflect"
)

// StringLength config options.
type StringLength struct {
	Min int
	Max int
}

type stringLengthValidator struct {
	opts StringLength
}

// NewStringLengthValidator constructor.
func NewStringLengthValidator(opts StringLength) Validator {
	if opts.Max == 0 {
		opts.Max = math.MaxInt64
	}

	return &stringLengthValidator{
		opts: opts,
	}
}

func (v stringLengthValidator) Validate(input interface{}) error {
	switch val := input.(type) {
	case string:
		size := len(val)

		if size < v.opts.Min {
			return fmt.Errorf("the input is less than %d characters long", v.opts.Min)
		}

		if size > v.opts.Max {
			return fmt.Errorf("the input is more than %d characters long", v.opts.Max)
		}
	default:
		return fmt.Errorf("invalid \"%s\" type given. String expected", reflect.TypeOf(val))
	}

	return nil
}
