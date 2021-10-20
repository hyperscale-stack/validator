// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDValidator(t *testing.T) {
	v := NewUUIDValidator()

	assert.EqualError(t, v.Validate(124), "invalid input type \"int\" for UUID validator")
	assert.Error(t, v.Validate("bad"))
	assert.Error(t, v.Validate([]byte{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa,
	}))

	assert.NoError(t, v.Validate("9D2C8507-5F9D-4CB0-A098-2E307B39DC91"))
	assert.NoError(t, v.Validate([]byte{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}))
}
