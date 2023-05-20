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

func TestDoZZSimple(t *testing.T) {
	var a string
	testFunc := func() {
		time.Sleep(1 * time.Millisecond)
		a = "test"
	}

	f := DoZZ(testFunc)

	time.Sleep(2 * time.Millisecond)

	f()
	if a != "test" {
		t.Error(a)
	}
}

func TestDoZOSimple(t *testing.T) {
	var a string
	testFunc := func() string {
		time.Sleep(1 * time.Millisecond)
		return "test"
	}

	f := DoZO(testFunc)

	time.Sleep(2 * time.Millisecond)

	a = f()
	if a != "test" {
		t.Error(a)
	}
}

func TestDoZTSimple(t *testing.T) {
	var a, b string
	testFunc := func() (string, string) {
		time.Sleep(1 * time.Millisecond)
		return "test", "test1"
	}

	f := DoZT(testFunc)

	time.Sleep(2 * time.Millisecond)

	a, b = f()
	if a != "test" || b != "test1" {
		t.Error(a, b)
	}
}

func TestDoZThSimple(t *testing.T) {
	var a, b, c string
	testFunc := func() (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return "test", "test1", "test2"
	}

	f := DoZTh(testFunc)

	time.Sleep(2 * time.Millisecond)

	a, b, c = f()
	if a != "test" || b != "test1" || c != "test2" {
		t.Error(a)
	}
}
