package ext

/////////////////////////////////////////////////////////////////////////////
// The `Some` and `None` struct for `Option` type.
/////////////////////////////////////////////////////////////////////////////

type Some[T any] struct {
	t T
}

type None struct{}

/////////////////////////////////////////////////////////////////////////////
// The `Option` type.
/////////////////////////////////////////////////////////////////////////////

type Option[T any] interface {
	Some[T] | None
}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func NewSome[T any, O Option[T]](t T) O {
	return any(Some[T]{t}).(O)
}

func NewNone[T any, O Option[T]]() O {
	return any(None{}).(O)
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsSome[T any, O Option[T]](o O) bool {
	var res bool
	switch any(o).(type) {
	case Some[T]:
		res = true
	case None:
		res = false
	}
	return res
}

func IsNone[T any, O Option[T]](o O) bool {
	return !IsSome[T](o)
}

func IsSomeAnd[T any, O Option[T]](o O, f func(T) bool) bool {
	var res bool
	switch x := any(o).(type) {
	case Some[T]:
		res = f(x.t)
	case None:
		res = false
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Getting to contained values
/////////////////////////////////////////////////////////////////////////

func UnwrapOption[T any, O Option[T]](o O) T {
	var res T
	switch x := any(o).(type) {
	case Some[T]:
		res = x.t
	}
	return res
}

func MapSome[T, U any, I Option[T], O Option[U]](in I, fn func(T) U) O {
	var res O
	switch x := any(in).(type) {
	case Some[T]:
		res = NewSome[U, O](fn(x.t))
	case None:
		res = NewNone[U, O]()
	}
	return res
}

func OkOrElse[T, E any, O Option[T], R Result[T, E]](o O, fail func() E) R {
	var res R
	switch x := any(o).(type) {
	case Some[T]:
		res = NewOk[T, E, R](x.t)
	case None:
		res = NewErr[T, E, R](fail())
	}
	return res
}
