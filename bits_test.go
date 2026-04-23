package bit

import (
	"fmt"
	"slices"
	"testing"
)

func wrapU(v, n int64) int64 {
	if n == 64 {
		return v
	}
	mask := (int64(1) << n) - 1
	return v & mask
}

func wrapI(v, n int64) int64 {
	if n == 64 {
		return v
	}
	mask := (int64(1) << n) - 1
	sign := int64(1) << (n - 1)
	return ((v & mask) ^ sign) - sign
}

func TestSet(t *testing.T) {
	{
		want := []U3{0, 1, 2, 3, 4, 5, 6, 7, 0}
		got := []U3{}
		for i := range 9 {
			var u3 U3
			u3.Set(U3(i))
			got = append(got, u3)
		}
		if !slices.Equal(want, got) {
			t.Errorf("\n got:  %#v\n want: %#v\n", got, want)
		}
	}
	{
		want := []I3{-1, 0, 1, 2, 3, -4, -3, -2, -1}
		got := []I3{}
		for i := range 9 {
			var i3 I3
			i3.Set(-9 + I3(i))
			got = append(got, i3)
		}
		if !slices.Equal(want, got) {
			t.Errorf("\n got:  %#v\n want: %#v\n", got, want)
		}
	}
}

func TestU1_Arithmetic(t *testing.T) {
	type row struct {
		name    string
		a, b, w int64
		fn      func(U1, U1) U1
	}
	tests := []row{
		{"Add 0+0", 0, 0, 0, func(a, b U1) U1 { return a.Add(b) }},
		{"Add 1+1 wrap", 1, 1, 0, func(a, b U1) U1 { return a.Add(b) }},
		{"Sub 0-1 wrap", 0, 1, 1, func(a, b U1) U1 { return a.Sub(b) }},
		{"Mul 1*1", 1, 1, 1, func(a, b U1) U1 { return a.Mul(b) }},
		{"And 1&0", 1, 0, 0, func(a, b U1) U1 { return a.And(b) }},
		{"Or  0|1", 0, 1, 1, func(a, b U1) U1 { return a.Or(b) }},
		{"Xor 1^1", 1, 1, 0, func(a, b U1) U1 { return a.Xor(b) }},
		{"Shr 1>>0", 1, 0, 1, func(a, b U1) U1 { return a.Shr(b) }},
		{"Shl 1<<0", 1, 0, 1, func(a, b U1) U1 { return a.Shl(b) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := int64(tt.fn(U1(tt.a), U1(tt.b)))
			if got != tt.w {
				t.Errorf("got %d, want %d", got, tt.w)
			}
		})
	}
}

func TestU1_IncDec(t *testing.T) {
	if U1(1).Inc() != 0 {
		t.Error("Inc wrap failed")
	}
	if U1(0).Dec() != 1 {
		t.Error("Dec wrap failed")
	}
}

func TestU1_Not(t *testing.T) {
	if U1(0).Not() != 1 {
		t.Errorf("Not(0) = %d, want 1", U1(0).Not())
	}
	if U1(1).Not() != 0 {
		t.Errorf("Not(1) = %d, want 0", U1(1).Not())
	}
}

func TestU1_BitBits(t *testing.T) {
	v := U1(1)
	if v.Bit(0) != 1 {
		t.Errorf("Bit(0) = %d, want 1", v.Bit(0))
	}
	if v.Bits(0, 0) != 1 {
		t.Errorf("Bits(0,0) = %d, want 1", v.Bits(0, 0))
	}
	v2 := U1(1)
	v2.SetBit(0, 0)
	if v2 != 0 {
		t.Errorf("setBit(0,0) = %d, want 0", v2)
	}
	v3 := U1(0)
	v3.SetBits(0, 0, 1)
	if v3 != 1 {
		t.Errorf("SetBits(0,0,1) = %d, want 1", v3)
	}
}

func TestU1_RbitRbits(t *testing.T) {
	v := U1(0)
	r := v.Bitref(0)
	r.Set(1)
	if v != 1 {
		t.Errorf("rbit.Set(1) => v=%d want 1", v)
	}
	v2 := U1(0)
	r2 := v2.Bitsref(0, 0)
	r2.Set(1)
	if v2 != 1 {
		t.Errorf("rbits.Set(1) => v=%d want 1", v2)
	}
}

func TestU1_Signed(t *testing.T) {
	s := U1(1).Signed()
	if s != -1 {
		t.Errorf("u1(1).Signed() = %d, want -1", s)
	}
	s0 := U1(0).Signed()
	if s0 != 0 {
		t.Errorf("u1(0).Signed() = %d, want 0", s0)
	}
}

func TestU1_Clamp(t *testing.T) {
	if U1(0).Clamp(0) != 0 {
		t.Error("Clamp(0)")
	}
	if U1(0).Clamp(1) != 1 {
		t.Error("Clamp(1)")
	}
	if U1(0).Clamp(99) != 1 {
		t.Errorf("Clamp(99) = %d, want 1", U1(0).Clamp(99))
	}
}

