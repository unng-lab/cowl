package cowl

import (
	"sync"
)

// twoZero
type twoZero[A, B any] struct {
	fn func(A, B)
	a  *A
	b  *B
	m  sync.RWMutex
}

func (s *twoZero[A, B]) Run() {
	s.fn(*s.a, *s.b)
	s.m.Unlock()
}

// DoTZ two in zero return
func DoTZ[A, B any](
	fn func(A, B),
	a A,
	b B,
) func() {
	res := twoZero[A, B]{}
	res.m.Lock()
	res.fn, res.a, res.b = fn, &a, &b
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// twoOne
type twoOne[A, B, C any] struct {
	fn func(A, B) C
	a  *A
	b  *B
	c  *C
	m  sync.RWMutex
}

func (s *twoOne[A, B, C]) Run() {
	*s.c = s.fn(*s.a, *s.b)
	s.m.Unlock()
}

// DoTO two in one return
func DoTO[A, B, C any](
	fn func(A, B) C,
	a A,
	b B,
) func() C {
	res := twoOne[A, B, C]{c: new(C)}
	res.m.Lock()
	res.fn, res.a, res.b = fn, &a, &b
	send(&res)
	return func() C {
		res.m.RLock()
		return *res.c
	}
}

type twoTwo[A, B, C, D any] struct {
	fn func(A, B) (C, D)
	a  *A
	b  *B
	c  *C
	d  *D
	m  sync.RWMutex
}

func (s *twoTwo[A, B, C, D]) Run() {
	*s.c, *s.d = s.fn(*s.a, *s.b)
	s.m.Unlock()
}

// DoTT two in two return
func DoTT[A, B, C, D any](
	fn func(A, B) (C, D),
	a A,
	b B,
) func() (C, D) {
	res := twoTwo[A, B, C, D]{c: new(C), d: new(D)}
	res.m.Lock()
	res.fn, res.a, res.b = fn, &a, &b
	send(&res)
	return func() (C, D) {
		res.m.RLock()
		return *res.c, *res.d
	}
}

type twoThree[A, B, C, D, E any] struct {
	fn func(A, B) (C, D, E)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	m  sync.RWMutex
}

func (s *twoThree[A, B, C, D, E]) Run() {
	*s.c, *s.d, *s.e = s.fn(*s.a, *s.b)
	s.m.Unlock()
}

// DoTTh two in three return
func DoTTh[A, B, C, D, E any](
	fn func(A, B) (C, D, E),
	a A,
	b B,
) func() (C, D, E) {
	res := twoThree[A, B, C, D, E]{c: new(C), d: new(D), e: new(E)}
	res.m.Lock()
	res.fn, res.a, res.b = fn, &a, &b
	send(&res)
	return func() (C, D, E) {
		res.m.RLock()
		return *res.c, *res.d, *res.e
	}
}
