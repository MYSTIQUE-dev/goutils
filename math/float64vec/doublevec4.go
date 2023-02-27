//go:build !amd64

package math

import std_math "math"

type DoubleVec4 [4]float64

func (a *DoubleVec4) Add(b *DoubleVec4, result *DoubleVec4) {
	v1 := *a
	v2 := *b

	*result = DoubleVec4([4]float64{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
		v1[3] + v2[3],
	})
}

func (a *DoubleVec4) Sub(b *DoubleVec4, result *DoubleVec4) {
	v1 := *a
	v2 := *b

	*result = DoubleVec4([4]float64{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
		v1[3] - v2[3],
	})
}

func (a *DoubleVec4) Mul(b *DoubleVec4, result *DoubleVec4) {
	v1 := *a
	v2 := *b

	*result = DoubleVec4([4]float64{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
		v1[3] * v2[3],
	})
}

func (a *DoubleVec4) MulAdd(b *DoubleVec4, c *DoubleVec4, result *DoubleVec4) {
	v1 := *a
	v2 := *b
	v3 := *c

	*result = DoubleVec4([4]float64{
		v1[0] * v2[0] + v3[0],
		v1[1] * v2[1] + v3[1],
		v1[2] * v2[2] + v3[2],
		v1[3] * v2[3] + v3[3],
	})
}

func (a *DoubleVec4) Div(b *DoubleVec4, result *DoubleVec4) {
	v1 := *a
	v2 := *b

	*result = DoubleVec4([4]float64{
		v1[0] / v2[0],
		v1[1] / v2[1],
		v1[2] / v2[2],
		v1[3] / v2[3],
	})
}

func (v *DoubleVec4) Min(mn *DoubleVec4, result *DoubleVec4) {
	v1 := *v
	v2 := *mn

	*result = DoubleVec4([4]float64{
		std_math.Min(v1[0], v2[0]),
		std_math.Min(v1[1], v2[1]),
		std_math.Min(v1[2], v2[2]),
		std_math.Min(v1[3], v2[3]),
	})
}

func (v *DoubleVec4) Max(mx *DoubleVec4, result *DoubleVec4) {
	v1 := *v
	v2 := *mx

	*result = DoubleVec4([4]float64{
		std_math.Max(v1[0], v2[0]),
		std_math.Max(v1[1], v2[1]),
		std_math.Max(v1[2], v2[2]),
		std_math.Max(v1[3], v2[3]),
	})
}

func (v *DoubleVec4) Clamp(mn *DoubleVec4, mx *DoubleVec4, result *DoubleVec4) {
	var maxResult DoubleVec4
	
	v.Max(mn, &maxResult)
	maxResult.Min(mx, result)
}

func (v *DoubleVec4) Ceil(result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Ceil(v1[0]),
		std_math.Ceil(v1[1]),
		std_math.Ceil(v1[2]),
		std_math.Ceil(v1[3]),
	})
}

func (v *DoubleVec4) Round(result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Round(v1[0]),
		std_math.Round(v1[1]),
		std_math.Round(v1[2]),
		std_math.Round(v1[3]),
	})
}

func (v *DoubleVec4) Floor(result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Floor(v1[0]),
		std_math.Floor(v1[1]),
		std_math.Floor(v1[2]),
		std_math.Floor(v1[3]),
	})
}
