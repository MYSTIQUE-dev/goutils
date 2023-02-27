package int8vec

import "github.com/notsatvrn/goutils/math"

type Int8Vec3 [3]int8

func (a *Int8Vec3) Add(b *Int8Vec3, result *Int8Vec3) {
	v1 := *a
	v2 := *b

	v1[0] += v2[0]
	v1[1] += v2[1]
	v1[2] += v2[2]
	
	*result = v1
}

func (a *Int8Vec3) Sub(b *Int8Vec3, result *Int8Vec3) {
	v1 := *a
	v2 := *b

	v1[0] -= v2[0]
	v1[1] -= v2[1]
	v1[2] -= v2[2]
	
	*result = v1
}

func (a *Int8Vec3) Mul(b *Int8Vec3, result *Int8Vec3) {
	v1 := *a
	v2 := *b

	v1[0] *= v2[0]
	v1[1] *= v2[1]
	v1[2] *= v2[2]
	
	*result = v1
}

func (a *Int8Vec3) Div(b *Int8Vec3, result *Int8Vec3) {
	v1 := *a
	v2 := *b

	v1[0] /= v2[0]
	v1[1] /= v2[1]
	v1[2] /= v2[2]
	
	*result = v1
}

func (v *Int8Vec3) Min(mn *Int8Vec3, result *Int8Vec3) {
	v1 := *v
	v2 := *mn

	v1[0] = math.Min(v1[0], v2[0])
	v1[1] = math.Min(v1[1], v2[1])
	v1[2] = math.Min(v1[2], v2[2])
	
	*result = v1
}

func (v *Int8Vec3) Max(mx *Int8Vec3, result *Int8Vec3) {
	v1 := *v
	v2 := *mx

	v1[0] = math.Max(v1[0], v2[0])
	v1[1] = math.Max(v1[1], v2[1])
	v1[2] = math.Max(v1[2], v2[2])
	
	*result = v1
}

func (v *Int8Vec3) Clamp(mn *Int8Vec3, mx *Int8Vec3, result *Int8Vec3) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[0] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[0] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	
	*result = v1
}
