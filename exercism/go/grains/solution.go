package grains

import "errors"

const testVersion = 1

var ErrBadInput = errors.New("bad input data")

func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, ErrBadInput
	}
	return 1 << uint(i-1), nil
}

func Total() uint64 {
	return ^uint64(0) // that's 2^64 - 1, but doing this expicility would overflow uint64
}
