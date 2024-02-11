package main

type board struct {
	cells  [][]int
	height int
	width  int
}

func gameOfLife(board [][]int) {
	if len(board) < 1 {
		return
	}
	b := newBoard(board)

	b = b.proceed()
	copy(board, b.cells)
}

func newBoard(brd [][]int) *board {
	if len(brd) == 0 || len(brd[0]) == 0 {
		return nil
	}
	res := board{
		cells:  make([][]int, len(brd)),
		height: len(brd),
		width:  len(brd[0]),
	}
	for i := range res.cells {
		res.cells[i] = make([]int, res.width)
		copy(res.cells[i], brd[i])
	}

	return &res
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
	nb := make([][]int, b.height)
	for i := range nb {
		nb[i] = make([]int, b.width)
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

			nb[y][x] = newAlive
		}
	}

	b.cells = nb
	return b
}
