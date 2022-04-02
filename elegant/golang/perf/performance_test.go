package main

import (
	"testing"
)

func TestReadArray1(t *testing.T) {
	ReadArray1()
}

func TestReadArray2(t *testing.T) {
	ReadArray2()
}

func BenchmarkReadArray1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadArray1()
	}
}

func BenchmarkReadArray2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadArray2()
	}
}
