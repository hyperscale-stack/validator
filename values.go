// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import "net/url"

// ValuesValidator interface.
type ValuesValidator interface {
	Validate(values url.Values) map[string][]error
}

type valuesValidator struct {
	validators map[string][]Validator
}

// NewValuesValidator constructor.
func NewValuesValidator(validators map[string][]Validator) ValuesValidator {
	return &valuesValidator{
		validators: validators,
	}
}

func (v valuesValidator) validateField(key string, value interface{}) []error {
	errs := []error{}

	validators, ok := v.validators[key]
	if !ok {
		return errs
	}

	for _, validator := range validators {
		if err := validator.Validate(value); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (v valuesValidator) Validate(values url.Values) map[string][]error {
	errs := make(map[string][]error)

	for key, vals := range values {
		for _, val := range vals {
			if err := v.validateField(key, val); len(err) > 0 {
				errs[key] = err
			}
		}
	}

	return errs
}
