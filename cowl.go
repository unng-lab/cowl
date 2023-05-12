package cowl

var channels []chan Runnable
var queue chan int
var num int

func init() {
	n := 5
	if num > 0 {
		n = num
	}
	queue = make(chan int, n)
	for i := 0; i < n; i++ {
		channels = append(channels, make(chan Runnable, 1))
		queue <- i
	}
	go run()
}

func SetNum(n int) {
	num = n
}

func run() {
	for i := range channels {
		go func(k int) {
			for {
				select {
				case s := <-channels[k]:
					s.Run()
					queue <- k
				}
			}
		}(i)
	}

}

type Runnable interface {
	Run()
}
