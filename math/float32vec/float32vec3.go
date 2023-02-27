package float32vec

import (
	"github.com/notsatvrn/goutils/math"
	std_math "math"
)

type Float32Vec3 [3]float32

func (a *Float32Vec3) Add(b *Float32Vec3, result *Float32Vec3) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	
	*result = v1
}

func (a *Float32Vec3) Sub(b *Float32Vec3, result *Float32Vec3) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	
	*result = v1
}

func (a *Float32Vec3) Mul(b *Float32Vec3, result *Float32Vec3) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	
	*result = v1
}

func (a *Float32Vec3) FMA(b *Float32Vec3, c *Float32Vec3, result *Float32Vec3) {
	v1 := *a
	v2 := *b
	v3 := *c

	v1[0] = float32(std_math.FMA(float64(v1[0]), float64(v2[0]), float64(v3[0])))
	v1[1] = float32(std_math.FMA(float64(v1[1]), float64(v2[1]), float64(v3[1])))
	v1[2] = float32(std_math.FMA(float64(v1[2]), float64(v2[2]), float64(v3[2])))
	
	*result = v1
}

func (a *Float32Vec3) Div(b *Float32Vec3, result *Float32Vec3) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	
	*result = v1
}

func (v *Float32Vec3) Min(mn *Float32Vec3, result *Float32Vec3) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	
	*result = v1
}

func (v *Float32Vec3) Max(mx *Float32Vec3, result *Float32Vec3) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	
	*result = v1
}

func (v *Float32Vec3) Clamp(mn *Float32Vec3, mx *Float32Vec3, result *Float32Vec3) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	
	*result = v1
}

func (v *Float32Vec3) Ceil(result *Float32Vec3) {
	v1 := *v

	v1[0] = float32(std_math.Ceil(float64(v1[0])))
	v1[1] = float32(std_math.Ceil(float64(v1[1])))
	v1[2] = float32(std_math.Ceil(float64(v1[2])))
	
	*result = v1
}

func (v *Float32Vec3) Round(result *Float32Vec3) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	v1[2] = float32(std_math.Round(float64(v1[2])))
	
	*result = v1
}

func (v *Float32Vec3) Floor(result *Float32Vec3) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	v1[2] = float32(std_math.Round(float64(v1[2])))
	
	*result = v1
}
