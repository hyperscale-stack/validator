// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordValidatorWithBadType(t *testing.T) {
	v := NewPasswordValidator(Password{})

	assert.EqualError(t, v.Validate(123), "invalid \"int\" type given. String expected")
}

func TestPasswordValidatorWithGoodPassword(t *testing.T) {
	v := NewPasswordValidator(Password{
		Min: 6,
		Max: 12,
	})

	for _, password := range []string{
		"Azerty@1",
		"A-ef546gfd&",
	} {
		assert.Nil(t, v.Validate(password))
	}
}

func TestPasswordValidatorWithBadPassword(t *testing.T) {
	v := NewPasswordValidator(Password{
		Min: 6,
		Max: 12,
	})

	for _, item := range []struct {
		password string
		err      string
	}{
		{
			password: "bad",
			err:      "the password is less than 6 characters long",
		},
		{
			password: "badfdsfsdffdd",
			err:      "the password is more than 12 characters long",
		},
		{
			password: "Azerty@",
			err:      "the password must have at least one numeric character",
		},
	} {
		assert.EqualError(t, v.Validate(item.password), item.err)
	}
}
