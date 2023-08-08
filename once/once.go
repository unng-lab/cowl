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

package once

import (
	"sync"

	"github.com/unng-lab/cowl/async"
)

type Listener interface {
	Trigger()
}

type subscriber struct {
	l Listener
	m sync.Mutex
}

func (s *subscriber) Do() {
	s.l.Trigger()
	s.m.Unlock()
}

// Repeatable once
type Repeatable struct {
	m         sync.Mutex
	listeners map[Listener]*subscriber
}

func NewRepeatable() *Repeatable {
	return &Repeatable{
		listeners: make(map[Listener]*subscriber),
	}
}

func (r *Repeatable) AddListener(listener Listener) {
	r.m.Lock()
	s := subscriber{l: listener}
	r.listeners[listener] = &s
	r.m.Unlock()
}

func (r *Repeatable) RemoveListener(listener Listener) {
	r.m.Lock()
	delete(r.listeners, listener)
	r.m.Unlock()
}

func (r *Repeatable) Trigger() {
	r.m.Lock()
	defer r.m.Unlock()
	for k := range r.listeners {
		s := r.listeners[k]
		if s.m.TryLock() {
			async.Async0to0(s.Do)
		}
	}
}
