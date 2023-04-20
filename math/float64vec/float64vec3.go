//go:build !amd64

package math

import std_math "math"

type Float64Vec3 [3]float64

func (a *Float64Vec3) Add(b *Float64Vec3, result *Float64Vec3) {
	v1 := *a
	v2 := *b

	*result = Float64Vec3([3]float64{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	})
}

func (a *Float64Vec3) Sub(b *Float64Vec3, result *Float64Vec3) {
	v1 := *a
	v2 := *b

	*result = Float64Vec3([3]float64{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
	})
}

func (a *Float64Vec3) Mul(b *Float64Vec3, result *Float64Vec3) {
	v1 := *a
	v2 := *b

	*result = Float64Vec3([3]float64{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
	})
}

func (a *Float64Vec3) MulAdd(b *Float64Vec3, c *Float64Vec3, result *Float64Vec3) {
	v1 := *a
	v2 := *b
	v3 := *c

	*result = Float64Vec3([3]float64{
		v1[0] * v2[0] + v3[0],
		v1[1] * v2[1] + v3[1],
		v1[2] * v2[2] + v3[2],
	})
}

func (a *Float64Vec3) Div(b *Float64Vec3, result *Float64Vec3) {
	v1 := *a
	v2 := *b

	*result = Float64Vec3([3]float64{
		v1[0] / v2[0],
		v1[1] / v2[1],
		v1[2] / v2[2],
	})
}

func (v *Float64Vec3) Min(mn *Float64Vec3, result *Float64Vec3) {
	v1 := *v
	v2 := *mn

	*result = Float64Vec3([3]float64{
		std_math.Min(v1[0], v2[0]),
		std_math.Min(v1[1], v2[1]),
		std_math.Min(v1[2], v2[2]),
	})
}

func (v *Float64Vec3) Max(mx *Float64Vec3, result *Float64Vec3) {
	v1 := *v
	v2 := *mx

	*result = Float64Vec3([3]float64{
		std_math.Max(v1[0], v2[0]),
		std_math.Max(v1[1], v2[1]),
		std_math.Max(v1[2], v2[2]),
	})
}

func (v *Float64Vec3) Clamp(mn *Float64Vec3, mx *Float64Vec3, result *Float64Vec3) {
	var maxResult Float64Vec3
	
	v.Max(mn, &maxResult)
	maxResult.Min(mx, result)
}

func (v *Float64Vec3) Ceil(result *Float64Vec3) {
	v1 := *v

	*result = Float64Vec3([3]float64{
		std_math.Ceil(v1[0]),
		std_math.Ceil(v1[1]),
		std_math.Ceil(v1[2]),
	})
}

func (v *Float64Vec3) Round(result *Float64Vec3) {
	v1 := *v

	*result = Float64Vec3([3]float64{
		std_math.Round(v1[0]),
		std_math.Round(v1[1]),
		std_math.Round(v1[2]),
	})
}

func (v *Float64Vec3) Floor(result *Float64Vec3) {
	v1 := *v

	*result = Float64Vec3([3]float64{
		std_math.Floor(v1[0]),
		std_math.Floor(v1[1]),
		std_math.Floor(v1[2]),
	})
}
