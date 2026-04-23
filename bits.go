// Package bit provides fixed-width integers u1–u64 and i1–i64 with
// width-correct arithmetic, bit fields, byte access, clamping, and ranges.
package bit

// precision is implemented by every fixed-width type for internal use.
type precision interface {
	nbits() int
}

// Un is an unsigned fixed-width integer.
type Un interface {
	Unsigned
	precision
}

// In is a signed fixed-width integer.
type In interface {
	Signed
	precision
}

// Integer is either Signed or Unsigned.
type Integer interface{ Signed | Unsigned }

// Unsigned is the set of unsigned storage types.
type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Signed is the set of signed storage types.
type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64
}

// Range is a read/write view of a contiguous bit field inside an integer.
type Range[T interface {
	Integer
	precision
}] struct {
	target    *T
	mask      T
	precision int
	shift     uint32
}

// Val returns the field value, right-aligned (same as Bit for multi-bit fields).
func (r Range[T]) Val() T {
	return (*r.target & r.mask) >> T(r.shift)
}

// Bit returns the field value (same as Val).
func (r Range[T]) Bit() T {
	return (*r.target & r.mask) >> T(r.shift)
}

// Assign copies the bit field from o into r's position and width.
func (r Range[T]) Assign(o *Range[T]) Range[T] {
	*r.target = (*r.target &^ r.mask) | ((((*o.target & o.mask) >> T(o.shift)) << r.shift) & r.mask)
	return r
}

// Set writes value o into the field (masked and shifted).
func (r Range[T]) Set(o T) Range[T] {
	*r.target = (*r.target &^ r.mask) | ((o << r.shift) & r.mask)
	return r
}

// Inc increments the field by one within the field mask.
func (r Range[T]) Inc() Range[T] {
	*r.target = (*r.target &^ r.mask) | ((*r.target + (1 << r.shift)) & r.mask)
	return r
}

// Dec decrements the field by one within the field mask.
func (r Range[T]) Dec() Range[T] {
	*r.target = (*r.target &^ r.mask) | ((*r.target - (1 << r.shift)) & r.mask)
	return r
}

// Add adds o to the field value within the field mask.
func (r Range[T]) Add(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) + o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Sub subtracts o from the field value within the field mask.
func (r Range[T]) Sub(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) - o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Mul multiplies the field value by o within the field mask.
func (r Range[T]) Mul(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) * o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Div divides the field value by o within the field mask.
func (r Range[T]) Div(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) / o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Mod replaces the field with (field % o) within the field mask.
func (r Range[T]) Mod(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) % o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Shl left-shifts the field value by o within the field mask.
func (r Range[T]) Shl(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) << o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// Shr right-shifts the field value by o within the field mask.
func (r Range[T]) Shr(o T) Range[T] {
	value := ((*r.target & r.mask) >> T(r.shift)) >> o
	*r.target = (*r.target &^ r.mask) | ((value << r.shift) & r.mask)
	return r
}

// And bitwise-ANDs the field with o (other bits unchanged outside the mask).
func (r Range[T]) And(o T) Range[T] {
	*r.target = *r.target & (^r.mask | ((o << r.shift) & r.mask))
	return r
}

// Xor bitwise-XORs the field with o.
func (r Range[T]) Xor(o T) Range[T] {
	*r.target = *r.target ^ ((o << r.shift) & r.mask)
	return r
}

// Or bitwise-ORs the field with o.
func (r Range[T]) Or(o T) Range[T] {
	*r.target = *r.target | ((o << r.shift) & r.mask)
	return r
}

// bits returns a Range over the inclusive bit indices lo..hi of *t.
// Negative indices count from the MSB; if lo > hi, lo and hi are swapped.
func bits[T interface {
	Integer
	precision
}](t *T, lo, hi int) Range[T] {
	p := (*t).nbits()
	if lo < 0 {
		lo = p + lo
	}
	if hi < 0 {
		hi = p + hi
	}
	if lo > hi {
		lo, hi = hi, lo
	}
	mask := T(int64((uint64(0xffffffffffffffff) >> (64 - (int64(hi) - int64(lo) + 1))) << lo))
	return Range[T]{
		target:    t,
		mask:      mask,
		precision: p,
		shift:     uint32(lo),
	}
}

// bit returns a Range over a single bit of *t (negative index from MSB).
func bit[T interface {
	Integer
	precision
}](t *T, index int) Range[T] {
	p := (*t).nbits()
	if index < 0 {
		index = p + index
	}
	return Range[T]{
		target:    t,
		mask:      T(1 << index),
		precision: p,
		shift:     uint32(index),
	}
}

// byte returns a Range over byte index of *t (0 = least significant byte).
func byte[T interface {
	Integer
	precision
}](t *T, index int) Range[T] {
	return bits(t, index*8+0, index*8+7)
}

func uclamp[U Un](value uint64) U {
	var u U
	b := 1 << (u.nbits() - 1)
	m := b*2 - 1
	if U(value) < U(m) {
		return U(value)
	}
	return U(m)
}

func iclamp[I In](value int64) I {
	var zero I
	b := int64(1) << (zero.nbits() - 1)
	m := b - 1
	if value > m {
		return I(m)
	}
	if value < -b {
		return I(-b)
	}
	return I(value)
}
