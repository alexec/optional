// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultUint64 is an result uint64.
type ResultUint64 struct {
	value *uint64
	err   error
}

func NewResultUint64(v uint64, err error) ResultUint64 {
	if err != nil {
		return ErrUint64(err)
	} else {
		return OkUint64(v)
	}
}

func NewResultUint64Ptr(v *uint64, err error) ResultUint64 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrUint64(err)
	} else {
		return OkUint64(*v)
	}
}

func OkUint64(v uint64) ResultUint64 {
	return ResultUint64{value: &v}
}

func ErrUint64(err error) ResultUint64 {
	return ResultUint64{err: err}
}

func (r ResultUint64) Value() uint64 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultUint64) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultUint64) Get() (uint64, error) {
	if r.value == nil {
		var zero uint64
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultUint64) Present() bool {
	return r.value != nil
}
