package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	result := 0
	games := parseInput(string(data))
	for _, game := range games {
		result += game.redMax * game.greenMax * game.blueMax
	}

	fmt.Println("result", result)
}

func solve(lines []string) int {
	result := 0

	return result
}

type game struct {
	order int
	blueMax,
	redMax,
	greenMax int
}

func parseInput(inp string) []game {
	games := strings.Split(inp, "\n")

	// fmt.Println("games len", len(games))
	result := make([]game, 0, len(games)-1)
	for _, v := range games {
		if len(v) < 2 {
			break
		}
		firstPart := strings.Split(v[:strings.Index(v, ":")], " ")
		// fmt.Println("firstPart ", firstPart[1], len(firstPart))

		order, _ := strconv.Atoi(firstPart[1])
		game := game{
			order: order,
		}

		secondPart := strings.Split(v[strings.Index(v, ":")+1:], ";")
		// fmt.Println("game secondPart", len(secondPart), secondPart)
		for _, sp := range secondPart {
			// fmt.Println("sp", len(sp), sp)

			gamePart := strings.Split(sp, ",")
			// fmt.Println("gamePart", len(gamePart), gamePart)
			for _, gp := range gamePart {
				part := strings.Split(strings.TrimSpace(gp), " ")

				amount, _ := strconv.Atoi(part[0])

				// fmt.Println("part", len(part), part[1], amount)
				switch part[1] {
				case "green":
					if amount > game.greenMax {
						game.greenMax = amount
					}
				case "red":
					if amount > game.redMax {
						game.redMax = amount
					}
				case "blue":
					if amount > game.blueMax {
						game.blueMax = amount
					}
				}
			}

		}
		// fmt.Println("game unparsed", v)
		// fmt.Printf("game parsed %+v \n", game)
		result = append(result, game)
		// break
	}

	return result
}
