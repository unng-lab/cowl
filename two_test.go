package cowl

import (
	"testing"
	"time"
)

func TestDoTZSimple(t *testing.T) {
	var z string
	testFunc := func(a, b string) {
		time.Sleep(1 * time.Millisecond)
		z = a + " " + b
	}

	f := DoTZ(testFunc, "a-test", "b-test")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "a-test b-test" {
		t.Error(z)
	}
}

func TestDoTOSimple(t *testing.T) {
	var z string
	testFunc := func(a, b string) string {
		time.Sleep(1 * time.Millisecond)
		return a + b
	}

	f := DoTO(testFunc, "a", "b")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "ab" {
		t.Error(z)
	}
}

func TestDoTTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a, b string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + " " + b + "1", a + " " + b + "2"
	}

	f := DoTT(testFunc, "a-test", "b-test")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "a-test b-test1" || y != "a-test b-test2" {
		t.Error(z, y)
	}
}

func TestDoTThSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a, b string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + " " + b + "1", a + " " + b + "2", a + " " + b + "3"
	}

	f := DoTTh(testFunc, "a-test", "b-test")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "a-test b-test1" || y != "a-test b-test2" || x != "a-test b-test3" {
		t.Error(z, x, y)
	}
}
