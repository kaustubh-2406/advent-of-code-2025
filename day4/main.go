package day4

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GridElement int

const (
	EMPTY GridElement = iota
	PAPER_ROLL
)

type Grid struct {
	row   int
	col   int
	items map[string]GridElement
}

func NewGrid() Grid {
	return Grid{
		row:   0,
		col:   0,
		items: make(map[string]GridElement),
	}
}

func (g Grid) decodeKey(key string) (row int, col int) {
	parts := strings.Split(key, ",")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a, b
}

func (g Grid) encodeKey(i, j int) string {
	key := fmt.Sprintf("%d,%d", i, j)
	return key
}

func (g *Grid) AddItem(i, j int, value GridElement) {
	key := g.encodeKey(i, j)
	g.items[key] = value
}

func (g *Grid) GetItem(i, j int) (element GridElement, ok bool) {
	key := g.encodeKey(i, j)
	value, ok := g.items[key]
	return value, ok
}

func (g *Grid) CountAdjacentPaperRolls(key string) int {
	//
	neighbour := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	neigbour_count := 0
	i, j := g.decodeKey(key)
	// fmt.Printf("(%d, %d) =================== \n", i, j)
	for _, n := range neighbour {
		dx := n[0]
		dy := n[1]

		value, ok := g.GetItem(i+dx, j+dy)
		// fmt.Println(i+dx, ",", j+dy, " => value =", value, ok)
		if ok && value == PAPER_ROLL {
			neigbour_count += 1
		}
	}

	// fmt.Println("key", key, "rollcount = ", neigbour_count)
	return neigbour_count
}

func ParseInput() (r Grid, err error) {
	// file, err := os.Open("day4/sample.txt")
	file, err := os.Open("day4/input.txt")
	if err != nil {
		return Grid{}, errors.New("soemthing went wrong while parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := NewGrid()

	for scanner.Scan() {
		line := scanner.Text()

		g.col = len(line)
		g.col = len(line)
		for col_count, r := range line {
			if string(r) == "@" {
				g.AddItem(g.row, col_count, PAPER_ROLL)
			}
		}

		g.col = len(line)
		g.row += 1
	}

	return g, nil
}

type Solution struct{}

func (day Solution) Part1() {
	grid, _ := ParseInput()

	keys := make([]string, 0, len(grid.items))
	for k := range grid.items {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	count := 0
	for _, k := range keys {
		rollCount := grid.CountAdjacentPaperRolls(k)
		if rollCount < 4 {
			count += 1
		}
	}

	fmt.Println("Part 1 Answer: ", count)
}

func (Solution) Part2() {
	count := 0
	grid, _ := ParseInput()

	for {
		keys := make([]string, 0, len(grid.items))
		for k := range grid.items {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		rollRemoved := false
		for _, k := range keys {
			fmt.Println("k = ", k)
			rollCount := grid.CountAdjacentPaperRolls(k)
			if rollCount < 4 {
				rollRemoved = true
				count += 1
				delete(grid.items, k)
			}
		}
		if !rollRemoved {
			break
		}
	}

	fmt.Println("Part 2 Answer: ", count)
}
