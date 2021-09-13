// Code generated by optional. DO NOT EDIT.

package optional

import (
	"errors"
)

// ResultRune is an result rune.
type ResultRune struct {
	value *rune
	err   error
}

func NewResultRune(v rune, err error) ResultRune {
	if err != nil {
		return ErrRune(err)
	} else {
		return OkRune(v)
	}
}

func NewResultRunePtr(v *rune, err error) ResultRune {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrRune(err)
	} else {
		return OkRune(*v)
	}
}

func OkRune(v rune) ResultRune {
	return ResultRune{value: &v}
}

func ErrRune(err error) ResultRune {
	return ResultRune{err: err}
}

func (r ResultRune) Value() rune {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultRune) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultRune) Get() (rune, error) {
	if r.value == nil {
		var zero rune
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultRune) Present() bool {
	return r.value != nil
}
