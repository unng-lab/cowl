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
	"testing"
	"time"
)

func TestDoFiZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d, e string) {
		time.Sleep(1 * time.Millisecond)
		z = a + b + c + d + e
	}

	f := DoFiZ(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "abcde" {
		t.Error(z)
	}
}

func TestDoFiOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d, e string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e
	}

	f := DoFiO(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "abcde" {
		t.Error(z)
	}
}

func TestDoFiTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b, c, d, e string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e + "1", a + b + c + d + e + "2"
	}

	f := DoFiT(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "abcde1" || y != "abcde2" {
		t.Error(z, y)
	}
}

func TestDoFiFSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b, c, d, e string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e + "1", a + b + c + d + e + "2", a + b + c + d + e + "3"
	}

	f := DoFiTh(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "abcde1" || y != "abcde2" || x != "abcde3" {
		t.Error(z, x, y)
	}
}
