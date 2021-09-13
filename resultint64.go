// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultInt64 is an result int64.
type ResultInt64 struct {
	value *int64
	err   error
}

func NewResultInt64(v int64, err error) ResultInt64 {
	if err != nil {
		return ErrInt64(err)
	} else {
		return OkInt64(v)
	}
}

func NewResultInt64Ptr(v *int64, err error) ResultInt64 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrInt64(err)
	} else {
		return OkInt64(*v)
	}
}

func OkInt64(v int64) ResultInt64 {
	return ResultInt64{value: &v}
}

func ErrInt64(err error) ResultInt64 {
	return ResultInt64{err: err}
}

func (r ResultInt64) Value() int64 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultInt64) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultInt64) Get() (int64, error) {
	if r.value == nil {
		var zero int64
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultInt64) Present() bool {
	return r.value != nil
}