// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorValidator(t *testing.T) {
	v := NewColorValidator()

	assert.EqualError(t, v.Validate(124), "invalid input type \"int\" for color validator")
	assert.Error(t, v.Validate("bad"))
	assert.Error(t, v.Validate("#"))
	assert.Error(t, v.Validate("#f1"))

	assert.NoError(t, v.Validate("#000"))
	assert.NoError(t, v.Validate("#000000"))
	assert.NoError(t, v.Validate("#f6f6f6"))
	assert.NoError(t, v.Validate("#f6f"))
	assert.NoError(t, v.Validate("#F6F6F6"))
	assert.NoError(t, v.Validate("#F6F"))
}
