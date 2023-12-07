package ext

type Opt[T any] struct {
	Value T
	Has   bool
}
