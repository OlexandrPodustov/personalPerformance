package main

import (
	"math"
	"strconv"
)

func main() {
	type Clock struct {
		in int
	}
	n1 := Clock{23}
	n2 := Clock{23}
	println(n1 == n2)
}

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	var c Clock
	//var newhh, newmm int

	//handling minutes
	if minute > 0 {
		c.m = minute % 60
	} else if minute < 0 {
		c.m = 60 + minute%60
		hour--
	}
	//if there are more minutes than 59
	if math.Abs(float64(minute)) > 59 {
		hour += minute / 60
	}

	//handling hours
	if hour < 0 {
		c.h = 24 + c.h%24
	} else {
		c.h = hour
	}
	return c

}

func (c Clock) String() string {

	//adding zero before hour or minute
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
	return Clock{c.h, c.m + minutes}
}
