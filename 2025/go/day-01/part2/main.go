package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

const (
	MIN = 0
	MAX = 100
)

func main() {
	count := 0
	current := 50

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		ins := buf.Text()
		val, _ := strconv.Atoi(ins[1:])

		if ins[0] == 'L' {
			val *= -1
		}

		vals := []int{current, val + current}
		slices.Sort(vals)

		// count times we cross 0
		for i := vals[0] + 1; i < vals[1]; i++ {
			if i == 0 || i%MAX == 0 {
				count++
			}
		}

		current = (val + current) % MAX

		// count if we land on 0
		if current == 0 {
			count++
		}
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Count: ", count)
}
