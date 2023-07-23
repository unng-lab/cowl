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

package seq

import (
	"time"
)

type ExpSmoothing struct {
	before       time.Duration
	smoothFactor float64
}

func NewExpSmoothing(startDuration time.Duration, smoothFactor float64) *ExpSmoothing {
	if smoothFactor <= 0 || smoothFactor >= 1 {
		panic("smooth factor must be between 0 1")
	}
	return &ExpSmoothing{
		before:       startDuration,
		smoothFactor: smoothFactor,
	}
}

func (e *ExpSmoothing) Next(curDur time.Duration) time.Duration {
	e.before = time.Duration(int64(float64(curDur.Nanoseconds())*e.smoothFactor)) +
		time.Duration(int64((1-e.smoothFactor)*float64(e.before.Nanoseconds())))
	return e.before
}
