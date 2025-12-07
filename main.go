package main

import (
	"fmt"
	"os"

	day1 "aoc-2025/day1"
	day2 "aoc-2025/day2"
	day3 "aoc-2025/day3"
	day4 "aoc-2025/day4"
	day5 "aoc-2025/day5"
)

type Day interface {
	Part1()
	Part2()
}

func main() {
	days := map[string]Day{
		"1": day1.Solution{},
		"2": day2.Solution{},
		"3": day3.Solution{},
		"4": day4.Solution{},
		"5": day5.Solution{},
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
