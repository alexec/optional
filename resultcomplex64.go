// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultComplex64 is an result complex64.
type ResultComplex64 struct {
	value *complex64
	err   error
}

func NewResultComplex64(v complex64, err error) ResultComplex64 {
	if err != nil {
		return ErrComplex64(err)
	} else {
		return OkComplex64(v)
	}
}

func NewResultComplex64Ptr(v *complex64, err error) ResultComplex64 {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrComplex64(err)
	} else {
		return OkComplex64(*v)
	}
}

func OkComplex64(v complex64) ResultComplex64 {
	return ResultComplex64{value: &v}
}

func ErrComplex64(err error) ResultComplex64 {
	return ResultComplex64{err: err}
}

func (r ResultComplex64) Value() complex64 {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultComplex64) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultComplex64) Get() (complex64, error) {
	if r.value == nil {
		var zero complex64
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultComplex64) Present() bool {
	return r.value != nil
}