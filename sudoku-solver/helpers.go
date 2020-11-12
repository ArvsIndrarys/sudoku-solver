package sudokusolver

import (
	"errors"
	"fmt"
)

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

func checkCorrectness(line []int) error {

	if len(line) != 9 {
		return errors.New(fmt.Sprintf("Wrong length of line/column/square. Expected 9, got %d", len(line)))
	}

	elements := make([]int, 9)
	copy(elements, possibleValues)

	for _, v := range line {

		if v < 0 || v > 9 {
			return errors.New(fmt.Sprintf("Wrong data present in grid: %d", v))
		}

		if v != 0 {
			if !checkExist(elements, v) {
				return errors.New(fmt.Sprintf("Duplicate %d found in line/column/square %v", v, line))
			}

			elements = removeValue(elements, v)
		}
	}
	return nil
}
