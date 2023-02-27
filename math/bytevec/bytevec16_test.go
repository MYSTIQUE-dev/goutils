package bytevec

import (
	"testing"
	"fmt"
)

// INPUTS

var (
	SimpleInput = ByteVec16([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	SimpleInputReverse = ByteVec16([16]byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	ClampMaxInput = ByteVec16([16]byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	ClampMinInput = ByteVec16([16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
)

// EXPECTED

var (
	AddExpected = ByteVec16([16]byte{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17})
	SubExpected = ByteVec16([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 5, 7, 9, 11, 13, 15})
	MulExpected = ByteVec16([16]byte{16, 30, 42, 52, 60, 66, 70, 72, 72, 70, 66, 60, 52, 42, 30, 16})
	DivExpected = ByteVec16([16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	MinExpected = ByteVec16([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	MaxExpected = ByteVec16([16]byte{16, 15, 14, 13, 12, 11, 10, 9, 9, 10, 11, 12, 13, 14, 15, 16})
	ClampExpected = ByteVec16([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9, 9, 9})
)

// ADD

func Benchmark_ByteVec16_Add(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Add(&SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_ByteVec16_Sub(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Sub(&SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_ByteVec16_Mul(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Mul(&SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_ByteVec16_Div(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Div(&SimpleInput, &result)
	}
	
	if result != DivExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_ByteVec16_Min(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Min(&SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_ByteVec16_Max(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Max(&SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_ByteVec16_Clamp(b *testing.B) {
	var result ByteVec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Clamp(&ClampMinInput, &ClampMaxInput, &result)
	}

	if result != ClampExpected {
		fmt.Println([16]byte(result))
		b.FailNow()
	}
}
