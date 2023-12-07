package ext

import (
	"fmt"
)

type Range[T Number] struct {
	start T
	end   T
	step  T
}

type RangeTo[T Number] struct {
	start T
	end   T
	step  T
}

func Until[T Number](start T, end T, step ...T) Range[T] {
	return Range[T]{start, end,
		obtainStep[T](start, end, step...)}
}

func To[T Number](start T, end T, step ...T) RangeTo[T] {
	return RangeTo[T]{start, end,
		obtainStep[T](start, end, step...)}
}

func (r Range[T]) ForEach(fn func(T)) {
	if !r.Empty() {
		if r.step > 0 {
			for i := r.start; i < r.end; i += r.step {
				fn(i)
			}
		} else if r.step < 0 {
			for i := r.start; i > r.end; i += r.step {
				fn(i)
			}
		}
	}
}

func (r Range[T]) Len() int {
	diff := r.end - r.start
	len_ := int(diff / r.step)
	if diff-r.step*T(len_) != 0 {
		len_ += 1
	}
	return len_
}

func (r Range[T]) Empty() bool {
	return r.Len() == 0
}

func (r Range[T]) ToVec() Vec[T] {
	vec := Vec_[T](r.Len())
	r.ForEach(func(t T) {
		vec.Append(t)
	})
	return vec
}

func (r Range[T]) String() string {
	if r.step == 1 {
		return fmt.Sprintf("range[%v..%v]", r.start, r.end)
	}
	return fmt.Sprintf("range[%v..%v:%v]", r.start, r.end, r.step)
}

func (r RangeTo[T]) ForEach(fn func(T)) {
	if r.Empty() {
		return
	}
	if r.step > 0 {
		for i := r.start; i <= r.end; i += r.step {
			fn(i)
		}
	} else if r.step < 0 {
		for i := r.start; i >= r.end; i += r.step {
			fn(i)
		}
	}
}

func (r RangeTo[T]) Len() int {
	return int((r.end-r.start)/r.step) + 1
}

func (r RangeTo[T]) Empty() bool {
	return r.Len() == 0
}

func (r RangeTo[T]) ToVec() Vec[T] {
	vec := Vec_[T](r.Len())
	r.ForEach(func(t T) {
		vec.Append(t)
	})
	return vec
}

func (r RangeTo[T]) String() string {
	if r.step == 1 {
		return fmt.Sprintf("range[%v..=%v]", r.start, r.end)
	}
	return fmt.Sprintf("range[%v..=%v:%v]", r.start, r.end, r.step)
}

func obtainStep[T Number](start T, end T, step_ ...T) T {
	step := T(1)
	if len(step_) > 0 {
		step = step_[0]
	}
	diff := end - start
	if step == 0 || diff*step < 0 {
		panic("bad range")
	}
	return step
}
