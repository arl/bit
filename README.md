# bit - Fixed-width integers for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/arl/bit.svg)](https://pkg.go.dev/github.com/arl/bit)
[![test](https://github.com/arl/bit/actions/workflows/test.yml/badge.svg)](https://github.com/arl/bit/actions/workflows/test.yml)

 - Unsigned **U1** to **U64**
 - Signed **I1** to **I64**
 - Each stored in the smallest standard `uint` or `int` type that fits.
 - Operations wrap or mask to the declared width.
 - **Clamp** saturates a wider value into range
 - **Clip** narrows by bit count
 - **Bit**, **SetBit** and **Bitref**
 - **Bits**, **SetBits** and **Bitsref**
 - **Byte**, **SetByte** and **Byteref**
 - Implements **Stringer**
 - and **fmt.Formatter** to format the canonical numeric value like a built-in integer.


Licensed under the MIT License; see [LICENSE](LICENSE).

API inspired by Ares emulator nall/primitives
