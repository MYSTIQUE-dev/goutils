//go:build !amd64

package math

import std_math "math"

type DoubleVec3 [3]float64

func (a *DoubleVec3) Add(b *DoubleVec3, result *DoubleVec3) {
	v1 := *a
	v2 := *b

	*result = DoubleVec3([3]float64{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	})
}

func (a *DoubleVec3) Sub(b *DoubleVec3, result *DoubleVec3) {
	v1 := *a
	v2 := *b

	*result = DoubleVec3([3]float64{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
	})
}

func (a *DoubleVec3) Mul(b *DoubleVec3, result *DoubleVec3) {
	v1 := *a
	v2 := *b

	*result = DoubleVec3([3]float64{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
	})
}

func (a *DoubleVec3) MulAdd(b *DoubleVec3, c *DoubleVec3, result *DoubleVec3) {
	v1 := *a
	v2 := *b
	v3 := *c

	*result = DoubleVec3([3]float64{
		v1[0] * v2[0] + v3[0],
		v1[1] * v2[1] + v3[1],
		v1[2] * v2[2] + v3[2],
	})
}

func (a *DoubleVec3) Div(b *DoubleVec3, result *DoubleVec3) {
	v1 := *a
	v2 := *b

	*result = DoubleVec3([3]float64{
		v1[0] / v2[0],
		v1[1] / v2[1],
		v1[2] / v2[2],
	})
}

func (v *DoubleVec3) Min(mn *DoubleVec3, result *DoubleVec3) {
	v1 := *v
	v2 := *mn

	*result = DoubleVec3([3]float64{
		std_math.Min(v1[0], v2[0]),
		std_math.Min(v1[1], v2[1]),
		std_math.Min(v1[2], v2[2]),
	})
}

func (v *DoubleVec3) Max(mx *DoubleVec3, result *DoubleVec3) {
	v1 := *v
	v2 := *mx

	*result = DoubleVec3([3]float64{
		std_math.Max(v1[0], v2[0]),
		std_math.Max(v1[1], v2[1]),
		std_math.Max(v1[2], v2[2]),
	})
}

func (v *DoubleVec3) Clamp(mn *DoubleVec3, mx *DoubleVec3, result *DoubleVec3) {
	var maxResult DoubleVec3
	
	v.Max(mn, &maxResult)
	maxResult.Min(mx, result)
}

func (v *DoubleVec3) Ceil(result *DoubleVec3) {
	v1 := *v

	*result = DoubleVec3([3]float64{
		std_math.Ceil(v1[0]),
		std_math.Ceil(v1[1]),
		std_math.Ceil(v1[2]),
	})
}

func (v *DoubleVec3) Round(result *DoubleVec3) {
	v1 := *v

	*result = DoubleVec3([3]float64{
		std_math.Round(v1[0]),
		std_math.Round(v1[1]),
		std_math.Round(v1[2]),
	})
}

func (v *DoubleVec3) Floor(result *DoubleVec3) {
	v1 := *v

	*result = DoubleVec3([3]float64{
		std_math.Floor(v1[0]),
		std_math.Floor(v1[1]),
		std_math.Floor(v1[2]),
	})
}
