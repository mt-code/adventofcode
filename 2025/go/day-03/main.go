package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/mt-code/adventofcode/utils"
)

func main() {
	utils.Solve(part1, part2)
}

func part1(scanner *bufio.Scanner) (answer int) {
	for scanner.Scan() {
		var digits []int
		for _, c := range strings.Split(scanner.Text(), "") {
			d, _ := strconv.Atoi(c)
			digits = append(digits, d)
		}
		answer += maxJoltage(digits, 2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}

func part2(scanner *bufio.Scanner) (answer int) {
	for scanner.Scan() {
		var digits []int
		for _, c := range strings.Split(scanner.Text(), "") {
			d, _ := strconv.Atoi(c)
			digits = append(digits, d)
		}
		answer += maxJoltage(digits, 12)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}

func maxJoltage(digits []int, n int) (total int) {
	start := 0
	for remaining := n; remaining > 0; remaining-- {
		best := start
		for i := start; i <= len(digits)-remaining; i++ {
			if digits[i] > digits[best] {
				best = i
			}
		}
		total = total*10 + digits[best]
		start = best + 1
	}
	return total
}
