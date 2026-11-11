package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/mt-code/adventofcode/utils"
)

func main() {
	utils.Solve(part1)
}

func part1(scanner *bufio.Scanner) (answer int) {
	for scanner.Scan() {
		var digits []int
		for _, c := range strings.Split(scanner.Text(), "") {
			d, _ := strconv.Atoi(c)
			digits = append(digits, d)
		}
		idx, d1 := getHighestDigit(digits[:len(digits)-1])
		_, d2 := getHighestDigit(digits[idx+1:])
		answer += (d1 * 10) + d2
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}

func getHighestDigit(digits []int) (idx int, digit int) {
	for i, d := range digits {
		if d > digit {
			digit, idx = d, i
		}
	}
	return idx, digit
}
