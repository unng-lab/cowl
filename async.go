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
	"github.com/unng-lab/cowl/internal/async"
)

// 0 to ...

func Async0to0(fn func()) func() {
	return async.Async0to0(fn)
}

func Async0to1[A any](fn func() A) func() A {
	return async.Async0to1(fn)
}

func Async0to2[A, B any](fn func() (A, B)) func() (A, B) {
	return async.Async0to2(fn)
}

func Async0to3[A, B, C any](fn func() (A, B, C)) func() (A, B, C) {
	return async.Async0to3(fn)
}

// 1 to ...

func Async1to0[A any](fn func(A), a A) func() {
	return async.Async1to0(fn, a)
}

func Async1to1[A, B any](fn func(A) B, a A) func() B {
	return async.Async1to1(fn, a)
}

func Async1to2[A, B, C any](fn func(A) (B, C), a A) func() (B, C) {
	return async.Async1to2(fn, a)
}

func Async1toO[A, B, C, D any](fn func(A) (B, C, D), a A) func() (B, C, D) {
	return async.Async1to3(fn, a)
}

// 2 to ...

func Async2toO[A, B any](fn func(A, B), a A, b B) func() {
	return async.Async2to0(fn, a, b)
}

func Async2to1[A, B, C any](fn func(A, B) C, a A, b B) func() C {
	return async.Async2to1(fn, a, b)
}

func Async2to2[A, B, C, D any](fn func(A, B) (C, D), a A, b B) func() (C, D) {
	return async.Async2to2(fn, a, b)
}

func Async2to3[A, B, C, D, E any](fn func(A, B) (C, D, E), a A, b B) func() (C, D, E) {
	return async.Async2to3(fn, a, b)
}

// 3 to ...

func Async3toO[A, B, C any](fn func(A, B, C), a A, b B, c C) func() {
	return async.Async3to0(fn, a, b, c)
}

func Async3to1[A, B, C, D any](fn func(A, B, C) D, a A, b B, c C) func() D {
	return async.Async3to1(fn, a, b, c)
}

func Async3to2[A, B, C, D, E any](fn func(A, B, C) (D, E), a A, b B, c C) func() (D, E) {
	return async.Async3to2(fn, a, b, c)
}

func Async3to3[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F), a A, b B, c C) func() (D, E, F) {
	return async.Async3to3(fn, a, b, c)
}

// 4 to ...

func Async4toO[A, B, C, D any](fn func(A, B, C, D), a A, b B, c C, d D) func() {
	return async.Async4to0(fn, a, b, c, d)
}

func Async4to1[A, B, C, D, E any](fn func(A, B, C, D) E, a A, b B, c C, d D) func() E {
	return async.Async4to1(fn, a, b, c, d)
}

func Async4to2[A, B, C, D, E, F any](fn func(A, B, C, D) (E, F), a A, b B, c C, d D) func() (E, F) {
	return async.Async4to2(fn, a, b, c, d)
}

func Async4to3[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G), a A, b B, c C, d D) func() (E, F, G) {
	return async.Async4to3(fn, a, b, c, d)
}

// 5 to ...

func Async5toO[A, B, C, D, E any](fn func(A, B, C, D, E), a A, b B, c C, d D, e E) func() {
	return async.Async5to0(fn, a, b, c, d, e)
}

func Async5to1[A, B, C, D, E, F any](fn func(A, B, C, D, E) F, a A, b B, c C, d D, e E) func() F {
	return async.Async5to1(fn, a, b, c, d, e)
}

func Async5to2[A, B, C, D, E, F, G any](fn func(A, B, C, D, E) (F, G), a A, b B, c C, d D, e E) func() (F, G) {
	return async.Async5to2(fn, a, b, c, d, e)
}

func Async5to3[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H), a A, b B, c C, d D, e E) func() (F, G, H) {
	return async.Async5to3(fn, a, b, c, d, e)
}
