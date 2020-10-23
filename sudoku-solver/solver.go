package sudokusolver

import (
	"errors"
	"fmt"
)

type grid struct {
	Groups  map[int][]int
	Lines   map[int][]int
	Columns map[int][]int
}

var possibleValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func generateGrid(input [][]int) (grid, error) {

	columns, err := generateColumns(input)
	if err != nil {
		return grid{}, err
	}
	lines, err := generateLines(input)
	if err != nil {
		return grid{}, err
	}
	groups, err := generateGroups(input)
	if err != nil {
		return grid{}, err
	}
	return grid{
		Columns: columns,
		Lines:   lines,
		Groups:  groups,
	}, nil
}

func generateLines(input [][]int) (map[int][]int, error) {
	lines := make(map[int][]int)

	for i, line := range input {

		if len(line) != 9 {
			return map[int][]int{}, errors.New(fmt.Sprintf("Wrong number of elements at line %d. Expected 9, got %d", i+1, len(line)))
		}

		if err := checkLineCorrectness(line); err != nil {
			return map[int][]int{}, err
		}

		lines[i] = line
	}

	return lines, nil
}

func generateColumns(input [][]int) (map[int][]int, error) {

	if len(input) != 9 {
		return map[int][]int{}, errors.New(fmt.Sprintf("Wrong number of columns. Expected 9, got %d", len(input)))
	}

	columns := make(map[int][]int)

	// columns map creation
	for _, line := range input {

		for columnIndex, cellValue := range line {
			columns[columnIndex] = append(columns[columnIndex], cellValue)
		}
	}

	// columns map verification
	for i, line := range columns {
		if len(line) != 9 {
			return map[int][]int{}, errors.New(fmt.Sprintf("Wrong number of elements at line %d. Expected 9, got %d", i+1, len(line)))
		}

		if err := checkLineCorrectness(line); err != nil {
			return map[int][]int{}, err
		}
	}

	return columns, nil
}

// generateGroups() call MUST BE AFTER generateColumns() and generateLines()
//as they already check data correctness, allowing us to avoid these checks
func generateGroups(input [][]int) (map[int][]int, error) {
	groups := make(map[int][]int)

	// We slice the line 3 elements by 3 elements
	// As line has 9 elements, we iterate 3 times
	for i := 0; i < 3; i++ {

		// We slice the columns 3 by 3 to get the 3x3 square and add it to our group
		for j := 0; j < 3; j++ {
			groups[3*i+j] = append(groups[3*i+j], input[3*i][3*j:3*j+3]...)
			groups[3*i+j] = append(groups[3*i+j], input[3*i+1][3*j:3*j+3]...)
			groups[3*i+j] = append(groups[3*i+j], input[3*i+2][3*j:3*j+3]...)
		}
	}

	// check the correctness of our groups
	for _, group := range groups {

		if err := checkLineCorrectness(group); err != nil {
			return map[int][]int{}, err
		}
	}

	return groups, nil
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
				return errors.New(fmt.Sprintf("Duplicate %d found in line %v", v, line))
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

func Map2DToString(grid map[int][]int) string {
	s := ""
	for _, v := range grid {
		s = s + fmt.Sprintf("%v\n", v)
	}

	return s
}
func Array2DToString(grid [][]int) string {
	s := ""
	for _, v := range grid {
		s = s + fmt.Sprintf("%v\n", v)
	}

	return s
}
