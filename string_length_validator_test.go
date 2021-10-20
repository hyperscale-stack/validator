// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLengthValidatorWithBadType(t *testing.T) {
	v := NewStringLengthValidator(StringLength{})

	assert.EqualError(t, v.Validate(123), "invalid \"int\" type given. String expected")
}

func TestStringLengthValidatorWithMaxOption(t *testing.T) {
	v := NewStringLengthValidator(StringLength{
		Max: 6,
	})

	assert.Nil(t, v.Validate("Test"))
	assert.EqualError(t, v.Validate("Testing"), "the input is more than 6 characters long")
}

func TestStringLengthValidatorWithMinOption(t *testing.T) {
	v := NewStringLengthValidator(StringLength{
		Min: 5,
	})

	assert.EqualError(t, v.Validate("Test"), "the input is less than 5 characters long")
	assert.Nil(t, v.Validate("Testing"))
}

func TestStringLengthValidatorWithMinAndMaxOption(t *testing.T) {
	v := NewStringLengthValidator(StringLength{
		Min: 5,
		Max: 6,
	})

	assert.EqualError(t, v.Validate("Test"), "the input is less than 5 characters long")
	assert.EqualError(t, v.Validate("Testing"), "the input is more than 6 characters long")
	assert.Nil(t, v.Validate("good t"))
}
