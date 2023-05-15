package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/unng-lab/cowl"
)

func main() {
	f := cowl.Do(hardLogic, "test", 1)
	sl := smallLogic()
	a, b := f()
	fmt.Println(a, b, sl)
}

func hardLogic(a string, b int) (string, error) {
	time.Sleep(5 * time.Second)
	return a + strconv.Itoa(b), nil
}

func smallLogic() string {
	time.Sleep(1 * time.Second)
	return "smallLogic test"
}
