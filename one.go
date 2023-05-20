package cowl

import (
	"sync"
)

type oneZero[A any] struct {
	fn func(A)
	a  *A
	m  sync.RWMutex
}

func (s *oneZero[A]) Run() {
	s.fn(*s.a)
	s.m.Unlock()
}

// DoOZ one in zero return
func DoOZ[A any](
	fn func(A),
	a A,
) func() {
	res := oneZero[A]{}
	res.m.Lock()
	res.fn, res.a = fn, &a
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// oneOne
type oneOne[A, B any] struct {
	fn func(A) B
	a  *A
	b  *B
	m  sync.RWMutex
}

func (s *oneOne[A, B]) Run() {
	*s.b = s.fn(*s.a)
	s.m.Unlock()
}

// DoOO one in one return
func DoOO[A, B any](
	fn func(A) B,
	a A,
) func() B {
	res := oneOne[A, B]{b: new(B)}
	res.m.Lock()
	res.fn, res.a = fn, &a
	send(&res)
	return func() B {
		res.m.RLock()
		return *res.b
	}
}

// oneTwo
type oneTwo[A, B, C any] struct {
	fn func(A) (B, C)
	a  *A
	b  *B
	c  *C
	m  sync.RWMutex
}

func (s *oneTwo[A, B, C]) Run() {
	*s.b, *s.c = s.fn(*s.a)
	s.m.Unlock()
}

// DoOT one in two return
func DoOT[A, B, C any](
	fn func(A) (B, C),
	a A,
) func() (B, C) {
	res := oneTwo[A, B, C]{b: new(B), c: new(C)}
	res.m.Lock()
	res.fn, res.a = fn, &a
	send(&res)
	return func() (B, C) {
		res.m.RLock()
		return *res.b, *res.c
	}
}

// oneThree
type oneThree[A, B, C, D any] struct {
	fn func(A) (B, C, D)
	a  *A
	b  *B
	c  *C
	d  *D
	m  sync.RWMutex
}

func (s *oneThree[A, B, C, D]) Run() {
	*s.b, *s.c, *s.d = s.fn(*s.a)
	s.m.Unlock()
}

// DoOTh one in three return
func DoOTh[A, B, C, D any](
	fn func(A) (B, C, D),
	a A,
) func() (B, C, D) {
	res := oneThree[A, B, C, D]{b: new(B), c: new(C), d: new(D)}
	res.m.Lock()
	res.fn, res.a = fn, &a
	send(&res)
	return func() (B, C, D) {
		res.m.RLock()
		return *res.b, *res.c, *res.d
	}
}
