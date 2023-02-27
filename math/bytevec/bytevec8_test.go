package bytevec

import (
	"testing"
	"fmt"
)

// INPUTS

var (
	bv8_SimpleInput = ByteVec8([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
	bv8_SimpleInputReverse = ByteVec8([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	bv8_ClampMaxInput = ByteVec8([8]byte{6, 6, 6, 6, 6, 6, 6, 6})
	bv8_ClampMinInput = ByteVec8([8]byte{1, 1, 1, 1, 1, 1, 1, 1})
)

// EXPECTED

var (
	bv8_AddExpected = ByteVec8([8]byte{9, 9, 9, 9, 9, 9, 9, 9})
	bv8_SubExpected = ByteVec8([8]byte{249, 251, 253, 255, 1, 3, 5, 7})
	bv8_MulExpected = ByteVec8([8]byte{8, 14, 18, 20, 20, 18, 14, 8})
	bv8_DivExpected = ByteVec8([8]byte{1, 1, 1, 1, 1, 1, 1, 1})
	bv8_MinExpected = ByteVec8([8]byte{1, 2, 3, 4, 4, 3, 2, 1})
	bv8_MaxExpected = ByteVec8([8]byte{8, 7, 6, 5, 5, 6, 7, 8})
	bv8_ClampExpected = ByteVec8([8]byte{1, 2, 3, 4, 5, 6, 6, 6})
)

// ADD

func Benchmark_ByteVec8_Add(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Add(&bv8_SimpleInputReverse, &result)
	}
	
	if result != bv8_AddExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_ByteVec8_Sub(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Sub(&bv8_SimpleInputReverse, &result)
	}
	
	if result != bv8_SubExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_ByteVec8_Mul(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Mul(&bv8_SimpleInputReverse, &result)
	}
	
	if result != bv8_MulExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_ByteVec8_Div(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Div(&bv8_SimpleInput, &result)
	}
	
	if result != bv8_DivExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_ByteVec8_Min(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Min(&bv8_SimpleInputReverse, &result)
	}
	
	if result != bv8_MinExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_ByteVec8_Max(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Max(&bv8_SimpleInputReverse, &result)
	}
	
	if result != bv8_MaxExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_ByteVec8_Clamp(b *testing.B) {
	var result ByteVec8

	for n := 0; n < b.N; n++ {
		bv8_SimpleInput.Clamp(&bv8_ClampMinInput, &bv8_ClampMaxInput, &result)
	}

	if result != bv8_ClampExpected {
		fmt.Println([8]byte(result))
		b.FailNow()
	}
}
