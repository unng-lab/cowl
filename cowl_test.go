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
	"strconv"
	"testing"
	"time"
)

func TestDoSimple(t *testing.T) {
	testFunc := func(a string, b int) (string, error) {
		time.Sleep(1 * time.Millisecond)
		return a + strconv.Itoa(b), nil
	}

	f := DoTT(testFunc, "test", 1)

	time.Sleep(2 * time.Millisecond)

	a, b := f()
	if a != "test1" || b != nil {
		t.Error(a, b)
	}
}

func TestDoMustWait(t *testing.T) {
	testFunc := func(a string, b int) (string, error) {
		time.Sleep(1 * time.Millisecond)
		return a + strconv.Itoa(b), nil
	}

	f := DoTT(testFunc, "test", 1)

	a, b := f()
	if a != "test1" || b != nil {
		t.Error(a, b)
	}
}

func TestDoMany(t *testing.T) {
	for i := 0; i < 10; i++ {
		testFunc := func(a string, b int) (string, error) {
			time.Sleep(1 * time.Millisecond)
			return a + strconv.Itoa(b), nil
		}

		f := DoTT(testFunc, "test", 1)

		a, b := f()
		if a != "test1" || b != nil {
			t.Error(a, b)
		}
	}
}
