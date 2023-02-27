//go:build amd64

package float32vec

import (
	"golang.org/x/sys/cpu"
	std_math "math"
)

type Float32Vec16 [16]float32

func f32v16_addAVX512F(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_addAVX(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_addSSE2(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func (a *Float32Vec16) Add(b *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_addAVX512F(a, b, result)
	case cpu.X86.HasAVX: f32v16_addAVX(a, b, result)
	default: f32v16_addSSE2(a, b, result)
	}
}

func f32v16_subAVX512F(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_subAVX(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_subSSE2(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func (a *Float32Vec16) Sub(b *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_subAVX512F(a, b, result)
	case cpu.X86.HasAVX: f32v16_subAVX(a, b, result)
	default: f32v16_subSSE2(a, b, result)
	}
}

func f32v16_mulAVX512F(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_mulAVX(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_mulSSE2(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func (a *Float32Vec16) Mul(b *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_mulAVX512F(a, b, result)
	case cpu.X86.HasAVX: f32v16_mulAVX(a, b, result)
	default: f32v16_mulSSE2(a, b, result)
	}
}

func f32v16_fmaFMA_AVX512F(a *Float32Vec16, b *Float32Vec16, c *Float32Vec16, result *Float32Vec16)
func f32v16_fmaFMA_AVX(a *Float32Vec16, b *Float32Vec16, c *Float32Vec16, result *Float32Vec16)
func f32v16_fmaGeneric(a *Float32Vec16, b *Float32Vec16, c *Float32Vec16, result *Float32Vec16) {
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
	v1[8] = float32(std_math.FMA(float64(v1[8]), float64(v2[8]), float64(v3[8])))
	v1[9] = float32(std_math.FMA(float64(v1[9]), float64(v2[9]), float64(v3[9])))
	v1[10] = float32(std_math.FMA(float64(v1[10]), float64(v2[10]), float64(v3[10])))
	v1[11] = float32(std_math.FMA(float64(v1[11]), float64(v2[11]), float64(v3[11])))
	v1[12] = float32(std_math.FMA(float64(v1[12]), float64(v2[12]), float64(v3[12])))
	v1[13] = float32(std_math.FMA(float64(v1[13]), float64(v2[13]), float64(v3[13])))
	v1[14] = float32(std_math.FMA(float64(v1[14]), float64(v2[14]), float64(v3[14])))
	v1[15] = float32(std_math.FMA(float64(v1[15]), float64(v2[15]), float64(v3[15])))
	
	*result = v1
}
func (a *Float32Vec16) FMA(b *Float32Vec16, c *Float32Vec16, result *Float32Vec16) {
	if cpu.X86.HasFMA {
		f32v16_fmaFMA(a, b, c, result)
	} else {
		f32v16_fmaGeneric(a, b, c, result)
	}
}

func f32v16_divAVX512F(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_divAVX(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func f32v16_divSSE2(a *Float32Vec16, b *Float32Vec16, result *Float32Vec16)
func (a *Float32Vec16) Div(b *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_divAVX512F(a, b, result)
	case cpu.X86.HasAVX: f32v16_divAVX(a, b, result)
	default: f32v16_divSSE2(a, b, result)
	}
}

func f32v16_minAVX512F(v *Float32Vec16, mn *Float32Vec16, result *Float32Vec16)
func f32v16_minAVX(v *Float32Vec16, mn *Float32Vec16, result *Float32Vec16)
func f32v16_minSSE2(v *Float32Vec16, mn *Float32Vec16, result *Float32Vec16)
func (v *Float32Vec16) Min(mn *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_minAVX512F(v, mn, result)
	case cpu.X86.HasAVX: f32v16_minAVX(v, mn, result)
	default: f32v16_minSSE2(v, mn, result)
	}
}

func f32v16_maxAVX512F(v *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func f32v16_maxAVX(v *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func f32v16_maxSSE2(v *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func (v *Float32Vec16) Max(mx *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_maxAVX512F(v, mx, result)
	case cpu.X86.HasAVX: f32v16_maxAVX(v, mx, result)
	default: f32v16_maxSSE2(v, mx, result)
	}
}

func f32v16_clampAVX512F(v *Float32Vec16, mn *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func f32v16_clampAVX(v *Float32Vec16, mn *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func f32v16_clampSSE2(v *Float32Vec16, mn *Float32Vec16, mx *Float32Vec16, result *Float32Vec16)
func (v *Float32Vec16) Clamp(mn *Float32Vec16, mx *Float32Vec16, result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_maxAVX512F(v, mn, mx, result)
	case cpu.X86.HasAVX: f32v16_maxAVX(v, mn, mx, result)
	default: f32v16_maxSSE2(v, mn, mx, result)
	}
}

func f32v16_ceilAVX512F(v *Float32Vec16, result *Float32Vec16)
func f32v16_ceilAVX(v *Float32Vec16, result *Float32Vec16)
func f32v16_ceilSSE41(v *Float32Vec16, result *Float32Vec16)
func f32v16_ceilGeneric(v *Float32Vec16, result *Float32Vec16) {
	v1 := *v

	v1[0] = float32(std_math.Ceil(float64(v1[0])))
	v1[1] = float32(std_math.Ceil(float64(v1[1])))
	v1[2] = float32(std_math.Ceil(float64(v1[2])))
	v1[3] = float32(std_math.Ceil(float64(v1[3])))
	v1[4] = float32(std_math.Ceil(float64(v1[4])))
	v1[5] = float32(std_math.Ceil(float64(v1[5])))
	v1[6] = float32(std_math.Ceil(float64(v1[6])))
	v1[7] = float32(std_math.Ceil(float64(v1[7])))
	v1[8] = float32(std_math.Ceil(float64(v1[8])))
	v1[9] = float32(std_math.Ceil(float64(v1[9])))
	v1[10] = float32(std_math.Ceil(float64(v1[10])))
	v1[11] = float32(std_math.Ceil(float64(v1[11])))
	v1[12] = float32(std_math.Ceil(float64(v1[12])))
	v1[13] = float32(std_math.Ceil(float64(v1[13])))
	v1[14] = float32(std_math.Ceil(float64(v1[14])))
	v1[15] = float32(std_math.Ceil(float64(v1[15])))
	
	*result = v1
}
func (v *Float32Vec16) Ceil(result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_ceilAVX512F(v, result)
	case cpu.X86.HasAVX: f32v16_ceilAVX(v, result)
	case cpu.X86.HasSSE41: f32v16_ceilSSE41(v, result)
	default: f32v16_ceilGeneric(v, result)
	}
}

func f32v16_roundAVX512F(v *Float32Vec16, result *Float32Vec16)
func f32v16_roundAVX(v *Float32Vec16, result *Float32Vec16)
func f32v16_roundSSE41(v *Float32Vec16, result *Float32Vec16)
func f32v16_roundGeneric(v *Float32Vec16, result *Float32Vec16) {
	v1 := *v

	v1[0] = float32(std_math.Round(float64(v1[0])))
	v1[1] = float32(std_math.Round(float64(v1[1])))
	v1[2] = float32(std_math.Round(float64(v1[2])))
	v1[3] = float32(std_math.Round(float64(v1[3])))
	v1[4] = float32(std_math.Round(float64(v1[4])))
	v1[5] = float32(std_math.Round(float64(v1[5])))
	v1[6] = float32(std_math.Round(float64(v1[6])))
	v1[7] = float32(std_math.Round(float64(v1[7])))
	v1[8] = float32(std_math.Round(float64(v1[8])))
	v1[9] = float32(std_math.Round(float64(v1[9])))
	v1[10] = float32(std_math.Round(float64(v1[10])))
	v1[11] = float32(std_math.Round(float64(v1[11])))
	v1[12] = float32(std_math.Round(float64(v1[12])))
	v1[13] = float32(std_math.Round(float64(v1[13])))
	v1[14] = float32(std_math.Round(float64(v1[14])))
	v1[15] = float32(std_math.Round(float64(v1[15])))
	
	*result = v1
}
func (v *Float32Vec16) Round(result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_roundAVX512F(v, result)
	case cpu.X86.HasAVX: f32v16_roundAVX(v, result)
	case cpu.X86.HasSSE41: f32v16_roundSSE41(v, result)
	default: f32v16_roundGeneric(v, result)
	}
}

func f32v16_floorAVX512F(v *Float32Vec16, result *Float32Vec16)
func f32v16_floorAVX(v *Float32Vec16, result *Float32Vec16)
func f32v16_floorSSE41(v *Float32Vec16, result *Float32Vec16)
func f32v16_floorGeneric(v *Float32Vec16, result *Float32Vec16) {
	v1 := *v

	v1[0] = float32(std_math.Floor(float64(v1[0])))
	v1[1] = float32(std_math.Floor(float64(v1[1])))
	v1[2] = float32(std_math.Floor(float64(v1[2])))
	v1[3] = float32(std_math.Floor(float64(v1[3])))
	v1[4] = float32(std_math.Floor(float64(v1[4])))
	v1[5] = float32(std_math.Floor(float64(v1[5])))
	v1[6] = float32(std_math.Floor(float64(v1[6])))
	v1[7] = float32(std_math.Floor(float64(v1[7])))
	v1[8] = float32(std_math.Floor(float64(v1[8])))
	v1[9] = float32(std_math.Floor(float64(v1[9])))
	v1[10] = float32(std_math.Floor(float64(v1[10])))
	v1[11] = float32(std_math.Floor(float64(v1[11])))
	v1[12] = float32(std_math.Floor(float64(v1[12])))
	v1[13] = float32(std_math.Floor(float64(v1[13])))
	v1[14] = float32(std_math.Floor(float64(v1[14])))
	v1[15] = float32(std_math.Floor(float64(v1[15])))
	
	*result = v1
}
func (v *Float32Vec16) Floor(result *Float32Vec16) {
	switch {
	case cpu.X86.HasAVX512F: f32v16_floorAVX512F(v, result)
	case cpu.X86.HasAVX: f32v16_floorAVX(v, result)
	case cpu.X86.HasSSE41: f32v16_floorSSE41(v, result)
	default: f32v16_floorGeneric(v, result)
	}
}
