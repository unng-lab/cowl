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

package cowl

import (
	"sync"
)

// threeZero
type threeZero[A, B, C any] struct {
	m sync.RWMutex
}

// DoThZ three in zero return
func DoThZ[A, B, C any](
	fn func(A, B, C),
	a A,
	b B,
	c C,
) func() {
	var res threeZero[A, B, C]
	res.m.Lock()
	go func() {
		fn(a, b, c)
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// threeOne
type threeOne[B, C, D any] struct {
	b B
	c C
	d D
	m sync.RWMutex
}

// DoThO three in one return
func DoThO[A, B, C, D any](
	fn func(A, B, C) D,
	a A,
	b B,
	c C,
) func() D {
	var res threeOne[B, C, D]
	res.m.Lock()
	go func() {
		res.d = fn(a, b, c)
		res.m.Unlock()
	}()
	return func() D {
		res.m.RLock()
		return res.d
	}
}

type threeTwo[D, E any] struct {
	d D
	e E
	m sync.RWMutex
}

// DoThT three in two return
func DoThT[A, B, C, D, E any](
	fn func(A, B, C) (D, E),
	a A,
	b B,
	c C,
) func() (D, E) {
	var res threeTwo[D, E]
	res.m.Lock()
	go func() {
		res.d, res.e = fn(a, b, c)
		res.m.Unlock()
	}()
	return func() (D, E) {
		res.m.RLock()
		return res.d, res.e
	}
}

type threeThree[D, E, F any] struct {
	d D
	e E
	f F
	m sync.RWMutex
}

// DoThTh three in three return
func DoThTh[A, B, C, D, E, F any](
	fn func(A, B, C) (D, E, F),
	a A,
	b B,
	c C,
) func() (D, E, F) {
	var res threeThree[D, E, F]
	res.m.Lock()
	go func() {
		res.d, res.e, res.f = fn(a, b, c)
		res.m.Unlock()
	}()
	return func() (D, E, F) {
		res.m.RLock()
		return res.d, res.e, res.f
	}
}
