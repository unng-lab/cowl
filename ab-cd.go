package cowl

import (
	"sync"
)

type s[A, B, C, D any] struct {
	f func(A, B) (C, D)
	a A
	b B
	c C
	d D
	m sync.Mutex
}

func (s *s[A, B, C, D]) Run() {
	s.c, s.d = s.f(s.a, s.b)
	s.m.Unlock()
}

func Do[A, B, C, D any](
	f func(A, B) (C, D),
	a A,
	b B,
) func() (C, D) {
	var res s[A, B, C, D]
	res.m.Lock()
	res.f, res.a, res.b = f, a, b
	send(&res)
	return func() (C, D) {
		res.m.Lock()
		res.m.Unlock()
		return res.c, res.d
	}
}
