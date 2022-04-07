package main

import (
	"fmt"
	//"math/bits"
	"regexp"
	"strconv"
	"strings"
)

func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

func isPowerOfThree2(n int) bool {
	s := strconv.FormatInt(int64(n), 3)
	ok, _ := regexp.Match(`^10*$`, []byte(s))
	return ok
}

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func hammingDistance(x int, y int) int {
	z := x ^ y
	n := 0
	for z != 0 {
		z &= (z - 1)
		n++
	}
	return n
}

func reverseBits1(num uint32) uint32 {
	m := uint32(0)
	for i := 0; i < 32; i++ {
		m = 2*m + num%2
		num >>= 1
	}
	return m
}

func reverseBits2(num uint32) uint32 {
	ret := uint64(0)
	power := uint64(24)

	for num != 0 {
		value := (uint64(num&0xff) * 0x0202020202 & 0x010884422010) % 1023
		ret += value << power
		num = num >> 8
		power -= 8
	}
	return uint32(ret)
}

func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xFF00FF00) >> 8) | ((num & 0x00FF00FF) << 8)
	num = ((num & 0xF0F0F0F0) >> 4) | ((num & 0x0F0F0F0F) << 4)
	num = ((num & 0xCCCCCCCC) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xAAAAAAAA) >> 1) | ((num & 0x55555555) << 1)
	return num
}

func main() {
	fmt.Println(strings.Split("asdasd,adasd,dasda,tg,rht,", ","))
}
