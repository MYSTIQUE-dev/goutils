package bytevec

import "github.com/notsatvrn/goutils/math"

type ByteVec8 [8]byte

func (a *ByteVec8) Add(b *ByteVec8, result *ByteVec8) {
	v1 := *a
	v2 := *b
	
	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	v1[3] += v2[3]
	v1[4] += v2[4]
	v1[5] += v2[5]
	v1[6] += v2[6]
	v1[7] += v2[7]
	
	*result = v1
}

func (a *ByteVec8) Sub(b *ByteVec8, result *ByteVec8) {
	v1 := *a
	v2 := *b
	
	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	v1[3] -= v2[3]
	v1[4] -= v2[4]
	v1[5] -= v2[5]
	v1[6] -= v2[6]
	v1[7] -= v2[7]
	
	*result = v1
}

func (a *ByteVec8) Mul(b *ByteVec8, result *ByteVec8) {
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
	
	*result = v1
}

func (a *ByteVec8) Div(b *ByteVec8, result *ByteVec8) {
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
	
	*result = v1
}

func (v *ByteVec8) Min(mn *ByteVec8, result *ByteVec8) {
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
	
	*result = v1
}

func (v *ByteVec8) Max(mx *ByteVec8, result *ByteVec8) {
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
	
	*result = v1
}

func (v *ByteVec8) Clamp(mn *ByteVec8, mx *ByteVec8, result *ByteVec8) {
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

	*result = v1
}
