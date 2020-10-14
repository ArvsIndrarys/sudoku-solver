package sudokusolver

import (
	"errors"
	"fmt"
	"log"
)

type grid struct {
	group  map[int][]int
	line   map[int][]int
	column map[int][]int
}

var possibleValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func solve(input [][]int) {

	err := verifyInputData(input)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateLines(input [][]int) map[int][]int {
	lines := make(map[int][]int)

	for i, line := range input {

		if len(line) != 9 {
			log.Fatalf("Wrong number of elements at line %d. Expected 9, got %d", i+1, len(line))
		}

		duplicate := getFirstDuplicateInLine(line)
		if duplicate != 0 {
			log.Fatalf("Found duplicate at line %d: %d", i+1, duplicate)
		}

		lines[i] = line
	}

	return lines
}

func generateColumns(input [][]int) map[int]int {

	if len(input) != 9 {
		log.Fatalf("Wrong number of columns. Expected 9, got %d", len(input))
	}

	columns := make(map[int][]int)

}

func generateGroups(input [][]int) {

}

func generateGrid(input [][]int) grid {

	return grid{
		group:   generateGroups(input),
		line:    generateLines(input),
		columns: generateColumns(input),
	}
}

func verifyInputData(input [][]int) error {

	if len(input) != 9 {
		return errors.New(fmt.Sprintf("Wrong number of lines on the grid. Expected 9, got %d", len(input)))
	}

	for i, line := range input {
		if len(line) != 9 {
			return errors.New(fmt.Sprintf("Wrong number of elements at line %d. Expected 9, got %d", i+1, len(line)))
		}
	}

	return checkDuplicatesInGrid(input)
}

func checkDuplicatesInGrid(input [][]int) error {

	// check on lines
	for _, v := range input {

		err := checkLineCorrectness(v)
		if err != nil {
			log.Fatal(err)
		}

	}

	// check on columns

	// check in squares

	return nil
}

func checkLineCorrectness(line []int) error {

	if len(line) != 9 {
		return errors.New(fmt.Sprintf("Wrong length of line/column. Expected 9, got %d", len(line)))
	}

	elements := make([]int, 9)
	copy(elements, possibleValues)

	for _, v := range line {

		if v < 0 || v > 9 {
			return errors.New(fmt.Sprintf("Wrong data present in grid: %d", v))
		}

		if v != 0 {
			if !checkExist(elements, v) {
				return errors.New(fmt.Sprintf("Duplicate %d found in line", v, input))
			}

			elements = removeValue(elements, v)
		}
	}
	return nil
}

func resolveLine(line []int) []int {

	var count, index int

	possible := make([]int, 9)
	copy(possible, possibleValues)

	for i, v := range line {
		if v == 0 {
			count++
			index = i
			continue
		}
		possible = removeValue(possible, v)
	}

	if count == 0 {
		return line
	}

	if count > 1 {
		return line
	}

	line[index] = possible[0]
	return line

}

func removeValue(slice []int, value int) []int {

	for i, v := range slice {
		if v == value {
			slice[i] = slice[len(slice)-1]

			tempSlice := slice[:len(slice)-1]
			return tempSlice
		}
	}

	return nil
}

func replaceValue(slice []int, index, value int) []int {

	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return nil
}

func checkExist(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
