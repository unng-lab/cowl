package cowl

import (
	"sync"
)

// zeroZero
type zeroZero struct {
	fn func()
	m  sync.RWMutex
}

func (s *zeroZero) Run() {
	s.fn()
	s.m.Unlock()
}

// DoZZ zero in zero return
func DoZZ(
	fn func(),
) func() {
	res := zeroZero{}
	res.m.Lock()
	res.fn = fn
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// zeroOne
type zeroOne[A any] struct {
	fn func() A
	a  *A
	m  sync.RWMutex
}

func (s *zeroOne[A]) Run() {
	*s.a = s.fn()
	s.m.Unlock()
}

// DoZO zero in one return
func DoZO[A any](
	fn func() A,
) func() A {
	res := zeroOne[A]{a: new(A)}
	res.m.Lock()
	res.fn = fn
	send(&res)
	return func() A {
		res.m.RLock()
		return *res.a
	}
}

// zeroTwo
type zeroTwo[A, B any] struct {
	fn func() (A, B)
	a  *A
	b  *B
	m  sync.RWMutex
}

func (s *zeroTwo[A, B]) Run() {
	*s.a, *s.b = s.fn()
	s.m.Unlock()
}

// DoZT zero in two return
func DoZT[A, B any](
	fn func() (A, B),
) func() (A, B) {
	res := zeroTwo[A, B]{a: new(A), b: new(B)}
	res.m.Lock()
	res.fn = fn
	send(&res)
	return func() (A, B) {
		res.m.RLock()
		return *res.a, *res.b
	}
}

// zeroThree
type zeroThree[A, B, C any] struct {
	fn func() (A, B, C)
	a  *A
	b  *B
	c  *C
	m  sync.RWMutex
}

func (s *zeroThree[A, B, C]) Run() {
	*s.a, *s.b, *s.c = s.fn()
	s.m.Unlock()
}

// DoZTh zero in three return
func DoZTh[A, B, C any](
	fn func() (A, B, C),
) func() (A, B, C) {
	res := zeroThree[A, B, C]{a: new(A), b: new(B), c: new(C)}
	res.m.Lock()
	res.fn = fn
	send(&res)
	return func() (A, B, C) {
		res.m.RLock()
		return *res.a, *res.b, *res.c
	}
}
