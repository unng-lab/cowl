package cowl

import (
	"sync"
)

// threeZero
type threeZero[A, B, C any] struct {
	fn func(A, B, C)
	a  *A
	b  *B
	c  *C
	m  sync.RWMutex
}

func (s *threeZero[A, B, C]) Run() {
	s.fn(*s.a, *s.b, *s.c)
	s.m.Unlock()
}

// DoThZ three in zero return
func DoThZ[A, B, C any](
	fn func(A, B, C),
	a A,
	b B,
	c C,
) func() {
	res := threeZero[A, B, C]{}
	res.m.Lock()
	res.fn, res.a, res.b, res.c = fn, &a, &b, &c
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// threeOne
type threeOne[A, B, C, D any] struct {
	fn func(A, B, C) D
	a  *A
	b  *B
	c  *C
	d  *D
	m  sync.RWMutex
}

func (s *threeOne[A, B, C, D]) Run() {
	*s.d = s.fn(*s.a, *s.b, *s.c)
	s.m.Unlock()
}

// DoThO three in one return
func DoThO[A, B, C, D any](
	fn func(A, B, C) D,
	a A,
	b B,
	c C,
) func() D {
	res := threeOne[A, B, C, D]{d: new(D)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c = fn, &a, &b, &c
	send(&res)
	return func() D {
		res.m.RLock()
		return *res.d
	}
}

type threeTwo[A, B, C, D, E any] struct {
	fn func(A, B, C) (D, E)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	m  sync.RWMutex
}

func (s *threeTwo[A, B, C, D, E]) Run() {
	*s.d, *s.e = s.fn(*s.a, *s.b, *s.c)
	s.m.Unlock()
}

// DoThT three in two return
func DoThT[A, B, C, D, E any](
	fn func(A, B, C) (D, E),
	a A,
	b B,
	c C,
) func() (D, E) {
	res := threeTwo[A, B, C, D, E]{d: new(D), e: new(E)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c = fn, &a, &b, &c
	send(&res)
	return func() (D, E) {
		res.m.RLock()
		return *res.d, *res.e
	}
}

type threeThree[A, B, C, D, E, F any] struct {
	fn func(A, B, C) (D, E, F)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	m  sync.RWMutex
}

func (s *threeThree[A, B, C, D, E, F]) Run() {
	*s.d, *s.e, *s.f = s.fn(*s.a, *s.b, *s.c)
	s.m.Unlock()
}

// DoThTh three in three return
func DoThTh[A, B, C, D, E, F any](
	fn func(A, B, C) (D, E, F),
	a A,
	b B,
	c C,
) func() (D, E, F) {
	res := threeThree[A, B, C, D, E, F]{d: new(D), e: new(E), f: new(F)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c = fn, &a, &b, &c
	send(&res)
	return func() (D, E, F) {
		res.m.RLock()
		return *res.d, *res.e, *res.f
	}
}
