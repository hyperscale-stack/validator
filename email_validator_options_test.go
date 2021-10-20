// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmailTimeout(t *testing.T) {
	v := &emailValidator{}

	EmailTimeout(1 * time.Second)(v)

	assert.Equal(t, 1*time.Second, v.timeout)
}
