package main

import (
	"fmt"
	"os"

	day1 "aoc-2025/day1"
)

type Day interface {
	Part1()
	Part2()
}

func main() {
	days := map[string]Day{
		"1": day1.Day1{},
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run . <day-number>")
		return
	}

	for _, d := range args {
		if solver, ok := days[d]; ok {
			fmt.Printf("== Day %s ==\n", d)
			solver.Part1()
			solver.Part2()
		} else {
			fmt.Printf("Unknown day: %s\n", d)
		}
	}

}
