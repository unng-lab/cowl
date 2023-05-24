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

	"github.com/samber/lo"
	"github.com/unng-lab/cowl/internal/async"
)

func BenchmarkSamberLoAsync(b *testing.B) {
	aa, bb, cc := "1", "2", "3"
	b.ReportAllocs()
	b.SetParallelism(4)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch := lo.Async3(func() (string, string, string) {
				return testFunc(aa, bb, cc)
			})
			x, y, z := lo.Unpack3(<-ch)
			_ = x + y + z
		}
	})
}

func BenchmarkDoThThTest(b *testing.B) {
	aa, bb, cc := "1", "2", "3"
	b.ReportAllocs()
	b.SetParallelism(4)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res := async.Async3to3(testFunc, aa, bb, cc)
			x, y, z := res()
			_ = x + y + z
		}
	})
}

func testFunc(aa, bb, cc string) (string, string, string) {
	return aa, bb, cc
}
