package ext

import "errors"

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Ok() T
	Err() error
}

func Ok[T any](t T) Result[T] {
	return ok[T]{t}
}

func Err[T any](e error) Result[T] {
	return err[T]{e}
}

type ok[T any] struct {
	t T
}

func (ok[T]) IsOk() bool {
	return true
}

func (ok[T]) IsErr() bool {
	return false
}

func (o ok[T]) Ok() T {
	return o.t
}

func (o ok[T]) Err() error {
	panic(errors.New("result is ok"))
}

type err[T any] struct {
	e error
}

func (err[T]) IsOk() bool {
	return false
}

func (err[T]) IsErr() bool {
	return true
}

func (e err[T]) Ok() T {
	panic(e.e)
}

func (e err[T]) Err() error {
	return e.e
}
