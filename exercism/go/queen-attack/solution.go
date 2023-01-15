package queenattack

import (
	"errors"
	"strconv"
)

const testVersion = 2

func CanQueenAttack(b, w string) (res bool, err error) {
	if b == w {
		err = errors.New("input values are equal, two figures can't stand on the same place")
		return
	}
	if len(b) > 2 || len(w) > 2 {
		err = errors.New("input values are longer than expected")
		return
	}

	letterBlack := int(b[:1][0])
	letterWhite := int(w[:1][0])
	if letterBlack < 97 || letterBlack > 104 || letterWhite < 97 || letterWhite > 104 {
		err = errors.New("incorrect input letter value")
		return
	}

	intBlack, err := strconv.Atoi(b[1:])
	if err != nil {
		return
	}

	intWhite, err := strconv.Atoi(w[1:])
	if err != nil {
		return
	}

	// check board boundaries
	if intBlack < 1 || intWhite < 1 || intBlack > 8 || intWhite > 8 {
		err = errors.New("some or even both of input values are bigger or smaller than chessboard")
	}
	// horisontal check
	if letterBlack == letterWhite {
		res = true
		return
	}
	// vertical check
	if intBlack == intWhite {
		res = true
		return
	}
	// diagonal check
	for i := 0; i < 8; i++ {
		if letterWhite == letterBlack+i || letterBlack == letterWhite+i {
			if intWhite == intBlack+i || intBlack == intWhite+i {
				res = true
				return
			}
		}
	}
	return
}
