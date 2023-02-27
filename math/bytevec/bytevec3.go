package bytevec

import "github.com/notsatvrn/goutils/math"

type ByteVec3 [3]byte

func (a *ByteVec3) Add(b *ByteVec3, result *ByteVec3) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	
	*result = v1
}

func (a *ByteVec3) Sub(b *ByteVec3, result *ByteVec3) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	
	*result = v1
}

func (a *ByteVec3) Mul(b *ByteVec3, result *ByteVec3) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	
	*result = v1
}

func (a *ByteVec3) Div(b *ByteVec3, result *ByteVec3) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	
	*result = v1
}

func (v *ByteVec3) Min(mn *ByteVec3, result *ByteVec3) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	
	*result = v1
}

func (v *ByteVec3) Max(mx *ByteVec3, result *ByteVec3) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	
	*result = v1
}

func (v *ByteVec3) Clamp(mn *ByteVec3, mx *ByteVec3, result *ByteVec3) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	
	*result = v1
}
