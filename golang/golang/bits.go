package main

import (
	"fmt"
	"math"
)

func do1() {
	var maxInt64 int64 = 0x7FFFFFFFFFFFFFFF
	fmt.Println("math.MaxInt64:", math.MaxInt64, "0x7FFFFFFFFFFFFFFF:", maxInt64)

	var minInt64 int64 = maxInt64 + 1
	fmt.Println("math.MinInt64", math.MinInt64, "minInt64:", minInt64)

	var maxUint64 uint64 = 0xFFFFFFFFFFFFFFFF
	fmt.Println("math.MaxUint64:", uint64(math.MaxUint64), "0xFFFFFFFFFFFFFFFF:", maxUint64)
}

func do2() {
	var n1 int64 = -1
	n2 := uint64(n1)
	fmt.Printf("n2 = 0x%X, n2 = %d\n", n2, n2)

	var n3 uint64 = 0xFFFFFFFFFFFFFFFF
	n4 := int64(n3)
	fmt.Printf("n4 = %d\n", n4)
}

func EncodeZigZag(x int64) uint64 {
	return uint64(x<<1) ^ uint64(x>>63)
}

/*
-1
1111 1111
1111 1110 ^ 1111 1111 = 0000 0001

-2
1111 1110
1111 1100 ^ 1111 1111 = 0000 0011

-3
1111 1101
1111 1010 ^ 1111 1111 = 0000 0101
*/
func do3() {
	fmt.Println(EncodeZigZag(-1))
}

func do4() {
	// 正数<<
	fmt.Println("正数<<:")
	var a int64 = 1
	fmt.Printf("old: %064b\n<<1: %064b\n<<2: %064b\n\n", uint64(a), uint64(a<<1), uint64(a<<2))
	var a1 int64 = 0x4000000000000000
	fmt.Printf("old: %064b\n<<1: %064b\n<<2: %064b\n\n", uint64(a1), uint64(a1<<1), uint64(a1<<2))

	// 负数<<
	fmt.Println("负数<<:")
	var b int64 = -1
	fmt.Printf("old: %064b\n<<1: %064b\n<<2: %064b\n\n", uint64(b), uint64(b<<1), uint64(b<<2))
	var b1 int64 = math.MinInt64 + 1
	fmt.Printf("old: %064b\n<<1: %064b\n<<2: %064b\n\n", uint64(b1), uint64(b1<<1), uint64(b1<<2))

	// 正数>>
	fmt.Println("正数>>:")
	var c int64 = 0x4000000000000000
	fmt.Printf("old: %064b\n>>1: %064b\n>>2: %064b\n\n", uint64(c), uint64(c>>1), uint64(c>>2))

	// 负数>>
	fmt.Println("负数>>:")
	var d int64 = -1
	fmt.Printf("old: %064b\n>>1: %064b\n>>2: %064b\n\n", uint64(d), uint64(d>>1), uint64(d>>2))
}

// int8范围
func do5() {
	/*
	   -1
	   1000 0001
	   1111 1110
	   1111 1111

	   -127
	   1111 1111
	   1000 0000
	   1000 0001

	   -128
	   1000 0000
	   1111 1111
	   1000 0000
	*/
	var a int8 = -1
	var b int8 = -127
	var c int8 = -128

	fmt.Printf("-1   = 0b%08b\n", uint8(a))
	fmt.Printf("-127 = 0b%08b\n", uint8(b))
	fmt.Printf("-128 = 0b%08b\n", uint8(c))
}

// int8 -> int64
func do6() {
	var a int8 = -1
	var b int8 = -127
	var c int8 = -128

	fmt.Printf("uint8(int8(-1))    = 0b%08b\n", uint8(a))
	fmt.Printf("uint64(int8(-1))   = 0b%064b\n", uint64(a))
	fmt.Printf("int64(int8(-1))    = %d\n\n", int64(a))

	fmt.Printf("uint8(int8(-127))  = 0b%08b\n", uint8(b))
	fmt.Printf("uint64(int8(-127)) = 0b%064b\n", uint64(b))
	fmt.Printf("int64(int8(-127))  = %d\n\n", int64(b))

	fmt.Printf("uint8(int8(-128))  = 0b%08b\n", uint8(c))
	fmt.Printf("uint64(int8(-128)) = 0b%064b\n", uint64(c))
	fmt.Printf("int64(int8(-128))  = %d\n\n", int64(c))

	var x uint8 = 1
	var y uint8 = 127
	var z uint8 = 128

	fmt.Printf("uint8(1)          = 0b%08b\n", x)
	fmt.Printf("uint64(int8(1))   = 0b%064b\n", uint64(x))
	fmt.Printf("int64(int8(1))    = %d\n\n", int64(x))

	fmt.Printf("uint8(127)        = 0b%08b\n", y)
	fmt.Printf("uint64(int8(127)) = 0b%064b\n", uint64(y))
	fmt.Printf("int64(int8(127))  = %d\n\n", int64(y))

	fmt.Printf("uint8(128)        = 0b%08b\n", z)
	fmt.Printf("uint64(int8(128)) = 0b%064b\n", uint64(z))
	fmt.Printf("int64(int8(128))  = %d\n\n", int64(z))
}

// int64/uint64 -> int8/uint8
func do7() {
	var a int64 = -1
	var b int64 = -128
	var c int64 = -129

	fmt.Printf("uint64(int64(-1))   = 0b%064b\n", uint64(a))
	fmt.Printf("uint8(int64(-1))    = 0b%08b\n", uint8(a))
	fmt.Printf("int8(int64(-1))     = %d\n\n", int8(a))

	fmt.Printf("uint64(int64(-128)) = 0b%064b\n", uint64(b))
	fmt.Printf("uint8(int64(-128))  = 0b%08b\n", uint8(b))
	fmt.Printf("int8(int64(-128))   = %d\n\n", int8(b))

	fmt.Printf("uint64(int64(-129)) = 0b%064b\n", uint64(c))
	fmt.Printf("uint8(int64(-129))  = 0b%08b\n", uint8(c))
	fmt.Printf("int8(int64(-129))   = %d\n\n", int8(c))

	var x uint64 = 1
	var y uint64 = 255
	var z uint64 = 256

	fmt.Printf("uint64(1)          = 0b%064b\n", x)
	fmt.Printf("uint8(uint64(1))   = 0b%08b\n", uint8(x))
	fmt.Printf("int8(uint64(1))    = %d\n\n", int8(x))

	fmt.Printf("uint64(255)        = 0b%064b\n", y)
	fmt.Printf("uint8(uint64(255)) = 0b%08b\n", uint8(y))
	fmt.Printf("int8(uint64(255))  = %d\n\n", int8(y))

	fmt.Printf("uint64(256)        = 0b%064b\n", z)
	fmt.Printf("uint8(uint64(256)) = 0b%08b\n", uint8(z))
	fmt.Printf("int8(uint64(256))  = %d\n\n", int8(z))
}

func main() {
	do4()
}
