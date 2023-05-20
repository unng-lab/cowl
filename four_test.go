package cowl

import (
	"testing"
	"time"
)

func TestDoFZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d string) {
		time.Sleep(1 * time.Millisecond)
		z = a + b + c + d
	}

	f := DoFZ(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "abcd" {
		t.Error(z)
	}
}

func TestDoFOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b, c, d string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d
	}

	f := DoFO(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "abcd" {
		t.Error(z)
	}
}

func TestDoFTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b, c, d string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + "1", a + b + c + d + "2"
	}

	f := DoFT(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "abcd1" || y != "abcd2" {
		t.Error(z, y)
	}
}

func TestDoFFSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b, c, d string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + b + c + d + "1", a + b + c + d + "2", a + b + c + d + "3"
	}

	f := DoFTh(testFunc, "a", "b", "c", "d")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "abcd1" || y != "abcd2" || x != "abcd3" {
		t.Error(z, x, y)
	}
}
