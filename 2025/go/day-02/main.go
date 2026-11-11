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
		ranges := strings.Split(scanner.Text(), ",")

		for _, r := range ranges {
			ids := strings.Split(r, "-")
			start, _ := strconv.Atoi(ids[0])
			end, _ := strconv.Atoi(ids[1])

			for i := start; i <= end; i += 1 {
				str := strconv.Itoa(i)

				if len(str)%2 != 0 {
					continue
				}
				if str[0:len(str)/2] == str[(len(str)/2):] {
					answer += i
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}

func part2(scanner *bufio.Scanner) (answer int) {
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")

		for _, r := range ranges {
			ids := strings.Split(r, "-")
			start, _ := strconv.Atoi(ids[0])
			end, _ := strconv.Atoi(ids[1])

		IdLoop:
			for i := start; i <= end; i += 1 {
				str := strconv.Itoa(i)

				// iterate through possible match sizes, size can't be more than half the string length
				for size := 1; size <= (len(str)/2)+1; size++ {
					// take base string of size length for comparison
					base := str[0:size]
					match := true

					if len(str)%size != 0 || len(str) <= 1 || len(str) == size {
						continue
					}

					// compare each chunk of the id with the base, if one doesn't match update flag
					for chunk := 1; chunk < len(str)/size; chunk++ {
						offset := size * chunk

						if str[offset:offset+size] != base {
							match = false
						}
					}

					// if all chunks matched, it is considered an invalid id and counted towards answer
					if match {
						answer += i
						continue IdLoop
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}
