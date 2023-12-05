package ext

type Sender[E any] chan<- E

type Receiver[E any] <-chan E

func Chan_[E any](cap int) (Sender[E], Receiver[E]) {
	ch := make(chan E, cap)
	return ch, ch
}

func (c Sender[E]) Send(e E) {
	c <- e
}

func (c Sender[E]) TrySend(e E) bool {
	select {
	case c <- e:
		return true
	default:
		return false
	}
}

func (c Sender[E]) Len() int {
	return len(c)
}

func (c Sender[E]) Cap() int {
	return cap(c)
}

func (c Sender[E]) Full() bool {
	return len(c) == cap(c)
}

func (c Sender[E]) Close() {
	close(c)
}

func (c Sender[E]) append_(element E) Sender[E] {
	c <- element
	return c
}

func (c Receiver[E]) Recv() (E, bool) {
	e, ok := <-c
	return e, ok
}

func (c Receiver[E]) TryRecv() (E, bool) {
	select {
	case e, ok := <-c:
		return e, ok
	default:
		return *new(E), false
	}
}

func (c Receiver[E]) ForEach(fn func(E)) {
	for e := range c {
		fn(e)
	}
}

func (c Receiver[E]) Len() int {
	return len(c)
}

func (c Receiver[E]) Cap() int {
	return cap(c)
}

func (c Receiver[E]) Empty() bool {
	return len(c) == 0
}
