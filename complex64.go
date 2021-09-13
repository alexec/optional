// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// Complex64 is an optional complex64.
type Complex64 struct {
	value *complex64
}

// NewComplex64 creates an optional.Complex64 from a complex64.
func NewComplex64(v complex64) Complex64 {
	return Complex64{&v}
}

// Set sets the complex64 value.
func (c *Complex64) Set(v complex64) {
	c.value = &v
}

// Get returns the complex64 value or an error if not present.
func (c Complex64) Get() (complex64, error) {
	if !c.Present() {
		var zero complex64
		return zero, errors.New("value not present")
	}
	return *c.value, nil
}

// MustGet returns the complex64 value or panics if not present.
func (c Complex64) MustGet() complex64 {
	if !c.Present() {
		panic("value not present")
	}
	return *c.value
}

// Present returns whether or not the value is present.
func (c Complex64) Present() bool {
	return c.value != nil
}

// OrElse returns the complex64 value or a default value if the value is not present.
func (c Complex64) OrElse(v complex64) complex64 {
	if c.Present() {
		return *c.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (c Complex64) If(fn func(complex64)) {
	if c.Present() {
		fn(*c.value)
	}
}
