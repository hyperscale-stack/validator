// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type uuidValidator struct {
}

// NewUUIDValidator constructor.
func NewUUIDValidator() Validator {
	return &uuidValidator{}
}

func (v uuidValidator) Validate(input interface{}) error {
	switch val := input.(type) {
	case string:
		if _, err := uuid.Parse(val); err != nil {
			return err
		}
	case []byte:
		if _, err := uuid.FromBytes(val); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid input type \"%v\" for UUID validator", reflect.TypeOf(val))
	}

	return nil
}
