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

package fifo

import (
	"testing"
)

func TestGetPut(t *testing.T) {
	st := NewStore("", "", 10)
	st.Put("key1", "value1")
	v, err := st.Get("key1")
	if err != nil {
		t.Error(st)
	}
	if v == nil || *v != "value1" {
		t.Error(st)
	}
}

func TestFiFo(t *testing.T) {
	st := NewStore("", "", 3)
	st.Put("key1", "value1")
	st.Put("key2", "value2")
	st.Put("key3", "value3")
	st.Put("key4", "value4")
	_, err := st.Get("key1")
	if err == nil {
		t.Error(st)
	}
	v2, err := st.Get("key2")
	if err != nil {
		t.Error(st)
	}
	v3, err := st.Get("key3")
	if err != nil {
		t.Error(st)
	}
	v4, err := st.Get("key4")
	if err != nil {
		t.Error(st)
	}
	if v2 == nil || *v2 != "value2" {
		t.Error(st)
	}
	if v3 == nil || *v3 != "value3" {
		t.Error(st)
	}
	if v4 == nil || *v4 != "value4" {
		t.Error(st)
	}

	if len(st.next) != 3 {
		t.Error(st)
	}

	if len(st.list) != 3 {
		t.Error(st)
	}

}