func TestI3_Clamp(t *testing.T) {
	tests := []struct {
		in   int64
		want int64
	}{
		{0, 0},
		{3, 3},
		{4, 3},
		{100, 3},
		{-4, -4},
		{-5, -4},
		{-100, -4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := int64(I3(0).Clamp(tt.in))
			if got != tt.want {
				t.Errorf("I3.Clamp(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestI3_Clip(t *testing.T) {
	tests := []struct {
		val  I3
		bits int
		want I3
	}{
		{0b011, 2, -1},
		{0b001, 2, 1},
		{0b010, 2, -2},
		{0b111, 3, -1},
		{0b011, 3, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Clip(%d,%d)", tt.val, tt.bits), func(t *testing.T) {
			got := tt.val.Clip(tt.bits)
			if got != tt.want {
				t.Errorf("I3(%d).Clip(%d) = %d, want %d", tt.val, tt.bits, got, tt.want)
			}
		})
	}
}

func TestI1_Clamp(t *testing.T) {
	if I1(0).Clamp(0) != 0 {
		t.Error("I1.Clamp(0)")
	}
	if I1(0).Clamp(-1) != -1 {
		t.Error("I1.Clamp(-1)")
	}
	if I1(0).Clamp(1) != 0 {
		t.Errorf("I1.Clamp(1) = %d, want 0", I1(0).Clamp(1))
	}
	if I1(0).Clamp(-99) != -1 {
		t.Errorf("I1.Clamp(-99) = %d, want -1", I1(0).Clamp(-99))
	}
}

func TestI1_Arithmetic(t *testing.T) {
	type row struct {
		name    string
		a, b, w int64
		fn      func(I1, I1) I1
	}
	tests := []row{
		{"Add 0+0", 0, 0, 0, func(a, b I1) I1 { return a.Add(b) }},
		{"Add -1+0", -1, 0, -1, func(a, b I1) I1 { return a.Add(b) }},
		{"Sub 0-(-1)", 0, -1, -1, func(a, b I1) I1 { return a.Sub(b) }},
		{"Mul -1*-1 wrap", -1, -1, -1, func(a, b I1) I1 { return a.Mul(b) }},
		{"And -1&0", -1, 0, 0, func(a, b I1) I1 { return a.And(b) }},
		{"Xor -1^-1", -1, -1, 0, func(a, b I1) I1 { return a.Xor(b) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := int64(tt.fn(I1(tt.a), I1(tt.b)))
			if got != tt.w {
				t.Errorf("got %d, want %d", got, tt.w)
			}
		})
	}
}

func TestI1_Not(t *testing.T) {
	if I1(0).Not() != -1 {
		t.Errorf("I1(0).Not() = %d, want -1", I1(0).Not())
	}
	if I1(-1).Not() != 0 {
		t.Errorf("I1(-1).Not() = %d, want 0", I1(-1).Not())
	}
}

func TestI1_BitBits(t *testing.T) {
	v := I1(-1)
	if v.Bit(0) != 1 {
		t.Errorf("Bit(0) of -1 = %d, want 1 (raw bit)", v.Bit(0))
	}
	v2 := I1(-1)
	v2.SetBit(0, 0)
	if v2 != 0 {
		t.Errorf("SetBit(0,0) = %d, want 0", v2)
	}
}

func TestI1_RbitRbits(t *testing.T) {
	v := I1(0)
	v.Bitref(0).Set(-1)
	if v != 1 {
		t.Errorf("Bitref.Set(-1) => %d, want 1 (raw bit)", v)
	}
}

func TestI1_Unsigned(t *testing.T) {
	u := I1(-1).Unsigned()
	if u != 1 {
		t.Errorf("I1(-1).Unsigned() = %d, want 1", u)
	}
}

func TestU3_WrapAround(t *testing.T) {
	tests := []struct {
		name string
		in   int64
		want int64
	}{
		{"7", 7, 7},
		{"8 wraps to 0", 8, 0},
		{"255 wraps", 255, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := U3(tt.in)
			got := int64(U3.cast(v))
			if got != tt.want {
				t.Errorf("cast(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestU3_AllOps(t *testing.T) {
	type row struct {
		name    string
		a, b, w int64
	}
	ops := []struct {
		op string
		fn func(a, b U3) U3
	}{
		{"Add", func(a, b U3) U3 { return a.Add(b) }},
		{"Sub", func(a, b U3) U3 { return a.Sub(b) }},
		{"Mul", func(a, b U3) U3 { return a.Mul(b) }},
		{"And", func(a, b U3) U3 { return a.And(b) }},
		{"Or", func(a, b U3) U3 { return a.Or(b) }},
		{"Xor", func(a, b U3) U3 { return a.Xor(b) }},
	}
	inputs := [][2]int64{{0, 0}, {3, 5}, {7, 1}, {5, 3}, {6, 7}}
	for _, op := range ops {
		for _, ab := range inputs {
			a, b := ab[0], ab[1]
			var want int64
			switch op.op {
			case "Add":
				want = wrapU(a+b, 3)
			case "Sub":
				want = wrapU(a-b, 3)
			case "Mul":
				want = wrapU(a*b, 3)
			case "And":
				want = a & b
			case "Or":
				want = a | b
			case "Xor":
				want = a ^ b
			}
			t.Run(op.op, func(t *testing.T) {
				got := int64(op.fn(U3(a), U3(b)))
				if got != want {
					t.Errorf("%s(%d,%d) = %d, want %d", op.op, a, b, got, want)
				}
			})
		}
	}
}

func TestU3_BitFields(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		lo    int
		hi    int
		want  int64
	}{
		{"Bits(0,1) of 0b110", 0b110, 0, 1, 0b10},
		{"Bits(1,2) of 0b110", 0b110, 1, 2, 0b11},
		{"Bits(0,0) of 0b101", 0b101, 0, 0, 1},
		{"Bits(2,2) of 0b100", 0b100, 2, 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := U3(tt.value)
			got := int64(v.Bits(tt.lo, tt.hi))
			if got != tt.want {
				t.Errorf("Bits(%d,%d) of %03b = %d, want %d", tt.lo, tt.hi, v, got, tt.want)
			}
		})
	}
}

func TestU3_SetBitFields(t *testing.T) {
	v := U3(0b101)
	v.SetBit(1, 1)
	if v != 0b111 {
		t.Errorf("SetBit(1,1) of 101 = %03b, want 111", v)
	}
	v.SetBits(0, 1, 0b00)
	if v != 0b100 {
		t.Errorf("SetBits(0,1,0) of 111 = %03b, want 100", v)
	}
}

func TestU3_RbitRbits(t *testing.T) {
	v := U3(0b000)
	v.Bitsref(0, 1).Set(0b11)
	if v != 0b011 {
		t.Errorf("Bitsref(0,1).Set(0b11) = %03b, want 011", v)
	}

	v2 := U3(0b000)
	v2.Bitref(2).Set(1)
	if v2 != 0b100 {
		t.Errorf("rbit(2).Set(1) = %03b, want 100", v2)
	}
}

func TestU3_NegativeIndex(t *testing.T) {
	v := U3(0b100)
	if v.Bit(-1) != 1 {
		t.Errorf("Bit(-1) of 100 = %d, want 1", v.Bit(-1))
	}
	got := v.Bits(-2, -1)
	if got != 0b10 {
		t.Errorf("bits(-2,-1) of 100 = %b, want 10", got)
	}
}

func TestI3_Range(t *testing.T) {
	tests := []struct {
		in, want int64
	}{
		{0, 0},
		{3, 3},
		{4, -4},

		{7, -1},

		{-1, -1},

		{-4, -4},

		{-5, 3},
	}
	for _, tt := range tests {
		v := I3(tt.in)
		got := int64(I3.cast(v))
		if got != tt.want {
			t.Errorf("I3(%d).cast() = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestI3_AllOps(t *testing.T) {
	pairs := [][2]int64{{0, 0}, {3, 1}, {-4, 1}, {2, -3}, {-1, -1}}
	ops := []struct {
		name string
		fn   func(I3, I3) I3
		ref  func(a, b int64) int64
	}{
		{"Add", func(a, b I3) I3 { return a.Add(b) }, func(a, b int64) int64 { return wrapI(a+b, 3) }},
		{"Sub", func(a, b I3) I3 { return a.Sub(b) }, func(a, b int64) int64 { return wrapI(a-b, 3) }},
		{"Mul", func(a, b I3) I3 { return a.Mul(b) }, func(a, b int64) int64 { return wrapI(a*b, 3) }},
		{"And", func(a, b I3) I3 { return a.And(b) }, func(a, b int64) int64 { return wrapI(a&b, 3) }},
		{"Or", func(a, b I3) I3 { return a.Or(b) }, func(a, b int64) int64 { return wrapI(a|b, 3) }},
		{"Xor", func(a, b I3) I3 { return a.Xor(b) }, func(a, b int64) int64 { return wrapI(a^b, 3) }},
	}
	for _, op := range ops {
		for _, ab := range pairs {
			a, b := ab[0], ab[1]
			want := op.ref(a, b)
			t.Run(op.name, func(t *testing.T) {
				got := int64(op.fn(I3(a), I3(b)))
				if got != want {
					t.Errorf("%s(%d,%d) = %d, want %d", op.name, a, b, got, want)
				}
			})
		}
	}
}

func TestI3_IncDec(t *testing.T) {
	if I3(3).Inc() != -4 {
		t.Errorf("I3(3).Inc() = %d, want -4", I3(3).Inc())
	}
	if I3(-4).Dec() != 3 {
		t.Errorf("I3(-4).Dec() = %d, want 3", I3(-4).Dec())
	}
}

func TestI3_Not(t *testing.T) {
	if I3(0).Not() != -1 {
		t.Errorf("I3(0).Not() = %d, want -1", I3(0).Not())
	}
}

func TestI3_UnsignedRoundtrip(t *testing.T) {
	for raw := int64(0); raw < 8; raw++ {
		u := U3(raw)
		s := u.Signed()
		back := s.Unsigned()
		if int64(back) != raw {
			t.Errorf("u3(%d).Signed().Unsigned() = %d, want %d", raw, back, raw)
		}
	}
}

func TestI3_BitBits(t *testing.T) {
	v := I3(-1)
	if v.Bit(0) != 1 {
		t.Errorf("I3(-1).Bit(0) = %d, want 1 (raw bit)", v.Bit(0))
	}
	got := v.Bits(0, 1)
	if got != 3 {
		t.Errorf("I3(-1).Bits(0,1) = %d, want 3 (raw bits)", got)
	}
}

func TestI3_SetBitBits(t *testing.T) {
	v := I3(0)
	v.SetBit(0, -1)
	if v != 1 {
		t.Errorf("SetBit(0,-1) = %d, want 1", v)
	}
}

func TestI3_RbitRbits(t *testing.T) {
	v := I3(0)
	v.Bitsref(0, 1).Set(3)
	if v != 3 {
		t.Errorf("v.Bitsref(0, 1).Set(3) => %d, want 3", v)
	}
}

func TestU8_Overflow(t *testing.T) {
	v := U8(255).Add(1)
	if v != 0 {
		t.Errorf("u8(255)+1 = %d, want 0", v)
	}
}

func TestI8_Overflow(t *testing.T) {
	v := I8(127).Inc()
	if v != -128 {
		t.Errorf("I8(127).Inc() = %d, want -128", v)
	}
}

func TestU64_MaxAdd(t *testing.T) {
	const maxU64 = ^U64(0)
	v := maxU64.Add(1)
	if v != 0 {
		t.Errorf("u64 max+1 != 0")
	}
}

func TestI64_MinDec(t *testing.T) {
	const minI64 = I64(-1 << 63)
	v := minI64.Dec()
	if v != I64(1<<63-1) {
		t.Errorf("I64 min-1 did not wrap to max: got %d", v)
	}
}

func TestRange_SetIncDec(t *testing.T) {
	var v U8
	r := bits(&v, 2, 4)

	r.Set(0b111)
	if v != 0b_0001_1100 {
		t.Errorf("Set(0b111) => %08b, want 00011100", v)
	}
	r.Inc()
	if v != 0 {
		t.Errorf("Inc() from max => %08b, want 0", v)
	}
	r.Dec()
	if v != 0b_0001_1100 {
		t.Errorf("Dec() from 0 => %08b, want 00011100", v)
	}
}

func TestRange_ArithOps(t *testing.T) {
	var v U8 = 0b_0000_0110
	r := bits(&v, 1, 2)

	r.Add(1)
	if v != 0 {
		t.Errorf("Add(1) overflow => %08b, want 0", v)
	}

	v = 0b_0000_0110
	r2 := bits(&v, 1, 2)
	r2.Sub(1)
	if v != 0b_0000_0100 {
		t.Errorf("Sub(1) => %08b, want 00000100", v)
	}
}

func TestRange_BitwiseOps(t *testing.T) {
	var v U8 = 0b_1111_1111
	r := bits(&v, 2, 5)

	r.And(0b0101)
	if v != 0b_1101_0111 {
		t.Errorf("And(0101) => %08b, want 11010111", v)
	}

	var v2 U8 = 0
	r2 := bits(&v2, 2, 5)
	r2.Or(0b1010)
	if v2 != 0b_0010_1000 {
		t.Errorf("Or(1010) => %08b, want 00101000", v2)
	}

	var v3 U8 = 0b_0011_1100
	r3 := bits(&v3, 2, 5)
	r3.Xor(0b1111)
	if v3 != 0 {
		t.Errorf("Xor(1111) => %08b, want 0", v3)
	}
}

func TestRange_Assign(t *testing.T) {
	var dst, src U8
	src = 0b_0000_0110
	s := bits(&src, 1, 2)
	d := bits(&dst, 4, 5)
	d.Assign(&s)
	if dst != 0b_0011_0000 {
		t.Errorf("Assign => %08b, want 00110000", dst)
	}
}

func TestRange_MulDiv(t *testing.T) {
	var v U8 = 0b_0000_0110
	r := bits(&v, 1, 2)
	r.Mul(2)
	if v != 0b_0000_0100 {
		t.Errorf("Mul(2) => %08b, want 00000100", v)
	}

	v = 0b_0000_0110
	r2 := bits(&v, 1, 2)
	r2.Div(3)
	if v != 0b_0000_0010 {
		t.Errorf("Div(3) => %08b, want 00000010", v)
	}
}

func TestRange_ShlShr(t *testing.T) {
	var v U8 = 0b_0000_0110
	r := bits(&v, 1, 2)
	r.Shl(1)
	if v != 0b_0000_0100 {
		t.Errorf("Shl(1) => %08b, want 00000100", v)
	}
}

func TestBit_SingleBit(t *testing.T) {
	var v U8 = 0b_1010_1010
	for i := range 8 {
		r := bit(&v, i)
		want := (v >> U8(i)) & 1
		if r.Val() != want {
			t.Errorf("Bit(%d).Val() = %d, want %d", i, r.Val(), want)
		}
	}
}

func TestU5_WrapAll(t *testing.T) {
	const N = 5
	const max = 1<<N - 1
	for a := int64(0); a <= max; a++ {
		for b := int64(0); b <= max; b++ {
			ua, ub := U5(a), U5(b)
			if int64(ua.Add(ub)) != wrapU(a+b, N) {
				t.Errorf("U5 Add(%d,%d) wrong", a, b)
			}
			if int64(ua.Xor(ub)) != (a ^ b) {
				t.Errorf("U5 Xor(%d,%d) wrong", a, b)
			}
		}
	}
}

func TestI5_WrapAll(t *testing.T) {
	const N = 5
	const hi = 1<<(N-1) - 1
	const lo = -(1 << (N - 1))
	for a := int64(lo); a <= hi; a++ {
		for b := int64(lo); b <= hi; b++ {
			ia, ib := I5(a), I5(b)
			got := int64(ia.Add(ib))
			want := wrapI(a+b, N)
			if got != want {
				t.Errorf("I5 Add(%d,%d) = %d, want %d", a, b, got, want)
			}
		}
	}
}

func TestU16_Bits(t *testing.T) {
	v := U16(0xABCD)
	if v.Bits(4, 7) != 0xC {
		t.Errorf("u16 Bits(4,7) of 0xABCD = %x, want 0xC", v.Bits(4, 7))
	}
	v.SetBits(4, 7, 0xF)
	if v != 0xABFD {
		t.Errorf("u16 SetBits(4,7,0xF) of 0xABCD = %x, want 0xABFD", v)
	}
}

func TestI16_Cast(t *testing.T) {
	v := I16(-32768)
	if v.cast() != -32768 {
		t.Errorf("I16(-32768).cast() = %d", v.cast())
	}
	v2 := I16(32767)
	if v2.cast() != 32767 {
		t.Errorf("I16(32767).cast() = %d", v2.cast())
	}
}

func TestU32_SetBitHighBit(t *testing.T) {
	v := U32(0)
	v.SetBit(31, 1)
	if v != 0x80000000 {
		t.Errorf("u32 SetBit(31,1) = %x, want 80000000", v)
	}
}

func TestI32_Shr(t *testing.T) {
	v := I32(-8)
	got := v.Shr(1)
	if got != -4 {
		t.Errorf("I32(-8).Shr(1) = %d, want -4", got)
	}
}

func TestU1_ExhaustiveAll(t *testing.T) {
	for a := U1(0); ; a++ {
		for b := U1(0); ; b++ {
			if a.Add(b) != U1(wrapU(int64(a)+int64(b), 1)) {
				t.Errorf("u1 Add(%d,%d)", a, b)
			}
			if a.Xor(b) != U1(int64(a)^int64(b)) {
				t.Errorf("u1 Xor(%d,%d)", a, b)
			}
			if b == 1 {
				break
			}
		}
		if a == 1 {
			break
		}
	}
}

func TestU3_Clip(t *testing.T) {
	tests := []struct {
		v    U3
		bits int
		want U3
	}{
		{0b111, 3, 0b111},
		{0b111, 2, 0b011},
		{0b111, 1, 0b001},
		{0b101, 2, 0b001},
	}
	for _, tt := range tests {
		if got := tt.v.Clip(tt.bits); got != tt.want {
			t.Errorf("u3(%03b).Clip(%d) = %03b, want %03b", tt.v, tt.bits, got, tt.want)
		}
	}
}

func TestU8_Clip(t *testing.T) {
	tests := []struct {
		v    U8
		bits int
		want U8
	}{
		{0xAB, 8, 0xAB},
		{0xAB, 4, 0x0B},
		{0xAB, 6, 0x2B},
		{0xFF, 1, 0x01},
	}
	for _, tt := range tests {
		if got := tt.v.Clip(tt.bits); got != tt.want {
			t.Errorf("u8(%#x).Clip(%d) = %#x, want %#x", tt.v, tt.bits, got, tt.want)
		}
	}
}

func TestU16_Clip(t *testing.T) {
	tests := []struct {
		v    U16
		bits int
		want U16
	}{
		{0xABCD, 16, 0xABCD},
		{0xABCD, 12, 0x0BCD},
		{0xABCD, 8, 0x00CD},
		{0xABCD, 4, 0x000D},
	}
	for _, tt := range tests {
		if got := tt.v.Clip(tt.bits); got != tt.want {
			t.Errorf("u16(%#x).Clip(%d) = %#x, want %#x", tt.v, tt.bits, got, tt.want)
		}
	}
}

func TestU32_Clip(t *testing.T) {
	v := U32(0xDEADBEEF)
	if got := v.Clip(16); got != 0xBEEF {
		t.Errorf("u32.Clip(16) = %#x, want 0xBEEF", got)
	}
	if got := v.Clip(32); got != 0xDEADBEEF {
		t.Errorf("u32.Clip(32) = %#x, want 0xDEADBEEF", got)
	}
}

func TestI8_Clip(t *testing.T) {
	tests := []struct {
		v    I8
		bits int
		want I8
	}{
		{31, 5, -1},
		{15, 5, 15},
		{-1, 8, -1},
		{0x55, 4, 5},
		{-120, 4, -8},
		{1, 1, -1},
		{0, 1, 0},
	}
	for _, tt := range tests {
		if got := tt.v.Clip(tt.bits); got != tt.want {
			t.Errorf("I8(%d).Clip(%d) = %d, want %d", tt.v, tt.bits, got, tt.want)
		}
	}
}

func TestI16_Clip(t *testing.T) {
	tests := []struct {
		v    I16
		bits int
		want I16
	}{
		{-1, 16, -1},
		{0x7FFF, 16, 0x7FFF},
		{0x00FF, 8, -1},
		{0x00FF, 9, 0xFF},
		{0x0100, 9, -256},
	}
	for _, tt := range tests {
		if got := tt.v.Clip(tt.bits); got != tt.want {
			t.Errorf("I16(%#x).Clip(%d) = %d, want %d", tt.v, tt.bits, got, tt.want)
		}
	}
}

func TestRange_Bit(t *testing.T) {
	v := U8(0b_0110_1010)
	r := bits(&v, 2, 5)
	if got := r.Bit(); got != 10 {
		t.Errorf("Range.Bit() = %04b, want 1010", got)
	}
	if r.Bit() != r.Val() {
		t.Errorf("Range.Bit() = %d != Range.Val() = %d", r.Bit(), r.Val())
	}

	v2 := U8(0b_0000_0100)
	rb := bit(&v2, 2)
	if got := rb.Bit(); got != 1 {
		t.Errorf("single-bit Range.Bit() = %d, want 1", got)
	}
	rb0 := bit(&v2, 0)
	if got := rb0.Bit(); got != 0 {
		t.Errorf("single-bit Range.Bit() at unset position = %d, want 0", got)
	}
}

func TestRange_Mod(t *testing.T) {
	v := U8(0b_1111_0000)
	r := bits(&v, 0, 3)
	r.Set(10)
	r.Mod(3)
	if r.Val() != 1 {
		t.Errorf("Range.Mod(3) of 10 = %d, want 1", r.Val())
	}
	if v>>4 != 0b1111 {
		t.Errorf("Mod touched outside bits: %08b", v)
	}

	v2 := U8(25)
	r2 := bits(&v2, 0, 7)
	r2.Mod(7)
	if v2 != 4 {
		t.Errorf("full-width Mod: got %d, want 4", v2)
	}
}

func TestRange_Shr(t *testing.T) {
	v := U8(0b_1010_0000)
	r := bits(&v, 0, 3)
	r.Set(0b1100)
	r.Shr(2)
	if r.Val() != 0b0011 {
		t.Errorf("Range.Shr(2) of 0b1100 = %04b, want 0011", r.Val())
	}
	if v>>4 != 0b1010 {
		t.Errorf("Shr touched outside bits: %08b", v)
	}

	r.Set(0b1111)
	r.Shr(4)
	if r.Val() != 0 {
		t.Errorf("Range.Shr(4) of 0b1111 = %d, want 0", r.Val())
	}
}

func TestBits_NegativeIndices(t *testing.T) {
	v := U8(0xA6)

	if got, want := v.Bits(-6, -3), v.Bits(2, 5); got != want {
		t.Errorf("Bits(-6,-3) = %04b, want %04b (== Bits(2,5))", got, want)
	}
	if got := v.Bits(-6, -3); got != 0b1001 {
		t.Errorf("Bits(-6,-3) of %08b = %04b, want 1001", v, got)
	}

	if got := v.Bit(-1); got != 1 {
		t.Errorf("Bit(-1) of %08b = %d, want 1", v, got)
	}
	if got := v.Bit(-8); got != 0 {
		t.Errorf("Bit(-8) of %08b = %d, want 0", v, got)
	}
}

func TestBits_MixedSignIndices(t *testing.T) {
	v := U8(0xA6)

	if got := v.Bits(2, -3); got != 0b1001 {
		t.Errorf("Bits(2,-3) of %08b = %04b, want 1001", v, got)
	}
	if got := v.Bits(-6, 5); got != 0b1001 {
		t.Errorf("Bits(-6,5) of %08b = %04b, want 1001", v, got)
	}
}

func TestBits_SwappedOrder(t *testing.T) {
	v := U8(0b_0110_1100)
	if got, want := v.Bits(5, 2), v.Bits(2, 5); got != want {
		t.Errorf("Bits(5,2) = %04b, want %04b (== Bits(2,5))", got, want)
	}
	if got := v.Bits(5, 2); got != 0b1011 {
		t.Errorf("Bits(5,2) of %08b = %04b, want 1011", v, got)
	}

	if got, want := v.Bits(-3, -6), v.Bits(2, 5); got != want {
		t.Errorf("Bits(-3,-6) = %04b, want %04b", got, want)
	}

	v2 := U8(0)
	v2.SetBits(5, 2, 0b1111)
	if v2 != 0b_0011_1100 {
		t.Errorf("SetBits(5,2,0xF) = %08b, want 00111100", v2)
	}
}

func TestU16_Byte(t *testing.T) {
	v := U16(0xABCD)
	if got := v.Byte(0); got != 0xCD {
		t.Errorf("u16(0xABCD).Byte(0) = %#x, want 0xCD", got)
	}
	if got := v.Byte(1); got != 0xAB {
		t.Errorf("u16(0xABCD).Byte(1) = %#x, want 0xAB", got)
	}
	v.SetByte(0, 0xFF)
	if v != 0xABFF {
		t.Errorf("u16.SetByte(0, 0xFF) = %#x, want 0xABFF", v)
	}
	v.SetByte(1, 0x12)
	if v != 0x12FF {
		t.Errorf("u16.SetByte(1, 0x12) = %#x, want 0x12FF", v)
	}
}

func TestU32_Byte(t *testing.T) {
	v := U32(0xDEADBEEF)
	wantBytes := [4]U32{0xEF, 0xBE, 0xAD, 0xDE}
	for i, want := range wantBytes {
		if got := v.Byte(i); got != want {
			t.Errorf("u32(0xDEADBEEF).Byte(%d) = %#x, want %#x", i, got, want)
		}
	}
	v.SetByte(2, 0xAA)
	if v != 0xDEAABEEF {
		t.Errorf("u32.SetByte(2,0xAA) = %#x, want 0xDEAABEEF", v)
	}
}

func TestU64_Byte(t *testing.T) {
	v := U64(0x0102030405060708)
	for i, want := range [8]U64{8, 7, 6, 5, 4, 3, 2, 1} {
		if got := v.Byte(i); got != want {
			t.Errorf("u64.Byte(%d) = %#x, want %#x", i, got, want)
		}
	}
	v.SetByte(7, 0xFF)
	if v != 0xFF02030405060708 {
		t.Errorf("u64.SetByte(7,0xFF) = %#x, want 0xFF02030405060708", v)
	}
}

func TestI16_Byte(t *testing.T) {
	v := I16(-1)
	if got := v.Byte(0); got != 0x00FF {
		t.Errorf("I16(-1).Byte(0) = %d, want 255", got)
	}
	if got := v.Byte(1); got != -1 {
		t.Errorf("I16(-1).Byte(1) = %d, want -1", got)
	}

	v2 := I16(0x7F00)
	if got := v2.Byte(1); got != 0x7F {
		t.Errorf("I16(0x7F00).Byte(1) = %d, want 0x7F", got)
	}

	v3 := I16(0)
	v3.SetByte(1, 0x12)
	if v3 != 0x1200 {
		t.Errorf("I16.SetByte(1,0x12) = %#x, want 0x1200", v3)
	}
}

func TestI32_Byte(t *testing.T) {
	v := I32(0x11223344)
	if got := v.Byte(2); got != 0x22 {
		t.Errorf("I32.Byte(2) = %#x, want 0x22", got)
	}
	v.SetByte(0, 0xFF)
	if v != 0x112233FF {
		t.Errorf("I32.SetByte(0,0xFF) = %#x, want 0x112233FF", v)
	}
}

func TestU32_Byteref(t *testing.T) {
	v := U32(0xDEADBEEF)
	r := v.Byteref(1)
	if got := r.Val(); got != 0xBE {
		t.Errorf("u32 Byteref(1).Val() = %#x, want 0xBE", got)
	}
	r.Set(0xAA)
	if v != 0xDEADAAEF {
		t.Errorf("u32 Byteref(1).Set(0xAA) = %#x, want 0xDEADAAEF", v)
	}
}

func TestByte_Generic(t *testing.T) {
	v := U32(0x11223344)
	r := byte(&v, 2)
	if got := r.Val(); got != 0x22 {
		t.Errorf("Byte(&v, 2).Val() = %#x, want 0x22", got)
	}
	r.Set(0xEE)
	if v != 0x11EE3344 {
		t.Errorf("Byte(&v, 2).Set(0xEE) = %#x, want 0x11EE3344", v)
	}
}

func BenchmarkU1_Add(b *testing.B) {
	v := U1(1)
	for b.Loop() {
		v = v.Add(1)
	}
}

func BenchmarkI1_Add(b *testing.B) {
	v := I1(-1)
	for b.Loop() {
		v = v.Add(-1)
	}
}

func BenchmarkU1_Inc(b *testing.B) {
	v := U1(0)
	for b.Loop() {
		v = v.Inc()
	}
}

func BenchmarkI1_Inc(b *testing.B) {
	v := I1(0)
	for b.Loop() {
		v = v.Inc()
	}
}

func BenchmarkU1_Not(b *testing.B) {
	v := U1(0)
	for b.Loop() {
		v = v.Not()
	}
}

func BenchmarkI1_Not(b *testing.B) {
	v := I1(0)
	for b.Loop() {
		v = v.Not()
	}
}

func BenchmarkU1_Xor(b *testing.B) {
	v := U1(1)
	for b.Loop() {
		v = v.Xor(1)
	}
}

func BenchmarkI1_Xor(b *testing.B) {
	v := I1(-1)
	for b.Loop() {
		v = v.Xor(-1)
	}
}

func BenchmarkU1_Clamp(b *testing.B) {
	v := U1(0)
	for b.Loop() {
		v = v.Clamp(99)
	}
}

func BenchmarkU1_Clip(b *testing.B) {
	v := U1(1)
	for b.Loop() {
		v = v.Clip(1)
	}
}

func BenchmarkI3_Clamp(b *testing.B) {
	v := I3(0)
	for b.Loop() {
		v = v.Clamp(int64(3))
	}
}

func BenchmarkI3_Clip(b *testing.B) {
	v := I3(-1)
	var r I3
	for b.Loop() {
		r = v.Clip(2)
	}
	_ = r
}

func BenchmarkU1_bit(b *testing.B) {
	v := U1(1)
	var r U1
	for b.Loop() {
		r = v.Bit(0)
	}
	_ = r
}

func BenchmarkI1_bit(b *testing.B) {
	v := I1(-1)
	var r I1
	for b.Loop() {
		r = v.Bit(0)
	}
	_ = r
}

func BenchmarkU1_bits(b *testing.B) {
	v := U1(1)
	var r U1
	for b.Loop() {
		r = v.Bits(0, 0)
	}
	_ = r
}

func BenchmarkI1_bits(b *testing.B) {
	v := I1(-1)
	var r I1
	for b.Loop() {
		r = v.Bits(0, 0)
	}
	_ = r
}

func BenchmarkU1_setBit(b *testing.B) {
	v := U1(0)
	for b.Loop() {
		v.SetBit(0, 1)
	}
}

func BenchmarkI1_setBit(b *testing.B) {
	v := I1(0)
	for b.Loop() {
		v.SetBit(0, -1)
	}
}

func BenchmarkU1_setBits(b *testing.B) {
	v := U1(0)
	for b.Loop() {
		v.SetBits(0, 0, 1)
	}
}

func BenchmarkI1_setBits(b *testing.B) {
	v := I1(0)
	for b.Loop() {
		v.SetBits(0, 0, -1)
	}
}

func BenchmarkU1_BitRef(b *testing.B) {
	v := U1(1)
	for b.Loop() {
		r := v.Bitref(0)
		_ = r.Val()
	}
}

func BenchmarkI1_BitRef(b *testing.B) {
	v := I1(-1)
	for b.Loop() {
		r := v.Bitref(0)
		_ = r.Val()
	}
}

func BenchmarkU1_BitsRef(b *testing.B) {
	v := U1(1)
	for b.Loop() {
		r := v.Bitsref(0, 0)
		_ = r.Val()
	}
}

func BenchmarkI1_BitsRef(b *testing.B) {
	v := I1(-1)
	for b.Loop() {
		r := v.Bitsref(0, 0)
		_ = r.Val()
	}
}

func BenchmarkU1_signed(b *testing.B) {
	v := U1(1)
	var r I1
	for b.Loop() {
		r = v.Signed()
	}
	_ = r
}

func BenchmarkI1_unsigned(b *testing.B) {
	v := I1(-1)
	var r U1
	for b.Loop() {
		r = v.Unsigned()
	}
	_ = r
}

func BenchmarkRange_Set(b *testing.B) {
	var v U8
	r := bits(&v, 2, 5)
	var i uint8
	for b.Loop() {
		r.Set(U8(i))
		i++
	}
}

func BenchmarkRange_Val(b *testing.B) {
	var v U8 = 0b_0011_1100
	r := bits(&v, 2, 5)
	for b.Loop() {
		r.Val()
	}
}

func BenchmarkRange_Inc(b *testing.B) {
	var v U8
	r := bits(&v, 2, 5)
	for b.Loop() {
		r.Inc()
	}
}
