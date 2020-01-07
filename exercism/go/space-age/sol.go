package space

import (
	"math"
)

const (
	Es      = 31557600
	Earth   = 365.25
	Mercury = 0.2408467 * Es
	Venus   = 0.61519726 * Es
	Mars    = 1.8808158 * Es
	Jupiter = 11.862615 * Es
	Saturn  = 29.447498 * Es
	Uranus  = 84.016846 * Es
	Neptune = 164.79132 * Es
)

type Planet string

func Age(s float64, p Planet) (a float64) {

	switch {
	case p == "Earth":
		a = s / Es
	case p == "Mercury":
		a = s / Mercury
	case p == "Venus":
		a = s / Venus
	case p == "Mars":
		a = s / Mars
	case p == "Jupiter":
		a = s / Jupiter
	case p == "Saturn":
		a = s / Saturn
	case p == "Uranus":
		a = s / Uranus
	case p == "Neptune":
		a = s / Neptune

	default:
		a = 0
	}
	a = round(a)
	return a
}

func round(i float64) float64 {
	roundedVal := float64(int(i*100)) / 100

	_, decimalpart := math.Modf(i * 100)

	dec := int(decimalpart * 10)
	rv := roundedVal * 100
	if dec+5 > 10 {
		rv++
	}
	roundedVal = rv / 100
	return roundedVal
}
