package main

import (
	"fmt"
	//"github.com/pkg/errors"
	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/errbase"
)

func main() {
	err := errors.New("hadoop")
	fmt.Println(err)

	if err1, ok := err.(errbase.SafeDetailer); ok {
		fmt.Println(err1.SafeDetails())
	}
}
