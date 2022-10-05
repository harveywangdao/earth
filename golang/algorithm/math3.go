package main

import (
	"fmt"
)

func add(n, m []byte) []byte {
	i, j := len(n)-1, len(m)-1
	var ret []byte
	c := byte(0)
	for {
		if i < 0 && j < 0 {
			break
		}

		var x, y byte
		if i >= 0 {
			x = n[i] - '0'
		}
		if j >= 0 {
			y = m[j] - '0'
		}

		z := x + y + c
		c = z / 10
		a := z % 10
		ret = append(ret, a+'0')
		i--
		j--
	}

	if c > 0 {
		ret = append(ret, c+'0')
	}

	var ret2 []byte
	for k := len(ret) - 1; k >= 0; k-- {
		ret2 = append(ret2, ret[k])
	}
	return ret2
}

/*
1234
  10
*/
func mul(n1, n2 string) string {
	var ret []byte
	for i := len(n1) - 1; i >= 0; i-- {
		if n1[i] == '0' {
			continue
		}

		var m []byte
		for k := len(n1) - 1; k > i; k-- {
			m = append(m, '0')
		}
		var s byte
		for j := len(n2) - 1; j >= 0; j-- {
			x := n1[i] - '0'
			y := n2[j] - '0'
			z := x*y + s
			d := z % 10
			m = append(m, d+'0')
			s = z / 10
		}
		if s > 0 {
			m = append(m, s+'0')
		}

		var ret2 []byte
		for k := len(m) - 1; k >= 0; k-- {
			ret2 = append(ret2, m[k])
		}

		if ret == nil {
			ret = ret2
		} else {
			ret = add(ret, ret2)
		}
	}

	return string(ret)
}

func main() {
	fmt.Println(9999 * 8888)
	fmt.Println(mul("9999", "8888"))
}
