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

// fourZero
type fourZero struct {
	m sync.RWMutex
}

// DoFZ four in zero return
func DoFZ[A, B, C, D any](
	fn func(A, B, C, D),
	a A,
	b B,
	c C,
	d D,
) func() {
	var res fourZero
	res.m.Lock()
	go func() {
		fn(a, b, c, d)
		res.m.Unlock()
	}()
	return func() {
		res.m.RLock()
	}
}

// fourOne
type fourOne[E any] struct {
	e E
	m sync.RWMutex
}

// DoFO four in one return
func DoFO[A, B, C, D, E any](
	fn func(A, B, C, D) E,
	a A,
	b B,
	c C,
	d D,
) func() E {
	var res fourOne[E]
	res.m.Lock()
	go func() {
		res.e = fn(a, b, c, d)
		res.m.Unlock()
	}()
	return func() E {
		res.m.RLock()
		return res.e
	}
}

type fourTwo[E, F any] struct {
	e E
	f F
	m sync.RWMutex
}

// DoFT four in two return
func DoFT[A, B, C, D, E, F any](
	fn func(A, B, C, D) (E, F),
	a A,
	b B,
	c C,
	d D,
) func() (E, F) {
	var res fourTwo[E, F]
	res.m.Lock()
	go func() {
		res.e, res.f = fn(a, b, c, d)
		res.m.Unlock()
	}()
	return func() (E, F) {
		res.m.RLock()
		return res.e, res.f
	}
}

type fourThree[E, F, G any] struct {
	e E
	f F
	g G
	m sync.RWMutex
}

// DoFTh four in three return
func DoFTh[A, B, C, D, E, F, G any](
	fn func(A, B, C, D) (E, F, G),
	a A,
	b B,
	c C,
	d D,
) func() (E, F, G) {
	var res fourThree[E, F, G]
	res.m.Lock()
	go func() {
		res.e, res.f, res.g = fn(a, b, c, d)
		res.m.Unlock()
	}()
	return func() (E, F, G) {
		res.m.RLock()
		return res.e, res.f, res.g
	}
}
