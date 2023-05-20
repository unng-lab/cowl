package cowl

import (
	"sync"
)

// fiveZero
type fiveZero[A, B, C, D, E any] struct {
	fn func(A, B, C, D, E)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	m  sync.RWMutex
}

func (s *fiveZero[A, B, C, D, E]) Run() {
	s.fn(*s.a, *s.b, *s.c, *s.d, *s.e)
	s.m.Unlock()
}

// DoFiZ five in zero return
func DoFiZ[A, B, C, D, E any](
	fn func(A, B, C, D, E),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() {
	res := fiveZero[A, B, C, D, E]{}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d, res.e = fn, &a, &b, &c, &d, &e
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// fiveOne
type fiveOne[A, B, C, D, E, F any] struct {
	fn func(A, B, C, D, E) F
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	m  sync.RWMutex
}

func (s *fiveOne[A, B, C, D, E, F]) Run() {
	*s.f = s.fn(*s.a, *s.b, *s.c, *s.d, *s.e)
	s.m.Unlock()
}

// DoFiO five in one return
func DoFiO[A, B, C, D, E, F any](
	fn func(A, B, C, D, E) F,
	a A,
	b B,
	c C,
	d D,
	e E,
) func() F {
	res := fiveOne[A, B, C, D, E, F]{f: new(F)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d, res.e = fn, &a, &b, &c, &d, &e
	send(&res)
	return func() F {
		res.m.RLock()
		return *res.f
	}
}

type fiveTwo[A, B, C, D, E, F, G any] struct {
	fn func(A, B, C, D, E) (F, G)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	g  *G
	m  sync.RWMutex
}

func (s *fiveTwo[A, B, C, D, E, F, G]) Run() {
	*s.f, *s.g = s.fn(*s.a, *s.b, *s.c, *s.d, *s.e)
	s.m.Unlock()
}

// DoFiT five in two return
func DoFiT[A, B, C, D, E, F, G any](
	fn func(A, B, C, D, E) (F, G),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() (F, G) {
	res := fiveTwo[A, B, C, D, E, F, G]{f: new(F), g: new(G)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d, res.e = fn, &a, &b, &c, &d, &e
	send(&res)
	return func() (F, G) {
		res.m.RLock()
		return *res.f, *res.g
	}
}

type fiveThree[A, B, C, D, E, F, G, H any] struct {
	fn func(A, B, C, D, E) (F, G, H)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	g  *G
	h  *H
	m  sync.RWMutex
}

func (s *fiveThree[A, B, C, D, E, F, G, H]) Run() {
	*s.f, *s.g, *s.h = s.fn(*s.a, *s.b, *s.c, *s.d, *s.e)
	s.m.Unlock()
}

// DoFiTh five in three return
func DoFiTh[A, B, C, D, E, F, G, H any](
	fn func(A, B, C, D, E) (F, G, H),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() (F, G, H) {
	res := fiveThree[A, B, C, D, E, F, G, H]{f: new(F), g: new(G), h: new(H)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d, res.e = fn, &a, &b, &c, &d, &e
	send(&res)
	return func() (F, G, H) {
		res.m.RLock()
		return *res.f, *res.g, *res.h
	}
}
