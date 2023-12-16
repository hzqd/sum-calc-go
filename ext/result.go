package ext

/////////////////////////////////////////////////////////////////////////////
// Error handling with the `Ok` and `Err` struct for `Result` type.
/////////////////////////////////////////////////////////////////////////////

type Ok[T any] struct {
	t T
}

type Err[E any] struct {
	e E
}

/////////////////////////////////////////////////////////////////////////////
// Error handling with the `Result` type.
/////////////////////////////////////////////////////////////////////////////

type Result[T, E any] interface {
	Ok[T] | Err[E]
}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func NewOk[T, E any, R Result[T, E]](t T) R {
	return any(Ok[T]{t}).(R)
}

func NewErr[T, E any, R Result[T, E]](e E) R {
	return any(Err[E]{e}).(R)
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsOk[T, E any, R Result[T, E]](r R) bool {
	var res bool
	switch any(r).(type) {
	case Ok[T]:
		res = true
	case Err[E]:
		res = false
	}
	return res
}

func IsErr[T, E any, R Result[T, E]](r R) bool {
	return !IsOk[T, E](r)
}

func isOkAnd[T, E any, R Result[T, E]](r R, f func(T) bool) bool {
	var res bool
	switch x := any(r).(type) {
	case Ok[T]:
		res = f(x.t)
	case Err[E]:
		res = false
	}
	return res
}

func isErrAnd[T, E any, R Result[T, E]](r R, f func(E) bool) bool {
	var res bool
	switch x := any(r).(type) {
	case Err[E]:
		res = f(x.e)
	case Ok[T]:
		res = false
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Adapter for each variant
/////////////////////////////////////////////////////////////////////////

func OkToOption[T, E any, R Result[T, E], O Option[T]](r R) O {
	var res O
	switch x := any(r).(type) {
	case Ok[T]:
		res = NewSome[T, O](x.t)
	case Err[E]:
		res = NewNone[T, O]()
	}
	return res
}

func ErrToOption[T, E any, R Result[T, E], O Option[E]](r R) O {
	var res O
	switch x := any(r).(type) {
	case Ok[T]:
		res = NewNone[E, O]()
	case Err[E]:
		res = NewSome[E, O](x.e)
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Transforming contained values
/////////////////////////////////////////////////////////////////////////

func MapOk[T, U, E any, Rin Result[T, E], Rout Result[U, E]](r Rin, f func(T) U) Rout {
	var res Rout
	switch x := any(r).(type) {
	case Ok[T]:
		res = NewOk[U, E, Rout](f(x.t))
	case Err[E]:
		res = NewErr[U, E, Rout](x.e)
	}
	return res
}

func MapErr[T, F, E any, Rin Result[T, E], Rout Result[T, F]](r Rin, f func(E) F) Rout {
	var res Rout
	switch x := any(r).(type) {
	case Ok[T]:
		res = NewOk[T, F, Rout](x.t)
	case Err[E]:
		res = NewErr[T, F, Rout](f(x.e))
	}
	return res
}

func MapOkOrElse[T, U, E any, Rin Result[T, E]](r Rin, fail func(E) U, succ func(T) U) U {
	var res U
	switch x := any(r).(type) {
	case Ok[T]:
		res = succ(x.t)
	case Err[E]:
		res = fail(x.e)
	}
	return res
}

func UnwrapOk[T, E any, R Result[T, E]](r R) T {
	var res T
	switch x := any(r).(type) {
	case Ok[T]:
		res = x.t
	}
	return res
}

func UnwrapErr[T, E any, R Result[T, E]](r R) E {
	var res E
	switch x := any(r).(type) {
	case Err[E]:
		res = x.e
	}
	return res
}
