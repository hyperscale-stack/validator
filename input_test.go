// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputValidator(t *testing.T) {
	i := NewInputValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	errs := i.ValidateMap(map[string]interface{}{
		"uuid": "9D2C8507-5F9D-4CB0-A098-2E307B39DC91",
		"name": "Title",
	})
	assert.Equal(t, 0, len(errs))

	values := url.Values{}
	values.Set("uuid", "9D2C8507-5F9D-4CB0-A098-2E307B39DC91")
	values.Set("name", "Title")

	errs = i.ValidateValues(values)
	assert.Equal(t, 0, len(errs))

}

func TestInputValidatorWithBadInput(t *testing.T) {
	i := NewInputValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	errs := i.ValidateMap(map[string]interface{}{
		"uuid": "9D2C8507-5F9D-7CB0-A098-2E307B391",
		"name": "Title",
	})
	assert.Equal(t, 1, len(errs))

	assert.Contains(t, errs, "uuid")
	assert.Error(t, errs["uuid"][0])

	values := url.Values{}
	values.Set("uuid", "9D2C8507-5F9D-7CB0-A098-2E307B391")
	values.Set("name", "Title")

	errs = i.ValidateValues(values)
	assert.Equal(t, 1, len(errs))

	assert.Contains(t, errs, "uuid")
	assert.Error(t, errs["uuid"][0])
}

func BenchmarkInputValidatorMap(b *testing.B) {
	v := NewInputValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	input := map[string]interface{}{
		"uuid": "9D2C8507-5F9D-4CB0-A098-2E307B39DC91",
		"name": "Title",
	}

	for i := 0; i < b.N; i++ {
		v.ValidateMap(input)
	}
}

func BenchmarkInputValidatorValues(b *testing.B) {
	v := NewInputValidator(map[string][]Validator{
		"uuid": {
			NewUUIDValidator(),
		},
	})

	values := url.Values{}
	values.Set("uuid", "9D2C8507-5F9D-4CB0-A098-2E307B39DC91")
	values.Set("name", "Title")

	for i := 0; i < b.N; i++ {
		v.ValidateValues(values)
	}
}
