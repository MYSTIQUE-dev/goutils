//go:build amd64

package int8vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	SimpleInput = Int8Vec16([16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	SimpleInputReverse = Int8Vec16([16]int8{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	ClampMaxInput = Int8Vec16([16]int8{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	ClampMinInput = Int8Vec16([16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
)

// EXPECTED

var (
	AddExpected = Int8Vec16([16]int8{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17})
	SubExpected = Int8Vec16([16]int8{-15, -13, -11, -9, -7, -5, -3, -1, 1, 3, 5, 7, 9, 11, 13, 15})
	MulExpected = Int8Vec16([16]int8{16, 30, 42, 52, 60, 66, 70, 72, 72, 70, 66, 60, 52, 42, 30, 16})
	DivExpected = Int8Vec16([16]int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	MinExpected = Int8Vec16([16]int8{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	MaxExpected = Int8Vec16([16]int8{16, 15, 14, 13, 12, 11, 10, 9, 9, 10, 11, 12, 13, 14, 15, 16})
	ClampExpected = Int8Vec16([16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9, 9, 9})
)

// ADD

func Benchmark_Int8Vec16_Add(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Add(&SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Int8Vec16_Sub(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Sub(&SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Int8Vec16_Mul(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Mul(&SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Int8Vec16_Div(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Div(&SimpleInput, &result)
	}
	
	if result != DivExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Int8Vec16_minSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow()
	}

	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		minSSE41(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_minGeneric(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		minGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_Min(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Min(&SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Int8Vec16_maxSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow()
	}

	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		maxSSE41(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_maxGeneric(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		maxGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_Max(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Max(&SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Int8Vec16_clampSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow()
	}

	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		clampSSE41(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_clampGeneric(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		clampGeneric(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}

func Benchmark_Int8Vec16_Clamp(b *testing.B) {
	var result Int8Vec16

	for n := 0; n < b.N; n++ {
		SimpleInput.Clamp(&ClampMinInput, &ClampMaxInput, &result)
	}

	if result != ClampExpected {
		fmt.Println([16]int8(result))
		b.FailNow()
	}
}
