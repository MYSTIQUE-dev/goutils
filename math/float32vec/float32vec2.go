package float32vec

import (
	"github.com/notsatvrn/goutils/math"
	std_math "math"
)

type Float32Vec2 [2]float32

func (a *Float32Vec2) Add(b *Float32Vec2, result *Float32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	
	*result = v1
}

func (a *Float32Vec2) Sub(b *Float32Vec2, result *Float32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	
	*result = v1
}

func (a *Float32Vec2) Mul(b *Float32Vec2, result *Float32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	
	*result = v1
}

func (a *Float32Vec2) FMA(b *Float32Vec2, c *Float32Vec2, result *Float32Vec2) {
	v1 := *a
	v2 := *b
	v3 := *c

	v1[0] = float32(std_math.FMA(float64(v1[0]), float64(v2[0]), float64(v3[0])))
	v1[1] = float32(std_math.FMA(float64(v1[1]), float64(v2[1]), float64(v3[1])))
	
	*result = v1
}

func (a *Float32Vec2) Div(b *Float32Vec2, result *Float32Vec2) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	
	*result = v1
}

func (v *Float32Vec2) Min(mn *Float32Vec2, result *Float32Vec2) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	
	*result = v1
}

func (v *Float32Vec2) Max(mx *Float32Vec2, result *Float32Vec2) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	
	*result = v1
}

func (v *Float32Vec2) Clamp(mn *Float32Vec2, mx *Float32Vec2, result *Float32Vec2) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	
	*result = v1
}

func (v *Float32Vec2) Ceil(result *Float32Vec2) {
	v1 := *v

	v1[0] = float32(std_math.Ceil(float64(v1[0])))
	v1[1] = float32(std_math.Ceil(float64(v1[1])))
	
	*result = v1
}

func (v *Float32Vec2) Round(result *Float32Vec2) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	
	*result = v1
}

func (v *Float32Vec2) Floor(result *Float32Vec2) {
	v1 := *v

	v1[0] = float32(std_math.Floor(float64(v1[0])))
	v1[1] = float32(std_math.Floor(float64(v1[1])))
	
	*result = v1
}
