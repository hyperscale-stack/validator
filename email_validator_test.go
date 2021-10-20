// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidator(t *testing.T) {
	v := NewEmailValidator(EmailTimeout(1 * time.Second))

	assert.EqualError(t, v.Validate(124), "invalid input type \"int\" for email validator")
	assert.Error(t, v.Validate("bad"))
	assert.Error(t, v.Validate("bad@bad-domain-name.tld"))
	assert.Error(t, v.Validate("user@perdu.com"))

	assert.NoError(t, v.Validate("euskadi31@gmail.com"))
	assert.NoError(t, v.Validate("axel@etcheverry.biz"))
}

func TestEmailValidatorSingle(t *testing.T) {
	v := NewEmailValidator(EmailTimeout(1 * time.Second))

	assert.NoError(t, v.Validate("euskadi31@gmail.com"))
}
