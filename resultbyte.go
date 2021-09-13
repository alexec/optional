// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultByte is an result byte.
type ResultByte struct {
	value *byte
	err   error
}

func NewResultByte(v byte, err error) ResultByte {
	if err != nil {
		return ErrByte(err)
	} else {
		return OkByte(v)
	}
}

func NewResultBytePtr(v *byte, err error) ResultByte {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrByte(err)
	} else {
		return OkByte(*v)
	}
}

func OkByte(v byte) ResultByte {
	return ResultByte{value: &v}
}

func ErrByte(err error) ResultByte {
	return ResultByte{err: err}
}

func (r ResultByte) Value() byte {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultByte) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultByte) Get() (byte, error) {
	if r.value == nil {
		var zero byte
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultByte) Present() bool {
	return r.value != nil
}