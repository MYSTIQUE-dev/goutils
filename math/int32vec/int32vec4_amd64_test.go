//go:build amd64

package int32vec

import (
	"golang.org/x/sys/cpu"
	"testing"
	"fmt"
)

// INPUTS

var (
	SimpleInput = Int32Vec4([4]int32{1, 2, 3, 4})
	SimpleInputReverse = Int32Vec4([4]int32{4, 3, 2, 1})
	ClampMaxInput = Int32Vec4([4]int32{3, 3, 3, 3})
	ClampMinInput = Int32Vec4([4]int32{1, 1, 1, 1})
)

// EXPECTED

var (
	AddExpected = Int32Vec4([4]int32{5, 5, 5, 5})
	SubExpected = Int32Vec4([4]int32{-3, -1, 1, 3})
	MulExpected = Int32Vec4([4]int32{4, 6, 6, 4})
	MulAddExpected = Int32Vec4([4]int32{2, 6, 12, 20})
	DivExpected = Int32Vec4([4]int32{1, 1, 1, 1})
	MinExpected = Int32Vec4([4]int32{1, 2, 2, 1})
	MaxExpected = Int32Vec4([4]int32{4, 3, 3, 4})
	ClampExpected = Int32Vec4([4]int32{1, 2, 3, 3})
)

// ADD

func Benchmark_Int32Vec4_Add(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Add(&SimpleInputReverse, &result)
	}
	
	if result != AddExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// SUB

func Benchmark_Int32Vec4_Sub(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Sub(&SimpleInputReverse, &result)
	}
	
	if result != SubExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// MUL

func Benchmark_Int32Vec4_Mul(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Mul(&SimpleInputReverse, &result)
	}
	
	if result != MulExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// DIV

func Benchmark_Int32Vec4_Div(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Div(&SimpleInput, &result)
	}
	
	if result != DivExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// MIN

func Benchmark_Int32Vec4_minSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		minSSE41(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_minGeneric(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		minGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_Min(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Min(&SimpleInputReverse, &result)
	}
	
	if result != MinExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// MAX

func Benchmark_Int32Vec4_maxSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		maxSSE41(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_maxGeneric(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		maxGeneric(&SimpleInput, &SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_Max(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Max(&SimpleInputReverse, &result)
	}
	
	if result != MaxExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

// CLAMP

func Benchmark_Int32Vec4_clampSSE41(b *testing.B) {
	if !cpu.X86.HasSSE41 {
		b.SkipNow();
	}

	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		clampSSE41(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_clampGeneric(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		clampGeneric(&SimpleInput, &ClampMinInput, &ClampMaxInput, &result)
	}
	
	if result != ClampExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}

func Benchmark_Int32Vec4_Clamp(b *testing.B) {
	var result Int32Vec4

	for n := 0; n < b.N; n++ {
		SimpleInput.Clamp(&ClampMinInput, &ClampMaxInput, &result)
	}

	if result != ClampExpected {
		fmt.Println([4]int32(result))
		b.FailNow()
	}
}
