package day2

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// =========================================
// Day 1
// =========================================

type Range struct{ start, end int }

func ParseRangeInput(s string) []Range {
	parts := strings.Split(s, ",")
	r := make([]Range, 0, len(parts))

	for _, part := range parts {
		nums := strings.Split(part, "-")
		a, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal("Cannot parse num input: ", part)
		}
		b, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal("Cannot parse num input: ", part)
		}
		r = append(r, Range{start: a, end: b})
	}

	return r
}

func ParseInput() (r []Range, err error) {
	// file, err := os.Open("day2/sample.txt")
	file, err := os.Open("day2/input.txt")
	if err != nil {
		return nil, errors.New("soemthing went wrong while parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := ParseRangeInput(line)
		return r, nil
	}

	return nil, errors.New("Should not come here")
}

func checkIsInvalid(num int) bool {
	s := strconv.Itoa(num)

	// nums with odd length could not be invalid
	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	is_invalid := true
	for i := range mid {
		is_invalid = is_invalid && (s[i] == s[i+mid])
	}
	return is_invalid
}

func checkIsInvalidAdvanced(num int) bool {
	s := strconv.Itoa(num)
	mid := len(s) / 2

	r := []rune(s)
	is_repeating := false

	for i := 1; i <= mid && !is_repeating; i++ {
		repeating := true
		firstChunk := string(r[:i])

		for chunk := range slices.Chunk(r, i) {
			if string(chunk) != firstChunk {
				repeating = false
				break
			}
		}

		if repeating {
			is_repeating = true
			break
		}
	}

	return is_repeating
}

type Solution struct{}

func (day Solution) Part1() {
	r, _ := ParseInput()
	counter := 0

	for _, part := range r {
		a := part.start
		b := part.end

		for i := a; i <= b; i++ {
			if checkIsInvalid(i) {
				counter += i
			}
		}
	}

	fmt.Println("Part 1 Answer: ", counter)
}

func (Solution) Part2() {
	counter := 0

	r, _ := ParseInput()
	for _, part := range r {
		a := part.start
		b := part.end

		for i := a; i <= b; i++ {
			if checkIsInvalidAdvanced(i) {
				counter += i
			}
		}
	}

	fmt.Println("Part 2 Answer: ", counter)
}
