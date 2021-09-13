// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultFloat32 is an result float32.
type ResultFloat32 struct {
	value *float32
	err   error
}

func NewResultFloat32(v float32, err error) ResultFloat32 {
	if err != nil {
		return ErrFloat32(err)
	} else {
		return OkFloat32(v)
	}
}

func NewResultFloat32Ptr(v *float32, err error) ResultFloat32 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrFloat32(err)
	} else {
		return OkFloat32(*v)
	}
}

func OkFloat32(v float32) ResultFloat32 {
	return ResultFloat32{value: &v}
}

func ErrFloat32(err error) ResultFloat32 {
	return ResultFloat32{err: err}
}

func (r ResultFloat32) Value() float32 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultFloat32) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultFloat32) Get() (float32, error) {
	if r.value == nil {
		var zero float32
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultFloat32) Present() bool {
	return r.value != nil
}