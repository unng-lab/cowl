package cowl

import (
	"testing"
	"time"
)

func TestDoZZSimple(t *testing.T) {
	var a string
	testFunc := func() {
		time.Sleep(1 * time.Millisecond)
		a = "test"
	}

	f := DoZZ(testFunc)

	time.Sleep(2 * time.Millisecond)

	f()
	if a != "test" {
		t.Error(a)
	}
}

func TestDoZOSimple(t *testing.T) {
	var a string
	testFunc := func() string {
		time.Sleep(1 * time.Millisecond)
		return "test"
	}

	f := DoZO(testFunc)

	time.Sleep(2 * time.Millisecond)

	a = f()
	if a != "test" {
		t.Error(a)
	}
}

func TestDoZTSimple(t *testing.T) {
	var a, b string
	testFunc := func() (string, string) {
		time.Sleep(1 * time.Millisecond)
		return "test", "test1"
	}

	f := DoZT(testFunc)

	time.Sleep(2 * time.Millisecond)

	a, b = f()
	if a != "test" || b != "test1" {
		t.Error(a, b)
	}
}

func TestDoZThSimple(t *testing.T) {
	var a, b, c string
	testFunc := func() (string, string, string) {
		time.Sleep(1 * time.Millisecond)
		return "test", "test1", "test2"
	}

	f := DoZTh(testFunc)

	time.Sleep(2 * time.Millisecond)

	a, b, c = f()
	if a != "test" || b != "test1" || c != "test2" {
		t.Error(a)
	}
}
