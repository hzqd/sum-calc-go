package ext

import (
	"sync"
	"sync/atomic"
)

type Future[T any] struct {
	*future[T]
}

type future[T any] struct {
	runFn   func(func(T))
	wait    sync.WaitGroup
	pending atomic.Bool
	result  T
}

func Async[T any](runFn func(resolve func(T))) Future[T] {
	f := Future[T]{&future[T]{
		runFn:   runFn,
		wait:    sync.WaitGroup{},
		pending: atomic.Bool{},
	}}
	f.wait.Add(1)
	f.pending.Store(true)
	go f.runFn(f.resolve)
	return f
}

func (f Future[T]) Await() T {
	if f.pending.Load() {
		f.wait.Wait()
	}
	return f.result
}

func (f Future[T]) TryGet() (T, bool) {
	return f.result, !f.pending.Load()
}

func (f Future[T]) resolve(value T) {
	if f.pending.Load() {
		f.result = value
		f.wait.Done()
	}
}
