package bytevec

import (
	"testing"
	"fmt"
)

// INPUTS

var (
	bv2_SimpleInput = ByteVec2([2]byte{1, 4})
	bv2_SimpleInputReverse = ByteVec2([2]byte{4, 1})
	bv2_ClampMaxInput = ByteVec2([2]byte{3, 3})
	bv2_ClampMinInput = ByteVec2([2]byte{2, 2})
)

// EXPECTED

var (
	bv2_AddExpected = ByteVec2([2]byte{5, 5})
	bv2_SubExpected = ByteVec2([2]byte{253, 3})
	bv2_MulExpected = ByteVec2([2]byte{4, 4})
	bv2_DivExpected = ByteVec2([2]byte{1, 1})
	bv2_MinExpected = ByteVec2([2]byte{1, 1})
	bv2_MaxExpected = ByteVec2([2]byte{4, 4})
	bv2_ClampExpected = ByteVec2([2]byte{2, 3})
)

// ADD

func Benchmark_ByteVec2_Add(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Add(&bv2_SimpleInputReverse, &result)
	}
	
	if result != bv2_AddExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_ByteVec2_Sub(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Sub(&bv2_SimpleInputReverse, &result)
	}
	
	if result != bv2_SubExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_ByteVec2_Mul(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Mul(&bv2_SimpleInputReverse, &result)
	}
	
	if result != bv2_MulExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_ByteVec2_Div(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Div(&bv2_SimpleInput, &result)
	}
	
	if result != bv2_DivExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_ByteVec2_Min(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Min(&bv2_SimpleInputReverse, &result)
	}
	
	if result != bv2_MinExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_ByteVec2_Max(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Max(&bv2_SimpleInputReverse, &result)
	}
	
	if result != bv2_MaxExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_ByteVec2_Clamp(b *testing.B) {
	var result ByteVec2

	for n := 0; n < b.N; n++ {
		bv2_SimpleInput.Clamp(&bv2_ClampMinInput, &bv2_ClampMaxInput, &result)
	}

	if result != bv2_ClampExpected {
		fmt.Println([2]byte(result))
		b.FailNow()
	}
}
