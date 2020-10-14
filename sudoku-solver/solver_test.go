package sudokusolver

import (
	"errors"
	"reflect"
	"testing"
)

var (
	correctInput = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	wrongInputLines = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	wrongInputColumns = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 9},
		{1, 2, 3, 4, 5, 6, 7, 8},
	}

	wrongInputColumnsOneLine = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
)

func TestCorrectInputData(t *testing.T) {

	err := verifyInputData(correctInput)
	if err != nil {
		t.Errorf("Got an error when not expected: %v", err)

	}
}

func TestWrongNumberOfLinesInputData(t *testing.T) {
	err := verifyInputData(wrongInputLines)
	expected := errors.New("Wrong number of lines on the grid. Expected 9, got 8")
	if err == nil {
		t.Errorf("Got no error when one was expected")

	}
	if err.Error() != expected.Error() {
		t.Errorf("Wrong error msg, got %v and expected %v", err, expected)
	}
}

func TestWrongNumberOfColumnsInputData(t *testing.T) {
	err := verifyInputData(wrongInputColumns)
	expected := errors.New("Wrong number of elements at line 1. Expected 9, got 8")
	if err == nil {
		t.Errorf("Got no error when one was expected")

	}
	if err.Error() != expected.Error() {
		t.Errorf("Wrong error msg, got %v and expected %v", err, expected)
	}
}

func TestWrongNumberOfColumnsOnOneLineInputData(t *testing.T) {
	err := verifyInputData(wrongInputColumnsOneLine)
	expected := errors.New("Wrong number of elements at line 3. Expected 9, got 8")
	if err == nil {
		t.Errorf("Got no error when one was expected")

	}
	if err.Error() != expected.Error() {
		t.Errorf("Wrong error msg, got %v and expected %v", err, expected)
	}
}
func TestResolveLine(t *testing.T) {

	exampleUnfilledLine := []int{2, 3, 5, 8, 4, 0, 9, 7, 1}

	expected := []int{2, 3, 5, 8, 4, 6, 9, 7, 1}

	solution := resolveLine(exampleUnfilledLine)

	if !reflect.DeepEqual(expected, solution) {
		t.Error("Slices are not equal : ", expected, solution)
	}

}

func TestCheckDuplicateLine(t *testing.T) {

	completedLine := []int{2, 3, 1, 8, 4, 7, 5, 6, 9}

	lineWithSeveralMissingNumbers := []int{2, 3, 0, 8, 4, 0, 5, 0, 0}

	lineWithDuplicates := []int{2, 3, 5, 8, 4, 0, 2, 7, 1}

	duplicate, hasDuplicates := checkDuplicatesInLine(completedLine)
	if hasDuplicates {
		t.Error("Should not have found duplicates on line ok ", completedLine)
	}

	duplicate, hasDuplicates = checkDuplicatesInLine(lineWithSeveralMissingNumbers)
	if hasDuplicates {
		t.Error("Should not have found duplicates on line with missing numbers: ", lineWithSeveralMissingNumbers)
	}

	duplicate, hasDuplicates = checkDuplicatesInLine(lineWithDuplicates)
	if !hasDuplicates && duplicate != 2 {
		t.Error("Should have reported duplicated 2 in line: ", lineWithDuplicates)
	}
}
