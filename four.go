package cowl

import (
	"sync"
)

// fourZero
type fourZero[A, B, C, D any] struct {
	fn func(A, B, C, D)
	a  *A
	b  *B
	c  *C
	d  *D
	m  sync.RWMutex
}

func (s *fourZero[A, B, C, D]) Run() {
	s.fn(*s.a, *s.b, *s.c, *s.d)
	s.m.Unlock()
}

// DoFZ four in zero return
func DoFZ[A, B, C, D any](
	fn func(A, B, C, D),
	a A,
	b B,
	c C,
	d D,
) func() {
	res := fourZero[A, B, C, D]{}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d = fn, &a, &b, &c, &d
	send(&res)
	return func() {
		res.m.RLock()
	}
}

// fourOne
type fourOne[A, B, C, D, E any] struct {
	fn func(A, B, C, D) E
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	m  sync.RWMutex
}

func (s *fourOne[A, B, C, D, E]) Run() {
	*s.e = s.fn(*s.a, *s.b, *s.c, *s.d)
	s.m.Unlock()
}

// DoFO four in one return
func DoFO[A, B, C, D, E any](
	fn func(A, B, C, D) E,
	a A,
	b B,
	c C,
	d D,
) func() E {
	res := fourOne[A, B, C, D, E]{e: new(E)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d = fn, &a, &b, &c, &d
	send(&res)
	return func() E {
		res.m.RLock()
		return *res.e
	}
}

type fourTwo[A, B, C, D, E, F any] struct {
	fn func(A, B, C, D) (E, F)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	m  sync.RWMutex
}

func (s *fourTwo[A, B, C, D, E, F]) Run() {
	*s.e, *s.f = s.fn(*s.a, *s.b, *s.c, *s.d)
	s.m.Unlock()
}

// DoFT four in two return
func DoFT[A, B, C, D, E, F any](
	fn func(A, B, C, D) (E, F),
	a A,
	b B,
	c C,
	d D,
) func() (E, F) {
	res := fourTwo[A, B, C, D, E, F]{e: new(E), f: new(F)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d = fn, &a, &b, &c, &d
	send(&res)
	return func() (E, F) {
		res.m.RLock()
		return *res.e, *res.f
	}
}

type fourThree[A, B, C, D, E, F, G any] struct {
	fn func(A, B, C, D) (E, F, G)
	a  *A
	b  *B
	c  *C
	d  *D
	e  *E
	f  *F
	g  *G
	m  sync.RWMutex
}

func (s *fourThree[A, B, C, D, E, F, G]) Run() {
	*s.e, *s.f, *s.g = s.fn(*s.a, *s.b, *s.c, *s.d)
	s.m.Unlock()
}

// DoFTh four in three return
func DoFTh[A, B, C, D, E, F, G any](
	fn func(A, B, C, D) (E, F, G),
	a A,
	b B,
	c C,
	d D,
) func() (E, F, G) {
	res := fourThree[A, B, C, D, E, F, G]{e: new(E), f: new(F), g: new(G)}
	res.m.Lock()
	res.fn, res.a, res.b, res.c, res.d = fn, &a, &b, &c, &d
	send(&res)
	return func() (E, F, G) {
		res.m.RLock()
		return *res.e, *res.f, *res.g
	}
}
