package sudokusolver

import (
	"errors"
	"fmt"
)

func removeValue(slice []int, value int) []int {

	for i, v := range slice {
		if v == value {
			slice[i] = slice[0]
			return slice[1:]
		}
	}
	return slice
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

	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

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
