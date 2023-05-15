package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/unng-lab/cowl"
)

func main() {
	f := cowl.Do(hardLogic, "test", 1)
	ll := lightLogic()
	a, b := f()
	fmt.Println(a, b, ll)
}

func hardLogic(a string, b int) (string, error) {
	time.Sleep(5 * time.Second)
	return a + strconv.Itoa(b), nil
}

func lightLogic() string {
	time.Sleep(1 * time.Second)
	return "lightLogic test"
}
