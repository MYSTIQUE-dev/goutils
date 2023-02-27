//go:build amd64

package float32vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	f32v8_SimpleInput = Float32Vec8([8]float32{1, 2, 3, 4, 5, 6, 7, 8})
	f32v8_SimpleInputReverse = Float32Vec8([8]float32{8, 7, 6, 5, 4, 3, 2, 1})
	f32v8_CeilInput = Float32Vec8([8]float32{0.9, 1.9, 2.9, 3.9, 4.9, 5.9, 6.9, 7.9})
	f32v8_RoundInput = Float32Vec8([8]float32{0.9, 2.1, 2.9, 4.1, 4.9, 6.1, 6.9, 8.1})
	f32v8_FloorInput = Float32Vec8([8]float32{1.9, 2.9, 3.9, 4.9, 5.9, 6.9, 7.9, 8.9})
	f32v8_ClampMaxInput = Float32Vec8([8]float32{6, 6, 6, 6, 6, 6, 6, 6})
	f32v8_ClampMinInput = Float32Vec8([8]float32{1, 1, 1, 1, 1, 1, 1, 1})
)

// EXPECTED

var (
	f32v8_AddExpected = Float32Vec8([8]float32{9, 9, 9, 9, 9, 9, 9, 9})
	f32v8_SubExpected = Float32Vec8([8]float32{-7, -5, -3, -1, 1, 3, 5, 7})
	f32v8_MulExpected = Float32Vec8([8]float32{8, 14, 18, 20, 20, 18, 14, 8})
	f32v8_FMAExpected = Float32Vec8([8]float32{2, 6, 12, 20, 30, 42, 56, 72})
	f32v8_DivExpected = Float32Vec8([8]float32{1, 1, 1, 1, 1, 1, 1, 1})
	f32v8_MinExpected = Float32Vec8([8]float32{1, 2, 3, 4, 4, 3, 2, 1})
	f32v8_MaxExpected = Float32Vec8([8]float32{8, 7, 6, 5, 5, 6, 7, 8})
	f32v8_ClampExpected = Float32Vec8([8]float32{1, 2, 3, 4, 5, 6, 6, 6})
	f32v8_RoundExpected = Float32Vec8([8]float32{1, 2, 3, 4, 5, 6, 7, 8})
)

// ADD

func Benchmark_Float32Vec8_Add(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Add(&f32v8_SimpleInputReverse, &result)
	}
	
	if result != f32v8_AddExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Float32Vec8_Sub(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Sub(&f32v8_SimpleInputReverse, &result)
	}
	
	if result != f32v8_SubExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Float32Vec8_Mul(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Mul(&f32v8_SimpleInputReverse, &result)
	}
	
	if result != f32v8_MulExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// FMA

func Benchmark_Float32Vec8_FMA_FMA(b *testing.B) {
	if !cpu.X86.HasFMA {
		b.SkipNow();
	}

	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_fmaFMA(&f32v8_SimpleInput, &f32v8_SimpleInput, &f32v8_SimpleInput, &result)
	}
	
	if result != f32v8_FMAExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_FMA_Generic(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_fmaGeneric(&f32v8_SimpleInput, &f32v8_SimpleInput, &f32v8_SimpleInput, &result)
	}
	
	if result != f32v8_FMAExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_FMA(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.FMA(&f32v8_SimpleInput, &f32v8_SimpleInput, &result)
	}
	
	if result != f32v8_FMAExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Float32Vec8_Div(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Div(&f32v8_SimpleInput, &result)
	}
	
	if result != f32v8_DivExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Float32Vec8_Min(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Min(&f32v8_SimpleInputReverse, &result)
	}
	
	if result != f32v8_MinExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Float32Vec8_Max(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Max(&f32v8_SimpleInputReverse, &result)
	}
	
	if result != f32v8_MaxExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Float32Vec8_Clamp(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_SimpleInput.Clamp(&f32v8_ClampMinInput, &f32v8_ClampMaxInput, &result)
	}

	if result != f32v8_ClampExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// CEIL

func Benchmark_Float32Vec8_Ceil_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_ceilSSE41(&f32v8_CeilInput, &result)
	}
	
	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Ceil_Generic(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_ceilGeneric(&f32v8_CeilInput, &result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Ceil(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_CeilInput.Ceil(&result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// ROUND

func Benchmark_Float32Vec8_Round_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_roundSSE41(&f32v8_RoundInput, &result)
	}
	
	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Round_Generic(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_roundGeneric(&f32v8_RoundInput, &result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Round(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_RoundInput.Round(&result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

// FLOOR

func Benchmark_Float32Vec8_Floor_SSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_floorSSE41(&f32v8_FloorInput, &result)
	}
	
	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Floor_Generic(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_floorGeneric(&f32v8_FloorInput, &result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}

func Benchmark_Float32Vec8_Floor(b *testing.B) {
	var result Float32Vec8

	for n := 0; n < b.N; n++ {
		f32v8_FloorInput.Floor(&result)
	}

	if result != f32v8_RoundExpected {
		fmt.Println([8]float32(result))
		b.FailNow()
	}
}