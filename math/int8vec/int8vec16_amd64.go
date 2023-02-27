//go:build amd64

package int8vec

import (
	"github.com/notsatvrn/goutils/math"
	"golang.org/x/sys/cpu"
)

type Int8Vec16 [16]int8

func add(a *Int8Vec16, b *Int8Vec16, result *Int8Vec16)
func (a *Int8Vec16) Add(b *Int8Vec16, result *Int8Vec16) {
	add(a, b, result)
}

func sub(a *Int8Vec16, b *Int8Vec16, result *Int8Vec16)
func (a *Int8Vec16) Sub(b *Int8Vec16, result *Int8Vec16) {
	sub(a, b, result)
}

func (a *Int8Vec16) Mul(b *Int8Vec16, result *Int8Vec16) {
	v1 := *a
	v2 := *b
	
	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	v1[3] *= v2[3]
	v1[4] *= v2[4]
	v1[5] *= v2[5]
	v1[6] *= v2[6]
	v1[7] *= v2[7]
	v1[8] *= v2[8]
	v1[9] *= v2[9]
	v1[10] *= v2[10]
	v1[11] *= v2[11]
	v1[12] *= v2[12]
	v1[13] *= v2[13]
	v1[14] *= v2[14]
	v1[15] *= v2[15]
	
	*result = v1
}

func (a *Int8Vec16) Div(b *Int8Vec16, result *Int8Vec16) {
	v1 := *a
	v2 := *b
	
	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	v1[3] /= v2[3]
	v1[4] /= v2[4]
	v1[5] /= v2[5]
	v1[6] /= v2[6]
	v1[7] /= v2[7]
	v1[8] /= v2[8]
	v1[9] /= v2[9]
	v1[10] /= v2[10]
	v1[11] /= v2[11]
	v1[12] /= v2[12]
	v1[13] /= v2[13]
	v1[14] /= v2[14]
	v1[15] /= v2[15]
	
	*result = v1
}

func minSSE41(v *Int8Vec16, mn *Int8Vec16, result *Int8Vec16)
func minGeneric(v *Int8Vec16, mn *Int8Vec16, result *Int8Vec16) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	v1[3] = math.Min(v1[3], v2[3])
	v1[4] = math.Min(v1[4], v2[4])
	v1[5] = math.Min(v1[5], v2[5])
	v1[6] = math.Min(v1[6], v2[6])
	v1[7] = math.Min(v1[7], v2[7])
	v1[8] = math.Min(v1[8], v2[8])
	v1[9] = math.Min(v1[9], v2[9])
	v1[10] = math.Min(v1[10], v2[10])
	v1[11] = math.Min(v1[11], v2[11])
	v1[12] = math.Min(v1[12], v2[12])
	v1[13] = math.Min(v1[13], v2[13])
	v1[14] = math.Min(v1[14], v2[14])
	v1[15] = math.Min(v1[15], v2[15])

	*result = v1
}
func (v *Int8Vec16) Min(mn *Int8Vec16, result *Int8Vec16) {
	if cpu.X86.HasSSE41 {
		minSSE41(v, mn, result)
	} else {
		minGeneric(v, mn, result)
	}
}

func maxSSE41(v *Int8Vec16, mx *Int8Vec16, result *Int8Vec16)
func maxGeneric(v *Int8Vec16, mx *Int8Vec16, result *Int8Vec16) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	v1[3] = math.Max(v1[3], v2[3])
	v1[4] = math.Max(v1[4], v2[4])
	v1[5] = math.Max(v1[5], v2[5])
	v1[6] = math.Max(v1[6], v2[6])
	v1[7] = math.Max(v1[7], v2[7])
	v1[8] = math.Max(v1[8], v2[8])
	v1[9] = math.Max(v1[9], v2[9])
	v1[10] = math.Max(v1[10], v2[10])
	v1[11] = math.Max(v1[11], v2[11])
	v1[12] = math.Max(v1[12], v2[12])
	v1[13] = math.Max(v1[13], v2[13])
	v1[14] = math.Max(v1[14], v2[14])
	v1[15] = math.Max(v1[15], v2[15])

	*result = v1
}
func (v *Int8Vec16) Max(mx *Int8Vec16, result *Int8Vec16) {
	if cpu.X86.HasSSE41 {
		maxSSE41(v, mx, result)
	} else {
		maxGeneric(v, mx, result)
	}
}

func clampSSE41(v *Int8Vec16, mn *Int8Vec16, mx *Int8Vec16, result *Int8Vec16)
func clampGeneric(v *Int8Vec16, mn *Int8Vec16, mx *Int8Vec16, result *Int8Vec16) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[3] = math.Max(v1[3], math.Min(v2[3], v3[3]))
	v1[4] = math.Max(v1[4], math.Min(v2[4], v3[4]))
	v1[5] = math.Max(v1[5], math.Min(v2[5], v3[5]))
	v1[6] = math.Max(v1[6], math.Min(v2[6], v3[6]))
	v1[7] = math.Max(v1[7], math.Min(v2[7], v3[7]))
	v1[8] = math.Max(v1[8], math.Min(v2[8], v3[8]))
	v1[9] = math.Max(v1[9], math.Min(v2[9], v3[9]))
	v1[10] = math.Max(v1[10], math.Min(v2[10], v3[10]))
	v1[11] = math.Max(v1[11], math.Min(v2[11], v3[11]))
	v1[12] = math.Max(v1[12], math.Min(v2[12], v3[12]))
	v1[13] = math.Max(v1[13], math.Min(v2[13], v3[13]))
	v1[14] = math.Max(v1[14], math.Min(v2[14], v3[14]))
	v1[15] = math.Max(v1[15], math.Min(v2[15], v3[15]))

	*result = v1
}
func (v *Int8Vec16) Clamp(mn *Int8Vec16, mx *Int8Vec16, result *Int8Vec16) {
	if cpu.X86.HasSSE41 {
		clampSSE41(v, mn, mx, result)
	} else {
		clampGeneric(v, mn, mx, result)
	}
}
