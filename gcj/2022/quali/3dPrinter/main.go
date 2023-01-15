package main

import (
	"fmt"
	"log"
	"math"
)

const imp = "IMPOSSIBLE"

type printer struct {
	cyan, magenta, yellow, black int
}

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		printers := []printer{}
		for i := 0; i < 3; i++ {
			var pr printer
			if _, err := fmt.Scan(&pr.cyan, &pr.magenta, &pr.yellow, &pr.black); err != nil {
				log.Println(testCaseNumber, "  dimensions:", err)
			}
			printers = append(printers, pr)
		}
		result := solve(printers)

		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(printers []printer) string {
	max := getMax()
	cm := findCommonMax(printers)
	if cm.Sum() < max {
		return imp
	}
	if cm.Sum() == max {
		return cm.String()
	}
	// for _, v := range printers {
	// 	fmt.Println("printer", v)
	// }
	// fmt.Println("findCommonMax", cm)
	commonPrinter := cm.extract()
	// fmt.Println("commonPrinter after extract", commonPrinter.Sum())

	return commonPrinter.String()
}

func findCommonMax(printers []printer) printer {
	max := getMax()

	maxCyan, maxMagenta, maxYellow, maxBlack := max, max, max, max

	for _, v := range printers {
		if v.cyan < maxCyan {
			maxCyan = v.cyan
		}
		if v.magenta < maxMagenta {
			maxMagenta = v.magenta
		}
		if v.yellow < maxYellow {
			maxYellow = v.yellow
		}
		if v.black < maxBlack {
			maxBlack = v.black
		}
	}

	return printer{
		cyan:    maxCyan,
		magenta: maxMagenta,
		yellow:  maxYellow,
		black:   maxBlack,
	}
}

func (p printer) Sum() int {
	return p.cyan + p.magenta + p.yellow + p.black
}

func (p printer) String() string {
	return fmt.Sprintf("%v %v %v %v", p.cyan, p.magenta, p.yellow, p.black)
}

func (p printer) extract() printer {
	max := getMax()
	if p.cyan-max >= 0 {
		p.cyan, p.magenta, p.yellow, p.black = 0, 0, 0, 0
		p.cyan = max
		return p
	}
	if p.magenta-max >= 0 {
		p.cyan, p.magenta, p.yellow, p.black = 0, 0, 0, 0
		p.magenta = max
		return p
	}
	if p.yellow-max >= 0 {
		p.cyan, p.magenta, p.yellow, p.black = 0, 0, 0, 0
		p.yellow = max
		return p
	}
	if p.black-max >= 0 {
		p.cyan, p.magenta, p.yellow, p.black = 0, 0, 0, 0
		p.black = max
		return p
	}

	max -= p.cyan
	if max < 0 {
		p.cyan += max
		p.magenta, p.yellow, p.black = 0, 0, 0
		return p
	}
	max -= p.magenta
	if max < 0 {
		p.magenta += max
		p.yellow, p.black = 0, 0
		return p
	}

	max -= p.yellow
	if max < 0 {
		p.yellow += max
		p.black = 0
		return p
	}
	max -= p.black
	if max < 0 {
		p.black += max
		return p
	}

	return printer{
		cyan:    0,
		magenta: 0,
		yellow:  0,
		black:   0,
	}
}

func getMax() int {
	return int(math.Pow10(6))
}
