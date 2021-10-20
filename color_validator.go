// Copyright 2020 UX Stack. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

var colorHexPattern = regexp.MustCompile(`^#([0-9a-fA-F]{3}){1,2}$`)

type colorValidator struct {
}

// NewColorValidator constructor.
func NewColorValidator() Validator {
	return &colorValidator{}
}

func (v colorValidator) Validate(input interface{}) error {
	switch color := input.(type) {
	case string:
		if !colorHexPattern.MatchString(color) {
			return fmt.Errorf("color %q is invalid", color)
		}

	default:
		return fmt.Errorf("invalid input type \"%v\" for color validator", reflect.TypeOf(color))
	}

	return nil
}
