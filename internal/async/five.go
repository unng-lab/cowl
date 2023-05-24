/*
 * Copyright (c) 2023. UNNG-Lab
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
 *  subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 *  substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 *  INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR
 *  A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 *  COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 *  ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH
 *  THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package async

import (
	"sync"
)

// fiveZero
type fiveZero struct {
	m sync.RWMutex
}

// Async5to0 five in zero return
func Async5to0[A, B, C, D, E any](
	fn func(A, B, C, D, E),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() {
	var res fiveZero
	res.m.Lock()
	go func() {
		fn(a, b, c, d, e)
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// fiveOne
type fiveOne[F any] struct {
	f F
	m sync.RWMutex
}

// Async5to1 five in one return
func Async5to1[A, B, C, D, E, F any](
	fn func(A, B, C, D, E) F,
	a A,
	b B,
	c C,
	d D,
	e E,
) func() F {
	var res fiveOne[F]
	res.m.Lock()
	go func() {
		res.f = fn(a, b, c, d, e)
		res.m.Unlock()
	}()
	return func() F {
		res.m.RLock()
		return res.f
	}
}

type fiveTwo[F, G any] struct {
	f F
	g G
	m sync.RWMutex
}

// Async5to2 five in two return
func Async5to2[A, B, C, D, E, F, G any](
	fn func(A, B, C, D, E) (F, G),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() (F, G) {
	var res fiveTwo[F, G]
	res.m.Lock()
	go func() {
		res.f, res.g = fn(a, b, c, d, e)
		res.m.Unlock()
	}()
	return func() (F, G) {
		res.m.RLock()
		return res.f, res.g
	}
}

type fiveThree[F, G, H any] struct {
	f F
	g G
	h H
	m sync.RWMutex
}

// Async5to3 five in three return
func Async5to3[A, B, C, D, E, F, G, H any](
	fn func(A, B, C, D, E) (F, G, H),
	a A,
	b B,
	c C,
	d D,
	e E,
) func() (F, G, H) {
	var res fiveThree[F, G, H]
	res.m.Lock()
	go func() {
		res.f, res.g, res.h = fn(a, b, c, d, e)
		res.m.Unlock()
	}()
	return func() (F, G, H) {
		res.m.RLock()
		return res.f, res.g, res.h
	}
}
