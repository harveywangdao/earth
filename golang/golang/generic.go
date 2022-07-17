package main

import (
	"fmt"
)

/*type Node[T] struct{
  value T
}

func do1() {
  n := Node[int] {
    value: 4,
  }
}

type Node2[T any] struct{
  value T
}

type Node3[T comparable] struct{
  value T
}

type Numeric interface{
  int | int32 | float64
}

type Node4[T Numeric] struct{
  value T
}

func min[T constraints.Ordered](x,y T) T {
  if x < y {
    return x
  }
  return y
}

func do2() {
  m := min[int](2,3)
}

type Tree[T interface{}] struct{
  left, right *Tree[T]
  data T
}

type Ordered interface{
  Integer | Float | ~string
}
*/

/*func Add1[T ~int](a, b T) T {
  return a + b
}

func Add2[T int](a, b T) T {
  return a + b
}

type MyInt int

func do3() {
	var a, b MyInt
	a = 2
	b = 3
	c := Add1(a, b)
	//Add2(a, b)

  fmt.Printf("%T %d\n", c, c)
}*/

type MyInter interface {
	int8 | int16 | int32 | int64
}

func Add[T MyInter](a, b T) T {
	return a + b
}

func Add2[T1, T2 MyInter](a T1, b T2) T1 {
	return a + T1(b)
}

func do4() {
	var a, b int8
	var c, d int32
	a, b = 1, 2
	c, d = 5, 3

	x := Add(a, b)
	y := Add(c, d)
	Add2(a, c)

	fmt.Printf("%T %T %d %d\n", x, y, x, y)
}

type Number interface {
	int64 | float64
}

func do5() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
  do5()
}
