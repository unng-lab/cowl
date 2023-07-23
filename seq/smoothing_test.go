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
	"testing"
	"time"
)

const (
	name = 10e7
)

func TestExpSmoothing_Next(t *testing.T) {
	type fields struct {
		before       time.Duration
		smoothFactor float64
	}
	type args struct {
		curDur time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name: "test1",
			fields: fields{
				before:       time.Second,
				smoothFactor: 0.3,
			},
			args: args{
				curDur: time.Second,
			},
			want: time.Second,
		},
		{
			name: "test2",
			fields: fields{
				before:       time.Second,
				smoothFactor: 0.3,
			},
			args: args{
				curDur: time.Second * 2,
			},
			want: time.Duration(int64(1.3 * 10e8)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExpSmoothing{
				before:       tt.fields.before,
				smoothFactor: tt.fields.smoothFactor,
			}
			var got time.Duration
			if got = e.Next(tt.args.curDur); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
			if e.before != got {
				t.Errorf("Next() = %v, want %v", e.before, got)
			}
		})
	}
}
