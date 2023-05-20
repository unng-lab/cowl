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

func TestDoFZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d string) {
		time.Sleep(1 * time.Millisecond)
		z = a + b + c + d
	}

	f := DoFZ(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "abcd" {
		t.Error(z)
	}
}

func TestDoFOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d
	}

	f := DoFO(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "abcd" {
		t.Error(z)
	}
}

func TestDoFTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b, c, d string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + "1", a + b + c + d + "2"
	}

	f := DoFT(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "abcd1" || y != "abcd2" {
		t.Error(z, y)
	}
}

func TestDoFFSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b, c, d string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + "1", a + b + c + d + "2", a + b + c + d + "3"
	}

	f := DoFTh(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "abcd1" || y != "abcd2" || x != "abcd3" {
		t.Error(z, x, y)
	}
}
