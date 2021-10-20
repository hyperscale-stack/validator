// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValuesValidator(t *testing.T) {
	i := NewValuesValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	values := url.Values{}
	values.Set("uuid", "9D2C8507-5F9D-4CB0-A098-2E307B39DC91")
	values.Set("name", "Title")

	errs := i.Validate(values)
	assert.Equal(t, 0, len(errs))
}

func TestValuesValidatorWithBadInput(t *testing.T) {
	i := NewValuesValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	values := url.Values{}
	values.Set("uuid", "9D2C8507-5F9D-7CB0-A098-2E307B391")
	values.Set("name", "Title")

	errs := i.Validate(values)
	assert.Equal(t, 1, len(errs))

	assert.Contains(t, errs, "uuid")
	assert.Error(t, errs["uuid"][0])
}
