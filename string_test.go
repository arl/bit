package bit

import (
	"fmt"
	"testing"
)

var (
	_ fmt.Stringer  = U1(0)
	_ fmt.Stringer  = U3(0)
	_ fmt.Stringer  = U8(0)
	_ fmt.Stringer  = U16(0)
	_ fmt.Stringer  = U32(0)
	_ fmt.Stringer  = U64(0)
	_ fmt.Stringer  = I1(0)
	_ fmt.Stringer  = I3(0)
	_ fmt.Stringer  = I8(0)
	_ fmt.Stringer  = I16(0)
	_ fmt.Stringer  = I32(0)
	_ fmt.Stringer  = I64(0)
	_ fmt.Formatter = U1(0)
	_ fmt.Formatter = U3(0)
	_ fmt.Formatter = U8(0)
	_ fmt.Formatter = U16(0)
	_ fmt.Formatter = U32(0)
	_ fmt.Formatter = U64(0)
	_ fmt.Formatter = I1(0)
	_ fmt.Formatter = I3(0)
	_ fmt.Formatter = I8(0)
	_ fmt.Formatter = I16(0)
	_ fmt.Formatter = I32(0)
	_ fmt.Formatter = I64(0)
)

func TestString_UnsignedDecimal(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"u1(0)", U1(0).String(), "0"},
		{"u1(1)", U1(1).String(), "1"},
		{"u3(0)", U3(0).String(), "0"},
		{"u3(5)", U3(5).String(), "5"},
		{"u3(7)", U3(7).String(), "7"},
		{"u8(0)", U8(0).String(), "0"},
		{"u8(255)", U8(255).String(), "255"},
		{"u16(0xABCD)", U16(0xABCD).String(), "43981"},
		{"u16 max", U16(0xFFFF).String(), "65535"},
		{"u32 max", U32(0xFFFFFFFF).String(), "4294967295"},
		{"u64 max", U64(0xFFFFFFFFFFFFFFFF).String(), "18446744073709551615"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestString_SignedDecimal(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"I1(0)", I1(0).String(), "0"},
		{"I1(-1)", I1(-1).String(), "-1"},
		{"I3(0)", I3(0).String(), "0"},
		{"I3(3) max", I3(3).String(), "3"},
		{"I3(-4) min", I3(-4).String(), "-4"},
		{"I3(-1)", I3(-1).String(), "-1"},
		{"I8(127) max", I8(127).String(), "127"},
		{"I8(-128) min", I8(-128).String(), "-128"},
		{"I16 max", I16(0x7FFF).String(), "32767"},
		{"I16 min", I16(-0x8000).String(), "-32768"},
		{"I32 max", I32(0x7FFFFFFF).String(), "2147483647"},
		{"I32 min", I32(-0x80000000).String(), "-2147483648"},
		{"I64 min", I64(-0x8000000000000000).String(), "-9223372036854775808"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestString_Canonicalizes(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{

		{"u3(15).String()", U3(15).String(), "7"},
		{"u3(9).String()", U3(9).String(), "1"},
		{"I3(7).String()", I3(7).String(), "-1"},
		{"I3(5).String()", I3(5).String(), "-3"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestFormat_NumericVerbs(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"%b u8", fmt.Sprintf("%b", U8(0xAB)), "10101011"},
		{"%o u8", fmt.Sprintf("%o", U8(0xAB)), "253"},
		{"%d u8", fmt.Sprintf("%d", U8(0xAB)), "171"},
		{"%x u8", fmt.Sprintf("%x", U8(0xAB)), "ab"},
		{"%X u8", fmt.Sprintf("%X", U8(0xAB)), "AB"},
		{"%v u8", fmt.Sprintf("%v", U8(0xAB)), "171"},
		{"%x u16", fmt.Sprintf("%x", U16(0xABFD)), "abfd"},
		{"%X u16", fmt.Sprintf("%X", U16(0xABFD)), "ABFD"},
		{"%x u32", fmt.Sprintf("%x", U32(0xDEADBEEF)), "deadbeef"},
		{"%d I8(-1)", fmt.Sprintf("%d", I8(-1)), "-1"},
		{"%b I8(-1)", fmt.Sprintf("%b", I8(-1)), "-1"},
		{"%x I8(-1)", fmt.Sprintf("%x", I8(-1)), "-1"},
		{"%d I16(-128)", fmt.Sprintf("%d", I16(-128)), "-128"},
		{"%x I16(0x7F)", fmt.Sprintf("%x", I16(0x7F)), "7f"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestFormat_WidthAndFlags(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"%08x u16", fmt.Sprintf("%08x", U16(0xABCD)), "0000abcd"},
		{"%08X u16", fmt.Sprintf("%08X", U16(0xABCD)), "0000ABCD"},
		{"%5d u8", fmt.Sprintf("%5d", U8(42)), "   42"},
		{"%-5d| u8", fmt.Sprintf("%-5d|", U8(42)), "42   |"},
		{"%+d I8 pos", fmt.Sprintf("%+d", I8(42)), "+42"},
		{"%+d I8 neg", fmt.Sprintf("%+d", I8(-42)), "-42"},
		{"%#x u8", fmt.Sprintf("%#x", U8(0xAB)), "0xab"},
		{"%#o u8", fmt.Sprintf("%#o", U8(0x10)), "020"},
		{"% d I8", fmt.Sprintf("% d", I8(5)), " 5"},
		{"%010d I32 neg", fmt.Sprintf("%010d", I32(-42)), "-000000042"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestFormat_Canonicalizes(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"%d -> U3(15)", fmt.Sprintf("%d", U3(15)), "7"},
		{"%b -> U3(15)", fmt.Sprintf("%b", U3(15)), "111"},
		{"%x -> U4(0xFF)", fmt.Sprintf("%x", U4(0xFF)), "f"},
		{"%d -> I3(7)", fmt.Sprintf("%d", I3(7)), "-1"},
		{"%d -> I5(31)", fmt.Sprintf("%d", I5(31)), "-1"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}

func TestFormat_MatchesBackingPrimitive(t *testing.T) {
	verbs := []string{"%d", "%b", "%o", "%x", "%X", "%v", "%+d", "%08x", "%#x"}
	checkU := func(name string, typed any, primitive any) {
		for _, v := range verbs {
			got := fmt.Sprintf(v, typed)
			want := fmt.Sprintf(v, primitive)
			if got != want {
				t.Errorf("%s %q: got %q, want %q", name, v, got, want)
			}
		}
	}
	checkU("u8(0xAB)", U8(0xAB), uint8(0xAB))
	checkU("u16(0xABCD)", U16(0xABCD), uint16(0xABCD))
	checkU("u32(0xDEADBEEF)", U32(0xDEADBEEF), uint32(0xDEADBEEF))
	checkU("u64(max)", U64(0xFFFFFFFFFFFFFFFF), uint64(0xFFFFFFFFFFFFFFFF))
	checkU("I8(-42)", I8(-42), int8(-42))
	checkU("I16(-1)", I16(-1), int16(-1))
	checkU("I32(min)", I32(-0x80000000), int32(-0x80000000))
	checkU("I64(max)", I64(0x7FFFFFFFFFFFFFFF), int64(0x7FFFFFFFFFFFFFFF))
}

func TestFormat_PrintAndSprint(t *testing.T) {
	v := U16(0xABCD)
	if got, want := fmt.Sprint(v), v.String(); got != want {
		t.Errorf("fmt.Sprint(u16) = %q, want %q", got, want)
	}
	if got, want := fmt.Sprintln(v), v.String()+"\n"; got != want {
		t.Errorf("fmt.Sprintln(u16) = %q, want %q", got, want)
	}
	if got, want := fmt.Sprint(I32(-1234)), "-1234"; got != want {
		t.Errorf("fmt.Sprint(I32) = %q, want %q", got, want)
	}
}

func TestFormat_NoStringHexEncoding(t *testing.T) {
	if got := fmt.Sprintf("%x", U16(0xABFD)); got == "3434303239" {
		t.Fatalf("%%x on u16 hex-encoded decimal string: %q", got)
	}
	if got := fmt.Sprintf("%x", U16(0xABFD)); got != "abfd" {
		t.Errorf("%%x u16(0xABFD) = %q, want %q", got, "abfd")
	}
}
