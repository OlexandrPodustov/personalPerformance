// Package clock provides very convenient conversion for minutes and hours.
package clock

import (
	"strconv"
)

const testVersion = 4

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	n := hour*60 + minute
	hh, mm := min2hr(n)
	return Clock{h: hh, m: mm}
}

func min2hr(min int) (h, m int) {
	var hh, mm int
	if min >= 0 {
		mm = min % (24 * 60)
	} else if min < 0 {
		mm = 60*24 + min%(24*60)
	}
	hh = mm / 60
	mm = mm % 60
	return hh, mm
}

func (c Clock) String() string {
	var hh, m string
	if c.m%60 < 10 {
		m = "0" + strconv.Itoa(c.m%60)
	} else {
		m = strconv.Itoa(c.m % 60)
	}

	if c.h%24 < 10 {
		hh = "0" + strconv.Itoa(c.h%24)
	} else {
		hh = strconv.Itoa(c.h % 24)
	}

	s := hh + ":" + m
	return s

}

func (c Clock) Add(minutes int) Clock {
	hh, mm := min2hr(c.h*60 + c.m + minutes)
	return Clock{hh, mm}
}
