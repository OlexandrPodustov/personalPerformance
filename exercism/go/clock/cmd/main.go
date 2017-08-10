// Package cook provides handy conversion methods for units typically used in recipes.
package main

import (
	"fmt"
)

const testVersion = 4

const mod = 24 * 60

type Clock struct {
	min int
}

func (c *Clock) normalize() {
	t := c.min
	println(t)
	if t < 0 {
		t -= (t / mod) * mod
		println(t)

		t += mod
	}
	c.min = t % mod
}

func New(hour, minute int) Clock {
	res := Clock{min: 60*hour + minute}

	res.normalize()
	println(res.String())
	return res
}

func (c Clock) String() string {
	m := c.min % 60
	h := c.min / 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (c Clock) Add(minutes int) Clock {
	c.min += minutes
	c.normalize()
	return c
}

func main() {
	nnn := New(-100, 15)
	println("result " + nnn.String())
}
