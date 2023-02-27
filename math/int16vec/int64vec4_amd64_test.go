//go:build amd64

package int64vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	int64v4_SimpleInput = Int64Vec4([4]int64{1, 2, 3, 4})
	int64v4_SimpleInputReverse = Int64Vec4([4]int64{4, 3, 2, 1})
	int64v4_ClampMaxInput = Int64Vec4([4]int64{3, 3, 3, 3})
	int64v4_ClampMinInput = Int64Vec4([4]int64{1, 1, 1, 1})
)

// EXPECTED

var (
	int64v4_AddExpected = Int64Vec4([4]int64{5, 5, 5, 5})
	int64v4_SubExpected = Int64Vec4([4]int64{-3, -1, 1, 3})
	int64v4_MulExpected = Int64Vec4([4]int64{4, 6, 6, 4})
	int64v4_MulAddExpected = Int64Vec4([4]int64{2, 6, 12, 20})
	int64v4_DivExpected = Int64Vec4([4]int64{1, 1, 1, 1})
	int64v4_MinExpected = Int64Vec4([4]int64{1, 2, 2, 1})
	int64v4_MaxExpected = Int64Vec4([4]int64{4, 3, 3, 4})
	int64v4_ClampExpected = Int64Vec4([4]int64{1, 2, 3, 3})
)

// ADD

func Benchmark_Int64Vec4_int64v4_AddAVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_AddAVX2(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_AddSSE2(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_AddSSE2(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Add(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Add(&int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Int64Vec4_int64v4_SubAVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SubAVX2(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_SubSSE2(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SubSSE2(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Sub(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Sub(&int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Int64Vec4_int64v4_MulAVX512DQ_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512DQ && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MulAVX512DQ_VL(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_MulGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MulGeneric(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Mul(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Mul(&int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Int64Vec4_Div(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Div(&int64v4_SimpleInput, &result)
	}
	
	if result != int64v4_DivExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Int64Vec4_int64v4_MinAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MinAVX512F_VL(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_MinGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MinGeneric(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Min(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Min(&int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Int64Vec4_int64v4_MaxAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MaxAVX512F_VL(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_MaxGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_MaxGeneric(&int64v4_SimpleInput, &int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Max(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Max(&int64v4_SimpleInputReverse, &result)
	}
	
	if result != int64v4_MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Int64Vec4_int64v4_ClampAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_ClampAVX512F_VL(&int64v4_SimpleInput, &int64v4_ClampMinInput, &int64v4_ClampMaxInput, &result)
	}
	
	if result != int64v4_ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_int64v4_ClampGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_ClampGeneric(&int64v4_SimpleInput, &int64v4_ClampMinInput, &int64v4_ClampMaxInput, &result)
	}
	
	if result != int64v4_ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Clamp(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		int64v4_SimpleInput.Clamp(&int64v4_ClampMinInput, &int64v4_ClampMaxInput, &result)
	}

	if result != int64v4_ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}
