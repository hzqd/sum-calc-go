package ext

type ForEach[E any] interface {
	ForEach(func(E))
	Len() int
	Empty() bool
}

type append_[E, ES any] interface {
	append_(E) ES
}

// Map 将Vec[T]转成Vec[R]
func Map[T, R any](vec Vec[T], fn func(T) R) Vec[R] {
	rs := Vec_[R](vec.Len())
	for _, t := range vec {
		rs.Append(fn(t))
	}
	return rs
}

func MapTo[TS ForEach[T], RS append_[R, RS], T, R any](
	ts TS, fn func(T) R, toFn func(int) RS) RS {
	rs := toFn(ts.Len())
	ts.ForEach(func(t T) {
		rs = rs.append_(fn(t))
	})
	return rs
}

// Flatten 将Vec[Vec[T]]转成Vec[T]
func Flatten[T any](vec Vec[Vec[T]]) Vec[T] {
	len_ := 0
	for _, v := range vec {
		len_ += v.Len()
	}
	rs := VecInit[T](len_)
	i, j := 0, 0
	for _, v := range vec {
		i, j = j, j+v.Len()
		copy(rs[i:j], v)
	}
	return rs
}

// FlatMap 将 Vec[Vec[T]] 平铺成 Vec[T]
func FlatMap[T, R any](vec Vec[T], fn func(T) Vec[R]) Vec[R] {
	rs := Vec_[Vec[R]](vec.Len())
	for _, t := range vec {
		rs.Append(fn(t))
	}
	return Flatten(rs)
}

// Filter 过滤Vec[T]中不需要的元素
func Filter[T any](vec Vec[T], fn func(T) bool) Vec[T] {
	rs := Vec_[T](vec.Len())
	for _, t := range vec {
		if fn(t) {
			rs.Append(t)
		}
	}
	return rs
}

func FilterTo[TS ForEach[T], RS append_[T, RS], T any](
	ts TS, fn func(T) bool, toFn func(int) RS) RS {
	rs := toFn(ts.Len())
	ts.ForEach(func(t T) {
		if fn(t) {
			rs = rs.append_(t)
		}
	})
	return rs
}

// FilterMap 将Vec[T]转成Vec[R] 并过滤不需要的元素
func FilterMap[T, R any](vec Vec[T], fn func(T) (R, bool)) Vec[R] {
	rs := Vec_[R](vec.Len())
	for _, t := range vec {
		if r, b := fn(t); b {
			rs.Append(r)
		}
	}
	return rs
}

func FilterMapTo[TS ForEach[T], RS append_[R, RS], T, R any](ts TS, fn func(T) (R, bool), toFn func(int) RS) RS {
	rs := toFn(ts.Len())
	ts.ForEach(func(t T) {
		if r, b := fn(t); b {
			rs = rs.append_(r)
		}
	})
	return rs
}

// Fold 对 Vec[T] 做合并操作 需要一个初始值
func Fold[TS ForEach[T], T, R any](ts TS, init R, fn func(R, T) R) R {
	ts.ForEach(func(t T) {
		init = fn(init, t)
	})
	return init
}

// RevFold 对 Vec[T] 做反向合并操作 需要一个初始值
func RevFold[T, R any](vec Vec[T], init R, fn func(R, T) R) R {
	for i := vec.Len() - 1; i >= 0; i-- {
		init = fn(init, vec[i])
	}
	return init
}

func FoldDefault[TS ForEach[T], T, R any](ts TS, fn func(R, T) R) R {
	return Fold(ts, *new(R), fn)
}

func RevFoldDefault[T, R any](vec Vec[T], fn func(R, T) R) R {
	return RevFold(vec, *new(R), fn)
}

// ReduceError 定义 Reduce 的错误类型
type ReduceError struct {
	Message string
}

// 为 Reduce 实现 error 接口
func (e *ReduceError) Error() string {
	return e.Message
}

// Reduce : unsafe func, error <-> panic
func Reduce[T any](vec Vec[T], fn func(T, T) T) (T, error) {
	var err error
	var acc T
	size := vec.Len()
	if size == 0 {
		err = &ReduceError{"ReduceError: no first element (index out of range '[0]')"}
		return acc, err
	}
	acc = vec[0]
	for i := 1; i < size; i++ {
		acc = fn(acc, vec[i])
	}
	return acc, err
}

// RevReduce : unsafe func, error <-> panic
func RevReduce[T any](vec Vec[T], fn func(T, T) T) (T, error) {
	var err error
	var acc T
	size := vec.Len()
	if size == 0 {
		err = &ReduceError{"ReduceError: no first element (index out of range '[0]')"}
		return acc, err
	}
	acc = vec[size-1]
	for i := size - 2; i >= 0; i-- {
		acc = fn(acc, vec[i])
	}
	return acc, err
}

// Any 判断切片中的值是否符合预期，一个符合为true
func Any[T any](vec Vec[T], fn func(T) bool) bool {
	for _, t := range vec {
		if fn(t) {
			return true
		}
	}
	return false
}

// All 判断切片中的值是否符合预期，全部符合为true
func All[T any](vec Vec[T], fn func(T) bool) bool {
	for _, t := range vec {
		if !fn(t) {
			return false
		}
	}
	return true
}

// ToDict 分组函数 可以对key映射
func ToDict[K comparable, T any](vec Vec[T], kFn func(T) K) Dict[K, T] {
	dict := Dict_[K, T](4)
	for _, t := range vec {
		k := kFn(t)
		dict.Store(k, t)
	}
	return dict
}

// VToDict 分组函数 可以对key和value映射
func VToDict[K comparable, V, T any](vec Vec[T], kvFn func(T) (K, V)) Dict[K, V] {
	dict := Dict_[K, V](4)
	for _, t := range vec {
		k, v := kvFn(t)
		dict.Store(k, v)
	}
	return dict
}

// GroupBy 分组函数 可以对key映射
func GroupBy[K comparable, T any](vec Vec[T], kFn func(T) K) Dict[K, *Vec[T]] {
	dict := Dict_[K, *Vec[T]](4)
	for _, t := range vec {
		k := kFn(t)
		vs, b := dict[k]
		if !b {
			vs = new(Vec[T])
			vs.Grow(4)
			dict[k] = vs
		}
		vs.Append(t)
	}
	return dict
}

// VGroupBy 分组函数 可以对key和value同时映射
func VGroupBy[K comparable, V, T any](vec Vec[T], kvFn func(T) (K, V)) Dict[K, *Vec[V]] {
	dict := Dict_[K, *Vec[V]](4)
	for _, t := range vec {
		k, v := kvFn(t)
		vs, b := dict[k]
		if !b {
			vs = new(Vec[V])
			vs.Grow(4)
			dict[k] = vs
		}
		vs.Append(v)
	}
	return dict
}

// FollowSort 跟随排序
func FollowSort[O comparable, T any](orders Vec[O], vec Vec[T], kFn func(T) O) Vec[T] {
	return FilterMap(orders, ToDict(vec, kFn).Load)
}
