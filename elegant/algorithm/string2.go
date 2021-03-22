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

type level struct {
	buf *bytes.Buffer
	k   int
}

func decodeString(s string) string {
	stack := &MyStack{}
	stack.Push(&level{
		buf: bytes.NewBuffer(nil),
		k:   1,
	})

	k := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			k = k*10 + int(c-'0')
			continue
		}

		if c == '[' {
			stack.Push(&level{
				buf: bytes.NewBuffer(nil),
				k:   k,
			})
			k = 0
			continue
		}
		if c == ']' {
			cur := stack.Pop().(*level)
			str := cur.buf.String()
			last := stack.Peek().(*level)
			for j := 0; j < cur.k; j++ {
				last.buf.WriteString(str)
			}
			continue
		}

		le := stack.Peek().(*level)
		le.buf.WriteByte(c)
	}
	le := stack.Pop().(*level)
	return le.buf.String()
}

func main() {

}
