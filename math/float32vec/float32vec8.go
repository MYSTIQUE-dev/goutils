//go:build !amd64

package float32vec

import (
	"github.com/notsatvrn/goutils/math"
	std_math "math"
)

type Float32Vec8 [8]float32

func (a *Float32Vec8) Add(b *Float32Vec8, result *Float32Vec8) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	v1[3] += v2[3]
	v1[4] += v2[4]
	v1[5] += v2[5]
	v1[6] += v2[6]
	v1[7] += v2[7]
	
	*result = v1
}

func (a *Float32Vec8) Sub(b *Float32Vec8, result *Float32Vec8) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	v1[3] -= v2[3]
	v1[4] -= v2[4]
	v1[5] -= v2[5]
	v1[6] -= v2[6]
	v1[7] -= v2[7]
	
	*result = v1
}

func (a *Float32Vec8) Mul(b *Float32Vec8, result *Float32Vec8) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	v1[3] *= v2[3]
	v1[4] *= v2[4]
	v1[5] *= v2[5]
	v1[6] *= v2[6]
	v1[7] *= v2[7]
	
	*result = v1
}

func (a *Float32Vec8) FMA(b *Float32Vec8, c *Float32Vec8, result *Float32Vec8) {
	v1 := *a
	v2 := *b
	v3 := *c

	v1[0] = float32(std_math.FMA(float64(v1[0]), float64(v2[0]), float64(v3[0])))
	v1[1] = float32(std_math.FMA(float64(v1[1]), float64(v2[1]), float64(v3[1])))
	v1[2] = float32(std_math.FMA(float64(v1[2]), float64(v2[2]), float64(v3[2])))
	v1[3] = float32(std_math.FMA(float64(v1[3]), float64(v2[3]), float64(v3[3])))
	v1[4] = float32(std_math.FMA(float64(v1[4]), float64(v2[4]), float64(v3[4])))
	v1[5] = float32(std_math.FMA(float64(v1[5]), float64(v2[5]), float64(v3[5])))
	v1[6] = float32(std_math.FMA(float64(v1[6]), float64(v2[6]), float64(v3[6])))
	v1[7] = float32(std_math.FMA(float64(v1[7]), float64(v2[7]), float64(v3[7])))
	
	*result = v1
}

func (a *Float32Vec8) Div(b *Float32Vec8, result *Float32Vec8) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	v1[3] /= v2[3]
	v1[4] /= v2[4]
	v1[5] /= v2[5]
	v1[6] /= v2[6]
	v1[7] /= v2[7]
	
	*result = v1
}

func (v *Float32Vec8) Min(mn *Float32Vec8, result *Float32Vec8) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	v1[3] = math.Min(v1[3], v2[3])
	v1[4] = math.Min(v1[4], v2[4])
	v1[5] = math.Min(v1[5], v2[5])
	v1[6] = math.Min(v1[6], v2[6])
	v1[7] = math.Min(v1[7], v2[7])

	*result = v1
}

func (v *Float32Vec8) Max(mx *Float32Vec8, result *Float32Vec8) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	v1[3] = math.Max(v1[3], v2[3])
	v1[4] = math.Max(v1[4], v2[4])
	v1[5] = math.Max(v1[5], v2[5])
	v1[6] = math.Max(v1[6], v2[6])
	v1[7] = math.Max(v1[7], v2[7])

	*result = v1
}

func (v *Float32Vec8) Clamp(mn *Float32Vec8, mx *Float32Vec8, result *Float32Vec8) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[3] = math.Max(v1[3], math.Min(v2[3], v3[3]))
	v1[4] = math.Max(v1[4], math.Min(v2[4], v3[4]))
	v1[5] = math.Max(v1[5], math.Min(v2[5], v3[5]))
	v1[6] = math.Max(v1[6], math.Min(v2[6], v3[6]))
	v1[7] = math.Max(v1[7], math.Min(v2[7], v3[7]))
	
	*result = v1
}

func (v *Float32Vec8) Ceil(result *Float32Vec8) {
	v1 := *v

	v1[0] = float32(std_math.Ceil(float64(v1[0])))
	v1[1] = float32(std_math.Ceil(float64(v1[1])))
	v1[2] = float32(std_math.Ceil(float64(v1[2])))
	v1[3] = float32(std_math.Ceil(float64(v1[3])))
	v1[4] = float32(std_math.Ceil(float64(v1[4])))
	v1[5] = float32(std_math.Ceil(float64(v1[5])))
	v1[6] = float32(std_math.Ceil(float64(v1[6])))
	v1[7] = float32(std_math.Ceil(float64(v1[7])))
	
	*result = v1
}

func (v *Float32Vec8) Round(result *Float32Vec8) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	v1[2] = float32(std_math.Round(float64(v1[2])))
	v1[3] = float32(std_math.Round(float64(v1[3])))
	v1[4] = float32(std_math.Round(float64(v1[4])))
	v1[5] = float32(std_math.Round(float64(v1[5])))
	v1[6] = float32(std_math.Round(float64(v1[6])))
	v1[7] = float32(std_math.Round(float64(v1[7])))
	
	*result = v1
}

func (v *Float32Vec8) Floor(result *Float32Vec8) {
	v1 := *v

	v1[0] = float32(std_math.Floor(float64(v1[0])))
	v1[1] = float32(std_math.Floor(float64(v1[1])))
	v1[2] = float32(std_math.Floor(float64(v1[2])))
	v1[3] = float32(std_math.Floor(float64(v1[3])))
	v1[4] = float32(std_math.Floor(float64(v1[4])))
	v1[5] = float32(std_math.Floor(float64(v1[5])))
	v1[6] = float32(std_math.Floor(float64(v1[6])))
	v1[7] = float32(std_math.Floor(float64(v1[7])))
	
	*result = v1
}
