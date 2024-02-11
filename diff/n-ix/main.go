package main

import "fmt"

const (
	height, width    = 25, 25
	iterationsAmount = 50
)

type board struct {
	cells  [][]int
	height int
	width  int
}

func main() {
	b := newBoard(height, width)
	b.setGliderPattern()
	b.print("initialized board with glider pattern")

	for i := 1; i <= iterationsAmount; i++ {
		b = b.proceed()
		b.print(fmt.Sprint("iteration", i))
	}
}

func newBoard(height, width int) *board {
	res := board{
		cells:  make([][]int, 0),
		height: height,
		width:  width,
	}

	for i := 0; i < height; i++ {
		res.cells = append(res.cells, make([]int, res.width))
	}

	return &res
}

func (b *board) setGliderPattern() {
	y := b.height / 2
	x := b.width / 2
	if x < 0 || y < 0 || x > b.width-1 || y > b.height-1 {
		return
	}

	b.cells[y-1][x] = 1   // up
	b.cells[y+1][x] = 1   // down
	b.cells[y][x+1] = 1   // right
	b.cells[y+1][x-1] = 1 // down left diagonal
	b.cells[y+1][x+1] = 1 // down right diagonal
}

func (b *board) aliveNeighbours(x, y int) int {
	result := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if b.get(nx, ny) == 1 {
				result++
			}
		}
	}

	return result
}

func (b *board) get(x, y int) int {
	if x < 0 || y < 0 || x > b.width-1 || y > b.height-1 {
		return 0
	}

	return b.cells[y][x]
}

func (b *board) proceed() *board {
	localBoard := make([][]int, b.height)
	for i := range localBoard {
		localBoard[i] = make([]int, b.width)
	}

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			nbAmount := b.aliveNeighbours(x, y)
			current := b.get(x, y)
			newAlive := 0
			switch {
			case nbAmount < 2: // 1st rule - underpopulation
			case current == 1 && (nbAmount == 2 || nbAmount == 3): // 2nd rule - keep alive
				newAlive = 1
			case current == 1 && nbAmount > 3: // 3rd rule - overcrowding
			case current == 0 && nbAmount == 3: // 4th rule - reproduction
				newAlive = 1
			}

			localBoard[y][x] = newAlive
		}
	}

	b.cells = localBoard
	return b
}

func (b *board) print(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	for _, row := range b.cells {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}
