package math

import std_math "math"

type Float64Vec2 [2]float64

func (a *Float64Vec2) Add(b *Float64Vec2, result *Float64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	
	*result = v1
}

func (a *Float64Vec2) Sub(b *Float64Vec2, result *Float64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	
	*result = v1
}

func (a *Float64Vec2) Mul(b *Float64Vec2, result *Float64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	
	*result = v1
}

func (a *Float64Vec2) FMA(b *Float64Vec2, c *Float64Vec2, result *Float64Vec2) {
	v1 := *a
	v2 := *b
	v3 := *c

	v1[0] = std_math.FMA(v1[0], v2[0], v3[0])
	v1[1] = std_math.FMA(v1[1], v2[1], v3[1])
	
	*result = v1
}

func (a *Float64Vec2) Div(b *Float64Vec2, result *Float64Vec2) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	
	*result = v1
}

func (v *Float64Vec2) Min(mn *Float64Vec2, result *Float64Vec2) {
	v1 := *v
	v2 := *mn

	v1[0] = std_math.Min(v1[0], v2[0])
	v1[1] = std_math.Min(v1[1], v2[1])
	
	*result = v1
}

func (v *Float64Vec2) Max(mx *Float64Vec2, result *Float64Vec2) {
	v1 := *v
	v2 := *mn

	v1[0] = std_math.Max(v1[0], v2[0])
	v1[1] = std_math.Max(v1[1], v2[1])
	
	*result = v1
}

func (v *Float64Vec2) Clamp(mn *Float64Vec2, mx *Float64Vec2, result *Float64Vec2) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = std_math.Max(v1[0], std_math.Min(v2[0], v3[0]))
	v1[1] = std_math.Max(v1[1], std_math.Min(v2[1], v3[1]))
	
	*result = v1
}

func (v *Float64Vec2) Ceil(result *Float64Vec2) {
	v1 := *v

	v1[0] = std_math.Ceil(v1[0])
	v1[1] = std_math.Ceil(v1[1])
	
	*result = v1
}

func (v *Float64Vec2) Round(result *Float64Vec2) {
	v1 := *v

	v1[0] = std_math.Round(v1[0])
	v1[1] = std_math.Round(v1[1])
	
	*result = v1
}

func (v *Float64Vec2) Floor(result *Float64Vec2) {
	v1 := *v

	v1[0] = std_math.Floor(v1[0])
	v1[1] = std_math.Floor(v1[1])
	
	*result = v1
}
