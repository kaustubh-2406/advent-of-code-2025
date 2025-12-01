package day1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction string

const (
	LEFT  Direction = "Left"
	RIGHT Direction = "Right"
)

type Operation struct {
	dir   Direction
	turns int
}

var operations []Operation = make([]Operation, 0, 10000)

// =========================================
// dial struct
// =========================================
type Dial struct {
	current int
}

func (d *Dial) TurnLeft(num int)  { d.current = (d.current - num + 100) % 100 }
func (d *Dial) TurnRight(num int) { d.current = (d.current + num + 100) % 100 }

// =========================================
// Day 1
// =========================================
type Day1 struct{}

func ParseInput() error {
	// file, err := os.Open("day1/sample.txt")
	file, err := os.Open("day1/input.txt")
	if err != nil {
		return errors.New("soemthing went wrong while parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch line[0] {
		case 'L':
			a, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal("cannot convert string to int")
			}
			operations = append(operations, Operation{dir: LEFT, turns: a})
		case 'R':
			a, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal("cannot convert string to int")
			}
			operations = append(operations, Operation{dir: RIGHT, turns: a})
		default:
			log.Fatal("cannot parse")
		}

	}
	return nil
}

var err error = ParseInput()

func (day Day1) Part1() {
	counter := 0
	dial := Dial{current: 50}

	for _, operation := range operations {
		switch operation.dir {
		case RIGHT:
			dial.TurnRight(operation.turns)
		case LEFT:
			dial.TurnLeft(operation.turns)
		default:
			log.Fatal("fatal")
		}

		if dial.current == 0 {
			counter += 1
		}
	}

	fmt.Println("Part 1 Answer: ", counter)
}

func (Day1) Part2() {
	counter := 0
	dial := Dial{current: 50}

	for _, operation := range operations {
		for range operation.turns {
			switch operation.dir {
			case RIGHT:
				dial.TurnRight(1)
				if dial.current == 0 {
					counter += 1
				}

			case LEFT:
				dial.TurnLeft(1)
				if dial.current == 0 {
					counter += 1
				}
			default:
				log.Fatal("fatal")
			}
		}
	}
	fmt.Println("Part 2 Answer: ", counter)
}
