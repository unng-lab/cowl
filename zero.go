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

// zeroZero
type zeroZero struct {
	m sync.RWMutex
}

// DoZZ zero in zero return
func DoZZ(
	fn func(),
) func() {
	var res zeroZero
	res.m.Lock()
	go func() {
		fn()
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// zeroOne
type zeroOne[A any] struct {
	a A
	m sync.RWMutex
}

// DoZO zero in one return
func DoZO[A any](
	fn func() A,
) func() A {
	var res zeroOne[A]
	res.m.Lock()
	go func() {
		res.a = fn()
		res.m.Unlock()
	}()
	return func() A {
		res.m.RLock()
		return res.a
	}
}

// zeroTwo
type zeroTwo[A, B any] struct {
	a A
	b B
	m sync.RWMutex
}

// DoZT zero in two return
func DoZT[A, B any](
	fn func() (A, B),
) func() (A, B) {
	var res zeroTwo[A, B]
	res.m.Lock()
	go func() {
		res.a, res.b = fn()
		res.m.Unlock()
	}()
	return func() (A, B) {
		res.m.RLock()
		return res.a, res.b
	}
}

// zeroThree
type zeroThree[A, B, C any] struct {
	a A
	b B
	c C
	m sync.RWMutex
}

// DoZTh zero in three return
func DoZTh[A, B, C any](
	fn func() (A, B, C),
) func() (A, B, C) {
	var res zeroThree[A, B, C]
	res.m.Lock()
	go func() {
		res.a, res.b, res.c = fn()
		res.m.Unlock()
	}()
	return func() (A, B, C) {
		res.m.RLock()
		return res.a, res.b, res.c
	}
}
