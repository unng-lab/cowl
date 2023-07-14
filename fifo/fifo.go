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
	"errors"
	"sync"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Store[T comparable, K any] struct {
	next chan T
	list map[T]K
	size int
	m    sync.RWMutex
}

func NewStore[T comparable, K any](_ T, _ K, size int) *Store[T, K] {
	if size < 1 {
		panic("size cant be less than 1")
	}
	return &Store[T, K]{
		next: make(chan T, size),
		list: make(map[T]K, size),
		size: size,
	}
}

func (s *Store[T, K]) Get(key T) (*K, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	if v, ok := s.list[key]; ok {
		return &v, nil
	}
	return nil, ErrKeyNotFound
}

func (s *Store[T, K]) Put(key T, value K) {
	s.m.Lock()
	defer s.m.Unlock()
	if len(s.next) == s.size {
		delete(s.list, <-s.next)
		s.list[key] = value
		s.next <- key
		return
	}
	s.list[key] = value
	return
}
