package main

import (
	"bufio"
	"log"

	"github.com/mt-code/adventofcode/utils"
)

func main() {
	utils.Solve(part1)
}

func part1(scanner *bufio.Scanner) (answer int) {
	y := 0
	grid := make([][]byte, 0)

	// create set of neighbour tiles around current x,y coord
	neighbours := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for y = 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// only process paper rolls
			if grid[y][x] == '.' {
				continue
			}

			count := 0
			for _, n := range neighbours {
				nX, nY := x+n[0], y+n[1]
				// if target neighbour tile is oob, skip and count as free space
				if nY < 0 || nY >= len(grid) || nX < 0 || nX >= len(grid[y]) {
					count++
					continue
				}
				if grid[nY][nX] == '.' {
					count++
				}
			}
			if count > 4 {
				answer++
			}
		}
	}

	return answer
}
