package int32vec

import "github.com/notsatvrn/goutils/math"

type Int32Vec2 [2]int32

func (a *Int32Vec2) Add(b *Int32Vec2, result *Int32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	
	*result = v1
}

func (a *Int32Vec2) Sub(b *Int32Vec2, result *Int32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	
	*result = v1
}

func (a *Int32Vec2) Mul(b *Int32Vec2, result *Int32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	
	*result = v1
}

func (a *Int32Vec2) Div(b *Int32Vec2, result *Int32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	
	*result = v1
}

func (v *Int32Vec2) Min(mn *Int32Vec2, result *Int32Vec2) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	
	*result = v1
}

func (v *Int32Vec2) Max(mx *Int32Vec2, result *Int32Vec2) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	
	*result = v1
}

func (v *Int32Vec2) Clamp(mn *Int32Vec2, mx *Int32Vec2, result *Int32Vec2) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	
	*result = v1
}
