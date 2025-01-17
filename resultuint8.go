// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultUint8 is an result uint8.
type ResultUint8 struct {
	value *uint8
	err   error
}

func NewResultUint8(v uint8, err error) ResultUint8 {
	if err != nil {
		return ErrUint8(err)
	} else {
		return OkUint8(v)
	}
}

func NewResultUint8Ptr(v *uint8, err error) ResultUint8 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrUint8(err)
	} else {
		return OkUint8(*v)
	}
}

func OkUint8(v uint8) ResultUint8 {
	return ResultUint8{value: &v}
}

func ErrUint8(err error) ResultUint8 {
	return ResultUint8{err: err}
}

func (r ResultUint8) Value() uint8 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultUint8) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultUint8) Get() (uint8, error) {
	if r.value == nil {
		var zero uint8
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultUint8) Present() bool {
	return r.value != nil
}
