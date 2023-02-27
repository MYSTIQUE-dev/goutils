//go:build amd64

package math

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	dv4_SimpleInput = DoubleVec4([4]float64{1, 2, 3, 4})
	dv4_SimpleInputReverse = DoubleVec4([4]float64{4, 3, 2, 1})
	dv4_CeilInput = DoubleVec4([4]float64{0.9, 1.9, 2.9, 3.9})
	dv4_RoundInput = DoubleVec4([4]float64{0.9, 2.1, 2.9, 4.1})
	dv4_FloorInput = DoubleVec4([4]float64{1.9, 2.9, 3.9, 4.9})
	dv4_ClampMaxInput = DoubleVec4([4]float64{3, 3, 3, 3})
	dv4_ClampMinInput = DoubleVec4([4]float64{1, 1, 1, 1})
)

// EXPECTED

var (
	dv4_AddExpected = DoubleVec4([4]float64{5, 5, 5, 5})
	dv4_SubExpected = DoubleVec4([4]float64{-3, -1, 1, 3})
	dv4_MulExpected = DoubleVec4([4]float64{4, 6, 6, 4})
	dv4_MulAddExpected = DoubleVec4([4]float64{2, 6, 12, 20})
	dv4_DivExpected = DoubleVec4([4]float64{1, 1, 1, 1})
	dv4_MinExpected = DoubleVec4([4]float64{1, 2, 2, 1})
	dv4_MaxExpected = DoubleVec4([4]float64{4, 3, 3, 4})
	dv4_ClampExpected = DoubleVec4([4]float64{1, 2, 3, 3})
	dv4_RoundExpected = DoubleVec4([4]float64{1, 2, 3, 4})
)

// ADD

func Benchmark_DoubleVec4_dv4_AddAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_AddAVX(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_AddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_AddSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_AddSSE2(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_AddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Add(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Add(&dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_AddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_DoubleVec4_dv4_SubAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SubAVX(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_SubExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_SubSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SubSSE2(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_SubExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Sub(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Sub(&dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_SubExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_DoubleVec4_dv4_MulAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MulAVX(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MulExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_MulSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MulSSE2(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MulExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Mul(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Mul(&dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MulExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// MULADD

func Benchmark_DoubleVec4_dv4_MulAddFMA(b *testing.B) {
	if !cpu.X86.HasFMA {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MulAddFMA(&dv4_SimpleInput, &dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_MulAddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_MulAddAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MulAddAVX(&dv4_SimpleInput, &dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_MulAddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_MulAddSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MulAddSSE2(&dv4_SimpleInput, &dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_MulAddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_MulAdd(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.MulAdd(&dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_MulAddExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_DoubleVec4_dv4_DivAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_DivAVX(&dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_DivExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_DivSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_DivSSE2(&dv4_SimpleInput, &dv4_SimpleInput, &result)
	}
	
	if result != dv4_DivExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Div(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Div(&dv4_SimpleInput, &result)
	}
	
	if result != dv4_DivExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_DoubleVec4_dv4_MinAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MinAVX(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MinExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_MinSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MinSSE2(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MinExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Min(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Min(&dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MinExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_DoubleVec4_dv4_MaxAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MaxAVX(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MaxExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_MaxSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_MaxSSE2(&dv4_SimpleInput, &dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MaxExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Max(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Max(&dv4_SimpleInputReverse, &result)
	}
	
	if result != dv4_MaxExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_DoubleVec4_dv4_ClampAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_ClampAVX(&dv4_SimpleInput, &dv4_ClampMinInput, &dv4_ClampMaxInput, &result)
	}
	
	if result != dv4_ClampExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_ClampSSE2(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_ClampSSE2(&dv4_SimpleInput, &dv4_ClampMinInput, &dv4_ClampMaxInput, &result)
	}
	
	if result != dv4_ClampExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Clamp(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_SimpleInput.Clamp(&dv4_ClampMinInput, &dv4_ClampMaxInput, &result)
	}

	if result != dv4_ClampExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// CEIL

func Benchmark_DoubleVec4_dv4_CeilAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_CeilAVX(&dv4_CeilInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_CeilSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_CeilSSE41(&dv4_CeilInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_CeilGeneric(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_CeilGeneric(&dv4_CeilInput, &result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Ceil(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_CeilInput.Ceil(&result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// ROUND

func Benchmark_DoubleVec4_dv4_RoundAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_RoundAVX(&dv4_RoundInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_RoundSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_RoundSSE41(&dv4_RoundInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_RoundGeneric(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_RoundGeneric(&dv4_RoundInput, &result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Round(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_RoundInput.Round(&result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

// FLOOR

func Benchmark_DoubleVec4_dv4_FloorAVX(b *testing.B) {
	if !cpu.X86.HasAVX {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_FloorAVX(&dv4_FloorInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_FloorSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_FloorSSE41(&dv4_FloorInput, &result)
	}
	
	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_dv4_FloorGeneric(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_FloorGeneric(&dv4_FloorInput, &result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}

func Benchmark_DoubleVec4_Floor(b *testing.B) {
	var result DoubleVec4

	for n := 0; n < b.N; n++ {
		dv4_FloorInput.Floor(&result)
	}

	if result != dv4_RoundExpected {
		fmt.Println([4]float64(result))
		b.FailNow()
	}
}