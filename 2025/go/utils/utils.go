package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver func(scanner *bufio.Scanner) int

func Solve(solvers ...Solver) {
	for i, solver := range solvers {
		f, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		answer := solver(scanner)
		fmt.Println(fmt.Sprintf("Part %d answer:", i+1), answer)
	}
}
