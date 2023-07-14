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

package mslice

import (
	"reflect"
	"testing"

	"github.com/samber/lo"
)

type TestStruct struct {
	id   int
	name string
}

func (t *TestStruct) GetData() *TestStruct {
	return t
}

func TestListAdd(t *testing.T) {
	tt := TestStruct{
		id:   1,
		name: "asd",
	}
	list := New(tt, 10)
	key := list.Add(tt)

	el, err := list.Get(key)
	if err != nil {
		t.Error(err)
	}

	if *el != tt {
		t.Error(tt)
	}

	if len(list.pieces) != 1 {
		t.Error(list)
	}

	if cap(list.pieces) != 10 {
		t.Error(list)
	}

	if cap(list.pieces) != 10 {
		t.Error(list)
	}
}
func TestNewList(t *testing.T) {
	tt := TestStruct{
		id:   1,
		name: "asd",
	}
	list := New(tt, 10)

	if len(list.pieces) > 0 {
		t.Error(list)
	}

	if cap(list.pieces) != 10 {
		t.Error(list)
	}

	if reflect.TypeOf(list.pieces).Name() != reflect.TypeOf([]TestStruct{}).Name() {
		t.Errorf("type error got = %v, want %v", reflect.TypeOf([]TestStruct{}), reflect.TypeOf(list.pieces))
	}
}

func TestListRemove(t *testing.T) {
	tt := TestStruct{
		id:   1,
		name: "asd",
	}
	list := New(tt, 10)
	list.Add(tt)
	key2 := list.Add(tt)
	list.Add(tt)

	if len(list.pieces) != 3 {
		t.Error(list)
	}

	if err := list.Remove(key2); err != nil {
		t.Error(err)
	}

	if len(list.pieces) != 3 {
		t.Error(list)
	}

	if lo.IndexOf(list.dList, key2) == -1 {
		t.Error(list)
	}

	key3 := list.Add(tt)

	if key3 != key2 {
		t.Error(list)
	}
}
