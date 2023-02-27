//go:build amd64

package int64vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	SimpleInput = Int64Vec4([4]int64{1, 2, 3, 4})
	SimpleInputReverse = Int64Vec4([4]int64{4, 3, 2, 1})
	ClampMaxInput = Int64Vec4([4]int64{3, 3, 3, 3})
	ClampMinInput = Int64Vec4([4]int64{1, 1, 1, 1})
)

// EXPECTED

var (
	AddExpected = Int64Vec4([4]int64{5, 5, 5, 5})
	SubExpected = Int64Vec4([4]int64{-3, -1, 1, 3})
	MulExpected = Int64Vec4([4]int64{4, 6, 6, 4})
	MulAddExpected = Int64Vec4([4]int64{2, 6, 12, 20})
	DivExpected = Int64Vec4([4]int64{1, 1, 1, 1})
	MinExpected = Int64Vec4([4]int64{1, 2, 2, 1})
	MaxExpected = Int64Vec4([4]int64{4, 3, 3, 4})
	ClampExpected = Int64Vec4([4]int64{1, 2, 3, 3})
)

// ADD

func Benchmark_Int64Vec4_addAVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		addAVX2(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_addSSE2(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		addSSE2(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Add(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Add(&SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Int64Vec4_subAVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		subAVX2(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_subSSE2(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		subSSE2(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Sub(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Sub(&SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Int64Vec4_mulAVX512DQ_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512DQ && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		mulAVX512DQ_VL(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_mulGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		mulGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Mul(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Mul(&SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Int64Vec4_Div(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Div(&SimpleInput, &result)
	}
	
	if result != DivExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Int64Vec4_minAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		minAVX512F_VL(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_minGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		minGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Min(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Min(&SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Int64Vec4_maxAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		maxAVX512F_VL(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_maxGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		maxGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Max(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Max(&SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Int64Vec4_clampAVX512F_VL(b *testing.B) {
	if !(cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL) {
		b.SkipNow();
	}

	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		clampAVX512F_VL(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_clampGeneric(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		clampGeneric(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}

func Benchmark_Int64Vec4_Clamp(b *testing.B) {
	var result Int64Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Clamp(&ClampMinInput, &ClampMaxInput, &result)
	}

	if result != ClampExpected {
		fmt.Println([4]int64(result))
		b.FailNow()
	}
}
