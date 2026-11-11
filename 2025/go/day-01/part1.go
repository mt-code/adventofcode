package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		isLeft := ins[0] == 'L'
		num, err := strconv.Atoi(ins[1:])
		if err != nil {
			log.Fatal(err)
		}

		if isLeft {
			current -= num
		} else {
			current += num
		}

		if current%MAX == MIN {
			current = MIN
		} else if current < MIN {
			current = MAX + (current % MAX)
		} else if current >= MAX {
			current = MIN + (current % MAX)
		}

		if current == 0 {
			count++
		}
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Count: ", count)
}
