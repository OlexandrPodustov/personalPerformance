package main

import "math/bits"

func findComplement(num int) int {
	if num == 1 {
		return 0
	}

	bitLength := bits.Len(uint(num))

	mask := (1 << bitLength) - 1

	return num ^ mask
}
