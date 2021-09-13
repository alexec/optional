// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultUint16 is an result uint16.
type ResultUint16 struct {
	value *uint16
	err   error
}

func NewResultUint16(v uint16, err error) ResultUint16 {
	if err != nil {
		return ErrUint16(err)
	} else {
		return OkUint16(v)
	}
}

func NewResultUint16Ptr(v *uint16, err error) ResultUint16 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrUint16(err)
	} else {
		return OkUint16(*v)
	}
}

func OkUint16(v uint16) ResultUint16 {
	return ResultUint16{value: &v}
}

func ErrUint16(err error) ResultUint16 {
	return ResultUint16{err: err}
}

func (r ResultUint16) Value() uint16 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultUint16) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultUint16) Get() (uint16, error) {
	if r.value == nil {
		var zero uint16
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultUint16) Present() bool {
	return r.value != nil
}