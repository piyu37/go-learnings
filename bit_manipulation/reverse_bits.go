package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res = res << 1
		bit := num & 1
		res += bit
		num = num >> 1
	}

	return res
}

// Let's break down how the bitmasks `0xff00ff00` and `0x00ff00ff` are used to isolate 8-bit blocks in the bit-reversal process.

// ### Understanding the Bitmasks

// - `0xff00ff00` in binary is:
//   ```
//   11111111 00000000 11111111 00000000
//   ```

// - `0x00ff00ff` in binary is:
//   ```
//   00000000 11111111 00000000 11111111
//   ```

// These bitmasks are designed to isolate specific 8-bit segments of a 32-bit integer.

// ### How They Work

// Given a 32-bit number `num`, let's apply these masks and observe the effect.

// #### Isolating with `0xff00ff00`
// When you apply `num & 0xff00ff00`, you keep the 8-bit blocks that align with the `1`s in the mask and set the
// rest to `0`. Specifically, this isolates the second and fourth 8-bit blocks:

// - Any bit in `num` corresponding to a `1` in the mask is retained.
// - Any bit in `num` corresponding to a `0` in the mask is set to `0`.

// For example, consider a 32-bit number `num`:
// ```
// num = 10101010 11001100 10101010 11001100
// ```

// Applying the mask `0xff00ff00`:
// ```
// num & 0xff00ff00 =
// 10101010 11001100 10101010 11001100 &
// 11111111 00000000 11111111 00000000 =
// 10101010 00000000 10101010 00000000
// ```

// #### Isolating with `0x00ff00ff`
// When you apply `num & 0x00ff00ff`, you keep the 8-bit blocks that align with the `1`s in the mask and set
// the rest to `0`. This isolates the first and third 8-bit blocks:

// - Any bit in `num` corresponding to a `1` in the mask is retained.
// - Any bit in `num` corresponding to a `0` in the mask is set to `0`.

// Using the same example number `num`:
// ```
// num = 10101010 11001100 10101010 11001100
// ```

// Applying the mask `0x00ff00ff`:
// ```
// num & 0x00ff00ff =
// 10101010 11001100 10101010 11001100 &
// 00000000 11111111 00000000 11111111 =
// 00000000 11001100 00000000 11001100
// ```

// ### Shifting and Combining

// Once the 8-bit blocks are isolated using the masks, the next step is to shift these isolated blocks to their
// new positions and combine them using the bitwise OR operator `|`.

// - `(num & 0xff00ff00) >> 8` shifts the second and fourth 8-bit blocks right by 8 bits.
// - `(num & 0x00ff00ff) << 8` shifts the first and third 8-bit blocks left by 8 bits.

// Then, combining these two results:
// ```
// num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
// ```
// This effectively swaps the 8-bit blocks within the 16-bit segments. For instance:
// ```
// ((10101010 00000000 10101010 00000000) >> 8) =
// 00000000 10101010 00000000 10101010

// ((00000000 11001100 00000000 11001100) << 8) =
// 11001100 00000000 11001100 00000000

// Combining these results with OR:
// 00000000 10101010 00000000 10101010 |
// 11001100 00000000 11001100 00000000 =
// 11001100 10101010 11001100 10101010
// ```

// Thus, this operation swaps the corresponding 8-bit blocks within the 32-bit integer. This process is part of a
// series of similar steps to fully reverse the bits in the integer.
func reverseBits2(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// https://leetcode.com/problems/reverse-bits/description/?envType=study-plan-v2&envId=top-interview-150
func reverseBitsMain() {
	var n uint32 = 0b00000010100101000001111010011100
	fmt.Println(reverseBits(n))
	fmt.Println(reverseBits2(n))
}
