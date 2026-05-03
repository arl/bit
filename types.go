// Generated Code; DO NOT EDIT.

package bit

import (
	"fmt"
	"strconv"
)

// U1 is an 1-bit unsigned integer.
type U1 uint8

// AsU1 returns a U1 representing v reduced to 1 bits.
func AsU1[T Unsigned](v T) U1 { return U1.cast(U1(v)) }

func (U1) nbits() int { return 1 }
func (u U1) mask() U1 { return 0x1 }
func (u U1) cast() U1 { return u & u.mask() }

// Set assigns v reduced to 1 bits to u.
func (u *U1) Set(v U1) U1 { *u = new(U1(v)).cast(); return U1.cast(v) }

// Add returns u+o reduced to 1 bits.
func (u U1) Add(o U1) U1 { return U1.cast(u + o) }

// Sub returns u-o reduced to 1 bits.
func (u U1) Sub(o U1) U1 { return U1.cast(u - o) }

// Inc returns u+1 reduced to 1 bits.
func (u U1) Inc() U1 { return U1.cast(u + 1) }

// Dec returns u-1 reduced to 1 bits.
func (u U1) Dec() U1 { return U1.cast(u - 1) }

// Mul returns u*o reduced to 1 bits.
func (u U1) Mul(o U1) U1 { return U1.cast(u * o) }

// Div returns u divided by o.
func (u U1) Div(o U1) U1 { return U1.cast(u / o) }

// Mod returns u modulo o.
func (u U1) Mod(o U1) U1 { return U1.cast(u % o) }

// And returns u&o reduced to 1 bits.
func (u U1) And(o U1) U1 { return U1.cast(u & o) }

// Or returns u|o reduced to 1 bits.
func (u U1) Or(o U1) U1 { return U1.cast(u | o) }

// Xor returns u^o reduced to 1 bits.
func (u U1) Xor(o U1) U1 { return U1.cast(u ^ o) }

// Shr returns u>>o reduced to 1 bits.
func (u U1) Shr(o U1) U1 { return U1.cast(u >> o) }

// Shl returns u<<o reduced to 1 bits.
func (u U1) Shl(o U1) U1 { return U1.cast(u << o) }

// Not returns bitwise complement of u reduced to 1 bits.
func (u U1) Not() U1 { return U1.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U1) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U1) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U1 value.
func (u U1) Signed() I1 { return I1.cast(I1(u)) }

// Clamp returns value saturated into the representable range of U1.
func (u U1) Clamp(value uint64) U1 { return uclamp[U1](value) }

// Clip masks u to a bits-wide low field.
func (u U1) Clip(bits int) U1 {
	b := 1 << (bits - 1)
	m := U1(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U1) Bit(index int) U1 {
	if index < 0 {
		index = 1 + index
	}
	return (u >> U1(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U1) SetBit(index int, v U1) {
	if index < 0 {
		index = 1 + index
	}
	bit := U1(1) << U1(index)
	*u = U1.cast((*u &^ bit) | ((v & 1) << U1(index)))
}

// Bitref returns a Range for bit index.
func (u *U1) Bitref(index int) Range[U1] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U1) Bits(lo, hi int) U1 {
	p := 1
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U1((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U1(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U1) SetBits(lo, hi int, v U1) {
	p := 1
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U1((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U1.cast((*u &^ mask) | ((v << U1(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U1) Bitsref(lo, hi int) Range[U1] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U1) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U1) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U1(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U1) Byteref(idx int) Range[U1] { return byte(u, idx) }

// I1 is an 1-bit signed integer in two's complement.
type I1 int8

// AsI1 returns a I1 representing v reduced to 1 bits.
func AsI1[T Signed](v T) I1 { return I1.cast(I1(v)) }

func (I1) nbits() int { return 1 }
func (i I1) mask() I1 { return 0x1 }
func (i I1) sign() I1 { return 1 << (i.nbits() - 1) }
func (i I1) cast() I1 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 1 bits to u.
func (i *I1) Set(v I1) I1 { *i = new(I1(v)).cast(); return I1.cast(v) }

// Add returns i+o reduced to 1 bits.
func (i I1) Add(o I1) I1 { return I1.cast(i + o) }

// Sub returns i-o reduced to 1 bits.
func (i I1) Sub(o I1) I1 { return I1.cast(i - o) }

// Inc returns i+1 reduced to 1 bits.
func (i I1) Inc() I1 { return I1.cast(i + 1) }

// Dec returns i-1 reduced to 1 bits.
func (i I1) Dec() I1 { return I1.cast(i - 1) }

// Mul returns i*o reduced to 1 bits.
func (i I1) Mul(o I1) I1 { return I1.cast(i * o) }

// Div returns i divided by o.
func (i I1) Div(o I1) I1 { return I1.cast(i / o) }

// Mod returns i modulo o.
func (i I1) Mod(o I1) I1 { return I1.cast(i % o) }

// And returns i&o reduced to 1 bits.
func (i I1) And(o I1) I1 { return I1.cast(i & o) }

// Or returns i|o reduced to 1 bits.
func (i I1) Or(o I1) I1 { return I1.cast(i | o) }

// Xor returns i^o reduced to 1 bits.
func (i I1) Xor(o I1) I1 { return I1.cast(i ^ o) }

// Shr returns i>>o reduced to 1 bits.
func (i I1) Shr(o I1) I1 { return I1.cast(i >> o) }

// Shl returns i<<o reduced to 1 bits.
func (i I1) Shl(o I1) I1 { return I1.cast(i << o) }

// Not returns bitwise complement of i reduced to 1 bits.
func (i I1) Not() I1 { return I1.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I1) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I1) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U1 bit pattern.
func (i I1) Unsigned() U1 { return U1.cast(U1(i)) }

// Clamp returns value saturated into the representable range of I1.
func (i I1) Clamp(value int64) I1 { return iclamp[I1](value) }

// Clip masks i to a bits-wide low field and sign-extends to 1 bits.
func (i I1) Clip(bits int) I1 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I1(uint64(i)&m^b) - I1(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I1) Bit(index int) I1 {
	if index < 0 {
		index = 1 + index
	}
	return (i >> I1(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I1) SetBit(index int, v I1) {
	if index < 0 {
		index = 1 + index
	}
	bit := I1(1) << I1(index)
	*i = I1.cast((*i &^ bit) | ((v & 1) << I1(index)))
}

// Bitref returns a Range for bit index.
func (i *I1) Bitref(index int) Range[I1] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I1) Bits(lo, hi int) I1 {
	p := 1
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I1((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I1(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I1) SetBits(lo, hi int, v I1) {
	p := 1
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I1((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I1.cast((*i &^ mask) | ((v << I1(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I1) Bitsref(lo, hi int) Range[I1] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I1) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I1) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I1(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I1) Byteref(idx int) Range[I1] { return byte(i, idx) }

// U2 is an 2-bit unsigned integer.
type U2 uint8

// AsU2 returns a U2 representing v reduced to 2 bits.
func AsU2[T Unsigned](v T) U2 { return U2.cast(U2(v)) }

func (U2) nbits() int { return 2 }
func (u U2) mask() U2 { return 0x3 }
func (u U2) cast() U2 { return u & u.mask() }

// Set assigns v reduced to 2 bits to u.
func (u *U2) Set(v U2) U2 { *u = new(U2(v)).cast(); return U2.cast(v) }

// Add returns u+o reduced to 2 bits.
func (u U2) Add(o U2) U2 { return U2.cast(u + o) }

// Sub returns u-o reduced to 2 bits.
func (u U2) Sub(o U2) U2 { return U2.cast(u - o) }

// Inc returns u+1 reduced to 2 bits.
func (u U2) Inc() U2 { return U2.cast(u + 1) }

// Dec returns u-1 reduced to 2 bits.
func (u U2) Dec() U2 { return U2.cast(u - 1) }

// Mul returns u*o reduced to 2 bits.
func (u U2) Mul(o U2) U2 { return U2.cast(u * o) }

// Div returns u divided by o.
func (u U2) Div(o U2) U2 { return U2.cast(u / o) }

// Mod returns u modulo o.
func (u U2) Mod(o U2) U2 { return U2.cast(u % o) }

// And returns u&o reduced to 2 bits.
func (u U2) And(o U2) U2 { return U2.cast(u & o) }

// Or returns u|o reduced to 2 bits.
func (u U2) Or(o U2) U2 { return U2.cast(u | o) }

// Xor returns u^o reduced to 2 bits.
func (u U2) Xor(o U2) U2 { return U2.cast(u ^ o) }

// Shr returns u>>o reduced to 2 bits.
func (u U2) Shr(o U2) U2 { return U2.cast(u >> o) }

// Shl returns u<<o reduced to 2 bits.
func (u U2) Shl(o U2) U2 { return U2.cast(u << o) }

// Not returns bitwise complement of u reduced to 2 bits.
func (u U2) Not() U2 { return U2.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U2) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U2) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U2 value.
func (u U2) Signed() I2 { return I2.cast(I2(u)) }

// Clamp returns value saturated into the representable range of U2.
func (u U2) Clamp(value uint64) U2 { return uclamp[U2](value) }

// Clip masks u to a bits-wide low field.
func (u U2) Clip(bits int) U2 {
	b := 1 << (bits - 1)
	m := U2(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U2) Bit(index int) U2 {
	if index < 0 {
		index = 2 + index
	}
	return (u >> U2(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U2) SetBit(index int, v U2) {
	if index < 0 {
		index = 2 + index
	}
	bit := U2(1) << U2(index)
	*u = U2.cast((*u &^ bit) | ((v & 1) << U2(index)))
}

// Bitref returns a Range for bit index.
func (u *U2) Bitref(index int) Range[U2] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U2) Bits(lo, hi int) U2 {
	p := 2
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U2((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U2(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U2) SetBits(lo, hi int, v U2) {
	p := 2
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U2((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U2.cast((*u &^ mask) | ((v << U2(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U2) Bitsref(lo, hi int) Range[U2] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U2) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U2) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U2(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U2) Byteref(idx int) Range[U2] { return byte(u, idx) }

// I2 is an 2-bit signed integer in two's complement.
type I2 int8

// AsI2 returns a I2 representing v reduced to 2 bits.
func AsI2[T Signed](v T) I2 { return I2.cast(I2(v)) }

func (I2) nbits() int { return 2 }
func (i I2) mask() I2 { return 0x3 }
func (i I2) sign() I2 { return 1 << (i.nbits() - 1) }
func (i I2) cast() I2 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 2 bits to u.
func (i *I2) Set(v I2) I2 { *i = new(I2(v)).cast(); return I2.cast(v) }

// Add returns i+o reduced to 2 bits.
func (i I2) Add(o I2) I2 { return I2.cast(i + o) }

// Sub returns i-o reduced to 2 bits.
func (i I2) Sub(o I2) I2 { return I2.cast(i - o) }

// Inc returns i+1 reduced to 2 bits.
func (i I2) Inc() I2 { return I2.cast(i + 1) }

// Dec returns i-1 reduced to 2 bits.
func (i I2) Dec() I2 { return I2.cast(i - 1) }

// Mul returns i*o reduced to 2 bits.
func (i I2) Mul(o I2) I2 { return I2.cast(i * o) }

// Div returns i divided by o.
func (i I2) Div(o I2) I2 { return I2.cast(i / o) }

// Mod returns i modulo o.
func (i I2) Mod(o I2) I2 { return I2.cast(i % o) }

// And returns i&o reduced to 2 bits.
func (i I2) And(o I2) I2 { return I2.cast(i & o) }

// Or returns i|o reduced to 2 bits.
func (i I2) Or(o I2) I2 { return I2.cast(i | o) }

// Xor returns i^o reduced to 2 bits.
func (i I2) Xor(o I2) I2 { return I2.cast(i ^ o) }

// Shr returns i>>o reduced to 2 bits.
func (i I2) Shr(o I2) I2 { return I2.cast(i >> o) }

// Shl returns i<<o reduced to 2 bits.
func (i I2) Shl(o I2) I2 { return I2.cast(i << o) }

// Not returns bitwise complement of i reduced to 2 bits.
func (i I2) Not() I2 { return I2.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I2) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I2) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U2 bit pattern.
func (i I2) Unsigned() U2 { return U2.cast(U2(i)) }

// Clamp returns value saturated into the representable range of I2.
func (i I2) Clamp(value int64) I2 { return iclamp[I2](value) }

// Clip masks i to a bits-wide low field and sign-extends to 2 bits.
func (i I2) Clip(bits int) I2 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I2(uint64(i)&m^b) - I2(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I2) Bit(index int) I2 {
	if index < 0 {
		index = 2 + index
	}
	return (i >> I2(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I2) SetBit(index int, v I2) {
	if index < 0 {
		index = 2 + index
	}
	bit := I2(1) << I2(index)
	*i = I2.cast((*i &^ bit) | ((v & 1) << I2(index)))
}

// Bitref returns a Range for bit index.
func (i *I2) Bitref(index int) Range[I2] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I2) Bits(lo, hi int) I2 {
	p := 2
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I2((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I2(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I2) SetBits(lo, hi int, v I2) {
	p := 2
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I2((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I2.cast((*i &^ mask) | ((v << I2(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I2) Bitsref(lo, hi int) Range[I2] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I2) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I2) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I2(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I2) Byteref(idx int) Range[I2] { return byte(i, idx) }

// U3 is an 3-bit unsigned integer.
type U3 uint8

// AsU3 returns a U3 representing v reduced to 3 bits.
func AsU3[T Unsigned](v T) U3 { return U3.cast(U3(v)) }

func (U3) nbits() int { return 3 }
func (u U3) mask() U3 { return 0x7 }
func (u U3) cast() U3 { return u & u.mask() }

// Set assigns v reduced to 3 bits to u.
func (u *U3) Set(v U3) U3 { *u = new(U3(v)).cast(); return U3.cast(v) }

// Add returns u+o reduced to 3 bits.
func (u U3) Add(o U3) U3 { return U3.cast(u + o) }

// Sub returns u-o reduced to 3 bits.
func (u U3) Sub(o U3) U3 { return U3.cast(u - o) }

// Inc returns u+1 reduced to 3 bits.
func (u U3) Inc() U3 { return U3.cast(u + 1) }

// Dec returns u-1 reduced to 3 bits.
func (u U3) Dec() U3 { return U3.cast(u - 1) }

// Mul returns u*o reduced to 3 bits.
func (u U3) Mul(o U3) U3 { return U3.cast(u * o) }

// Div returns u divided by o.
func (u U3) Div(o U3) U3 { return U3.cast(u / o) }

// Mod returns u modulo o.
func (u U3) Mod(o U3) U3 { return U3.cast(u % o) }

// And returns u&o reduced to 3 bits.
func (u U3) And(o U3) U3 { return U3.cast(u & o) }

// Or returns u|o reduced to 3 bits.
func (u U3) Or(o U3) U3 { return U3.cast(u | o) }

// Xor returns u^o reduced to 3 bits.
func (u U3) Xor(o U3) U3 { return U3.cast(u ^ o) }

// Shr returns u>>o reduced to 3 bits.
func (u U3) Shr(o U3) U3 { return U3.cast(u >> o) }

// Shl returns u<<o reduced to 3 bits.
func (u U3) Shl(o U3) U3 { return U3.cast(u << o) }

// Not returns bitwise complement of u reduced to 3 bits.
func (u U3) Not() U3 { return U3.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U3) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U3) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U3 value.
func (u U3) Signed() I3 { return I3.cast(I3(u)) }

// Clamp returns value saturated into the representable range of U3.
func (u U3) Clamp(value uint64) U3 { return uclamp[U3](value) }

// Clip masks u to a bits-wide low field.
func (u U3) Clip(bits int) U3 {
	b := 1 << (bits - 1)
	m := U3(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U3) Bit(index int) U3 {
	if index < 0 {
		index = 3 + index
	}
	return (u >> U3(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U3) SetBit(index int, v U3) {
	if index < 0 {
		index = 3 + index
	}
	bit := U3(1) << U3(index)
	*u = U3.cast((*u &^ bit) | ((v & 1) << U3(index)))
}

// Bitref returns a Range for bit index.
func (u *U3) Bitref(index int) Range[U3] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U3) Bits(lo, hi int) U3 {
	p := 3
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U3((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U3(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U3) SetBits(lo, hi int, v U3) {
	p := 3
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U3((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U3.cast((*u &^ mask) | ((v << U3(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U3) Bitsref(lo, hi int) Range[U3] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U3) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U3) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U3(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U3) Byteref(idx int) Range[U3] { return byte(u, idx) }

// I3 is an 3-bit signed integer in two's complement.
type I3 int8

// AsI3 returns a I3 representing v reduced to 3 bits.
func AsI3[T Signed](v T) I3 { return I3.cast(I3(v)) }

func (I3) nbits() int { return 3 }
func (i I3) mask() I3 { return 0x7 }
func (i I3) sign() I3 { return 1 << (i.nbits() - 1) }
func (i I3) cast() I3 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 3 bits to u.
func (i *I3) Set(v I3) I3 { *i = new(I3(v)).cast(); return I3.cast(v) }

// Add returns i+o reduced to 3 bits.
func (i I3) Add(o I3) I3 { return I3.cast(i + o) }

// Sub returns i-o reduced to 3 bits.
func (i I3) Sub(o I3) I3 { return I3.cast(i - o) }

// Inc returns i+1 reduced to 3 bits.
func (i I3) Inc() I3 { return I3.cast(i + 1) }

// Dec returns i-1 reduced to 3 bits.
func (i I3) Dec() I3 { return I3.cast(i - 1) }

// Mul returns i*o reduced to 3 bits.
func (i I3) Mul(o I3) I3 { return I3.cast(i * o) }

// Div returns i divided by o.
func (i I3) Div(o I3) I3 { return I3.cast(i / o) }

// Mod returns i modulo o.
func (i I3) Mod(o I3) I3 { return I3.cast(i % o) }

// And returns i&o reduced to 3 bits.
func (i I3) And(o I3) I3 { return I3.cast(i & o) }

// Or returns i|o reduced to 3 bits.
func (i I3) Or(o I3) I3 { return I3.cast(i | o) }

// Xor returns i^o reduced to 3 bits.
func (i I3) Xor(o I3) I3 { return I3.cast(i ^ o) }

// Shr returns i>>o reduced to 3 bits.
func (i I3) Shr(o I3) I3 { return I3.cast(i >> o) }

// Shl returns i<<o reduced to 3 bits.
func (i I3) Shl(o I3) I3 { return I3.cast(i << o) }

// Not returns bitwise complement of i reduced to 3 bits.
func (i I3) Not() I3 { return I3.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I3) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I3) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U3 bit pattern.
func (i I3) Unsigned() U3 { return U3.cast(U3(i)) }

// Clamp returns value saturated into the representable range of I3.
func (i I3) Clamp(value int64) I3 { return iclamp[I3](value) }

// Clip masks i to a bits-wide low field and sign-extends to 3 bits.
func (i I3) Clip(bits int) I3 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I3(uint64(i)&m^b) - I3(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I3) Bit(index int) I3 {
	if index < 0 {
		index = 3 + index
	}
	return (i >> I3(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I3) SetBit(index int, v I3) {
	if index < 0 {
		index = 3 + index
	}
	bit := I3(1) << I3(index)
	*i = I3.cast((*i &^ bit) | ((v & 1) << I3(index)))
}

// Bitref returns a Range for bit index.
func (i *I3) Bitref(index int) Range[I3] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I3) Bits(lo, hi int) I3 {
	p := 3
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I3((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I3(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I3) SetBits(lo, hi int, v I3) {
	p := 3
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I3((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I3.cast((*i &^ mask) | ((v << I3(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I3) Bitsref(lo, hi int) Range[I3] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I3) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I3) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I3(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I3) Byteref(idx int) Range[I3] { return byte(i, idx) }

// U4 is an 4-bit unsigned integer.
type U4 uint8

// AsU4 returns a U4 representing v reduced to 4 bits.
func AsU4[T Unsigned](v T) U4 { return U4.cast(U4(v)) }

func (U4) nbits() int { return 4 }
func (u U4) mask() U4 { return 0xf }
func (u U4) cast() U4 { return u & u.mask() }

// Set assigns v reduced to 4 bits to u.
func (u *U4) Set(v U4) U4 { *u = new(U4(v)).cast(); return U4.cast(v) }

// Add returns u+o reduced to 4 bits.
func (u U4) Add(o U4) U4 { return U4.cast(u + o) }

// Sub returns u-o reduced to 4 bits.
func (u U4) Sub(o U4) U4 { return U4.cast(u - o) }

// Inc returns u+1 reduced to 4 bits.
func (u U4) Inc() U4 { return U4.cast(u + 1) }

// Dec returns u-1 reduced to 4 bits.
func (u U4) Dec() U4 { return U4.cast(u - 1) }

// Mul returns u*o reduced to 4 bits.
func (u U4) Mul(o U4) U4 { return U4.cast(u * o) }

// Div returns u divided by o.
func (u U4) Div(o U4) U4 { return U4.cast(u / o) }

// Mod returns u modulo o.
func (u U4) Mod(o U4) U4 { return U4.cast(u % o) }

// And returns u&o reduced to 4 bits.
func (u U4) And(o U4) U4 { return U4.cast(u & o) }

// Or returns u|o reduced to 4 bits.
func (u U4) Or(o U4) U4 { return U4.cast(u | o) }

// Xor returns u^o reduced to 4 bits.
func (u U4) Xor(o U4) U4 { return U4.cast(u ^ o) }

// Shr returns u>>o reduced to 4 bits.
func (u U4) Shr(o U4) U4 { return U4.cast(u >> o) }

// Shl returns u<<o reduced to 4 bits.
func (u U4) Shl(o U4) U4 { return U4.cast(u << o) }

// Not returns bitwise complement of u reduced to 4 bits.
func (u U4) Not() U4 { return U4.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U4) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U4) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U4 value.
func (u U4) Signed() I4 { return I4.cast(I4(u)) }

// Clamp returns value saturated into the representable range of U4.
func (u U4) Clamp(value uint64) U4 { return uclamp[U4](value) }

// Clip masks u to a bits-wide low field.
func (u U4) Clip(bits int) U4 {
	b := 1 << (bits - 1)
	m := U4(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U4) Bit(index int) U4 {
	if index < 0 {
		index = 4 + index
	}
	return (u >> U4(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U4) SetBit(index int, v U4) {
	if index < 0 {
		index = 4 + index
	}
	bit := U4(1) << U4(index)
	*u = U4.cast((*u &^ bit) | ((v & 1) << U4(index)))
}

// Bitref returns a Range for bit index.
func (u *U4) Bitref(index int) Range[U4] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U4) Bits(lo, hi int) U4 {
	p := 4
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U4((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U4(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U4) SetBits(lo, hi int, v U4) {
	p := 4
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U4((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U4.cast((*u &^ mask) | ((v << U4(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U4) Bitsref(lo, hi int) Range[U4] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U4) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U4) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U4(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U4) Byteref(idx int) Range[U4] { return byte(u, idx) }

// I4 is an 4-bit signed integer in two's complement.
type I4 int8

// AsI4 returns a I4 representing v reduced to 4 bits.
func AsI4[T Signed](v T) I4 { return I4.cast(I4(v)) }

func (I4) nbits() int { return 4 }
func (i I4) mask() I4 { return 0xf }
func (i I4) sign() I4 { return 1 << (i.nbits() - 1) }
func (i I4) cast() I4 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 4 bits to u.
func (i *I4) Set(v I4) I4 { *i = new(I4(v)).cast(); return I4.cast(v) }

// Add returns i+o reduced to 4 bits.
func (i I4) Add(o I4) I4 { return I4.cast(i + o) }

// Sub returns i-o reduced to 4 bits.
func (i I4) Sub(o I4) I4 { return I4.cast(i - o) }

// Inc returns i+1 reduced to 4 bits.
func (i I4) Inc() I4 { return I4.cast(i + 1) }

// Dec returns i-1 reduced to 4 bits.
func (i I4) Dec() I4 { return I4.cast(i - 1) }

// Mul returns i*o reduced to 4 bits.
func (i I4) Mul(o I4) I4 { return I4.cast(i * o) }

// Div returns i divided by o.
func (i I4) Div(o I4) I4 { return I4.cast(i / o) }

// Mod returns i modulo o.
func (i I4) Mod(o I4) I4 { return I4.cast(i % o) }

// And returns i&o reduced to 4 bits.
func (i I4) And(o I4) I4 { return I4.cast(i & o) }

// Or returns i|o reduced to 4 bits.
func (i I4) Or(o I4) I4 { return I4.cast(i | o) }

// Xor returns i^o reduced to 4 bits.
func (i I4) Xor(o I4) I4 { return I4.cast(i ^ o) }

// Shr returns i>>o reduced to 4 bits.
func (i I4) Shr(o I4) I4 { return I4.cast(i >> o) }

// Shl returns i<<o reduced to 4 bits.
func (i I4) Shl(o I4) I4 { return I4.cast(i << o) }

// Not returns bitwise complement of i reduced to 4 bits.
func (i I4) Not() I4 { return I4.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I4) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I4) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U4 bit pattern.
func (i I4) Unsigned() U4 { return U4.cast(U4(i)) }

// Clamp returns value saturated into the representable range of I4.
func (i I4) Clamp(value int64) I4 { return iclamp[I4](value) }

// Clip masks i to a bits-wide low field and sign-extends to 4 bits.
func (i I4) Clip(bits int) I4 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I4(uint64(i)&m^b) - I4(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I4) Bit(index int) I4 {
	if index < 0 {
		index = 4 + index
	}
	return (i >> I4(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I4) SetBit(index int, v I4) {
	if index < 0 {
		index = 4 + index
	}
	bit := I4(1) << I4(index)
	*i = I4.cast((*i &^ bit) | ((v & 1) << I4(index)))
}

// Bitref returns a Range for bit index.
func (i *I4) Bitref(index int) Range[I4] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I4) Bits(lo, hi int) I4 {
	p := 4
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I4((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I4(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I4) SetBits(lo, hi int, v I4) {
	p := 4
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I4((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I4.cast((*i &^ mask) | ((v << I4(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I4) Bitsref(lo, hi int) Range[I4] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I4) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I4) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I4(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I4) Byteref(idx int) Range[I4] { return byte(i, idx) }

// U5 is an 5-bit unsigned integer.
type U5 uint8

// AsU5 returns a U5 representing v reduced to 5 bits.
func AsU5[T Unsigned](v T) U5 { return U5.cast(U5(v)) }

func (U5) nbits() int { return 5 }
func (u U5) mask() U5 { return 0x1f }
func (u U5) cast() U5 { return u & u.mask() }

// Set assigns v reduced to 5 bits to u.
func (u *U5) Set(v U5) U5 { *u = new(U5(v)).cast(); return U5.cast(v) }

// Add returns u+o reduced to 5 bits.
func (u U5) Add(o U5) U5 { return U5.cast(u + o) }

// Sub returns u-o reduced to 5 bits.
func (u U5) Sub(o U5) U5 { return U5.cast(u - o) }

// Inc returns u+1 reduced to 5 bits.
func (u U5) Inc() U5 { return U5.cast(u + 1) }

// Dec returns u-1 reduced to 5 bits.
func (u U5) Dec() U5 { return U5.cast(u - 1) }

// Mul returns u*o reduced to 5 bits.
func (u U5) Mul(o U5) U5 { return U5.cast(u * o) }

// Div returns u divided by o.
func (u U5) Div(o U5) U5 { return U5.cast(u / o) }

// Mod returns u modulo o.
func (u U5) Mod(o U5) U5 { return U5.cast(u % o) }

// And returns u&o reduced to 5 bits.
func (u U5) And(o U5) U5 { return U5.cast(u & o) }

// Or returns u|o reduced to 5 bits.
func (u U5) Or(o U5) U5 { return U5.cast(u | o) }

// Xor returns u^o reduced to 5 bits.
func (u U5) Xor(o U5) U5 { return U5.cast(u ^ o) }

// Shr returns u>>o reduced to 5 bits.
func (u U5) Shr(o U5) U5 { return U5.cast(u >> o) }

// Shl returns u<<o reduced to 5 bits.
func (u U5) Shl(o U5) U5 { return U5.cast(u << o) }

// Not returns bitwise complement of u reduced to 5 bits.
func (u U5) Not() U5 { return U5.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U5) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U5) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U5 value.
func (u U5) Signed() I5 { return I5.cast(I5(u)) }

// Clamp returns value saturated into the representable range of U5.
func (u U5) Clamp(value uint64) U5 { return uclamp[U5](value) }

// Clip masks u to a bits-wide low field.
func (u U5) Clip(bits int) U5 {
	b := 1 << (bits - 1)
	m := U5(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U5) Bit(index int) U5 {
	if index < 0 {
		index = 5 + index
	}
	return (u >> U5(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U5) SetBit(index int, v U5) {
	if index < 0 {
		index = 5 + index
	}
	bit := U5(1) << U5(index)
	*u = U5.cast((*u &^ bit) | ((v & 1) << U5(index)))
}

// Bitref returns a Range for bit index.
func (u *U5) Bitref(index int) Range[U5] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U5) Bits(lo, hi int) U5 {
	p := 5
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U5((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U5(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U5) SetBits(lo, hi int, v U5) {
	p := 5
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U5((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U5.cast((*u &^ mask) | ((v << U5(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U5) Bitsref(lo, hi int) Range[U5] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U5) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U5) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U5(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U5) Byteref(idx int) Range[U5] { return byte(u, idx) }

// I5 is an 5-bit signed integer in two's complement.
type I5 int8

// AsI5 returns a I5 representing v reduced to 5 bits.
func AsI5[T Signed](v T) I5 { return I5.cast(I5(v)) }

func (I5) nbits() int { return 5 }
func (i I5) mask() I5 { return 0x1f }
func (i I5) sign() I5 { return 1 << (i.nbits() - 1) }
func (i I5) cast() I5 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 5 bits to u.
func (i *I5) Set(v I5) I5 { *i = new(I5(v)).cast(); return I5.cast(v) }

// Add returns i+o reduced to 5 bits.
func (i I5) Add(o I5) I5 { return I5.cast(i + o) }

// Sub returns i-o reduced to 5 bits.
func (i I5) Sub(o I5) I5 { return I5.cast(i - o) }

// Inc returns i+1 reduced to 5 bits.
func (i I5) Inc() I5 { return I5.cast(i + 1) }

// Dec returns i-1 reduced to 5 bits.
func (i I5) Dec() I5 { return I5.cast(i - 1) }

// Mul returns i*o reduced to 5 bits.
func (i I5) Mul(o I5) I5 { return I5.cast(i * o) }

// Div returns i divided by o.
func (i I5) Div(o I5) I5 { return I5.cast(i / o) }

// Mod returns i modulo o.
func (i I5) Mod(o I5) I5 { return I5.cast(i % o) }

// And returns i&o reduced to 5 bits.
func (i I5) And(o I5) I5 { return I5.cast(i & o) }

// Or returns i|o reduced to 5 bits.
func (i I5) Or(o I5) I5 { return I5.cast(i | o) }

// Xor returns i^o reduced to 5 bits.
func (i I5) Xor(o I5) I5 { return I5.cast(i ^ o) }

// Shr returns i>>o reduced to 5 bits.
func (i I5) Shr(o I5) I5 { return I5.cast(i >> o) }

// Shl returns i<<o reduced to 5 bits.
func (i I5) Shl(o I5) I5 { return I5.cast(i << o) }

// Not returns bitwise complement of i reduced to 5 bits.
func (i I5) Not() I5 { return I5.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I5) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I5) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U5 bit pattern.
func (i I5) Unsigned() U5 { return U5.cast(U5(i)) }

// Clamp returns value saturated into the representable range of I5.
func (i I5) Clamp(value int64) I5 { return iclamp[I5](value) }

// Clip masks i to a bits-wide low field and sign-extends to 5 bits.
func (i I5) Clip(bits int) I5 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I5(uint64(i)&m^b) - I5(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I5) Bit(index int) I5 {
	if index < 0 {
		index = 5 + index
	}
	return (i >> I5(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I5) SetBit(index int, v I5) {
	if index < 0 {
		index = 5 + index
	}
	bit := I5(1) << I5(index)
	*i = I5.cast((*i &^ bit) | ((v & 1) << I5(index)))
}

// Bitref returns a Range for bit index.
func (i *I5) Bitref(index int) Range[I5] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I5) Bits(lo, hi int) I5 {
	p := 5
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I5((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I5(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I5) SetBits(lo, hi int, v I5) {
	p := 5
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I5((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I5.cast((*i &^ mask) | ((v << I5(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I5) Bitsref(lo, hi int) Range[I5] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I5) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I5) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I5(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I5) Byteref(idx int) Range[I5] { return byte(i, idx) }

// U6 is an 6-bit unsigned integer.
type U6 uint8

// AsU6 returns a U6 representing v reduced to 6 bits.
func AsU6[T Unsigned](v T) U6 { return U6.cast(U6(v)) }

func (U6) nbits() int { return 6 }
func (u U6) mask() U6 { return 0x3f }
func (u U6) cast() U6 { return u & u.mask() }

// Set assigns v reduced to 6 bits to u.
func (u *U6) Set(v U6) U6 { *u = new(U6(v)).cast(); return U6.cast(v) }

// Add returns u+o reduced to 6 bits.
func (u U6) Add(o U6) U6 { return U6.cast(u + o) }

// Sub returns u-o reduced to 6 bits.
func (u U6) Sub(o U6) U6 { return U6.cast(u - o) }

// Inc returns u+1 reduced to 6 bits.
func (u U6) Inc() U6 { return U6.cast(u + 1) }

// Dec returns u-1 reduced to 6 bits.
func (u U6) Dec() U6 { return U6.cast(u - 1) }

// Mul returns u*o reduced to 6 bits.
func (u U6) Mul(o U6) U6 { return U6.cast(u * o) }

// Div returns u divided by o.
func (u U6) Div(o U6) U6 { return U6.cast(u / o) }

// Mod returns u modulo o.
func (u U6) Mod(o U6) U6 { return U6.cast(u % o) }

// And returns u&o reduced to 6 bits.
func (u U6) And(o U6) U6 { return U6.cast(u & o) }

// Or returns u|o reduced to 6 bits.
func (u U6) Or(o U6) U6 { return U6.cast(u | o) }

// Xor returns u^o reduced to 6 bits.
func (u U6) Xor(o U6) U6 { return U6.cast(u ^ o) }

// Shr returns u>>o reduced to 6 bits.
func (u U6) Shr(o U6) U6 { return U6.cast(u >> o) }

// Shl returns u<<o reduced to 6 bits.
func (u U6) Shl(o U6) U6 { return U6.cast(u << o) }

// Not returns bitwise complement of u reduced to 6 bits.
func (u U6) Not() U6 { return U6.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U6) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U6) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U6 value.
func (u U6) Signed() I6 { return I6.cast(I6(u)) }

// Clamp returns value saturated into the representable range of U6.
func (u U6) Clamp(value uint64) U6 { return uclamp[U6](value) }

// Clip masks u to a bits-wide low field.
func (u U6) Clip(bits int) U6 {
	b := 1 << (bits - 1)
	m := U6(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U6) Bit(index int) U6 {
	if index < 0 {
		index = 6 + index
	}
	return (u >> U6(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U6) SetBit(index int, v U6) {
	if index < 0 {
		index = 6 + index
	}
	bit := U6(1) << U6(index)
	*u = U6.cast((*u &^ bit) | ((v & 1) << U6(index)))
}

// Bitref returns a Range for bit index.
func (u *U6) Bitref(index int) Range[U6] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U6) Bits(lo, hi int) U6 {
	p := 6
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U6((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U6(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U6) SetBits(lo, hi int, v U6) {
	p := 6
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U6((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U6.cast((*u &^ mask) | ((v << U6(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U6) Bitsref(lo, hi int) Range[U6] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U6) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U6) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U6(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U6) Byteref(idx int) Range[U6] { return byte(u, idx) }

// I6 is an 6-bit signed integer in two's complement.
type I6 int8

// AsI6 returns a I6 representing v reduced to 6 bits.
func AsI6[T Signed](v T) I6 { return I6.cast(I6(v)) }

func (I6) nbits() int { return 6 }
func (i I6) mask() I6 { return 0x3f }
func (i I6) sign() I6 { return 1 << (i.nbits() - 1) }
func (i I6) cast() I6 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 6 bits to u.
func (i *I6) Set(v I6) I6 { *i = new(I6(v)).cast(); return I6.cast(v) }

// Add returns i+o reduced to 6 bits.
func (i I6) Add(o I6) I6 { return I6.cast(i + o) }

// Sub returns i-o reduced to 6 bits.
func (i I6) Sub(o I6) I6 { return I6.cast(i - o) }

// Inc returns i+1 reduced to 6 bits.
func (i I6) Inc() I6 { return I6.cast(i + 1) }

// Dec returns i-1 reduced to 6 bits.
func (i I6) Dec() I6 { return I6.cast(i - 1) }

// Mul returns i*o reduced to 6 bits.
func (i I6) Mul(o I6) I6 { return I6.cast(i * o) }

// Div returns i divided by o.
func (i I6) Div(o I6) I6 { return I6.cast(i / o) }

// Mod returns i modulo o.
func (i I6) Mod(o I6) I6 { return I6.cast(i % o) }

// And returns i&o reduced to 6 bits.
func (i I6) And(o I6) I6 { return I6.cast(i & o) }

// Or returns i|o reduced to 6 bits.
func (i I6) Or(o I6) I6 { return I6.cast(i | o) }

// Xor returns i^o reduced to 6 bits.
func (i I6) Xor(o I6) I6 { return I6.cast(i ^ o) }

// Shr returns i>>o reduced to 6 bits.
func (i I6) Shr(o I6) I6 { return I6.cast(i >> o) }

// Shl returns i<<o reduced to 6 bits.
func (i I6) Shl(o I6) I6 { return I6.cast(i << o) }

// Not returns bitwise complement of i reduced to 6 bits.
func (i I6) Not() I6 { return I6.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I6) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I6) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U6 bit pattern.
func (i I6) Unsigned() U6 { return U6.cast(U6(i)) }

// Clamp returns value saturated into the representable range of I6.
func (i I6) Clamp(value int64) I6 { return iclamp[I6](value) }

// Clip masks i to a bits-wide low field and sign-extends to 6 bits.
func (i I6) Clip(bits int) I6 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I6(uint64(i)&m^b) - I6(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I6) Bit(index int) I6 {
	if index < 0 {
		index = 6 + index
	}
	return (i >> I6(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I6) SetBit(index int, v I6) {
	if index < 0 {
		index = 6 + index
	}
	bit := I6(1) << I6(index)
	*i = I6.cast((*i &^ bit) | ((v & 1) << I6(index)))
}

// Bitref returns a Range for bit index.
func (i *I6) Bitref(index int) Range[I6] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I6) Bits(lo, hi int) I6 {
	p := 6
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I6((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I6(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I6) SetBits(lo, hi int, v I6) {
	p := 6
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I6((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I6.cast((*i &^ mask) | ((v << I6(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I6) Bitsref(lo, hi int) Range[I6] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I6) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I6) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I6(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I6) Byteref(idx int) Range[I6] { return byte(i, idx) }

// U7 is an 7-bit unsigned integer.
type U7 uint8

// AsU7 returns a U7 representing v reduced to 7 bits.
func AsU7[T Unsigned](v T) U7 { return U7.cast(U7(v)) }

func (U7) nbits() int { return 7 }
func (u U7) mask() U7 { return 0x7f }
func (u U7) cast() U7 { return u & u.mask() }

// Set assigns v reduced to 7 bits to u.
func (u *U7) Set(v U7) U7 { *u = new(U7(v)).cast(); return U7.cast(v) }

// Add returns u+o reduced to 7 bits.
func (u U7) Add(o U7) U7 { return U7.cast(u + o) }

// Sub returns u-o reduced to 7 bits.
func (u U7) Sub(o U7) U7 { return U7.cast(u - o) }

// Inc returns u+1 reduced to 7 bits.
func (u U7) Inc() U7 { return U7.cast(u + 1) }

// Dec returns u-1 reduced to 7 bits.
func (u U7) Dec() U7 { return U7.cast(u - 1) }

// Mul returns u*o reduced to 7 bits.
func (u U7) Mul(o U7) U7 { return U7.cast(u * o) }

// Div returns u divided by o.
func (u U7) Div(o U7) U7 { return U7.cast(u / o) }

// Mod returns u modulo o.
func (u U7) Mod(o U7) U7 { return U7.cast(u % o) }

// And returns u&o reduced to 7 bits.
func (u U7) And(o U7) U7 { return U7.cast(u & o) }

// Or returns u|o reduced to 7 bits.
func (u U7) Or(o U7) U7 { return U7.cast(u | o) }

// Xor returns u^o reduced to 7 bits.
func (u U7) Xor(o U7) U7 { return U7.cast(u ^ o) }

// Shr returns u>>o reduced to 7 bits.
func (u U7) Shr(o U7) U7 { return U7.cast(u >> o) }

// Shl returns u<<o reduced to 7 bits.
func (u U7) Shl(o U7) U7 { return U7.cast(u << o) }

// Not returns bitwise complement of u reduced to 7 bits.
func (u U7) Not() U7 { return U7.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U7) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U7) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U7 value.
func (u U7) Signed() I7 { return I7.cast(I7(u)) }

// Clamp returns value saturated into the representable range of U7.
func (u U7) Clamp(value uint64) U7 { return uclamp[U7](value) }

// Clip masks u to a bits-wide low field.
func (u U7) Clip(bits int) U7 {
	b := 1 << (bits - 1)
	m := U7(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U7) Bit(index int) U7 {
	if index < 0 {
		index = 7 + index
	}
	return (u >> U7(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U7) SetBit(index int, v U7) {
	if index < 0 {
		index = 7 + index
	}
	bit := U7(1) << U7(index)
	*u = U7.cast((*u &^ bit) | ((v & 1) << U7(index)))
}

// Bitref returns a Range for bit index.
func (u *U7) Bitref(index int) Range[U7] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U7) Bits(lo, hi int) U7 {
	p := 7
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U7((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U7(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U7) SetBits(lo, hi int, v U7) {
	p := 7
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U7((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U7.cast((*u &^ mask) | ((v << U7(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U7) Bitsref(lo, hi int) Range[U7] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U7) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U7) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U7(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U7) Byteref(idx int) Range[U7] { return byte(u, idx) }

// I7 is an 7-bit signed integer in two's complement.
type I7 int8

// AsI7 returns a I7 representing v reduced to 7 bits.
func AsI7[T Signed](v T) I7 { return I7.cast(I7(v)) }

func (I7) nbits() int { return 7 }
func (i I7) mask() I7 { return 0x7f }
func (i I7) sign() I7 { return 1 << (i.nbits() - 1) }
func (i I7) cast() I7 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 7 bits to u.
func (i *I7) Set(v I7) I7 { *i = new(I7(v)).cast(); return I7.cast(v) }

// Add returns i+o reduced to 7 bits.
func (i I7) Add(o I7) I7 { return I7.cast(i + o) }

// Sub returns i-o reduced to 7 bits.
func (i I7) Sub(o I7) I7 { return I7.cast(i - o) }

// Inc returns i+1 reduced to 7 bits.
func (i I7) Inc() I7 { return I7.cast(i + 1) }

// Dec returns i-1 reduced to 7 bits.
func (i I7) Dec() I7 { return I7.cast(i - 1) }

// Mul returns i*o reduced to 7 bits.
func (i I7) Mul(o I7) I7 { return I7.cast(i * o) }

// Div returns i divided by o.
func (i I7) Div(o I7) I7 { return I7.cast(i / o) }

// Mod returns i modulo o.
func (i I7) Mod(o I7) I7 { return I7.cast(i % o) }

// And returns i&o reduced to 7 bits.
func (i I7) And(o I7) I7 { return I7.cast(i & o) }

// Or returns i|o reduced to 7 bits.
func (i I7) Or(o I7) I7 { return I7.cast(i | o) }

// Xor returns i^o reduced to 7 bits.
func (i I7) Xor(o I7) I7 { return I7.cast(i ^ o) }

// Shr returns i>>o reduced to 7 bits.
func (i I7) Shr(o I7) I7 { return I7.cast(i >> o) }

// Shl returns i<<o reduced to 7 bits.
func (i I7) Shl(o I7) I7 { return I7.cast(i << o) }

// Not returns bitwise complement of i reduced to 7 bits.
func (i I7) Not() I7 { return I7.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I7) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I7) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U7 bit pattern.
func (i I7) Unsigned() U7 { return U7.cast(U7(i)) }

// Clamp returns value saturated into the representable range of I7.
func (i I7) Clamp(value int64) I7 { return iclamp[I7](value) }

// Clip masks i to a bits-wide low field and sign-extends to 7 bits.
func (i I7) Clip(bits int) I7 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I7(uint64(i)&m^b) - I7(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I7) Bit(index int) I7 {
	if index < 0 {
		index = 7 + index
	}
	return (i >> I7(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I7) SetBit(index int, v I7) {
	if index < 0 {
		index = 7 + index
	}
	bit := I7(1) << I7(index)
	*i = I7.cast((*i &^ bit) | ((v & 1) << I7(index)))
}

// Bitref returns a Range for bit index.
func (i *I7) Bitref(index int) Range[I7] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I7) Bits(lo, hi int) I7 {
	p := 7
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I7((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I7(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I7) SetBits(lo, hi int, v I7) {
	p := 7
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I7((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I7.cast((*i &^ mask) | ((v << I7(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I7) Bitsref(lo, hi int) Range[I7] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I7) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I7) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I7(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I7) Byteref(idx int) Range[I7] { return byte(i, idx) }

// U8 is an 8-bit unsigned integer.
type U8 uint8

// AsU8 returns a U8 representing v reduced to 8 bits.
func AsU8[T Unsigned](v T) U8 { return U8.cast(U8(v)) }

func (U8) nbits() int { return 8 }
func (u U8) mask() U8 { return 0xff }
func (u U8) cast() U8 { return u & u.mask() }

// Set assigns v reduced to 8 bits to u.
func (u *U8) Set(v U8) U8 { *u = new(U8(v)).cast(); return U8.cast(v) }

// Add returns u+o reduced to 8 bits.
func (u U8) Add(o U8) U8 { return U8.cast(u + o) }

// Sub returns u-o reduced to 8 bits.
func (u U8) Sub(o U8) U8 { return U8.cast(u - o) }

// Inc returns u+1 reduced to 8 bits.
func (u U8) Inc() U8 { return U8.cast(u + 1) }

// Dec returns u-1 reduced to 8 bits.
func (u U8) Dec() U8 { return U8.cast(u - 1) }

// Mul returns u*o reduced to 8 bits.
func (u U8) Mul(o U8) U8 { return U8.cast(u * o) }

// Div returns u divided by o.
func (u U8) Div(o U8) U8 { return U8.cast(u / o) }

// Mod returns u modulo o.
func (u U8) Mod(o U8) U8 { return U8.cast(u % o) }

// And returns u&o reduced to 8 bits.
func (u U8) And(o U8) U8 { return U8.cast(u & o) }

// Or returns u|o reduced to 8 bits.
func (u U8) Or(o U8) U8 { return U8.cast(u | o) }

// Xor returns u^o reduced to 8 bits.
func (u U8) Xor(o U8) U8 { return U8.cast(u ^ o) }

// Shr returns u>>o reduced to 8 bits.
func (u U8) Shr(o U8) U8 { return U8.cast(u >> o) }

// Shl returns u<<o reduced to 8 bits.
func (u U8) Shl(o U8) U8 { return U8.cast(u << o) }

// Not returns bitwise complement of u reduced to 8 bits.
func (u U8) Not() U8 { return U8.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U8) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U8) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U8 value.
func (u U8) Signed() I8 { return I8.cast(I8(u)) }

// Clamp returns value saturated into the representable range of U8.
func (u U8) Clamp(value uint64) U8 { return uclamp[U8](value) }

// Clip masks u to a bits-wide low field.
func (u U8) Clip(bits int) U8 {
	b := 1 << (bits - 1)
	m := U8(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U8) Bit(index int) U8 {
	if index < 0 {
		index = 8 + index
	}
	return (u >> U8(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U8) SetBit(index int, v U8) {
	if index < 0 {
		index = 8 + index
	}
	bit := U8(1) << U8(index)
	*u = U8.cast((*u &^ bit) | ((v & 1) << U8(index)))
}

// Bitref returns a Range for bit index.
func (u *U8) Bitref(index int) Range[U8] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U8) Bits(lo, hi int) U8 {
	p := 8
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U8((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U8(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U8) SetBits(lo, hi int, v U8) {
	p := 8
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U8((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U8.cast((*u &^ mask) | ((v << U8(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U8) Bitsref(lo, hi int) Range[U8] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U8) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U8) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U8(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U8) Byteref(idx int) Range[U8] { return byte(u, idx) }

// I8 is an 8-bit signed integer in two's complement.
type I8 int8

// AsI8 returns a I8 representing v reduced to 8 bits.
func AsI8[T Signed](v T) I8 { return I8.cast(I8(v)) }

func (I8) nbits() int { return 8 }
func (i I8) mask() I8 { return -1 }
func (i I8) sign() I8 { return 1 << (i.nbits() - 1) }
func (i I8) cast() I8 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 8 bits to u.
func (i *I8) Set(v I8) I8 { *i = new(I8(v)).cast(); return I8.cast(v) }

// Add returns i+o reduced to 8 bits.
func (i I8) Add(o I8) I8 { return I8.cast(i + o) }

// Sub returns i-o reduced to 8 bits.
func (i I8) Sub(o I8) I8 { return I8.cast(i - o) }

// Inc returns i+1 reduced to 8 bits.
func (i I8) Inc() I8 { return I8.cast(i + 1) }

// Dec returns i-1 reduced to 8 bits.
func (i I8) Dec() I8 { return I8.cast(i - 1) }

// Mul returns i*o reduced to 8 bits.
func (i I8) Mul(o I8) I8 { return I8.cast(i * o) }

// Div returns i divided by o.
func (i I8) Div(o I8) I8 { return I8.cast(i / o) }

// Mod returns i modulo o.
func (i I8) Mod(o I8) I8 { return I8.cast(i % o) }

// And returns i&o reduced to 8 bits.
func (i I8) And(o I8) I8 { return I8.cast(i & o) }

// Or returns i|o reduced to 8 bits.
func (i I8) Or(o I8) I8 { return I8.cast(i | o) }

// Xor returns i^o reduced to 8 bits.
func (i I8) Xor(o I8) I8 { return I8.cast(i ^ o) }

// Shr returns i>>o reduced to 8 bits.
func (i I8) Shr(o I8) I8 { return I8.cast(i >> o) }

// Shl returns i<<o reduced to 8 bits.
func (i I8) Shl(o I8) I8 { return I8.cast(i << o) }

// Not returns bitwise complement of i reduced to 8 bits.
func (i I8) Not() I8 { return I8.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I8) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I8) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U8 bit pattern.
func (i I8) Unsigned() U8 { return U8.cast(U8(i)) }

// Clamp returns value saturated into the representable range of I8.
func (i I8) Clamp(value int64) I8 { return iclamp[I8](value) }

// Clip masks i to a bits-wide low field and sign-extends to 8 bits.
func (i I8) Clip(bits int) I8 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I8(uint64(i)&m^b) - I8(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I8) Bit(index int) I8 {
	if index < 0 {
		index = 8 + index
	}
	return (i >> I8(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I8) SetBit(index int, v I8) {
	if index < 0 {
		index = 8 + index
	}
	bit := I8(1) << I8(index)
	*i = I8.cast((*i &^ bit) | ((v & 1) << I8(index)))
}

// Bitref returns a Range for bit index.
func (i *I8) Bitref(index int) Range[I8] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I8) Bits(lo, hi int) I8 {
	p := 8
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I8((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I8(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I8) SetBits(lo, hi int, v I8) {
	p := 8
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I8((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I8.cast((*i &^ mask) | ((v << I8(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I8) Bitsref(lo, hi int) Range[I8] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I8) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I8) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I8(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I8) Byteref(idx int) Range[I8] { return byte(i, idx) }

// U9 is an 9-bit unsigned integer.
type U9 uint16

// AsU9 returns a U9 representing v reduced to 9 bits.
func AsU9[T Unsigned](v T) U9 { return U9.cast(U9(v)) }

func (U9) nbits() int { return 9 }
func (u U9) mask() U9 { return 0x1ff }
func (u U9) cast() U9 { return u & u.mask() }

// Set assigns v reduced to 9 bits to u.
func (u *U9) Set(v U9) U9 { *u = new(U9(v)).cast(); return U9.cast(v) }

// Add returns u+o reduced to 9 bits.
func (u U9) Add(o U9) U9 { return U9.cast(u + o) }

// Sub returns u-o reduced to 9 bits.
func (u U9) Sub(o U9) U9 { return U9.cast(u - o) }

// Inc returns u+1 reduced to 9 bits.
func (u U9) Inc() U9 { return U9.cast(u + 1) }

// Dec returns u-1 reduced to 9 bits.
func (u U9) Dec() U9 { return U9.cast(u - 1) }

// Mul returns u*o reduced to 9 bits.
func (u U9) Mul(o U9) U9 { return U9.cast(u * o) }

// Div returns u divided by o.
func (u U9) Div(o U9) U9 { return U9.cast(u / o) }

// Mod returns u modulo o.
func (u U9) Mod(o U9) U9 { return U9.cast(u % o) }

// And returns u&o reduced to 9 bits.
func (u U9) And(o U9) U9 { return U9.cast(u & o) }

// Or returns u|o reduced to 9 bits.
func (u U9) Or(o U9) U9 { return U9.cast(u | o) }

// Xor returns u^o reduced to 9 bits.
func (u U9) Xor(o U9) U9 { return U9.cast(u ^ o) }

// Shr returns u>>o reduced to 9 bits.
func (u U9) Shr(o U9) U9 { return U9.cast(u >> o) }

// Shl returns u<<o reduced to 9 bits.
func (u U9) Shl(o U9) U9 { return U9.cast(u << o) }

// Not returns bitwise complement of u reduced to 9 bits.
func (u U9) Not() U9 { return U9.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U9) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U9) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U9 value.
func (u U9) Signed() I9 { return I9.cast(I9(u)) }

// Clamp returns value saturated into the representable range of U9.
func (u U9) Clamp(value uint64) U9 { return uclamp[U9](value) }

// Clip masks u to a bits-wide low field.
func (u U9) Clip(bits int) U9 {
	b := 1 << (bits - 1)
	m := U9(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U9) Bit(index int) U9 {
	if index < 0 {
		index = 9 + index
	}
	return (u >> U9(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U9) SetBit(index int, v U9) {
	if index < 0 {
		index = 9 + index
	}
	bit := U9(1) << U9(index)
	*u = U9.cast((*u &^ bit) | ((v & 1) << U9(index)))
}

// Bitref returns a Range for bit index.
func (u *U9) Bitref(index int) Range[U9] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U9) Bits(lo, hi int) U9 {
	p := 9
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U9((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U9(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U9) SetBits(lo, hi int, v U9) {
	p := 9
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U9((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U9.cast((*u &^ mask) | ((v << U9(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U9) Bitsref(lo, hi int) Range[U9] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U9) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U9) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U9(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U9) Byteref(idx int) Range[U9] { return byte(u, idx) }

// I9 is an 9-bit signed integer in two's complement.
type I9 int16

// AsI9 returns a I9 representing v reduced to 9 bits.
func AsI9[T Signed](v T) I9 { return I9.cast(I9(v)) }

func (I9) nbits() int { return 9 }
func (i I9) mask() I9 { return 0x1ff }
func (i I9) sign() I9 { return 1 << (i.nbits() - 1) }
func (i I9) cast() I9 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 9 bits to u.
func (i *I9) Set(v I9) I9 { *i = new(I9(v)).cast(); return I9.cast(v) }

// Add returns i+o reduced to 9 bits.
func (i I9) Add(o I9) I9 { return I9.cast(i + o) }

// Sub returns i-o reduced to 9 bits.
func (i I9) Sub(o I9) I9 { return I9.cast(i - o) }

// Inc returns i+1 reduced to 9 bits.
func (i I9) Inc() I9 { return I9.cast(i + 1) }

// Dec returns i-1 reduced to 9 bits.
func (i I9) Dec() I9 { return I9.cast(i - 1) }

// Mul returns i*o reduced to 9 bits.
func (i I9) Mul(o I9) I9 { return I9.cast(i * o) }

// Div returns i divided by o.
func (i I9) Div(o I9) I9 { return I9.cast(i / o) }

// Mod returns i modulo o.
func (i I9) Mod(o I9) I9 { return I9.cast(i % o) }

// And returns i&o reduced to 9 bits.
func (i I9) And(o I9) I9 { return I9.cast(i & o) }

// Or returns i|o reduced to 9 bits.
func (i I9) Or(o I9) I9 { return I9.cast(i | o) }

// Xor returns i^o reduced to 9 bits.
func (i I9) Xor(o I9) I9 { return I9.cast(i ^ o) }

// Shr returns i>>o reduced to 9 bits.
func (i I9) Shr(o I9) I9 { return I9.cast(i >> o) }

// Shl returns i<<o reduced to 9 bits.
func (i I9) Shl(o I9) I9 { return I9.cast(i << o) }

// Not returns bitwise complement of i reduced to 9 bits.
func (i I9) Not() I9 { return I9.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I9) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I9) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U9 bit pattern.
func (i I9) Unsigned() U9 { return U9.cast(U9(i)) }

// Clamp returns value saturated into the representable range of I9.
func (i I9) Clamp(value int64) I9 { return iclamp[I9](value) }

// Clip masks i to a bits-wide low field and sign-extends to 9 bits.
func (i I9) Clip(bits int) I9 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I9(uint64(i)&m^b) - I9(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I9) Bit(index int) I9 {
	if index < 0 {
		index = 9 + index
	}
	return (i >> I9(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I9) SetBit(index int, v I9) {
	if index < 0 {
		index = 9 + index
	}
	bit := I9(1) << I9(index)
	*i = I9.cast((*i &^ bit) | ((v & 1) << I9(index)))
}

// Bitref returns a Range for bit index.
func (i *I9) Bitref(index int) Range[I9] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I9) Bits(lo, hi int) I9 {
	p := 9
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I9((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I9(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I9) SetBits(lo, hi int, v I9) {
	p := 9
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I9((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I9.cast((*i &^ mask) | ((v << I9(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I9) Bitsref(lo, hi int) Range[I9] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I9) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I9) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I9(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I9) Byteref(idx int) Range[I9] { return byte(i, idx) }

// U10 is an 10-bit unsigned integer.
type U10 uint16

// AsU10 returns a U10 representing v reduced to 10 bits.
func AsU10[T Unsigned](v T) U10 { return U10.cast(U10(v)) }

func (U10) nbits() int  { return 10 }
func (u U10) mask() U10 { return 0x3ff }
func (u U10) cast() U10 { return u & u.mask() }

// Set assigns v reduced to 10 bits to u.
func (u *U10) Set(v U10) U10 { *u = new(U10(v)).cast(); return U10.cast(v) }

// Add returns u+o reduced to 10 bits.
func (u U10) Add(o U10) U10 { return U10.cast(u + o) }

// Sub returns u-o reduced to 10 bits.
func (u U10) Sub(o U10) U10 { return U10.cast(u - o) }

// Inc returns u+1 reduced to 10 bits.
func (u U10) Inc() U10 { return U10.cast(u + 1) }

// Dec returns u-1 reduced to 10 bits.
func (u U10) Dec() U10 { return U10.cast(u - 1) }

// Mul returns u*o reduced to 10 bits.
func (u U10) Mul(o U10) U10 { return U10.cast(u * o) }

// Div returns u divided by o.
func (u U10) Div(o U10) U10 { return U10.cast(u / o) }

// Mod returns u modulo o.
func (u U10) Mod(o U10) U10 { return U10.cast(u % o) }

// And returns u&o reduced to 10 bits.
func (u U10) And(o U10) U10 { return U10.cast(u & o) }

// Or returns u|o reduced to 10 bits.
func (u U10) Or(o U10) U10 { return U10.cast(u | o) }

// Xor returns u^o reduced to 10 bits.
func (u U10) Xor(o U10) U10 { return U10.cast(u ^ o) }

// Shr returns u>>o reduced to 10 bits.
func (u U10) Shr(o U10) U10 { return U10.cast(u >> o) }

// Shl returns u<<o reduced to 10 bits.
func (u U10) Shl(o U10) U10 { return U10.cast(u << o) }

// Not returns bitwise complement of u reduced to 10 bits.
func (u U10) Not() U10 { return U10.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U10) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U10) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U10 value.
func (u U10) Signed() I10 { return I10.cast(I10(u)) }

// Clamp returns value saturated into the representable range of U10.
func (u U10) Clamp(value uint64) U10 { return uclamp[U10](value) }

// Clip masks u to a bits-wide low field.
func (u U10) Clip(bits int) U10 {
	b := 1 << (bits - 1)
	m := U10(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U10) Bit(index int) U10 {
	if index < 0 {
		index = 10 + index
	}
	return (u >> U10(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U10) SetBit(index int, v U10) {
	if index < 0 {
		index = 10 + index
	}
	bit := U10(1) << U10(index)
	*u = U10.cast((*u &^ bit) | ((v & 1) << U10(index)))
}

// Bitref returns a Range for bit index.
func (u *U10) Bitref(index int) Range[U10] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U10) Bits(lo, hi int) U10 {
	p := 10
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U10((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U10(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U10) SetBits(lo, hi int, v U10) {
	p := 10
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U10((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U10.cast((*u &^ mask) | ((v << U10(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U10) Bitsref(lo, hi int) Range[U10] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U10) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U10) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U10(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U10) Byteref(idx int) Range[U10] { return byte(u, idx) }

// I10 is an 10-bit signed integer in two's complement.
type I10 int16

// AsI10 returns a I10 representing v reduced to 10 bits.
func AsI10[T Signed](v T) I10 { return I10.cast(I10(v)) }

func (I10) nbits() int  { return 10 }
func (i I10) mask() I10 { return 0x3ff }
func (i I10) sign() I10 { return 1 << (i.nbits() - 1) }
func (i I10) cast() I10 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 10 bits to u.
func (i *I10) Set(v I10) I10 { *i = new(I10(v)).cast(); return I10.cast(v) }

// Add returns i+o reduced to 10 bits.
func (i I10) Add(o I10) I10 { return I10.cast(i + o) }

// Sub returns i-o reduced to 10 bits.
func (i I10) Sub(o I10) I10 { return I10.cast(i - o) }

// Inc returns i+1 reduced to 10 bits.
func (i I10) Inc() I10 { return I10.cast(i + 1) }

// Dec returns i-1 reduced to 10 bits.
func (i I10) Dec() I10 { return I10.cast(i - 1) }

// Mul returns i*o reduced to 10 bits.
func (i I10) Mul(o I10) I10 { return I10.cast(i * o) }

// Div returns i divided by o.
func (i I10) Div(o I10) I10 { return I10.cast(i / o) }

// Mod returns i modulo o.
func (i I10) Mod(o I10) I10 { return I10.cast(i % o) }

// And returns i&o reduced to 10 bits.
func (i I10) And(o I10) I10 { return I10.cast(i & o) }

// Or returns i|o reduced to 10 bits.
func (i I10) Or(o I10) I10 { return I10.cast(i | o) }

// Xor returns i^o reduced to 10 bits.
func (i I10) Xor(o I10) I10 { return I10.cast(i ^ o) }

// Shr returns i>>o reduced to 10 bits.
func (i I10) Shr(o I10) I10 { return I10.cast(i >> o) }

// Shl returns i<<o reduced to 10 bits.
func (i I10) Shl(o I10) I10 { return I10.cast(i << o) }

// Not returns bitwise complement of i reduced to 10 bits.
func (i I10) Not() I10 { return I10.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I10) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I10) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U10 bit pattern.
func (i I10) Unsigned() U10 { return U10.cast(U10(i)) }

// Clamp returns value saturated into the representable range of I10.
func (i I10) Clamp(value int64) I10 { return iclamp[I10](value) }

// Clip masks i to a bits-wide low field and sign-extends to 10 bits.
func (i I10) Clip(bits int) I10 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I10(uint64(i)&m^b) - I10(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I10) Bit(index int) I10 {
	if index < 0 {
		index = 10 + index
	}
	return (i >> I10(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I10) SetBit(index int, v I10) {
	if index < 0 {
		index = 10 + index
	}
	bit := I10(1) << I10(index)
	*i = I10.cast((*i &^ bit) | ((v & 1) << I10(index)))
}

// Bitref returns a Range for bit index.
func (i *I10) Bitref(index int) Range[I10] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I10) Bits(lo, hi int) I10 {
	p := 10
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I10((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I10(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I10) SetBits(lo, hi int, v I10) {
	p := 10
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I10((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I10.cast((*i &^ mask) | ((v << I10(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I10) Bitsref(lo, hi int) Range[I10] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I10) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I10) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I10(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I10) Byteref(idx int) Range[I10] { return byte(i, idx) }

// U11 is an 11-bit unsigned integer.
type U11 uint16

// AsU11 returns a U11 representing v reduced to 11 bits.
func AsU11[T Unsigned](v T) U11 { return U11.cast(U11(v)) }

func (U11) nbits() int  { return 11 }
func (u U11) mask() U11 { return 0x7ff }
func (u U11) cast() U11 { return u & u.mask() }

// Set assigns v reduced to 11 bits to u.
func (u *U11) Set(v U11) U11 { *u = new(U11(v)).cast(); return U11.cast(v) }

// Add returns u+o reduced to 11 bits.
func (u U11) Add(o U11) U11 { return U11.cast(u + o) }

// Sub returns u-o reduced to 11 bits.
func (u U11) Sub(o U11) U11 { return U11.cast(u - o) }

// Inc returns u+1 reduced to 11 bits.
func (u U11) Inc() U11 { return U11.cast(u + 1) }

// Dec returns u-1 reduced to 11 bits.
func (u U11) Dec() U11 { return U11.cast(u - 1) }

// Mul returns u*o reduced to 11 bits.
func (u U11) Mul(o U11) U11 { return U11.cast(u * o) }

// Div returns u divided by o.
func (u U11) Div(o U11) U11 { return U11.cast(u / o) }

// Mod returns u modulo o.
func (u U11) Mod(o U11) U11 { return U11.cast(u % o) }

// And returns u&o reduced to 11 bits.
func (u U11) And(o U11) U11 { return U11.cast(u & o) }

// Or returns u|o reduced to 11 bits.
func (u U11) Or(o U11) U11 { return U11.cast(u | o) }

// Xor returns u^o reduced to 11 bits.
func (u U11) Xor(o U11) U11 { return U11.cast(u ^ o) }

// Shr returns u>>o reduced to 11 bits.
func (u U11) Shr(o U11) U11 { return U11.cast(u >> o) }

// Shl returns u<<o reduced to 11 bits.
func (u U11) Shl(o U11) U11 { return U11.cast(u << o) }

// Not returns bitwise complement of u reduced to 11 bits.
func (u U11) Not() U11 { return U11.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U11) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U11) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U11 value.
func (u U11) Signed() I11 { return I11.cast(I11(u)) }

// Clamp returns value saturated into the representable range of U11.
func (u U11) Clamp(value uint64) U11 { return uclamp[U11](value) }

// Clip masks u to a bits-wide low field.
func (u U11) Clip(bits int) U11 {
	b := 1 << (bits - 1)
	m := U11(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U11) Bit(index int) U11 {
	if index < 0 {
		index = 11 + index
	}
	return (u >> U11(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U11) SetBit(index int, v U11) {
	if index < 0 {
		index = 11 + index
	}
	bit := U11(1) << U11(index)
	*u = U11.cast((*u &^ bit) | ((v & 1) << U11(index)))
}

// Bitref returns a Range for bit index.
func (u *U11) Bitref(index int) Range[U11] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U11) Bits(lo, hi int) U11 {
	p := 11
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U11((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U11(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U11) SetBits(lo, hi int, v U11) {
	p := 11
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U11((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U11.cast((*u &^ mask) | ((v << U11(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U11) Bitsref(lo, hi int) Range[U11] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U11) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U11) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U11(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U11) Byteref(idx int) Range[U11] { return byte(u, idx) }

// I11 is an 11-bit signed integer in two's complement.
type I11 int16

// AsI11 returns a I11 representing v reduced to 11 bits.
func AsI11[T Signed](v T) I11 { return I11.cast(I11(v)) }

func (I11) nbits() int  { return 11 }
func (i I11) mask() I11 { return 0x7ff }
func (i I11) sign() I11 { return 1 << (i.nbits() - 1) }
func (i I11) cast() I11 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 11 bits to u.
func (i *I11) Set(v I11) I11 { *i = new(I11(v)).cast(); return I11.cast(v) }

// Add returns i+o reduced to 11 bits.
func (i I11) Add(o I11) I11 { return I11.cast(i + o) }

// Sub returns i-o reduced to 11 bits.
func (i I11) Sub(o I11) I11 { return I11.cast(i - o) }

// Inc returns i+1 reduced to 11 bits.
func (i I11) Inc() I11 { return I11.cast(i + 1) }

// Dec returns i-1 reduced to 11 bits.
func (i I11) Dec() I11 { return I11.cast(i - 1) }

// Mul returns i*o reduced to 11 bits.
func (i I11) Mul(o I11) I11 { return I11.cast(i * o) }

// Div returns i divided by o.
func (i I11) Div(o I11) I11 { return I11.cast(i / o) }

// Mod returns i modulo o.
func (i I11) Mod(o I11) I11 { return I11.cast(i % o) }

// And returns i&o reduced to 11 bits.
func (i I11) And(o I11) I11 { return I11.cast(i & o) }

// Or returns i|o reduced to 11 bits.
func (i I11) Or(o I11) I11 { return I11.cast(i | o) }

// Xor returns i^o reduced to 11 bits.
func (i I11) Xor(o I11) I11 { return I11.cast(i ^ o) }

// Shr returns i>>o reduced to 11 bits.
func (i I11) Shr(o I11) I11 { return I11.cast(i >> o) }

// Shl returns i<<o reduced to 11 bits.
func (i I11) Shl(o I11) I11 { return I11.cast(i << o) }

// Not returns bitwise complement of i reduced to 11 bits.
func (i I11) Not() I11 { return I11.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I11) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I11) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U11 bit pattern.
func (i I11) Unsigned() U11 { return U11.cast(U11(i)) }

// Clamp returns value saturated into the representable range of I11.
func (i I11) Clamp(value int64) I11 { return iclamp[I11](value) }

// Clip masks i to a bits-wide low field and sign-extends to 11 bits.
func (i I11) Clip(bits int) I11 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I11(uint64(i)&m^b) - I11(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I11) Bit(index int) I11 {
	if index < 0 {
		index = 11 + index
	}
	return (i >> I11(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I11) SetBit(index int, v I11) {
	if index < 0 {
		index = 11 + index
	}
	bit := I11(1) << I11(index)
	*i = I11.cast((*i &^ bit) | ((v & 1) << I11(index)))
}

// Bitref returns a Range for bit index.
func (i *I11) Bitref(index int) Range[I11] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I11) Bits(lo, hi int) I11 {
	p := 11
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I11((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I11(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I11) SetBits(lo, hi int, v I11) {
	p := 11
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I11((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I11.cast((*i &^ mask) | ((v << I11(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I11) Bitsref(lo, hi int) Range[I11] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I11) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I11) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I11(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I11) Byteref(idx int) Range[I11] { return byte(i, idx) }

// U12 is an 12-bit unsigned integer.
type U12 uint16

// AsU12 returns a U12 representing v reduced to 12 bits.
func AsU12[T Unsigned](v T) U12 { return U12.cast(U12(v)) }

func (U12) nbits() int  { return 12 }
func (u U12) mask() U12 { return 0xfff }
func (u U12) cast() U12 { return u & u.mask() }

// Set assigns v reduced to 12 bits to u.
func (u *U12) Set(v U12) U12 { *u = new(U12(v)).cast(); return U12.cast(v) }

// Add returns u+o reduced to 12 bits.
func (u U12) Add(o U12) U12 { return U12.cast(u + o) }

// Sub returns u-o reduced to 12 bits.
func (u U12) Sub(o U12) U12 { return U12.cast(u - o) }

// Inc returns u+1 reduced to 12 bits.
func (u U12) Inc() U12 { return U12.cast(u + 1) }

// Dec returns u-1 reduced to 12 bits.
func (u U12) Dec() U12 { return U12.cast(u - 1) }

// Mul returns u*o reduced to 12 bits.
func (u U12) Mul(o U12) U12 { return U12.cast(u * o) }

// Div returns u divided by o.
func (u U12) Div(o U12) U12 { return U12.cast(u / o) }

// Mod returns u modulo o.
func (u U12) Mod(o U12) U12 { return U12.cast(u % o) }

// And returns u&o reduced to 12 bits.
func (u U12) And(o U12) U12 { return U12.cast(u & o) }

// Or returns u|o reduced to 12 bits.
func (u U12) Or(o U12) U12 { return U12.cast(u | o) }

// Xor returns u^o reduced to 12 bits.
func (u U12) Xor(o U12) U12 { return U12.cast(u ^ o) }

// Shr returns u>>o reduced to 12 bits.
func (u U12) Shr(o U12) U12 { return U12.cast(u >> o) }

// Shl returns u<<o reduced to 12 bits.
func (u U12) Shl(o U12) U12 { return U12.cast(u << o) }

// Not returns bitwise complement of u reduced to 12 bits.
func (u U12) Not() U12 { return U12.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U12) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U12) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U12 value.
func (u U12) Signed() I12 { return I12.cast(I12(u)) }

// Clamp returns value saturated into the representable range of U12.
func (u U12) Clamp(value uint64) U12 { return uclamp[U12](value) }

// Clip masks u to a bits-wide low field.
func (u U12) Clip(bits int) U12 {
	b := 1 << (bits - 1)
	m := U12(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U12) Bit(index int) U12 {
	if index < 0 {
		index = 12 + index
	}
	return (u >> U12(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U12) SetBit(index int, v U12) {
	if index < 0 {
		index = 12 + index
	}
	bit := U12(1) << U12(index)
	*u = U12.cast((*u &^ bit) | ((v & 1) << U12(index)))
}

// Bitref returns a Range for bit index.
func (u *U12) Bitref(index int) Range[U12] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U12) Bits(lo, hi int) U12 {
	p := 12
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U12((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U12(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U12) SetBits(lo, hi int, v U12) {
	p := 12
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U12((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U12.cast((*u &^ mask) | ((v << U12(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U12) Bitsref(lo, hi int) Range[U12] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U12) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U12) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U12(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U12) Byteref(idx int) Range[U12] { return byte(u, idx) }

// I12 is an 12-bit signed integer in two's complement.
type I12 int16

// AsI12 returns a I12 representing v reduced to 12 bits.
func AsI12[T Signed](v T) I12 { return I12.cast(I12(v)) }

func (I12) nbits() int  { return 12 }
func (i I12) mask() I12 { return 0xfff }
func (i I12) sign() I12 { return 1 << (i.nbits() - 1) }
func (i I12) cast() I12 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 12 bits to u.
func (i *I12) Set(v I12) I12 { *i = new(I12(v)).cast(); return I12.cast(v) }

// Add returns i+o reduced to 12 bits.
func (i I12) Add(o I12) I12 { return I12.cast(i + o) }

// Sub returns i-o reduced to 12 bits.
func (i I12) Sub(o I12) I12 { return I12.cast(i - o) }

// Inc returns i+1 reduced to 12 bits.
func (i I12) Inc() I12 { return I12.cast(i + 1) }

// Dec returns i-1 reduced to 12 bits.
func (i I12) Dec() I12 { return I12.cast(i - 1) }

// Mul returns i*o reduced to 12 bits.
func (i I12) Mul(o I12) I12 { return I12.cast(i * o) }

// Div returns i divided by o.
func (i I12) Div(o I12) I12 { return I12.cast(i / o) }

// Mod returns i modulo o.
func (i I12) Mod(o I12) I12 { return I12.cast(i % o) }

// And returns i&o reduced to 12 bits.
func (i I12) And(o I12) I12 { return I12.cast(i & o) }

// Or returns i|o reduced to 12 bits.
func (i I12) Or(o I12) I12 { return I12.cast(i | o) }

// Xor returns i^o reduced to 12 bits.
func (i I12) Xor(o I12) I12 { return I12.cast(i ^ o) }

// Shr returns i>>o reduced to 12 bits.
func (i I12) Shr(o I12) I12 { return I12.cast(i >> o) }

// Shl returns i<<o reduced to 12 bits.
func (i I12) Shl(o I12) I12 { return I12.cast(i << o) }

// Not returns bitwise complement of i reduced to 12 bits.
func (i I12) Not() I12 { return I12.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I12) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I12) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U12 bit pattern.
func (i I12) Unsigned() U12 { return U12.cast(U12(i)) }

// Clamp returns value saturated into the representable range of I12.
func (i I12) Clamp(value int64) I12 { return iclamp[I12](value) }

// Clip masks i to a bits-wide low field and sign-extends to 12 bits.
func (i I12) Clip(bits int) I12 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I12(uint64(i)&m^b) - I12(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I12) Bit(index int) I12 {
	if index < 0 {
		index = 12 + index
	}
	return (i >> I12(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I12) SetBit(index int, v I12) {
	if index < 0 {
		index = 12 + index
	}
	bit := I12(1) << I12(index)
	*i = I12.cast((*i &^ bit) | ((v & 1) << I12(index)))
}

// Bitref returns a Range for bit index.
func (i *I12) Bitref(index int) Range[I12] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I12) Bits(lo, hi int) I12 {
	p := 12
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I12((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I12(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I12) SetBits(lo, hi int, v I12) {
	p := 12
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I12((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I12.cast((*i &^ mask) | ((v << I12(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I12) Bitsref(lo, hi int) Range[I12] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I12) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I12) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I12(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I12) Byteref(idx int) Range[I12] { return byte(i, idx) }

// U13 is an 13-bit unsigned integer.
type U13 uint16

// AsU13 returns a U13 representing v reduced to 13 bits.
func AsU13[T Unsigned](v T) U13 { return U13.cast(U13(v)) }

func (U13) nbits() int  { return 13 }
func (u U13) mask() U13 { return 0x1fff }
func (u U13) cast() U13 { return u & u.mask() }

// Set assigns v reduced to 13 bits to u.
func (u *U13) Set(v U13) U13 { *u = new(U13(v)).cast(); return U13.cast(v) }

// Add returns u+o reduced to 13 bits.
func (u U13) Add(o U13) U13 { return U13.cast(u + o) }

// Sub returns u-o reduced to 13 bits.
func (u U13) Sub(o U13) U13 { return U13.cast(u - o) }

// Inc returns u+1 reduced to 13 bits.
func (u U13) Inc() U13 { return U13.cast(u + 1) }

// Dec returns u-1 reduced to 13 bits.
func (u U13) Dec() U13 { return U13.cast(u - 1) }

// Mul returns u*o reduced to 13 bits.
func (u U13) Mul(o U13) U13 { return U13.cast(u * o) }

// Div returns u divided by o.
func (u U13) Div(o U13) U13 { return U13.cast(u / o) }

// Mod returns u modulo o.
func (u U13) Mod(o U13) U13 { return U13.cast(u % o) }

// And returns u&o reduced to 13 bits.
func (u U13) And(o U13) U13 { return U13.cast(u & o) }

// Or returns u|o reduced to 13 bits.
func (u U13) Or(o U13) U13 { return U13.cast(u | o) }

// Xor returns u^o reduced to 13 bits.
func (u U13) Xor(o U13) U13 { return U13.cast(u ^ o) }

// Shr returns u>>o reduced to 13 bits.
func (u U13) Shr(o U13) U13 { return U13.cast(u >> o) }

// Shl returns u<<o reduced to 13 bits.
func (u U13) Shl(o U13) U13 { return U13.cast(u << o) }

// Not returns bitwise complement of u reduced to 13 bits.
func (u U13) Not() U13 { return U13.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U13) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U13) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U13 value.
func (u U13) Signed() I13 { return I13.cast(I13(u)) }

// Clamp returns value saturated into the representable range of U13.
func (u U13) Clamp(value uint64) U13 { return uclamp[U13](value) }

// Clip masks u to a bits-wide low field.
func (u U13) Clip(bits int) U13 {
	b := 1 << (bits - 1)
	m := U13(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U13) Bit(index int) U13 {
	if index < 0 {
		index = 13 + index
	}
	return (u >> U13(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U13) SetBit(index int, v U13) {
	if index < 0 {
		index = 13 + index
	}
	bit := U13(1) << U13(index)
	*u = U13.cast((*u &^ bit) | ((v & 1) << U13(index)))
}

// Bitref returns a Range for bit index.
func (u *U13) Bitref(index int) Range[U13] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U13) Bits(lo, hi int) U13 {
	p := 13
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U13((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U13(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U13) SetBits(lo, hi int, v U13) {
	p := 13
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U13((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U13.cast((*u &^ mask) | ((v << U13(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U13) Bitsref(lo, hi int) Range[U13] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U13) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U13) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U13(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U13) Byteref(idx int) Range[U13] { return byte(u, idx) }

// I13 is an 13-bit signed integer in two's complement.
type I13 int16

// AsI13 returns a I13 representing v reduced to 13 bits.
func AsI13[T Signed](v T) I13 { return I13.cast(I13(v)) }

func (I13) nbits() int  { return 13 }
func (i I13) mask() I13 { return 0x1fff }
func (i I13) sign() I13 { return 1 << (i.nbits() - 1) }
func (i I13) cast() I13 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 13 bits to u.
func (i *I13) Set(v I13) I13 { *i = new(I13(v)).cast(); return I13.cast(v) }

// Add returns i+o reduced to 13 bits.
func (i I13) Add(o I13) I13 { return I13.cast(i + o) }

// Sub returns i-o reduced to 13 bits.
func (i I13) Sub(o I13) I13 { return I13.cast(i - o) }

// Inc returns i+1 reduced to 13 bits.
func (i I13) Inc() I13 { return I13.cast(i + 1) }

// Dec returns i-1 reduced to 13 bits.
func (i I13) Dec() I13 { return I13.cast(i - 1) }

// Mul returns i*o reduced to 13 bits.
func (i I13) Mul(o I13) I13 { return I13.cast(i * o) }

// Div returns i divided by o.
func (i I13) Div(o I13) I13 { return I13.cast(i / o) }

// Mod returns i modulo o.
func (i I13) Mod(o I13) I13 { return I13.cast(i % o) }

// And returns i&o reduced to 13 bits.
func (i I13) And(o I13) I13 { return I13.cast(i & o) }

// Or returns i|o reduced to 13 bits.
func (i I13) Or(o I13) I13 { return I13.cast(i | o) }

// Xor returns i^o reduced to 13 bits.
func (i I13) Xor(o I13) I13 { return I13.cast(i ^ o) }

// Shr returns i>>o reduced to 13 bits.
func (i I13) Shr(o I13) I13 { return I13.cast(i >> o) }

// Shl returns i<<o reduced to 13 bits.
func (i I13) Shl(o I13) I13 { return I13.cast(i << o) }

// Not returns bitwise complement of i reduced to 13 bits.
func (i I13) Not() I13 { return I13.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I13) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I13) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U13 bit pattern.
func (i I13) Unsigned() U13 { return U13.cast(U13(i)) }

// Clamp returns value saturated into the representable range of I13.
func (i I13) Clamp(value int64) I13 { return iclamp[I13](value) }

// Clip masks i to a bits-wide low field and sign-extends to 13 bits.
func (i I13) Clip(bits int) I13 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I13(uint64(i)&m^b) - I13(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I13) Bit(index int) I13 {
	if index < 0 {
		index = 13 + index
	}
	return (i >> I13(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I13) SetBit(index int, v I13) {
	if index < 0 {
		index = 13 + index
	}
	bit := I13(1) << I13(index)
	*i = I13.cast((*i &^ bit) | ((v & 1) << I13(index)))
}

// Bitref returns a Range for bit index.
func (i *I13) Bitref(index int) Range[I13] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I13) Bits(lo, hi int) I13 {
	p := 13
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I13((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I13(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I13) SetBits(lo, hi int, v I13) {
	p := 13
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I13((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I13.cast((*i &^ mask) | ((v << I13(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I13) Bitsref(lo, hi int) Range[I13] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I13) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I13) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I13(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I13) Byteref(idx int) Range[I13] { return byte(i, idx) }

// U14 is an 14-bit unsigned integer.
type U14 uint16

// AsU14 returns a U14 representing v reduced to 14 bits.
func AsU14[T Unsigned](v T) U14 { return U14.cast(U14(v)) }

func (U14) nbits() int  { return 14 }
func (u U14) mask() U14 { return 0x3fff }
func (u U14) cast() U14 { return u & u.mask() }

// Set assigns v reduced to 14 bits to u.
func (u *U14) Set(v U14) U14 { *u = new(U14(v)).cast(); return U14.cast(v) }

// Add returns u+o reduced to 14 bits.
func (u U14) Add(o U14) U14 { return U14.cast(u + o) }

// Sub returns u-o reduced to 14 bits.
func (u U14) Sub(o U14) U14 { return U14.cast(u - o) }

// Inc returns u+1 reduced to 14 bits.
func (u U14) Inc() U14 { return U14.cast(u + 1) }

// Dec returns u-1 reduced to 14 bits.
func (u U14) Dec() U14 { return U14.cast(u - 1) }

// Mul returns u*o reduced to 14 bits.
func (u U14) Mul(o U14) U14 { return U14.cast(u * o) }

// Div returns u divided by o.
func (u U14) Div(o U14) U14 { return U14.cast(u / o) }

// Mod returns u modulo o.
func (u U14) Mod(o U14) U14 { return U14.cast(u % o) }

// And returns u&o reduced to 14 bits.
func (u U14) And(o U14) U14 { return U14.cast(u & o) }

// Or returns u|o reduced to 14 bits.
func (u U14) Or(o U14) U14 { return U14.cast(u | o) }

// Xor returns u^o reduced to 14 bits.
func (u U14) Xor(o U14) U14 { return U14.cast(u ^ o) }

// Shr returns u>>o reduced to 14 bits.
func (u U14) Shr(o U14) U14 { return U14.cast(u >> o) }

// Shl returns u<<o reduced to 14 bits.
func (u U14) Shl(o U14) U14 { return U14.cast(u << o) }

// Not returns bitwise complement of u reduced to 14 bits.
func (u U14) Not() U14 { return U14.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U14) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U14) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U14 value.
func (u U14) Signed() I14 { return I14.cast(I14(u)) }

// Clamp returns value saturated into the representable range of U14.
func (u U14) Clamp(value uint64) U14 { return uclamp[U14](value) }

// Clip masks u to a bits-wide low field.
func (u U14) Clip(bits int) U14 {
	b := 1 << (bits - 1)
	m := U14(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U14) Bit(index int) U14 {
	if index < 0 {
		index = 14 + index
	}
	return (u >> U14(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U14) SetBit(index int, v U14) {
	if index < 0 {
		index = 14 + index
	}
	bit := U14(1) << U14(index)
	*u = U14.cast((*u &^ bit) | ((v & 1) << U14(index)))
}

// Bitref returns a Range for bit index.
func (u *U14) Bitref(index int) Range[U14] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U14) Bits(lo, hi int) U14 {
	p := 14
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U14((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U14(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U14) SetBits(lo, hi int, v U14) {
	p := 14
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U14((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U14.cast((*u &^ mask) | ((v << U14(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U14) Bitsref(lo, hi int) Range[U14] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U14) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U14) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U14(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U14) Byteref(idx int) Range[U14] { return byte(u, idx) }

// I14 is an 14-bit signed integer in two's complement.
type I14 int16

// AsI14 returns a I14 representing v reduced to 14 bits.
func AsI14[T Signed](v T) I14 { return I14.cast(I14(v)) }

func (I14) nbits() int  { return 14 }
func (i I14) mask() I14 { return 0x3fff }
func (i I14) sign() I14 { return 1 << (i.nbits() - 1) }
func (i I14) cast() I14 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 14 bits to u.
func (i *I14) Set(v I14) I14 { *i = new(I14(v)).cast(); return I14.cast(v) }

// Add returns i+o reduced to 14 bits.
func (i I14) Add(o I14) I14 { return I14.cast(i + o) }

// Sub returns i-o reduced to 14 bits.
func (i I14) Sub(o I14) I14 { return I14.cast(i - o) }

// Inc returns i+1 reduced to 14 bits.
func (i I14) Inc() I14 { return I14.cast(i + 1) }

// Dec returns i-1 reduced to 14 bits.
func (i I14) Dec() I14 { return I14.cast(i - 1) }

// Mul returns i*o reduced to 14 bits.
func (i I14) Mul(o I14) I14 { return I14.cast(i * o) }

// Div returns i divided by o.
func (i I14) Div(o I14) I14 { return I14.cast(i / o) }

// Mod returns i modulo o.
func (i I14) Mod(o I14) I14 { return I14.cast(i % o) }

// And returns i&o reduced to 14 bits.
func (i I14) And(o I14) I14 { return I14.cast(i & o) }

// Or returns i|o reduced to 14 bits.
func (i I14) Or(o I14) I14 { return I14.cast(i | o) }

// Xor returns i^o reduced to 14 bits.
func (i I14) Xor(o I14) I14 { return I14.cast(i ^ o) }

// Shr returns i>>o reduced to 14 bits.
func (i I14) Shr(o I14) I14 { return I14.cast(i >> o) }

// Shl returns i<<o reduced to 14 bits.
func (i I14) Shl(o I14) I14 { return I14.cast(i << o) }

// Not returns bitwise complement of i reduced to 14 bits.
func (i I14) Not() I14 { return I14.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I14) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I14) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U14 bit pattern.
func (i I14) Unsigned() U14 { return U14.cast(U14(i)) }

// Clamp returns value saturated into the representable range of I14.
func (i I14) Clamp(value int64) I14 { return iclamp[I14](value) }

// Clip masks i to a bits-wide low field and sign-extends to 14 bits.
func (i I14) Clip(bits int) I14 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I14(uint64(i)&m^b) - I14(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I14) Bit(index int) I14 {
	if index < 0 {
		index = 14 + index
	}
	return (i >> I14(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I14) SetBit(index int, v I14) {
	if index < 0 {
		index = 14 + index
	}
	bit := I14(1) << I14(index)
	*i = I14.cast((*i &^ bit) | ((v & 1) << I14(index)))
}

// Bitref returns a Range for bit index.
func (i *I14) Bitref(index int) Range[I14] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I14) Bits(lo, hi int) I14 {
	p := 14
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I14((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I14(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I14) SetBits(lo, hi int, v I14) {
	p := 14
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I14((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I14.cast((*i &^ mask) | ((v << I14(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I14) Bitsref(lo, hi int) Range[I14] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I14) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I14) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I14(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I14) Byteref(idx int) Range[I14] { return byte(i, idx) }

// U15 is an 15-bit unsigned integer.
type U15 uint16

// AsU15 returns a U15 representing v reduced to 15 bits.
func AsU15[T Unsigned](v T) U15 { return U15.cast(U15(v)) }

func (U15) nbits() int  { return 15 }
func (u U15) mask() U15 { return 0x7fff }
func (u U15) cast() U15 { return u & u.mask() }

// Set assigns v reduced to 15 bits to u.
func (u *U15) Set(v U15) U15 { *u = new(U15(v)).cast(); return U15.cast(v) }

// Add returns u+o reduced to 15 bits.
func (u U15) Add(o U15) U15 { return U15.cast(u + o) }

// Sub returns u-o reduced to 15 bits.
func (u U15) Sub(o U15) U15 { return U15.cast(u - o) }

// Inc returns u+1 reduced to 15 bits.
func (u U15) Inc() U15 { return U15.cast(u + 1) }

// Dec returns u-1 reduced to 15 bits.
func (u U15) Dec() U15 { return U15.cast(u - 1) }

// Mul returns u*o reduced to 15 bits.
func (u U15) Mul(o U15) U15 { return U15.cast(u * o) }

// Div returns u divided by o.
func (u U15) Div(o U15) U15 { return U15.cast(u / o) }

// Mod returns u modulo o.
func (u U15) Mod(o U15) U15 { return U15.cast(u % o) }

// And returns u&o reduced to 15 bits.
func (u U15) And(o U15) U15 { return U15.cast(u & o) }

// Or returns u|o reduced to 15 bits.
func (u U15) Or(o U15) U15 { return U15.cast(u | o) }

// Xor returns u^o reduced to 15 bits.
func (u U15) Xor(o U15) U15 { return U15.cast(u ^ o) }

// Shr returns u>>o reduced to 15 bits.
func (u U15) Shr(o U15) U15 { return U15.cast(u >> o) }

// Shl returns u<<o reduced to 15 bits.
func (u U15) Shl(o U15) U15 { return U15.cast(u << o) }

// Not returns bitwise complement of u reduced to 15 bits.
func (u U15) Not() U15 { return U15.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U15) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U15) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U15 value.
func (u U15) Signed() I15 { return I15.cast(I15(u)) }

// Clamp returns value saturated into the representable range of U15.
func (u U15) Clamp(value uint64) U15 { return uclamp[U15](value) }

// Clip masks u to a bits-wide low field.
func (u U15) Clip(bits int) U15 {
	b := 1 << (bits - 1)
	m := U15(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U15) Bit(index int) U15 {
	if index < 0 {
		index = 15 + index
	}
	return (u >> U15(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U15) SetBit(index int, v U15) {
	if index < 0 {
		index = 15 + index
	}
	bit := U15(1) << U15(index)
	*u = U15.cast((*u &^ bit) | ((v & 1) << U15(index)))
}

// Bitref returns a Range for bit index.
func (u *U15) Bitref(index int) Range[U15] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U15) Bits(lo, hi int) U15 {
	p := 15
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U15((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U15(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U15) SetBits(lo, hi int, v U15) {
	p := 15
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U15((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U15.cast((*u &^ mask) | ((v << U15(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U15) Bitsref(lo, hi int) Range[U15] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U15) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U15) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U15(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U15) Byteref(idx int) Range[U15] { return byte(u, idx) }

// I15 is an 15-bit signed integer in two's complement.
type I15 int16

// AsI15 returns a I15 representing v reduced to 15 bits.
func AsI15[T Signed](v T) I15 { return I15.cast(I15(v)) }

func (I15) nbits() int  { return 15 }
func (i I15) mask() I15 { return 0x7fff }
func (i I15) sign() I15 { return 1 << (i.nbits() - 1) }
func (i I15) cast() I15 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 15 bits to u.
func (i *I15) Set(v I15) I15 { *i = new(I15(v)).cast(); return I15.cast(v) }

// Add returns i+o reduced to 15 bits.
func (i I15) Add(o I15) I15 { return I15.cast(i + o) }

// Sub returns i-o reduced to 15 bits.
func (i I15) Sub(o I15) I15 { return I15.cast(i - o) }

// Inc returns i+1 reduced to 15 bits.
func (i I15) Inc() I15 { return I15.cast(i + 1) }

// Dec returns i-1 reduced to 15 bits.
func (i I15) Dec() I15 { return I15.cast(i - 1) }

// Mul returns i*o reduced to 15 bits.
func (i I15) Mul(o I15) I15 { return I15.cast(i * o) }

// Div returns i divided by o.
func (i I15) Div(o I15) I15 { return I15.cast(i / o) }

// Mod returns i modulo o.
func (i I15) Mod(o I15) I15 { return I15.cast(i % o) }

// And returns i&o reduced to 15 bits.
func (i I15) And(o I15) I15 { return I15.cast(i & o) }

// Or returns i|o reduced to 15 bits.
func (i I15) Or(o I15) I15 { return I15.cast(i | o) }

// Xor returns i^o reduced to 15 bits.
func (i I15) Xor(o I15) I15 { return I15.cast(i ^ o) }

// Shr returns i>>o reduced to 15 bits.
func (i I15) Shr(o I15) I15 { return I15.cast(i >> o) }

// Shl returns i<<o reduced to 15 bits.
func (i I15) Shl(o I15) I15 { return I15.cast(i << o) }

// Not returns bitwise complement of i reduced to 15 bits.
func (i I15) Not() I15 { return I15.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I15) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I15) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U15 bit pattern.
func (i I15) Unsigned() U15 { return U15.cast(U15(i)) }

// Clamp returns value saturated into the representable range of I15.
func (i I15) Clamp(value int64) I15 { return iclamp[I15](value) }

// Clip masks i to a bits-wide low field and sign-extends to 15 bits.
func (i I15) Clip(bits int) I15 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I15(uint64(i)&m^b) - I15(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I15) Bit(index int) I15 {
	if index < 0 {
		index = 15 + index
	}
	return (i >> I15(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I15) SetBit(index int, v I15) {
	if index < 0 {
		index = 15 + index
	}
	bit := I15(1) << I15(index)
	*i = I15.cast((*i &^ bit) | ((v & 1) << I15(index)))
}

// Bitref returns a Range for bit index.
func (i *I15) Bitref(index int) Range[I15] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I15) Bits(lo, hi int) I15 {
	p := 15
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I15((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I15(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I15) SetBits(lo, hi int, v I15) {
	p := 15
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I15((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I15.cast((*i &^ mask) | ((v << I15(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I15) Bitsref(lo, hi int) Range[I15] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I15) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I15) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I15(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I15) Byteref(idx int) Range[I15] { return byte(i, idx) }

// U16 is an 16-bit unsigned integer.
type U16 uint16

// AsU16 returns a U16 representing v reduced to 16 bits.
func AsU16[T Unsigned](v T) U16 { return U16.cast(U16(v)) }

func (U16) nbits() int  { return 16 }
func (u U16) mask() U16 { return 0xffff }
func (u U16) cast() U16 { return u & u.mask() }

// Set assigns v reduced to 16 bits to u.
func (u *U16) Set(v U16) U16 { *u = new(U16(v)).cast(); return U16.cast(v) }

// Add returns u+o reduced to 16 bits.
func (u U16) Add(o U16) U16 { return U16.cast(u + o) }

// Sub returns u-o reduced to 16 bits.
func (u U16) Sub(o U16) U16 { return U16.cast(u - o) }

// Inc returns u+1 reduced to 16 bits.
func (u U16) Inc() U16 { return U16.cast(u + 1) }

// Dec returns u-1 reduced to 16 bits.
func (u U16) Dec() U16 { return U16.cast(u - 1) }

// Mul returns u*o reduced to 16 bits.
func (u U16) Mul(o U16) U16 { return U16.cast(u * o) }

// Div returns u divided by o.
func (u U16) Div(o U16) U16 { return U16.cast(u / o) }

// Mod returns u modulo o.
func (u U16) Mod(o U16) U16 { return U16.cast(u % o) }

// And returns u&o reduced to 16 bits.
func (u U16) And(o U16) U16 { return U16.cast(u & o) }

// Or returns u|o reduced to 16 bits.
func (u U16) Or(o U16) U16 { return U16.cast(u | o) }

// Xor returns u^o reduced to 16 bits.
func (u U16) Xor(o U16) U16 { return U16.cast(u ^ o) }

// Shr returns u>>o reduced to 16 bits.
func (u U16) Shr(o U16) U16 { return U16.cast(u >> o) }

// Shl returns u<<o reduced to 16 bits.
func (u U16) Shl(o U16) U16 { return U16.cast(u << o) }

// Not returns bitwise complement of u reduced to 16 bits.
func (u U16) Not() U16 { return U16.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U16) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U16) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U16 value.
func (u U16) Signed() I16 { return I16.cast(I16(u)) }

// Clamp returns value saturated into the representable range of U16.
func (u U16) Clamp(value uint64) U16 { return uclamp[U16](value) }

// Clip masks u to a bits-wide low field.
func (u U16) Clip(bits int) U16 {
	b := 1 << (bits - 1)
	m := U16(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U16) Bit(index int) U16 {
	if index < 0 {
		index = 16 + index
	}
	return (u >> U16(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U16) SetBit(index int, v U16) {
	if index < 0 {
		index = 16 + index
	}
	bit := U16(1) << U16(index)
	*u = U16.cast((*u &^ bit) | ((v & 1) << U16(index)))
}

// Bitref returns a Range for bit index.
func (u *U16) Bitref(index int) Range[U16] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U16) Bits(lo, hi int) U16 {
	p := 16
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U16((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U16(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U16) SetBits(lo, hi int, v U16) {
	p := 16
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U16((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U16.cast((*u &^ mask) | ((v << U16(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U16) Bitsref(lo, hi int) Range[U16] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U16) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U16) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U16(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U16) Byteref(idx int) Range[U16] { return byte(u, idx) }

// I16 is an 16-bit signed integer in two's complement.
type I16 int16

// AsI16 returns a I16 representing v reduced to 16 bits.
func AsI16[T Signed](v T) I16 { return I16.cast(I16(v)) }

func (I16) nbits() int  { return 16 }
func (i I16) mask() I16 { return -1 }
func (i I16) sign() I16 { return 1 << (i.nbits() - 1) }
func (i I16) cast() I16 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 16 bits to u.
func (i *I16) Set(v I16) I16 { *i = new(I16(v)).cast(); return I16.cast(v) }

// Add returns i+o reduced to 16 bits.
func (i I16) Add(o I16) I16 { return I16.cast(i + o) }

// Sub returns i-o reduced to 16 bits.
func (i I16) Sub(o I16) I16 { return I16.cast(i - o) }

// Inc returns i+1 reduced to 16 bits.
func (i I16) Inc() I16 { return I16.cast(i + 1) }

// Dec returns i-1 reduced to 16 bits.
func (i I16) Dec() I16 { return I16.cast(i - 1) }

// Mul returns i*o reduced to 16 bits.
func (i I16) Mul(o I16) I16 { return I16.cast(i * o) }

// Div returns i divided by o.
func (i I16) Div(o I16) I16 { return I16.cast(i / o) }

// Mod returns i modulo o.
func (i I16) Mod(o I16) I16 { return I16.cast(i % o) }

// And returns i&o reduced to 16 bits.
func (i I16) And(o I16) I16 { return I16.cast(i & o) }

// Or returns i|o reduced to 16 bits.
func (i I16) Or(o I16) I16 { return I16.cast(i | o) }

// Xor returns i^o reduced to 16 bits.
func (i I16) Xor(o I16) I16 { return I16.cast(i ^ o) }

// Shr returns i>>o reduced to 16 bits.
func (i I16) Shr(o I16) I16 { return I16.cast(i >> o) }

// Shl returns i<<o reduced to 16 bits.
func (i I16) Shl(o I16) I16 { return I16.cast(i << o) }

// Not returns bitwise complement of i reduced to 16 bits.
func (i I16) Not() I16 { return I16.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I16) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I16) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U16 bit pattern.
func (i I16) Unsigned() U16 { return U16.cast(U16(i)) }

// Clamp returns value saturated into the representable range of I16.
func (i I16) Clamp(value int64) I16 { return iclamp[I16](value) }

// Clip masks i to a bits-wide low field and sign-extends to 16 bits.
func (i I16) Clip(bits int) I16 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I16(uint64(i)&m^b) - I16(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I16) Bit(index int) I16 {
	if index < 0 {
		index = 16 + index
	}
	return (i >> I16(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I16) SetBit(index int, v I16) {
	if index < 0 {
		index = 16 + index
	}
	bit := I16(1) << I16(index)
	*i = I16.cast((*i &^ bit) | ((v & 1) << I16(index)))
}

// Bitref returns a Range for bit index.
func (i *I16) Bitref(index int) Range[I16] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I16) Bits(lo, hi int) I16 {
	p := 16
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I16((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I16(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I16) SetBits(lo, hi int, v I16) {
	p := 16
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I16((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I16.cast((*i &^ mask) | ((v << I16(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I16) Bitsref(lo, hi int) Range[I16] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I16) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I16) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I16(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I16) Byteref(idx int) Range[I16] { return byte(i, idx) }

// U17 is an 17-bit unsigned integer.
type U17 uint32

// AsU17 returns a U17 representing v reduced to 17 bits.
func AsU17[T Unsigned](v T) U17 { return U17.cast(U17(v)) }

func (U17) nbits() int  { return 17 }
func (u U17) mask() U17 { return 0x1ffff }
func (u U17) cast() U17 { return u & u.mask() }

// Set assigns v reduced to 17 bits to u.
func (u *U17) Set(v U17) U17 { *u = new(U17(v)).cast(); return U17.cast(v) }

// Add returns u+o reduced to 17 bits.
func (u U17) Add(o U17) U17 { return U17.cast(u + o) }

// Sub returns u-o reduced to 17 bits.
func (u U17) Sub(o U17) U17 { return U17.cast(u - o) }

// Inc returns u+1 reduced to 17 bits.
func (u U17) Inc() U17 { return U17.cast(u + 1) }

// Dec returns u-1 reduced to 17 bits.
func (u U17) Dec() U17 { return U17.cast(u - 1) }

// Mul returns u*o reduced to 17 bits.
func (u U17) Mul(o U17) U17 { return U17.cast(u * o) }

// Div returns u divided by o.
func (u U17) Div(o U17) U17 { return U17.cast(u / o) }

// Mod returns u modulo o.
func (u U17) Mod(o U17) U17 { return U17.cast(u % o) }

// And returns u&o reduced to 17 bits.
func (u U17) And(o U17) U17 { return U17.cast(u & o) }

// Or returns u|o reduced to 17 bits.
func (u U17) Or(o U17) U17 { return U17.cast(u | o) }

// Xor returns u^o reduced to 17 bits.
func (u U17) Xor(o U17) U17 { return U17.cast(u ^ o) }

// Shr returns u>>o reduced to 17 bits.
func (u U17) Shr(o U17) U17 { return U17.cast(u >> o) }

// Shl returns u<<o reduced to 17 bits.
func (u U17) Shl(o U17) U17 { return U17.cast(u << o) }

// Not returns bitwise complement of u reduced to 17 bits.
func (u U17) Not() U17 { return U17.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U17) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U17) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U17 value.
func (u U17) Signed() I17 { return I17.cast(I17(u)) }

// Clamp returns value saturated into the representable range of U17.
func (u U17) Clamp(value uint64) U17 { return uclamp[U17](value) }

// Clip masks u to a bits-wide low field.
func (u U17) Clip(bits int) U17 {
	b := 1 << (bits - 1)
	m := U17(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U17) Bit(index int) U17 {
	if index < 0 {
		index = 17 + index
	}
	return (u >> U17(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U17) SetBit(index int, v U17) {
	if index < 0 {
		index = 17 + index
	}
	bit := U17(1) << U17(index)
	*u = U17.cast((*u &^ bit) | ((v & 1) << U17(index)))
}

// Bitref returns a Range for bit index.
func (u *U17) Bitref(index int) Range[U17] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U17) Bits(lo, hi int) U17 {
	p := 17
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U17((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U17(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U17) SetBits(lo, hi int, v U17) {
	p := 17
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U17((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U17.cast((*u &^ mask) | ((v << U17(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U17) Bitsref(lo, hi int) Range[U17] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U17) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U17) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U17(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U17) Byteref(idx int) Range[U17] { return byte(u, idx) }

// I17 is an 17-bit signed integer in two's complement.
type I17 int32

// AsI17 returns a I17 representing v reduced to 17 bits.
func AsI17[T Signed](v T) I17 { return I17.cast(I17(v)) }

func (I17) nbits() int  { return 17 }
func (i I17) mask() I17 { return 0x1ffff }
func (i I17) sign() I17 { return 1 << (i.nbits() - 1) }
func (i I17) cast() I17 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 17 bits to u.
func (i *I17) Set(v I17) I17 { *i = new(I17(v)).cast(); return I17.cast(v) }

// Add returns i+o reduced to 17 bits.
func (i I17) Add(o I17) I17 { return I17.cast(i + o) }

// Sub returns i-o reduced to 17 bits.
func (i I17) Sub(o I17) I17 { return I17.cast(i - o) }

// Inc returns i+1 reduced to 17 bits.
func (i I17) Inc() I17 { return I17.cast(i + 1) }

// Dec returns i-1 reduced to 17 bits.
func (i I17) Dec() I17 { return I17.cast(i - 1) }

// Mul returns i*o reduced to 17 bits.
func (i I17) Mul(o I17) I17 { return I17.cast(i * o) }

// Div returns i divided by o.
func (i I17) Div(o I17) I17 { return I17.cast(i / o) }

// Mod returns i modulo o.
func (i I17) Mod(o I17) I17 { return I17.cast(i % o) }

// And returns i&o reduced to 17 bits.
func (i I17) And(o I17) I17 { return I17.cast(i & o) }

// Or returns i|o reduced to 17 bits.
func (i I17) Or(o I17) I17 { return I17.cast(i | o) }

// Xor returns i^o reduced to 17 bits.
func (i I17) Xor(o I17) I17 { return I17.cast(i ^ o) }

// Shr returns i>>o reduced to 17 bits.
func (i I17) Shr(o I17) I17 { return I17.cast(i >> o) }

// Shl returns i<<o reduced to 17 bits.
func (i I17) Shl(o I17) I17 { return I17.cast(i << o) }

// Not returns bitwise complement of i reduced to 17 bits.
func (i I17) Not() I17 { return I17.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I17) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I17) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U17 bit pattern.
func (i I17) Unsigned() U17 { return U17.cast(U17(i)) }

// Clamp returns value saturated into the representable range of I17.
func (i I17) Clamp(value int64) I17 { return iclamp[I17](value) }

// Clip masks i to a bits-wide low field and sign-extends to 17 bits.
func (i I17) Clip(bits int) I17 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I17(uint64(i)&m^b) - I17(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I17) Bit(index int) I17 {
	if index < 0 {
		index = 17 + index
	}
	return (i >> I17(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I17) SetBit(index int, v I17) {
	if index < 0 {
		index = 17 + index
	}
	bit := I17(1) << I17(index)
	*i = I17.cast((*i &^ bit) | ((v & 1) << I17(index)))
}

// Bitref returns a Range for bit index.
func (i *I17) Bitref(index int) Range[I17] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I17) Bits(lo, hi int) I17 {
	p := 17
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I17((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I17(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I17) SetBits(lo, hi int, v I17) {
	p := 17
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I17((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I17.cast((*i &^ mask) | ((v << I17(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I17) Bitsref(lo, hi int) Range[I17] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I17) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I17) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I17(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I17) Byteref(idx int) Range[I17] { return byte(i, idx) }

// U18 is an 18-bit unsigned integer.
type U18 uint32

// AsU18 returns a U18 representing v reduced to 18 bits.
func AsU18[T Unsigned](v T) U18 { return U18.cast(U18(v)) }

func (U18) nbits() int  { return 18 }
func (u U18) mask() U18 { return 0x3ffff }
func (u U18) cast() U18 { return u & u.mask() }

// Set assigns v reduced to 18 bits to u.
func (u *U18) Set(v U18) U18 { *u = new(U18(v)).cast(); return U18.cast(v) }

// Add returns u+o reduced to 18 bits.
func (u U18) Add(o U18) U18 { return U18.cast(u + o) }

// Sub returns u-o reduced to 18 bits.
func (u U18) Sub(o U18) U18 { return U18.cast(u - o) }

// Inc returns u+1 reduced to 18 bits.
func (u U18) Inc() U18 { return U18.cast(u + 1) }

// Dec returns u-1 reduced to 18 bits.
func (u U18) Dec() U18 { return U18.cast(u - 1) }

// Mul returns u*o reduced to 18 bits.
func (u U18) Mul(o U18) U18 { return U18.cast(u * o) }

// Div returns u divided by o.
func (u U18) Div(o U18) U18 { return U18.cast(u / o) }

// Mod returns u modulo o.
func (u U18) Mod(o U18) U18 { return U18.cast(u % o) }

// And returns u&o reduced to 18 bits.
func (u U18) And(o U18) U18 { return U18.cast(u & o) }

// Or returns u|o reduced to 18 bits.
func (u U18) Or(o U18) U18 { return U18.cast(u | o) }

// Xor returns u^o reduced to 18 bits.
func (u U18) Xor(o U18) U18 { return U18.cast(u ^ o) }

// Shr returns u>>o reduced to 18 bits.
func (u U18) Shr(o U18) U18 { return U18.cast(u >> o) }

// Shl returns u<<o reduced to 18 bits.
func (u U18) Shl(o U18) U18 { return U18.cast(u << o) }

// Not returns bitwise complement of u reduced to 18 bits.
func (u U18) Not() U18 { return U18.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U18) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U18) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U18 value.
func (u U18) Signed() I18 { return I18.cast(I18(u)) }

// Clamp returns value saturated into the representable range of U18.
func (u U18) Clamp(value uint64) U18 { return uclamp[U18](value) }

// Clip masks u to a bits-wide low field.
func (u U18) Clip(bits int) U18 {
	b := 1 << (bits - 1)
	m := U18(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U18) Bit(index int) U18 {
	if index < 0 {
		index = 18 + index
	}
	return (u >> U18(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U18) SetBit(index int, v U18) {
	if index < 0 {
		index = 18 + index
	}
	bit := U18(1) << U18(index)
	*u = U18.cast((*u &^ bit) | ((v & 1) << U18(index)))
}

// Bitref returns a Range for bit index.
func (u *U18) Bitref(index int) Range[U18] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U18) Bits(lo, hi int) U18 {
	p := 18
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U18((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U18(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U18) SetBits(lo, hi int, v U18) {
	p := 18
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U18((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U18.cast((*u &^ mask) | ((v << U18(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U18) Bitsref(lo, hi int) Range[U18] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U18) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U18) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U18(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U18) Byteref(idx int) Range[U18] { return byte(u, idx) }

// I18 is an 18-bit signed integer in two's complement.
type I18 int32

// AsI18 returns a I18 representing v reduced to 18 bits.
func AsI18[T Signed](v T) I18 { return I18.cast(I18(v)) }

func (I18) nbits() int  { return 18 }
func (i I18) mask() I18 { return 0x3ffff }
func (i I18) sign() I18 { return 1 << (i.nbits() - 1) }
func (i I18) cast() I18 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 18 bits to u.
func (i *I18) Set(v I18) I18 { *i = new(I18(v)).cast(); return I18.cast(v) }

// Add returns i+o reduced to 18 bits.
func (i I18) Add(o I18) I18 { return I18.cast(i + o) }

// Sub returns i-o reduced to 18 bits.
func (i I18) Sub(o I18) I18 { return I18.cast(i - o) }

// Inc returns i+1 reduced to 18 bits.
func (i I18) Inc() I18 { return I18.cast(i + 1) }

// Dec returns i-1 reduced to 18 bits.
func (i I18) Dec() I18 { return I18.cast(i - 1) }

// Mul returns i*o reduced to 18 bits.
func (i I18) Mul(o I18) I18 { return I18.cast(i * o) }

// Div returns i divided by o.
func (i I18) Div(o I18) I18 { return I18.cast(i / o) }

// Mod returns i modulo o.
func (i I18) Mod(o I18) I18 { return I18.cast(i % o) }

// And returns i&o reduced to 18 bits.
func (i I18) And(o I18) I18 { return I18.cast(i & o) }

// Or returns i|o reduced to 18 bits.
func (i I18) Or(o I18) I18 { return I18.cast(i | o) }

// Xor returns i^o reduced to 18 bits.
func (i I18) Xor(o I18) I18 { return I18.cast(i ^ o) }

// Shr returns i>>o reduced to 18 bits.
func (i I18) Shr(o I18) I18 { return I18.cast(i >> o) }

// Shl returns i<<o reduced to 18 bits.
func (i I18) Shl(o I18) I18 { return I18.cast(i << o) }

// Not returns bitwise complement of i reduced to 18 bits.
func (i I18) Not() I18 { return I18.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I18) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I18) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U18 bit pattern.
func (i I18) Unsigned() U18 { return U18.cast(U18(i)) }

// Clamp returns value saturated into the representable range of I18.
func (i I18) Clamp(value int64) I18 { return iclamp[I18](value) }

// Clip masks i to a bits-wide low field and sign-extends to 18 bits.
func (i I18) Clip(bits int) I18 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I18(uint64(i)&m^b) - I18(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I18) Bit(index int) I18 {
	if index < 0 {
		index = 18 + index
	}
	return (i >> I18(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I18) SetBit(index int, v I18) {
	if index < 0 {
		index = 18 + index
	}
	bit := I18(1) << I18(index)
	*i = I18.cast((*i &^ bit) | ((v & 1) << I18(index)))
}

// Bitref returns a Range for bit index.
func (i *I18) Bitref(index int) Range[I18] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I18) Bits(lo, hi int) I18 {
	p := 18
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I18((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I18(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I18) SetBits(lo, hi int, v I18) {
	p := 18
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I18((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I18.cast((*i &^ mask) | ((v << I18(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I18) Bitsref(lo, hi int) Range[I18] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I18) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I18) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I18(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I18) Byteref(idx int) Range[I18] { return byte(i, idx) }

// U19 is an 19-bit unsigned integer.
type U19 uint32

// AsU19 returns a U19 representing v reduced to 19 bits.
func AsU19[T Unsigned](v T) U19 { return U19.cast(U19(v)) }

func (U19) nbits() int  { return 19 }
func (u U19) mask() U19 { return 0x7ffff }
func (u U19) cast() U19 { return u & u.mask() }

// Set assigns v reduced to 19 bits to u.
func (u *U19) Set(v U19) U19 { *u = new(U19(v)).cast(); return U19.cast(v) }

// Add returns u+o reduced to 19 bits.
func (u U19) Add(o U19) U19 { return U19.cast(u + o) }

// Sub returns u-o reduced to 19 bits.
func (u U19) Sub(o U19) U19 { return U19.cast(u - o) }

// Inc returns u+1 reduced to 19 bits.
func (u U19) Inc() U19 { return U19.cast(u + 1) }

// Dec returns u-1 reduced to 19 bits.
func (u U19) Dec() U19 { return U19.cast(u - 1) }

// Mul returns u*o reduced to 19 bits.
func (u U19) Mul(o U19) U19 { return U19.cast(u * o) }

// Div returns u divided by o.
func (u U19) Div(o U19) U19 { return U19.cast(u / o) }

// Mod returns u modulo o.
func (u U19) Mod(o U19) U19 { return U19.cast(u % o) }

// And returns u&o reduced to 19 bits.
func (u U19) And(o U19) U19 { return U19.cast(u & o) }

// Or returns u|o reduced to 19 bits.
func (u U19) Or(o U19) U19 { return U19.cast(u | o) }

// Xor returns u^o reduced to 19 bits.
func (u U19) Xor(o U19) U19 { return U19.cast(u ^ o) }

// Shr returns u>>o reduced to 19 bits.
func (u U19) Shr(o U19) U19 { return U19.cast(u >> o) }

// Shl returns u<<o reduced to 19 bits.
func (u U19) Shl(o U19) U19 { return U19.cast(u << o) }

// Not returns bitwise complement of u reduced to 19 bits.
func (u U19) Not() U19 { return U19.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U19) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U19) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U19 value.
func (u U19) Signed() I19 { return I19.cast(I19(u)) }

// Clamp returns value saturated into the representable range of U19.
func (u U19) Clamp(value uint64) U19 { return uclamp[U19](value) }

// Clip masks u to a bits-wide low field.
func (u U19) Clip(bits int) U19 {
	b := 1 << (bits - 1)
	m := U19(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U19) Bit(index int) U19 {
	if index < 0 {
		index = 19 + index
	}
	return (u >> U19(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U19) SetBit(index int, v U19) {
	if index < 0 {
		index = 19 + index
	}
	bit := U19(1) << U19(index)
	*u = U19.cast((*u &^ bit) | ((v & 1) << U19(index)))
}

// Bitref returns a Range for bit index.
func (u *U19) Bitref(index int) Range[U19] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U19) Bits(lo, hi int) U19 {
	p := 19
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U19((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U19(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U19) SetBits(lo, hi int, v U19) {
	p := 19
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U19((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U19.cast((*u &^ mask) | ((v << U19(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U19) Bitsref(lo, hi int) Range[U19] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U19) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U19) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U19(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U19) Byteref(idx int) Range[U19] { return byte(u, idx) }

// I19 is an 19-bit signed integer in two's complement.
type I19 int32

// AsI19 returns a I19 representing v reduced to 19 bits.
func AsI19[T Signed](v T) I19 { return I19.cast(I19(v)) }

func (I19) nbits() int  { return 19 }
func (i I19) mask() I19 { return 0x7ffff }
func (i I19) sign() I19 { return 1 << (i.nbits() - 1) }
func (i I19) cast() I19 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 19 bits to u.
func (i *I19) Set(v I19) I19 { *i = new(I19(v)).cast(); return I19.cast(v) }

// Add returns i+o reduced to 19 bits.
func (i I19) Add(o I19) I19 { return I19.cast(i + o) }

// Sub returns i-o reduced to 19 bits.
func (i I19) Sub(o I19) I19 { return I19.cast(i - o) }

// Inc returns i+1 reduced to 19 bits.
func (i I19) Inc() I19 { return I19.cast(i + 1) }

// Dec returns i-1 reduced to 19 bits.
func (i I19) Dec() I19 { return I19.cast(i - 1) }

// Mul returns i*o reduced to 19 bits.
func (i I19) Mul(o I19) I19 { return I19.cast(i * o) }

// Div returns i divided by o.
func (i I19) Div(o I19) I19 { return I19.cast(i / o) }

// Mod returns i modulo o.
func (i I19) Mod(o I19) I19 { return I19.cast(i % o) }

// And returns i&o reduced to 19 bits.
func (i I19) And(o I19) I19 { return I19.cast(i & o) }

// Or returns i|o reduced to 19 bits.
func (i I19) Or(o I19) I19 { return I19.cast(i | o) }

// Xor returns i^o reduced to 19 bits.
func (i I19) Xor(o I19) I19 { return I19.cast(i ^ o) }

// Shr returns i>>o reduced to 19 bits.
func (i I19) Shr(o I19) I19 { return I19.cast(i >> o) }

// Shl returns i<<o reduced to 19 bits.
func (i I19) Shl(o I19) I19 { return I19.cast(i << o) }

// Not returns bitwise complement of i reduced to 19 bits.
func (i I19) Not() I19 { return I19.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I19) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I19) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U19 bit pattern.
func (i I19) Unsigned() U19 { return U19.cast(U19(i)) }

// Clamp returns value saturated into the representable range of I19.
func (i I19) Clamp(value int64) I19 { return iclamp[I19](value) }

// Clip masks i to a bits-wide low field and sign-extends to 19 bits.
func (i I19) Clip(bits int) I19 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I19(uint64(i)&m^b) - I19(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I19) Bit(index int) I19 {
	if index < 0 {
		index = 19 + index
	}
	return (i >> I19(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I19) SetBit(index int, v I19) {
	if index < 0 {
		index = 19 + index
	}
	bit := I19(1) << I19(index)
	*i = I19.cast((*i &^ bit) | ((v & 1) << I19(index)))
}

// Bitref returns a Range for bit index.
func (i *I19) Bitref(index int) Range[I19] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I19) Bits(lo, hi int) I19 {
	p := 19
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I19((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I19(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I19) SetBits(lo, hi int, v I19) {
	p := 19
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I19((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I19.cast((*i &^ mask) | ((v << I19(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I19) Bitsref(lo, hi int) Range[I19] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I19) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I19) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I19(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I19) Byteref(idx int) Range[I19] { return byte(i, idx) }

// U20 is an 20-bit unsigned integer.
type U20 uint32

// AsU20 returns a U20 representing v reduced to 20 bits.
func AsU20[T Unsigned](v T) U20 { return U20.cast(U20(v)) }

func (U20) nbits() int  { return 20 }
func (u U20) mask() U20 { return 0xfffff }
func (u U20) cast() U20 { return u & u.mask() }

// Set assigns v reduced to 20 bits to u.
func (u *U20) Set(v U20) U20 { *u = new(U20(v)).cast(); return U20.cast(v) }

// Add returns u+o reduced to 20 bits.
func (u U20) Add(o U20) U20 { return U20.cast(u + o) }

// Sub returns u-o reduced to 20 bits.
func (u U20) Sub(o U20) U20 { return U20.cast(u - o) }

// Inc returns u+1 reduced to 20 bits.
func (u U20) Inc() U20 { return U20.cast(u + 1) }

// Dec returns u-1 reduced to 20 bits.
func (u U20) Dec() U20 { return U20.cast(u - 1) }

// Mul returns u*o reduced to 20 bits.
func (u U20) Mul(o U20) U20 { return U20.cast(u * o) }

// Div returns u divided by o.
func (u U20) Div(o U20) U20 { return U20.cast(u / o) }

// Mod returns u modulo o.
func (u U20) Mod(o U20) U20 { return U20.cast(u % o) }

// And returns u&o reduced to 20 bits.
func (u U20) And(o U20) U20 { return U20.cast(u & o) }

// Or returns u|o reduced to 20 bits.
func (u U20) Or(o U20) U20 { return U20.cast(u | o) }

// Xor returns u^o reduced to 20 bits.
func (u U20) Xor(o U20) U20 { return U20.cast(u ^ o) }

// Shr returns u>>o reduced to 20 bits.
func (u U20) Shr(o U20) U20 { return U20.cast(u >> o) }

// Shl returns u<<o reduced to 20 bits.
func (u U20) Shl(o U20) U20 { return U20.cast(u << o) }

// Not returns bitwise complement of u reduced to 20 bits.
func (u U20) Not() U20 { return U20.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U20) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U20) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U20 value.
func (u U20) Signed() I20 { return I20.cast(I20(u)) }

// Clamp returns value saturated into the representable range of U20.
func (u U20) Clamp(value uint64) U20 { return uclamp[U20](value) }

// Clip masks u to a bits-wide low field.
func (u U20) Clip(bits int) U20 {
	b := 1 << (bits - 1)
	m := U20(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U20) Bit(index int) U20 {
	if index < 0 {
		index = 20 + index
	}
	return (u >> U20(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U20) SetBit(index int, v U20) {
	if index < 0 {
		index = 20 + index
	}
	bit := U20(1) << U20(index)
	*u = U20.cast((*u &^ bit) | ((v & 1) << U20(index)))
}

// Bitref returns a Range for bit index.
func (u *U20) Bitref(index int) Range[U20] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U20) Bits(lo, hi int) U20 {
	p := 20
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U20((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U20(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U20) SetBits(lo, hi int, v U20) {
	p := 20
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U20((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U20.cast((*u &^ mask) | ((v << U20(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U20) Bitsref(lo, hi int) Range[U20] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U20) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U20) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U20(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U20) Byteref(idx int) Range[U20] { return byte(u, idx) }

// I20 is an 20-bit signed integer in two's complement.
type I20 int32

// AsI20 returns a I20 representing v reduced to 20 bits.
func AsI20[T Signed](v T) I20 { return I20.cast(I20(v)) }

func (I20) nbits() int  { return 20 }
func (i I20) mask() I20 { return 0xfffff }
func (i I20) sign() I20 { return 1 << (i.nbits() - 1) }
func (i I20) cast() I20 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 20 bits to u.
func (i *I20) Set(v I20) I20 { *i = new(I20(v)).cast(); return I20.cast(v) }

// Add returns i+o reduced to 20 bits.
func (i I20) Add(o I20) I20 { return I20.cast(i + o) }

// Sub returns i-o reduced to 20 bits.
func (i I20) Sub(o I20) I20 { return I20.cast(i - o) }

// Inc returns i+1 reduced to 20 bits.
func (i I20) Inc() I20 { return I20.cast(i + 1) }

// Dec returns i-1 reduced to 20 bits.
func (i I20) Dec() I20 { return I20.cast(i - 1) }

// Mul returns i*o reduced to 20 bits.
func (i I20) Mul(o I20) I20 { return I20.cast(i * o) }

// Div returns i divided by o.
func (i I20) Div(o I20) I20 { return I20.cast(i / o) }

// Mod returns i modulo o.
func (i I20) Mod(o I20) I20 { return I20.cast(i % o) }

// And returns i&o reduced to 20 bits.
func (i I20) And(o I20) I20 { return I20.cast(i & o) }

// Or returns i|o reduced to 20 bits.
func (i I20) Or(o I20) I20 { return I20.cast(i | o) }

// Xor returns i^o reduced to 20 bits.
func (i I20) Xor(o I20) I20 { return I20.cast(i ^ o) }

// Shr returns i>>o reduced to 20 bits.
func (i I20) Shr(o I20) I20 { return I20.cast(i >> o) }

// Shl returns i<<o reduced to 20 bits.
func (i I20) Shl(o I20) I20 { return I20.cast(i << o) }

// Not returns bitwise complement of i reduced to 20 bits.
func (i I20) Not() I20 { return I20.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I20) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I20) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U20 bit pattern.
func (i I20) Unsigned() U20 { return U20.cast(U20(i)) }

// Clamp returns value saturated into the representable range of I20.
func (i I20) Clamp(value int64) I20 { return iclamp[I20](value) }

// Clip masks i to a bits-wide low field and sign-extends to 20 bits.
func (i I20) Clip(bits int) I20 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I20(uint64(i)&m^b) - I20(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I20) Bit(index int) I20 {
	if index < 0 {
		index = 20 + index
	}
	return (i >> I20(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I20) SetBit(index int, v I20) {
	if index < 0 {
		index = 20 + index
	}
	bit := I20(1) << I20(index)
	*i = I20.cast((*i &^ bit) | ((v & 1) << I20(index)))
}

// Bitref returns a Range for bit index.
func (i *I20) Bitref(index int) Range[I20] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I20) Bits(lo, hi int) I20 {
	p := 20
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I20((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I20(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I20) SetBits(lo, hi int, v I20) {
	p := 20
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I20((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I20.cast((*i &^ mask) | ((v << I20(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I20) Bitsref(lo, hi int) Range[I20] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I20) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I20) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I20(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I20) Byteref(idx int) Range[I20] { return byte(i, idx) }

// U21 is an 21-bit unsigned integer.
type U21 uint32

// AsU21 returns a U21 representing v reduced to 21 bits.
func AsU21[T Unsigned](v T) U21 { return U21.cast(U21(v)) }

func (U21) nbits() int  { return 21 }
func (u U21) mask() U21 { return 0x1fffff }
func (u U21) cast() U21 { return u & u.mask() }

// Set assigns v reduced to 21 bits to u.
func (u *U21) Set(v U21) U21 { *u = new(U21(v)).cast(); return U21.cast(v) }

// Add returns u+o reduced to 21 bits.
func (u U21) Add(o U21) U21 { return U21.cast(u + o) }

// Sub returns u-o reduced to 21 bits.
func (u U21) Sub(o U21) U21 { return U21.cast(u - o) }

// Inc returns u+1 reduced to 21 bits.
func (u U21) Inc() U21 { return U21.cast(u + 1) }

// Dec returns u-1 reduced to 21 bits.
func (u U21) Dec() U21 { return U21.cast(u - 1) }

// Mul returns u*o reduced to 21 bits.
func (u U21) Mul(o U21) U21 { return U21.cast(u * o) }

// Div returns u divided by o.
func (u U21) Div(o U21) U21 { return U21.cast(u / o) }

// Mod returns u modulo o.
func (u U21) Mod(o U21) U21 { return U21.cast(u % o) }

// And returns u&o reduced to 21 bits.
func (u U21) And(o U21) U21 { return U21.cast(u & o) }

// Or returns u|o reduced to 21 bits.
func (u U21) Or(o U21) U21 { return U21.cast(u | o) }

// Xor returns u^o reduced to 21 bits.
func (u U21) Xor(o U21) U21 { return U21.cast(u ^ o) }

// Shr returns u>>o reduced to 21 bits.
func (u U21) Shr(o U21) U21 { return U21.cast(u >> o) }

// Shl returns u<<o reduced to 21 bits.
func (u U21) Shl(o U21) U21 { return U21.cast(u << o) }

// Not returns bitwise complement of u reduced to 21 bits.
func (u U21) Not() U21 { return U21.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U21) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U21) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U21 value.
func (u U21) Signed() I21 { return I21.cast(I21(u)) }

// Clamp returns value saturated into the representable range of U21.
func (u U21) Clamp(value uint64) U21 { return uclamp[U21](value) }

// Clip masks u to a bits-wide low field.
func (u U21) Clip(bits int) U21 {
	b := 1 << (bits - 1)
	m := U21(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U21) Bit(index int) U21 {
	if index < 0 {
		index = 21 + index
	}
	return (u >> U21(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U21) SetBit(index int, v U21) {
	if index < 0 {
		index = 21 + index
	}
	bit := U21(1) << U21(index)
	*u = U21.cast((*u &^ bit) | ((v & 1) << U21(index)))
}

// Bitref returns a Range for bit index.
func (u *U21) Bitref(index int) Range[U21] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U21) Bits(lo, hi int) U21 {
	p := 21
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U21((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U21(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U21) SetBits(lo, hi int, v U21) {
	p := 21
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U21((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U21.cast((*u &^ mask) | ((v << U21(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U21) Bitsref(lo, hi int) Range[U21] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U21) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U21) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U21(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U21) Byteref(idx int) Range[U21] { return byte(u, idx) }

// I21 is an 21-bit signed integer in two's complement.
type I21 int32

// AsI21 returns a I21 representing v reduced to 21 bits.
func AsI21[T Signed](v T) I21 { return I21.cast(I21(v)) }

func (I21) nbits() int  { return 21 }
func (i I21) mask() I21 { return 0x1fffff }
func (i I21) sign() I21 { return 1 << (i.nbits() - 1) }
func (i I21) cast() I21 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 21 bits to u.
func (i *I21) Set(v I21) I21 { *i = new(I21(v)).cast(); return I21.cast(v) }

// Add returns i+o reduced to 21 bits.
func (i I21) Add(o I21) I21 { return I21.cast(i + o) }

// Sub returns i-o reduced to 21 bits.
func (i I21) Sub(o I21) I21 { return I21.cast(i - o) }

// Inc returns i+1 reduced to 21 bits.
func (i I21) Inc() I21 { return I21.cast(i + 1) }

// Dec returns i-1 reduced to 21 bits.
func (i I21) Dec() I21 { return I21.cast(i - 1) }

// Mul returns i*o reduced to 21 bits.
func (i I21) Mul(o I21) I21 { return I21.cast(i * o) }

// Div returns i divided by o.
func (i I21) Div(o I21) I21 { return I21.cast(i / o) }

// Mod returns i modulo o.
func (i I21) Mod(o I21) I21 { return I21.cast(i % o) }

// And returns i&o reduced to 21 bits.
func (i I21) And(o I21) I21 { return I21.cast(i & o) }

// Or returns i|o reduced to 21 bits.
func (i I21) Or(o I21) I21 { return I21.cast(i | o) }

// Xor returns i^o reduced to 21 bits.
func (i I21) Xor(o I21) I21 { return I21.cast(i ^ o) }

// Shr returns i>>o reduced to 21 bits.
func (i I21) Shr(o I21) I21 { return I21.cast(i >> o) }

// Shl returns i<<o reduced to 21 bits.
func (i I21) Shl(o I21) I21 { return I21.cast(i << o) }

// Not returns bitwise complement of i reduced to 21 bits.
func (i I21) Not() I21 { return I21.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I21) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I21) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U21 bit pattern.
func (i I21) Unsigned() U21 { return U21.cast(U21(i)) }

// Clamp returns value saturated into the representable range of I21.
func (i I21) Clamp(value int64) I21 { return iclamp[I21](value) }

// Clip masks i to a bits-wide low field and sign-extends to 21 bits.
func (i I21) Clip(bits int) I21 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I21(uint64(i)&m^b) - I21(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I21) Bit(index int) I21 {
	if index < 0 {
		index = 21 + index
	}
	return (i >> I21(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I21) SetBit(index int, v I21) {
	if index < 0 {
		index = 21 + index
	}
	bit := I21(1) << I21(index)
	*i = I21.cast((*i &^ bit) | ((v & 1) << I21(index)))
}

// Bitref returns a Range for bit index.
func (i *I21) Bitref(index int) Range[I21] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I21) Bits(lo, hi int) I21 {
	p := 21
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I21((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I21(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I21) SetBits(lo, hi int, v I21) {
	p := 21
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I21((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I21.cast((*i &^ mask) | ((v << I21(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I21) Bitsref(lo, hi int) Range[I21] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I21) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I21) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I21(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I21) Byteref(idx int) Range[I21] { return byte(i, idx) }

// U22 is an 22-bit unsigned integer.
type U22 uint32

// AsU22 returns a U22 representing v reduced to 22 bits.
func AsU22[T Unsigned](v T) U22 { return U22.cast(U22(v)) }

func (U22) nbits() int  { return 22 }
func (u U22) mask() U22 { return 0x3fffff }
func (u U22) cast() U22 { return u & u.mask() }

// Set assigns v reduced to 22 bits to u.
func (u *U22) Set(v U22) U22 { *u = new(U22(v)).cast(); return U22.cast(v) }

// Add returns u+o reduced to 22 bits.
func (u U22) Add(o U22) U22 { return U22.cast(u + o) }

// Sub returns u-o reduced to 22 bits.
func (u U22) Sub(o U22) U22 { return U22.cast(u - o) }

// Inc returns u+1 reduced to 22 bits.
func (u U22) Inc() U22 { return U22.cast(u + 1) }

// Dec returns u-1 reduced to 22 bits.
func (u U22) Dec() U22 { return U22.cast(u - 1) }

// Mul returns u*o reduced to 22 bits.
func (u U22) Mul(o U22) U22 { return U22.cast(u * o) }

// Div returns u divided by o.
func (u U22) Div(o U22) U22 { return U22.cast(u / o) }

// Mod returns u modulo o.
func (u U22) Mod(o U22) U22 { return U22.cast(u % o) }

// And returns u&o reduced to 22 bits.
func (u U22) And(o U22) U22 { return U22.cast(u & o) }

// Or returns u|o reduced to 22 bits.
func (u U22) Or(o U22) U22 { return U22.cast(u | o) }

// Xor returns u^o reduced to 22 bits.
func (u U22) Xor(o U22) U22 { return U22.cast(u ^ o) }

// Shr returns u>>o reduced to 22 bits.
func (u U22) Shr(o U22) U22 { return U22.cast(u >> o) }

// Shl returns u<<o reduced to 22 bits.
func (u U22) Shl(o U22) U22 { return U22.cast(u << o) }

// Not returns bitwise complement of u reduced to 22 bits.
func (u U22) Not() U22 { return U22.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U22) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U22) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U22 value.
func (u U22) Signed() I22 { return I22.cast(I22(u)) }

// Clamp returns value saturated into the representable range of U22.
func (u U22) Clamp(value uint64) U22 { return uclamp[U22](value) }

// Clip masks u to a bits-wide low field.
func (u U22) Clip(bits int) U22 {
	b := 1 << (bits - 1)
	m := U22(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U22) Bit(index int) U22 {
	if index < 0 {
		index = 22 + index
	}
	return (u >> U22(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U22) SetBit(index int, v U22) {
	if index < 0 {
		index = 22 + index
	}
	bit := U22(1) << U22(index)
	*u = U22.cast((*u &^ bit) | ((v & 1) << U22(index)))
}

// Bitref returns a Range for bit index.
func (u *U22) Bitref(index int) Range[U22] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U22) Bits(lo, hi int) U22 {
	p := 22
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U22((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U22(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U22) SetBits(lo, hi int, v U22) {
	p := 22
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U22((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U22.cast((*u &^ mask) | ((v << U22(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U22) Bitsref(lo, hi int) Range[U22] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U22) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U22) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U22(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U22) Byteref(idx int) Range[U22] { return byte(u, idx) }

// I22 is an 22-bit signed integer in two's complement.
type I22 int32

// AsI22 returns a I22 representing v reduced to 22 bits.
func AsI22[T Signed](v T) I22 { return I22.cast(I22(v)) }

func (I22) nbits() int  { return 22 }
func (i I22) mask() I22 { return 0x3fffff }
func (i I22) sign() I22 { return 1 << (i.nbits() - 1) }
func (i I22) cast() I22 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 22 bits to u.
func (i *I22) Set(v I22) I22 { *i = new(I22(v)).cast(); return I22.cast(v) }

// Add returns i+o reduced to 22 bits.
func (i I22) Add(o I22) I22 { return I22.cast(i + o) }

// Sub returns i-o reduced to 22 bits.
func (i I22) Sub(o I22) I22 { return I22.cast(i - o) }

// Inc returns i+1 reduced to 22 bits.
func (i I22) Inc() I22 { return I22.cast(i + 1) }

// Dec returns i-1 reduced to 22 bits.
func (i I22) Dec() I22 { return I22.cast(i - 1) }

// Mul returns i*o reduced to 22 bits.
func (i I22) Mul(o I22) I22 { return I22.cast(i * o) }

// Div returns i divided by o.
func (i I22) Div(o I22) I22 { return I22.cast(i / o) }

// Mod returns i modulo o.
func (i I22) Mod(o I22) I22 { return I22.cast(i % o) }

// And returns i&o reduced to 22 bits.
func (i I22) And(o I22) I22 { return I22.cast(i & o) }

// Or returns i|o reduced to 22 bits.
func (i I22) Or(o I22) I22 { return I22.cast(i | o) }

// Xor returns i^o reduced to 22 bits.
func (i I22) Xor(o I22) I22 { return I22.cast(i ^ o) }

// Shr returns i>>o reduced to 22 bits.
func (i I22) Shr(o I22) I22 { return I22.cast(i >> o) }

// Shl returns i<<o reduced to 22 bits.
func (i I22) Shl(o I22) I22 { return I22.cast(i << o) }

// Not returns bitwise complement of i reduced to 22 bits.
func (i I22) Not() I22 { return I22.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I22) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I22) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U22 bit pattern.
func (i I22) Unsigned() U22 { return U22.cast(U22(i)) }

// Clamp returns value saturated into the representable range of I22.
func (i I22) Clamp(value int64) I22 { return iclamp[I22](value) }

// Clip masks i to a bits-wide low field and sign-extends to 22 bits.
func (i I22) Clip(bits int) I22 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I22(uint64(i)&m^b) - I22(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I22) Bit(index int) I22 {
	if index < 0 {
		index = 22 + index
	}
	return (i >> I22(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I22) SetBit(index int, v I22) {
	if index < 0 {
		index = 22 + index
	}
	bit := I22(1) << I22(index)
	*i = I22.cast((*i &^ bit) | ((v & 1) << I22(index)))
}

// Bitref returns a Range for bit index.
func (i *I22) Bitref(index int) Range[I22] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I22) Bits(lo, hi int) I22 {
	p := 22
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I22((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I22(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I22) SetBits(lo, hi int, v I22) {
	p := 22
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I22((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I22.cast((*i &^ mask) | ((v << I22(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I22) Bitsref(lo, hi int) Range[I22] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I22) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I22) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I22(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I22) Byteref(idx int) Range[I22] { return byte(i, idx) }

// U23 is an 23-bit unsigned integer.
type U23 uint32

// AsU23 returns a U23 representing v reduced to 23 bits.
func AsU23[T Unsigned](v T) U23 { return U23.cast(U23(v)) }

func (U23) nbits() int  { return 23 }
func (u U23) mask() U23 { return 0x7fffff }
func (u U23) cast() U23 { return u & u.mask() }

// Set assigns v reduced to 23 bits to u.
func (u *U23) Set(v U23) U23 { *u = new(U23(v)).cast(); return U23.cast(v) }

// Add returns u+o reduced to 23 bits.
func (u U23) Add(o U23) U23 { return U23.cast(u + o) }

// Sub returns u-o reduced to 23 bits.
func (u U23) Sub(o U23) U23 { return U23.cast(u - o) }

// Inc returns u+1 reduced to 23 bits.
func (u U23) Inc() U23 { return U23.cast(u + 1) }

// Dec returns u-1 reduced to 23 bits.
func (u U23) Dec() U23 { return U23.cast(u - 1) }

// Mul returns u*o reduced to 23 bits.
func (u U23) Mul(o U23) U23 { return U23.cast(u * o) }

// Div returns u divided by o.
func (u U23) Div(o U23) U23 { return U23.cast(u / o) }

// Mod returns u modulo o.
func (u U23) Mod(o U23) U23 { return U23.cast(u % o) }

// And returns u&o reduced to 23 bits.
func (u U23) And(o U23) U23 { return U23.cast(u & o) }

// Or returns u|o reduced to 23 bits.
func (u U23) Or(o U23) U23 { return U23.cast(u | o) }

// Xor returns u^o reduced to 23 bits.
func (u U23) Xor(o U23) U23 { return U23.cast(u ^ o) }

// Shr returns u>>o reduced to 23 bits.
func (u U23) Shr(o U23) U23 { return U23.cast(u >> o) }

// Shl returns u<<o reduced to 23 bits.
func (u U23) Shl(o U23) U23 { return U23.cast(u << o) }

// Not returns bitwise complement of u reduced to 23 bits.
func (u U23) Not() U23 { return U23.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U23) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U23) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U23 value.
func (u U23) Signed() I23 { return I23.cast(I23(u)) }

// Clamp returns value saturated into the representable range of U23.
func (u U23) Clamp(value uint64) U23 { return uclamp[U23](value) }

// Clip masks u to a bits-wide low field.
func (u U23) Clip(bits int) U23 {
	b := 1 << (bits - 1)
	m := U23(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U23) Bit(index int) U23 {
	if index < 0 {
		index = 23 + index
	}
	return (u >> U23(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U23) SetBit(index int, v U23) {
	if index < 0 {
		index = 23 + index
	}
	bit := U23(1) << U23(index)
	*u = U23.cast((*u &^ bit) | ((v & 1) << U23(index)))
}

// Bitref returns a Range for bit index.
func (u *U23) Bitref(index int) Range[U23] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U23) Bits(lo, hi int) U23 {
	p := 23
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U23((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U23(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U23) SetBits(lo, hi int, v U23) {
	p := 23
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U23((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U23.cast((*u &^ mask) | ((v << U23(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U23) Bitsref(lo, hi int) Range[U23] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U23) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U23) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U23(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U23) Byteref(idx int) Range[U23] { return byte(u, idx) }

// I23 is an 23-bit signed integer in two's complement.
type I23 int32

// AsI23 returns a I23 representing v reduced to 23 bits.
func AsI23[T Signed](v T) I23 { return I23.cast(I23(v)) }

func (I23) nbits() int  { return 23 }
func (i I23) mask() I23 { return 0x7fffff }
func (i I23) sign() I23 { return 1 << (i.nbits() - 1) }
func (i I23) cast() I23 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 23 bits to u.
func (i *I23) Set(v I23) I23 { *i = new(I23(v)).cast(); return I23.cast(v) }

// Add returns i+o reduced to 23 bits.
func (i I23) Add(o I23) I23 { return I23.cast(i + o) }

// Sub returns i-o reduced to 23 bits.
func (i I23) Sub(o I23) I23 { return I23.cast(i - o) }

// Inc returns i+1 reduced to 23 bits.
func (i I23) Inc() I23 { return I23.cast(i + 1) }

// Dec returns i-1 reduced to 23 bits.
func (i I23) Dec() I23 { return I23.cast(i - 1) }

// Mul returns i*o reduced to 23 bits.
func (i I23) Mul(o I23) I23 { return I23.cast(i * o) }

// Div returns i divided by o.
func (i I23) Div(o I23) I23 { return I23.cast(i / o) }

// Mod returns i modulo o.
func (i I23) Mod(o I23) I23 { return I23.cast(i % o) }

// And returns i&o reduced to 23 bits.
func (i I23) And(o I23) I23 { return I23.cast(i & o) }

// Or returns i|o reduced to 23 bits.
func (i I23) Or(o I23) I23 { return I23.cast(i | o) }

// Xor returns i^o reduced to 23 bits.
func (i I23) Xor(o I23) I23 { return I23.cast(i ^ o) }

// Shr returns i>>o reduced to 23 bits.
func (i I23) Shr(o I23) I23 { return I23.cast(i >> o) }

// Shl returns i<<o reduced to 23 bits.
func (i I23) Shl(o I23) I23 { return I23.cast(i << o) }

// Not returns bitwise complement of i reduced to 23 bits.
func (i I23) Not() I23 { return I23.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I23) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I23) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U23 bit pattern.
func (i I23) Unsigned() U23 { return U23.cast(U23(i)) }

// Clamp returns value saturated into the representable range of I23.
func (i I23) Clamp(value int64) I23 { return iclamp[I23](value) }

// Clip masks i to a bits-wide low field and sign-extends to 23 bits.
func (i I23) Clip(bits int) I23 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I23(uint64(i)&m^b) - I23(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I23) Bit(index int) I23 {
	if index < 0 {
		index = 23 + index
	}
	return (i >> I23(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I23) SetBit(index int, v I23) {
	if index < 0 {
		index = 23 + index
	}
	bit := I23(1) << I23(index)
	*i = I23.cast((*i &^ bit) | ((v & 1) << I23(index)))
}

// Bitref returns a Range for bit index.
func (i *I23) Bitref(index int) Range[I23] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I23) Bits(lo, hi int) I23 {
	p := 23
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I23((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I23(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I23) SetBits(lo, hi int, v I23) {
	p := 23
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I23((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I23.cast((*i &^ mask) | ((v << I23(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I23) Bitsref(lo, hi int) Range[I23] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I23) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I23) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I23(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I23) Byteref(idx int) Range[I23] { return byte(i, idx) }

// U24 is an 24-bit unsigned integer.
type U24 uint32

// AsU24 returns a U24 representing v reduced to 24 bits.
func AsU24[T Unsigned](v T) U24 { return U24.cast(U24(v)) }

func (U24) nbits() int  { return 24 }
func (u U24) mask() U24 { return 0xffffff }
func (u U24) cast() U24 { return u & u.mask() }

// Set assigns v reduced to 24 bits to u.
func (u *U24) Set(v U24) U24 { *u = new(U24(v)).cast(); return U24.cast(v) }

// Add returns u+o reduced to 24 bits.
func (u U24) Add(o U24) U24 { return U24.cast(u + o) }

// Sub returns u-o reduced to 24 bits.
func (u U24) Sub(o U24) U24 { return U24.cast(u - o) }

// Inc returns u+1 reduced to 24 bits.
func (u U24) Inc() U24 { return U24.cast(u + 1) }

// Dec returns u-1 reduced to 24 bits.
func (u U24) Dec() U24 { return U24.cast(u - 1) }

// Mul returns u*o reduced to 24 bits.
func (u U24) Mul(o U24) U24 { return U24.cast(u * o) }

// Div returns u divided by o.
func (u U24) Div(o U24) U24 { return U24.cast(u / o) }

// Mod returns u modulo o.
func (u U24) Mod(o U24) U24 { return U24.cast(u % o) }

// And returns u&o reduced to 24 bits.
func (u U24) And(o U24) U24 { return U24.cast(u & o) }

// Or returns u|o reduced to 24 bits.
func (u U24) Or(o U24) U24 { return U24.cast(u | o) }

// Xor returns u^o reduced to 24 bits.
func (u U24) Xor(o U24) U24 { return U24.cast(u ^ o) }

// Shr returns u>>o reduced to 24 bits.
func (u U24) Shr(o U24) U24 { return U24.cast(u >> o) }

// Shl returns u<<o reduced to 24 bits.
func (u U24) Shl(o U24) U24 { return U24.cast(u << o) }

// Not returns bitwise complement of u reduced to 24 bits.
func (u U24) Not() U24 { return U24.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U24) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U24) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U24 value.
func (u U24) Signed() I24 { return I24.cast(I24(u)) }

// Clamp returns value saturated into the representable range of U24.
func (u U24) Clamp(value uint64) U24 { return uclamp[U24](value) }

// Clip masks u to a bits-wide low field.
func (u U24) Clip(bits int) U24 {
	b := 1 << (bits - 1)
	m := U24(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U24) Bit(index int) U24 {
	if index < 0 {
		index = 24 + index
	}
	return (u >> U24(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U24) SetBit(index int, v U24) {
	if index < 0 {
		index = 24 + index
	}
	bit := U24(1) << U24(index)
	*u = U24.cast((*u &^ bit) | ((v & 1) << U24(index)))
}

// Bitref returns a Range for bit index.
func (u *U24) Bitref(index int) Range[U24] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U24) Bits(lo, hi int) U24 {
	p := 24
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U24((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U24(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U24) SetBits(lo, hi int, v U24) {
	p := 24
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U24((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U24.cast((*u &^ mask) | ((v << U24(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U24) Bitsref(lo, hi int) Range[U24] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U24) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U24) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U24(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U24) Byteref(idx int) Range[U24] { return byte(u, idx) }

// I24 is an 24-bit signed integer in two's complement.
type I24 int32

// AsI24 returns a I24 representing v reduced to 24 bits.
func AsI24[T Signed](v T) I24 { return I24.cast(I24(v)) }

func (I24) nbits() int  { return 24 }
func (i I24) mask() I24 { return 0xffffff }
func (i I24) sign() I24 { return 1 << (i.nbits() - 1) }
func (i I24) cast() I24 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 24 bits to u.
func (i *I24) Set(v I24) I24 { *i = new(I24(v)).cast(); return I24.cast(v) }

// Add returns i+o reduced to 24 bits.
func (i I24) Add(o I24) I24 { return I24.cast(i + o) }

// Sub returns i-o reduced to 24 bits.
func (i I24) Sub(o I24) I24 { return I24.cast(i - o) }

// Inc returns i+1 reduced to 24 bits.
func (i I24) Inc() I24 { return I24.cast(i + 1) }

// Dec returns i-1 reduced to 24 bits.
func (i I24) Dec() I24 { return I24.cast(i - 1) }

// Mul returns i*o reduced to 24 bits.
func (i I24) Mul(o I24) I24 { return I24.cast(i * o) }

// Div returns i divided by o.
func (i I24) Div(o I24) I24 { return I24.cast(i / o) }

// Mod returns i modulo o.
func (i I24) Mod(o I24) I24 { return I24.cast(i % o) }

// And returns i&o reduced to 24 bits.
func (i I24) And(o I24) I24 { return I24.cast(i & o) }

// Or returns i|o reduced to 24 bits.
func (i I24) Or(o I24) I24 { return I24.cast(i | o) }

// Xor returns i^o reduced to 24 bits.
func (i I24) Xor(o I24) I24 { return I24.cast(i ^ o) }

// Shr returns i>>o reduced to 24 bits.
func (i I24) Shr(o I24) I24 { return I24.cast(i >> o) }

// Shl returns i<<o reduced to 24 bits.
func (i I24) Shl(o I24) I24 { return I24.cast(i << o) }

// Not returns bitwise complement of i reduced to 24 bits.
func (i I24) Not() I24 { return I24.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I24) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I24) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U24 bit pattern.
func (i I24) Unsigned() U24 { return U24.cast(U24(i)) }

// Clamp returns value saturated into the representable range of I24.
func (i I24) Clamp(value int64) I24 { return iclamp[I24](value) }

// Clip masks i to a bits-wide low field and sign-extends to 24 bits.
func (i I24) Clip(bits int) I24 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I24(uint64(i)&m^b) - I24(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I24) Bit(index int) I24 {
	if index < 0 {
		index = 24 + index
	}
	return (i >> I24(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I24) SetBit(index int, v I24) {
	if index < 0 {
		index = 24 + index
	}
	bit := I24(1) << I24(index)
	*i = I24.cast((*i &^ bit) | ((v & 1) << I24(index)))
}

// Bitref returns a Range for bit index.
func (i *I24) Bitref(index int) Range[I24] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I24) Bits(lo, hi int) I24 {
	p := 24
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I24((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I24(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I24) SetBits(lo, hi int, v I24) {
	p := 24
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I24((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I24.cast((*i &^ mask) | ((v << I24(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I24) Bitsref(lo, hi int) Range[I24] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I24) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I24) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I24(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I24) Byteref(idx int) Range[I24] { return byte(i, idx) }

// U25 is an 25-bit unsigned integer.
type U25 uint32

// AsU25 returns a U25 representing v reduced to 25 bits.
func AsU25[T Unsigned](v T) U25 { return U25.cast(U25(v)) }

func (U25) nbits() int  { return 25 }
func (u U25) mask() U25 { return 0x1ffffff }
func (u U25) cast() U25 { return u & u.mask() }

// Set assigns v reduced to 25 bits to u.
func (u *U25) Set(v U25) U25 { *u = new(U25(v)).cast(); return U25.cast(v) }

// Add returns u+o reduced to 25 bits.
func (u U25) Add(o U25) U25 { return U25.cast(u + o) }

// Sub returns u-o reduced to 25 bits.
func (u U25) Sub(o U25) U25 { return U25.cast(u - o) }

// Inc returns u+1 reduced to 25 bits.
func (u U25) Inc() U25 { return U25.cast(u + 1) }

// Dec returns u-1 reduced to 25 bits.
func (u U25) Dec() U25 { return U25.cast(u - 1) }

// Mul returns u*o reduced to 25 bits.
func (u U25) Mul(o U25) U25 { return U25.cast(u * o) }

// Div returns u divided by o.
func (u U25) Div(o U25) U25 { return U25.cast(u / o) }

// Mod returns u modulo o.
func (u U25) Mod(o U25) U25 { return U25.cast(u % o) }

// And returns u&o reduced to 25 bits.
func (u U25) And(o U25) U25 { return U25.cast(u & o) }

// Or returns u|o reduced to 25 bits.
func (u U25) Or(o U25) U25 { return U25.cast(u | o) }

// Xor returns u^o reduced to 25 bits.
func (u U25) Xor(o U25) U25 { return U25.cast(u ^ o) }

// Shr returns u>>o reduced to 25 bits.
func (u U25) Shr(o U25) U25 { return U25.cast(u >> o) }

// Shl returns u<<o reduced to 25 bits.
func (u U25) Shl(o U25) U25 { return U25.cast(u << o) }

// Not returns bitwise complement of u reduced to 25 bits.
func (u U25) Not() U25 { return U25.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U25) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U25) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U25 value.
func (u U25) Signed() I25 { return I25.cast(I25(u)) }

// Clamp returns value saturated into the representable range of U25.
func (u U25) Clamp(value uint64) U25 { return uclamp[U25](value) }

// Clip masks u to a bits-wide low field.
func (u U25) Clip(bits int) U25 {
	b := 1 << (bits - 1)
	m := U25(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U25) Bit(index int) U25 {
	if index < 0 {
		index = 25 + index
	}
	return (u >> U25(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U25) SetBit(index int, v U25) {
	if index < 0 {
		index = 25 + index
	}
	bit := U25(1) << U25(index)
	*u = U25.cast((*u &^ bit) | ((v & 1) << U25(index)))
}

// Bitref returns a Range for bit index.
func (u *U25) Bitref(index int) Range[U25] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U25) Bits(lo, hi int) U25 {
	p := 25
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U25((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U25(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U25) SetBits(lo, hi int, v U25) {
	p := 25
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U25((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U25.cast((*u &^ mask) | ((v << U25(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U25) Bitsref(lo, hi int) Range[U25] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U25) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U25) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U25(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U25) Byteref(idx int) Range[U25] { return byte(u, idx) }

// I25 is an 25-bit signed integer in two's complement.
type I25 int32

// AsI25 returns a I25 representing v reduced to 25 bits.
func AsI25[T Signed](v T) I25 { return I25.cast(I25(v)) }

func (I25) nbits() int  { return 25 }
func (i I25) mask() I25 { return 0x1ffffff }
func (i I25) sign() I25 { return 1 << (i.nbits() - 1) }
func (i I25) cast() I25 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 25 bits to u.
func (i *I25) Set(v I25) I25 { *i = new(I25(v)).cast(); return I25.cast(v) }

// Add returns i+o reduced to 25 bits.
func (i I25) Add(o I25) I25 { return I25.cast(i + o) }

// Sub returns i-o reduced to 25 bits.
func (i I25) Sub(o I25) I25 { return I25.cast(i - o) }

// Inc returns i+1 reduced to 25 bits.
func (i I25) Inc() I25 { return I25.cast(i + 1) }

// Dec returns i-1 reduced to 25 bits.
func (i I25) Dec() I25 { return I25.cast(i - 1) }

// Mul returns i*o reduced to 25 bits.
func (i I25) Mul(o I25) I25 { return I25.cast(i * o) }

// Div returns i divided by o.
func (i I25) Div(o I25) I25 { return I25.cast(i / o) }

// Mod returns i modulo o.
func (i I25) Mod(o I25) I25 { return I25.cast(i % o) }

// And returns i&o reduced to 25 bits.
func (i I25) And(o I25) I25 { return I25.cast(i & o) }

// Or returns i|o reduced to 25 bits.
func (i I25) Or(o I25) I25 { return I25.cast(i | o) }

// Xor returns i^o reduced to 25 bits.
func (i I25) Xor(o I25) I25 { return I25.cast(i ^ o) }

// Shr returns i>>o reduced to 25 bits.
func (i I25) Shr(o I25) I25 { return I25.cast(i >> o) }

// Shl returns i<<o reduced to 25 bits.
func (i I25) Shl(o I25) I25 { return I25.cast(i << o) }

// Not returns bitwise complement of i reduced to 25 bits.
func (i I25) Not() I25 { return I25.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I25) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I25) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U25 bit pattern.
func (i I25) Unsigned() U25 { return U25.cast(U25(i)) }

// Clamp returns value saturated into the representable range of I25.
func (i I25) Clamp(value int64) I25 { return iclamp[I25](value) }

// Clip masks i to a bits-wide low field and sign-extends to 25 bits.
func (i I25) Clip(bits int) I25 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I25(uint64(i)&m^b) - I25(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I25) Bit(index int) I25 {
	if index < 0 {
		index = 25 + index
	}
	return (i >> I25(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I25) SetBit(index int, v I25) {
	if index < 0 {
		index = 25 + index
	}
	bit := I25(1) << I25(index)
	*i = I25.cast((*i &^ bit) | ((v & 1) << I25(index)))
}

// Bitref returns a Range for bit index.
func (i *I25) Bitref(index int) Range[I25] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I25) Bits(lo, hi int) I25 {
	p := 25
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I25((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I25(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I25) SetBits(lo, hi int, v I25) {
	p := 25
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I25((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I25.cast((*i &^ mask) | ((v << I25(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I25) Bitsref(lo, hi int) Range[I25] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I25) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I25) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I25(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I25) Byteref(idx int) Range[I25] { return byte(i, idx) }

// U26 is an 26-bit unsigned integer.
type U26 uint32

// AsU26 returns a U26 representing v reduced to 26 bits.
func AsU26[T Unsigned](v T) U26 { return U26.cast(U26(v)) }

func (U26) nbits() int  { return 26 }
func (u U26) mask() U26 { return 0x3ffffff }
func (u U26) cast() U26 { return u & u.mask() }

// Set assigns v reduced to 26 bits to u.
func (u *U26) Set(v U26) U26 { *u = new(U26(v)).cast(); return U26.cast(v) }

// Add returns u+o reduced to 26 bits.
func (u U26) Add(o U26) U26 { return U26.cast(u + o) }

// Sub returns u-o reduced to 26 bits.
func (u U26) Sub(o U26) U26 { return U26.cast(u - o) }

// Inc returns u+1 reduced to 26 bits.
func (u U26) Inc() U26 { return U26.cast(u + 1) }

// Dec returns u-1 reduced to 26 bits.
func (u U26) Dec() U26 { return U26.cast(u - 1) }

// Mul returns u*o reduced to 26 bits.
func (u U26) Mul(o U26) U26 { return U26.cast(u * o) }

// Div returns u divided by o.
func (u U26) Div(o U26) U26 { return U26.cast(u / o) }

// Mod returns u modulo o.
func (u U26) Mod(o U26) U26 { return U26.cast(u % o) }

// And returns u&o reduced to 26 bits.
func (u U26) And(o U26) U26 { return U26.cast(u & o) }

// Or returns u|o reduced to 26 bits.
func (u U26) Or(o U26) U26 { return U26.cast(u | o) }

// Xor returns u^o reduced to 26 bits.
func (u U26) Xor(o U26) U26 { return U26.cast(u ^ o) }

// Shr returns u>>o reduced to 26 bits.
func (u U26) Shr(o U26) U26 { return U26.cast(u >> o) }

// Shl returns u<<o reduced to 26 bits.
func (u U26) Shl(o U26) U26 { return U26.cast(u << o) }

// Not returns bitwise complement of u reduced to 26 bits.
func (u U26) Not() U26 { return U26.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U26) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U26) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U26 value.
func (u U26) Signed() I26 { return I26.cast(I26(u)) }

// Clamp returns value saturated into the representable range of U26.
func (u U26) Clamp(value uint64) U26 { return uclamp[U26](value) }

// Clip masks u to a bits-wide low field.
func (u U26) Clip(bits int) U26 {
	b := 1 << (bits - 1)
	m := U26(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U26) Bit(index int) U26 {
	if index < 0 {
		index = 26 + index
	}
	return (u >> U26(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U26) SetBit(index int, v U26) {
	if index < 0 {
		index = 26 + index
	}
	bit := U26(1) << U26(index)
	*u = U26.cast((*u &^ bit) | ((v & 1) << U26(index)))
}

// Bitref returns a Range for bit index.
func (u *U26) Bitref(index int) Range[U26] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U26) Bits(lo, hi int) U26 {
	p := 26
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U26((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U26(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U26) SetBits(lo, hi int, v U26) {
	p := 26
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U26((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U26.cast((*u &^ mask) | ((v << U26(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U26) Bitsref(lo, hi int) Range[U26] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U26) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U26) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U26(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U26) Byteref(idx int) Range[U26] { return byte(u, idx) }

// I26 is an 26-bit signed integer in two's complement.
type I26 int32

// AsI26 returns a I26 representing v reduced to 26 bits.
func AsI26[T Signed](v T) I26 { return I26.cast(I26(v)) }

func (I26) nbits() int  { return 26 }
func (i I26) mask() I26 { return 0x3ffffff }
func (i I26) sign() I26 { return 1 << (i.nbits() - 1) }
func (i I26) cast() I26 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 26 bits to u.
func (i *I26) Set(v I26) I26 { *i = new(I26(v)).cast(); return I26.cast(v) }

// Add returns i+o reduced to 26 bits.
func (i I26) Add(o I26) I26 { return I26.cast(i + o) }

// Sub returns i-o reduced to 26 bits.
func (i I26) Sub(o I26) I26 { return I26.cast(i - o) }

// Inc returns i+1 reduced to 26 bits.
func (i I26) Inc() I26 { return I26.cast(i + 1) }

// Dec returns i-1 reduced to 26 bits.
func (i I26) Dec() I26 { return I26.cast(i - 1) }

// Mul returns i*o reduced to 26 bits.
func (i I26) Mul(o I26) I26 { return I26.cast(i * o) }

// Div returns i divided by o.
func (i I26) Div(o I26) I26 { return I26.cast(i / o) }

// Mod returns i modulo o.
func (i I26) Mod(o I26) I26 { return I26.cast(i % o) }

// And returns i&o reduced to 26 bits.
func (i I26) And(o I26) I26 { return I26.cast(i & o) }

// Or returns i|o reduced to 26 bits.
func (i I26) Or(o I26) I26 { return I26.cast(i | o) }

// Xor returns i^o reduced to 26 bits.
func (i I26) Xor(o I26) I26 { return I26.cast(i ^ o) }

// Shr returns i>>o reduced to 26 bits.
func (i I26) Shr(o I26) I26 { return I26.cast(i >> o) }

// Shl returns i<<o reduced to 26 bits.
func (i I26) Shl(o I26) I26 { return I26.cast(i << o) }

// Not returns bitwise complement of i reduced to 26 bits.
func (i I26) Not() I26 { return I26.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I26) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I26) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U26 bit pattern.
func (i I26) Unsigned() U26 { return U26.cast(U26(i)) }

// Clamp returns value saturated into the representable range of I26.
func (i I26) Clamp(value int64) I26 { return iclamp[I26](value) }

// Clip masks i to a bits-wide low field and sign-extends to 26 bits.
func (i I26) Clip(bits int) I26 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I26(uint64(i)&m^b) - I26(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I26) Bit(index int) I26 {
	if index < 0 {
		index = 26 + index
	}
	return (i >> I26(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I26) SetBit(index int, v I26) {
	if index < 0 {
		index = 26 + index
	}
	bit := I26(1) << I26(index)
	*i = I26.cast((*i &^ bit) | ((v & 1) << I26(index)))
}

// Bitref returns a Range for bit index.
func (i *I26) Bitref(index int) Range[I26] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I26) Bits(lo, hi int) I26 {
	p := 26
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I26((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I26(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I26) SetBits(lo, hi int, v I26) {
	p := 26
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I26((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I26.cast((*i &^ mask) | ((v << I26(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I26) Bitsref(lo, hi int) Range[I26] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I26) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I26) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I26(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I26) Byteref(idx int) Range[I26] { return byte(i, idx) }

// U27 is an 27-bit unsigned integer.
type U27 uint32

// AsU27 returns a U27 representing v reduced to 27 bits.
func AsU27[T Unsigned](v T) U27 { return U27.cast(U27(v)) }

func (U27) nbits() int  { return 27 }
func (u U27) mask() U27 { return 0x7ffffff }
func (u U27) cast() U27 { return u & u.mask() }

// Set assigns v reduced to 27 bits to u.
func (u *U27) Set(v U27) U27 { *u = new(U27(v)).cast(); return U27.cast(v) }

// Add returns u+o reduced to 27 bits.
func (u U27) Add(o U27) U27 { return U27.cast(u + o) }

// Sub returns u-o reduced to 27 bits.
func (u U27) Sub(o U27) U27 { return U27.cast(u - o) }

// Inc returns u+1 reduced to 27 bits.
func (u U27) Inc() U27 { return U27.cast(u + 1) }

// Dec returns u-1 reduced to 27 bits.
func (u U27) Dec() U27 { return U27.cast(u - 1) }

// Mul returns u*o reduced to 27 bits.
func (u U27) Mul(o U27) U27 { return U27.cast(u * o) }

// Div returns u divided by o.
func (u U27) Div(o U27) U27 { return U27.cast(u / o) }

// Mod returns u modulo o.
func (u U27) Mod(o U27) U27 { return U27.cast(u % o) }

// And returns u&o reduced to 27 bits.
func (u U27) And(o U27) U27 { return U27.cast(u & o) }

// Or returns u|o reduced to 27 bits.
func (u U27) Or(o U27) U27 { return U27.cast(u | o) }

// Xor returns u^o reduced to 27 bits.
func (u U27) Xor(o U27) U27 { return U27.cast(u ^ o) }

// Shr returns u>>o reduced to 27 bits.
func (u U27) Shr(o U27) U27 { return U27.cast(u >> o) }

// Shl returns u<<o reduced to 27 bits.
func (u U27) Shl(o U27) U27 { return U27.cast(u << o) }

// Not returns bitwise complement of u reduced to 27 bits.
func (u U27) Not() U27 { return U27.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U27) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U27) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U27 value.
func (u U27) Signed() I27 { return I27.cast(I27(u)) }

// Clamp returns value saturated into the representable range of U27.
func (u U27) Clamp(value uint64) U27 { return uclamp[U27](value) }

// Clip masks u to a bits-wide low field.
func (u U27) Clip(bits int) U27 {
	b := 1 << (bits - 1)
	m := U27(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U27) Bit(index int) U27 {
	if index < 0 {
		index = 27 + index
	}
	return (u >> U27(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U27) SetBit(index int, v U27) {
	if index < 0 {
		index = 27 + index
	}
	bit := U27(1) << U27(index)
	*u = U27.cast((*u &^ bit) | ((v & 1) << U27(index)))
}

// Bitref returns a Range for bit index.
func (u *U27) Bitref(index int) Range[U27] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U27) Bits(lo, hi int) U27 {
	p := 27
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U27((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U27(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U27) SetBits(lo, hi int, v U27) {
	p := 27
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U27((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U27.cast((*u &^ mask) | ((v << U27(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U27) Bitsref(lo, hi int) Range[U27] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U27) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U27) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U27(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U27) Byteref(idx int) Range[U27] { return byte(u, idx) }

// I27 is an 27-bit signed integer in two's complement.
type I27 int32

// AsI27 returns a I27 representing v reduced to 27 bits.
func AsI27[T Signed](v T) I27 { return I27.cast(I27(v)) }

func (I27) nbits() int  { return 27 }
func (i I27) mask() I27 { return 0x7ffffff }
func (i I27) sign() I27 { return 1 << (i.nbits() - 1) }
func (i I27) cast() I27 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 27 bits to u.
func (i *I27) Set(v I27) I27 { *i = new(I27(v)).cast(); return I27.cast(v) }

// Add returns i+o reduced to 27 bits.
func (i I27) Add(o I27) I27 { return I27.cast(i + o) }

// Sub returns i-o reduced to 27 bits.
func (i I27) Sub(o I27) I27 { return I27.cast(i - o) }

// Inc returns i+1 reduced to 27 bits.
func (i I27) Inc() I27 { return I27.cast(i + 1) }

// Dec returns i-1 reduced to 27 bits.
func (i I27) Dec() I27 { return I27.cast(i - 1) }

// Mul returns i*o reduced to 27 bits.
func (i I27) Mul(o I27) I27 { return I27.cast(i * o) }

// Div returns i divided by o.
func (i I27) Div(o I27) I27 { return I27.cast(i / o) }

// Mod returns i modulo o.
func (i I27) Mod(o I27) I27 { return I27.cast(i % o) }

// And returns i&o reduced to 27 bits.
func (i I27) And(o I27) I27 { return I27.cast(i & o) }

// Or returns i|o reduced to 27 bits.
func (i I27) Or(o I27) I27 { return I27.cast(i | o) }

// Xor returns i^o reduced to 27 bits.
func (i I27) Xor(o I27) I27 { return I27.cast(i ^ o) }

// Shr returns i>>o reduced to 27 bits.
func (i I27) Shr(o I27) I27 { return I27.cast(i >> o) }

// Shl returns i<<o reduced to 27 bits.
func (i I27) Shl(o I27) I27 { return I27.cast(i << o) }

// Not returns bitwise complement of i reduced to 27 bits.
func (i I27) Not() I27 { return I27.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I27) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I27) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U27 bit pattern.
func (i I27) Unsigned() U27 { return U27.cast(U27(i)) }

// Clamp returns value saturated into the representable range of I27.
func (i I27) Clamp(value int64) I27 { return iclamp[I27](value) }

// Clip masks i to a bits-wide low field and sign-extends to 27 bits.
func (i I27) Clip(bits int) I27 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I27(uint64(i)&m^b) - I27(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I27) Bit(index int) I27 {
	if index < 0 {
		index = 27 + index
	}
	return (i >> I27(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I27) SetBit(index int, v I27) {
	if index < 0 {
		index = 27 + index
	}
	bit := I27(1) << I27(index)
	*i = I27.cast((*i &^ bit) | ((v & 1) << I27(index)))
}

// Bitref returns a Range for bit index.
func (i *I27) Bitref(index int) Range[I27] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I27) Bits(lo, hi int) I27 {
	p := 27
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I27((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I27(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I27) SetBits(lo, hi int, v I27) {
	p := 27
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I27((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I27.cast((*i &^ mask) | ((v << I27(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I27) Bitsref(lo, hi int) Range[I27] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I27) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I27) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I27(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I27) Byteref(idx int) Range[I27] { return byte(i, idx) }

// U28 is an 28-bit unsigned integer.
type U28 uint32

// AsU28 returns a U28 representing v reduced to 28 bits.
func AsU28[T Unsigned](v T) U28 { return U28.cast(U28(v)) }

func (U28) nbits() int  { return 28 }
func (u U28) mask() U28 { return 0xfffffff }
func (u U28) cast() U28 { return u & u.mask() }

// Set assigns v reduced to 28 bits to u.
func (u *U28) Set(v U28) U28 { *u = new(U28(v)).cast(); return U28.cast(v) }

// Add returns u+o reduced to 28 bits.
func (u U28) Add(o U28) U28 { return U28.cast(u + o) }

// Sub returns u-o reduced to 28 bits.
func (u U28) Sub(o U28) U28 { return U28.cast(u - o) }

// Inc returns u+1 reduced to 28 bits.
func (u U28) Inc() U28 { return U28.cast(u + 1) }

// Dec returns u-1 reduced to 28 bits.
func (u U28) Dec() U28 { return U28.cast(u - 1) }

// Mul returns u*o reduced to 28 bits.
func (u U28) Mul(o U28) U28 { return U28.cast(u * o) }

// Div returns u divided by o.
func (u U28) Div(o U28) U28 { return U28.cast(u / o) }

// Mod returns u modulo o.
func (u U28) Mod(o U28) U28 { return U28.cast(u % o) }

// And returns u&o reduced to 28 bits.
func (u U28) And(o U28) U28 { return U28.cast(u & o) }

// Or returns u|o reduced to 28 bits.
func (u U28) Or(o U28) U28 { return U28.cast(u | o) }

// Xor returns u^o reduced to 28 bits.
func (u U28) Xor(o U28) U28 { return U28.cast(u ^ o) }

// Shr returns u>>o reduced to 28 bits.
func (u U28) Shr(o U28) U28 { return U28.cast(u >> o) }

// Shl returns u<<o reduced to 28 bits.
func (u U28) Shl(o U28) U28 { return U28.cast(u << o) }

// Not returns bitwise complement of u reduced to 28 bits.
func (u U28) Not() U28 { return U28.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U28) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U28) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U28 value.
func (u U28) Signed() I28 { return I28.cast(I28(u)) }

// Clamp returns value saturated into the representable range of U28.
func (u U28) Clamp(value uint64) U28 { return uclamp[U28](value) }

// Clip masks u to a bits-wide low field.
func (u U28) Clip(bits int) U28 {
	b := 1 << (bits - 1)
	m := U28(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U28) Bit(index int) U28 {
	if index < 0 {
		index = 28 + index
	}
	return (u >> U28(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U28) SetBit(index int, v U28) {
	if index < 0 {
		index = 28 + index
	}
	bit := U28(1) << U28(index)
	*u = U28.cast((*u &^ bit) | ((v & 1) << U28(index)))
}

// Bitref returns a Range for bit index.
func (u *U28) Bitref(index int) Range[U28] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U28) Bits(lo, hi int) U28 {
	p := 28
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U28((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U28(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U28) SetBits(lo, hi int, v U28) {
	p := 28
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U28((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U28.cast((*u &^ mask) | ((v << U28(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U28) Bitsref(lo, hi int) Range[U28] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U28) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U28) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U28(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U28) Byteref(idx int) Range[U28] { return byte(u, idx) }

// I28 is an 28-bit signed integer in two's complement.
type I28 int32

// AsI28 returns a I28 representing v reduced to 28 bits.
func AsI28[T Signed](v T) I28 { return I28.cast(I28(v)) }

func (I28) nbits() int  { return 28 }
func (i I28) mask() I28 { return 0xfffffff }
func (i I28) sign() I28 { return 1 << (i.nbits() - 1) }
func (i I28) cast() I28 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 28 bits to u.
func (i *I28) Set(v I28) I28 { *i = new(I28(v)).cast(); return I28.cast(v) }

// Add returns i+o reduced to 28 bits.
func (i I28) Add(o I28) I28 { return I28.cast(i + o) }

// Sub returns i-o reduced to 28 bits.
func (i I28) Sub(o I28) I28 { return I28.cast(i - o) }

// Inc returns i+1 reduced to 28 bits.
func (i I28) Inc() I28 { return I28.cast(i + 1) }

// Dec returns i-1 reduced to 28 bits.
func (i I28) Dec() I28 { return I28.cast(i - 1) }

// Mul returns i*o reduced to 28 bits.
func (i I28) Mul(o I28) I28 { return I28.cast(i * o) }

// Div returns i divided by o.
func (i I28) Div(o I28) I28 { return I28.cast(i / o) }

// Mod returns i modulo o.
func (i I28) Mod(o I28) I28 { return I28.cast(i % o) }

// And returns i&o reduced to 28 bits.
func (i I28) And(o I28) I28 { return I28.cast(i & o) }

// Or returns i|o reduced to 28 bits.
func (i I28) Or(o I28) I28 { return I28.cast(i | o) }

// Xor returns i^o reduced to 28 bits.
func (i I28) Xor(o I28) I28 { return I28.cast(i ^ o) }

// Shr returns i>>o reduced to 28 bits.
func (i I28) Shr(o I28) I28 { return I28.cast(i >> o) }

// Shl returns i<<o reduced to 28 bits.
func (i I28) Shl(o I28) I28 { return I28.cast(i << o) }

// Not returns bitwise complement of i reduced to 28 bits.
func (i I28) Not() I28 { return I28.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I28) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I28) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U28 bit pattern.
func (i I28) Unsigned() U28 { return U28.cast(U28(i)) }

// Clamp returns value saturated into the representable range of I28.
func (i I28) Clamp(value int64) I28 { return iclamp[I28](value) }

// Clip masks i to a bits-wide low field and sign-extends to 28 bits.
func (i I28) Clip(bits int) I28 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I28(uint64(i)&m^b) - I28(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I28) Bit(index int) I28 {
	if index < 0 {
		index = 28 + index
	}
	return (i >> I28(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I28) SetBit(index int, v I28) {
	if index < 0 {
		index = 28 + index
	}
	bit := I28(1) << I28(index)
	*i = I28.cast((*i &^ bit) | ((v & 1) << I28(index)))
}

// Bitref returns a Range for bit index.
func (i *I28) Bitref(index int) Range[I28] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I28) Bits(lo, hi int) I28 {
	p := 28
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I28((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I28(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I28) SetBits(lo, hi int, v I28) {
	p := 28
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I28((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I28.cast((*i &^ mask) | ((v << I28(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I28) Bitsref(lo, hi int) Range[I28] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I28) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I28) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I28(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I28) Byteref(idx int) Range[I28] { return byte(i, idx) }

// U29 is an 29-bit unsigned integer.
type U29 uint32

// AsU29 returns a U29 representing v reduced to 29 bits.
func AsU29[T Unsigned](v T) U29 { return U29.cast(U29(v)) }

func (U29) nbits() int  { return 29 }
func (u U29) mask() U29 { return 0x1fffffff }
func (u U29) cast() U29 { return u & u.mask() }

// Set assigns v reduced to 29 bits to u.
func (u *U29) Set(v U29) U29 { *u = new(U29(v)).cast(); return U29.cast(v) }

// Add returns u+o reduced to 29 bits.
func (u U29) Add(o U29) U29 { return U29.cast(u + o) }

// Sub returns u-o reduced to 29 bits.
func (u U29) Sub(o U29) U29 { return U29.cast(u - o) }

// Inc returns u+1 reduced to 29 bits.
func (u U29) Inc() U29 { return U29.cast(u + 1) }

// Dec returns u-1 reduced to 29 bits.
func (u U29) Dec() U29 { return U29.cast(u - 1) }

// Mul returns u*o reduced to 29 bits.
func (u U29) Mul(o U29) U29 { return U29.cast(u * o) }

// Div returns u divided by o.
func (u U29) Div(o U29) U29 { return U29.cast(u / o) }

// Mod returns u modulo o.
func (u U29) Mod(o U29) U29 { return U29.cast(u % o) }

// And returns u&o reduced to 29 bits.
func (u U29) And(o U29) U29 { return U29.cast(u & o) }

// Or returns u|o reduced to 29 bits.
func (u U29) Or(o U29) U29 { return U29.cast(u | o) }

// Xor returns u^o reduced to 29 bits.
func (u U29) Xor(o U29) U29 { return U29.cast(u ^ o) }

// Shr returns u>>o reduced to 29 bits.
func (u U29) Shr(o U29) U29 { return U29.cast(u >> o) }

// Shl returns u<<o reduced to 29 bits.
func (u U29) Shl(o U29) U29 { return U29.cast(u << o) }

// Not returns bitwise complement of u reduced to 29 bits.
func (u U29) Not() U29 { return U29.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U29) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U29) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U29 value.
func (u U29) Signed() I29 { return I29.cast(I29(u)) }

// Clamp returns value saturated into the representable range of U29.
func (u U29) Clamp(value uint64) U29 { return uclamp[U29](value) }

// Clip masks u to a bits-wide low field.
func (u U29) Clip(bits int) U29 {
	b := 1 << (bits - 1)
	m := U29(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U29) Bit(index int) U29 {
	if index < 0 {
		index = 29 + index
	}
	return (u >> U29(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U29) SetBit(index int, v U29) {
	if index < 0 {
		index = 29 + index
	}
	bit := U29(1) << U29(index)
	*u = U29.cast((*u &^ bit) | ((v & 1) << U29(index)))
}

// Bitref returns a Range for bit index.
func (u *U29) Bitref(index int) Range[U29] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U29) Bits(lo, hi int) U29 {
	p := 29
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U29((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U29(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U29) SetBits(lo, hi int, v U29) {
	p := 29
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U29((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U29.cast((*u &^ mask) | ((v << U29(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U29) Bitsref(lo, hi int) Range[U29] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U29) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U29) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U29(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U29) Byteref(idx int) Range[U29] { return byte(u, idx) }

// I29 is an 29-bit signed integer in two's complement.
type I29 int32

// AsI29 returns a I29 representing v reduced to 29 bits.
func AsI29[T Signed](v T) I29 { return I29.cast(I29(v)) }

func (I29) nbits() int  { return 29 }
func (i I29) mask() I29 { return 0x1fffffff }
func (i I29) sign() I29 { return 1 << (i.nbits() - 1) }
func (i I29) cast() I29 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 29 bits to u.
func (i *I29) Set(v I29) I29 { *i = new(I29(v)).cast(); return I29.cast(v) }

// Add returns i+o reduced to 29 bits.
func (i I29) Add(o I29) I29 { return I29.cast(i + o) }

// Sub returns i-o reduced to 29 bits.
func (i I29) Sub(o I29) I29 { return I29.cast(i - o) }

// Inc returns i+1 reduced to 29 bits.
func (i I29) Inc() I29 { return I29.cast(i + 1) }

// Dec returns i-1 reduced to 29 bits.
func (i I29) Dec() I29 { return I29.cast(i - 1) }

// Mul returns i*o reduced to 29 bits.
func (i I29) Mul(o I29) I29 { return I29.cast(i * o) }

// Div returns i divided by o.
func (i I29) Div(o I29) I29 { return I29.cast(i / o) }

// Mod returns i modulo o.
func (i I29) Mod(o I29) I29 { return I29.cast(i % o) }

// And returns i&o reduced to 29 bits.
func (i I29) And(o I29) I29 { return I29.cast(i & o) }

// Or returns i|o reduced to 29 bits.
func (i I29) Or(o I29) I29 { return I29.cast(i | o) }

// Xor returns i^o reduced to 29 bits.
func (i I29) Xor(o I29) I29 { return I29.cast(i ^ o) }

// Shr returns i>>o reduced to 29 bits.
func (i I29) Shr(o I29) I29 { return I29.cast(i >> o) }

// Shl returns i<<o reduced to 29 bits.
func (i I29) Shl(o I29) I29 { return I29.cast(i << o) }

// Not returns bitwise complement of i reduced to 29 bits.
func (i I29) Not() I29 { return I29.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I29) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I29) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U29 bit pattern.
func (i I29) Unsigned() U29 { return U29.cast(U29(i)) }

// Clamp returns value saturated into the representable range of I29.
func (i I29) Clamp(value int64) I29 { return iclamp[I29](value) }

// Clip masks i to a bits-wide low field and sign-extends to 29 bits.
func (i I29) Clip(bits int) I29 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I29(uint64(i)&m^b) - I29(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I29) Bit(index int) I29 {
	if index < 0 {
		index = 29 + index
	}
	return (i >> I29(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I29) SetBit(index int, v I29) {
	if index < 0 {
		index = 29 + index
	}
	bit := I29(1) << I29(index)
	*i = I29.cast((*i &^ bit) | ((v & 1) << I29(index)))
}

// Bitref returns a Range for bit index.
func (i *I29) Bitref(index int) Range[I29] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I29) Bits(lo, hi int) I29 {
	p := 29
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I29((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I29(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I29) SetBits(lo, hi int, v I29) {
	p := 29
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I29((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I29.cast((*i &^ mask) | ((v << I29(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I29) Bitsref(lo, hi int) Range[I29] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I29) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I29) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I29(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I29) Byteref(idx int) Range[I29] { return byte(i, idx) }

// U30 is an 30-bit unsigned integer.
type U30 uint32

// AsU30 returns a U30 representing v reduced to 30 bits.
func AsU30[T Unsigned](v T) U30 { return U30.cast(U30(v)) }

func (U30) nbits() int  { return 30 }
func (u U30) mask() U30 { return 0x3fffffff }
func (u U30) cast() U30 { return u & u.mask() }

// Set assigns v reduced to 30 bits to u.
func (u *U30) Set(v U30) U30 { *u = new(U30(v)).cast(); return U30.cast(v) }

// Add returns u+o reduced to 30 bits.
func (u U30) Add(o U30) U30 { return U30.cast(u + o) }

// Sub returns u-o reduced to 30 bits.
func (u U30) Sub(o U30) U30 { return U30.cast(u - o) }

// Inc returns u+1 reduced to 30 bits.
func (u U30) Inc() U30 { return U30.cast(u + 1) }

// Dec returns u-1 reduced to 30 bits.
func (u U30) Dec() U30 { return U30.cast(u - 1) }

// Mul returns u*o reduced to 30 bits.
func (u U30) Mul(o U30) U30 { return U30.cast(u * o) }

// Div returns u divided by o.
func (u U30) Div(o U30) U30 { return U30.cast(u / o) }

// Mod returns u modulo o.
func (u U30) Mod(o U30) U30 { return U30.cast(u % o) }

// And returns u&o reduced to 30 bits.
func (u U30) And(o U30) U30 { return U30.cast(u & o) }

// Or returns u|o reduced to 30 bits.
func (u U30) Or(o U30) U30 { return U30.cast(u | o) }

// Xor returns u^o reduced to 30 bits.
func (u U30) Xor(o U30) U30 { return U30.cast(u ^ o) }

// Shr returns u>>o reduced to 30 bits.
func (u U30) Shr(o U30) U30 { return U30.cast(u >> o) }

// Shl returns u<<o reduced to 30 bits.
func (u U30) Shl(o U30) U30 { return U30.cast(u << o) }

// Not returns bitwise complement of u reduced to 30 bits.
func (u U30) Not() U30 { return U30.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U30) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U30) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U30 value.
func (u U30) Signed() I30 { return I30.cast(I30(u)) }

// Clamp returns value saturated into the representable range of U30.
func (u U30) Clamp(value uint64) U30 { return uclamp[U30](value) }

// Clip masks u to a bits-wide low field.
func (u U30) Clip(bits int) U30 {
	b := 1 << (bits - 1)
	m := U30(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U30) Bit(index int) U30 {
	if index < 0 {
		index = 30 + index
	}
	return (u >> U30(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U30) SetBit(index int, v U30) {
	if index < 0 {
		index = 30 + index
	}
	bit := U30(1) << U30(index)
	*u = U30.cast((*u &^ bit) | ((v & 1) << U30(index)))
}

// Bitref returns a Range for bit index.
func (u *U30) Bitref(index int) Range[U30] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U30) Bits(lo, hi int) U30 {
	p := 30
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U30((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U30(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U30) SetBits(lo, hi int, v U30) {
	p := 30
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U30((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U30.cast((*u &^ mask) | ((v << U30(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U30) Bitsref(lo, hi int) Range[U30] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U30) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U30) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U30(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U30) Byteref(idx int) Range[U30] { return byte(u, idx) }

// I30 is an 30-bit signed integer in two's complement.
type I30 int32

// AsI30 returns a I30 representing v reduced to 30 bits.
func AsI30[T Signed](v T) I30 { return I30.cast(I30(v)) }

func (I30) nbits() int  { return 30 }
func (i I30) mask() I30 { return 0x3fffffff }
func (i I30) sign() I30 { return 1 << (i.nbits() - 1) }
func (i I30) cast() I30 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 30 bits to u.
func (i *I30) Set(v I30) I30 { *i = new(I30(v)).cast(); return I30.cast(v) }

// Add returns i+o reduced to 30 bits.
func (i I30) Add(o I30) I30 { return I30.cast(i + o) }

// Sub returns i-o reduced to 30 bits.
func (i I30) Sub(o I30) I30 { return I30.cast(i - o) }

// Inc returns i+1 reduced to 30 bits.
func (i I30) Inc() I30 { return I30.cast(i + 1) }

// Dec returns i-1 reduced to 30 bits.
func (i I30) Dec() I30 { return I30.cast(i - 1) }

// Mul returns i*o reduced to 30 bits.
func (i I30) Mul(o I30) I30 { return I30.cast(i * o) }

// Div returns i divided by o.
func (i I30) Div(o I30) I30 { return I30.cast(i / o) }

// Mod returns i modulo o.
func (i I30) Mod(o I30) I30 { return I30.cast(i % o) }

// And returns i&o reduced to 30 bits.
func (i I30) And(o I30) I30 { return I30.cast(i & o) }

// Or returns i|o reduced to 30 bits.
func (i I30) Or(o I30) I30 { return I30.cast(i | o) }

// Xor returns i^o reduced to 30 bits.
func (i I30) Xor(o I30) I30 { return I30.cast(i ^ o) }

// Shr returns i>>o reduced to 30 bits.
func (i I30) Shr(o I30) I30 { return I30.cast(i >> o) }

// Shl returns i<<o reduced to 30 bits.
func (i I30) Shl(o I30) I30 { return I30.cast(i << o) }

// Not returns bitwise complement of i reduced to 30 bits.
func (i I30) Not() I30 { return I30.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I30) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I30) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U30 bit pattern.
func (i I30) Unsigned() U30 { return U30.cast(U30(i)) }

// Clamp returns value saturated into the representable range of I30.
func (i I30) Clamp(value int64) I30 { return iclamp[I30](value) }

// Clip masks i to a bits-wide low field and sign-extends to 30 bits.
func (i I30) Clip(bits int) I30 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I30(uint64(i)&m^b) - I30(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I30) Bit(index int) I30 {
	if index < 0 {
		index = 30 + index
	}
	return (i >> I30(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I30) SetBit(index int, v I30) {
	if index < 0 {
		index = 30 + index
	}
	bit := I30(1) << I30(index)
	*i = I30.cast((*i &^ bit) | ((v & 1) << I30(index)))
}

// Bitref returns a Range for bit index.
func (i *I30) Bitref(index int) Range[I30] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I30) Bits(lo, hi int) I30 {
	p := 30
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I30((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I30(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I30) SetBits(lo, hi int, v I30) {
	p := 30
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I30((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I30.cast((*i &^ mask) | ((v << I30(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I30) Bitsref(lo, hi int) Range[I30] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I30) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I30) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I30(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I30) Byteref(idx int) Range[I30] { return byte(i, idx) }

// U31 is an 31-bit unsigned integer.
type U31 uint32

// AsU31 returns a U31 representing v reduced to 31 bits.
func AsU31[T Unsigned](v T) U31 { return U31.cast(U31(v)) }

func (U31) nbits() int  { return 31 }
func (u U31) mask() U31 { return 0x7fffffff }
func (u U31) cast() U31 { return u & u.mask() }

// Set assigns v reduced to 31 bits to u.
func (u *U31) Set(v U31) U31 { *u = new(U31(v)).cast(); return U31.cast(v) }

// Add returns u+o reduced to 31 bits.
func (u U31) Add(o U31) U31 { return U31.cast(u + o) }

// Sub returns u-o reduced to 31 bits.
func (u U31) Sub(o U31) U31 { return U31.cast(u - o) }

// Inc returns u+1 reduced to 31 bits.
func (u U31) Inc() U31 { return U31.cast(u + 1) }

// Dec returns u-1 reduced to 31 bits.
func (u U31) Dec() U31 { return U31.cast(u - 1) }

// Mul returns u*o reduced to 31 bits.
func (u U31) Mul(o U31) U31 { return U31.cast(u * o) }

// Div returns u divided by o.
func (u U31) Div(o U31) U31 { return U31.cast(u / o) }

// Mod returns u modulo o.
func (u U31) Mod(o U31) U31 { return U31.cast(u % o) }

// And returns u&o reduced to 31 bits.
func (u U31) And(o U31) U31 { return U31.cast(u & o) }

// Or returns u|o reduced to 31 bits.
func (u U31) Or(o U31) U31 { return U31.cast(u | o) }

// Xor returns u^o reduced to 31 bits.
func (u U31) Xor(o U31) U31 { return U31.cast(u ^ o) }

// Shr returns u>>o reduced to 31 bits.
func (u U31) Shr(o U31) U31 { return U31.cast(u >> o) }

// Shl returns u<<o reduced to 31 bits.
func (u U31) Shl(o U31) U31 { return U31.cast(u << o) }

// Not returns bitwise complement of u reduced to 31 bits.
func (u U31) Not() U31 { return U31.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U31) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U31) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U31 value.
func (u U31) Signed() I31 { return I31.cast(I31(u)) }

// Clamp returns value saturated into the representable range of U31.
func (u U31) Clamp(value uint64) U31 { return uclamp[U31](value) }

// Clip masks u to a bits-wide low field.
func (u U31) Clip(bits int) U31 {
	b := 1 << (bits - 1)
	m := U31(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U31) Bit(index int) U31 {
	if index < 0 {
		index = 31 + index
	}
	return (u >> U31(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U31) SetBit(index int, v U31) {
	if index < 0 {
		index = 31 + index
	}
	bit := U31(1) << U31(index)
	*u = U31.cast((*u &^ bit) | ((v & 1) << U31(index)))
}

// Bitref returns a Range for bit index.
func (u *U31) Bitref(index int) Range[U31] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U31) Bits(lo, hi int) U31 {
	p := 31
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U31((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U31(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U31) SetBits(lo, hi int, v U31) {
	p := 31
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U31((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U31.cast((*u &^ mask) | ((v << U31(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U31) Bitsref(lo, hi int) Range[U31] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U31) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U31) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U31(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U31) Byteref(idx int) Range[U31] { return byte(u, idx) }

// I31 is an 31-bit signed integer in two's complement.
type I31 int32

// AsI31 returns a I31 representing v reduced to 31 bits.
func AsI31[T Signed](v T) I31 { return I31.cast(I31(v)) }

func (I31) nbits() int  { return 31 }
func (i I31) mask() I31 { return 0x7fffffff }
func (i I31) sign() I31 { return 1 << (i.nbits() - 1) }
func (i I31) cast() I31 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 31 bits to u.
func (i *I31) Set(v I31) I31 { *i = new(I31(v)).cast(); return I31.cast(v) }

// Add returns i+o reduced to 31 bits.
func (i I31) Add(o I31) I31 { return I31.cast(i + o) }

// Sub returns i-o reduced to 31 bits.
func (i I31) Sub(o I31) I31 { return I31.cast(i - o) }

// Inc returns i+1 reduced to 31 bits.
func (i I31) Inc() I31 { return I31.cast(i + 1) }

// Dec returns i-1 reduced to 31 bits.
func (i I31) Dec() I31 { return I31.cast(i - 1) }

// Mul returns i*o reduced to 31 bits.
func (i I31) Mul(o I31) I31 { return I31.cast(i * o) }

// Div returns i divided by o.
func (i I31) Div(o I31) I31 { return I31.cast(i / o) }

// Mod returns i modulo o.
func (i I31) Mod(o I31) I31 { return I31.cast(i % o) }

// And returns i&o reduced to 31 bits.
func (i I31) And(o I31) I31 { return I31.cast(i & o) }

// Or returns i|o reduced to 31 bits.
func (i I31) Or(o I31) I31 { return I31.cast(i | o) }

// Xor returns i^o reduced to 31 bits.
func (i I31) Xor(o I31) I31 { return I31.cast(i ^ o) }

// Shr returns i>>o reduced to 31 bits.
func (i I31) Shr(o I31) I31 { return I31.cast(i >> o) }

// Shl returns i<<o reduced to 31 bits.
func (i I31) Shl(o I31) I31 { return I31.cast(i << o) }

// Not returns bitwise complement of i reduced to 31 bits.
func (i I31) Not() I31 { return I31.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I31) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I31) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U31 bit pattern.
func (i I31) Unsigned() U31 { return U31.cast(U31(i)) }

// Clamp returns value saturated into the representable range of I31.
func (i I31) Clamp(value int64) I31 { return iclamp[I31](value) }

// Clip masks i to a bits-wide low field and sign-extends to 31 bits.
func (i I31) Clip(bits int) I31 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I31(uint64(i)&m^b) - I31(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I31) Bit(index int) I31 {
	if index < 0 {
		index = 31 + index
	}
	return (i >> I31(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I31) SetBit(index int, v I31) {
	if index < 0 {
		index = 31 + index
	}
	bit := I31(1) << I31(index)
	*i = I31.cast((*i &^ bit) | ((v & 1) << I31(index)))
}

// Bitref returns a Range for bit index.
func (i *I31) Bitref(index int) Range[I31] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I31) Bits(lo, hi int) I31 {
	p := 31
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I31((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I31(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I31) SetBits(lo, hi int, v I31) {
	p := 31
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I31((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I31.cast((*i &^ mask) | ((v << I31(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I31) Bitsref(lo, hi int) Range[I31] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I31) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I31) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I31(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I31) Byteref(idx int) Range[I31] { return byte(i, idx) }

// U32 is an 32-bit unsigned integer.
type U32 uint32

// AsU32 returns a U32 representing v reduced to 32 bits.
func AsU32[T Unsigned](v T) U32 { return U32.cast(U32(v)) }

func (U32) nbits() int  { return 32 }
func (u U32) mask() U32 { return 0xffffffff }
func (u U32) cast() U32 { return u & u.mask() }

// Set assigns v reduced to 32 bits to u.
func (u *U32) Set(v U32) U32 { *u = new(U32(v)).cast(); return U32.cast(v) }

// Add returns u+o reduced to 32 bits.
func (u U32) Add(o U32) U32 { return U32.cast(u + o) }

// Sub returns u-o reduced to 32 bits.
func (u U32) Sub(o U32) U32 { return U32.cast(u - o) }

// Inc returns u+1 reduced to 32 bits.
func (u U32) Inc() U32 { return U32.cast(u + 1) }

// Dec returns u-1 reduced to 32 bits.
func (u U32) Dec() U32 { return U32.cast(u - 1) }

// Mul returns u*o reduced to 32 bits.
func (u U32) Mul(o U32) U32 { return U32.cast(u * o) }

// Div returns u divided by o.
func (u U32) Div(o U32) U32 { return U32.cast(u / o) }

// Mod returns u modulo o.
func (u U32) Mod(o U32) U32 { return U32.cast(u % o) }

// And returns u&o reduced to 32 bits.
func (u U32) And(o U32) U32 { return U32.cast(u & o) }

// Or returns u|o reduced to 32 bits.
func (u U32) Or(o U32) U32 { return U32.cast(u | o) }

// Xor returns u^o reduced to 32 bits.
func (u U32) Xor(o U32) U32 { return U32.cast(u ^ o) }

// Shr returns u>>o reduced to 32 bits.
func (u U32) Shr(o U32) U32 { return U32.cast(u >> o) }

// Shl returns u<<o reduced to 32 bits.
func (u U32) Shl(o U32) U32 { return U32.cast(u << o) }

// Not returns bitwise complement of u reduced to 32 bits.
func (u U32) Not() U32 { return U32.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U32) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U32) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U32 value.
func (u U32) Signed() I32 { return I32.cast(I32(u)) }

// Clamp returns value saturated into the representable range of U32.
func (u U32) Clamp(value uint64) U32 { return uclamp[U32](value) }

// Clip masks u to a bits-wide low field.
func (u U32) Clip(bits int) U32 {
	b := 1 << (bits - 1)
	m := U32(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U32) Bit(index int) U32 {
	if index < 0 {
		index = 32 + index
	}
	return (u >> U32(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U32) SetBit(index int, v U32) {
	if index < 0 {
		index = 32 + index
	}
	bit := U32(1) << U32(index)
	*u = U32.cast((*u &^ bit) | ((v & 1) << U32(index)))
}

// Bitref returns a Range for bit index.
func (u *U32) Bitref(index int) Range[U32] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U32) Bits(lo, hi int) U32 {
	p := 32
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U32((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U32(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U32) SetBits(lo, hi int, v U32) {
	p := 32
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U32((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U32.cast((*u &^ mask) | ((v << U32(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U32) Bitsref(lo, hi int) Range[U32] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U32) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U32) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U32(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U32) Byteref(idx int) Range[U32] { return byte(u, idx) }

// I32 is an 32-bit signed integer in two's complement.
type I32 int32

// AsI32 returns a I32 representing v reduced to 32 bits.
func AsI32[T Signed](v T) I32 { return I32.cast(I32(v)) }

func (I32) nbits() int  { return 32 }
func (i I32) mask() I32 { return -1 }
func (i I32) sign() I32 { return 1 << (i.nbits() - 1) }
func (i I32) cast() I32 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 32 bits to u.
func (i *I32) Set(v I32) I32 { *i = new(I32(v)).cast(); return I32.cast(v) }

// Add returns i+o reduced to 32 bits.
func (i I32) Add(o I32) I32 { return I32.cast(i + o) }

// Sub returns i-o reduced to 32 bits.
func (i I32) Sub(o I32) I32 { return I32.cast(i - o) }

// Inc returns i+1 reduced to 32 bits.
func (i I32) Inc() I32 { return I32.cast(i + 1) }

// Dec returns i-1 reduced to 32 bits.
func (i I32) Dec() I32 { return I32.cast(i - 1) }

// Mul returns i*o reduced to 32 bits.
func (i I32) Mul(o I32) I32 { return I32.cast(i * o) }

// Div returns i divided by o.
func (i I32) Div(o I32) I32 { return I32.cast(i / o) }

// Mod returns i modulo o.
func (i I32) Mod(o I32) I32 { return I32.cast(i % o) }

// And returns i&o reduced to 32 bits.
func (i I32) And(o I32) I32 { return I32.cast(i & o) }

// Or returns i|o reduced to 32 bits.
func (i I32) Or(o I32) I32 { return I32.cast(i | o) }

// Xor returns i^o reduced to 32 bits.
func (i I32) Xor(o I32) I32 { return I32.cast(i ^ o) }

// Shr returns i>>o reduced to 32 bits.
func (i I32) Shr(o I32) I32 { return I32.cast(i >> o) }

// Shl returns i<<o reduced to 32 bits.
func (i I32) Shl(o I32) I32 { return I32.cast(i << o) }

// Not returns bitwise complement of i reduced to 32 bits.
func (i I32) Not() I32 { return I32.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I32) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I32) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U32 bit pattern.
func (i I32) Unsigned() U32 { return U32.cast(U32(i)) }

// Clamp returns value saturated into the representable range of I32.
func (i I32) Clamp(value int64) I32 { return iclamp[I32](value) }

// Clip masks i to a bits-wide low field and sign-extends to 32 bits.
func (i I32) Clip(bits int) I32 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I32(uint64(i)&m^b) - I32(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I32) Bit(index int) I32 {
	if index < 0 {
		index = 32 + index
	}
	return (i >> I32(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I32) SetBit(index int, v I32) {
	if index < 0 {
		index = 32 + index
	}
	bit := I32(1) << I32(index)
	*i = I32.cast((*i &^ bit) | ((v & 1) << I32(index)))
}

// Bitref returns a Range for bit index.
func (i *I32) Bitref(index int) Range[I32] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I32) Bits(lo, hi int) I32 {
	p := 32
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I32((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I32(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I32) SetBits(lo, hi int, v I32) {
	p := 32
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I32((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I32.cast((*i &^ mask) | ((v << I32(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I32) Bitsref(lo, hi int) Range[I32] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I32) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I32) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I32(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I32) Byteref(idx int) Range[I32] { return byte(i, idx) }

// U33 is an 33-bit unsigned integer.
type U33 uint64

// AsU33 returns a U33 representing v reduced to 33 bits.
func AsU33[T Unsigned](v T) U33 { return U33.cast(U33(v)) }

func (U33) nbits() int  { return 33 }
func (u U33) mask() U33 { return 0x1ffffffff }
func (u U33) cast() U33 { return u & u.mask() }

// Set assigns v reduced to 33 bits to u.
func (u *U33) Set(v U33) U33 { *u = new(U33(v)).cast(); return U33.cast(v) }

// Add returns u+o reduced to 33 bits.
func (u U33) Add(o U33) U33 { return U33.cast(u + o) }

// Sub returns u-o reduced to 33 bits.
func (u U33) Sub(o U33) U33 { return U33.cast(u - o) }

// Inc returns u+1 reduced to 33 bits.
func (u U33) Inc() U33 { return U33.cast(u + 1) }

// Dec returns u-1 reduced to 33 bits.
func (u U33) Dec() U33 { return U33.cast(u - 1) }

// Mul returns u*o reduced to 33 bits.
func (u U33) Mul(o U33) U33 { return U33.cast(u * o) }

// Div returns u divided by o.
func (u U33) Div(o U33) U33 { return U33.cast(u / o) }

// Mod returns u modulo o.
func (u U33) Mod(o U33) U33 { return U33.cast(u % o) }

// And returns u&o reduced to 33 bits.
func (u U33) And(o U33) U33 { return U33.cast(u & o) }

// Or returns u|o reduced to 33 bits.
func (u U33) Or(o U33) U33 { return U33.cast(u | o) }

// Xor returns u^o reduced to 33 bits.
func (u U33) Xor(o U33) U33 { return U33.cast(u ^ o) }

// Shr returns u>>o reduced to 33 bits.
func (u U33) Shr(o U33) U33 { return U33.cast(u >> o) }

// Shl returns u<<o reduced to 33 bits.
func (u U33) Shl(o U33) U33 { return U33.cast(u << o) }

// Not returns bitwise complement of u reduced to 33 bits.
func (u U33) Not() U33 { return U33.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U33) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U33) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U33 value.
func (u U33) Signed() I33 { return I33.cast(I33(u)) }

// Clamp returns value saturated into the representable range of U33.
func (u U33) Clamp(value uint64) U33 { return uclamp[U33](value) }

// Clip masks u to a bits-wide low field.
func (u U33) Clip(bits int) U33 {
	b := 1 << (bits - 1)
	m := U33(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U33) Bit(index int) U33 {
	if index < 0 {
		index = 33 + index
	}
	return (u >> U33(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U33) SetBit(index int, v U33) {
	if index < 0 {
		index = 33 + index
	}
	bit := U33(1) << U33(index)
	*u = U33.cast((*u &^ bit) | ((v & 1) << U33(index)))
}

// Bitref returns a Range for bit index.
func (u *U33) Bitref(index int) Range[U33] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U33) Bits(lo, hi int) U33 {
	p := 33
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U33((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U33(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U33) SetBits(lo, hi int, v U33) {
	p := 33
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U33((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U33.cast((*u &^ mask) | ((v << U33(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U33) Bitsref(lo, hi int) Range[U33] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U33) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U33) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U33(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U33) Byteref(idx int) Range[U33] { return byte(u, idx) }

// I33 is an 33-bit signed integer in two's complement.
type I33 int64

// AsI33 returns a I33 representing v reduced to 33 bits.
func AsI33[T Signed](v T) I33 { return I33.cast(I33(v)) }

func (I33) nbits() int  { return 33 }
func (i I33) mask() I33 { return 0x1ffffffff }
func (i I33) sign() I33 { return 1 << (i.nbits() - 1) }
func (i I33) cast() I33 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 33 bits to u.
func (i *I33) Set(v I33) I33 { *i = new(I33(v)).cast(); return I33.cast(v) }

// Add returns i+o reduced to 33 bits.
func (i I33) Add(o I33) I33 { return I33.cast(i + o) }

// Sub returns i-o reduced to 33 bits.
func (i I33) Sub(o I33) I33 { return I33.cast(i - o) }

// Inc returns i+1 reduced to 33 bits.
func (i I33) Inc() I33 { return I33.cast(i + 1) }

// Dec returns i-1 reduced to 33 bits.
func (i I33) Dec() I33 { return I33.cast(i - 1) }

// Mul returns i*o reduced to 33 bits.
func (i I33) Mul(o I33) I33 { return I33.cast(i * o) }

// Div returns i divided by o.
func (i I33) Div(o I33) I33 { return I33.cast(i / o) }

// Mod returns i modulo o.
func (i I33) Mod(o I33) I33 { return I33.cast(i % o) }

// And returns i&o reduced to 33 bits.
func (i I33) And(o I33) I33 { return I33.cast(i & o) }

// Or returns i|o reduced to 33 bits.
func (i I33) Or(o I33) I33 { return I33.cast(i | o) }

// Xor returns i^o reduced to 33 bits.
func (i I33) Xor(o I33) I33 { return I33.cast(i ^ o) }

// Shr returns i>>o reduced to 33 bits.
func (i I33) Shr(o I33) I33 { return I33.cast(i >> o) }

// Shl returns i<<o reduced to 33 bits.
func (i I33) Shl(o I33) I33 { return I33.cast(i << o) }

// Not returns bitwise complement of i reduced to 33 bits.
func (i I33) Not() I33 { return I33.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I33) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I33) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U33 bit pattern.
func (i I33) Unsigned() U33 { return U33.cast(U33(i)) }

// Clamp returns value saturated into the representable range of I33.
func (i I33) Clamp(value int64) I33 { return iclamp[I33](value) }

// Clip masks i to a bits-wide low field and sign-extends to 33 bits.
func (i I33) Clip(bits int) I33 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I33(uint64(i)&m^b) - I33(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I33) Bit(index int) I33 {
	if index < 0 {
		index = 33 + index
	}
	return (i >> I33(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I33) SetBit(index int, v I33) {
	if index < 0 {
		index = 33 + index
	}
	bit := I33(1) << I33(index)
	*i = I33.cast((*i &^ bit) | ((v & 1) << I33(index)))
}

// Bitref returns a Range for bit index.
func (i *I33) Bitref(index int) Range[I33] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I33) Bits(lo, hi int) I33 {
	p := 33
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I33((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I33(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I33) SetBits(lo, hi int, v I33) {
	p := 33
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I33((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I33.cast((*i &^ mask) | ((v << I33(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I33) Bitsref(lo, hi int) Range[I33] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I33) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I33) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I33(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I33) Byteref(idx int) Range[I33] { return byte(i, idx) }

// U34 is an 34-bit unsigned integer.
type U34 uint64

// AsU34 returns a U34 representing v reduced to 34 bits.
func AsU34[T Unsigned](v T) U34 { return U34.cast(U34(v)) }

func (U34) nbits() int  { return 34 }
func (u U34) mask() U34 { return 0x3ffffffff }
func (u U34) cast() U34 { return u & u.mask() }

// Set assigns v reduced to 34 bits to u.
func (u *U34) Set(v U34) U34 { *u = new(U34(v)).cast(); return U34.cast(v) }

// Add returns u+o reduced to 34 bits.
func (u U34) Add(o U34) U34 { return U34.cast(u + o) }

// Sub returns u-o reduced to 34 bits.
func (u U34) Sub(o U34) U34 { return U34.cast(u - o) }

// Inc returns u+1 reduced to 34 bits.
func (u U34) Inc() U34 { return U34.cast(u + 1) }

// Dec returns u-1 reduced to 34 bits.
func (u U34) Dec() U34 { return U34.cast(u - 1) }

// Mul returns u*o reduced to 34 bits.
func (u U34) Mul(o U34) U34 { return U34.cast(u * o) }

// Div returns u divided by o.
func (u U34) Div(o U34) U34 { return U34.cast(u / o) }

// Mod returns u modulo o.
func (u U34) Mod(o U34) U34 { return U34.cast(u % o) }

// And returns u&o reduced to 34 bits.
func (u U34) And(o U34) U34 { return U34.cast(u & o) }

// Or returns u|o reduced to 34 bits.
func (u U34) Or(o U34) U34 { return U34.cast(u | o) }

// Xor returns u^o reduced to 34 bits.
func (u U34) Xor(o U34) U34 { return U34.cast(u ^ o) }

// Shr returns u>>o reduced to 34 bits.
func (u U34) Shr(o U34) U34 { return U34.cast(u >> o) }

// Shl returns u<<o reduced to 34 bits.
func (u U34) Shl(o U34) U34 { return U34.cast(u << o) }

// Not returns bitwise complement of u reduced to 34 bits.
func (u U34) Not() U34 { return U34.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U34) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U34) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U34 value.
func (u U34) Signed() I34 { return I34.cast(I34(u)) }

// Clamp returns value saturated into the representable range of U34.
func (u U34) Clamp(value uint64) U34 { return uclamp[U34](value) }

// Clip masks u to a bits-wide low field.
func (u U34) Clip(bits int) U34 {
	b := 1 << (bits - 1)
	m := U34(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U34) Bit(index int) U34 {
	if index < 0 {
		index = 34 + index
	}
	return (u >> U34(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U34) SetBit(index int, v U34) {
	if index < 0 {
		index = 34 + index
	}
	bit := U34(1) << U34(index)
	*u = U34.cast((*u &^ bit) | ((v & 1) << U34(index)))
}

// Bitref returns a Range for bit index.
func (u *U34) Bitref(index int) Range[U34] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U34) Bits(lo, hi int) U34 {
	p := 34
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U34((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U34(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U34) SetBits(lo, hi int, v U34) {
	p := 34
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U34((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U34.cast((*u &^ mask) | ((v << U34(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U34) Bitsref(lo, hi int) Range[U34] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U34) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U34) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U34(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U34) Byteref(idx int) Range[U34] { return byte(u, idx) }

// I34 is an 34-bit signed integer in two's complement.
type I34 int64

// AsI34 returns a I34 representing v reduced to 34 bits.
func AsI34[T Signed](v T) I34 { return I34.cast(I34(v)) }

func (I34) nbits() int  { return 34 }
func (i I34) mask() I34 { return 0x3ffffffff }
func (i I34) sign() I34 { return 1 << (i.nbits() - 1) }
func (i I34) cast() I34 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 34 bits to u.
func (i *I34) Set(v I34) I34 { *i = new(I34(v)).cast(); return I34.cast(v) }

// Add returns i+o reduced to 34 bits.
func (i I34) Add(o I34) I34 { return I34.cast(i + o) }

// Sub returns i-o reduced to 34 bits.
func (i I34) Sub(o I34) I34 { return I34.cast(i - o) }

// Inc returns i+1 reduced to 34 bits.
func (i I34) Inc() I34 { return I34.cast(i + 1) }

// Dec returns i-1 reduced to 34 bits.
func (i I34) Dec() I34 { return I34.cast(i - 1) }

// Mul returns i*o reduced to 34 bits.
func (i I34) Mul(o I34) I34 { return I34.cast(i * o) }

// Div returns i divided by o.
func (i I34) Div(o I34) I34 { return I34.cast(i / o) }

// Mod returns i modulo o.
func (i I34) Mod(o I34) I34 { return I34.cast(i % o) }

// And returns i&o reduced to 34 bits.
func (i I34) And(o I34) I34 { return I34.cast(i & o) }

// Or returns i|o reduced to 34 bits.
func (i I34) Or(o I34) I34 { return I34.cast(i | o) }

// Xor returns i^o reduced to 34 bits.
func (i I34) Xor(o I34) I34 { return I34.cast(i ^ o) }

// Shr returns i>>o reduced to 34 bits.
func (i I34) Shr(o I34) I34 { return I34.cast(i >> o) }

// Shl returns i<<o reduced to 34 bits.
func (i I34) Shl(o I34) I34 { return I34.cast(i << o) }

// Not returns bitwise complement of i reduced to 34 bits.
func (i I34) Not() I34 { return I34.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I34) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I34) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U34 bit pattern.
func (i I34) Unsigned() U34 { return U34.cast(U34(i)) }

// Clamp returns value saturated into the representable range of I34.
func (i I34) Clamp(value int64) I34 { return iclamp[I34](value) }

// Clip masks i to a bits-wide low field and sign-extends to 34 bits.
func (i I34) Clip(bits int) I34 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I34(uint64(i)&m^b) - I34(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I34) Bit(index int) I34 {
	if index < 0 {
		index = 34 + index
	}
	return (i >> I34(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I34) SetBit(index int, v I34) {
	if index < 0 {
		index = 34 + index
	}
	bit := I34(1) << I34(index)
	*i = I34.cast((*i &^ bit) | ((v & 1) << I34(index)))
}

// Bitref returns a Range for bit index.
func (i *I34) Bitref(index int) Range[I34] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I34) Bits(lo, hi int) I34 {
	p := 34
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I34((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I34(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I34) SetBits(lo, hi int, v I34) {
	p := 34
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I34((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I34.cast((*i &^ mask) | ((v << I34(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I34) Bitsref(lo, hi int) Range[I34] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I34) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I34) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I34(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I34) Byteref(idx int) Range[I34] { return byte(i, idx) }

// U35 is an 35-bit unsigned integer.
type U35 uint64

// AsU35 returns a U35 representing v reduced to 35 bits.
func AsU35[T Unsigned](v T) U35 { return U35.cast(U35(v)) }

func (U35) nbits() int  { return 35 }
func (u U35) mask() U35 { return 0x7ffffffff }
func (u U35) cast() U35 { return u & u.mask() }

// Set assigns v reduced to 35 bits to u.
func (u *U35) Set(v U35) U35 { *u = new(U35(v)).cast(); return U35.cast(v) }

// Add returns u+o reduced to 35 bits.
func (u U35) Add(o U35) U35 { return U35.cast(u + o) }

// Sub returns u-o reduced to 35 bits.
func (u U35) Sub(o U35) U35 { return U35.cast(u - o) }

// Inc returns u+1 reduced to 35 bits.
func (u U35) Inc() U35 { return U35.cast(u + 1) }

// Dec returns u-1 reduced to 35 bits.
func (u U35) Dec() U35 { return U35.cast(u - 1) }

// Mul returns u*o reduced to 35 bits.
func (u U35) Mul(o U35) U35 { return U35.cast(u * o) }

// Div returns u divided by o.
func (u U35) Div(o U35) U35 { return U35.cast(u / o) }

// Mod returns u modulo o.
func (u U35) Mod(o U35) U35 { return U35.cast(u % o) }

// And returns u&o reduced to 35 bits.
func (u U35) And(o U35) U35 { return U35.cast(u & o) }

// Or returns u|o reduced to 35 bits.
func (u U35) Or(o U35) U35 { return U35.cast(u | o) }

// Xor returns u^o reduced to 35 bits.
func (u U35) Xor(o U35) U35 { return U35.cast(u ^ o) }

// Shr returns u>>o reduced to 35 bits.
func (u U35) Shr(o U35) U35 { return U35.cast(u >> o) }

// Shl returns u<<o reduced to 35 bits.
func (u U35) Shl(o U35) U35 { return U35.cast(u << o) }

// Not returns bitwise complement of u reduced to 35 bits.
func (u U35) Not() U35 { return U35.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U35) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U35) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U35 value.
func (u U35) Signed() I35 { return I35.cast(I35(u)) }

// Clamp returns value saturated into the representable range of U35.
func (u U35) Clamp(value uint64) U35 { return uclamp[U35](value) }

// Clip masks u to a bits-wide low field.
func (u U35) Clip(bits int) U35 {
	b := 1 << (bits - 1)
	m := U35(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U35) Bit(index int) U35 {
	if index < 0 {
		index = 35 + index
	}
	return (u >> U35(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U35) SetBit(index int, v U35) {
	if index < 0 {
		index = 35 + index
	}
	bit := U35(1) << U35(index)
	*u = U35.cast((*u &^ bit) | ((v & 1) << U35(index)))
}

// Bitref returns a Range for bit index.
func (u *U35) Bitref(index int) Range[U35] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U35) Bits(lo, hi int) U35 {
	p := 35
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U35((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U35(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U35) SetBits(lo, hi int, v U35) {
	p := 35
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U35((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U35.cast((*u &^ mask) | ((v << U35(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U35) Bitsref(lo, hi int) Range[U35] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U35) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U35) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U35(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U35) Byteref(idx int) Range[U35] { return byte(u, idx) }

// I35 is an 35-bit signed integer in two's complement.
type I35 int64

// AsI35 returns a I35 representing v reduced to 35 bits.
func AsI35[T Signed](v T) I35 { return I35.cast(I35(v)) }

func (I35) nbits() int  { return 35 }
func (i I35) mask() I35 { return 0x7ffffffff }
func (i I35) sign() I35 { return 1 << (i.nbits() - 1) }
func (i I35) cast() I35 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 35 bits to u.
func (i *I35) Set(v I35) I35 { *i = new(I35(v)).cast(); return I35.cast(v) }

// Add returns i+o reduced to 35 bits.
func (i I35) Add(o I35) I35 { return I35.cast(i + o) }

// Sub returns i-o reduced to 35 bits.
func (i I35) Sub(o I35) I35 { return I35.cast(i - o) }

// Inc returns i+1 reduced to 35 bits.
func (i I35) Inc() I35 { return I35.cast(i + 1) }

// Dec returns i-1 reduced to 35 bits.
func (i I35) Dec() I35 { return I35.cast(i - 1) }

// Mul returns i*o reduced to 35 bits.
func (i I35) Mul(o I35) I35 { return I35.cast(i * o) }

// Div returns i divided by o.
func (i I35) Div(o I35) I35 { return I35.cast(i / o) }

// Mod returns i modulo o.
func (i I35) Mod(o I35) I35 { return I35.cast(i % o) }

// And returns i&o reduced to 35 bits.
func (i I35) And(o I35) I35 { return I35.cast(i & o) }

// Or returns i|o reduced to 35 bits.
func (i I35) Or(o I35) I35 { return I35.cast(i | o) }

// Xor returns i^o reduced to 35 bits.
func (i I35) Xor(o I35) I35 { return I35.cast(i ^ o) }

// Shr returns i>>o reduced to 35 bits.
func (i I35) Shr(o I35) I35 { return I35.cast(i >> o) }

// Shl returns i<<o reduced to 35 bits.
func (i I35) Shl(o I35) I35 { return I35.cast(i << o) }

// Not returns bitwise complement of i reduced to 35 bits.
func (i I35) Not() I35 { return I35.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I35) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I35) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U35 bit pattern.
func (i I35) Unsigned() U35 { return U35.cast(U35(i)) }

// Clamp returns value saturated into the representable range of I35.
func (i I35) Clamp(value int64) I35 { return iclamp[I35](value) }

// Clip masks i to a bits-wide low field and sign-extends to 35 bits.
func (i I35) Clip(bits int) I35 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I35(uint64(i)&m^b) - I35(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I35) Bit(index int) I35 {
	if index < 0 {
		index = 35 + index
	}
	return (i >> I35(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I35) SetBit(index int, v I35) {
	if index < 0 {
		index = 35 + index
	}
	bit := I35(1) << I35(index)
	*i = I35.cast((*i &^ bit) | ((v & 1) << I35(index)))
}

// Bitref returns a Range for bit index.
func (i *I35) Bitref(index int) Range[I35] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I35) Bits(lo, hi int) I35 {
	p := 35
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I35((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I35(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I35) SetBits(lo, hi int, v I35) {
	p := 35
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I35((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I35.cast((*i &^ mask) | ((v << I35(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I35) Bitsref(lo, hi int) Range[I35] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I35) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I35) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I35(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I35) Byteref(idx int) Range[I35] { return byte(i, idx) }

// U36 is an 36-bit unsigned integer.
type U36 uint64

// AsU36 returns a U36 representing v reduced to 36 bits.
func AsU36[T Unsigned](v T) U36 { return U36.cast(U36(v)) }

func (U36) nbits() int  { return 36 }
func (u U36) mask() U36 { return 0xfffffffff }
func (u U36) cast() U36 { return u & u.mask() }

// Set assigns v reduced to 36 bits to u.
func (u *U36) Set(v U36) U36 { *u = new(U36(v)).cast(); return U36.cast(v) }

// Add returns u+o reduced to 36 bits.
func (u U36) Add(o U36) U36 { return U36.cast(u + o) }

// Sub returns u-o reduced to 36 bits.
func (u U36) Sub(o U36) U36 { return U36.cast(u - o) }

// Inc returns u+1 reduced to 36 bits.
func (u U36) Inc() U36 { return U36.cast(u + 1) }

// Dec returns u-1 reduced to 36 bits.
func (u U36) Dec() U36 { return U36.cast(u - 1) }

// Mul returns u*o reduced to 36 bits.
func (u U36) Mul(o U36) U36 { return U36.cast(u * o) }

// Div returns u divided by o.
func (u U36) Div(o U36) U36 { return U36.cast(u / o) }

// Mod returns u modulo o.
func (u U36) Mod(o U36) U36 { return U36.cast(u % o) }

// And returns u&o reduced to 36 bits.
func (u U36) And(o U36) U36 { return U36.cast(u & o) }

// Or returns u|o reduced to 36 bits.
func (u U36) Or(o U36) U36 { return U36.cast(u | o) }

// Xor returns u^o reduced to 36 bits.
func (u U36) Xor(o U36) U36 { return U36.cast(u ^ o) }

// Shr returns u>>o reduced to 36 bits.
func (u U36) Shr(o U36) U36 { return U36.cast(u >> o) }

// Shl returns u<<o reduced to 36 bits.
func (u U36) Shl(o U36) U36 { return U36.cast(u << o) }

// Not returns bitwise complement of u reduced to 36 bits.
func (u U36) Not() U36 { return U36.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U36) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U36) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U36 value.
func (u U36) Signed() I36 { return I36.cast(I36(u)) }

// Clamp returns value saturated into the representable range of U36.
func (u U36) Clamp(value uint64) U36 { return uclamp[U36](value) }

// Clip masks u to a bits-wide low field.
func (u U36) Clip(bits int) U36 {
	b := 1 << (bits - 1)
	m := U36(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U36) Bit(index int) U36 {
	if index < 0 {
		index = 36 + index
	}
	return (u >> U36(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U36) SetBit(index int, v U36) {
	if index < 0 {
		index = 36 + index
	}
	bit := U36(1) << U36(index)
	*u = U36.cast((*u &^ bit) | ((v & 1) << U36(index)))
}

// Bitref returns a Range for bit index.
func (u *U36) Bitref(index int) Range[U36] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U36) Bits(lo, hi int) U36 {
	p := 36
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U36((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U36(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U36) SetBits(lo, hi int, v U36) {
	p := 36
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U36((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U36.cast((*u &^ mask) | ((v << U36(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U36) Bitsref(lo, hi int) Range[U36] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U36) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U36) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U36(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U36) Byteref(idx int) Range[U36] { return byte(u, idx) }

// I36 is an 36-bit signed integer in two's complement.
type I36 int64

// AsI36 returns a I36 representing v reduced to 36 bits.
func AsI36[T Signed](v T) I36 { return I36.cast(I36(v)) }

func (I36) nbits() int  { return 36 }
func (i I36) mask() I36 { return 0xfffffffff }
func (i I36) sign() I36 { return 1 << (i.nbits() - 1) }
func (i I36) cast() I36 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 36 bits to u.
func (i *I36) Set(v I36) I36 { *i = new(I36(v)).cast(); return I36.cast(v) }

// Add returns i+o reduced to 36 bits.
func (i I36) Add(o I36) I36 { return I36.cast(i + o) }

// Sub returns i-o reduced to 36 bits.
func (i I36) Sub(o I36) I36 { return I36.cast(i - o) }

// Inc returns i+1 reduced to 36 bits.
func (i I36) Inc() I36 { return I36.cast(i + 1) }

// Dec returns i-1 reduced to 36 bits.
func (i I36) Dec() I36 { return I36.cast(i - 1) }

// Mul returns i*o reduced to 36 bits.
func (i I36) Mul(o I36) I36 { return I36.cast(i * o) }

// Div returns i divided by o.
func (i I36) Div(o I36) I36 { return I36.cast(i / o) }

// Mod returns i modulo o.
func (i I36) Mod(o I36) I36 { return I36.cast(i % o) }

// And returns i&o reduced to 36 bits.
func (i I36) And(o I36) I36 { return I36.cast(i & o) }

// Or returns i|o reduced to 36 bits.
func (i I36) Or(o I36) I36 { return I36.cast(i | o) }

// Xor returns i^o reduced to 36 bits.
func (i I36) Xor(o I36) I36 { return I36.cast(i ^ o) }

// Shr returns i>>o reduced to 36 bits.
func (i I36) Shr(o I36) I36 { return I36.cast(i >> o) }

// Shl returns i<<o reduced to 36 bits.
func (i I36) Shl(o I36) I36 { return I36.cast(i << o) }

// Not returns bitwise complement of i reduced to 36 bits.
func (i I36) Not() I36 { return I36.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I36) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I36) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U36 bit pattern.
func (i I36) Unsigned() U36 { return U36.cast(U36(i)) }

// Clamp returns value saturated into the representable range of I36.
func (i I36) Clamp(value int64) I36 { return iclamp[I36](value) }

// Clip masks i to a bits-wide low field and sign-extends to 36 bits.
func (i I36) Clip(bits int) I36 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I36(uint64(i)&m^b) - I36(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I36) Bit(index int) I36 {
	if index < 0 {
		index = 36 + index
	}
	return (i >> I36(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I36) SetBit(index int, v I36) {
	if index < 0 {
		index = 36 + index
	}
	bit := I36(1) << I36(index)
	*i = I36.cast((*i &^ bit) | ((v & 1) << I36(index)))
}

// Bitref returns a Range for bit index.
func (i *I36) Bitref(index int) Range[I36] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I36) Bits(lo, hi int) I36 {
	p := 36
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I36((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I36(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I36) SetBits(lo, hi int, v I36) {
	p := 36
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I36((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I36.cast((*i &^ mask) | ((v << I36(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I36) Bitsref(lo, hi int) Range[I36] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I36) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I36) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I36(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I36) Byteref(idx int) Range[I36] { return byte(i, idx) }

// U37 is an 37-bit unsigned integer.
type U37 uint64

// AsU37 returns a U37 representing v reduced to 37 bits.
func AsU37[T Unsigned](v T) U37 { return U37.cast(U37(v)) }

func (U37) nbits() int  { return 37 }
func (u U37) mask() U37 { return 0x1fffffffff }
func (u U37) cast() U37 { return u & u.mask() }

// Set assigns v reduced to 37 bits to u.
func (u *U37) Set(v U37) U37 { *u = new(U37(v)).cast(); return U37.cast(v) }

// Add returns u+o reduced to 37 bits.
func (u U37) Add(o U37) U37 { return U37.cast(u + o) }

// Sub returns u-o reduced to 37 bits.
func (u U37) Sub(o U37) U37 { return U37.cast(u - o) }

// Inc returns u+1 reduced to 37 bits.
func (u U37) Inc() U37 { return U37.cast(u + 1) }

// Dec returns u-1 reduced to 37 bits.
func (u U37) Dec() U37 { return U37.cast(u - 1) }

// Mul returns u*o reduced to 37 bits.
func (u U37) Mul(o U37) U37 { return U37.cast(u * o) }

// Div returns u divided by o.
func (u U37) Div(o U37) U37 { return U37.cast(u / o) }

// Mod returns u modulo o.
func (u U37) Mod(o U37) U37 { return U37.cast(u % o) }

// And returns u&o reduced to 37 bits.
func (u U37) And(o U37) U37 { return U37.cast(u & o) }

// Or returns u|o reduced to 37 bits.
func (u U37) Or(o U37) U37 { return U37.cast(u | o) }

// Xor returns u^o reduced to 37 bits.
func (u U37) Xor(o U37) U37 { return U37.cast(u ^ o) }

// Shr returns u>>o reduced to 37 bits.
func (u U37) Shr(o U37) U37 { return U37.cast(u >> o) }

// Shl returns u<<o reduced to 37 bits.
func (u U37) Shl(o U37) U37 { return U37.cast(u << o) }

// Not returns bitwise complement of u reduced to 37 bits.
func (u U37) Not() U37 { return U37.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U37) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U37) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U37 value.
func (u U37) Signed() I37 { return I37.cast(I37(u)) }

// Clamp returns value saturated into the representable range of U37.
func (u U37) Clamp(value uint64) U37 { return uclamp[U37](value) }

// Clip masks u to a bits-wide low field.
func (u U37) Clip(bits int) U37 {
	b := 1 << (bits - 1)
	m := U37(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U37) Bit(index int) U37 {
	if index < 0 {
		index = 37 + index
	}
	return (u >> U37(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U37) SetBit(index int, v U37) {
	if index < 0 {
		index = 37 + index
	}
	bit := U37(1) << U37(index)
	*u = U37.cast((*u &^ bit) | ((v & 1) << U37(index)))
}

// Bitref returns a Range for bit index.
func (u *U37) Bitref(index int) Range[U37] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U37) Bits(lo, hi int) U37 {
	p := 37
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U37((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U37(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U37) SetBits(lo, hi int, v U37) {
	p := 37
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U37((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U37.cast((*u &^ mask) | ((v << U37(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U37) Bitsref(lo, hi int) Range[U37] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U37) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U37) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U37(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U37) Byteref(idx int) Range[U37] { return byte(u, idx) }

// I37 is an 37-bit signed integer in two's complement.
type I37 int64

// AsI37 returns a I37 representing v reduced to 37 bits.
func AsI37[T Signed](v T) I37 { return I37.cast(I37(v)) }

func (I37) nbits() int  { return 37 }
func (i I37) mask() I37 { return 0x1fffffffff }
func (i I37) sign() I37 { return 1 << (i.nbits() - 1) }
func (i I37) cast() I37 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 37 bits to u.
func (i *I37) Set(v I37) I37 { *i = new(I37(v)).cast(); return I37.cast(v) }

// Add returns i+o reduced to 37 bits.
func (i I37) Add(o I37) I37 { return I37.cast(i + o) }

// Sub returns i-o reduced to 37 bits.
func (i I37) Sub(o I37) I37 { return I37.cast(i - o) }

// Inc returns i+1 reduced to 37 bits.
func (i I37) Inc() I37 { return I37.cast(i + 1) }

// Dec returns i-1 reduced to 37 bits.
func (i I37) Dec() I37 { return I37.cast(i - 1) }

// Mul returns i*o reduced to 37 bits.
func (i I37) Mul(o I37) I37 { return I37.cast(i * o) }

// Div returns i divided by o.
func (i I37) Div(o I37) I37 { return I37.cast(i / o) }

// Mod returns i modulo o.
func (i I37) Mod(o I37) I37 { return I37.cast(i % o) }

// And returns i&o reduced to 37 bits.
func (i I37) And(o I37) I37 { return I37.cast(i & o) }

// Or returns i|o reduced to 37 bits.
func (i I37) Or(o I37) I37 { return I37.cast(i | o) }

// Xor returns i^o reduced to 37 bits.
func (i I37) Xor(o I37) I37 { return I37.cast(i ^ o) }

// Shr returns i>>o reduced to 37 bits.
func (i I37) Shr(o I37) I37 { return I37.cast(i >> o) }

// Shl returns i<<o reduced to 37 bits.
func (i I37) Shl(o I37) I37 { return I37.cast(i << o) }

// Not returns bitwise complement of i reduced to 37 bits.
func (i I37) Not() I37 { return I37.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I37) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I37) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U37 bit pattern.
func (i I37) Unsigned() U37 { return U37.cast(U37(i)) }

// Clamp returns value saturated into the representable range of I37.
func (i I37) Clamp(value int64) I37 { return iclamp[I37](value) }

// Clip masks i to a bits-wide low field and sign-extends to 37 bits.
func (i I37) Clip(bits int) I37 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I37(uint64(i)&m^b) - I37(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I37) Bit(index int) I37 {
	if index < 0 {
		index = 37 + index
	}
	return (i >> I37(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I37) SetBit(index int, v I37) {
	if index < 0 {
		index = 37 + index
	}
	bit := I37(1) << I37(index)
	*i = I37.cast((*i &^ bit) | ((v & 1) << I37(index)))
}

// Bitref returns a Range for bit index.
func (i *I37) Bitref(index int) Range[I37] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I37) Bits(lo, hi int) I37 {
	p := 37
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I37((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I37(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I37) SetBits(lo, hi int, v I37) {
	p := 37
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I37((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I37.cast((*i &^ mask) | ((v << I37(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I37) Bitsref(lo, hi int) Range[I37] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I37) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I37) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I37(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I37) Byteref(idx int) Range[I37] { return byte(i, idx) }

// U38 is an 38-bit unsigned integer.
type U38 uint64

// AsU38 returns a U38 representing v reduced to 38 bits.
func AsU38[T Unsigned](v T) U38 { return U38.cast(U38(v)) }

func (U38) nbits() int  { return 38 }
func (u U38) mask() U38 { return 0x3fffffffff }
func (u U38) cast() U38 { return u & u.mask() }

// Set assigns v reduced to 38 bits to u.
func (u *U38) Set(v U38) U38 { *u = new(U38(v)).cast(); return U38.cast(v) }

// Add returns u+o reduced to 38 bits.
func (u U38) Add(o U38) U38 { return U38.cast(u + o) }

// Sub returns u-o reduced to 38 bits.
func (u U38) Sub(o U38) U38 { return U38.cast(u - o) }

// Inc returns u+1 reduced to 38 bits.
func (u U38) Inc() U38 { return U38.cast(u + 1) }

// Dec returns u-1 reduced to 38 bits.
func (u U38) Dec() U38 { return U38.cast(u - 1) }

// Mul returns u*o reduced to 38 bits.
func (u U38) Mul(o U38) U38 { return U38.cast(u * o) }

// Div returns u divided by o.
func (u U38) Div(o U38) U38 { return U38.cast(u / o) }

// Mod returns u modulo o.
func (u U38) Mod(o U38) U38 { return U38.cast(u % o) }

// And returns u&o reduced to 38 bits.
func (u U38) And(o U38) U38 { return U38.cast(u & o) }

// Or returns u|o reduced to 38 bits.
func (u U38) Or(o U38) U38 { return U38.cast(u | o) }

// Xor returns u^o reduced to 38 bits.
func (u U38) Xor(o U38) U38 { return U38.cast(u ^ o) }

// Shr returns u>>o reduced to 38 bits.
func (u U38) Shr(o U38) U38 { return U38.cast(u >> o) }

// Shl returns u<<o reduced to 38 bits.
func (u U38) Shl(o U38) U38 { return U38.cast(u << o) }

// Not returns bitwise complement of u reduced to 38 bits.
func (u U38) Not() U38 { return U38.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U38) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U38) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U38 value.
func (u U38) Signed() I38 { return I38.cast(I38(u)) }

// Clamp returns value saturated into the representable range of U38.
func (u U38) Clamp(value uint64) U38 { return uclamp[U38](value) }

// Clip masks u to a bits-wide low field.
func (u U38) Clip(bits int) U38 {
	b := 1 << (bits - 1)
	m := U38(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U38) Bit(index int) U38 {
	if index < 0 {
		index = 38 + index
	}
	return (u >> U38(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U38) SetBit(index int, v U38) {
	if index < 0 {
		index = 38 + index
	}
	bit := U38(1) << U38(index)
	*u = U38.cast((*u &^ bit) | ((v & 1) << U38(index)))
}

// Bitref returns a Range for bit index.
func (u *U38) Bitref(index int) Range[U38] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U38) Bits(lo, hi int) U38 {
	p := 38
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U38((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U38(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U38) SetBits(lo, hi int, v U38) {
	p := 38
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U38((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U38.cast((*u &^ mask) | ((v << U38(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U38) Bitsref(lo, hi int) Range[U38] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U38) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U38) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U38(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U38) Byteref(idx int) Range[U38] { return byte(u, idx) }

// I38 is an 38-bit signed integer in two's complement.
type I38 int64

// AsI38 returns a I38 representing v reduced to 38 bits.
func AsI38[T Signed](v T) I38 { return I38.cast(I38(v)) }

func (I38) nbits() int  { return 38 }
func (i I38) mask() I38 { return 0x3fffffffff }
func (i I38) sign() I38 { return 1 << (i.nbits() - 1) }
func (i I38) cast() I38 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 38 bits to u.
func (i *I38) Set(v I38) I38 { *i = new(I38(v)).cast(); return I38.cast(v) }

// Add returns i+o reduced to 38 bits.
func (i I38) Add(o I38) I38 { return I38.cast(i + o) }

// Sub returns i-o reduced to 38 bits.
func (i I38) Sub(o I38) I38 { return I38.cast(i - o) }

// Inc returns i+1 reduced to 38 bits.
func (i I38) Inc() I38 { return I38.cast(i + 1) }

// Dec returns i-1 reduced to 38 bits.
func (i I38) Dec() I38 { return I38.cast(i - 1) }

// Mul returns i*o reduced to 38 bits.
func (i I38) Mul(o I38) I38 { return I38.cast(i * o) }

// Div returns i divided by o.
func (i I38) Div(o I38) I38 { return I38.cast(i / o) }

// Mod returns i modulo o.
func (i I38) Mod(o I38) I38 { return I38.cast(i % o) }

// And returns i&o reduced to 38 bits.
func (i I38) And(o I38) I38 { return I38.cast(i & o) }

// Or returns i|o reduced to 38 bits.
func (i I38) Or(o I38) I38 { return I38.cast(i | o) }

// Xor returns i^o reduced to 38 bits.
func (i I38) Xor(o I38) I38 { return I38.cast(i ^ o) }

// Shr returns i>>o reduced to 38 bits.
func (i I38) Shr(o I38) I38 { return I38.cast(i >> o) }

// Shl returns i<<o reduced to 38 bits.
func (i I38) Shl(o I38) I38 { return I38.cast(i << o) }

// Not returns bitwise complement of i reduced to 38 bits.
func (i I38) Not() I38 { return I38.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I38) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I38) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U38 bit pattern.
func (i I38) Unsigned() U38 { return U38.cast(U38(i)) }

// Clamp returns value saturated into the representable range of I38.
func (i I38) Clamp(value int64) I38 { return iclamp[I38](value) }

// Clip masks i to a bits-wide low field and sign-extends to 38 bits.
func (i I38) Clip(bits int) I38 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I38(uint64(i)&m^b) - I38(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I38) Bit(index int) I38 {
	if index < 0 {
		index = 38 + index
	}
	return (i >> I38(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I38) SetBit(index int, v I38) {
	if index < 0 {
		index = 38 + index
	}
	bit := I38(1) << I38(index)
	*i = I38.cast((*i &^ bit) | ((v & 1) << I38(index)))
}

// Bitref returns a Range for bit index.
func (i *I38) Bitref(index int) Range[I38] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I38) Bits(lo, hi int) I38 {
	p := 38
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I38((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I38(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I38) SetBits(lo, hi int, v I38) {
	p := 38
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I38((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I38.cast((*i &^ mask) | ((v << I38(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I38) Bitsref(lo, hi int) Range[I38] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I38) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I38) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I38(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I38) Byteref(idx int) Range[I38] { return byte(i, idx) }

// U39 is an 39-bit unsigned integer.
type U39 uint64

// AsU39 returns a U39 representing v reduced to 39 bits.
func AsU39[T Unsigned](v T) U39 { return U39.cast(U39(v)) }

func (U39) nbits() int  { return 39 }
func (u U39) mask() U39 { return 0x7fffffffff }
func (u U39) cast() U39 { return u & u.mask() }

// Set assigns v reduced to 39 bits to u.
func (u *U39) Set(v U39) U39 { *u = new(U39(v)).cast(); return U39.cast(v) }

// Add returns u+o reduced to 39 bits.
func (u U39) Add(o U39) U39 { return U39.cast(u + o) }

// Sub returns u-o reduced to 39 bits.
func (u U39) Sub(o U39) U39 { return U39.cast(u - o) }

// Inc returns u+1 reduced to 39 bits.
func (u U39) Inc() U39 { return U39.cast(u + 1) }

// Dec returns u-1 reduced to 39 bits.
func (u U39) Dec() U39 { return U39.cast(u - 1) }

// Mul returns u*o reduced to 39 bits.
func (u U39) Mul(o U39) U39 { return U39.cast(u * o) }

// Div returns u divided by o.
func (u U39) Div(o U39) U39 { return U39.cast(u / o) }

// Mod returns u modulo o.
func (u U39) Mod(o U39) U39 { return U39.cast(u % o) }

// And returns u&o reduced to 39 bits.
func (u U39) And(o U39) U39 { return U39.cast(u & o) }

// Or returns u|o reduced to 39 bits.
func (u U39) Or(o U39) U39 { return U39.cast(u | o) }

// Xor returns u^o reduced to 39 bits.
func (u U39) Xor(o U39) U39 { return U39.cast(u ^ o) }

// Shr returns u>>o reduced to 39 bits.
func (u U39) Shr(o U39) U39 { return U39.cast(u >> o) }

// Shl returns u<<o reduced to 39 bits.
func (u U39) Shl(o U39) U39 { return U39.cast(u << o) }

// Not returns bitwise complement of u reduced to 39 bits.
func (u U39) Not() U39 { return U39.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U39) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U39) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U39 value.
func (u U39) Signed() I39 { return I39.cast(I39(u)) }

// Clamp returns value saturated into the representable range of U39.
func (u U39) Clamp(value uint64) U39 { return uclamp[U39](value) }

// Clip masks u to a bits-wide low field.
func (u U39) Clip(bits int) U39 {
	b := 1 << (bits - 1)
	m := U39(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U39) Bit(index int) U39 {
	if index < 0 {
		index = 39 + index
	}
	return (u >> U39(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U39) SetBit(index int, v U39) {
	if index < 0 {
		index = 39 + index
	}
	bit := U39(1) << U39(index)
	*u = U39.cast((*u &^ bit) | ((v & 1) << U39(index)))
}

// Bitref returns a Range for bit index.
func (u *U39) Bitref(index int) Range[U39] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U39) Bits(lo, hi int) U39 {
	p := 39
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U39((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U39(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U39) SetBits(lo, hi int, v U39) {
	p := 39
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U39((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U39.cast((*u &^ mask) | ((v << U39(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U39) Bitsref(lo, hi int) Range[U39] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U39) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U39) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U39(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U39) Byteref(idx int) Range[U39] { return byte(u, idx) }

// I39 is an 39-bit signed integer in two's complement.
type I39 int64

// AsI39 returns a I39 representing v reduced to 39 bits.
func AsI39[T Signed](v T) I39 { return I39.cast(I39(v)) }

func (I39) nbits() int  { return 39 }
func (i I39) mask() I39 { return 0x7fffffffff }
func (i I39) sign() I39 { return 1 << (i.nbits() - 1) }
func (i I39) cast() I39 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 39 bits to u.
func (i *I39) Set(v I39) I39 { *i = new(I39(v)).cast(); return I39.cast(v) }

// Add returns i+o reduced to 39 bits.
func (i I39) Add(o I39) I39 { return I39.cast(i + o) }

// Sub returns i-o reduced to 39 bits.
func (i I39) Sub(o I39) I39 { return I39.cast(i - o) }

// Inc returns i+1 reduced to 39 bits.
func (i I39) Inc() I39 { return I39.cast(i + 1) }

// Dec returns i-1 reduced to 39 bits.
func (i I39) Dec() I39 { return I39.cast(i - 1) }

// Mul returns i*o reduced to 39 bits.
func (i I39) Mul(o I39) I39 { return I39.cast(i * o) }

// Div returns i divided by o.
func (i I39) Div(o I39) I39 { return I39.cast(i / o) }

// Mod returns i modulo o.
func (i I39) Mod(o I39) I39 { return I39.cast(i % o) }

// And returns i&o reduced to 39 bits.
func (i I39) And(o I39) I39 { return I39.cast(i & o) }

// Or returns i|o reduced to 39 bits.
func (i I39) Or(o I39) I39 { return I39.cast(i | o) }

// Xor returns i^o reduced to 39 bits.
func (i I39) Xor(o I39) I39 { return I39.cast(i ^ o) }

// Shr returns i>>o reduced to 39 bits.
func (i I39) Shr(o I39) I39 { return I39.cast(i >> o) }

// Shl returns i<<o reduced to 39 bits.
func (i I39) Shl(o I39) I39 { return I39.cast(i << o) }

// Not returns bitwise complement of i reduced to 39 bits.
func (i I39) Not() I39 { return I39.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I39) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I39) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U39 bit pattern.
func (i I39) Unsigned() U39 { return U39.cast(U39(i)) }

// Clamp returns value saturated into the representable range of I39.
func (i I39) Clamp(value int64) I39 { return iclamp[I39](value) }

// Clip masks i to a bits-wide low field and sign-extends to 39 bits.
func (i I39) Clip(bits int) I39 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I39(uint64(i)&m^b) - I39(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I39) Bit(index int) I39 {
	if index < 0 {
		index = 39 + index
	}
	return (i >> I39(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I39) SetBit(index int, v I39) {
	if index < 0 {
		index = 39 + index
	}
	bit := I39(1) << I39(index)
	*i = I39.cast((*i &^ bit) | ((v & 1) << I39(index)))
}

// Bitref returns a Range for bit index.
func (i *I39) Bitref(index int) Range[I39] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I39) Bits(lo, hi int) I39 {
	p := 39
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I39((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I39(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I39) SetBits(lo, hi int, v I39) {
	p := 39
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I39((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I39.cast((*i &^ mask) | ((v << I39(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I39) Bitsref(lo, hi int) Range[I39] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I39) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I39) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I39(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I39) Byteref(idx int) Range[I39] { return byte(i, idx) }

// U40 is an 40-bit unsigned integer.
type U40 uint64

// AsU40 returns a U40 representing v reduced to 40 bits.
func AsU40[T Unsigned](v T) U40 { return U40.cast(U40(v)) }

func (U40) nbits() int  { return 40 }
func (u U40) mask() U40 { return 0xffffffffff }
func (u U40) cast() U40 { return u & u.mask() }

// Set assigns v reduced to 40 bits to u.
func (u *U40) Set(v U40) U40 { *u = new(U40(v)).cast(); return U40.cast(v) }

// Add returns u+o reduced to 40 bits.
func (u U40) Add(o U40) U40 { return U40.cast(u + o) }

// Sub returns u-o reduced to 40 bits.
func (u U40) Sub(o U40) U40 { return U40.cast(u - o) }

// Inc returns u+1 reduced to 40 bits.
func (u U40) Inc() U40 { return U40.cast(u + 1) }

// Dec returns u-1 reduced to 40 bits.
func (u U40) Dec() U40 { return U40.cast(u - 1) }

// Mul returns u*o reduced to 40 bits.
func (u U40) Mul(o U40) U40 { return U40.cast(u * o) }

// Div returns u divided by o.
func (u U40) Div(o U40) U40 { return U40.cast(u / o) }

// Mod returns u modulo o.
func (u U40) Mod(o U40) U40 { return U40.cast(u % o) }

// And returns u&o reduced to 40 bits.
func (u U40) And(o U40) U40 { return U40.cast(u & o) }

// Or returns u|o reduced to 40 bits.
func (u U40) Or(o U40) U40 { return U40.cast(u | o) }

// Xor returns u^o reduced to 40 bits.
func (u U40) Xor(o U40) U40 { return U40.cast(u ^ o) }

// Shr returns u>>o reduced to 40 bits.
func (u U40) Shr(o U40) U40 { return U40.cast(u >> o) }

// Shl returns u<<o reduced to 40 bits.
func (u U40) Shl(o U40) U40 { return U40.cast(u << o) }

// Not returns bitwise complement of u reduced to 40 bits.
func (u U40) Not() U40 { return U40.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U40) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U40) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U40 value.
func (u U40) Signed() I40 { return I40.cast(I40(u)) }

// Clamp returns value saturated into the representable range of U40.
func (u U40) Clamp(value uint64) U40 { return uclamp[U40](value) }

// Clip masks u to a bits-wide low field.
func (u U40) Clip(bits int) U40 {
	b := 1 << (bits - 1)
	m := U40(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U40) Bit(index int) U40 {
	if index < 0 {
		index = 40 + index
	}
	return (u >> U40(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U40) SetBit(index int, v U40) {
	if index < 0 {
		index = 40 + index
	}
	bit := U40(1) << U40(index)
	*u = U40.cast((*u &^ bit) | ((v & 1) << U40(index)))
}

// Bitref returns a Range for bit index.
func (u *U40) Bitref(index int) Range[U40] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U40) Bits(lo, hi int) U40 {
	p := 40
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U40((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U40(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U40) SetBits(lo, hi int, v U40) {
	p := 40
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U40((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U40.cast((*u &^ mask) | ((v << U40(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U40) Bitsref(lo, hi int) Range[U40] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U40) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U40) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U40(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U40) Byteref(idx int) Range[U40] { return byte(u, idx) }

// I40 is an 40-bit signed integer in two's complement.
type I40 int64

// AsI40 returns a I40 representing v reduced to 40 bits.
func AsI40[T Signed](v T) I40 { return I40.cast(I40(v)) }

func (I40) nbits() int  { return 40 }
func (i I40) mask() I40 { return 0xffffffffff }
func (i I40) sign() I40 { return 1 << (i.nbits() - 1) }
func (i I40) cast() I40 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 40 bits to u.
func (i *I40) Set(v I40) I40 { *i = new(I40(v)).cast(); return I40.cast(v) }

// Add returns i+o reduced to 40 bits.
func (i I40) Add(o I40) I40 { return I40.cast(i + o) }

// Sub returns i-o reduced to 40 bits.
func (i I40) Sub(o I40) I40 { return I40.cast(i - o) }

// Inc returns i+1 reduced to 40 bits.
func (i I40) Inc() I40 { return I40.cast(i + 1) }

// Dec returns i-1 reduced to 40 bits.
func (i I40) Dec() I40 { return I40.cast(i - 1) }

// Mul returns i*o reduced to 40 bits.
func (i I40) Mul(o I40) I40 { return I40.cast(i * o) }

// Div returns i divided by o.
func (i I40) Div(o I40) I40 { return I40.cast(i / o) }

// Mod returns i modulo o.
func (i I40) Mod(o I40) I40 { return I40.cast(i % o) }

// And returns i&o reduced to 40 bits.
func (i I40) And(o I40) I40 { return I40.cast(i & o) }

// Or returns i|o reduced to 40 bits.
func (i I40) Or(o I40) I40 { return I40.cast(i | o) }

// Xor returns i^o reduced to 40 bits.
func (i I40) Xor(o I40) I40 { return I40.cast(i ^ o) }

// Shr returns i>>o reduced to 40 bits.
func (i I40) Shr(o I40) I40 { return I40.cast(i >> o) }

// Shl returns i<<o reduced to 40 bits.
func (i I40) Shl(o I40) I40 { return I40.cast(i << o) }

// Not returns bitwise complement of i reduced to 40 bits.
func (i I40) Not() I40 { return I40.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I40) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I40) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U40 bit pattern.
func (i I40) Unsigned() U40 { return U40.cast(U40(i)) }

// Clamp returns value saturated into the representable range of I40.
func (i I40) Clamp(value int64) I40 { return iclamp[I40](value) }

// Clip masks i to a bits-wide low field and sign-extends to 40 bits.
func (i I40) Clip(bits int) I40 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I40(uint64(i)&m^b) - I40(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I40) Bit(index int) I40 {
	if index < 0 {
		index = 40 + index
	}
	return (i >> I40(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I40) SetBit(index int, v I40) {
	if index < 0 {
		index = 40 + index
	}
	bit := I40(1) << I40(index)
	*i = I40.cast((*i &^ bit) | ((v & 1) << I40(index)))
}

// Bitref returns a Range for bit index.
func (i *I40) Bitref(index int) Range[I40] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I40) Bits(lo, hi int) I40 {
	p := 40
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I40((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I40(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I40) SetBits(lo, hi int, v I40) {
	p := 40
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I40((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I40.cast((*i &^ mask) | ((v << I40(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I40) Bitsref(lo, hi int) Range[I40] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I40) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I40) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I40(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I40) Byteref(idx int) Range[I40] { return byte(i, idx) }

// U41 is an 41-bit unsigned integer.
type U41 uint64

// AsU41 returns a U41 representing v reduced to 41 bits.
func AsU41[T Unsigned](v T) U41 { return U41.cast(U41(v)) }

func (U41) nbits() int  { return 41 }
func (u U41) mask() U41 { return 0x1ffffffffff }
func (u U41) cast() U41 { return u & u.mask() }

// Set assigns v reduced to 41 bits to u.
func (u *U41) Set(v U41) U41 { *u = new(U41(v)).cast(); return U41.cast(v) }

// Add returns u+o reduced to 41 bits.
func (u U41) Add(o U41) U41 { return U41.cast(u + o) }

// Sub returns u-o reduced to 41 bits.
func (u U41) Sub(o U41) U41 { return U41.cast(u - o) }

// Inc returns u+1 reduced to 41 bits.
func (u U41) Inc() U41 { return U41.cast(u + 1) }

// Dec returns u-1 reduced to 41 bits.
func (u U41) Dec() U41 { return U41.cast(u - 1) }

// Mul returns u*o reduced to 41 bits.
func (u U41) Mul(o U41) U41 { return U41.cast(u * o) }

// Div returns u divided by o.
func (u U41) Div(o U41) U41 { return U41.cast(u / o) }

// Mod returns u modulo o.
func (u U41) Mod(o U41) U41 { return U41.cast(u % o) }

// And returns u&o reduced to 41 bits.
func (u U41) And(o U41) U41 { return U41.cast(u & o) }

// Or returns u|o reduced to 41 bits.
func (u U41) Or(o U41) U41 { return U41.cast(u | o) }

// Xor returns u^o reduced to 41 bits.
func (u U41) Xor(o U41) U41 { return U41.cast(u ^ o) }

// Shr returns u>>o reduced to 41 bits.
func (u U41) Shr(o U41) U41 { return U41.cast(u >> o) }

// Shl returns u<<o reduced to 41 bits.
func (u U41) Shl(o U41) U41 { return U41.cast(u << o) }

// Not returns bitwise complement of u reduced to 41 bits.
func (u U41) Not() U41 { return U41.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U41) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U41) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U41 value.
func (u U41) Signed() I41 { return I41.cast(I41(u)) }

// Clamp returns value saturated into the representable range of U41.
func (u U41) Clamp(value uint64) U41 { return uclamp[U41](value) }

// Clip masks u to a bits-wide low field.
func (u U41) Clip(bits int) U41 {
	b := 1 << (bits - 1)
	m := U41(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U41) Bit(index int) U41 {
	if index < 0 {
		index = 41 + index
	}
	return (u >> U41(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U41) SetBit(index int, v U41) {
	if index < 0 {
		index = 41 + index
	}
	bit := U41(1) << U41(index)
	*u = U41.cast((*u &^ bit) | ((v & 1) << U41(index)))
}

// Bitref returns a Range for bit index.
func (u *U41) Bitref(index int) Range[U41] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U41) Bits(lo, hi int) U41 {
	p := 41
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U41((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U41(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U41) SetBits(lo, hi int, v U41) {
	p := 41
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U41((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U41.cast((*u &^ mask) | ((v << U41(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U41) Bitsref(lo, hi int) Range[U41] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U41) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U41) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U41(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U41) Byteref(idx int) Range[U41] { return byte(u, idx) }

// I41 is an 41-bit signed integer in two's complement.
type I41 int64

// AsI41 returns a I41 representing v reduced to 41 bits.
func AsI41[T Signed](v T) I41 { return I41.cast(I41(v)) }

func (I41) nbits() int  { return 41 }
func (i I41) mask() I41 { return 0x1ffffffffff }
func (i I41) sign() I41 { return 1 << (i.nbits() - 1) }
func (i I41) cast() I41 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 41 bits to u.
func (i *I41) Set(v I41) I41 { *i = new(I41(v)).cast(); return I41.cast(v) }

// Add returns i+o reduced to 41 bits.
func (i I41) Add(o I41) I41 { return I41.cast(i + o) }

// Sub returns i-o reduced to 41 bits.
func (i I41) Sub(o I41) I41 { return I41.cast(i - o) }

// Inc returns i+1 reduced to 41 bits.
func (i I41) Inc() I41 { return I41.cast(i + 1) }

// Dec returns i-1 reduced to 41 bits.
func (i I41) Dec() I41 { return I41.cast(i - 1) }

// Mul returns i*o reduced to 41 bits.
func (i I41) Mul(o I41) I41 { return I41.cast(i * o) }

// Div returns i divided by o.
func (i I41) Div(o I41) I41 { return I41.cast(i / o) }

// Mod returns i modulo o.
func (i I41) Mod(o I41) I41 { return I41.cast(i % o) }

// And returns i&o reduced to 41 bits.
func (i I41) And(o I41) I41 { return I41.cast(i & o) }

// Or returns i|o reduced to 41 bits.
func (i I41) Or(o I41) I41 { return I41.cast(i | o) }

// Xor returns i^o reduced to 41 bits.
func (i I41) Xor(o I41) I41 { return I41.cast(i ^ o) }

// Shr returns i>>o reduced to 41 bits.
func (i I41) Shr(o I41) I41 { return I41.cast(i >> o) }

// Shl returns i<<o reduced to 41 bits.
func (i I41) Shl(o I41) I41 { return I41.cast(i << o) }

// Not returns bitwise complement of i reduced to 41 bits.
func (i I41) Not() I41 { return I41.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I41) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I41) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U41 bit pattern.
func (i I41) Unsigned() U41 { return U41.cast(U41(i)) }

// Clamp returns value saturated into the representable range of I41.
func (i I41) Clamp(value int64) I41 { return iclamp[I41](value) }

// Clip masks i to a bits-wide low field and sign-extends to 41 bits.
func (i I41) Clip(bits int) I41 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I41(uint64(i)&m^b) - I41(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I41) Bit(index int) I41 {
	if index < 0 {
		index = 41 + index
	}
	return (i >> I41(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I41) SetBit(index int, v I41) {
	if index < 0 {
		index = 41 + index
	}
	bit := I41(1) << I41(index)
	*i = I41.cast((*i &^ bit) | ((v & 1) << I41(index)))
}

// Bitref returns a Range for bit index.
func (i *I41) Bitref(index int) Range[I41] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I41) Bits(lo, hi int) I41 {
	p := 41
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I41((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I41(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I41) SetBits(lo, hi int, v I41) {
	p := 41
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I41((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I41.cast((*i &^ mask) | ((v << I41(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I41) Bitsref(lo, hi int) Range[I41] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I41) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I41) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I41(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I41) Byteref(idx int) Range[I41] { return byte(i, idx) }

// U42 is an 42-bit unsigned integer.
type U42 uint64

// AsU42 returns a U42 representing v reduced to 42 bits.
func AsU42[T Unsigned](v T) U42 { return U42.cast(U42(v)) }

func (U42) nbits() int  { return 42 }
func (u U42) mask() U42 { return 0x3ffffffffff }
func (u U42) cast() U42 { return u & u.mask() }

// Set assigns v reduced to 42 bits to u.
func (u *U42) Set(v U42) U42 { *u = new(U42(v)).cast(); return U42.cast(v) }

// Add returns u+o reduced to 42 bits.
func (u U42) Add(o U42) U42 { return U42.cast(u + o) }

// Sub returns u-o reduced to 42 bits.
func (u U42) Sub(o U42) U42 { return U42.cast(u - o) }

// Inc returns u+1 reduced to 42 bits.
func (u U42) Inc() U42 { return U42.cast(u + 1) }

// Dec returns u-1 reduced to 42 bits.
func (u U42) Dec() U42 { return U42.cast(u - 1) }

// Mul returns u*o reduced to 42 bits.
func (u U42) Mul(o U42) U42 { return U42.cast(u * o) }

// Div returns u divided by o.
func (u U42) Div(o U42) U42 { return U42.cast(u / o) }

// Mod returns u modulo o.
func (u U42) Mod(o U42) U42 { return U42.cast(u % o) }

// And returns u&o reduced to 42 bits.
func (u U42) And(o U42) U42 { return U42.cast(u & o) }

// Or returns u|o reduced to 42 bits.
func (u U42) Or(o U42) U42 { return U42.cast(u | o) }

// Xor returns u^o reduced to 42 bits.
func (u U42) Xor(o U42) U42 { return U42.cast(u ^ o) }

// Shr returns u>>o reduced to 42 bits.
func (u U42) Shr(o U42) U42 { return U42.cast(u >> o) }

// Shl returns u<<o reduced to 42 bits.
func (u U42) Shl(o U42) U42 { return U42.cast(u << o) }

// Not returns bitwise complement of u reduced to 42 bits.
func (u U42) Not() U42 { return U42.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U42) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U42) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U42 value.
func (u U42) Signed() I42 { return I42.cast(I42(u)) }

// Clamp returns value saturated into the representable range of U42.
func (u U42) Clamp(value uint64) U42 { return uclamp[U42](value) }

// Clip masks u to a bits-wide low field.
func (u U42) Clip(bits int) U42 {
	b := 1 << (bits - 1)
	m := U42(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U42) Bit(index int) U42 {
	if index < 0 {
		index = 42 + index
	}
	return (u >> U42(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U42) SetBit(index int, v U42) {
	if index < 0 {
		index = 42 + index
	}
	bit := U42(1) << U42(index)
	*u = U42.cast((*u &^ bit) | ((v & 1) << U42(index)))
}

// Bitref returns a Range for bit index.
func (u *U42) Bitref(index int) Range[U42] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U42) Bits(lo, hi int) U42 {
	p := 42
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U42((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U42(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U42) SetBits(lo, hi int, v U42) {
	p := 42
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U42((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U42.cast((*u &^ mask) | ((v << U42(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U42) Bitsref(lo, hi int) Range[U42] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U42) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U42) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U42(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U42) Byteref(idx int) Range[U42] { return byte(u, idx) }

// I42 is an 42-bit signed integer in two's complement.
type I42 int64

// AsI42 returns a I42 representing v reduced to 42 bits.
func AsI42[T Signed](v T) I42 { return I42.cast(I42(v)) }

func (I42) nbits() int  { return 42 }
func (i I42) mask() I42 { return 0x3ffffffffff }
func (i I42) sign() I42 { return 1 << (i.nbits() - 1) }
func (i I42) cast() I42 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 42 bits to u.
func (i *I42) Set(v I42) I42 { *i = new(I42(v)).cast(); return I42.cast(v) }

// Add returns i+o reduced to 42 bits.
func (i I42) Add(o I42) I42 { return I42.cast(i + o) }

// Sub returns i-o reduced to 42 bits.
func (i I42) Sub(o I42) I42 { return I42.cast(i - o) }

// Inc returns i+1 reduced to 42 bits.
func (i I42) Inc() I42 { return I42.cast(i + 1) }

// Dec returns i-1 reduced to 42 bits.
func (i I42) Dec() I42 { return I42.cast(i - 1) }

// Mul returns i*o reduced to 42 bits.
func (i I42) Mul(o I42) I42 { return I42.cast(i * o) }

// Div returns i divided by o.
func (i I42) Div(o I42) I42 { return I42.cast(i / o) }

// Mod returns i modulo o.
func (i I42) Mod(o I42) I42 { return I42.cast(i % o) }

// And returns i&o reduced to 42 bits.
func (i I42) And(o I42) I42 { return I42.cast(i & o) }

// Or returns i|o reduced to 42 bits.
func (i I42) Or(o I42) I42 { return I42.cast(i | o) }

// Xor returns i^o reduced to 42 bits.
func (i I42) Xor(o I42) I42 { return I42.cast(i ^ o) }

// Shr returns i>>o reduced to 42 bits.
func (i I42) Shr(o I42) I42 { return I42.cast(i >> o) }

// Shl returns i<<o reduced to 42 bits.
func (i I42) Shl(o I42) I42 { return I42.cast(i << o) }

// Not returns bitwise complement of i reduced to 42 bits.
func (i I42) Not() I42 { return I42.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I42) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I42) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U42 bit pattern.
func (i I42) Unsigned() U42 { return U42.cast(U42(i)) }

// Clamp returns value saturated into the representable range of I42.
func (i I42) Clamp(value int64) I42 { return iclamp[I42](value) }

// Clip masks i to a bits-wide low field and sign-extends to 42 bits.
func (i I42) Clip(bits int) I42 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I42(uint64(i)&m^b) - I42(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I42) Bit(index int) I42 {
	if index < 0 {
		index = 42 + index
	}
	return (i >> I42(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I42) SetBit(index int, v I42) {
	if index < 0 {
		index = 42 + index
	}
	bit := I42(1) << I42(index)
	*i = I42.cast((*i &^ bit) | ((v & 1) << I42(index)))
}

// Bitref returns a Range for bit index.
func (i *I42) Bitref(index int) Range[I42] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I42) Bits(lo, hi int) I42 {
	p := 42
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I42((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I42(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I42) SetBits(lo, hi int, v I42) {
	p := 42
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I42((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I42.cast((*i &^ mask) | ((v << I42(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I42) Bitsref(lo, hi int) Range[I42] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I42) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I42) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I42(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I42) Byteref(idx int) Range[I42] { return byte(i, idx) }

// U43 is an 43-bit unsigned integer.
type U43 uint64

// AsU43 returns a U43 representing v reduced to 43 bits.
func AsU43[T Unsigned](v T) U43 { return U43.cast(U43(v)) }

func (U43) nbits() int  { return 43 }
func (u U43) mask() U43 { return 0x7ffffffffff }
func (u U43) cast() U43 { return u & u.mask() }

// Set assigns v reduced to 43 bits to u.
func (u *U43) Set(v U43) U43 { *u = new(U43(v)).cast(); return U43.cast(v) }

// Add returns u+o reduced to 43 bits.
func (u U43) Add(o U43) U43 { return U43.cast(u + o) }

// Sub returns u-o reduced to 43 bits.
func (u U43) Sub(o U43) U43 { return U43.cast(u - o) }

// Inc returns u+1 reduced to 43 bits.
func (u U43) Inc() U43 { return U43.cast(u + 1) }

// Dec returns u-1 reduced to 43 bits.
func (u U43) Dec() U43 { return U43.cast(u - 1) }

// Mul returns u*o reduced to 43 bits.
func (u U43) Mul(o U43) U43 { return U43.cast(u * o) }

// Div returns u divided by o.
func (u U43) Div(o U43) U43 { return U43.cast(u / o) }

// Mod returns u modulo o.
func (u U43) Mod(o U43) U43 { return U43.cast(u % o) }

// And returns u&o reduced to 43 bits.
func (u U43) And(o U43) U43 { return U43.cast(u & o) }

// Or returns u|o reduced to 43 bits.
func (u U43) Or(o U43) U43 { return U43.cast(u | o) }

// Xor returns u^o reduced to 43 bits.
func (u U43) Xor(o U43) U43 { return U43.cast(u ^ o) }

// Shr returns u>>o reduced to 43 bits.
func (u U43) Shr(o U43) U43 { return U43.cast(u >> o) }

// Shl returns u<<o reduced to 43 bits.
func (u U43) Shl(o U43) U43 { return U43.cast(u << o) }

// Not returns bitwise complement of u reduced to 43 bits.
func (u U43) Not() U43 { return U43.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U43) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U43) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U43 value.
func (u U43) Signed() I43 { return I43.cast(I43(u)) }

// Clamp returns value saturated into the representable range of U43.
func (u U43) Clamp(value uint64) U43 { return uclamp[U43](value) }

// Clip masks u to a bits-wide low field.
func (u U43) Clip(bits int) U43 {
	b := 1 << (bits - 1)
	m := U43(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U43) Bit(index int) U43 {
	if index < 0 {
		index = 43 + index
	}
	return (u >> U43(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U43) SetBit(index int, v U43) {
	if index < 0 {
		index = 43 + index
	}
	bit := U43(1) << U43(index)
	*u = U43.cast((*u &^ bit) | ((v & 1) << U43(index)))
}

// Bitref returns a Range for bit index.
func (u *U43) Bitref(index int) Range[U43] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U43) Bits(lo, hi int) U43 {
	p := 43
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U43((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U43(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U43) SetBits(lo, hi int, v U43) {
	p := 43
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U43((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U43.cast((*u &^ mask) | ((v << U43(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U43) Bitsref(lo, hi int) Range[U43] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U43) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U43) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U43(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U43) Byteref(idx int) Range[U43] { return byte(u, idx) }

// I43 is an 43-bit signed integer in two's complement.
type I43 int64

// AsI43 returns a I43 representing v reduced to 43 bits.
func AsI43[T Signed](v T) I43 { return I43.cast(I43(v)) }

func (I43) nbits() int  { return 43 }
func (i I43) mask() I43 { return 0x7ffffffffff }
func (i I43) sign() I43 { return 1 << (i.nbits() - 1) }
func (i I43) cast() I43 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 43 bits to u.
func (i *I43) Set(v I43) I43 { *i = new(I43(v)).cast(); return I43.cast(v) }

// Add returns i+o reduced to 43 bits.
func (i I43) Add(o I43) I43 { return I43.cast(i + o) }

// Sub returns i-o reduced to 43 bits.
func (i I43) Sub(o I43) I43 { return I43.cast(i - o) }

// Inc returns i+1 reduced to 43 bits.
func (i I43) Inc() I43 { return I43.cast(i + 1) }

// Dec returns i-1 reduced to 43 bits.
func (i I43) Dec() I43 { return I43.cast(i - 1) }

// Mul returns i*o reduced to 43 bits.
func (i I43) Mul(o I43) I43 { return I43.cast(i * o) }

// Div returns i divided by o.
func (i I43) Div(o I43) I43 { return I43.cast(i / o) }

// Mod returns i modulo o.
func (i I43) Mod(o I43) I43 { return I43.cast(i % o) }

// And returns i&o reduced to 43 bits.
func (i I43) And(o I43) I43 { return I43.cast(i & o) }

// Or returns i|o reduced to 43 bits.
func (i I43) Or(o I43) I43 { return I43.cast(i | o) }

// Xor returns i^o reduced to 43 bits.
func (i I43) Xor(o I43) I43 { return I43.cast(i ^ o) }

// Shr returns i>>o reduced to 43 bits.
func (i I43) Shr(o I43) I43 { return I43.cast(i >> o) }

// Shl returns i<<o reduced to 43 bits.
func (i I43) Shl(o I43) I43 { return I43.cast(i << o) }

// Not returns bitwise complement of i reduced to 43 bits.
func (i I43) Not() I43 { return I43.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I43) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I43) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U43 bit pattern.
func (i I43) Unsigned() U43 { return U43.cast(U43(i)) }

// Clamp returns value saturated into the representable range of I43.
func (i I43) Clamp(value int64) I43 { return iclamp[I43](value) }

// Clip masks i to a bits-wide low field and sign-extends to 43 bits.
func (i I43) Clip(bits int) I43 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I43(uint64(i)&m^b) - I43(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I43) Bit(index int) I43 {
	if index < 0 {
		index = 43 + index
	}
	return (i >> I43(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I43) SetBit(index int, v I43) {
	if index < 0 {
		index = 43 + index
	}
	bit := I43(1) << I43(index)
	*i = I43.cast((*i &^ bit) | ((v & 1) << I43(index)))
}

// Bitref returns a Range for bit index.
func (i *I43) Bitref(index int) Range[I43] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I43) Bits(lo, hi int) I43 {
	p := 43
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I43((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I43(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I43) SetBits(lo, hi int, v I43) {
	p := 43
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I43((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I43.cast((*i &^ mask) | ((v << I43(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I43) Bitsref(lo, hi int) Range[I43] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I43) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I43) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I43(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I43) Byteref(idx int) Range[I43] { return byte(i, idx) }

// U44 is an 44-bit unsigned integer.
type U44 uint64

// AsU44 returns a U44 representing v reduced to 44 bits.
func AsU44[T Unsigned](v T) U44 { return U44.cast(U44(v)) }

func (U44) nbits() int  { return 44 }
func (u U44) mask() U44 { return 0xfffffffffff }
func (u U44) cast() U44 { return u & u.mask() }

// Set assigns v reduced to 44 bits to u.
func (u *U44) Set(v U44) U44 { *u = new(U44(v)).cast(); return U44.cast(v) }

// Add returns u+o reduced to 44 bits.
func (u U44) Add(o U44) U44 { return U44.cast(u + o) }

// Sub returns u-o reduced to 44 bits.
func (u U44) Sub(o U44) U44 { return U44.cast(u - o) }

// Inc returns u+1 reduced to 44 bits.
func (u U44) Inc() U44 { return U44.cast(u + 1) }

// Dec returns u-1 reduced to 44 bits.
func (u U44) Dec() U44 { return U44.cast(u - 1) }

// Mul returns u*o reduced to 44 bits.
func (u U44) Mul(o U44) U44 { return U44.cast(u * o) }

// Div returns u divided by o.
func (u U44) Div(o U44) U44 { return U44.cast(u / o) }

// Mod returns u modulo o.
func (u U44) Mod(o U44) U44 { return U44.cast(u % o) }

// And returns u&o reduced to 44 bits.
func (u U44) And(o U44) U44 { return U44.cast(u & o) }

// Or returns u|o reduced to 44 bits.
func (u U44) Or(o U44) U44 { return U44.cast(u | o) }

// Xor returns u^o reduced to 44 bits.
func (u U44) Xor(o U44) U44 { return U44.cast(u ^ o) }

// Shr returns u>>o reduced to 44 bits.
func (u U44) Shr(o U44) U44 { return U44.cast(u >> o) }

// Shl returns u<<o reduced to 44 bits.
func (u U44) Shl(o U44) U44 { return U44.cast(u << o) }

// Not returns bitwise complement of u reduced to 44 bits.
func (u U44) Not() U44 { return U44.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U44) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U44) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U44 value.
func (u U44) Signed() I44 { return I44.cast(I44(u)) }

// Clamp returns value saturated into the representable range of U44.
func (u U44) Clamp(value uint64) U44 { return uclamp[U44](value) }

// Clip masks u to a bits-wide low field.
func (u U44) Clip(bits int) U44 {
	b := 1 << (bits - 1)
	m := U44(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U44) Bit(index int) U44 {
	if index < 0 {
		index = 44 + index
	}
	return (u >> U44(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U44) SetBit(index int, v U44) {
	if index < 0 {
		index = 44 + index
	}
	bit := U44(1) << U44(index)
	*u = U44.cast((*u &^ bit) | ((v & 1) << U44(index)))
}

// Bitref returns a Range for bit index.
func (u *U44) Bitref(index int) Range[U44] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U44) Bits(lo, hi int) U44 {
	p := 44
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U44((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U44(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U44) SetBits(lo, hi int, v U44) {
	p := 44
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U44((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U44.cast((*u &^ mask) | ((v << U44(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U44) Bitsref(lo, hi int) Range[U44] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U44) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U44) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U44(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U44) Byteref(idx int) Range[U44] { return byte(u, idx) }

// I44 is an 44-bit signed integer in two's complement.
type I44 int64

// AsI44 returns a I44 representing v reduced to 44 bits.
func AsI44[T Signed](v T) I44 { return I44.cast(I44(v)) }

func (I44) nbits() int  { return 44 }
func (i I44) mask() I44 { return 0xfffffffffff }
func (i I44) sign() I44 { return 1 << (i.nbits() - 1) }
func (i I44) cast() I44 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 44 bits to u.
func (i *I44) Set(v I44) I44 { *i = new(I44(v)).cast(); return I44.cast(v) }

// Add returns i+o reduced to 44 bits.
func (i I44) Add(o I44) I44 { return I44.cast(i + o) }

// Sub returns i-o reduced to 44 bits.
func (i I44) Sub(o I44) I44 { return I44.cast(i - o) }

// Inc returns i+1 reduced to 44 bits.
func (i I44) Inc() I44 { return I44.cast(i + 1) }

// Dec returns i-1 reduced to 44 bits.
func (i I44) Dec() I44 { return I44.cast(i - 1) }

// Mul returns i*o reduced to 44 bits.
func (i I44) Mul(o I44) I44 { return I44.cast(i * o) }

// Div returns i divided by o.
func (i I44) Div(o I44) I44 { return I44.cast(i / o) }

// Mod returns i modulo o.
func (i I44) Mod(o I44) I44 { return I44.cast(i % o) }

// And returns i&o reduced to 44 bits.
func (i I44) And(o I44) I44 { return I44.cast(i & o) }

// Or returns i|o reduced to 44 bits.
func (i I44) Or(o I44) I44 { return I44.cast(i | o) }

// Xor returns i^o reduced to 44 bits.
func (i I44) Xor(o I44) I44 { return I44.cast(i ^ o) }

// Shr returns i>>o reduced to 44 bits.
func (i I44) Shr(o I44) I44 { return I44.cast(i >> o) }

// Shl returns i<<o reduced to 44 bits.
func (i I44) Shl(o I44) I44 { return I44.cast(i << o) }

// Not returns bitwise complement of i reduced to 44 bits.
func (i I44) Not() I44 { return I44.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I44) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I44) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U44 bit pattern.
func (i I44) Unsigned() U44 { return U44.cast(U44(i)) }

// Clamp returns value saturated into the representable range of I44.
func (i I44) Clamp(value int64) I44 { return iclamp[I44](value) }

// Clip masks i to a bits-wide low field and sign-extends to 44 bits.
func (i I44) Clip(bits int) I44 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I44(uint64(i)&m^b) - I44(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I44) Bit(index int) I44 {
	if index < 0 {
		index = 44 + index
	}
	return (i >> I44(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I44) SetBit(index int, v I44) {
	if index < 0 {
		index = 44 + index
	}
	bit := I44(1) << I44(index)
	*i = I44.cast((*i &^ bit) | ((v & 1) << I44(index)))
}

// Bitref returns a Range for bit index.
func (i *I44) Bitref(index int) Range[I44] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I44) Bits(lo, hi int) I44 {
	p := 44
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I44((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I44(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I44) SetBits(lo, hi int, v I44) {
	p := 44
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I44((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I44.cast((*i &^ mask) | ((v << I44(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I44) Bitsref(lo, hi int) Range[I44] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I44) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I44) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I44(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I44) Byteref(idx int) Range[I44] { return byte(i, idx) }

// U45 is an 45-bit unsigned integer.
type U45 uint64

// AsU45 returns a U45 representing v reduced to 45 bits.
func AsU45[T Unsigned](v T) U45 { return U45.cast(U45(v)) }

func (U45) nbits() int  { return 45 }
func (u U45) mask() U45 { return 0x1fffffffffff }
func (u U45) cast() U45 { return u & u.mask() }

// Set assigns v reduced to 45 bits to u.
func (u *U45) Set(v U45) U45 { *u = new(U45(v)).cast(); return U45.cast(v) }

// Add returns u+o reduced to 45 bits.
func (u U45) Add(o U45) U45 { return U45.cast(u + o) }

// Sub returns u-o reduced to 45 bits.
func (u U45) Sub(o U45) U45 { return U45.cast(u - o) }

// Inc returns u+1 reduced to 45 bits.
func (u U45) Inc() U45 { return U45.cast(u + 1) }

// Dec returns u-1 reduced to 45 bits.
func (u U45) Dec() U45 { return U45.cast(u - 1) }

// Mul returns u*o reduced to 45 bits.
func (u U45) Mul(o U45) U45 { return U45.cast(u * o) }

// Div returns u divided by o.
func (u U45) Div(o U45) U45 { return U45.cast(u / o) }

// Mod returns u modulo o.
func (u U45) Mod(o U45) U45 { return U45.cast(u % o) }

// And returns u&o reduced to 45 bits.
func (u U45) And(o U45) U45 { return U45.cast(u & o) }

// Or returns u|o reduced to 45 bits.
func (u U45) Or(o U45) U45 { return U45.cast(u | o) }

// Xor returns u^o reduced to 45 bits.
func (u U45) Xor(o U45) U45 { return U45.cast(u ^ o) }

// Shr returns u>>o reduced to 45 bits.
func (u U45) Shr(o U45) U45 { return U45.cast(u >> o) }

// Shl returns u<<o reduced to 45 bits.
func (u U45) Shl(o U45) U45 { return U45.cast(u << o) }

// Not returns bitwise complement of u reduced to 45 bits.
func (u U45) Not() U45 { return U45.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U45) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U45) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U45 value.
func (u U45) Signed() I45 { return I45.cast(I45(u)) }

// Clamp returns value saturated into the representable range of U45.
func (u U45) Clamp(value uint64) U45 { return uclamp[U45](value) }

// Clip masks u to a bits-wide low field.
func (u U45) Clip(bits int) U45 {
	b := 1 << (bits - 1)
	m := U45(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U45) Bit(index int) U45 {
	if index < 0 {
		index = 45 + index
	}
	return (u >> U45(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U45) SetBit(index int, v U45) {
	if index < 0 {
		index = 45 + index
	}
	bit := U45(1) << U45(index)
	*u = U45.cast((*u &^ bit) | ((v & 1) << U45(index)))
}

// Bitref returns a Range for bit index.
func (u *U45) Bitref(index int) Range[U45] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U45) Bits(lo, hi int) U45 {
	p := 45
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U45((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U45(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U45) SetBits(lo, hi int, v U45) {
	p := 45
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U45((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U45.cast((*u &^ mask) | ((v << U45(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U45) Bitsref(lo, hi int) Range[U45] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U45) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U45) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U45(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U45) Byteref(idx int) Range[U45] { return byte(u, idx) }

// I45 is an 45-bit signed integer in two's complement.
type I45 int64

// AsI45 returns a I45 representing v reduced to 45 bits.
func AsI45[T Signed](v T) I45 { return I45.cast(I45(v)) }

func (I45) nbits() int  { return 45 }
func (i I45) mask() I45 { return 0x1fffffffffff }
func (i I45) sign() I45 { return 1 << (i.nbits() - 1) }
func (i I45) cast() I45 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 45 bits to u.
func (i *I45) Set(v I45) I45 { *i = new(I45(v)).cast(); return I45.cast(v) }

// Add returns i+o reduced to 45 bits.
func (i I45) Add(o I45) I45 { return I45.cast(i + o) }

// Sub returns i-o reduced to 45 bits.
func (i I45) Sub(o I45) I45 { return I45.cast(i - o) }

// Inc returns i+1 reduced to 45 bits.
func (i I45) Inc() I45 { return I45.cast(i + 1) }

// Dec returns i-1 reduced to 45 bits.
func (i I45) Dec() I45 { return I45.cast(i - 1) }

// Mul returns i*o reduced to 45 bits.
func (i I45) Mul(o I45) I45 { return I45.cast(i * o) }

// Div returns i divided by o.
func (i I45) Div(o I45) I45 { return I45.cast(i / o) }

// Mod returns i modulo o.
func (i I45) Mod(o I45) I45 { return I45.cast(i % o) }

// And returns i&o reduced to 45 bits.
func (i I45) And(o I45) I45 { return I45.cast(i & o) }

// Or returns i|o reduced to 45 bits.
func (i I45) Or(o I45) I45 { return I45.cast(i | o) }

// Xor returns i^o reduced to 45 bits.
func (i I45) Xor(o I45) I45 { return I45.cast(i ^ o) }

// Shr returns i>>o reduced to 45 bits.
func (i I45) Shr(o I45) I45 { return I45.cast(i >> o) }

// Shl returns i<<o reduced to 45 bits.
func (i I45) Shl(o I45) I45 { return I45.cast(i << o) }

// Not returns bitwise complement of i reduced to 45 bits.
func (i I45) Not() I45 { return I45.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I45) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I45) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U45 bit pattern.
func (i I45) Unsigned() U45 { return U45.cast(U45(i)) }

// Clamp returns value saturated into the representable range of I45.
func (i I45) Clamp(value int64) I45 { return iclamp[I45](value) }

// Clip masks i to a bits-wide low field and sign-extends to 45 bits.
func (i I45) Clip(bits int) I45 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I45(uint64(i)&m^b) - I45(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I45) Bit(index int) I45 {
	if index < 0 {
		index = 45 + index
	}
	return (i >> I45(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I45) SetBit(index int, v I45) {
	if index < 0 {
		index = 45 + index
	}
	bit := I45(1) << I45(index)
	*i = I45.cast((*i &^ bit) | ((v & 1) << I45(index)))
}

// Bitref returns a Range for bit index.
func (i *I45) Bitref(index int) Range[I45] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I45) Bits(lo, hi int) I45 {
	p := 45
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I45((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I45(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I45) SetBits(lo, hi int, v I45) {
	p := 45
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I45((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I45.cast((*i &^ mask) | ((v << I45(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I45) Bitsref(lo, hi int) Range[I45] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I45) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I45) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I45(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I45) Byteref(idx int) Range[I45] { return byte(i, idx) }

// U46 is an 46-bit unsigned integer.
type U46 uint64

// AsU46 returns a U46 representing v reduced to 46 bits.
func AsU46[T Unsigned](v T) U46 { return U46.cast(U46(v)) }

func (U46) nbits() int  { return 46 }
func (u U46) mask() U46 { return 0x3fffffffffff }
func (u U46) cast() U46 { return u & u.mask() }

// Set assigns v reduced to 46 bits to u.
func (u *U46) Set(v U46) U46 { *u = new(U46(v)).cast(); return U46.cast(v) }

// Add returns u+o reduced to 46 bits.
func (u U46) Add(o U46) U46 { return U46.cast(u + o) }

// Sub returns u-o reduced to 46 bits.
func (u U46) Sub(o U46) U46 { return U46.cast(u - o) }

// Inc returns u+1 reduced to 46 bits.
func (u U46) Inc() U46 { return U46.cast(u + 1) }

// Dec returns u-1 reduced to 46 bits.
func (u U46) Dec() U46 { return U46.cast(u - 1) }

// Mul returns u*o reduced to 46 bits.
func (u U46) Mul(o U46) U46 { return U46.cast(u * o) }

// Div returns u divided by o.
func (u U46) Div(o U46) U46 { return U46.cast(u / o) }

// Mod returns u modulo o.
func (u U46) Mod(o U46) U46 { return U46.cast(u % o) }

// And returns u&o reduced to 46 bits.
func (u U46) And(o U46) U46 { return U46.cast(u & o) }

// Or returns u|o reduced to 46 bits.
func (u U46) Or(o U46) U46 { return U46.cast(u | o) }

// Xor returns u^o reduced to 46 bits.
func (u U46) Xor(o U46) U46 { return U46.cast(u ^ o) }

// Shr returns u>>o reduced to 46 bits.
func (u U46) Shr(o U46) U46 { return U46.cast(u >> o) }

// Shl returns u<<o reduced to 46 bits.
func (u U46) Shl(o U46) U46 { return U46.cast(u << o) }

// Not returns bitwise complement of u reduced to 46 bits.
func (u U46) Not() U46 { return U46.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U46) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U46) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U46 value.
func (u U46) Signed() I46 { return I46.cast(I46(u)) }

// Clamp returns value saturated into the representable range of U46.
func (u U46) Clamp(value uint64) U46 { return uclamp[U46](value) }

// Clip masks u to a bits-wide low field.
func (u U46) Clip(bits int) U46 {
	b := 1 << (bits - 1)
	m := U46(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U46) Bit(index int) U46 {
	if index < 0 {
		index = 46 + index
	}
	return (u >> U46(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U46) SetBit(index int, v U46) {
	if index < 0 {
		index = 46 + index
	}
	bit := U46(1) << U46(index)
	*u = U46.cast((*u &^ bit) | ((v & 1) << U46(index)))
}

// Bitref returns a Range for bit index.
func (u *U46) Bitref(index int) Range[U46] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U46) Bits(lo, hi int) U46 {
	p := 46
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U46((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U46(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U46) SetBits(lo, hi int, v U46) {
	p := 46
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U46((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U46.cast((*u &^ mask) | ((v << U46(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U46) Bitsref(lo, hi int) Range[U46] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U46) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U46) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U46(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U46) Byteref(idx int) Range[U46] { return byte(u, idx) }

// I46 is an 46-bit signed integer in two's complement.
type I46 int64

// AsI46 returns a I46 representing v reduced to 46 bits.
func AsI46[T Signed](v T) I46 { return I46.cast(I46(v)) }

func (I46) nbits() int  { return 46 }
func (i I46) mask() I46 { return 0x3fffffffffff }
func (i I46) sign() I46 { return 1 << (i.nbits() - 1) }
func (i I46) cast() I46 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 46 bits to u.
func (i *I46) Set(v I46) I46 { *i = new(I46(v)).cast(); return I46.cast(v) }

// Add returns i+o reduced to 46 bits.
func (i I46) Add(o I46) I46 { return I46.cast(i + o) }

// Sub returns i-o reduced to 46 bits.
func (i I46) Sub(o I46) I46 { return I46.cast(i - o) }

// Inc returns i+1 reduced to 46 bits.
func (i I46) Inc() I46 { return I46.cast(i + 1) }

// Dec returns i-1 reduced to 46 bits.
func (i I46) Dec() I46 { return I46.cast(i - 1) }

// Mul returns i*o reduced to 46 bits.
func (i I46) Mul(o I46) I46 { return I46.cast(i * o) }

// Div returns i divided by o.
func (i I46) Div(o I46) I46 { return I46.cast(i / o) }

// Mod returns i modulo o.
func (i I46) Mod(o I46) I46 { return I46.cast(i % o) }

// And returns i&o reduced to 46 bits.
func (i I46) And(o I46) I46 { return I46.cast(i & o) }

// Or returns i|o reduced to 46 bits.
func (i I46) Or(o I46) I46 { return I46.cast(i | o) }

// Xor returns i^o reduced to 46 bits.
func (i I46) Xor(o I46) I46 { return I46.cast(i ^ o) }

// Shr returns i>>o reduced to 46 bits.
func (i I46) Shr(o I46) I46 { return I46.cast(i >> o) }

// Shl returns i<<o reduced to 46 bits.
func (i I46) Shl(o I46) I46 { return I46.cast(i << o) }

// Not returns bitwise complement of i reduced to 46 bits.
func (i I46) Not() I46 { return I46.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I46) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I46) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U46 bit pattern.
func (i I46) Unsigned() U46 { return U46.cast(U46(i)) }

// Clamp returns value saturated into the representable range of I46.
func (i I46) Clamp(value int64) I46 { return iclamp[I46](value) }

// Clip masks i to a bits-wide low field and sign-extends to 46 bits.
func (i I46) Clip(bits int) I46 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I46(uint64(i)&m^b) - I46(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I46) Bit(index int) I46 {
	if index < 0 {
		index = 46 + index
	}
	return (i >> I46(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I46) SetBit(index int, v I46) {
	if index < 0 {
		index = 46 + index
	}
	bit := I46(1) << I46(index)
	*i = I46.cast((*i &^ bit) | ((v & 1) << I46(index)))
}

// Bitref returns a Range for bit index.
func (i *I46) Bitref(index int) Range[I46] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I46) Bits(lo, hi int) I46 {
	p := 46
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I46((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I46(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I46) SetBits(lo, hi int, v I46) {
	p := 46
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I46((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I46.cast((*i &^ mask) | ((v << I46(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I46) Bitsref(lo, hi int) Range[I46] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I46) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I46) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I46(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I46) Byteref(idx int) Range[I46] { return byte(i, idx) }

// U47 is an 47-bit unsigned integer.
type U47 uint64

// AsU47 returns a U47 representing v reduced to 47 bits.
func AsU47[T Unsigned](v T) U47 { return U47.cast(U47(v)) }

func (U47) nbits() int  { return 47 }
func (u U47) mask() U47 { return 0x7fffffffffff }
func (u U47) cast() U47 { return u & u.mask() }

// Set assigns v reduced to 47 bits to u.
func (u *U47) Set(v U47) U47 { *u = new(U47(v)).cast(); return U47.cast(v) }

// Add returns u+o reduced to 47 bits.
func (u U47) Add(o U47) U47 { return U47.cast(u + o) }

// Sub returns u-o reduced to 47 bits.
func (u U47) Sub(o U47) U47 { return U47.cast(u - o) }

// Inc returns u+1 reduced to 47 bits.
func (u U47) Inc() U47 { return U47.cast(u + 1) }

// Dec returns u-1 reduced to 47 bits.
func (u U47) Dec() U47 { return U47.cast(u - 1) }

// Mul returns u*o reduced to 47 bits.
func (u U47) Mul(o U47) U47 { return U47.cast(u * o) }

// Div returns u divided by o.
func (u U47) Div(o U47) U47 { return U47.cast(u / o) }

// Mod returns u modulo o.
func (u U47) Mod(o U47) U47 { return U47.cast(u % o) }

// And returns u&o reduced to 47 bits.
func (u U47) And(o U47) U47 { return U47.cast(u & o) }

// Or returns u|o reduced to 47 bits.
func (u U47) Or(o U47) U47 { return U47.cast(u | o) }

// Xor returns u^o reduced to 47 bits.
func (u U47) Xor(o U47) U47 { return U47.cast(u ^ o) }

// Shr returns u>>o reduced to 47 bits.
func (u U47) Shr(o U47) U47 { return U47.cast(u >> o) }

// Shl returns u<<o reduced to 47 bits.
func (u U47) Shl(o U47) U47 { return U47.cast(u << o) }

// Not returns bitwise complement of u reduced to 47 bits.
func (u U47) Not() U47 { return U47.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U47) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U47) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U47 value.
func (u U47) Signed() I47 { return I47.cast(I47(u)) }

// Clamp returns value saturated into the representable range of U47.
func (u U47) Clamp(value uint64) U47 { return uclamp[U47](value) }

// Clip masks u to a bits-wide low field.
func (u U47) Clip(bits int) U47 {
	b := 1 << (bits - 1)
	m := U47(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U47) Bit(index int) U47 {
	if index < 0 {
		index = 47 + index
	}
	return (u >> U47(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U47) SetBit(index int, v U47) {
	if index < 0 {
		index = 47 + index
	}
	bit := U47(1) << U47(index)
	*u = U47.cast((*u &^ bit) | ((v & 1) << U47(index)))
}

// Bitref returns a Range for bit index.
func (u *U47) Bitref(index int) Range[U47] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U47) Bits(lo, hi int) U47 {
	p := 47
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U47((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U47(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U47) SetBits(lo, hi int, v U47) {
	p := 47
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U47((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U47.cast((*u &^ mask) | ((v << U47(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U47) Bitsref(lo, hi int) Range[U47] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U47) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U47) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U47(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U47) Byteref(idx int) Range[U47] { return byte(u, idx) }

// I47 is an 47-bit signed integer in two's complement.
type I47 int64

// AsI47 returns a I47 representing v reduced to 47 bits.
func AsI47[T Signed](v T) I47 { return I47.cast(I47(v)) }

func (I47) nbits() int  { return 47 }
func (i I47) mask() I47 { return 0x7fffffffffff }
func (i I47) sign() I47 { return 1 << (i.nbits() - 1) }
func (i I47) cast() I47 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 47 bits to u.
func (i *I47) Set(v I47) I47 { *i = new(I47(v)).cast(); return I47.cast(v) }

// Add returns i+o reduced to 47 bits.
func (i I47) Add(o I47) I47 { return I47.cast(i + o) }

// Sub returns i-o reduced to 47 bits.
func (i I47) Sub(o I47) I47 { return I47.cast(i - o) }

// Inc returns i+1 reduced to 47 bits.
func (i I47) Inc() I47 { return I47.cast(i + 1) }

// Dec returns i-1 reduced to 47 bits.
func (i I47) Dec() I47 { return I47.cast(i - 1) }

// Mul returns i*o reduced to 47 bits.
func (i I47) Mul(o I47) I47 { return I47.cast(i * o) }

// Div returns i divided by o.
func (i I47) Div(o I47) I47 { return I47.cast(i / o) }

// Mod returns i modulo o.
func (i I47) Mod(o I47) I47 { return I47.cast(i % o) }

// And returns i&o reduced to 47 bits.
func (i I47) And(o I47) I47 { return I47.cast(i & o) }

// Or returns i|o reduced to 47 bits.
func (i I47) Or(o I47) I47 { return I47.cast(i | o) }

// Xor returns i^o reduced to 47 bits.
func (i I47) Xor(o I47) I47 { return I47.cast(i ^ o) }

// Shr returns i>>o reduced to 47 bits.
func (i I47) Shr(o I47) I47 { return I47.cast(i >> o) }

// Shl returns i<<o reduced to 47 bits.
func (i I47) Shl(o I47) I47 { return I47.cast(i << o) }

// Not returns bitwise complement of i reduced to 47 bits.
func (i I47) Not() I47 { return I47.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I47) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I47) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U47 bit pattern.
func (i I47) Unsigned() U47 { return U47.cast(U47(i)) }

// Clamp returns value saturated into the representable range of I47.
func (i I47) Clamp(value int64) I47 { return iclamp[I47](value) }

// Clip masks i to a bits-wide low field and sign-extends to 47 bits.
func (i I47) Clip(bits int) I47 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I47(uint64(i)&m^b) - I47(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I47) Bit(index int) I47 {
	if index < 0 {
		index = 47 + index
	}
	return (i >> I47(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I47) SetBit(index int, v I47) {
	if index < 0 {
		index = 47 + index
	}
	bit := I47(1) << I47(index)
	*i = I47.cast((*i &^ bit) | ((v & 1) << I47(index)))
}

// Bitref returns a Range for bit index.
func (i *I47) Bitref(index int) Range[I47] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I47) Bits(lo, hi int) I47 {
	p := 47
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I47((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I47(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I47) SetBits(lo, hi int, v I47) {
	p := 47
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I47((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I47.cast((*i &^ mask) | ((v << I47(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I47) Bitsref(lo, hi int) Range[I47] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I47) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I47) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I47(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I47) Byteref(idx int) Range[I47] { return byte(i, idx) }

// U48 is an 48-bit unsigned integer.
type U48 uint64

// AsU48 returns a U48 representing v reduced to 48 bits.
func AsU48[T Unsigned](v T) U48 { return U48.cast(U48(v)) }

func (U48) nbits() int  { return 48 }
func (u U48) mask() U48 { return 0xffffffffffff }
func (u U48) cast() U48 { return u & u.mask() }

// Set assigns v reduced to 48 bits to u.
func (u *U48) Set(v U48) U48 { *u = new(U48(v)).cast(); return U48.cast(v) }

// Add returns u+o reduced to 48 bits.
func (u U48) Add(o U48) U48 { return U48.cast(u + o) }

// Sub returns u-o reduced to 48 bits.
func (u U48) Sub(o U48) U48 { return U48.cast(u - o) }

// Inc returns u+1 reduced to 48 bits.
func (u U48) Inc() U48 { return U48.cast(u + 1) }

// Dec returns u-1 reduced to 48 bits.
func (u U48) Dec() U48 { return U48.cast(u - 1) }

// Mul returns u*o reduced to 48 bits.
func (u U48) Mul(o U48) U48 { return U48.cast(u * o) }

// Div returns u divided by o.
func (u U48) Div(o U48) U48 { return U48.cast(u / o) }

// Mod returns u modulo o.
func (u U48) Mod(o U48) U48 { return U48.cast(u % o) }

// And returns u&o reduced to 48 bits.
func (u U48) And(o U48) U48 { return U48.cast(u & o) }

// Or returns u|o reduced to 48 bits.
func (u U48) Or(o U48) U48 { return U48.cast(u | o) }

// Xor returns u^o reduced to 48 bits.
func (u U48) Xor(o U48) U48 { return U48.cast(u ^ o) }

// Shr returns u>>o reduced to 48 bits.
func (u U48) Shr(o U48) U48 { return U48.cast(u >> o) }

// Shl returns u<<o reduced to 48 bits.
func (u U48) Shl(o U48) U48 { return U48.cast(u << o) }

// Not returns bitwise complement of u reduced to 48 bits.
func (u U48) Not() U48 { return U48.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U48) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U48) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U48 value.
func (u U48) Signed() I48 { return I48.cast(I48(u)) }

// Clamp returns value saturated into the representable range of U48.
func (u U48) Clamp(value uint64) U48 { return uclamp[U48](value) }

// Clip masks u to a bits-wide low field.
func (u U48) Clip(bits int) U48 {
	b := 1 << (bits - 1)
	m := U48(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U48) Bit(index int) U48 {
	if index < 0 {
		index = 48 + index
	}
	return (u >> U48(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U48) SetBit(index int, v U48) {
	if index < 0 {
		index = 48 + index
	}
	bit := U48(1) << U48(index)
	*u = U48.cast((*u &^ bit) | ((v & 1) << U48(index)))
}

// Bitref returns a Range for bit index.
func (u *U48) Bitref(index int) Range[U48] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U48) Bits(lo, hi int) U48 {
	p := 48
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U48((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U48(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U48) SetBits(lo, hi int, v U48) {
	p := 48
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U48((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U48.cast((*u &^ mask) | ((v << U48(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U48) Bitsref(lo, hi int) Range[U48] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U48) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U48) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U48(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U48) Byteref(idx int) Range[U48] { return byte(u, idx) }

// I48 is an 48-bit signed integer in two's complement.
type I48 int64

// AsI48 returns a I48 representing v reduced to 48 bits.
func AsI48[T Signed](v T) I48 { return I48.cast(I48(v)) }

func (I48) nbits() int  { return 48 }
func (i I48) mask() I48 { return 0xffffffffffff }
func (i I48) sign() I48 { return 1 << (i.nbits() - 1) }
func (i I48) cast() I48 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 48 bits to u.
func (i *I48) Set(v I48) I48 { *i = new(I48(v)).cast(); return I48.cast(v) }

// Add returns i+o reduced to 48 bits.
func (i I48) Add(o I48) I48 { return I48.cast(i + o) }

// Sub returns i-o reduced to 48 bits.
func (i I48) Sub(o I48) I48 { return I48.cast(i - o) }

// Inc returns i+1 reduced to 48 bits.
func (i I48) Inc() I48 { return I48.cast(i + 1) }

// Dec returns i-1 reduced to 48 bits.
func (i I48) Dec() I48 { return I48.cast(i - 1) }

// Mul returns i*o reduced to 48 bits.
func (i I48) Mul(o I48) I48 { return I48.cast(i * o) }

// Div returns i divided by o.
func (i I48) Div(o I48) I48 { return I48.cast(i / o) }

// Mod returns i modulo o.
func (i I48) Mod(o I48) I48 { return I48.cast(i % o) }

// And returns i&o reduced to 48 bits.
func (i I48) And(o I48) I48 { return I48.cast(i & o) }

// Or returns i|o reduced to 48 bits.
func (i I48) Or(o I48) I48 { return I48.cast(i | o) }

// Xor returns i^o reduced to 48 bits.
func (i I48) Xor(o I48) I48 { return I48.cast(i ^ o) }

// Shr returns i>>o reduced to 48 bits.
func (i I48) Shr(o I48) I48 { return I48.cast(i >> o) }

// Shl returns i<<o reduced to 48 bits.
func (i I48) Shl(o I48) I48 { return I48.cast(i << o) }

// Not returns bitwise complement of i reduced to 48 bits.
func (i I48) Not() I48 { return I48.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I48) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I48) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U48 bit pattern.
func (i I48) Unsigned() U48 { return U48.cast(U48(i)) }

// Clamp returns value saturated into the representable range of I48.
func (i I48) Clamp(value int64) I48 { return iclamp[I48](value) }

// Clip masks i to a bits-wide low field and sign-extends to 48 bits.
func (i I48) Clip(bits int) I48 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I48(uint64(i)&m^b) - I48(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I48) Bit(index int) I48 {
	if index < 0 {
		index = 48 + index
	}
	return (i >> I48(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I48) SetBit(index int, v I48) {
	if index < 0 {
		index = 48 + index
	}
	bit := I48(1) << I48(index)
	*i = I48.cast((*i &^ bit) | ((v & 1) << I48(index)))
}

// Bitref returns a Range for bit index.
func (i *I48) Bitref(index int) Range[I48] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I48) Bits(lo, hi int) I48 {
	p := 48
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I48((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I48(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I48) SetBits(lo, hi int, v I48) {
	p := 48
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I48((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I48.cast((*i &^ mask) | ((v << I48(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I48) Bitsref(lo, hi int) Range[I48] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I48) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I48) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I48(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I48) Byteref(idx int) Range[I48] { return byte(i, idx) }

// U49 is an 49-bit unsigned integer.
type U49 uint64

// AsU49 returns a U49 representing v reduced to 49 bits.
func AsU49[T Unsigned](v T) U49 { return U49.cast(U49(v)) }

func (U49) nbits() int  { return 49 }
func (u U49) mask() U49 { return 0x1ffffffffffff }
func (u U49) cast() U49 { return u & u.mask() }

// Set assigns v reduced to 49 bits to u.
func (u *U49) Set(v U49) U49 { *u = new(U49(v)).cast(); return U49.cast(v) }

// Add returns u+o reduced to 49 bits.
func (u U49) Add(o U49) U49 { return U49.cast(u + o) }

// Sub returns u-o reduced to 49 bits.
func (u U49) Sub(o U49) U49 { return U49.cast(u - o) }

// Inc returns u+1 reduced to 49 bits.
func (u U49) Inc() U49 { return U49.cast(u + 1) }

// Dec returns u-1 reduced to 49 bits.
func (u U49) Dec() U49 { return U49.cast(u - 1) }

// Mul returns u*o reduced to 49 bits.
func (u U49) Mul(o U49) U49 { return U49.cast(u * o) }

// Div returns u divided by o.
func (u U49) Div(o U49) U49 { return U49.cast(u / o) }

// Mod returns u modulo o.
func (u U49) Mod(o U49) U49 { return U49.cast(u % o) }

// And returns u&o reduced to 49 bits.
func (u U49) And(o U49) U49 { return U49.cast(u & o) }

// Or returns u|o reduced to 49 bits.
func (u U49) Or(o U49) U49 { return U49.cast(u | o) }

// Xor returns u^o reduced to 49 bits.
func (u U49) Xor(o U49) U49 { return U49.cast(u ^ o) }

// Shr returns u>>o reduced to 49 bits.
func (u U49) Shr(o U49) U49 { return U49.cast(u >> o) }

// Shl returns u<<o reduced to 49 bits.
func (u U49) Shl(o U49) U49 { return U49.cast(u << o) }

// Not returns bitwise complement of u reduced to 49 bits.
func (u U49) Not() U49 { return U49.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U49) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U49) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U49 value.
func (u U49) Signed() I49 { return I49.cast(I49(u)) }

// Clamp returns value saturated into the representable range of U49.
func (u U49) Clamp(value uint64) U49 { return uclamp[U49](value) }

// Clip masks u to a bits-wide low field.
func (u U49) Clip(bits int) U49 {
	b := 1 << (bits - 1)
	m := U49(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U49) Bit(index int) U49 {
	if index < 0 {
		index = 49 + index
	}
	return (u >> U49(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U49) SetBit(index int, v U49) {
	if index < 0 {
		index = 49 + index
	}
	bit := U49(1) << U49(index)
	*u = U49.cast((*u &^ bit) | ((v & 1) << U49(index)))
}

// Bitref returns a Range for bit index.
func (u *U49) Bitref(index int) Range[U49] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U49) Bits(lo, hi int) U49 {
	p := 49
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U49((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U49(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U49) SetBits(lo, hi int, v U49) {
	p := 49
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U49((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U49.cast((*u &^ mask) | ((v << U49(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U49) Bitsref(lo, hi int) Range[U49] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U49) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U49) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U49(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U49) Byteref(idx int) Range[U49] { return byte(u, idx) }

// I49 is an 49-bit signed integer in two's complement.
type I49 int64

// AsI49 returns a I49 representing v reduced to 49 bits.
func AsI49[T Signed](v T) I49 { return I49.cast(I49(v)) }

func (I49) nbits() int  { return 49 }
func (i I49) mask() I49 { return 0x1ffffffffffff }
func (i I49) sign() I49 { return 1 << (i.nbits() - 1) }
func (i I49) cast() I49 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 49 bits to u.
func (i *I49) Set(v I49) I49 { *i = new(I49(v)).cast(); return I49.cast(v) }

// Add returns i+o reduced to 49 bits.
func (i I49) Add(o I49) I49 { return I49.cast(i + o) }

// Sub returns i-o reduced to 49 bits.
func (i I49) Sub(o I49) I49 { return I49.cast(i - o) }

// Inc returns i+1 reduced to 49 bits.
func (i I49) Inc() I49 { return I49.cast(i + 1) }

// Dec returns i-1 reduced to 49 bits.
func (i I49) Dec() I49 { return I49.cast(i - 1) }

// Mul returns i*o reduced to 49 bits.
func (i I49) Mul(o I49) I49 { return I49.cast(i * o) }

// Div returns i divided by o.
func (i I49) Div(o I49) I49 { return I49.cast(i / o) }

// Mod returns i modulo o.
func (i I49) Mod(o I49) I49 { return I49.cast(i % o) }

// And returns i&o reduced to 49 bits.
func (i I49) And(o I49) I49 { return I49.cast(i & o) }

// Or returns i|o reduced to 49 bits.
func (i I49) Or(o I49) I49 { return I49.cast(i | o) }

// Xor returns i^o reduced to 49 bits.
func (i I49) Xor(o I49) I49 { return I49.cast(i ^ o) }

// Shr returns i>>o reduced to 49 bits.
func (i I49) Shr(o I49) I49 { return I49.cast(i >> o) }

// Shl returns i<<o reduced to 49 bits.
func (i I49) Shl(o I49) I49 { return I49.cast(i << o) }

// Not returns bitwise complement of i reduced to 49 bits.
func (i I49) Not() I49 { return I49.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I49) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I49) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U49 bit pattern.
func (i I49) Unsigned() U49 { return U49.cast(U49(i)) }

// Clamp returns value saturated into the representable range of I49.
func (i I49) Clamp(value int64) I49 { return iclamp[I49](value) }

// Clip masks i to a bits-wide low field and sign-extends to 49 bits.
func (i I49) Clip(bits int) I49 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I49(uint64(i)&m^b) - I49(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I49) Bit(index int) I49 {
	if index < 0 {
		index = 49 + index
	}
	return (i >> I49(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I49) SetBit(index int, v I49) {
	if index < 0 {
		index = 49 + index
	}
	bit := I49(1) << I49(index)
	*i = I49.cast((*i &^ bit) | ((v & 1) << I49(index)))
}

// Bitref returns a Range for bit index.
func (i *I49) Bitref(index int) Range[I49] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I49) Bits(lo, hi int) I49 {
	p := 49
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I49((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I49(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I49) SetBits(lo, hi int, v I49) {
	p := 49
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I49((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I49.cast((*i &^ mask) | ((v << I49(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I49) Bitsref(lo, hi int) Range[I49] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I49) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I49) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I49(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I49) Byteref(idx int) Range[I49] { return byte(i, idx) }

// U50 is an 50-bit unsigned integer.
type U50 uint64

// AsU50 returns a U50 representing v reduced to 50 bits.
func AsU50[T Unsigned](v T) U50 { return U50.cast(U50(v)) }

func (U50) nbits() int  { return 50 }
func (u U50) mask() U50 { return 0x3ffffffffffff }
func (u U50) cast() U50 { return u & u.mask() }

// Set assigns v reduced to 50 bits to u.
func (u *U50) Set(v U50) U50 { *u = new(U50(v)).cast(); return U50.cast(v) }

// Add returns u+o reduced to 50 bits.
func (u U50) Add(o U50) U50 { return U50.cast(u + o) }

// Sub returns u-o reduced to 50 bits.
func (u U50) Sub(o U50) U50 { return U50.cast(u - o) }

// Inc returns u+1 reduced to 50 bits.
func (u U50) Inc() U50 { return U50.cast(u + 1) }

// Dec returns u-1 reduced to 50 bits.
func (u U50) Dec() U50 { return U50.cast(u - 1) }

// Mul returns u*o reduced to 50 bits.
func (u U50) Mul(o U50) U50 { return U50.cast(u * o) }

// Div returns u divided by o.
func (u U50) Div(o U50) U50 { return U50.cast(u / o) }

// Mod returns u modulo o.
func (u U50) Mod(o U50) U50 { return U50.cast(u % o) }

// And returns u&o reduced to 50 bits.
func (u U50) And(o U50) U50 { return U50.cast(u & o) }

// Or returns u|o reduced to 50 bits.
func (u U50) Or(o U50) U50 { return U50.cast(u | o) }

// Xor returns u^o reduced to 50 bits.
func (u U50) Xor(o U50) U50 { return U50.cast(u ^ o) }

// Shr returns u>>o reduced to 50 bits.
func (u U50) Shr(o U50) U50 { return U50.cast(u >> o) }

// Shl returns u<<o reduced to 50 bits.
func (u U50) Shl(o U50) U50 { return U50.cast(u << o) }

// Not returns bitwise complement of u reduced to 50 bits.
func (u U50) Not() U50 { return U50.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U50) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U50) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U50 value.
func (u U50) Signed() I50 { return I50.cast(I50(u)) }

// Clamp returns value saturated into the representable range of U50.
func (u U50) Clamp(value uint64) U50 { return uclamp[U50](value) }

// Clip masks u to a bits-wide low field.
func (u U50) Clip(bits int) U50 {
	b := 1 << (bits - 1)
	m := U50(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U50) Bit(index int) U50 {
	if index < 0 {
		index = 50 + index
	}
	return (u >> U50(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U50) SetBit(index int, v U50) {
	if index < 0 {
		index = 50 + index
	}
	bit := U50(1) << U50(index)
	*u = U50.cast((*u &^ bit) | ((v & 1) << U50(index)))
}

// Bitref returns a Range for bit index.
func (u *U50) Bitref(index int) Range[U50] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U50) Bits(lo, hi int) U50 {
	p := 50
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U50((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U50(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U50) SetBits(lo, hi int, v U50) {
	p := 50
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U50((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U50.cast((*u &^ mask) | ((v << U50(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U50) Bitsref(lo, hi int) Range[U50] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U50) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U50) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U50(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U50) Byteref(idx int) Range[U50] { return byte(u, idx) }

// I50 is an 50-bit signed integer in two's complement.
type I50 int64

// AsI50 returns a I50 representing v reduced to 50 bits.
func AsI50[T Signed](v T) I50 { return I50.cast(I50(v)) }

func (I50) nbits() int  { return 50 }
func (i I50) mask() I50 { return 0x3ffffffffffff }
func (i I50) sign() I50 { return 1 << (i.nbits() - 1) }
func (i I50) cast() I50 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 50 bits to u.
func (i *I50) Set(v I50) I50 { *i = new(I50(v)).cast(); return I50.cast(v) }

// Add returns i+o reduced to 50 bits.
func (i I50) Add(o I50) I50 { return I50.cast(i + o) }

// Sub returns i-o reduced to 50 bits.
func (i I50) Sub(o I50) I50 { return I50.cast(i - o) }

// Inc returns i+1 reduced to 50 bits.
func (i I50) Inc() I50 { return I50.cast(i + 1) }

// Dec returns i-1 reduced to 50 bits.
func (i I50) Dec() I50 { return I50.cast(i - 1) }

// Mul returns i*o reduced to 50 bits.
func (i I50) Mul(o I50) I50 { return I50.cast(i * o) }

// Div returns i divided by o.
func (i I50) Div(o I50) I50 { return I50.cast(i / o) }

// Mod returns i modulo o.
func (i I50) Mod(o I50) I50 { return I50.cast(i % o) }

// And returns i&o reduced to 50 bits.
func (i I50) And(o I50) I50 { return I50.cast(i & o) }

// Or returns i|o reduced to 50 bits.
func (i I50) Or(o I50) I50 { return I50.cast(i | o) }

// Xor returns i^o reduced to 50 bits.
func (i I50) Xor(o I50) I50 { return I50.cast(i ^ o) }

// Shr returns i>>o reduced to 50 bits.
func (i I50) Shr(o I50) I50 { return I50.cast(i >> o) }

// Shl returns i<<o reduced to 50 bits.
func (i I50) Shl(o I50) I50 { return I50.cast(i << o) }

// Not returns bitwise complement of i reduced to 50 bits.
func (i I50) Not() I50 { return I50.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I50) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I50) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U50 bit pattern.
func (i I50) Unsigned() U50 { return U50.cast(U50(i)) }

// Clamp returns value saturated into the representable range of I50.
func (i I50) Clamp(value int64) I50 { return iclamp[I50](value) }

// Clip masks i to a bits-wide low field and sign-extends to 50 bits.
func (i I50) Clip(bits int) I50 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I50(uint64(i)&m^b) - I50(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I50) Bit(index int) I50 {
	if index < 0 {
		index = 50 + index
	}
	return (i >> I50(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I50) SetBit(index int, v I50) {
	if index < 0 {
		index = 50 + index
	}
	bit := I50(1) << I50(index)
	*i = I50.cast((*i &^ bit) | ((v & 1) << I50(index)))
}

// Bitref returns a Range for bit index.
func (i *I50) Bitref(index int) Range[I50] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I50) Bits(lo, hi int) I50 {
	p := 50
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I50((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I50(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I50) SetBits(lo, hi int, v I50) {
	p := 50
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I50((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I50.cast((*i &^ mask) | ((v << I50(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I50) Bitsref(lo, hi int) Range[I50] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I50) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I50) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I50(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I50) Byteref(idx int) Range[I50] { return byte(i, idx) }

// U51 is an 51-bit unsigned integer.
type U51 uint64

// AsU51 returns a U51 representing v reduced to 51 bits.
func AsU51[T Unsigned](v T) U51 { return U51.cast(U51(v)) }

func (U51) nbits() int  { return 51 }
func (u U51) mask() U51 { return 0x7ffffffffffff }
func (u U51) cast() U51 { return u & u.mask() }

// Set assigns v reduced to 51 bits to u.
func (u *U51) Set(v U51) U51 { *u = new(U51(v)).cast(); return U51.cast(v) }

// Add returns u+o reduced to 51 bits.
func (u U51) Add(o U51) U51 { return U51.cast(u + o) }

// Sub returns u-o reduced to 51 bits.
func (u U51) Sub(o U51) U51 { return U51.cast(u - o) }

// Inc returns u+1 reduced to 51 bits.
func (u U51) Inc() U51 { return U51.cast(u + 1) }

// Dec returns u-1 reduced to 51 bits.
func (u U51) Dec() U51 { return U51.cast(u - 1) }

// Mul returns u*o reduced to 51 bits.
func (u U51) Mul(o U51) U51 { return U51.cast(u * o) }

// Div returns u divided by o.
func (u U51) Div(o U51) U51 { return U51.cast(u / o) }

// Mod returns u modulo o.
func (u U51) Mod(o U51) U51 { return U51.cast(u % o) }

// And returns u&o reduced to 51 bits.
func (u U51) And(o U51) U51 { return U51.cast(u & o) }

// Or returns u|o reduced to 51 bits.
func (u U51) Or(o U51) U51 { return U51.cast(u | o) }

// Xor returns u^o reduced to 51 bits.
func (u U51) Xor(o U51) U51 { return U51.cast(u ^ o) }

// Shr returns u>>o reduced to 51 bits.
func (u U51) Shr(o U51) U51 { return U51.cast(u >> o) }

// Shl returns u<<o reduced to 51 bits.
func (u U51) Shl(o U51) U51 { return U51.cast(u << o) }

// Not returns bitwise complement of u reduced to 51 bits.
func (u U51) Not() U51 { return U51.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U51) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U51) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U51 value.
func (u U51) Signed() I51 { return I51.cast(I51(u)) }

// Clamp returns value saturated into the representable range of U51.
func (u U51) Clamp(value uint64) U51 { return uclamp[U51](value) }

// Clip masks u to a bits-wide low field.
func (u U51) Clip(bits int) U51 {
	b := 1 << (bits - 1)
	m := U51(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U51) Bit(index int) U51 {
	if index < 0 {
		index = 51 + index
	}
	return (u >> U51(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U51) SetBit(index int, v U51) {
	if index < 0 {
		index = 51 + index
	}
	bit := U51(1) << U51(index)
	*u = U51.cast((*u &^ bit) | ((v & 1) << U51(index)))
}

// Bitref returns a Range for bit index.
func (u *U51) Bitref(index int) Range[U51] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U51) Bits(lo, hi int) U51 {
	p := 51
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U51((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U51(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U51) SetBits(lo, hi int, v U51) {
	p := 51
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U51((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U51.cast((*u &^ mask) | ((v << U51(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U51) Bitsref(lo, hi int) Range[U51] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U51) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U51) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U51(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U51) Byteref(idx int) Range[U51] { return byte(u, idx) }

// I51 is an 51-bit signed integer in two's complement.
type I51 int64

// AsI51 returns a I51 representing v reduced to 51 bits.
func AsI51[T Signed](v T) I51 { return I51.cast(I51(v)) }

func (I51) nbits() int  { return 51 }
func (i I51) mask() I51 { return 0x7ffffffffffff }
func (i I51) sign() I51 { return 1 << (i.nbits() - 1) }
func (i I51) cast() I51 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 51 bits to u.
func (i *I51) Set(v I51) I51 { *i = new(I51(v)).cast(); return I51.cast(v) }

// Add returns i+o reduced to 51 bits.
func (i I51) Add(o I51) I51 { return I51.cast(i + o) }

// Sub returns i-o reduced to 51 bits.
func (i I51) Sub(o I51) I51 { return I51.cast(i - o) }

// Inc returns i+1 reduced to 51 bits.
func (i I51) Inc() I51 { return I51.cast(i + 1) }

// Dec returns i-1 reduced to 51 bits.
func (i I51) Dec() I51 { return I51.cast(i - 1) }

// Mul returns i*o reduced to 51 bits.
func (i I51) Mul(o I51) I51 { return I51.cast(i * o) }

// Div returns i divided by o.
func (i I51) Div(o I51) I51 { return I51.cast(i / o) }

// Mod returns i modulo o.
func (i I51) Mod(o I51) I51 { return I51.cast(i % o) }

// And returns i&o reduced to 51 bits.
func (i I51) And(o I51) I51 { return I51.cast(i & o) }

// Or returns i|o reduced to 51 bits.
func (i I51) Or(o I51) I51 { return I51.cast(i | o) }

// Xor returns i^o reduced to 51 bits.
func (i I51) Xor(o I51) I51 { return I51.cast(i ^ o) }

// Shr returns i>>o reduced to 51 bits.
func (i I51) Shr(o I51) I51 { return I51.cast(i >> o) }

// Shl returns i<<o reduced to 51 bits.
func (i I51) Shl(o I51) I51 { return I51.cast(i << o) }

// Not returns bitwise complement of i reduced to 51 bits.
func (i I51) Not() I51 { return I51.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I51) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I51) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U51 bit pattern.
func (i I51) Unsigned() U51 { return U51.cast(U51(i)) }

// Clamp returns value saturated into the representable range of I51.
func (i I51) Clamp(value int64) I51 { return iclamp[I51](value) }

// Clip masks i to a bits-wide low field and sign-extends to 51 bits.
func (i I51) Clip(bits int) I51 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I51(uint64(i)&m^b) - I51(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I51) Bit(index int) I51 {
	if index < 0 {
		index = 51 + index
	}
	return (i >> I51(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I51) SetBit(index int, v I51) {
	if index < 0 {
		index = 51 + index
	}
	bit := I51(1) << I51(index)
	*i = I51.cast((*i &^ bit) | ((v & 1) << I51(index)))
}

// Bitref returns a Range for bit index.
func (i *I51) Bitref(index int) Range[I51] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I51) Bits(lo, hi int) I51 {
	p := 51
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I51((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I51(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I51) SetBits(lo, hi int, v I51) {
	p := 51
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I51((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I51.cast((*i &^ mask) | ((v << I51(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I51) Bitsref(lo, hi int) Range[I51] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I51) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I51) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I51(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I51) Byteref(idx int) Range[I51] { return byte(i, idx) }

// U52 is an 52-bit unsigned integer.
type U52 uint64

// AsU52 returns a U52 representing v reduced to 52 bits.
func AsU52[T Unsigned](v T) U52 { return U52.cast(U52(v)) }

func (U52) nbits() int  { return 52 }
func (u U52) mask() U52 { return 0xfffffffffffff }
func (u U52) cast() U52 { return u & u.mask() }

// Set assigns v reduced to 52 bits to u.
func (u *U52) Set(v U52) U52 { *u = new(U52(v)).cast(); return U52.cast(v) }

// Add returns u+o reduced to 52 bits.
func (u U52) Add(o U52) U52 { return U52.cast(u + o) }

// Sub returns u-o reduced to 52 bits.
func (u U52) Sub(o U52) U52 { return U52.cast(u - o) }

// Inc returns u+1 reduced to 52 bits.
func (u U52) Inc() U52 { return U52.cast(u + 1) }

// Dec returns u-1 reduced to 52 bits.
func (u U52) Dec() U52 { return U52.cast(u - 1) }

// Mul returns u*o reduced to 52 bits.
func (u U52) Mul(o U52) U52 { return U52.cast(u * o) }

// Div returns u divided by o.
func (u U52) Div(o U52) U52 { return U52.cast(u / o) }

// Mod returns u modulo o.
func (u U52) Mod(o U52) U52 { return U52.cast(u % o) }

// And returns u&o reduced to 52 bits.
func (u U52) And(o U52) U52 { return U52.cast(u & o) }

// Or returns u|o reduced to 52 bits.
func (u U52) Or(o U52) U52 { return U52.cast(u | o) }

// Xor returns u^o reduced to 52 bits.
func (u U52) Xor(o U52) U52 { return U52.cast(u ^ o) }

// Shr returns u>>o reduced to 52 bits.
func (u U52) Shr(o U52) U52 { return U52.cast(u >> o) }

// Shl returns u<<o reduced to 52 bits.
func (u U52) Shl(o U52) U52 { return U52.cast(u << o) }

// Not returns bitwise complement of u reduced to 52 bits.
func (u U52) Not() U52 { return U52.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U52) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U52) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U52 value.
func (u U52) Signed() I52 { return I52.cast(I52(u)) }

// Clamp returns value saturated into the representable range of U52.
func (u U52) Clamp(value uint64) U52 { return uclamp[U52](value) }

// Clip masks u to a bits-wide low field.
func (u U52) Clip(bits int) U52 {
	b := 1 << (bits - 1)
	m := U52(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U52) Bit(index int) U52 {
	if index < 0 {
		index = 52 + index
	}
	return (u >> U52(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U52) SetBit(index int, v U52) {
	if index < 0 {
		index = 52 + index
	}
	bit := U52(1) << U52(index)
	*u = U52.cast((*u &^ bit) | ((v & 1) << U52(index)))
}

// Bitref returns a Range for bit index.
func (u *U52) Bitref(index int) Range[U52] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U52) Bits(lo, hi int) U52 {
	p := 52
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U52((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U52(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U52) SetBits(lo, hi int, v U52) {
	p := 52
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U52((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U52.cast((*u &^ mask) | ((v << U52(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U52) Bitsref(lo, hi int) Range[U52] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U52) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U52) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U52(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U52) Byteref(idx int) Range[U52] { return byte(u, idx) }

// I52 is an 52-bit signed integer in two's complement.
type I52 int64

// AsI52 returns a I52 representing v reduced to 52 bits.
func AsI52[T Signed](v T) I52 { return I52.cast(I52(v)) }

func (I52) nbits() int  { return 52 }
func (i I52) mask() I52 { return 0xfffffffffffff }
func (i I52) sign() I52 { return 1 << (i.nbits() - 1) }
func (i I52) cast() I52 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 52 bits to u.
func (i *I52) Set(v I52) I52 { *i = new(I52(v)).cast(); return I52.cast(v) }

// Add returns i+o reduced to 52 bits.
func (i I52) Add(o I52) I52 { return I52.cast(i + o) }

// Sub returns i-o reduced to 52 bits.
func (i I52) Sub(o I52) I52 { return I52.cast(i - o) }

// Inc returns i+1 reduced to 52 bits.
func (i I52) Inc() I52 { return I52.cast(i + 1) }

// Dec returns i-1 reduced to 52 bits.
func (i I52) Dec() I52 { return I52.cast(i - 1) }

// Mul returns i*o reduced to 52 bits.
func (i I52) Mul(o I52) I52 { return I52.cast(i * o) }

// Div returns i divided by o.
func (i I52) Div(o I52) I52 { return I52.cast(i / o) }

// Mod returns i modulo o.
func (i I52) Mod(o I52) I52 { return I52.cast(i % o) }

// And returns i&o reduced to 52 bits.
func (i I52) And(o I52) I52 { return I52.cast(i & o) }

// Or returns i|o reduced to 52 bits.
func (i I52) Or(o I52) I52 { return I52.cast(i | o) }

// Xor returns i^o reduced to 52 bits.
func (i I52) Xor(o I52) I52 { return I52.cast(i ^ o) }

// Shr returns i>>o reduced to 52 bits.
func (i I52) Shr(o I52) I52 { return I52.cast(i >> o) }

// Shl returns i<<o reduced to 52 bits.
func (i I52) Shl(o I52) I52 { return I52.cast(i << o) }

// Not returns bitwise complement of i reduced to 52 bits.
func (i I52) Not() I52 { return I52.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I52) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I52) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U52 bit pattern.
func (i I52) Unsigned() U52 { return U52.cast(U52(i)) }

// Clamp returns value saturated into the representable range of I52.
func (i I52) Clamp(value int64) I52 { return iclamp[I52](value) }

// Clip masks i to a bits-wide low field and sign-extends to 52 bits.
func (i I52) Clip(bits int) I52 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I52(uint64(i)&m^b) - I52(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I52) Bit(index int) I52 {
	if index < 0 {
		index = 52 + index
	}
	return (i >> I52(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I52) SetBit(index int, v I52) {
	if index < 0 {
		index = 52 + index
	}
	bit := I52(1) << I52(index)
	*i = I52.cast((*i &^ bit) | ((v & 1) << I52(index)))
}

// Bitref returns a Range for bit index.
func (i *I52) Bitref(index int) Range[I52] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I52) Bits(lo, hi int) I52 {
	p := 52
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I52((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I52(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I52) SetBits(lo, hi int, v I52) {
	p := 52
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I52((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I52.cast((*i &^ mask) | ((v << I52(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I52) Bitsref(lo, hi int) Range[I52] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I52) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I52) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I52(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I52) Byteref(idx int) Range[I52] { return byte(i, idx) }

// U53 is an 53-bit unsigned integer.
type U53 uint64

// AsU53 returns a U53 representing v reduced to 53 bits.
func AsU53[T Unsigned](v T) U53 { return U53.cast(U53(v)) }

func (U53) nbits() int  { return 53 }
func (u U53) mask() U53 { return 0x1fffffffffffff }
func (u U53) cast() U53 { return u & u.mask() }

// Set assigns v reduced to 53 bits to u.
func (u *U53) Set(v U53) U53 { *u = new(U53(v)).cast(); return U53.cast(v) }

// Add returns u+o reduced to 53 bits.
func (u U53) Add(o U53) U53 { return U53.cast(u + o) }

// Sub returns u-o reduced to 53 bits.
func (u U53) Sub(o U53) U53 { return U53.cast(u - o) }

// Inc returns u+1 reduced to 53 bits.
func (u U53) Inc() U53 { return U53.cast(u + 1) }

// Dec returns u-1 reduced to 53 bits.
func (u U53) Dec() U53 { return U53.cast(u - 1) }

// Mul returns u*o reduced to 53 bits.
func (u U53) Mul(o U53) U53 { return U53.cast(u * o) }

// Div returns u divided by o.
func (u U53) Div(o U53) U53 { return U53.cast(u / o) }

// Mod returns u modulo o.
func (u U53) Mod(o U53) U53 { return U53.cast(u % o) }

// And returns u&o reduced to 53 bits.
func (u U53) And(o U53) U53 { return U53.cast(u & o) }

// Or returns u|o reduced to 53 bits.
func (u U53) Or(o U53) U53 { return U53.cast(u | o) }

// Xor returns u^o reduced to 53 bits.
func (u U53) Xor(o U53) U53 { return U53.cast(u ^ o) }

// Shr returns u>>o reduced to 53 bits.
func (u U53) Shr(o U53) U53 { return U53.cast(u >> o) }

// Shl returns u<<o reduced to 53 bits.
func (u U53) Shl(o U53) U53 { return U53.cast(u << o) }

// Not returns bitwise complement of u reduced to 53 bits.
func (u U53) Not() U53 { return U53.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U53) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U53) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U53 value.
func (u U53) Signed() I53 { return I53.cast(I53(u)) }

// Clamp returns value saturated into the representable range of U53.
func (u U53) Clamp(value uint64) U53 { return uclamp[U53](value) }

// Clip masks u to a bits-wide low field.
func (u U53) Clip(bits int) U53 {
	b := 1 << (bits - 1)
	m := U53(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U53) Bit(index int) U53 {
	if index < 0 {
		index = 53 + index
	}
	return (u >> U53(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U53) SetBit(index int, v U53) {
	if index < 0 {
		index = 53 + index
	}
	bit := U53(1) << U53(index)
	*u = U53.cast((*u &^ bit) | ((v & 1) << U53(index)))
}

// Bitref returns a Range for bit index.
func (u *U53) Bitref(index int) Range[U53] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U53) Bits(lo, hi int) U53 {
	p := 53
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U53((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U53(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U53) SetBits(lo, hi int, v U53) {
	p := 53
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U53((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U53.cast((*u &^ mask) | ((v << U53(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U53) Bitsref(lo, hi int) Range[U53] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U53) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U53) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U53(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U53) Byteref(idx int) Range[U53] { return byte(u, idx) }

// I53 is an 53-bit signed integer in two's complement.
type I53 int64

// AsI53 returns a I53 representing v reduced to 53 bits.
func AsI53[T Signed](v T) I53 { return I53.cast(I53(v)) }

func (I53) nbits() int  { return 53 }
func (i I53) mask() I53 { return 0x1fffffffffffff }
func (i I53) sign() I53 { return 1 << (i.nbits() - 1) }
func (i I53) cast() I53 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 53 bits to u.
func (i *I53) Set(v I53) I53 { *i = new(I53(v)).cast(); return I53.cast(v) }

// Add returns i+o reduced to 53 bits.
func (i I53) Add(o I53) I53 { return I53.cast(i + o) }

// Sub returns i-o reduced to 53 bits.
func (i I53) Sub(o I53) I53 { return I53.cast(i - o) }

// Inc returns i+1 reduced to 53 bits.
func (i I53) Inc() I53 { return I53.cast(i + 1) }

// Dec returns i-1 reduced to 53 bits.
func (i I53) Dec() I53 { return I53.cast(i - 1) }

// Mul returns i*o reduced to 53 bits.
func (i I53) Mul(o I53) I53 { return I53.cast(i * o) }

// Div returns i divided by o.
func (i I53) Div(o I53) I53 { return I53.cast(i / o) }

// Mod returns i modulo o.
func (i I53) Mod(o I53) I53 { return I53.cast(i % o) }

// And returns i&o reduced to 53 bits.
func (i I53) And(o I53) I53 { return I53.cast(i & o) }

// Or returns i|o reduced to 53 bits.
func (i I53) Or(o I53) I53 { return I53.cast(i | o) }

// Xor returns i^o reduced to 53 bits.
func (i I53) Xor(o I53) I53 { return I53.cast(i ^ o) }

// Shr returns i>>o reduced to 53 bits.
func (i I53) Shr(o I53) I53 { return I53.cast(i >> o) }

// Shl returns i<<o reduced to 53 bits.
func (i I53) Shl(o I53) I53 { return I53.cast(i << o) }

// Not returns bitwise complement of i reduced to 53 bits.
func (i I53) Not() I53 { return I53.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I53) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I53) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U53 bit pattern.
func (i I53) Unsigned() U53 { return U53.cast(U53(i)) }

// Clamp returns value saturated into the representable range of I53.
func (i I53) Clamp(value int64) I53 { return iclamp[I53](value) }

// Clip masks i to a bits-wide low field and sign-extends to 53 bits.
func (i I53) Clip(bits int) I53 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I53(uint64(i)&m^b) - I53(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I53) Bit(index int) I53 {
	if index < 0 {
		index = 53 + index
	}
	return (i >> I53(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I53) SetBit(index int, v I53) {
	if index < 0 {
		index = 53 + index
	}
	bit := I53(1) << I53(index)
	*i = I53.cast((*i &^ bit) | ((v & 1) << I53(index)))
}

// Bitref returns a Range for bit index.
func (i *I53) Bitref(index int) Range[I53] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I53) Bits(lo, hi int) I53 {
	p := 53
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I53((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I53(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I53) SetBits(lo, hi int, v I53) {
	p := 53
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I53((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I53.cast((*i &^ mask) | ((v << I53(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I53) Bitsref(lo, hi int) Range[I53] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I53) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I53) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I53(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I53) Byteref(idx int) Range[I53] { return byte(i, idx) }

// U54 is an 54-bit unsigned integer.
type U54 uint64

// AsU54 returns a U54 representing v reduced to 54 bits.
func AsU54[T Unsigned](v T) U54 { return U54.cast(U54(v)) }

func (U54) nbits() int  { return 54 }
func (u U54) mask() U54 { return 0x3fffffffffffff }
func (u U54) cast() U54 { return u & u.mask() }

// Set assigns v reduced to 54 bits to u.
func (u *U54) Set(v U54) U54 { *u = new(U54(v)).cast(); return U54.cast(v) }

// Add returns u+o reduced to 54 bits.
func (u U54) Add(o U54) U54 { return U54.cast(u + o) }

// Sub returns u-o reduced to 54 bits.
func (u U54) Sub(o U54) U54 { return U54.cast(u - o) }

// Inc returns u+1 reduced to 54 bits.
func (u U54) Inc() U54 { return U54.cast(u + 1) }

// Dec returns u-1 reduced to 54 bits.
func (u U54) Dec() U54 { return U54.cast(u - 1) }

// Mul returns u*o reduced to 54 bits.
func (u U54) Mul(o U54) U54 { return U54.cast(u * o) }

// Div returns u divided by o.
func (u U54) Div(o U54) U54 { return U54.cast(u / o) }

// Mod returns u modulo o.
func (u U54) Mod(o U54) U54 { return U54.cast(u % o) }

// And returns u&o reduced to 54 bits.
func (u U54) And(o U54) U54 { return U54.cast(u & o) }

// Or returns u|o reduced to 54 bits.
func (u U54) Or(o U54) U54 { return U54.cast(u | o) }

// Xor returns u^o reduced to 54 bits.
func (u U54) Xor(o U54) U54 { return U54.cast(u ^ o) }

// Shr returns u>>o reduced to 54 bits.
func (u U54) Shr(o U54) U54 { return U54.cast(u >> o) }

// Shl returns u<<o reduced to 54 bits.
func (u U54) Shl(o U54) U54 { return U54.cast(u << o) }

// Not returns bitwise complement of u reduced to 54 bits.
func (u U54) Not() U54 { return U54.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U54) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U54) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U54 value.
func (u U54) Signed() I54 { return I54.cast(I54(u)) }

// Clamp returns value saturated into the representable range of U54.
func (u U54) Clamp(value uint64) U54 { return uclamp[U54](value) }

// Clip masks u to a bits-wide low field.
func (u U54) Clip(bits int) U54 {
	b := 1 << (bits - 1)
	m := U54(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U54) Bit(index int) U54 {
	if index < 0 {
		index = 54 + index
	}
	return (u >> U54(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U54) SetBit(index int, v U54) {
	if index < 0 {
		index = 54 + index
	}
	bit := U54(1) << U54(index)
	*u = U54.cast((*u &^ bit) | ((v & 1) << U54(index)))
}

// Bitref returns a Range for bit index.
func (u *U54) Bitref(index int) Range[U54] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U54) Bits(lo, hi int) U54 {
	p := 54
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U54((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U54(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U54) SetBits(lo, hi int, v U54) {
	p := 54
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U54((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U54.cast((*u &^ mask) | ((v << U54(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U54) Bitsref(lo, hi int) Range[U54] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U54) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U54) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U54(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U54) Byteref(idx int) Range[U54] { return byte(u, idx) }

// I54 is an 54-bit signed integer in two's complement.
type I54 int64

// AsI54 returns a I54 representing v reduced to 54 bits.
func AsI54[T Signed](v T) I54 { return I54.cast(I54(v)) }

func (I54) nbits() int  { return 54 }
func (i I54) mask() I54 { return 0x3fffffffffffff }
func (i I54) sign() I54 { return 1 << (i.nbits() - 1) }
func (i I54) cast() I54 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 54 bits to u.
func (i *I54) Set(v I54) I54 { *i = new(I54(v)).cast(); return I54.cast(v) }

// Add returns i+o reduced to 54 bits.
func (i I54) Add(o I54) I54 { return I54.cast(i + o) }

// Sub returns i-o reduced to 54 bits.
func (i I54) Sub(o I54) I54 { return I54.cast(i - o) }

// Inc returns i+1 reduced to 54 bits.
func (i I54) Inc() I54 { return I54.cast(i + 1) }

// Dec returns i-1 reduced to 54 bits.
func (i I54) Dec() I54 { return I54.cast(i - 1) }

// Mul returns i*o reduced to 54 bits.
func (i I54) Mul(o I54) I54 { return I54.cast(i * o) }

// Div returns i divided by o.
func (i I54) Div(o I54) I54 { return I54.cast(i / o) }

// Mod returns i modulo o.
func (i I54) Mod(o I54) I54 { return I54.cast(i % o) }

// And returns i&o reduced to 54 bits.
func (i I54) And(o I54) I54 { return I54.cast(i & o) }

// Or returns i|o reduced to 54 bits.
func (i I54) Or(o I54) I54 { return I54.cast(i | o) }

// Xor returns i^o reduced to 54 bits.
func (i I54) Xor(o I54) I54 { return I54.cast(i ^ o) }

// Shr returns i>>o reduced to 54 bits.
func (i I54) Shr(o I54) I54 { return I54.cast(i >> o) }

// Shl returns i<<o reduced to 54 bits.
func (i I54) Shl(o I54) I54 { return I54.cast(i << o) }

// Not returns bitwise complement of i reduced to 54 bits.
func (i I54) Not() I54 { return I54.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I54) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I54) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U54 bit pattern.
func (i I54) Unsigned() U54 { return U54.cast(U54(i)) }

// Clamp returns value saturated into the representable range of I54.
func (i I54) Clamp(value int64) I54 { return iclamp[I54](value) }

// Clip masks i to a bits-wide low field and sign-extends to 54 bits.
func (i I54) Clip(bits int) I54 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I54(uint64(i)&m^b) - I54(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I54) Bit(index int) I54 {
	if index < 0 {
		index = 54 + index
	}
	return (i >> I54(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I54) SetBit(index int, v I54) {
	if index < 0 {
		index = 54 + index
	}
	bit := I54(1) << I54(index)
	*i = I54.cast((*i &^ bit) | ((v & 1) << I54(index)))
}

// Bitref returns a Range for bit index.
func (i *I54) Bitref(index int) Range[I54] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I54) Bits(lo, hi int) I54 {
	p := 54
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I54((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I54(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I54) SetBits(lo, hi int, v I54) {
	p := 54
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I54((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I54.cast((*i &^ mask) | ((v << I54(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I54) Bitsref(lo, hi int) Range[I54] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I54) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I54) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I54(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I54) Byteref(idx int) Range[I54] { return byte(i, idx) }

// U55 is an 55-bit unsigned integer.
type U55 uint64

// AsU55 returns a U55 representing v reduced to 55 bits.
func AsU55[T Unsigned](v T) U55 { return U55.cast(U55(v)) }

func (U55) nbits() int  { return 55 }
func (u U55) mask() U55 { return 0x7fffffffffffff }
func (u U55) cast() U55 { return u & u.mask() }

// Set assigns v reduced to 55 bits to u.
func (u *U55) Set(v U55) U55 { *u = new(U55(v)).cast(); return U55.cast(v) }

// Add returns u+o reduced to 55 bits.
func (u U55) Add(o U55) U55 { return U55.cast(u + o) }

// Sub returns u-o reduced to 55 bits.
func (u U55) Sub(o U55) U55 { return U55.cast(u - o) }

// Inc returns u+1 reduced to 55 bits.
func (u U55) Inc() U55 { return U55.cast(u + 1) }

// Dec returns u-1 reduced to 55 bits.
func (u U55) Dec() U55 { return U55.cast(u - 1) }

// Mul returns u*o reduced to 55 bits.
func (u U55) Mul(o U55) U55 { return U55.cast(u * o) }

// Div returns u divided by o.
func (u U55) Div(o U55) U55 { return U55.cast(u / o) }

// Mod returns u modulo o.
func (u U55) Mod(o U55) U55 { return U55.cast(u % o) }

// And returns u&o reduced to 55 bits.
func (u U55) And(o U55) U55 { return U55.cast(u & o) }

// Or returns u|o reduced to 55 bits.
func (u U55) Or(o U55) U55 { return U55.cast(u | o) }

// Xor returns u^o reduced to 55 bits.
func (u U55) Xor(o U55) U55 { return U55.cast(u ^ o) }

// Shr returns u>>o reduced to 55 bits.
func (u U55) Shr(o U55) U55 { return U55.cast(u >> o) }

// Shl returns u<<o reduced to 55 bits.
func (u U55) Shl(o U55) U55 { return U55.cast(u << o) }

// Not returns bitwise complement of u reduced to 55 bits.
func (u U55) Not() U55 { return U55.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U55) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U55) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U55 value.
func (u U55) Signed() I55 { return I55.cast(I55(u)) }

// Clamp returns value saturated into the representable range of U55.
func (u U55) Clamp(value uint64) U55 { return uclamp[U55](value) }

// Clip masks u to a bits-wide low field.
func (u U55) Clip(bits int) U55 {
	b := 1 << (bits - 1)
	m := U55(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U55) Bit(index int) U55 {
	if index < 0 {
		index = 55 + index
	}
	return (u >> U55(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U55) SetBit(index int, v U55) {
	if index < 0 {
		index = 55 + index
	}
	bit := U55(1) << U55(index)
	*u = U55.cast((*u &^ bit) | ((v & 1) << U55(index)))
}

// Bitref returns a Range for bit index.
func (u *U55) Bitref(index int) Range[U55] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U55) Bits(lo, hi int) U55 {
	p := 55
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U55((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U55(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U55) SetBits(lo, hi int, v U55) {
	p := 55
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U55((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U55.cast((*u &^ mask) | ((v << U55(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U55) Bitsref(lo, hi int) Range[U55] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U55) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U55) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U55(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U55) Byteref(idx int) Range[U55] { return byte(u, idx) }

// I55 is an 55-bit signed integer in two's complement.
type I55 int64

// AsI55 returns a I55 representing v reduced to 55 bits.
func AsI55[T Signed](v T) I55 { return I55.cast(I55(v)) }

func (I55) nbits() int  { return 55 }
func (i I55) mask() I55 { return 0x7fffffffffffff }
func (i I55) sign() I55 { return 1 << (i.nbits() - 1) }
func (i I55) cast() I55 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 55 bits to u.
func (i *I55) Set(v I55) I55 { *i = new(I55(v)).cast(); return I55.cast(v) }

// Add returns i+o reduced to 55 bits.
func (i I55) Add(o I55) I55 { return I55.cast(i + o) }

// Sub returns i-o reduced to 55 bits.
func (i I55) Sub(o I55) I55 { return I55.cast(i - o) }

// Inc returns i+1 reduced to 55 bits.
func (i I55) Inc() I55 { return I55.cast(i + 1) }

// Dec returns i-1 reduced to 55 bits.
func (i I55) Dec() I55 { return I55.cast(i - 1) }

// Mul returns i*o reduced to 55 bits.
func (i I55) Mul(o I55) I55 { return I55.cast(i * o) }

// Div returns i divided by o.
func (i I55) Div(o I55) I55 { return I55.cast(i / o) }

// Mod returns i modulo o.
func (i I55) Mod(o I55) I55 { return I55.cast(i % o) }

// And returns i&o reduced to 55 bits.
func (i I55) And(o I55) I55 { return I55.cast(i & o) }

// Or returns i|o reduced to 55 bits.
func (i I55) Or(o I55) I55 { return I55.cast(i | o) }

// Xor returns i^o reduced to 55 bits.
func (i I55) Xor(o I55) I55 { return I55.cast(i ^ o) }

// Shr returns i>>o reduced to 55 bits.
func (i I55) Shr(o I55) I55 { return I55.cast(i >> o) }

// Shl returns i<<o reduced to 55 bits.
func (i I55) Shl(o I55) I55 { return I55.cast(i << o) }

// Not returns bitwise complement of i reduced to 55 bits.
func (i I55) Not() I55 { return I55.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I55) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I55) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U55 bit pattern.
func (i I55) Unsigned() U55 { return U55.cast(U55(i)) }

// Clamp returns value saturated into the representable range of I55.
func (i I55) Clamp(value int64) I55 { return iclamp[I55](value) }

// Clip masks i to a bits-wide low field and sign-extends to 55 bits.
func (i I55) Clip(bits int) I55 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I55(uint64(i)&m^b) - I55(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I55) Bit(index int) I55 {
	if index < 0 {
		index = 55 + index
	}
	return (i >> I55(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I55) SetBit(index int, v I55) {
	if index < 0 {
		index = 55 + index
	}
	bit := I55(1) << I55(index)
	*i = I55.cast((*i &^ bit) | ((v & 1) << I55(index)))
}

// Bitref returns a Range for bit index.
func (i *I55) Bitref(index int) Range[I55] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I55) Bits(lo, hi int) I55 {
	p := 55
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I55((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I55(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I55) SetBits(lo, hi int, v I55) {
	p := 55
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I55((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I55.cast((*i &^ mask) | ((v << I55(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I55) Bitsref(lo, hi int) Range[I55] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I55) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I55) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I55(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I55) Byteref(idx int) Range[I55] { return byte(i, idx) }

// U56 is an 56-bit unsigned integer.
type U56 uint64

// AsU56 returns a U56 representing v reduced to 56 bits.
func AsU56[T Unsigned](v T) U56 { return U56.cast(U56(v)) }

func (U56) nbits() int  { return 56 }
func (u U56) mask() U56 { return 0xffffffffffffff }
func (u U56) cast() U56 { return u & u.mask() }

// Set assigns v reduced to 56 bits to u.
func (u *U56) Set(v U56) U56 { *u = new(U56(v)).cast(); return U56.cast(v) }

// Add returns u+o reduced to 56 bits.
func (u U56) Add(o U56) U56 { return U56.cast(u + o) }

// Sub returns u-o reduced to 56 bits.
func (u U56) Sub(o U56) U56 { return U56.cast(u - o) }

// Inc returns u+1 reduced to 56 bits.
func (u U56) Inc() U56 { return U56.cast(u + 1) }

// Dec returns u-1 reduced to 56 bits.
func (u U56) Dec() U56 { return U56.cast(u - 1) }

// Mul returns u*o reduced to 56 bits.
func (u U56) Mul(o U56) U56 { return U56.cast(u * o) }

// Div returns u divided by o.
func (u U56) Div(o U56) U56 { return U56.cast(u / o) }

// Mod returns u modulo o.
func (u U56) Mod(o U56) U56 { return U56.cast(u % o) }

// And returns u&o reduced to 56 bits.
func (u U56) And(o U56) U56 { return U56.cast(u & o) }

// Or returns u|o reduced to 56 bits.
func (u U56) Or(o U56) U56 { return U56.cast(u | o) }

// Xor returns u^o reduced to 56 bits.
func (u U56) Xor(o U56) U56 { return U56.cast(u ^ o) }

// Shr returns u>>o reduced to 56 bits.
func (u U56) Shr(o U56) U56 { return U56.cast(u >> o) }

// Shl returns u<<o reduced to 56 bits.
func (u U56) Shl(o U56) U56 { return U56.cast(u << o) }

// Not returns bitwise complement of u reduced to 56 bits.
func (u U56) Not() U56 { return U56.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U56) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U56) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U56 value.
func (u U56) Signed() I56 { return I56.cast(I56(u)) }

// Clamp returns value saturated into the representable range of U56.
func (u U56) Clamp(value uint64) U56 { return uclamp[U56](value) }

// Clip masks u to a bits-wide low field.
func (u U56) Clip(bits int) U56 {
	b := 1 << (bits - 1)
	m := U56(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U56) Bit(index int) U56 {
	if index < 0 {
		index = 56 + index
	}
	return (u >> U56(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U56) SetBit(index int, v U56) {
	if index < 0 {
		index = 56 + index
	}
	bit := U56(1) << U56(index)
	*u = U56.cast((*u &^ bit) | ((v & 1) << U56(index)))
}

// Bitref returns a Range for bit index.
func (u *U56) Bitref(index int) Range[U56] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U56) Bits(lo, hi int) U56 {
	p := 56
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U56((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U56(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U56) SetBits(lo, hi int, v U56) {
	p := 56
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U56((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U56.cast((*u &^ mask) | ((v << U56(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U56) Bitsref(lo, hi int) Range[U56] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U56) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U56) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U56(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U56) Byteref(idx int) Range[U56] { return byte(u, idx) }

// I56 is an 56-bit signed integer in two's complement.
type I56 int64

// AsI56 returns a I56 representing v reduced to 56 bits.
func AsI56[T Signed](v T) I56 { return I56.cast(I56(v)) }

func (I56) nbits() int  { return 56 }
func (i I56) mask() I56 { return 0xffffffffffffff }
func (i I56) sign() I56 { return 1 << (i.nbits() - 1) }
func (i I56) cast() I56 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 56 bits to u.
func (i *I56) Set(v I56) I56 { *i = new(I56(v)).cast(); return I56.cast(v) }

// Add returns i+o reduced to 56 bits.
func (i I56) Add(o I56) I56 { return I56.cast(i + o) }

// Sub returns i-o reduced to 56 bits.
func (i I56) Sub(o I56) I56 { return I56.cast(i - o) }

// Inc returns i+1 reduced to 56 bits.
func (i I56) Inc() I56 { return I56.cast(i + 1) }

// Dec returns i-1 reduced to 56 bits.
func (i I56) Dec() I56 { return I56.cast(i - 1) }

// Mul returns i*o reduced to 56 bits.
func (i I56) Mul(o I56) I56 { return I56.cast(i * o) }

// Div returns i divided by o.
func (i I56) Div(o I56) I56 { return I56.cast(i / o) }

// Mod returns i modulo o.
func (i I56) Mod(o I56) I56 { return I56.cast(i % o) }

// And returns i&o reduced to 56 bits.
func (i I56) And(o I56) I56 { return I56.cast(i & o) }

// Or returns i|o reduced to 56 bits.
func (i I56) Or(o I56) I56 { return I56.cast(i | o) }

// Xor returns i^o reduced to 56 bits.
func (i I56) Xor(o I56) I56 { return I56.cast(i ^ o) }

// Shr returns i>>o reduced to 56 bits.
func (i I56) Shr(o I56) I56 { return I56.cast(i >> o) }

// Shl returns i<<o reduced to 56 bits.
func (i I56) Shl(o I56) I56 { return I56.cast(i << o) }

// Not returns bitwise complement of i reduced to 56 bits.
func (i I56) Not() I56 { return I56.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I56) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I56) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U56 bit pattern.
func (i I56) Unsigned() U56 { return U56.cast(U56(i)) }

// Clamp returns value saturated into the representable range of I56.
func (i I56) Clamp(value int64) I56 { return iclamp[I56](value) }

// Clip masks i to a bits-wide low field and sign-extends to 56 bits.
func (i I56) Clip(bits int) I56 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I56(uint64(i)&m^b) - I56(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I56) Bit(index int) I56 {
	if index < 0 {
		index = 56 + index
	}
	return (i >> I56(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I56) SetBit(index int, v I56) {
	if index < 0 {
		index = 56 + index
	}
	bit := I56(1) << I56(index)
	*i = I56.cast((*i &^ bit) | ((v & 1) << I56(index)))
}

// Bitref returns a Range for bit index.
func (i *I56) Bitref(index int) Range[I56] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I56) Bits(lo, hi int) I56 {
	p := 56
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I56((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I56(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I56) SetBits(lo, hi int, v I56) {
	p := 56
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I56((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I56.cast((*i &^ mask) | ((v << I56(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I56) Bitsref(lo, hi int) Range[I56] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I56) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I56) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I56(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I56) Byteref(idx int) Range[I56] { return byte(i, idx) }

// U57 is an 57-bit unsigned integer.
type U57 uint64

// AsU57 returns a U57 representing v reduced to 57 bits.
func AsU57[T Unsigned](v T) U57 { return U57.cast(U57(v)) }

func (U57) nbits() int  { return 57 }
func (u U57) mask() U57 { return 0x1ffffffffffffff }
func (u U57) cast() U57 { return u & u.mask() }

// Set assigns v reduced to 57 bits to u.
func (u *U57) Set(v U57) U57 { *u = new(U57(v)).cast(); return U57.cast(v) }

// Add returns u+o reduced to 57 bits.
func (u U57) Add(o U57) U57 { return U57.cast(u + o) }

// Sub returns u-o reduced to 57 bits.
func (u U57) Sub(o U57) U57 { return U57.cast(u - o) }

// Inc returns u+1 reduced to 57 bits.
func (u U57) Inc() U57 { return U57.cast(u + 1) }

// Dec returns u-1 reduced to 57 bits.
func (u U57) Dec() U57 { return U57.cast(u - 1) }

// Mul returns u*o reduced to 57 bits.
func (u U57) Mul(o U57) U57 { return U57.cast(u * o) }

// Div returns u divided by o.
func (u U57) Div(o U57) U57 { return U57.cast(u / o) }

// Mod returns u modulo o.
func (u U57) Mod(o U57) U57 { return U57.cast(u % o) }

// And returns u&o reduced to 57 bits.
func (u U57) And(o U57) U57 { return U57.cast(u & o) }

// Or returns u|o reduced to 57 bits.
func (u U57) Or(o U57) U57 { return U57.cast(u | o) }

// Xor returns u^o reduced to 57 bits.
func (u U57) Xor(o U57) U57 { return U57.cast(u ^ o) }

// Shr returns u>>o reduced to 57 bits.
func (u U57) Shr(o U57) U57 { return U57.cast(u >> o) }

// Shl returns u<<o reduced to 57 bits.
func (u U57) Shl(o U57) U57 { return U57.cast(u << o) }

// Not returns bitwise complement of u reduced to 57 bits.
func (u U57) Not() U57 { return U57.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U57) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U57) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U57 value.
func (u U57) Signed() I57 { return I57.cast(I57(u)) }

// Clamp returns value saturated into the representable range of U57.
func (u U57) Clamp(value uint64) U57 { return uclamp[U57](value) }

// Clip masks u to a bits-wide low field.
func (u U57) Clip(bits int) U57 {
	b := 1 << (bits - 1)
	m := U57(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U57) Bit(index int) U57 {
	if index < 0 {
		index = 57 + index
	}
	return (u >> U57(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U57) SetBit(index int, v U57) {
	if index < 0 {
		index = 57 + index
	}
	bit := U57(1) << U57(index)
	*u = U57.cast((*u &^ bit) | ((v & 1) << U57(index)))
}

// Bitref returns a Range for bit index.
func (u *U57) Bitref(index int) Range[U57] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U57) Bits(lo, hi int) U57 {
	p := 57
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U57((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U57(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U57) SetBits(lo, hi int, v U57) {
	p := 57
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U57((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U57.cast((*u &^ mask) | ((v << U57(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U57) Bitsref(lo, hi int) Range[U57] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U57) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U57) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U57(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U57) Byteref(idx int) Range[U57] { return byte(u, idx) }

// I57 is an 57-bit signed integer in two's complement.
type I57 int64

// AsI57 returns a I57 representing v reduced to 57 bits.
func AsI57[T Signed](v T) I57 { return I57.cast(I57(v)) }

func (I57) nbits() int  { return 57 }
func (i I57) mask() I57 { return 0x1ffffffffffffff }
func (i I57) sign() I57 { return 1 << (i.nbits() - 1) }
func (i I57) cast() I57 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 57 bits to u.
func (i *I57) Set(v I57) I57 { *i = new(I57(v)).cast(); return I57.cast(v) }

// Add returns i+o reduced to 57 bits.
func (i I57) Add(o I57) I57 { return I57.cast(i + o) }

// Sub returns i-o reduced to 57 bits.
func (i I57) Sub(o I57) I57 { return I57.cast(i - o) }

// Inc returns i+1 reduced to 57 bits.
func (i I57) Inc() I57 { return I57.cast(i + 1) }

// Dec returns i-1 reduced to 57 bits.
func (i I57) Dec() I57 { return I57.cast(i - 1) }

// Mul returns i*o reduced to 57 bits.
func (i I57) Mul(o I57) I57 { return I57.cast(i * o) }

// Div returns i divided by o.
func (i I57) Div(o I57) I57 { return I57.cast(i / o) }

// Mod returns i modulo o.
func (i I57) Mod(o I57) I57 { return I57.cast(i % o) }

// And returns i&o reduced to 57 bits.
func (i I57) And(o I57) I57 { return I57.cast(i & o) }

// Or returns i|o reduced to 57 bits.
func (i I57) Or(o I57) I57 { return I57.cast(i | o) }

// Xor returns i^o reduced to 57 bits.
func (i I57) Xor(o I57) I57 { return I57.cast(i ^ o) }

// Shr returns i>>o reduced to 57 bits.
func (i I57) Shr(o I57) I57 { return I57.cast(i >> o) }

// Shl returns i<<o reduced to 57 bits.
func (i I57) Shl(o I57) I57 { return I57.cast(i << o) }

// Not returns bitwise complement of i reduced to 57 bits.
func (i I57) Not() I57 { return I57.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I57) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I57) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U57 bit pattern.
func (i I57) Unsigned() U57 { return U57.cast(U57(i)) }

// Clamp returns value saturated into the representable range of I57.
func (i I57) Clamp(value int64) I57 { return iclamp[I57](value) }

// Clip masks i to a bits-wide low field and sign-extends to 57 bits.
func (i I57) Clip(bits int) I57 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I57(uint64(i)&m^b) - I57(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I57) Bit(index int) I57 {
	if index < 0 {
		index = 57 + index
	}
	return (i >> I57(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I57) SetBit(index int, v I57) {
	if index < 0 {
		index = 57 + index
	}
	bit := I57(1) << I57(index)
	*i = I57.cast((*i &^ bit) | ((v & 1) << I57(index)))
}

// Bitref returns a Range for bit index.
func (i *I57) Bitref(index int) Range[I57] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I57) Bits(lo, hi int) I57 {
	p := 57
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I57((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I57(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I57) SetBits(lo, hi int, v I57) {
	p := 57
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I57((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I57.cast((*i &^ mask) | ((v << I57(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I57) Bitsref(lo, hi int) Range[I57] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I57) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I57) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I57(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I57) Byteref(idx int) Range[I57] { return byte(i, idx) }

// U58 is an 58-bit unsigned integer.
type U58 uint64

// AsU58 returns a U58 representing v reduced to 58 bits.
func AsU58[T Unsigned](v T) U58 { return U58.cast(U58(v)) }

func (U58) nbits() int  { return 58 }
func (u U58) mask() U58 { return 0x3ffffffffffffff }
func (u U58) cast() U58 { return u & u.mask() }

// Set assigns v reduced to 58 bits to u.
func (u *U58) Set(v U58) U58 { *u = new(U58(v)).cast(); return U58.cast(v) }

// Add returns u+o reduced to 58 bits.
func (u U58) Add(o U58) U58 { return U58.cast(u + o) }

// Sub returns u-o reduced to 58 bits.
func (u U58) Sub(o U58) U58 { return U58.cast(u - o) }

// Inc returns u+1 reduced to 58 bits.
func (u U58) Inc() U58 { return U58.cast(u + 1) }

// Dec returns u-1 reduced to 58 bits.
func (u U58) Dec() U58 { return U58.cast(u - 1) }

// Mul returns u*o reduced to 58 bits.
func (u U58) Mul(o U58) U58 { return U58.cast(u * o) }

// Div returns u divided by o.
func (u U58) Div(o U58) U58 { return U58.cast(u / o) }

// Mod returns u modulo o.
func (u U58) Mod(o U58) U58 { return U58.cast(u % o) }

// And returns u&o reduced to 58 bits.
func (u U58) And(o U58) U58 { return U58.cast(u & o) }

// Or returns u|o reduced to 58 bits.
func (u U58) Or(o U58) U58 { return U58.cast(u | o) }

// Xor returns u^o reduced to 58 bits.
func (u U58) Xor(o U58) U58 { return U58.cast(u ^ o) }

// Shr returns u>>o reduced to 58 bits.
func (u U58) Shr(o U58) U58 { return U58.cast(u >> o) }

// Shl returns u<<o reduced to 58 bits.
func (u U58) Shl(o U58) U58 { return U58.cast(u << o) }

// Not returns bitwise complement of u reduced to 58 bits.
func (u U58) Not() U58 { return U58.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U58) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U58) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U58 value.
func (u U58) Signed() I58 { return I58.cast(I58(u)) }

// Clamp returns value saturated into the representable range of U58.
func (u U58) Clamp(value uint64) U58 { return uclamp[U58](value) }

// Clip masks u to a bits-wide low field.
func (u U58) Clip(bits int) U58 {
	b := 1 << (bits - 1)
	m := U58(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U58) Bit(index int) U58 {
	if index < 0 {
		index = 58 + index
	}
	return (u >> U58(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U58) SetBit(index int, v U58) {
	if index < 0 {
		index = 58 + index
	}
	bit := U58(1) << U58(index)
	*u = U58.cast((*u &^ bit) | ((v & 1) << U58(index)))
}

// Bitref returns a Range for bit index.
func (u *U58) Bitref(index int) Range[U58] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U58) Bits(lo, hi int) U58 {
	p := 58
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U58((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U58(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U58) SetBits(lo, hi int, v U58) {
	p := 58
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U58((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U58.cast((*u &^ mask) | ((v << U58(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U58) Bitsref(lo, hi int) Range[U58] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U58) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U58) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U58(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U58) Byteref(idx int) Range[U58] { return byte(u, idx) }

// I58 is an 58-bit signed integer in two's complement.
type I58 int64

// AsI58 returns a I58 representing v reduced to 58 bits.
func AsI58[T Signed](v T) I58 { return I58.cast(I58(v)) }

func (I58) nbits() int  { return 58 }
func (i I58) mask() I58 { return 0x3ffffffffffffff }
func (i I58) sign() I58 { return 1 << (i.nbits() - 1) }
func (i I58) cast() I58 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 58 bits to u.
func (i *I58) Set(v I58) I58 { *i = new(I58(v)).cast(); return I58.cast(v) }

// Add returns i+o reduced to 58 bits.
func (i I58) Add(o I58) I58 { return I58.cast(i + o) }

// Sub returns i-o reduced to 58 bits.
func (i I58) Sub(o I58) I58 { return I58.cast(i - o) }

// Inc returns i+1 reduced to 58 bits.
func (i I58) Inc() I58 { return I58.cast(i + 1) }

// Dec returns i-1 reduced to 58 bits.
func (i I58) Dec() I58 { return I58.cast(i - 1) }

// Mul returns i*o reduced to 58 bits.
func (i I58) Mul(o I58) I58 { return I58.cast(i * o) }

// Div returns i divided by o.
func (i I58) Div(o I58) I58 { return I58.cast(i / o) }

// Mod returns i modulo o.
func (i I58) Mod(o I58) I58 { return I58.cast(i % o) }

// And returns i&o reduced to 58 bits.
func (i I58) And(o I58) I58 { return I58.cast(i & o) }

// Or returns i|o reduced to 58 bits.
func (i I58) Or(o I58) I58 { return I58.cast(i | o) }

// Xor returns i^o reduced to 58 bits.
func (i I58) Xor(o I58) I58 { return I58.cast(i ^ o) }

// Shr returns i>>o reduced to 58 bits.
func (i I58) Shr(o I58) I58 { return I58.cast(i >> o) }

// Shl returns i<<o reduced to 58 bits.
func (i I58) Shl(o I58) I58 { return I58.cast(i << o) }

// Not returns bitwise complement of i reduced to 58 bits.
func (i I58) Not() I58 { return I58.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I58) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I58) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U58 bit pattern.
func (i I58) Unsigned() U58 { return U58.cast(U58(i)) }

// Clamp returns value saturated into the representable range of I58.
func (i I58) Clamp(value int64) I58 { return iclamp[I58](value) }

// Clip masks i to a bits-wide low field and sign-extends to 58 bits.
func (i I58) Clip(bits int) I58 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I58(uint64(i)&m^b) - I58(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I58) Bit(index int) I58 {
	if index < 0 {
		index = 58 + index
	}
	return (i >> I58(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I58) SetBit(index int, v I58) {
	if index < 0 {
		index = 58 + index
	}
	bit := I58(1) << I58(index)
	*i = I58.cast((*i &^ bit) | ((v & 1) << I58(index)))
}

// Bitref returns a Range for bit index.
func (i *I58) Bitref(index int) Range[I58] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I58) Bits(lo, hi int) I58 {
	p := 58
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I58((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I58(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I58) SetBits(lo, hi int, v I58) {
	p := 58
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I58((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I58.cast((*i &^ mask) | ((v << I58(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I58) Bitsref(lo, hi int) Range[I58] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I58) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I58) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I58(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I58) Byteref(idx int) Range[I58] { return byte(i, idx) }

// U59 is an 59-bit unsigned integer.
type U59 uint64

// AsU59 returns a U59 representing v reduced to 59 bits.
func AsU59[T Unsigned](v T) U59 { return U59.cast(U59(v)) }

func (U59) nbits() int  { return 59 }
func (u U59) mask() U59 { return 0x7ffffffffffffff }
func (u U59) cast() U59 { return u & u.mask() }

// Set assigns v reduced to 59 bits to u.
func (u *U59) Set(v U59) U59 { *u = new(U59(v)).cast(); return U59.cast(v) }

// Add returns u+o reduced to 59 bits.
func (u U59) Add(o U59) U59 { return U59.cast(u + o) }

// Sub returns u-o reduced to 59 bits.
func (u U59) Sub(o U59) U59 { return U59.cast(u - o) }

// Inc returns u+1 reduced to 59 bits.
func (u U59) Inc() U59 { return U59.cast(u + 1) }

// Dec returns u-1 reduced to 59 bits.
func (u U59) Dec() U59 { return U59.cast(u - 1) }

// Mul returns u*o reduced to 59 bits.
func (u U59) Mul(o U59) U59 { return U59.cast(u * o) }

// Div returns u divided by o.
func (u U59) Div(o U59) U59 { return U59.cast(u / o) }

// Mod returns u modulo o.
func (u U59) Mod(o U59) U59 { return U59.cast(u % o) }

// And returns u&o reduced to 59 bits.
func (u U59) And(o U59) U59 { return U59.cast(u & o) }

// Or returns u|o reduced to 59 bits.
func (u U59) Or(o U59) U59 { return U59.cast(u | o) }

// Xor returns u^o reduced to 59 bits.
func (u U59) Xor(o U59) U59 { return U59.cast(u ^ o) }

// Shr returns u>>o reduced to 59 bits.
func (u U59) Shr(o U59) U59 { return U59.cast(u >> o) }

// Shl returns u<<o reduced to 59 bits.
func (u U59) Shl(o U59) U59 { return U59.cast(u << o) }

// Not returns bitwise complement of u reduced to 59 bits.
func (u U59) Not() U59 { return U59.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U59) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U59) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U59 value.
func (u U59) Signed() I59 { return I59.cast(I59(u)) }

// Clamp returns value saturated into the representable range of U59.
func (u U59) Clamp(value uint64) U59 { return uclamp[U59](value) }

// Clip masks u to a bits-wide low field.
func (u U59) Clip(bits int) U59 {
	b := 1 << (bits - 1)
	m := U59(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U59) Bit(index int) U59 {
	if index < 0 {
		index = 59 + index
	}
	return (u >> U59(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U59) SetBit(index int, v U59) {
	if index < 0 {
		index = 59 + index
	}
	bit := U59(1) << U59(index)
	*u = U59.cast((*u &^ bit) | ((v & 1) << U59(index)))
}

// Bitref returns a Range for bit index.
func (u *U59) Bitref(index int) Range[U59] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U59) Bits(lo, hi int) U59 {
	p := 59
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U59((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U59(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U59) SetBits(lo, hi int, v U59) {
	p := 59
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U59((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U59.cast((*u &^ mask) | ((v << U59(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U59) Bitsref(lo, hi int) Range[U59] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U59) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U59) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U59(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U59) Byteref(idx int) Range[U59] { return byte(u, idx) }

// I59 is an 59-bit signed integer in two's complement.
type I59 int64

// AsI59 returns a I59 representing v reduced to 59 bits.
func AsI59[T Signed](v T) I59 { return I59.cast(I59(v)) }

func (I59) nbits() int  { return 59 }
func (i I59) mask() I59 { return 0x7ffffffffffffff }
func (i I59) sign() I59 { return 1 << (i.nbits() - 1) }
func (i I59) cast() I59 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 59 bits to u.
func (i *I59) Set(v I59) I59 { *i = new(I59(v)).cast(); return I59.cast(v) }

// Add returns i+o reduced to 59 bits.
func (i I59) Add(o I59) I59 { return I59.cast(i + o) }

// Sub returns i-o reduced to 59 bits.
func (i I59) Sub(o I59) I59 { return I59.cast(i - o) }

// Inc returns i+1 reduced to 59 bits.
func (i I59) Inc() I59 { return I59.cast(i + 1) }

// Dec returns i-1 reduced to 59 bits.
func (i I59) Dec() I59 { return I59.cast(i - 1) }

// Mul returns i*o reduced to 59 bits.
func (i I59) Mul(o I59) I59 { return I59.cast(i * o) }

// Div returns i divided by o.
func (i I59) Div(o I59) I59 { return I59.cast(i / o) }

// Mod returns i modulo o.
func (i I59) Mod(o I59) I59 { return I59.cast(i % o) }

// And returns i&o reduced to 59 bits.
func (i I59) And(o I59) I59 { return I59.cast(i & o) }

// Or returns i|o reduced to 59 bits.
func (i I59) Or(o I59) I59 { return I59.cast(i | o) }

// Xor returns i^o reduced to 59 bits.
func (i I59) Xor(o I59) I59 { return I59.cast(i ^ o) }

// Shr returns i>>o reduced to 59 bits.
func (i I59) Shr(o I59) I59 { return I59.cast(i >> o) }

// Shl returns i<<o reduced to 59 bits.
func (i I59) Shl(o I59) I59 { return I59.cast(i << o) }

// Not returns bitwise complement of i reduced to 59 bits.
func (i I59) Not() I59 { return I59.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I59) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I59) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U59 bit pattern.
func (i I59) Unsigned() U59 { return U59.cast(U59(i)) }

// Clamp returns value saturated into the representable range of I59.
func (i I59) Clamp(value int64) I59 { return iclamp[I59](value) }

// Clip masks i to a bits-wide low field and sign-extends to 59 bits.
func (i I59) Clip(bits int) I59 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I59(uint64(i)&m^b) - I59(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I59) Bit(index int) I59 {
	if index < 0 {
		index = 59 + index
	}
	return (i >> I59(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I59) SetBit(index int, v I59) {
	if index < 0 {
		index = 59 + index
	}
	bit := I59(1) << I59(index)
	*i = I59.cast((*i &^ bit) | ((v & 1) << I59(index)))
}

// Bitref returns a Range for bit index.
func (i *I59) Bitref(index int) Range[I59] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I59) Bits(lo, hi int) I59 {
	p := 59
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I59((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I59(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I59) SetBits(lo, hi int, v I59) {
	p := 59
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I59((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I59.cast((*i &^ mask) | ((v << I59(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I59) Bitsref(lo, hi int) Range[I59] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I59) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I59) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I59(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I59) Byteref(idx int) Range[I59] { return byte(i, idx) }

// U60 is an 60-bit unsigned integer.
type U60 uint64

// AsU60 returns a U60 representing v reduced to 60 bits.
func AsU60[T Unsigned](v T) U60 { return U60.cast(U60(v)) }

func (U60) nbits() int  { return 60 }
func (u U60) mask() U60 { return 0xfffffffffffffff }
func (u U60) cast() U60 { return u & u.mask() }

// Set assigns v reduced to 60 bits to u.
func (u *U60) Set(v U60) U60 { *u = new(U60(v)).cast(); return U60.cast(v) }

// Add returns u+o reduced to 60 bits.
func (u U60) Add(o U60) U60 { return U60.cast(u + o) }

// Sub returns u-o reduced to 60 bits.
func (u U60) Sub(o U60) U60 { return U60.cast(u - o) }

// Inc returns u+1 reduced to 60 bits.
func (u U60) Inc() U60 { return U60.cast(u + 1) }

// Dec returns u-1 reduced to 60 bits.
func (u U60) Dec() U60 { return U60.cast(u - 1) }

// Mul returns u*o reduced to 60 bits.
func (u U60) Mul(o U60) U60 { return U60.cast(u * o) }

// Div returns u divided by o.
func (u U60) Div(o U60) U60 { return U60.cast(u / o) }

// Mod returns u modulo o.
func (u U60) Mod(o U60) U60 { return U60.cast(u % o) }

// And returns u&o reduced to 60 bits.
func (u U60) And(o U60) U60 { return U60.cast(u & o) }

// Or returns u|o reduced to 60 bits.
func (u U60) Or(o U60) U60 { return U60.cast(u | o) }

// Xor returns u^o reduced to 60 bits.
func (u U60) Xor(o U60) U60 { return U60.cast(u ^ o) }

// Shr returns u>>o reduced to 60 bits.
func (u U60) Shr(o U60) U60 { return U60.cast(u >> o) }

// Shl returns u<<o reduced to 60 bits.
func (u U60) Shl(o U60) U60 { return U60.cast(u << o) }

// Not returns bitwise complement of u reduced to 60 bits.
func (u U60) Not() U60 { return U60.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U60) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U60) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U60 value.
func (u U60) Signed() I60 { return I60.cast(I60(u)) }

// Clamp returns value saturated into the representable range of U60.
func (u U60) Clamp(value uint64) U60 { return uclamp[U60](value) }

// Clip masks u to a bits-wide low field.
func (u U60) Clip(bits int) U60 {
	b := 1 << (bits - 1)
	m := U60(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U60) Bit(index int) U60 {
	if index < 0 {
		index = 60 + index
	}
	return (u >> U60(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U60) SetBit(index int, v U60) {
	if index < 0 {
		index = 60 + index
	}
	bit := U60(1) << U60(index)
	*u = U60.cast((*u &^ bit) | ((v & 1) << U60(index)))
}

// Bitref returns a Range for bit index.
func (u *U60) Bitref(index int) Range[U60] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U60) Bits(lo, hi int) U60 {
	p := 60
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U60((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U60(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U60) SetBits(lo, hi int, v U60) {
	p := 60
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U60((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U60.cast((*u &^ mask) | ((v << U60(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U60) Bitsref(lo, hi int) Range[U60] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U60) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U60) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U60(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U60) Byteref(idx int) Range[U60] { return byte(u, idx) }

// I60 is an 60-bit signed integer in two's complement.
type I60 int64

// AsI60 returns a I60 representing v reduced to 60 bits.
func AsI60[T Signed](v T) I60 { return I60.cast(I60(v)) }

func (I60) nbits() int  { return 60 }
func (i I60) mask() I60 { return 0xfffffffffffffff }
func (i I60) sign() I60 { return 1 << (i.nbits() - 1) }
func (i I60) cast() I60 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 60 bits to u.
func (i *I60) Set(v I60) I60 { *i = new(I60(v)).cast(); return I60.cast(v) }

// Add returns i+o reduced to 60 bits.
func (i I60) Add(o I60) I60 { return I60.cast(i + o) }

// Sub returns i-o reduced to 60 bits.
func (i I60) Sub(o I60) I60 { return I60.cast(i - o) }

// Inc returns i+1 reduced to 60 bits.
func (i I60) Inc() I60 { return I60.cast(i + 1) }

// Dec returns i-1 reduced to 60 bits.
func (i I60) Dec() I60 { return I60.cast(i - 1) }

// Mul returns i*o reduced to 60 bits.
func (i I60) Mul(o I60) I60 { return I60.cast(i * o) }

// Div returns i divided by o.
func (i I60) Div(o I60) I60 { return I60.cast(i / o) }

// Mod returns i modulo o.
func (i I60) Mod(o I60) I60 { return I60.cast(i % o) }

// And returns i&o reduced to 60 bits.
func (i I60) And(o I60) I60 { return I60.cast(i & o) }

// Or returns i|o reduced to 60 bits.
func (i I60) Or(o I60) I60 { return I60.cast(i | o) }

// Xor returns i^o reduced to 60 bits.
func (i I60) Xor(o I60) I60 { return I60.cast(i ^ o) }

// Shr returns i>>o reduced to 60 bits.
func (i I60) Shr(o I60) I60 { return I60.cast(i >> o) }

// Shl returns i<<o reduced to 60 bits.
func (i I60) Shl(o I60) I60 { return I60.cast(i << o) }

// Not returns bitwise complement of i reduced to 60 bits.
func (i I60) Not() I60 { return I60.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I60) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I60) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U60 bit pattern.
func (i I60) Unsigned() U60 { return U60.cast(U60(i)) }

// Clamp returns value saturated into the representable range of I60.
func (i I60) Clamp(value int64) I60 { return iclamp[I60](value) }

// Clip masks i to a bits-wide low field and sign-extends to 60 bits.
func (i I60) Clip(bits int) I60 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I60(uint64(i)&m^b) - I60(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I60) Bit(index int) I60 {
	if index < 0 {
		index = 60 + index
	}
	return (i >> I60(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I60) SetBit(index int, v I60) {
	if index < 0 {
		index = 60 + index
	}
	bit := I60(1) << I60(index)
	*i = I60.cast((*i &^ bit) | ((v & 1) << I60(index)))
}

// Bitref returns a Range for bit index.
func (i *I60) Bitref(index int) Range[I60] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I60) Bits(lo, hi int) I60 {
	p := 60
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I60((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I60(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I60) SetBits(lo, hi int, v I60) {
	p := 60
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I60((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I60.cast((*i &^ mask) | ((v << I60(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I60) Bitsref(lo, hi int) Range[I60] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I60) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I60) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I60(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I60) Byteref(idx int) Range[I60] { return byte(i, idx) }

// U61 is an 61-bit unsigned integer.
type U61 uint64

// AsU61 returns a U61 representing v reduced to 61 bits.
func AsU61[T Unsigned](v T) U61 { return U61.cast(U61(v)) }

func (U61) nbits() int  { return 61 }
func (u U61) mask() U61 { return 0x1fffffffffffffff }
func (u U61) cast() U61 { return u & u.mask() }

// Set assigns v reduced to 61 bits to u.
func (u *U61) Set(v U61) U61 { *u = new(U61(v)).cast(); return U61.cast(v) }

// Add returns u+o reduced to 61 bits.
func (u U61) Add(o U61) U61 { return U61.cast(u + o) }

// Sub returns u-o reduced to 61 bits.
func (u U61) Sub(o U61) U61 { return U61.cast(u - o) }

// Inc returns u+1 reduced to 61 bits.
func (u U61) Inc() U61 { return U61.cast(u + 1) }

// Dec returns u-1 reduced to 61 bits.
func (u U61) Dec() U61 { return U61.cast(u - 1) }

// Mul returns u*o reduced to 61 bits.
func (u U61) Mul(o U61) U61 { return U61.cast(u * o) }

// Div returns u divided by o.
func (u U61) Div(o U61) U61 { return U61.cast(u / o) }

// Mod returns u modulo o.
func (u U61) Mod(o U61) U61 { return U61.cast(u % o) }

// And returns u&o reduced to 61 bits.
func (u U61) And(o U61) U61 { return U61.cast(u & o) }

// Or returns u|o reduced to 61 bits.
func (u U61) Or(o U61) U61 { return U61.cast(u | o) }

// Xor returns u^o reduced to 61 bits.
func (u U61) Xor(o U61) U61 { return U61.cast(u ^ o) }

// Shr returns u>>o reduced to 61 bits.
func (u U61) Shr(o U61) U61 { return U61.cast(u >> o) }

// Shl returns u<<o reduced to 61 bits.
func (u U61) Shl(o U61) U61 { return U61.cast(u << o) }

// Not returns bitwise complement of u reduced to 61 bits.
func (u U61) Not() U61 { return U61.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U61) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U61) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U61 value.
func (u U61) Signed() I61 { return I61.cast(I61(u)) }

// Clamp returns value saturated into the representable range of U61.
func (u U61) Clamp(value uint64) U61 { return uclamp[U61](value) }

// Clip masks u to a bits-wide low field.
func (u U61) Clip(bits int) U61 {
	b := 1 << (bits - 1)
	m := U61(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U61) Bit(index int) U61 {
	if index < 0 {
		index = 61 + index
	}
	return (u >> U61(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U61) SetBit(index int, v U61) {
	if index < 0 {
		index = 61 + index
	}
	bit := U61(1) << U61(index)
	*u = U61.cast((*u &^ bit) | ((v & 1) << U61(index)))
}

// Bitref returns a Range for bit index.
func (u *U61) Bitref(index int) Range[U61] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U61) Bits(lo, hi int) U61 {
	p := 61
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U61((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U61(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U61) SetBits(lo, hi int, v U61) {
	p := 61
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U61((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U61.cast((*u &^ mask) | ((v << U61(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U61) Bitsref(lo, hi int) Range[U61] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U61) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U61) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U61(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U61) Byteref(idx int) Range[U61] { return byte(u, idx) }

// I61 is an 61-bit signed integer in two's complement.
type I61 int64

// AsI61 returns a I61 representing v reduced to 61 bits.
func AsI61[T Signed](v T) I61 { return I61.cast(I61(v)) }

func (I61) nbits() int  { return 61 }
func (i I61) mask() I61 { return 0x1fffffffffffffff }
func (i I61) sign() I61 { return 1 << (i.nbits() - 1) }
func (i I61) cast() I61 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 61 bits to u.
func (i *I61) Set(v I61) I61 { *i = new(I61(v)).cast(); return I61.cast(v) }

// Add returns i+o reduced to 61 bits.
func (i I61) Add(o I61) I61 { return I61.cast(i + o) }

// Sub returns i-o reduced to 61 bits.
func (i I61) Sub(o I61) I61 { return I61.cast(i - o) }

// Inc returns i+1 reduced to 61 bits.
func (i I61) Inc() I61 { return I61.cast(i + 1) }

// Dec returns i-1 reduced to 61 bits.
func (i I61) Dec() I61 { return I61.cast(i - 1) }

// Mul returns i*o reduced to 61 bits.
func (i I61) Mul(o I61) I61 { return I61.cast(i * o) }

// Div returns i divided by o.
func (i I61) Div(o I61) I61 { return I61.cast(i / o) }

// Mod returns i modulo o.
func (i I61) Mod(o I61) I61 { return I61.cast(i % o) }

// And returns i&o reduced to 61 bits.
func (i I61) And(o I61) I61 { return I61.cast(i & o) }

// Or returns i|o reduced to 61 bits.
func (i I61) Or(o I61) I61 { return I61.cast(i | o) }

// Xor returns i^o reduced to 61 bits.
func (i I61) Xor(o I61) I61 { return I61.cast(i ^ o) }

// Shr returns i>>o reduced to 61 bits.
func (i I61) Shr(o I61) I61 { return I61.cast(i >> o) }

// Shl returns i<<o reduced to 61 bits.
func (i I61) Shl(o I61) I61 { return I61.cast(i << o) }

// Not returns bitwise complement of i reduced to 61 bits.
func (i I61) Not() I61 { return I61.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I61) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I61) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U61 bit pattern.
func (i I61) Unsigned() U61 { return U61.cast(U61(i)) }

// Clamp returns value saturated into the representable range of I61.
func (i I61) Clamp(value int64) I61 { return iclamp[I61](value) }

// Clip masks i to a bits-wide low field and sign-extends to 61 bits.
func (i I61) Clip(bits int) I61 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I61(uint64(i)&m^b) - I61(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I61) Bit(index int) I61 {
	if index < 0 {
		index = 61 + index
	}
	return (i >> I61(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I61) SetBit(index int, v I61) {
	if index < 0 {
		index = 61 + index
	}
	bit := I61(1) << I61(index)
	*i = I61.cast((*i &^ bit) | ((v & 1) << I61(index)))
}

// Bitref returns a Range for bit index.
func (i *I61) Bitref(index int) Range[I61] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I61) Bits(lo, hi int) I61 {
	p := 61
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I61((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I61(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I61) SetBits(lo, hi int, v I61) {
	p := 61
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I61((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I61.cast((*i &^ mask) | ((v << I61(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I61) Bitsref(lo, hi int) Range[I61] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I61) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I61) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I61(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I61) Byteref(idx int) Range[I61] { return byte(i, idx) }

// U62 is an 62-bit unsigned integer.
type U62 uint64

// AsU62 returns a U62 representing v reduced to 62 bits.
func AsU62[T Unsigned](v T) U62 { return U62.cast(U62(v)) }

func (U62) nbits() int  { return 62 }
func (u U62) mask() U62 { return 0x3fffffffffffffff }
func (u U62) cast() U62 { return u & u.mask() }

// Set assigns v reduced to 62 bits to u.
func (u *U62) Set(v U62) U62 { *u = new(U62(v)).cast(); return U62.cast(v) }

// Add returns u+o reduced to 62 bits.
func (u U62) Add(o U62) U62 { return U62.cast(u + o) }

// Sub returns u-o reduced to 62 bits.
func (u U62) Sub(o U62) U62 { return U62.cast(u - o) }

// Inc returns u+1 reduced to 62 bits.
func (u U62) Inc() U62 { return U62.cast(u + 1) }

// Dec returns u-1 reduced to 62 bits.
func (u U62) Dec() U62 { return U62.cast(u - 1) }

// Mul returns u*o reduced to 62 bits.
func (u U62) Mul(o U62) U62 { return U62.cast(u * o) }

// Div returns u divided by o.
func (u U62) Div(o U62) U62 { return U62.cast(u / o) }

// Mod returns u modulo o.
func (u U62) Mod(o U62) U62 { return U62.cast(u % o) }

// And returns u&o reduced to 62 bits.
func (u U62) And(o U62) U62 { return U62.cast(u & o) }

// Or returns u|o reduced to 62 bits.
func (u U62) Or(o U62) U62 { return U62.cast(u | o) }

// Xor returns u^o reduced to 62 bits.
func (u U62) Xor(o U62) U62 { return U62.cast(u ^ o) }

// Shr returns u>>o reduced to 62 bits.
func (u U62) Shr(o U62) U62 { return U62.cast(u >> o) }

// Shl returns u<<o reduced to 62 bits.
func (u U62) Shl(o U62) U62 { return U62.cast(u << o) }

// Not returns bitwise complement of u reduced to 62 bits.
func (u U62) Not() U62 { return U62.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U62) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U62) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U62 value.
func (u U62) Signed() I62 { return I62.cast(I62(u)) }

// Clamp returns value saturated into the representable range of U62.
func (u U62) Clamp(value uint64) U62 { return uclamp[U62](value) }

// Clip masks u to a bits-wide low field.
func (u U62) Clip(bits int) U62 {
	b := 1 << (bits - 1)
	m := U62(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U62) Bit(index int) U62 {
	if index < 0 {
		index = 62 + index
	}
	return (u >> U62(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U62) SetBit(index int, v U62) {
	if index < 0 {
		index = 62 + index
	}
	bit := U62(1) << U62(index)
	*u = U62.cast((*u &^ bit) | ((v & 1) << U62(index)))
}

// Bitref returns a Range for bit index.
func (u *U62) Bitref(index int) Range[U62] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U62) Bits(lo, hi int) U62 {
	p := 62
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U62((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U62(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U62) SetBits(lo, hi int, v U62) {
	p := 62
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U62((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U62.cast((*u &^ mask) | ((v << U62(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U62) Bitsref(lo, hi int) Range[U62] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U62) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U62) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U62(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U62) Byteref(idx int) Range[U62] { return byte(u, idx) }

// I62 is an 62-bit signed integer in two's complement.
type I62 int64

// AsI62 returns a I62 representing v reduced to 62 bits.
func AsI62[T Signed](v T) I62 { return I62.cast(I62(v)) }

func (I62) nbits() int  { return 62 }
func (i I62) mask() I62 { return 0x3fffffffffffffff }
func (i I62) sign() I62 { return 1 << (i.nbits() - 1) }
func (i I62) cast() I62 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 62 bits to u.
func (i *I62) Set(v I62) I62 { *i = new(I62(v)).cast(); return I62.cast(v) }

// Add returns i+o reduced to 62 bits.
func (i I62) Add(o I62) I62 { return I62.cast(i + o) }

// Sub returns i-o reduced to 62 bits.
func (i I62) Sub(o I62) I62 { return I62.cast(i - o) }

// Inc returns i+1 reduced to 62 bits.
func (i I62) Inc() I62 { return I62.cast(i + 1) }

// Dec returns i-1 reduced to 62 bits.
func (i I62) Dec() I62 { return I62.cast(i - 1) }

// Mul returns i*o reduced to 62 bits.
func (i I62) Mul(o I62) I62 { return I62.cast(i * o) }

// Div returns i divided by o.
func (i I62) Div(o I62) I62 { return I62.cast(i / o) }

// Mod returns i modulo o.
func (i I62) Mod(o I62) I62 { return I62.cast(i % o) }

// And returns i&o reduced to 62 bits.
func (i I62) And(o I62) I62 { return I62.cast(i & o) }

// Or returns i|o reduced to 62 bits.
func (i I62) Or(o I62) I62 { return I62.cast(i | o) }

// Xor returns i^o reduced to 62 bits.
func (i I62) Xor(o I62) I62 { return I62.cast(i ^ o) }

// Shr returns i>>o reduced to 62 bits.
func (i I62) Shr(o I62) I62 { return I62.cast(i >> o) }

// Shl returns i<<o reduced to 62 bits.
func (i I62) Shl(o I62) I62 { return I62.cast(i << o) }

// Not returns bitwise complement of i reduced to 62 bits.
func (i I62) Not() I62 { return I62.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I62) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I62) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U62 bit pattern.
func (i I62) Unsigned() U62 { return U62.cast(U62(i)) }

// Clamp returns value saturated into the representable range of I62.
func (i I62) Clamp(value int64) I62 { return iclamp[I62](value) }

// Clip masks i to a bits-wide low field and sign-extends to 62 bits.
func (i I62) Clip(bits int) I62 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I62(uint64(i)&m^b) - I62(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I62) Bit(index int) I62 {
	if index < 0 {
		index = 62 + index
	}
	return (i >> I62(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I62) SetBit(index int, v I62) {
	if index < 0 {
		index = 62 + index
	}
	bit := I62(1) << I62(index)
	*i = I62.cast((*i &^ bit) | ((v & 1) << I62(index)))
}

// Bitref returns a Range for bit index.
func (i *I62) Bitref(index int) Range[I62] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I62) Bits(lo, hi int) I62 {
	p := 62
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I62((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I62(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I62) SetBits(lo, hi int, v I62) {
	p := 62
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I62((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I62.cast((*i &^ mask) | ((v << I62(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I62) Bitsref(lo, hi int) Range[I62] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I62) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I62) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I62(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I62) Byteref(idx int) Range[I62] { return byte(i, idx) }

// U63 is an 63-bit unsigned integer.
type U63 uint64

// AsU63 returns a U63 representing v reduced to 63 bits.
func AsU63[T Unsigned](v T) U63 { return U63.cast(U63(v)) }

func (U63) nbits() int  { return 63 }
func (u U63) mask() U63 { return 0x7fffffffffffffff }
func (u U63) cast() U63 { return u & u.mask() }

// Set assigns v reduced to 63 bits to u.
func (u *U63) Set(v U63) U63 { *u = new(U63(v)).cast(); return U63.cast(v) }

// Add returns u+o reduced to 63 bits.
func (u U63) Add(o U63) U63 { return U63.cast(u + o) }

// Sub returns u-o reduced to 63 bits.
func (u U63) Sub(o U63) U63 { return U63.cast(u - o) }

// Inc returns u+1 reduced to 63 bits.
func (u U63) Inc() U63 { return U63.cast(u + 1) }

// Dec returns u-1 reduced to 63 bits.
func (u U63) Dec() U63 { return U63.cast(u - 1) }

// Mul returns u*o reduced to 63 bits.
func (u U63) Mul(o U63) U63 { return U63.cast(u * o) }

// Div returns u divided by o.
func (u U63) Div(o U63) U63 { return U63.cast(u / o) }

// Mod returns u modulo o.
func (u U63) Mod(o U63) U63 { return U63.cast(u % o) }

// And returns u&o reduced to 63 bits.
func (u U63) And(o U63) U63 { return U63.cast(u & o) }

// Or returns u|o reduced to 63 bits.
func (u U63) Or(o U63) U63 { return U63.cast(u | o) }

// Xor returns u^o reduced to 63 bits.
func (u U63) Xor(o U63) U63 { return U63.cast(u ^ o) }

// Shr returns u>>o reduced to 63 bits.
func (u U63) Shr(o U63) U63 { return U63.cast(u >> o) }

// Shl returns u<<o reduced to 63 bits.
func (u U63) Shl(o U63) U63 { return U63.cast(u << o) }

// Not returns bitwise complement of u reduced to 63 bits.
func (u U63) Not() U63 { return U63.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U63) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U63) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U63 value.
func (u U63) Signed() I63 { return I63.cast(I63(u)) }

// Clamp returns value saturated into the representable range of U63.
func (u U63) Clamp(value uint64) U63 { return uclamp[U63](value) }

// Clip masks u to a bits-wide low field.
func (u U63) Clip(bits int) U63 {
	b := 1 << (bits - 1)
	m := U63(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U63) Bit(index int) U63 {
	if index < 0 {
		index = 63 + index
	}
	return (u >> U63(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U63) SetBit(index int, v U63) {
	if index < 0 {
		index = 63 + index
	}
	bit := U63(1) << U63(index)
	*u = U63.cast((*u &^ bit) | ((v & 1) << U63(index)))
}

// Bitref returns a Range for bit index.
func (u *U63) Bitref(index int) Range[U63] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U63) Bits(lo, hi int) U63 {
	p := 63
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U63((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U63(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U63) SetBits(lo, hi int, v U63) {
	p := 63
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U63((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U63.cast((*u &^ mask) | ((v << U63(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U63) Bitsref(lo, hi int) Range[U63] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U63) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U63) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U63(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U63) Byteref(idx int) Range[U63] { return byte(u, idx) }

// I63 is an 63-bit signed integer in two's complement.
type I63 int64

// AsI63 returns a I63 representing v reduced to 63 bits.
func AsI63[T Signed](v T) I63 { return I63.cast(I63(v)) }

func (I63) nbits() int  { return 63 }
func (i I63) mask() I63 { return 0x7fffffffffffffff }
func (i I63) sign() I63 { return 1 << (i.nbits() - 1) }
func (i I63) cast() I63 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 63 bits to u.
func (i *I63) Set(v I63) I63 { *i = new(I63(v)).cast(); return I63.cast(v) }

// Add returns i+o reduced to 63 bits.
func (i I63) Add(o I63) I63 { return I63.cast(i + o) }

// Sub returns i-o reduced to 63 bits.
func (i I63) Sub(o I63) I63 { return I63.cast(i - o) }

// Inc returns i+1 reduced to 63 bits.
func (i I63) Inc() I63 { return I63.cast(i + 1) }

// Dec returns i-1 reduced to 63 bits.
func (i I63) Dec() I63 { return I63.cast(i - 1) }

// Mul returns i*o reduced to 63 bits.
func (i I63) Mul(o I63) I63 { return I63.cast(i * o) }

// Div returns i divided by o.
func (i I63) Div(o I63) I63 { return I63.cast(i / o) }

// Mod returns i modulo o.
func (i I63) Mod(o I63) I63 { return I63.cast(i % o) }

// And returns i&o reduced to 63 bits.
func (i I63) And(o I63) I63 { return I63.cast(i & o) }

// Or returns i|o reduced to 63 bits.
func (i I63) Or(o I63) I63 { return I63.cast(i | o) }

// Xor returns i^o reduced to 63 bits.
func (i I63) Xor(o I63) I63 { return I63.cast(i ^ o) }

// Shr returns i>>o reduced to 63 bits.
func (i I63) Shr(o I63) I63 { return I63.cast(i >> o) }

// Shl returns i<<o reduced to 63 bits.
func (i I63) Shl(o I63) I63 { return I63.cast(i << o) }

// Not returns bitwise complement of i reduced to 63 bits.
func (i I63) Not() I63 { return I63.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I63) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I63) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U63 bit pattern.
func (i I63) Unsigned() U63 { return U63.cast(U63(i)) }

// Clamp returns value saturated into the representable range of I63.
func (i I63) Clamp(value int64) I63 { return iclamp[I63](value) }

// Clip masks i to a bits-wide low field and sign-extends to 63 bits.
func (i I63) Clip(bits int) I63 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I63(uint64(i)&m^b) - I63(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I63) Bit(index int) I63 {
	if index < 0 {
		index = 63 + index
	}
	return (i >> I63(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I63) SetBit(index int, v I63) {
	if index < 0 {
		index = 63 + index
	}
	bit := I63(1) << I63(index)
	*i = I63.cast((*i &^ bit) | ((v & 1) << I63(index)))
}

// Bitref returns a Range for bit index.
func (i *I63) Bitref(index int) Range[I63] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I63) Bits(lo, hi int) I63 {
	p := 63
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I63((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I63(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I63) SetBits(lo, hi int, v I63) {
	p := 63
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I63((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I63.cast((*i &^ mask) | ((v << I63(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I63) Bitsref(lo, hi int) Range[I63] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I63) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I63) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I63(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I63) Byteref(idx int) Range[I63] { return byte(i, idx) }

// U64 is an 64-bit unsigned integer.
type U64 uint64

// AsU64 returns a U64 representing v reduced to 64 bits.
func AsU64[T Unsigned](v T) U64 { return U64.cast(U64(v)) }

func (U64) nbits() int  { return 64 }
func (u U64) mask() U64 { return 0xffffffffffffffff }
func (u U64) cast() U64 { return u & u.mask() }

// Set assigns v reduced to 64 bits to u.
func (u *U64) Set(v U64) U64 { *u = new(U64(v)).cast(); return U64.cast(v) }

// Add returns u+o reduced to 64 bits.
func (u U64) Add(o U64) U64 { return U64.cast(u + o) }

// Sub returns u-o reduced to 64 bits.
func (u U64) Sub(o U64) U64 { return U64.cast(u - o) }

// Inc returns u+1 reduced to 64 bits.
func (u U64) Inc() U64 { return U64.cast(u + 1) }

// Dec returns u-1 reduced to 64 bits.
func (u U64) Dec() U64 { return U64.cast(u - 1) }

// Mul returns u*o reduced to 64 bits.
func (u U64) Mul(o U64) U64 { return U64.cast(u * o) }

// Div returns u divided by o.
func (u U64) Div(o U64) U64 { return U64.cast(u / o) }

// Mod returns u modulo o.
func (u U64) Mod(o U64) U64 { return U64.cast(u % o) }

// And returns u&o reduced to 64 bits.
func (u U64) And(o U64) U64 { return U64.cast(u & o) }

// Or returns u|o reduced to 64 bits.
func (u U64) Or(o U64) U64 { return U64.cast(u | o) }

// Xor returns u^o reduced to 64 bits.
func (u U64) Xor(o U64) U64 { return U64.cast(u ^ o) }

// Shr returns u>>o reduced to 64 bits.
func (u U64) Shr(o U64) U64 { return U64.cast(u >> o) }

// Shl returns u<<o reduced to 64 bits.
func (u U64) Shl(o U64) U64 { return U64.cast(u << o) }

// Not returns bitwise complement of u reduced to 64 bits.
func (u U64) Not() U64 { return U64.cast(^u) }

// String returns the canonical decimal representation of u.
func (u U64) String() string { return strconv.FormatUint(uint64(u.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting.
func (u U64) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))
}

// Signed returns u interpreted as a signed U64 value.
func (u U64) Signed() I64 { return I64.cast(I64(u)) }

// Clamp returns value saturated into the representable range of U64.
func (u U64) Clamp(value uint64) U64 { return uclamp[U64](value) }

// Clip masks u to a bits-wide low field.
func (u U64) Clip(bits int) U64 {
	b := 1 << (bits - 1)
	m := U64(b*2 - 1)
	return u & m
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (u U64) Bit(index int) U64 {
	if index < 0 {
		index = 64 + index
	}
	return (u >> U64(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (u *U64) SetBit(index int, v U64) {
	if index < 0 {
		index = 64 + index
	}
	bit := U64(1) << U64(index)
	*u = U64.cast((*u &^ bit) | ((v & 1) << U64(index)))
}

// Bitref returns a Range for bit index.
func (u *U64) Bitref(index int) Range[U64] { return bit(u, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (u U64) Bits(lo, hi int) U64 {
	p := 64
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U64((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (u & mask) >> U64(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (u *U64) SetBits(lo, hi int, v U64) {
	p := 64
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := U64((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*u = U64.cast((*u &^ mask) | ((v << U64(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (u *U64) Bitsref(lo, hi int) Range[U64] { return bits(u, lo, hi) }

// Byte returns the byte at index idx.
func (u U64) Byte(index int) U8 { return U8(u.Bits(index*8, index*8+7)) }

// SetByte sets byte index (0 = least significant) to v.
func (u *U64) SetByte(index int, v U8) { u.SetBits(index*8, index*8+7, U64(v)) }

// Byteref returns a Range for the byte at index idx.
func (u *U64) Byteref(idx int) Range[U64] { return byte(u, idx) }

// I64 is an 64-bit signed integer in two's complement.
type I64 int64

// AsI64 returns a I64 representing v reduced to 64 bits.
func AsI64[T Signed](v T) I64 { return I64.cast(I64(v)) }

func (I64) nbits() int  { return 64 }
func (i I64) mask() I64 { return -1 }
func (i I64) sign() I64 { return 1 << (i.nbits() - 1) }
func (i I64) cast() I64 { return (i&i.mask() ^ i.sign()) - i.sign() }

// Set assigns v reduced to 64 bits to u.
func (i *I64) Set(v I64) I64 { *i = new(I64(v)).cast(); return I64.cast(v) }

// Add returns i+o reduced to 64 bits.
func (i I64) Add(o I64) I64 { return I64.cast(i + o) }

// Sub returns i-o reduced to 64 bits.
func (i I64) Sub(o I64) I64 { return I64.cast(i - o) }

// Inc returns i+1 reduced to 64 bits.
func (i I64) Inc() I64 { return I64.cast(i + 1) }

// Dec returns i-1 reduced to 64 bits.
func (i I64) Dec() I64 { return I64.cast(i - 1) }

// Mul returns i*o reduced to 64 bits.
func (i I64) Mul(o I64) I64 { return I64.cast(i * o) }

// Div returns i divided by o.
func (i I64) Div(o I64) I64 { return I64.cast(i / o) }

// Mod returns i modulo o.
func (i I64) Mod(o I64) I64 { return I64.cast(i % o) }

// And returns i&o reduced to 64 bits.
func (i I64) And(o I64) I64 { return I64.cast(i & o) }

// Or returns i|o reduced to 64 bits.
func (i I64) Or(o I64) I64 { return I64.cast(i | o) }

// Xor returns i^o reduced to 64 bits.
func (i I64) Xor(o I64) I64 { return I64.cast(i ^ o) }

// Shr returns i>>o reduced to 64 bits.
func (i I64) Shr(o I64) I64 { return I64.cast(i >> o) }

// Shl returns i<<o reduced to 64 bits.
func (i I64) Shl(o I64) I64 { return I64.cast(i << o) }

// Not returns bitwise complement of i reduced to 64 bits.
func (i I64) Not() I64 { return I64.cast(^i) }

// String returns the canonical decimal representation of i.
func (i I64) String() string { return strconv.FormatInt(int64(i.cast()), 10) }

// Format implements fmt.Formatter using the canonical value as int64,
// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.
func (i I64) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))
}

// Unsigned returns i as an unsigned U64 bit pattern.
func (i I64) Unsigned() U64 { return U64.cast(U64(i)) }

// Clamp returns value saturated into the representable range of I64.
func (i I64) Clamp(value int64) I64 { return iclamp[I64](value) }

// Clip masks i to a bits-wide low field and sign-extends to 64 bits.
func (i I64) Clip(bits int) I64 {
	b := uint64(1) << (bits - 1)
	m := b*2 - 1
	return (I64(uint64(i)&m^b) - I64(b))
}

// Bit returns bit index as 0 or 1; negative indices count from the MSB.
func (i I64) Bit(index int) I64 {
	if index < 0 {
		index = 64 + index
	}
	return (i >> I64(index)) & 1
}

// SetBit sets bit index to v&1 in place; negative indices count from the MSB.
func (i *I64) SetBit(index int, v I64) {
	if index < 0 {
		index = 64 + index
	}
	bit := I64(1) << I64(index)
	*i = I64.cast((*i &^ bit) | ((v & 1) << I64(index)))
}

// Bitref returns a Range for bit index.
func (i *I64) Bitref(index int) Range[I64] { return bit(i, index) }

// Bits returns the inclusive bit field from lo through hi,
// negative indices count from the MSB, and if lo>hi the range is swapped.
func (i I64) Bits(lo, hi int) I64 {
	p := 64
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I64((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	return (i & mask) >> I64(lo)
}

// SetBits sets bits lo through hi to v in place; index rules match Bits.
func (i *I64) SetBits(lo, hi int, v I64) {
	p := 64
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := I64((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)
	*i = I64.cast((*i &^ mask) | ((v << I64(lo)) & mask))
}

// Bitsref returns a Range for bits lo through hi.
func (i *I64) Bitsref(lo, hi int) Range[I64] { return bits(i, lo, hi) }

// Byte returns the byte at index idx.
func (i I64) Byte(idx int) U8 { return U8(i.Bits(idx*8, idx*8+7)) }

// SetByte sets the byte at index idx to v.
func (i *I64) SetByte(idx int, v U8) { i.SetBits(idx*8, idx*8+7, I64(v)) }

// Byteref returns a Range for the byte at index idx.
func (i *I64) Byteref(idx int) Range[I64] { return byte(i, idx) }
