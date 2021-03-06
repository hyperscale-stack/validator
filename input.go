// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import "net/url"

// InputValidator interface.
type InputValidator interface {
	ValidateMap(input map[string]interface{}) map[string][]error
	ValidateValues(values url.Values) map[string][]error
}

type inputValidator struct {
	validators map[string][]Validator
}

// NewInputValidator constructor.
func NewInputValidator(validators map[string][]Validator) InputValidator {
	return &inputValidator{
		validators: validators,
	}
}

func (v inputValidator) validateField(key string, value interface{}) []error {
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

func (v inputValidator) ValidateMap(input map[string]interface{}) map[string][]error {
	errs := make(map[string][]error)

	for key, val := range input {
		if err := v.validateField(key, val); len(err) > 0 {
			errs[key] = err
		}
	}

	return errs
}

func (v inputValidator) ValidateValues(values url.Values) map[string][]error {
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
