package cowl

import (
	"strconv"
	"testing"
	"time"
)

func TestDoSimple(t *testing.T) {
	testFunc := func(a string, b int) (string, error) {
		time.Sleep(1 * time.Millisecond)
		return a + strconv.Itoa(b), nil
	}

	f := Do(testFunc, "test", 1)

	time.Sleep(2 * time.Millisecond)

	a, b := f()
	if a != "test1" || b != nil {
		t.Error(a, b)
	}
}

func TestDoMustWait(t *testing.T) {
	testFunc := func(a string, b int) (string, error) {
		time.Sleep(1 * time.Millisecond)
		return a + strconv.Itoa(b), nil
	}

	f := Do(testFunc, "test", 1)

	a, b := f()
	if a != "test1" || b != nil {
		t.Error(a, b)
	}
}

func TestDoMany(t *testing.T) {
	for i := 0; i < 10; i++ {
		testFunc := func(a string, b int) (string, error) {
			time.Sleep(1 * time.Millisecond)
			return a + strconv.Itoa(b), nil
		}

		f := Do(testFunc, "test", 1)

		a, b := f()
		if a != "test1" || b != nil {
			t.Error(a, b)
		}
	}
}
