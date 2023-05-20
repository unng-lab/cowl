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

type oneZero struct {
	m sync.RWMutex
}

// DoOZ one in zero return
func DoOZ[A any](
	fn func(A),
	a A,
) func() {
	var res oneZero
	res.m.Lock()
	go func() {
		fn(a)
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// oneOne
type oneOne[B any] struct {
	b B
	m sync.RWMutex
}

// DoOO one in one return
func DoOO[A, B any](
	fn func(A) B,
	a A,
) func() B {
	var res oneOne[B]
	res.m.Lock()
	go func() {
		res.b = fn(a)
		res.m.Unlock()
	}()
	return func() B {
		res.m.RLock()
		return res.b
	}
}

// oneTwo
type oneTwo[B, C any] struct {
	b B
	c C
	m sync.RWMutex
}

// DoOT one in two return
func DoOT[A, B, C any](
	fn func(A) (B, C),
	a A,
) func() (B, C) {
	var res oneTwo[B, C]
	res.m.Lock()
	go func() {
		res.b, res.c = fn(a)
		res.m.Unlock()
	}()
	return func() (B, C) {
		res.m.RLock()
		return res.b, res.c
	}
}

// oneThree
type oneThree[B, C, D any] struct {
	b B
	c C
	d D
	m sync.RWMutex
}

// DoOTh one in three return
func DoOTh[A, B, C, D any](
	fn func(A) (B, C, D),
	a A,
) func() (B, C, D) {
	var res oneThree[B, C, D]
	res.m.Lock()
	go func() {
		res.b, res.c, res.d = fn(a)
		res.m.Unlock()
	}()
	return func() (B, C, D) {
		res.m.RLock()
		return res.b, res.c, res.d
	}
}
