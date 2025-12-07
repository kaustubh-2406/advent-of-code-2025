package day5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IngredientId int
type IngredientList []IngredientId
type IngridientRange struct {
	start IngredientId
	end   IngredientId
}

// Corrected comparator (note: not pointer receiver required)
func (ir IngridientRange) CompareTo(id IngredientId) int {
	// Return negative if ir < id
	if ir.end < id {
		return -1 // range ends before id -> range precedes target
	}
	// Return positive if ir > id
	if ir.start > id {
		return 1 // range starts after id -> range follows target
	}
	// Else id is within [start, end]
	return 0
}

type Inventory struct {
	FreshIngridients   []IngridientRange
	IngridientsToCheck IngredientList
}

func CheckIngridientInInventory(inventory *Inventory, id IngredientId) bool {
	slices.SortFunc(inventory.FreshIngridients, func(a, b IngridientRange) int {
		if a.start < b.start {
			return -1
		}
		if a.start > b.start {
			return 1
		}
		return 0
	})

	found := false
	for _, r := range inventory.FreshIngridients {
		if r.start < id && r.end > id {
			found = true
			break
		}
	}
	return found

	// TODO: need to figure out how we can make the binary search approach work
	// _, found := slices.BinarySearchFunc(
	// 	inventory.FreshIngridients,
	// 	ingredientId,
	// 	func(freshIngridient IngridientRange, id IngredientId) int {
	// 		return freshIngridient.CompareTo(id)
	// 	})
	// return found
}

func ParseInput() (r Inventory, err error) {
	// file, err := os.Open("day5/sample.txt")
	file, err := os.Open("day5/input.txt")
	if err != nil {
		return Inventory{}, errors.New("soemthing went wrong while parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := Inventory{}
	encounteredEmptyLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 { // empty line
			encounteredEmptyLine = true
			continue
		}

		if encounteredEmptyLine {
			id := parseIngridientToCheck(line)
			input.IngridientsToCheck = append(input.IngridientsToCheck, id)
		} else {
			ingridients := parseFreshIngridients(line)
			input.FreshIngridients = append(input.FreshIngridients, ingridients)
		}
	}

	return input, nil
}

func parseIngridientToCheck(line string) IngredientId {
	num, _ := strconv.Atoi(line)
	return IngredientId(num)
}

func parseFreshIngridients(line string) IngridientRange {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		panic("error: expected 2 items to be parts")
	}

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return IngridientRange{start: IngredientId(start), end: IngredientId(end)}
}

type Solution struct{}

func (day Solution) Part1() {
	count := 0
	inventory, _ := ParseInput()

	for _, id := range inventory.IngridientsToCheck {
		found := CheckIngridientInInventory(&inventory, id)
		// fmt.Println("found ?", found)
		if found {
			count += 1
		}
	}

	fmt.Println("Part 1 Answer: ", count)
}

func (Solution) Part2() {
	count := 0

	// TODO: optimize it, as it crashes without producing response

	// ingridients, _ := ParseInput()
	// set := make(map[IngredientId]struct{})
	//
	// for _, ingredientRange := range ingridients.FreshIngridients {
	// 	for i := ingredientRange.start; i <= ingredientRange.end; i++ {
	// 		_, found := set[i]
	// 		if found {
	// 			continue
	// 		}
	//
	// 		count += 1
	// 		set[i] = struct{}{}
	// 	}
	// }

	fmt.Println("Part 2 Answer: ", count)
}
