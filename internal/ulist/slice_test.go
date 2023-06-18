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

package ulist

import (
	"testing"
)

type TestStruct struct {
	id   int
	name string
}

func (t *TestStruct) GetData() *TestStruct {
	return t
}

func TestUListAdd(t *testing.T) {
	tt := TestStruct{
		id:   1,
		name: "asd",
	}
	list := New(tt, 10)
	key := list.Add(tt)

	if tt.id != list.Pieces[key].Entity.id || tt.name != list.Pieces[key].Entity.name {
		t.Error(tt)
	}

	if cap(list.Pieces) != 10 {
		t.Error(list)
	}

	if cap(list.Pieces) != 10 {
		t.Error(list)
	}
}
