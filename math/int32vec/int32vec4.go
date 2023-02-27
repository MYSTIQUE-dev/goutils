//go:build !amd64

package int32vec

type Int32Vec4 [4]int32

func (a *Int32Vec4) Add(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	*result = Int32Vec4([4]int32{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
		v1[3] + v2[3],
	})
}

func (a *Int32Vec4) Sub(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	*result = Int32Vec4([4]int32{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
		v1[3] - v2[3],
	})
}

func (a *Int32Vec4) Mul(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	*result = Int32Vec4([4]int32{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
		v1[3] * v2[3],
	})
}

func (a *Int32Vec4) MulAdd(b *Int32Vec4, c *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b
	v3 := *c

	*result = Int32Vec4([4]int32{
		v1[0] * v2[0] + v3[0],
		v1[1] * v2[1] + v3[1],
		v1[2] * v2[2] + v3[2],
		v1[3] * v2[3] + v3[3],
	})
}

func (a *Int32Vec4) Div(b *Int32Vec4, result *Int32Vec4) {
	v1 := *a
	v2 := *b

	*result = Int32Vec4([4]int32{
		v1[0] / v2[0],
		v1[1] / v2[1],
		v1[2] / v2[2],
		v1[3] / v2[3],
	})
}

func (v *Int32Vec4) Min(mn *Int32Vec4, result *Int32Vec4) {
	v1 := *v
	v2 := *mn

	*result = Int32Vec4([4]int32{
		Min(v1[0], v2[0]),
		Min(v1[1], v2[1]),
		Min(v1[2], v2[2]),
		Min(v1[3], v2[3]),
	})
}

func (v *Int32Vec4) Max(mx *Int32Vec4, result *Int32Vec4) {
	v1 := *v
	v2 := *mx

	*result = Int32Vec4([4]int32{
		Max(v1[0], v2[0]),
		Max(v1[1], v2[1]),
		Max(v1[2], v2[2]),
		Max(v1[3], v2[3]),
	})
}

func (v *Int32Vec4) Clamp(mn *Int32Vec4, mx *Int32Vec4, result *Int32Vec4) {
	v1 := *mn
	v2 := *mx
	v3 := *v

	v1[0] = math.Max(v1[0], math.Min(v2[0], v3[0]))
	v1[1] = math.Max(v1[1], math.Min(v2[1], v3[1]))
	v1[2] = math.Max(v1[2], math.Min(v2[2], v3[2]))
	v1[3] = math.Max(v1[3], math.Min(v2[3], v3[3]))

	*result = v1
}
