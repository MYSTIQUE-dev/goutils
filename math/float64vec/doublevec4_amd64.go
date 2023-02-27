//go:build amd64

package math

import (
	"golang.org/x/sys/cpu"
	std_math "math"
)

type DoubleVec4 [4]float64

func dv4_AddAVX(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func dv4_AddSSE2(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func (a *DoubleVec4) Add(b *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_AddAVX(a, b, result)
	} else {
		dv4_AddSSE2(a, b, result)
	}
}

func dv4_SubAVX(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func dv4_SubSSE2(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func (a *DoubleVec4) Sub(b *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_SubAVX(a, b, result)
	} else {
		dv4_SubSSE2(a, b, result)
	}
}

func dv4_MulAVX(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func dv4_MulSSE2(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func (a *DoubleVec4) Mul(b *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_MulAVX(a, b, result)
	} else {
		dv4_MulSSE2(a, b, result)
	}
}

func dv4_MulAddFMA(a *DoubleVec4, b *DoubleVec4, c *DoubleVec4, result *DoubleVec4)
func dv4_MulAddAVX(a *DoubleVec4, b *DoubleVec4, c *DoubleVec4, result *DoubleVec4)
func dv4_MulAddSSE2(a *DoubleVec4, b *DoubleVec4, c *DoubleVec4, result *DoubleVec4)
func (a *DoubleVec4) MulAdd(b *DoubleVec4, c *DoubleVec4, result *DoubleVec4) {
	switch {
	case cpu.X86.HasFMA: dv4_MulAddFMA(a, b, c, result)
	case cpu.X86.HasAVX: dv4_MulAddAVX(a, b, c, result)
	default: dv4_MulAddSSE2(a, b, c, result)
	}
}

func dv4_DivAVX(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func dv4_DivSSE2(a *DoubleVec4, b *DoubleVec4, result *DoubleVec4)
func (a *DoubleVec4) Div(b *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_DivAVX(a, b, result)
	} else {
		dv4_DivSSE2(a, b, result)
	}
}

func dv4_MinAVX(v *DoubleVec4, mn *DoubleVec4, result *DoubleVec4)
func dv4_MinSSE2(v *DoubleVec4, mn *DoubleVec4, result *DoubleVec4)
func (v *DoubleVec4) Min(mn *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_MinAVX(v, mn, result)
	} else {
		dv4_MinSSE2(v, mn, result)
	}
}

func dv4_MaxAVX(v *DoubleVec4, mx *DoubleVec4, result *DoubleVec4)
func dv4_MaxSSE2(v *DoubleVec4, mx *DoubleVec4, result *DoubleVec4)
func (v *DoubleVec4) Max(mx *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_MaxAVX(v, mx, result)
	} else {
		dv4_MaxSSE2(v, mx, result)
	}
}

func dv4_ClampAVX(v *DoubleVec4, mn *DoubleVec4, mx *DoubleVec4, result *DoubleVec4)
func dv4_ClampSSE2(v *DoubleVec4, mn *DoubleVec4, mx *DoubleVec4, result *DoubleVec4)
func (v *DoubleVec4) Clamp(mn *DoubleVec4, mx *DoubleVec4, result *DoubleVec4) {
	if cpu.X86.HasAVX {
		dv4_ClampAVX(v, mn, mx, result)
	} else {
		dv4_ClampSSE2(v, mn, mx, result)
	}
}

func dv4_CeilAVX(v *DoubleVec4, result *DoubleVec4)
func dv4_CeilSSE41(v *DoubleVec4, result *DoubleVec4)
func dv4_CeilGeneric(v *DoubleVec4, result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Ceil(v1[0]),
		std_math.Ceil(v1[1]),
		std_math.Ceil(v1[2]),
		std_math.Ceil(v1[3]),
	})
}
func (v *DoubleVec4) Ceil(result *DoubleVec4) {
	switch {
	case cpu.X86.HasAVX: dv4_CeilAVX(v, result)
	case cpu.X86.HasSSE41: dv4_CeilSSE41(v, result)
	default: dv4_CeilGeneric(v, result)
	}
}

func dv4_RoundAVX(v *DoubleVec4, result *DoubleVec4)
func dv4_RoundSSE41(v *DoubleVec4, result *DoubleVec4)
func dv4_RoundGeneric(v *DoubleVec4, result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Round(v1[0]),
		std_math.Round(v1[1]),
		std_math.Round(v1[2]),
		std_math.Round(v1[3]),
	})
}
func (v *DoubleVec4) Round(result *DoubleVec4) {
	switch {
	case cpu.X86.HasAVX: dv4_RoundAVX(v, result)
	case cpu.X86.HasSSE41: dv4_RoundSSE41(v, result)
	default: dv4_RoundGeneric(v, result)
	}
}

func dv4_FloorAVX(v *DoubleVec4, result *DoubleVec4)
func dv4_FloorSSE41(v *DoubleVec4, result *DoubleVec4)
func dv4_FloorGeneric(v *DoubleVec4, result *DoubleVec4) {
	v1 := *v

	*result = DoubleVec4([4]float64{
		std_math.Floor(v1[0]),
		std_math.Floor(v1[1]),
		std_math.Floor(v1[2]),
		std_math.Floor(v1[3]),
	})
}
func (v *DoubleVec4) Floor(result *DoubleVec4) {
	switch {
	case cpu.X86.HasAVX: dv4_FloorAVX(v, result)
	case cpu.X86.HasSSE41: dv4_FloorSSE41(v, result)
	default: dv4_FloorGeneric(v, result)
	}
}
