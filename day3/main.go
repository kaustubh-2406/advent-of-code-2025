package day3

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ParseInput() (r []string, err error) {
	// file, err := os.Open("day3/sample.txt")
	file, err := os.Open("day3/input.txt")
	if err != nil {
		return nil, errors.New("soemthing went wrong while parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}

	return lines, nil
}

func maxSubsequenceDigits(s string, k int) string {
	stack := make([]byte, 0, k) // basket to hold chosen digits
	toRemove := len(s) - k      // how many digits Grug can throw away

	for i := 0; i < len(s); i++ {
		d := s[i] // current digit

		// While basket not empty,
		// and Grug allowed to remove more,
		// and new digit bigger than last in basket → pop it
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1] // throw last small digit
			toRemove--
		}

		stack = append(stack, d) // take new digit
	}

	// maybe we kept too many → trim to exactly k
	return string(stack[:k])
}

type Solution struct{}

func (day Solution) Part1() {
	sum := 0
	lines, _ := ParseInput()

	for _, line := range lines {
		num := 0
		for i, ch := range line {
			for j := i + 1; j < len(line); j++ {
				s := string(ch) + string(line[j])
				n, _ := strconv.Atoi(s)
				if n > num {
					num = n
				}
			}
		}

		sum += num
	}

	fmt.Println("Part 1 Answer: ", sum)
}

func (Solution) Part2() {
	sum := 0
	lines, _ := ParseInput()
	for _, line := range lines {
		num := maxSubsequenceDigits(line, 12)
		n, _ := strconv.Atoi(num)
		sum += n
	}

	fmt.Println("Part 2 Answer: ", sum)
}
