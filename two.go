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

// twoZero
type twoZero[A, B any] struct {
	m sync.RWMutex
}

// DoTZ two in zero return
func DoTZ[A, B any](
	fn func(A, B),
	a A,
	b B,
) func() {
	var res twoZero[A, B]
	res.m.Lock()
	go func() {
		fn(a, b)
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// twoOne
type twoOne[C any] struct {
	c C
	m sync.RWMutex
}

// DoTO two in one return
func DoTO[A, B, C any](
	fn func(A, B) C,
	a A,
	b B,
) func() C {
	var res twoOne[C]
	res.m.Lock()
	go func() {
		res.c = fn(a, b)
		res.m.Unlock()
	}()
	return func() C {
		res.m.RLock()
		return res.c
	}
}

type twoTwo[C, D any] struct {
	c C
	d D
	m sync.RWMutex
}

// DoTT two in two return
func DoTT[A, B, C, D any](
	fn func(A, B) (C, D),
	a A,
	b B,
) func() (C, D) {
	var res twoTwo[C, D]
	res.m.Lock()
	go func() {
		res.c, res.d = fn(a, b)
		res.m.Unlock()
	}()
	return func() (C, D) {
		res.m.RLock()
		return res.c, res.d
	}
}

type twoThree[C, D, E any] struct {
	c C
	d D
	e E
	m sync.RWMutex
}

// DoTTh two in three return
func DoTTh[A, B, C, D, E any](
	fn func(A, B) (C, D, E),
	a A,
	b B,
) func() (C, D, E) {
	var res twoThree[C, D, E]
	res.m.Lock()
	go func() {
		res.c, res.d, res.e = fn(a, b)
		res.m.Unlock()
	}()
	return func() (C, D, E) {
		res.m.RLock()
		return res.c, res.d, res.e
	}
}
