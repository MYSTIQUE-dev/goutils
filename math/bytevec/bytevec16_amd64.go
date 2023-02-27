//go:build amd64

package bytevec

type ByteVec16 [16]byte

func bv16_Add(a *ByteVec16, b *ByteVec16, result *ByteVec16)
func (a *ByteVec16) Add(b *ByteVec16, result *ByteVec16) {
	bv16_Add(a, b, result)
}

func bv16_Sub(a *ByteVec16, b *ByteVec16, result *ByteVec16)
func (a *ByteVec16) Sub(b *ByteVec16, result *ByteVec16) {
	bv16_Sub(a, b, result)
}

func (a *ByteVec16) Mul(b *ByteVec16, result *ByteVec16) {
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
	v1[8] *= v2[8]
	v1[9] *= v2[9]
	v1[10] *= v2[10]
	v1[11] *= v2[11]
	v1[12] *= v2[12]
	v1[13] *= v2[13]
	v1[14] *= v2[14]
	v1[15] *= v2[15]
	
	*result = v1
}

func (a *ByteVec16) Div(b *ByteVec16, result *ByteVec16) {
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
	v1[8] /= v2[8]
	v1[9] /= v2[9]
	v1[10] /= v2[10]
	v1[11] /= v2[11]
	v1[12] /= v2[12]
	v1[13] /= v2[13]
	v1[14] /= v2[14]
	v1[15] /= v2[15]
	
	*result = v1
}

func bv16_Min(v *ByteVec16, mn *ByteVec16, result *ByteVec16)
func (v *ByteVec16) Min(mn *ByteVec16, result *ByteVec16) {
	bv16_Min(v, mn, result)
}

func bv16_Max(v *ByteVec16, mx *ByteVec16, result *ByteVec16)
func (v *ByteVec16) Max(mx *ByteVec16, result *ByteVec16) {
	bv16_Max(v, mx, result)
}

func bv16_Clamp(v *ByteVec16, mn *ByteVec16, mx *ByteVec16, result *ByteVec16)
func (v *ByteVec16) Clamp(mn *ByteVec16, mx *ByteVec16, result *ByteVec16) {
	bv16_Clamp(v, mn, mx, result)
}
