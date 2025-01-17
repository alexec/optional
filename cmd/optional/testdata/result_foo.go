// Code generated by optional. DO NOT EDIT.

package testdata

import (
	"errors"
)

// ResultFoo is an result Foo.
type ResultFoo struct {
	value *Foo
	err   error
}

func NewResultFoo(v Foo, err error) ResultFoo {
	if err != nil {
		return ErrFoo(err)
	} else {
		return OkFoo(v)
	}
}

func NewResultFooPtr(v *Foo, err error) ResultFoo {
	if v == nil && err == nil {
		panic("both value and err are nil")
	}
	if err != nil {
		return ErrFoo(err)
	} else {
		return OkFoo(*v)
	}
}

func OkFoo(v Foo) ResultFoo {
	return ResultFoo{value: &v}
}

func ErrFoo(err error) ResultFoo {
	return ResultFoo{err: err}
}

func (r ResultFoo) Value() Foo {
	if r.value == nil {
		panic("value not present")
	}
	return *r.value
}

func (r ResultFoo) Err() error {
	if r.err == nil {
		panic("err not present")
	}
	return r.err
}

func (r ResultFoo) Get() (Foo, error) {
	if r.value == nil {
		var zero Foo
		return zero, errors.New("value not present")
	}
	return *r.value, nil
}

func (r ResultFoo) Present() bool {
	return r.value != nil
}
