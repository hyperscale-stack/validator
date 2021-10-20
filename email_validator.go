// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"context"
	"fmt"
	"net"
	"net/mail"
	"reflect"
	"strings"
	"time"
)

type emailValidator struct {
	timeout time.Duration
}

// NewEmailValidator constructor.
func NewEmailValidator(opts ...EmailOption) Validator {
	v := &emailValidator{
		timeout: 100 * time.Millisecond,
	}

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func (v emailValidator) Validate(input interface{}) error {
	switch email := input.(type) {
	case string:
		addr, err := mail.ParseAddress(email)
		if err != nil {
			return err
		}

		parts := strings.Split(addr.Address, "@")

		ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
		defer cancel()

		if _, err := net.DefaultResolver.LookupMX(ctx, parts[1]); err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid input type \"%v\" for email validator", reflect.TypeOf(email))
	}

	return nil
}
