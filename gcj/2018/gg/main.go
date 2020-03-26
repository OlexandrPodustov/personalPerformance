package main

import (
	"fmt"
	"log"
	"math"
)

//1 ≤ T ≤ 20.
//A = 20.
func main() {
	var t int
	var size = 1000
	var matrix = make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}
	//fmt.Println("m:", matrix)
	//fmt.Print("T:")
	fmt.Scanf("%d", &t)
	for ij := 1; ij <= t; ij++ {
		var i, j, n, c, r int
		fmt.Scanf("%d", &n)
		sqr := math.Sqrt(float64(n))
		c = int(sqr)
		if c*c == n {
			r = c
		} else if sqr > float64(c) {
			r = c + 1
		}

		fmt.Println("n", n)
		fmt.Println("r", r)
		fmt.Println("c", c)
		r0, c0 := 10, 10
		for {
			fmt.Println(r0, c0)
			//fmt.Printf("i and j:")
			fmt.Scanf("%d %d", &i, &j)
			if i == 0 && j == 0 {
				break
			} else if i == -1 && j == -1 {
				log.Fatalln("-1 -1 received from the judge")
			} else {
				fmt.Println("i j", i, j)
				matrix[i][j] = 1
			}
		}
		rect := matrix[r0 : r0+r]
		for ii := 0; ii < len(rect); ii++ {
			rect[ii] = rect[ii][c0 : c0+c]
		}
		fmt.Println("rect", len(rect))
		//fmt.Println("rect", rect)
	}
}
