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
	"errors"
)

var (
	ErrNotExist       = errors.New("not exist")
	ErrAlreadyDeleted = errors.New("already deleted")
)

type Piece[T any] struct {
	entity  T
	deleted bool
}

type List[T any] struct {
	pieces []Piece[T]
	dList  []int
}

func New[T any](_ T, cap int) List[T] {
	res := List[T]{
		pieces: make([]Piece[T], 0, cap),
		dList:  make([]int, 0),
	}
	return res
}

func (l *List[T]) Add(piece T) int {
	if c := len(l.dList); c > 0 {
		key := l.dList[c-1]
		l.dList = l.dList[:c-1]
		return key
	} else {
		l.pieces = append(l.pieces, Piece[T]{entity: piece})
		return len(l.pieces) - 1
	}
}

func (l *List[T]) Remove(pieceKey int) error {
	if pieceKey < 0 || pieceKey > len(l.pieces)-1 {
		return ErrNotExist
	}
	if l.pieces[pieceKey].deleted {
		return ErrAlreadyDeleted
	}
	l.pieces[pieceKey].deleted = true
	l.dList = append(l.dList, pieceKey)
	return nil
}

func (l *List[T]) Range(f func(key int, value T) bool) {
	for k := range l.pieces {
		f(k, l.pieces[k].entity)
	}
}

func (l *List[T]) Get(pieceKey int) (*T, error) {
	if pieceKey < 0 || pieceKey > len(l.pieces)-1 {
		return nil, ErrNotExist
	}

	if l.pieces[pieceKey].deleted {
		return nil, ErrAlreadyDeleted
	}

	return &l.pieces[pieceKey].entity, nil
}
