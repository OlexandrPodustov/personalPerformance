package main

import (
	"fmt"
	"log"
)

//1 ≤ T ≤ 20.
//A = 20.
func main() {
	var t int
	var size int = 1000
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

		if n%3 == 0 {
			r = n / 3
			c = r
		}

		fmt.Println("n", n)
		fmt.Println("r", r)
		fmt.Println("c", c)
		for {
			fmt.Println("10 10")
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
		rect := matrix[i : j+1]
		for ii := 0; ii < len(rect); ii++ {
			rect[ii] = rect[ii][i : j+1]
		}
		fmt.Println("rect", len(rect))
		fmt.Println("rect", rect)
	}
}
