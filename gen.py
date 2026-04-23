#!/usr/bin/env python3
"""Generate types.go: u1..u64 and i1..i64 fixed-width integer types."""

import sys


def backing_unsigned(n: int) -> str:
    if n <= 8:
        return "uint8"
    if n <= 16:
        return "uint16"
    if n <= 32:
        return "uint32"
    return "uint64"


def backing_signed(n: int) -> str:
    if n <= 8:
        return "int8"
    if n <= 16:
        return "int16"
    if n <= 32:
        return "int32"
    return "int64"


def hex_mask(n: int) -> str:
    if n == 64:
        return "0xffffffffffffffff"
    return hex((1 << n) - 1)


def signed_mask(n: int) -> str:
    if n in (8, 16, 32, 64):
        return "-1"
    return hex((1 << n) - 1)


def write_unsigned(n: int) -> str:
    back = backing_unsigned(n)
    nm = f"U{n}"
    sig = f"I{n}"
    mask = hex_mask(n)
    lines = []
    a = lines.append

    a(f"")
    a(f"// {nm} is an {n}-bit unsigned integer.")
    a(f"type {nm} {back}")
    a(f"")
    a(f"func ({nm}) nbits() int           {{ return {n} }}")
    a(f"func (u {nm}) mask() {nm}         {{ return {mask} }}")
    a(f"func (u {nm}) cast() {nm}         {{ return u & u.mask() }}")
    a(f"")
    a(f"// Add returns u+o reduced to {n} bits.")
    a(f"func (u {nm}) Add(o {nm}) {nm}    {{ return {nm}.cast(u + o) }}")
    a(f"// Sub returns u-o reduced to {n} bits.")
    a(f"func (u {nm}) Sub(o {nm}) {nm}    {{ return {nm}.cast(u - o) }}")
    a(f"// Inc returns u+1 reduced to {n} bits.")
    a(f"func (u {nm}) Inc() {nm}          {{ return {nm}.cast(u + 1) }}")
    a(f"// Dec returns u-1 reduced to {n} bits.")
    a(f"func (u {nm}) Dec() {nm}          {{ return {nm}.cast(u - 1) }}")
    a(f"// Mul returns u*o reduced to {n} bits.")
    a(f"func (u {nm}) Mul(o {nm}) {nm}    {{ return {nm}.cast(u * o) }}")
    a(f"// Div returns u divided by o.")
    a(f"func (u {nm}) Div(o {nm}) {nm}    {{ return {nm}.cast(u / o) }}")
    a(f"// Mod returns u modulo o.")
    a(f"func (u {nm}) Mod(o {nm}) {nm}    {{ return {nm}.cast(u % o) }}")
    a(f"// And returns u&o reduced to {n} bits.")
    a(f"func (u {nm}) And(o {nm}) {nm}    {{ return {nm}.cast(u & o) }}")
    a(f"// Or returns u|o reduced to {n} bits.")
    a(f"func (u {nm}) Or(o {nm}) {nm}     {{ return {nm}.cast(u | o) }}")
    a(f"// Xor returns u^o reduced to {n} bits.")
    a(f"func (u {nm}) Xor(o {nm}) {nm}    {{ return {nm}.cast(u ^ o) }}")
    a(f"// Shr returns u>>o reduced to {n} bits.")
    a(f"func (u {nm}) Shr(o {nm}) {nm}    {{ return {nm}.cast(u >> o) }}")
    a(f"// Shl returns u<<o reduced to {n} bits.")
    a(f"func (u {nm}) Shl(o {nm}) {nm}    {{ return {nm}.cast(u << o) }}")
    a(f"// Not returns bitwise complement of u reduced to {n} bits.")
    a(f"func (u {nm}) Not() {nm}          {{ return {nm}.cast(^u) }}")
    a(f"")
    a(f"// String returns the canonical decimal representation of u.")
    a(
        f"func (u {nm}) String() string {{ return strconv.FormatUint(uint64(u.cast()), 10) }}"
    )
    a(f"")
    a(
        f"// Format implements fmt.Formatter using the canonical value as uint64, so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in unsigned formatting."
    )
    a(f"func (u {nm}) Format(f fmt.State, c rune) {{")
    a(f"\tfmt.Fprintf(f, fmt.FormatString(f, c), uint64(u.cast()))")
    a(f"}}")
    a(f"")
    a(f"// Signed returns u interpreted as a signed {nm} value.")
    a(f"func (u {nm}) Signed() {sig}     {{ return {sig}.cast({sig}(u)) }}")
    a(f"// Clamp returns value saturated into the representable range of {nm}.")
    a(f"func (u {nm}) Clamp(value uint64) {nm} {{ return uclamp[{nm}](value) }}")
    a(f"// Clip masks u to a bits-wide low field.")
    a(f"func (u {nm}) Clip(bits int) {nm} {{")
    a(f"\tb := 1 << (bits - 1)")
    a(f"\tm := {nm}(b*2 - 1)")
    a(f"\treturn u & m")
    a(f"}}")
    a(f"")
    a(f"// Bit returns bit index as 0 or 1; negative indices count from the MSB.")
    a(f"func (u {nm}) Bit(index int) {nm} {{")
    a(f"\tif index < 0 {{ index = {n} + index }}")
    a(f"\treturn (u >> {nm}(index)) & 1")
    a(f"}}")
    a(f"")
    a(f"// SetBit sets bit index to v&1 in place; negative indices count from the MSB.")
    a(f"func (u *{nm}) SetBit(index int, v {nm}) {{")
    a(f"\tif index < 0 {{ index = {n} + index }}")
    a(f"\tbit := {nm}(1) << {nm}(index)")
    a(f"\t*u = {nm}.cast((*u &^ bit) | ((v & 1) << {nm}(index)))")
    a(f"}}")
    a(f"")
    a(f"// Bitref returns a Range for bit index.")
    a(f"func (u *{nm}) Bitref(index int) Range[{nm}] {{ return bit(u, index) }}")
    a(f"")
    a(f"// Bits returns the inclusive bit field from lo through hi,")
    a(f"// negative indices count from the MSB, and if lo>hi the range is swapped.")
    a(f"func (u {nm}) Bits(lo, hi int) {nm} {{")
    a(f"\tp := {n}")
    a(f"\tif lo < 0 {{ lo = p + lo }}")
    a(f"\tif hi < 0 {{ hi = p + hi }}")
    a(f"\tif lo > hi {{ lo, hi = hi, lo }}")
    a(f"\tmask := {nm}((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)")
    a(f"\treturn (u & mask) >> {nm}(lo)")
    a(f"}}")
    a(f"")
    a(f"// SetBits sets bits lo through hi to v in place; index rules match Bits.")
    a(f"func (u *{nm}) SetBits(lo, hi int, v {nm}) {{")
    a(f"\tp := {n}")
    a(f"\tif lo < 0 {{ lo = p + lo }}")
    a(f"\tif hi < 0 {{ hi = p + hi }}")
    a(f"\tif lo > hi {{ lo, hi = hi, lo }}")
    a(f"\tmask := {nm}((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)")
    a(f"\t*u = {nm}.cast((*u &^ mask) | ((v << {nm}(lo)) & mask))")
    a(f"}}")
    a(f"")
    a(f"// Bitsref returns a Range for bits lo through hi.")
    a(f"func (u *{nm}) Bitsref(lo, hi int) Range[{nm}] {{ return bits(u, lo, hi) }}")
    a(f"")
    a(f"// Byte returns byte index (0 = least significant) as a bit field read.")
    a(f"func (u {nm}) Byte(index int) {nm} {{ return u.Bits(index*8, index*8+7) }}")
    a(f"")
    a(f"// SetByte sets byte index (0 = least significant) to v.")
    a(
        f"func (u *{nm}) SetByte(index int, v {nm}) {{ u.SetBits(index*8, index*8+7, v) }}"
    )
    a(f"")
    a(f"// Byteref returns a Range for byte index.")
    a(f"func (u *{nm}) Byteref(index int) Range[{nm}] {{ return byte(u, index) }}")

    return "\n".join(lines)


