// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultInt16 is an result int16.
type ResultInt16 struct {
	value *int16
	err   error
}

func NewResultInt16(v int16, err error) ResultInt16 {
	if err != nil {
		return ErrInt16(err)
	} else {
		return OkInt16(v)
	}
}

func NewResultInt16Ptr(v *int16, err error) ResultInt16 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrInt16(err)
	} else {
		return OkInt16(*v)
	}
}

func OkInt16(v int16) ResultInt16 {
	return ResultInt16{value: &v}
}

func ErrInt16(err error) ResultInt16 {
	return ResultInt16{err: err}
}

func (r ResultInt16) Value() int16 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultInt16) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultInt16) Get() (int16, error) {
	if r.value == nil {
		var zero int16
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultInt16) Present() bool {
	return r.value != nil
}
