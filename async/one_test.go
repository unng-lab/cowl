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
	"testing"
	"time"
)

func TestAsync1to0Simple(t *testing.T) {
	var z string
	testFunc := func(a string) {
		time.Sleep(1 * time.Millisecond)
		z = a
	}

	f := Async1to0(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "test" {
		t.Error(z)
	}
}

func TestAsync1to1Simple(t *testing.T) {
	var z string
	testFunc := func(a string) string {
		time.Sleep(1 * time.Millisecond)
		return a
	}

	f := Async1to1(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "test" {
		t.Error(z)
	}
}

func TestAsync1to2Simple(t *testing.T) {
	var z, y string
	testFunc := func(a string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + "1", a + "2"
	}

	f := Async1to2(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "test1" || y != "test2" {
		t.Error(z, y)
	}
}

func TestAsync1to3Simple(t *testing.T) {
	var z, y, x string
	testFunc := func(a string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + "1", a + "2", a + "3"
	}

	f := Async1to3(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "test1" || y != "test2" || x != "test3" {
		t.Error(z, x, y)
	}
}
