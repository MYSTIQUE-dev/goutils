//go:build !amd64

package int64vec

import "github.com/notsatvrn/goutils/math"

type Int64Vec4 [4]int64

func (a *Int64Vec4) Add(b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	v1[3] += v2[3]
	
	*result = v1
}

func (a *Int64Vec4) Sub(b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	v1[3] -= v2[3]
	
	*result = v1
}

func (a *Int64Vec4) Mul(b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	v1[3] *= v2[3]
	
	*result = v1
}

func (a *Int64Vec4) Div(b *Int64Vec4, result *Int64Vec4) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	v1[3] /= v2[3]
	
	*result = v1
}

func (v *Int64Vec4) Min(mn *Int64Vec4, result *Int64Vec4) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	v1[3] = math.Min(v1[3], v2[3])
	
	*result = v1
}

func (v *Int64Vec4) Max(mx *Int64Vec4, result *Int64Vec4) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	v1[3] = math.Max(v1[3], v2[3])
	
	*result = v1
}

func (v *Int64Vec4) Clamp(mn *Int64Vec4, mx *Int64Vec4, result *Int64Vec4) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[0] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[0] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[0] = math.Max(v1[3], math.Min(v2[3], v3[3]))

	*result = v1
}
