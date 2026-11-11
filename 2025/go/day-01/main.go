package main

import (
	"bufio"
	"log"
	"slices"
	"strconv"

	"github.com/mt-code/adventofcode/utils"
)

const (
	MAX = 100
)

func main() {
	utils.Solve(part1, part2)
}

func part1(scanner *bufio.Scanner) (answer int) {
	current := 50

	for scanner.Scan() {
		ins := scanner.Text()
		val, _ := strconv.Atoi(ins[1:])

		if ins[0] == 'L' {
			val *= -1
		}

		current = (val + current) % MAX
		if current == 0 {
			answer++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}

func part2(scanner *bufio.Scanner) (answer int) {
	current := 50

	for scanner.Scan() {
		ins := scanner.Text()
		val, _ := strconv.Atoi(ins[1:])

		if ins[0] == 'L' {
			val *= -1
		}

		vals := []int{current, val + current}
		slices.Sort(vals)

		// count times we cross 0
		for i := vals[0] + 1; i < vals[1]; i++ {
			if i == 0 || i%MAX == 0 {
				answer++
			}
		}

		current = (val + current) % MAX

		// count if we land on 0
		if current == 0 {
			answer++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}
