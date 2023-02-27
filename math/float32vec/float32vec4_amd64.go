//go:build amd64

package float32vec

import (
	"golang.org/x/sys/cpu"
	std_math "math"
)

type Float32Vec4 [4]float32

func f32v4_add(a *Float32Vec4, b *Float32Vec4, result *Float32Vec4)
func (a *Float32Vec4) Add(b *Float32Vec4, result *Float32Vec4) {
	f32v4_add(a, b, result)
}

func f32v4_sub(a *Float32Vec4, b *Float32Vec4, result *Float32Vec4)
func (a *Float32Vec4) Sub(b *Float32Vec4, result *Float32Vec4) {
	f32v4_sub(a, b, result)
}

func f32v4_mul(a *Float32Vec4, b *Float32Vec4, result *Float32Vec4)
func (a *Float32Vec4) Mul(b *Float32Vec4, result *Float32Vec4) {
	f32v4_mul(a, b, result)
}

func f32v4_fmaFMA(a *Float32Vec4, b *Float32Vec4, c *Float32Vec4, result *Float32Vec4)
func f32v4_fmaGeneric(a *Float32Vec4, b *Float32Vec4, c *Float32Vec4, result *Float32Vec4) {
	v1 := *a
	v2 := *b
	v3 := *c

	v1[0] = float32(std_math.FMA(float64(v1[0]), float64(v2[0]), float64(v3[0])))
	v1[1] = float32(std_math.FMA(float64(v1[1]), float64(v2[1]), float64(v3[1])))
	v1[2] = float32(std_math.FMA(float64(v1[2]), float64(v2[2]), float64(v3[2])))
	v1[3] = float32(std_math.FMA(float64(v1[3]), float64(v2[3]), float64(v3[3])))
	
	*result = v1
}
func (a *Float32Vec4) FMA(b *Float32Vec4, c *Float32Vec4, result *Float32Vec4) {
	if cpu.X86.HasFMA {
		f32v4_fmaFMA(a, b, c, result)
	} else {
		f32v4_fmaGeneric(a, b, c, result)
	}
}

func f32v4_div(a *Float32Vec4, b *Float32Vec4, result *Float32Vec4)
func (a *Float32Vec4) Div(b *Float32Vec4, result *Float32Vec4) {
	f32v4_div(a, b, result)
}

func f32v4_min(v *Float32Vec4, mn *Float32Vec4, result *Float32Vec4)
func (v *Float32Vec4) Min(mn *Float32Vec4, result *Float32Vec4) {
	f32v4_min(v, mn, result)
}

func f32v4_max(v *Float32Vec4, mx *Float32Vec4, result *Float32Vec4)
func (v *Float32Vec4) Max(mx *Float32Vec4, result *Float32Vec4) {
	f32v4_max(v, mx, result)
}

func f32v4_clamp(v *Float32Vec4, mn *Float32Vec4, mx *Float32Vec4, result *Float32Vec4)
func (v *Float32Vec4) Clamp(mn *Float32Vec4, mx *Float32Vec4, result *Float32Vec4) {
	f32v4_clamp(v, mn, mx, result)
}

func f32v4_ceilSSE41(v *Float32Vec4, result *Float32Vec4)
func f32v4_ceilGeneric(v *Float32Vec4, result *Float32Vec4) {
	v1 := *v

	v1[0] = float32(std_math.Ceil(float64(v1[0])))
	v1[1] = float32(std_math.Ceil(float64(v1[1])))
	v1[2] = float32(std_math.Ceil(float64(v1[2])))
	v1[3] = float32(std_math.Ceil(float64(v1[3])))
	
	*result = v1
}
func (v *Float32Vec4) Ceil(result *Float32Vec4) {
	if cpu.X86.HasSSE41 {
		f32v4_ceilSSE41(v, result)
	} else {
		f32v4_ceilGeneric(v, result)
	}
}

func f32v4_roundSSE41(v *Float32Vec4, result *Float32Vec4)
func f32v4_roundGeneric(v *Float32Vec4, result *Float32Vec4) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	v1[2] = float32(std_math.Round(float64(v1[2])))
	v1[3] = float32(std_math.Round(float64(v1[3])))
	
	*result = v1
}
func (v *Float32Vec4) Round(result *Float32Vec4) {
	if cpu.X86.HasSSE41 {
		f32v4_roundSSE41(v, result)
	} else {
		f32v4_roundGeneric(v, result)
	}
}

func f32v4_floorSSE41(v *Float32Vec4, result *Float32Vec4)
func f32v4_floorGeneric(v *Float32Vec4, result *Float32Vec4) {
	v1 := *v

	v1[0] = float32(std_math.Floor(float64(v1[0])))
	v1[1] = float32(std_math.Floor(float64(v1[1])))
	v1[2] = float32(std_math.Floor(float64(v1[2])))
	v1[3] = float32(std_math.Floor(float64(v1[3])))
	
	*result = v1
}
func (v *Float32Vec4) Floor(result *Float32Vec4) {
	if cpu.X86.HasSSE41 {
		f32v4_floorSSE41(v, result)
	} else {
		f32v4_floorGeneric(v, result)
	}
}
