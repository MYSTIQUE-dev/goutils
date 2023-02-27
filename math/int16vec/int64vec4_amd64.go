//go:build amd64

package int64vec

import (
	"github.com/notsatvrn/goutils/math"
	"golang.org/x/sys/cpu"
)

type Int64Vec4 [4]int64

func int64v4_AddAVX2(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4)
func int64v4_AddSSE2(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4)
func (a *Int64Vec4) Add(b *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX2 {
		int64v4_AddAVX2(a, b, result)
	} else {
		int64v4_AddSSE2(a, b, result)
	}
}

func int64v4_SubAVX2(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4)
func int64v4_SubSSE2(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4)
func (a *Int64Vec4) Sub(b *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX2 {
		int64v4_SubAVX2(a, b, result)
	} else {
		int64v4_SubSSE2(a, b, result)
	}
}

func int64v4_MulAVX512DQ_VL(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4)
func int64v4_MulGeneric(a *Int64Vec4, b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b
	
	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	v1[3] *= v2[3]
	
	*result = v1
}
func (a *Int64Vec4) Mul(b *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX512DQ && cpu.X86.HasAVX512VL {
		int64v4_MulAVX512DQ_VL(a, b, result)
	} else {
		int64v4_MulGeneric(a, b, result)
	}
}

func (a *Int64Vec4) Div(b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b
	
	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	v1[3] /= v2[3]
	
	*result = v1
}

func int64v4_MinAVX512F_VL(v *Int64Vec4, mn *Int64Vec4, result *Int64Vec4)
func int64v4_MinGeneric(v *Int64Vec4, mn *Int64Vec4, result *Int64Vec4) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	v1[3] = math.Min(v1[3], v2[3])
	
	*result = v1
}
func (v *Int64Vec4) Min(mn *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL {
		int64v4_MinAVX512F_VL(v, mn, result)
	} else {
		int64v4_MinGeneric(v, mn, result)
	}
}

func int64v4_MaxAVX512F_VL(v *Int64Vec4, mx *Int64Vec4, result *Int64Vec4)
func int64v4_MaxGeneric(v *Int64Vec4, mx *Int64Vec4, result *Int64Vec4) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	v1[3] = math.Max(v1[3], v2[3])
	
	*result = v1
}
func (v *Int64Vec4) Max(mx *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL {
		int64v4_MaxAVX512F_VL(v, mx, result)
	} else {
		int64v4_MaxGeneric(v, mx, result)
	}
}

func int64v4_ClampAVX512F_VL(v *Int64Vec4, mn *Int64Vec4, mx *Int64Vec4, result *Int64Vec4)
func int64v4_ClampGeneric(v *Int64Vec4, mn *Int64Vec4, mx *Int64Vec4, result *Int64Vec4) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[0] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[0] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[0] = math.Max(v1[3], math.Min(v2[3], v3[3]))

	*result = v1
}
func (v *Int64Vec4) Clamp(mn *Int64Vec4, mx *Int64Vec4, result *Int64Vec4) {
	if cpu.X86.HasAVX512F && cpu.X86.HasAVX512VL {
		int64v4_ClampAVX512F_VL(v, mn, mx, result)
	} else {
		int64v4_ClampGeneric(v, mn, mx, result)
	}
}
