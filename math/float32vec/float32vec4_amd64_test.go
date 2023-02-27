//go:build amd64

package float32vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	f32v4_SimpleInput = Float32Vec4([4]float32{1, 2, 3, 4})
	f32v4_SimpleInputReverse = Float32Vec4([4]float32{4, 3, 2, 1})
	f32v4_CeilInput = Float32Vec4([4]float32{0.9, 1.9, 2.9, 3.9})
	f32v4_RoundInput = Float32Vec4([4]float32{0.9, 2.1, 2.9, 4.1})
	f32v4_FloorInput = Float32Vec4([4]float32{1.9, 2.9, 3.9, 4.9})
	f32v4_ClampMaxInput = Float32Vec4([4]float32{3, 3, 3, 3})
	f32v4_ClampMinInput = Float32Vec4([4]float32{1, 1, 1, 1})
)

// EXPECTED

var (
	f32v4_AddExpected = Float32Vec4([4]float32{5, 5, 5, 5})
	f32v4_SubExpected = Float32Vec4([4]float32{-3, -1, 1, 3})
	f32v4_MulExpected = Float32Vec4([4]float32{4, 6, 6, 4})
	f32v4_FMAExpected = Float32Vec4([4]float32{2, 6, 12, 20})
	f32v4_DivExpected = Float32Vec4([4]float32{1, 1, 1, 1})
	f32v4_MinExpected = Float32Vec4([4]float32{1, 2, 2, 1})
	f32v4_MaxExpected = Float32Vec4([4]float32{4, 3, 3, 4})
	f32v4_ClampExpected = Float32Vec4([4]float32{1, 2, 3, 3})
	f32v4_RoundExpected = Float32Vec4([4]float32{1, 2, 3, 4})
)

// ADD

func Benchmark_Float32Vec4_Add(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Add(&f32v4_SimpleInputReverse, &result)
	}
	
	if result != f32v4_AddExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Float32Vec4_Sub(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Sub(&f32v4_SimpleInputReverse, &result)
	}
	
	if result != f32v4_SubExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Float32Vec4_Mul(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Mul(&f32v4_SimpleInputReverse, &result)
	}
	
	if result != f32v4_MulExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// FMA

func Benchmark_Float32Vec4_FMA_FMA(b *testing.B) {
	if !cpu.X86.HasFMA {
		b.SkipNow();
	}

	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_fmaFMA(&f32v4_SimpleInput, &f32v4_SimpleInput, &f32v4_SimpleInput, &result)
	}
	
	if result != f32v4_FMAExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_FMA_Generic(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_fmaGeneric(&f32v4_SimpleInput, &f32v4_SimpleInput, &f32v4_SimpleInput, &result)
	}
	
	if result != f32v4_FMAExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_FMA(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.FMA(&f32v4_SimpleInput, &f32v4_SimpleInput, &result)
	}
	
	if result != f32v4_FMAExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Float32Vec4_Div(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Div(&f32v4_SimpleInput, &result)
	}
	
	if result != f32v4_DivExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Float32Vec4_Min(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Min(&f32v4_SimpleInputReverse, &result)
	}
	
	if result != f32v4_MinExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Float32Vec4_Max(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Max(&f32v4_SimpleInputReverse, &result)
	}
	
	if result != f32v4_MaxExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Float32Vec4_Clamp(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_SimpleInput.Clamp(&f32v4_ClampMinInput, &f32v4_ClampMaxInput, &result)
	}

	if result != f32v4_ClampExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// CEIL

func Benchmark_Float32Vec4_Ceil_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_ceilSSE41(&f32v4_CeilInput, &result)
	}
	
	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Ceil_Generic(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_ceilGeneric(&f32v4_CeilInput, &result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Ceil(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_CeilInput.Ceil(&result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// ROUND

func Benchmark_Float32Vec4_Round_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_roundSSE41(&f32v4_RoundInput, &result)
	}
	
	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Round_Generic(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_roundGeneric(&f32v4_RoundInput, &result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Round(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_RoundInput.Round(&result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

// FLOOR

func Benchmark_Float32Vec4_Floor_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_floorSSE41(&f32v4_FloorInput, &result)
	}
	
	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Floor_Generic(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_floorGeneric(&f32v4_FloorInput, &result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec4_Floor(b *testing.B) {
	var result Float32Vec4

	for n := 0; n < b.N; n++ {
		f32v4_FloorInput.Floor(&result)
	}

	if result != f32v4_RoundExpected {
		fmt.Println([4]float32(result))
		b.FailNow()
	}
}