Hyperscale Validator [![Last release](https://img.shields.io/github/release/hyperscale-stack/validator.svg)](https://github.com/hyperscale-stack/validator/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/validator?status.svg)](https://godoc.org/github.com/hyperscale-stack/validator)
====================

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/validator)](https://goreportcard.com/report/github.com/hyperscale-stack/validator)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://github.com/hyperscale-stack/validator/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/validator/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/validator/master.svg)](https://coveralls.io/github/hyperscale-stack/validator?branch=master) |

The Hyperscale Validator library provides a set of commonly needed data validators. It also provides a simple validator chaining mechanism by which multiple validators may be applied to a single datum in a user-defined order. 

## Example

Validate by `map[string]interface{}`

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/validator"
)

func main() {
    i := NewInputValidator(map[string][]Validator{
		"email": {
			NewEmailValidator(EmailTimeout(1 * time.Second)),
		},
	})

	errs := i.Validate(map[string]interface{}{
		"email":  "bad",
    })
    // return 
    // map[string][]error{
	//     "email": []error{...},
    // }
}

```


Validate by `url.Values`

```go
package main

import (
    "fmt"
    "github.com/hyperscale-stack/validator"
)

func main() {
    i := NewValuesValidator(map[string][]Validator{
		"email": {
			NewEmailValidator(EmailTimeout(1 * time.Second)),
		},
	})

    values := url.Values{}
    values.Set("email", "bad")

	errs := i.Validate(values)
    // return 
    // map[string][]error{
	//     "email": []error{...},
    // }
}

```


## License

Hyperscale Validator is licensed under [the MIT license](LICENSE.md).
