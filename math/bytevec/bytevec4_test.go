package bytevec

import (
	"testing"
	"fmt"
)

// INPUTS

var (
	bv4_SimpleInput = ByteVec4([4]byte{1, 2, 3, 4})
	bv4_SimpleInputReverse = ByteVec4([4]byte{4, 3, 2, 1})
	bv4_ClampMaxInput = ByteVec4([4]byte{3, 3, 3, 3})
	bv4_ClampMinInput = ByteVec4([4]byte{1, 1, 1, 1})
)

// EXPECTED

var (
	bv4_AddExpected = ByteVec4([4]byte{5, 5, 5, 5})
	bv4_SubExpected = ByteVec4([4]byte{253, 255, 1, 3})
	bv4_MulExpected = ByteVec4([4]byte{4, 6, 6, 4})
	bv4_DivExpected = ByteVec4([4]byte{1, 1, 1, 1})
	bv4_MinExpected = ByteVec4([4]byte{1, 2, 2, 1})
	bv4_MaxExpected = ByteVec4([4]byte{4, 3, 3, 4})
	bv4_ClampExpected = ByteVec4([4]byte{1, 2, 3, 3})
)

// ADD

func Benchmark_ByteVec4_Add(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Add(&bv4_SimpleInputReverse, &result)
	}
	
	if result != bv4_AddExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_ByteVec4_Sub(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Sub(&bv4_SimpleInputReverse, &result)
	}
	
	if result != bv4_SubExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_ByteVec4_Mul(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Mul(&bv4_SimpleInputReverse, &result)
	}
	
	if result != bv4_MulExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_ByteVec4_Div(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Div(&bv4_SimpleInput, &result)
	}
	
	if result != bv4_DivExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_ByteVec4_Min(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Min(&bv4_SimpleInputReverse, &result)
	}
	
	if result != bv4_MinExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_ByteVec4_Max(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Max(&bv4_SimpleInputReverse, &result)
	}
	
	if result != bv4_MaxExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_ByteVec4_Clamp(b *testing.B) {
	var result ByteVec4

	for n := 0; n < b.N; n++ {
		bv4_SimpleInput.Clamp(&bv4_ClampMinInput, &bv4_ClampMaxInput, &result)
	}

	if result != bv4_ClampExpected {
		fmt.Println([4]byte(result))
		b.FailNow()
	}
}
