// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import "time"

// EmailOption type.
type EmailOption func(*emailValidator)

// EmailTimeout remove all utm_* query parameters.
func EmailTimeout(timeout time.Duration) EmailOption {
	return func(v *emailValidator) {
		v.timeout = timeout
	}
}
