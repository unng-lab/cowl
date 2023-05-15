# Cowl

[![Test Status](https://github.com/unng-lab/cowl/actions/workflows/test.yml/badge.svg)](https://github.com/unng-lab/cowl/actions/workflows/test.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/unng-lab/cowl)
[![Go Reference](https://pkg.go.dev/badge/github.com/unng-lab/cowl.svg)](https://pkg.go.dev/github.com/unng-lab/cowl)
[![Go Report Card](https://goreportcard.com/badge/github.com/unng-lab/cowl)](https://goreportcard.com/report/github.com/unng-lab/cowl)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/unng-lab/cowl)
[![Coverage Status](https://coveralls.io/repos/github/unng-lab/cowl/badge.svg)](https://coveralls.io/github/unng-lab/cowl)

Cowl allows you to run a function in a parallel goroutine

## Getting Started

1. Install the dependency

```sh
go get -u github.com/unng-lab/cowl
```

2. Import the package and create a new function

```go
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
```

As u see normal execution would take 6 seconds and with cowl u can do it in 5 sec

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also
simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the Apache-2.0 License. See `LICENSE` for more information.
