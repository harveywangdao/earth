package mymath_test

import (
	"test/testing/mymath"
	"testing"
)

func TestAdd(t *testing.T) {
	ret := mymath.Add(2, 3)
	if ret != 2 {
		t.Error("Expected 5, got ", ret)
	}
}

func TestMinus(t *testing.T) {
	ret := mymath.Minus(2, 3)
	if ret != -5 {
		t.Error("Expected -1, got ", ret)
	}
}
