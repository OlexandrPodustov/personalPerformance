package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	cc := image[sr][sc]
	if cc == color {
		return image
	}
	fill(image, sr, sc, cc, color)

	return image
}

func fill(image [][]int, r, c, colorSubstitute, color int) {
	if r >= 0 && r < len(image) {
		if c >= 0 && c < len(image[0]) {
			if image[r][c] != colorSubstitute {
				return
			}
			image[r][c] = color

			// move up
			fill(image, r-1, c, colorSubstitute, color)

			// move down
			fill(image, r+1, c, colorSubstitute, color)

			// move left
			fill(image, r, c-1, colorSubstitute, color)

			// move right
			fill(image, r, c+1, colorSubstitute, color)
		}
	}

	return
}
