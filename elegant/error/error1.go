package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := errors.Errorf("%s", "hadoop")
	fmt.Println(err)

	err2 := errors.Wrap(err, "spark")
	fmt.Println(err2)

	err2 = errors.Wrap(err2, "storm")
	fmt.Println(err2)

	err2 = errors.Wrap(err2, "flink")
	fmt.Println(err2)

	err2 = errors.Wrap(err2, "flume")
	fmt.Println(err2)

	i := 0
	if err3, ok := err2.(stackTracer); ok {
		for _, f := range err3.StackTrace() {
			fmt.Printf("%+s:%d-------%d\n", f, f, i)
			i++
		}
	}
}
