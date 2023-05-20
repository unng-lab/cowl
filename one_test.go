package cowl

import (
	"testing"
	"time"
)

func TestDoOZSimple(t *testing.T) {
	var z string
	testFunc := func(a string) {
		time.Sleep(1 * time.Millisecond)
		z = a
	}

	f := DoOZ(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	f()
	if z != "test" {
		t.Error(z)
	}
}

func TestDoOOSimple(t *testing.T) {
	var z string
	testFunc := func(a string) string {
		time.Sleep(1 * time.Millisecond)
		return a
	}

	f := DoOO(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z = f()
	if z != "test" {
		t.Error(z)
	}
}

func TestDoOTSimple(t *testing.T) {
	var z, y string
	testFunc := func(a string) (string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + "1", a + "2"
	}

	f := DoOT(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z, y = f()
	if z != "test1" || y != "test2" {
		t.Error(z, y)
	}
}

func TestDoOThSimple(t *testing.T) {
	var z, y, x string
	testFunc := func(a string) (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return a + "1", a + "2", a + "3"
	}

	f := DoOTh(testFunc, "test")

	time.Sleep(2 * time.Millisecond)

	z, y, x = f()
	if z != "test1" || y != "test2" || x != "test3" {
		t.Error(z, x, y)
	}
}
