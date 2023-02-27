package bytevec

import "github.com/notsatvrn/goutils/math"

type ByteVec2 [2]byte

func (a *ByteVec2) Add(b *ByteVec2, result *ByteVec2) {
	v1 := *a
	v2 := *b
	
	v1[0] += v2[0]
	v1[1] += v2[1]
	
	*result = v1
}

func (a *ByteVec2) Sub(b *ByteVec2, result *ByteVec2) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	
	*result = v1
}

func (a *ByteVec2) Mul(b *ByteVec2, result *ByteVec2) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	
	*result = v1
}

func (a *ByteVec2) Div(b *ByteVec2, result *ByteVec2) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	
	*result = v1
}

func (v *ByteVec2) Min(mn *ByteVec2, result *ByteVec2) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	
	*result = v1
}

func (v *ByteVec2) Max(mx *ByteVec2, result *ByteVec2) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	
	*result = v1
}

func (v *ByteVec2) Clamp(mn *ByteVec2, mx *ByteVec2, result *ByteVec2) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	
	*result = v1
}
