package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	stack := &MyStack{}

	for i := 0; i < n; i++ {
		c := s[i]
		if stack.Size() > 0 {
			cc := stack.Peek().(byte)
			if cc == '{' && c == '}' {
				stack.Pop()
				continue
			} else if cc == '(' && c == ')' {
				stack.Pop()
				continue
			} else if cc == '[' && c == ']' {
				stack.Pop()
				continue
			}
		}

		if c == '}' || c == ')' || c == ']' {
			return false
		}

		stack.Push(c)
	}

	if stack.IsEmpty() {
		return true
	}
	return false
}

func main() {

}
