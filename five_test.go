package cowl

import (
	"testing"
	"time"
)

func TestDoFiZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d, e string) {
		time.Sleep(1 * time.Millisecond)
		z = a + b + c + d + e
	}

	f := DoFiZ(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "abcde" {
		t.Error(z)
	}
}

func TestDoFiOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d, e string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e
	}

	f := DoFiO(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "abcde" {
		t.Error(z)
	}
}

func TestDoFiTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b, c, d, e string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e + "1", a + b + c + d + e + "2"
	}

	f := DoFiT(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "abcde1" || y != "abcde2" {
		t.Error(z, y)
	}
}

func TestDoFiFSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b, c, d, e string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + e + "1", a + b + c + d + e + "2", a + b + c + d + e + "3"
	}

	f := DoFiTh(testFunc, "a", "b", "c", "d", "e")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "abcde1" || y != "abcde2" || x != "abcde3" {
		t.Error(z, x, y)
	}
}
