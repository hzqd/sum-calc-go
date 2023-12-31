package ext

import (
	"slices"
)

type Vec[E any] []E

func Vec_[E any](cap int) Vec[E] {
	return make([]E, 0, cap)
}

func VecGen[E any](len_ int, fn_ ...func(int) E) Vec[E] {
	if len(fn_) > 0 {
		vec := Vec_[E](len_)
		fn := fn_[0]
		for i := 0; i < len_; i++ {
			vec.Append(fn(i))
		}
		return vec
	} else {
		return make([]E, len_)
	}
}

func VecOf[E any](es ...E) Vec[E] {
	return es
}

func (v Vec[E]) ForEach(fn func(E)) {
	for _, e := range v {
		fn(e)
	}
}

func (v Vec[E]) Len() int {
	return len(v)
}

func (v Vec[E]) Cap() int {
	return cap(v)
}

func (v Vec[E]) Empty() bool {
	return len(v) == 0
}

func (v Vec[E]) Get(index int) E {
	return v[index]
}

func (v Vec[E]) GetOr(index int, or E) E {
	if index < len(v) {
		return v[index]
	}
	return or
}

func (v Vec[E]) GetElse(index int, fn func() E) E {
	if index < len(v) {
		return v[index]
	}
	return fn()
}

func (v Vec[E]) Reverse() {
	slices.Reverse(v)
}

func (v *Vec[E]) Append(element E) {
	*v = append(*v, element)
}

func (v *Vec[E]) Appends(elements ...E) {
	*v = append(*v, elements...)
}

func (v *Vec[E]) Insert(index int, elements ...E) {
	*v = slices.Insert(*v, index, elements...)
}

func (v *Vec[E]) RemoveAt(index int) {
	*v = slices.Delete(*v, index, index+1)
}

func (v *Vec[E]) RemoveRange(start, end int) {
	*v = slices.Delete(*v, start, end)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func (v *Vec[E]) Grow(n int) {
	*v = slices.Grow(*v, n)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func (v *Vec[E]) Clip() {
	*v = slices.Clip(*v)
}

func (v Vec[E]) Rev() RevVec[E] {
	return RevVec[E]{v}
}

func (v Vec[E]) append_(element E) Vec[E] {
	return append(v, element)
}

type RevVec[E any] struct {
	Vec[E]
}

func (v RevVec[E]) ForEach(fn func(E)) {
	for i := v.Len() - 1; i >= 0; i-- {
		fn(v.Vec[i])
	}
}

func (v RevVec[E]) Get(index int) E {
	return v.Vec.Get(v.Len() - index - 1)
}
