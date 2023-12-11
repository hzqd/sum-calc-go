package ext

import (
	"sync"
)

type Future[T any] struct {
	*future[T]
}

type future[T any] struct {
	runFn  func() T
	wait   sync.WaitGroup
	result T
}

func Spawn[T any](runFn func() T) Future[T] {
	f := Future[T]{&future[T]{
		runFn: runFn,
		wait:  sync.WaitGroup{},
	}}
	f.wait.Add(1)
	go func() {
		f.result = f.runFn()
		f.runFn = nil
		f.wait.Done()
	}()
	return f
}

func (f Future[T]) Await() T {
	f.wait.Wait()
	return f.result
}

func (f Future[T]) TryGet() (T, bool) {
	return f.result, f.runFn == nil
}
