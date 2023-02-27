package int64vec

import "github.com/notsatvrn/goutils/math"

type Int64Vec2 [2]int64

func (a *Int64Vec2) Add(b *Int64Vec2, result *Int64Vec2) {
	v1 := *a
	v2 := *b
	
	v1[0] += v2[0]
	v1[1] += v2[1]
	
	*result = v1
}

func (a *Int64Vec2) Sub(b *Int64Vec2, result *Int64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	
	*result = v1
}

func (a *Int64Vec2) Mul(b *Int64Vec2, result *Int64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	
	*result = v1
}

func (a *Int64Vec2) Div(b *Int64Vec2, result *Int64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	
	*result = v1
}

func (v *Int64Vec2) Min(mn *Int64Vec2, result *Int64Vec2) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	
	*result = v1
}

func (v *Int64Vec2) Max(mx *Int64Vec2, result *Int64Vec2) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	
	*result = v1
}

func (v *Int64Vec2) Clamp(mn *Int64Vec2, mx *Int64Vec2, result *Int64Vec2) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	
	*result = v1
}
