package cowl

import (
	"testing"
	"time"
)

func TestDoThZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c string) {
		time.Sleep(1 * time.Millisecond)
		z = a + b + c
	}

	f := DoThZ(testFunc, "a", "b", "c")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "abc" {
		t.Error(z)
	}
}

func TestDoThOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b + c
	}

	f := DoThO(testFunc, "a", "b", "c")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "abc" {
		t.Error(z)
	}
}

func TestDoThTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b, c string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + "1", a + b + c + "2"
	}

	f := DoThT(testFunc, "a", "b", "c")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "abc1" || y != "abc2" {
		t.Error(z, y)
	}
}

func TestDoThThSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b, c string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + "1", a + b + c + "2", a + b + c + "3"
	}

	f := DoThTh(testFunc, "a", "b", "c")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "abc1" || y != "abc2" || x != "abc3" {
		t.Error(z, x, y)
	}
}
