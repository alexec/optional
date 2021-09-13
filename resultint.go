// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultInt is an result int.
type ResultInt struct {
	value *int
	err   error
}

func NewResultInt(v int, err error) ResultInt {
	if err != nil {
		return ErrInt(err)
	} else {
		return OkInt(v)
	}
}

func NewResultIntPtr(v *int, err error) ResultInt {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrInt(err)
	} else {
		return OkInt(*v)
	}
}

func OkInt(v int) ResultInt {
	return ResultInt{value: &v}
}

func ErrInt(err error) ResultInt {
	return ResultInt{err: err}
}

func (r ResultInt) Value() int {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultInt) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultInt) Get() (int, error) {
	if r.value == nil {
		var zero int
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultInt) Present() bool {
	return r.value != nil
}
