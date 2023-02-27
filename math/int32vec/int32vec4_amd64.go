//go:build amd64

// ONLY Min, Max, and Clamp use SIMD.
// Some functions were impossible to implement with SIMD. Others had worse performance.
// Min, Max, and Clamp were the only ones possible to implement that had real gains.

package int32vec

import (
	"github.com/notsatvrn/goutils/math"
	"golang.org/x/sys/cpu"
)

type Int32Vec4 [4]int32

func (a *Int32Vec4) Add(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	v1[3] += v2[3]
	
	*result = v1
}

func (a *Int32Vec4) Sub(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	v1[3] -= v2[3]
	
	*result = v1
}

func (a *Int32Vec4) Mul(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	v1[3] *= v2[3]
	
	*result = v1
}

func (a *Int32Vec4) Div(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	v1[3] /= v2[3]
	
	*result = v1
}

func minSSE41(v *Int32Vec4, mn *Int32Vec4, result *Int32Vec4)
func minGeneric(v *Int32Vec4, mn *Int32Vec4, result *Int32Vec4) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	v1[3] = math.Min(v1[3], v2[3])
	
	*result = v1
}
func (v *Int32Vec4) Min(mn *Int32Vec4, result *Int32Vec4) {
	if cpu.X86.HasSSE41 {
		minSSE41(v, mn, result)
	} else {
		minGeneric(v, mn, result)
	}
}

func maxSSE41(v *Int32Vec4, mx *Int32Vec4, result *Int32Vec4)
func maxGeneric(v *Int32Vec4, mx *Int32Vec4, result *Int32Vec4) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	v1[3] = math.Max(v1[3], v2[3])
	
	*result = v1
}
func (v *Int32Vec4) Max(mx *Int32Vec4, result *Int32Vec4) {
	if cpu.X86.HasSSE41 {
		maxSSE41(v, mx, result)
	} else {
		maxGeneric(v, mx, result)
	}
}

func clampSSE41(v *Int32Vec4, mn *Int32Vec4, mx *Int32Vec4, result *Int32Vec4)
func clampGeneric(v *Int32Vec4, mn *Int32Vec4, mx *Int32Vec4, result *Int32Vec4) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[3] = math.Max(v1[3], math.Min(v2[3], v3[3]))

	*result = v1
}
func (v *Int32Vec4) Clamp(mn *Int32Vec4, mx *Int32Vec4, result *Int32Vec4){
	if cpu.X86.HasSSE41 {
		clampSSE41(v, mn, mx, result)
	} else {
		clampGeneric(v, mn, mx, result)
	}
}
