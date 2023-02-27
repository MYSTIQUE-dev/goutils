package bytevec

import (
	"testing"
	"fmt"
)

// INPUTS

var (
	bv3_SimpleInput = ByteVec3([3]byte{1, 2, 3})
	bv3_SimpleInputReverse = ByteVec3([3]byte{3, 2, 1})
	bv3_ClampMaxInput = ByteVec3([3]byte{3, 3, 3})
	bv3_ClampMinInput = ByteVec3([3]byte{2, 2, 2})
)

// EXPECTED

var (
	bv3_AddExpected = ByteVec3([3]byte{4, 4, 4})
	bv3_SubExpected = ByteVec3([3]byte{254, 0, 2})
	bv3_MulExpected = ByteVec3([3]byte{3, 4, 3})
	bv3_DivExpected = ByteVec3([3]byte{1, 1, 1})
	bv3_MinExpected = ByteVec3([3]byte{1, 2, 1})
	bv3_MaxExpected = ByteVec3([3]byte{3, 2, 3})
	bv3_ClampExpected = ByteVec3([3]byte{2, 2, 3})
)

// ADD

func Benchmark_ByteVec3_Add(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Add(&bv3_SimpleInputReverse, &result)
	}
	
	if result != bv3_AddExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_ByteVec3_Sub(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Sub(&bv3_SimpleInputReverse, &result)
	}
	
	if result != bv3_SubExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_ByteVec3_Mul(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Mul(&bv3_SimpleInputReverse, &result)
	}
	
	if result != bv3_MulExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_ByteVec3_Div(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Div(&bv3_SimpleInput, &result)
	}
	
	if result != bv3_DivExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_ByteVec3_Min(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Min(&bv3_SimpleInputReverse, &result)
	}
	
	if result != bv3_MinExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_ByteVec3_Max(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Max(&bv3_SimpleInputReverse, &result)
	}
	
	if result != bv3_MaxExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_ByteVec3_Clamp(b *testing.B) {
	var result ByteVec3

	for n := 0; n < b.N; n++ {
		bv3_SimpleInput.Clamp(&bv3_ClampMinInput, &bv3_ClampMaxInput, &result)
	}

	if result != bv3_ClampExpected {
		fmt.Println([3]byte(result))
		b.FailNow()
	}
}
