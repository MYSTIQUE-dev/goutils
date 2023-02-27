//go:build amd64

package float32vec

import (
	"golang.org/x/sys/cpu"
	std_math "math"
)

type Float32Vec8 [8]float32

func f32v8_addAVX(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func f32v8_addSSE2(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func (a *Float32Vec8) Add(b *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_addAVX(a, b, result)
	} else {
		f32v8_addSSE2(a, b, result)
	}
}

func f32v8_subAVX(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func f32v8_subSSE2(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func (a *Float32Vec8) Sub(b *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_subAVX(a, b, result)
	} else {
		f32v8_subSSE2(a, b, result)
	}
}

func f32v8_mulAVX(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func f32v8_mulSSE2(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func (a *Float32Vec8) Mul(b *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_mulAVX(a, b, result)
	} else {
		f32v8_mulSSE2(a, b, result)
	}
}

func f32v8_fmaFMA(a *Float32Vec8, b *Float32Vec8, c *Float32Vec8, result *Float32Vec8)
func f32v8_fmaGeneric(a *Float32Vec8, b *Float32Vec8, c *Float32Vec8, result *Float32Vec8) {
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
func (a *Float32Vec8) FMA(b *Float32Vec8, c *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasFMA {
		f32v8_fmaFMA(a, b, c, result)
	} else {
		f32v8_fmaGeneric(a, b, c, result)
	}
}

func f32v8_divAVX(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func f32v8_divSSE2(a *Float32Vec8, b *Float32Vec8, result *Float32Vec8)
func (a *Float32Vec8) Div(b *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_divAVX(a, b, result)
	} else {
		f32v8_divSSE2(a, b, result)
	}
}

func f32v8_minAVX(v *Float32Vec8, mn *Float32Vec8, result *Float32Vec8)
func f32v8_minSSE2(v *Float32Vec8, mn *Float32Vec8, result *Float32Vec8)
func (v *Float32Vec8) Min(mn *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_minAVX(v, mn, result)
	} else {
		f32v8_minSSE2(v, mn, result)
	}
}

func f32v8_maxAVX(v *Float32Vec8, mx *Float32Vec8, result *Float32Vec8)
func f32v8_maxSSE2(v *Float32Vec8, mx *Float32Vec8, result *Float32Vec8)
func (v *Float32Vec8) Max(mx *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_maxAVX(v, mx, result)
	} else {
		f32v8_maxSSE2(v, mx, result)
	}
}

func f32v8_clampAVX(v *Float32Vec8, mn *Float32Vec8, mx *Float32Vec8, result *Float32Vec8)
func f32v8_clampSSE2(v *Float32Vec8, mn *Float32Vec8, mx *Float32Vec8, result *Float32Vec8)
func (v *Float32Vec8) Clamp(mn *Float32Vec8, mx *Float32Vec8, result *Float32Vec8) {
	if cpu.X86.HasAVX {
		f32v8_clampAVX(v, mn, mx, result)
	} else {
		f32v8_clampSSE2(v, mn, mx, result)
	}
}

func f32v8_ceilAVX(v *Float32Vec8, result *Float32Vec8)
func f32v8_ceilSSE41(v *Float32Vec8, result *Float32Vec8)
func f32v8_ceilGeneric(v *Float32Vec8, result *Float32Vec8) {
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
func (v *Float32Vec8) Ceil(result *Float32Vec8) {
	switch {
	case cpu.X86.HasAVX: f32v8_ceilAVX(v, result)
	case cpu.X86.HasSSE41: f32v8_ceilSSE41(v, result)
	default: f32v8_ceilGeneric(v, result)
	}
}

func f32v8_roundAVX(v *Float32Vec8, result *Float32Vec8)
func f32v8_roundSSE41(v *Float32Vec8, result *Float32Vec8)
func f32v8_roundGeneric(v *Float32Vec8, result *Float32Vec8) {
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
func (v *Float32Vec8) Round(result *Float32Vec8) {
	switch {
	case cpu.X86.HasAVX: f32v8_roundAVX(v, result)
	case cpu.X86.HasSSE41: f32v8_roundSSE41(v, result)
	default: f32v8_roundGeneric(v, result)
	}
}

func f32v8_floorAVX(v *Float32Vec8, result *Float32Vec8)
func f32v8_floorSSE41(v *Float32Vec8, result *Float32Vec8)
func f32v8_floorGeneric(v *Float32Vec8, result *Float32Vec8) {
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
func (v *Float32Vec8) Floor(result *Float32Vec8) {
	switch {
	case cpu.X86.HasAVX: f32v8_floorAVX(v, result)
	case cpu.X86.HasSSE41: f32v8_floorSSE41(v, result)
	default: f32v8_floorGeneric(v, result)
	}
}
