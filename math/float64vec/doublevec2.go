package math

import std_math "math"

type DoubleVec2 [2]float64

func (a *DoubleVec2) Add(b *DoubleVec2, result *DoubleVec2) {
	v1 := *a
	v2 := *b

	*result = DoubleVec2([2]float64{
		v1[0] + v2[0],
		v1[1] + v2[1],
	})
}

func (a *DoubleVec2) Sub(b *DoubleVec2, result *DoubleVec2) {
	v1 := *a
	v2 := *b

	*result = DoubleVec2([2]float64{
		v1[0] - v2[0],
		v1[1] - v2[1],
	})
}

func (a *DoubleVec2) Mul(b *DoubleVec2, result *DoubleVec2) {
	v1 := *a
	v2 := *b

	*result = DoubleVec2([2]float64{
		v1[0] * v2[0],
		v1[1] * v2[1],
	})
}

func (a *DoubleVec2) MulAdd(b *DoubleVec2, c *DoubleVec2, result *DoubleVec2) {
	v1 := *a
	v2 := *b
	v3 := *c

	*result = DoubleVec2([2]float64{
		v1[0] * v2[0] + v3[0],
		v1[1] * v2[1] + v3[1],
	})
}

func (a *DoubleVec2) Div(b *DoubleVec2, result *DoubleVec2) {
	v1 := *a
	v2 := *b

	*result = DoubleVec2([2]float64{
		v1[0] / v2[0],
		v1[1] / v2[1],
	})
}

func (v *DoubleVec2) Min(mn *DoubleVec2, result *DoubleVec2) {
	v1 := *v
	v2 := *mn

	*result = DoubleVec2([2]float64{
		std_math.Min(v1[0], v2[0]),
		std_math.Min(v1[1], v2[1]),
	})
}

func (v *DoubleVec2) Max(mx *DoubleVec2, result *DoubleVec2) {
	v1 := *v
	v2 := *mx

	*result = DoubleVec2([2]float64{
		std_math.Max(v1[0], v2[0]),
		std_math.Max(v1[1], v2[1]),
	})
}

func (v *DoubleVec2) Clamp(mn *DoubleVec2, mx *DoubleVec2, result *DoubleVec2) {
	var maxResult DoubleVec2
	
	v.Max(mn, &maxResult)
	maxResult.Min(mx, result)
}

func (v *DoubleVec2) Ceil(result *DoubleVec2) {
	v1 := *v

	*result = DoubleVec2([2]float64{
		std_math.Ceil(v1[0]),
		std_math.Ceil(v1[1]),
	})
}

func (v *DoubleVec2) Round(result *DoubleVec2) {
	v1 := *v

	*result = DoubleVec2([2]float64{
		std_math.Round(v1[0]),
		std_math.Round(v1[1]),
	})
}

func (v *DoubleVec2) Floor(result *DoubleVec2) {
	v1 := *v

	*result = DoubleVec2([2]float64{
		std_math.Floor(v1[0]),
		std_math.Floor(v1[1]),
	})
}