def write_signed(n: int) -> str:
    back = backing_signed(n)
    nm = f"I{n}"
    uns = f"U{n}"
    mask = signed_mask(n)
    lines = []
    a = lines.append

    a(f"")
    a(f"// {nm} is an {n}-bit signed integer in two's complement.")
    a(f"type {nm} {back}")
    a(f"")
    a(f"func ({nm}) nbits() int        {{ return {n} }}")
    a(f"func (i {nm}) mask() {nm}      {{ return {mask} }}")
    a(f"func (i {nm}) sign() {nm}      {{ return 1 << (i.nbits() - 1) }}")
    a(f"func (i {nm}) cast() {nm}      {{ return (i&i.mask() ^ i.sign()) - i.sign() }}")
    a(f"")
    a(f"// Add returns i+o reduced to {n} bits.")
    a(f"func (i {nm}) Add(o {nm}) {nm} {{ return {nm}.cast(i + o) }}")
    a(f"// Sub returns i-o reduced to {n} bits.")
    a(f"func (i {nm}) Sub(o {nm}) {nm} {{ return {nm}.cast(i - o) }}")
    a(f"// Inc returns i+1 reduced to {n} bits.")
    a(f"func (i {nm}) Inc() {nm}       {{ return {nm}.cast(i + 1) }}")
    a(f"// Dec returns i-1 reduced to {n} bits.")
    a(f"func (i {nm}) Dec() {nm}       {{ return {nm}.cast(i - 1) }}")
    a(f"// Mul returns i*o reduced to {n} bits.")
    a(f"func (i {nm}) Mul(o {nm}) {nm} {{ return {nm}.cast(i * o) }}")
    a(f"// Div returns i divided by o.")
    a(f"func (i {nm}) Div(o {nm}) {nm} {{ return {nm}.cast(i / o) }}")
    a(f"// Mod returns i modulo o.")
    a(f"func (i {nm}) Mod(o {nm}) {nm} {{ return {nm}.cast(i % o) }}")
    a(f"// And returns i&o reduced to {n} bits.")
    a(f"func (i {nm}) And(o {nm}) {nm} {{ return {nm}.cast(i & o) }}")
    a(f"// Or returns i|o reduced to {n} bits.")
    a(f"func (i {nm}) Or(o {nm}) {nm}  {{ return {nm}.cast(i | o) }}")
    a(f"// Xor returns i^o reduced to {n} bits.")
    a(f"func (i {nm}) Xor(o {nm}) {nm} {{ return {nm}.cast(i ^ o) }}")
    a(f"// Shr returns i>>o reduced to {n} bits.")
    a(f"func (i {nm}) Shr(o {nm}) {nm} {{ return {nm}.cast(i >> o) }}")
    a(f"// Shl returns i<<o reduced to {n} bits.")
    a(f"func (i {nm}) Shl(o {nm}) {nm} {{ return {nm}.cast(i << o) }}")
    a(f"// Not returns bitwise complement of i reduced to {n} bits.")
    a(f"func (i {nm}) Not() {nm}       {{ return {nm}.cast(^i) }}")
    a(f"")
    a(f"// String returns the canonical decimal representation of i.")
    a(
        f"func (i {nm}) String() string {{ return strconv.FormatInt(int64(i.cast()), 10) }}"
    )
    a(f"")
    a(f"// Format implements fmt.Formatter using the canonical value as int64, ")
    a(f"// so %%b, %%o, %%d, %%x, %%X, and %%v match Go's built-in signed formatting.")
    a(f"func (i {nm}) Format(f fmt.State, c rune) {{")
    a(f"\tfmt.Fprintf(f, fmt.FormatString(f, c), int64(i.cast()))")
    a(f"}}")
    a(f"")
    a(f"// Unsigned returns i as an unsigned {uns} bit pattern.")
    a(f"func (i {nm}) Unsigned() {uns}    {{ return {uns}.cast({uns}(i)) }}")
    a(f"// Clamp returns value saturated into the representable range of {nm}.")
    a(f"func (i {nm}) Clamp(value int64) {nm} {{ return iclamp[{nm}](value) }}")
    a(f"// Clip masks i to a bits-wide low field and sign-extends to {n} bits.")
    a(f"func (i {nm}) Clip(bits int) {nm} {{")
    a(f"\tb := uint64(1) << (bits - 1)")
    a(f"\tm := b*2 - 1")
    a(f"\treturn ({nm}(uint64(i) & m ^ b) - {nm}(b))")
    a(f"}}")
    a(f"")
    a(f"// Bit returns bit index as 0 or 1; negative indices count from the MSB.")
    a(f"func (i {nm}) Bit(index int) {nm} {{")
    a(f"\tif index < 0 {{ index = {n} + index }}")
    a(f"\treturn (i >> {nm}(index)) & 1")
    a(f"}}")
    a(f"")
    a(f"// SetBit sets bit index to v&1 in place; negative indices count from the MSB.")
    a(f"func (i *{nm}) SetBit(index int, v {nm}) {{")
    a(f"\tif index < 0 {{ index = {n} + index }}")
    a(f"\tbit := {nm}(1) << {nm}(index)")
    a(f"\t*i = {nm}.cast((*i &^ bit) | ((v & 1) << {nm}(index)))")
    a(f"}}")
    a(f"")
    a(f"// Bitref returns a Range for bit index.")
    a(f"func (i *{nm}) Bitref(index int) Range[{nm}] {{ return bit(i, index) }}")
    a(f"")
    a(f"// Bits returns the inclusive bit field from lo through hi,")
    a(f"// negative indices count from the MSB, and if lo>hi the range is swapped.")
    a(f"func (i {nm}) Bits(lo, hi int) {nm} {{")
    a(f"\tp := {n}")
    a(f"\tif lo < 0 {{ lo = p + lo }}")
    a(f"\tif hi < 0 {{ hi = p + hi }}")
    a(f"\tif lo > hi {{ lo, hi = hi, lo }}")
    a(f"\tmask := {nm}((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)")
    a(f"\treturn (i & mask) >> {nm}(lo)")
    a(f"}}")
    a(f"")
    a(f"// SetBits sets bits lo through hi to v in place; index rules match Bits.")
    a(f"func (i *{nm}) SetBits(lo, hi int, v {nm}) {{")
    a(f"\tp := {n}")
    a(f"\tif lo < 0 {{ lo = p + lo }}")
    a(f"\tif hi < 0 {{ hi = p + hi }}")
    a(f"\tif lo > hi {{ lo, hi = hi, lo }}")
    a(f"\tmask := {nm}((uint64(0xffffffffffffffff) >> (64 - uint(hi-lo+1))) << lo)")
    a(f"\t*i = {nm}.cast((*i &^ mask) | ((v << {nm}(lo)) & mask))")
    a(f"}}")
    a(f"")
    a(f"// Bitsref returns a Range for bits lo through hi.")
    a(f"func (i *{nm}) Bitsref(lo, hi int) Range[{nm}] {{ return bits(i, lo, hi) }}")
    a(f"")
    a(f"// Byte returns byte index (0 = least significant) as a bit field read.")
    a(f"func (i {nm}) Byte(index int) {nm} {{ return i.Bits(index*8, index*8+7) }}")
    a(f"")
    a(f"// SetByte sets byte index (0 = least significant) to v.")
    a(
        f"func (i *{nm}) SetByte(index int, v {nm}) {{ i.SetBits(index*8, index*8+7, v) }}"
    )
    a(f"")
    a(f"// Byteref returns a Range for byte index.")
    a(f"func (i *{nm}) Byteref(index int) Range[{nm}] {{ return byte(i, index) }}")

    return "\n".join(lines)


def main():
    out = [
        "// Generated Code; DO NOT EDIT.",
        "",
        "package bit",
        "",
        "import (",
        '\t"fmt"',
        '\t"strconv"',
        ")",
        "",
    ]
    for n in range(1, 65):
        out.append(write_unsigned(n))
        out.append(write_signed(n))
    text = "\n".join(out) + "\n"
    path = sys.argv[1] if len(sys.argv) > 1 else "types.go"
    with open(path, "w") as f:
        f.write(text)
    print(f"wrote {path} ({len(text):,} bytes, {text.count(chr(10))} lines)")


if __name__ == "__main__":
    main()
