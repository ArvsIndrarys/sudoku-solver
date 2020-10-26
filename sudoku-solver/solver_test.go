package sudokusolver

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var (
	correctOneLineOrdered   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	correctOneLineUnordered = []int{9, 6, 3, 2, 1, 7, 5, 8, 4}
	correctOneLineOnlyHoles = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	correctOneLineWithHoles = []int{9, 6, 0, 2, 1, 7, 0, 8, 4}

	incorrectOneLineTooShort               = []int{1, 2, 3, 4, 5, 6, 7}
	incorrectOneLineWithDuplicates         = []int{2, 9, 1, 3, 4, 5, 7, 8, 9}
	incorrectOneLineWithDuplicatesAndHoles = []int{9, 6, 0, 2, 1, 7, 0, 4, 4}
	incorrectOneLineWithNumberSup9         = []int{9, 6, 0, 2, 1, 7, 19, 4, 4}
	incorrectOneLineWithNumberInf0         = []int{9, 6, 0, 2, 1, 7, -5, 4, 4}

	correctGrid = [][]int{
		{5, 4, 3, 9, 2, 1, 8, 7, 6},
		{2, 1, 9, 6, 8, 7, 5, 4, 3},
		{8, 7, 6, 3, 5, 4, 2, 1, 9},
		{9, 8, 7, 4, 6, 5, 3, 2, 1},
		{3, 2, 1, 7, 9, 8, 6, 5, 4},
		{6, 5, 4, 1, 3, 2, 9, 8, 7},
		{7, 6, 5, 2, 4, 3, 1, 9, 8},
		{4, 3, 2, 8, 1, 9, 7, 6, 5},
		{1, 9, 8, 5, 7, 6, 4, 3, 2},
	}

	correctGridAsColumns = [][]int{
		{5, 2, 8, 9, 3, 6, 7, 4, 1},
		{4, 1, 7, 8, 2, 5, 6, 3, 9},
		{3, 9, 6, 7, 1, 4, 5, 2, 8},
		{9, 6, 3, 4, 7, 1, 2, 8, 5},
		{2, 8, 5, 6, 9, 3, 4, 1, 7},
		{1, 7, 4, 5, 8, 2, 3, 9, 6},
		{8, 5, 2, 3, 6, 9, 1, 7, 4},
		{7, 4, 1, 2, 5, 8, 9, 6, 3},
		{6, 3, 9, 1, 4, 7, 8, 5, 2},
	}

	correctGridAsGroups = [][]int{
		{5, 4, 3, 2, 1, 9, 8, 7, 6},
		{9, 2, 1, 6, 8, 7, 3, 5, 4},
		{8, 7, 6, 5, 4, 3, 2, 1, 9},
		{9, 8, 7, 3, 2, 1, 6, 5, 4},
		{4, 6, 5, 7, 9, 8, 1, 3, 2},
		{3, 2, 1, 6, 5, 4, 9, 8, 7},
		{7, 6, 5, 4, 3, 2, 1, 9, 8},
		{2, 4, 3, 8, 1, 9, 5, 7, 6},
		{1, 9, 8, 7, 6, 5, 4, 3, 2},
	}

	incorrectGridWrongNumberOfLines = [][]int{
		{5, 4, 3, 9, 2, 1, 8, 7, 6},
		{2, 1, 9, 6, 8, 7, 5, 4, 3},
		{8, 7, 6, 3, 5, 4, 2, 1, 9},
		{9, 8, 7, 4, 6, 5, 3, 2, 1},
		{3, 2, 1, 7, 9, 8, 6, 5, 4},
		{6, 5, 4, 1, 3, 2, 9, 8, 7},
		{7, 6, 5, 2, 4, 3, 1, 9, 8},
		{4, 3, 2, 8, 1, 9, 7, 6, 5},
	}

	incorrectGridWrongNumberOfColumns = [][]int{
		{5, 4, 3, 9, 2, 1, 8, 7},
		{2, 1, 9, 6, 8, 7, 5, 4},
		{8, 7, 6, 3, 5, 4, 2, 1},
		{9, 8, 7, 4, 6, 5, 3, 2},
		{3, 2, 1, 7, 9, 8, 6, 5},
		{6, 5, 4, 1, 3, 2, 9, 8},
		{7, 6, 5, 2, 4, 3, 1, 9},
		{4, 3, 2, 8, 1, 9, 7, 6},
		{1, 9, 8, 5, 7, 6, 4, 3},
	}

	incorrectGridDuplicateInColumn = [][]int{
		{5, 4, 3, 9, 2, 1, 8, 7, 6},
		{2, 1, 9, 4, 8, 7, 5, 6, 3},
		{8, 7, 6, 3, 5, 4, 2, 1, 9},
		{9, 8, 7, 4, 6, 5, 3, 2, 1},
		{3, 2, 1, 7, 9, 8, 6, 5, 4},
		{6, 5, 4, 1, 3, 2, 9, 8, 7},
		{7, 6, 5, 2, 4, 3, 1, 9, 8},
		{4, 3, 2, 8, 1, 9, 7, 6, 5},
		{1, 9, 8, 7, 6, 5, 4, 3, 2},
	}

	incorrectGridDuplicateInLine = [][]int{
		{5, 4, 3, 9, 2, 1, 8, 7, 6},
		{2, 1, 9, 6, 8, 7, 5, 4, 3},
		{8, 7, 6, 3, 5, 4, 2, 1, 8},
		{9, 8, 7, 4, 6, 5, 3, 2, 1},
		{3, 2, 1, 7, 9, 8, 6, 5, 4},
		{6, 5, 4, 1, 3, 2, 9, 8, 7},
		{7, 6, 5, 2, 4, 3, 1, 9, 9},
		{4, 3, 2, 8, 1, 9, 7, 6, 5},
		{1, 9, 8, 5, 7, 6, 4, 3, 2},
	}

	gridMock grid
)

//DUPLICATE
func generateGridMock(input [][]int) grid {

	return grid{
		Entries: generateEntries(input),
	}
}

func TestOneLine(t *testing.T) {

	if checkCorrectness(correctOneLineOrdered) != nil {
		t.Errorf("Should not have found an error on correct line %v", correctOneLineOrdered)
	}
	if checkCorrectness(correctOneLineUnordered) != nil {
		t.Errorf("Should not have found an error on correct line %v", correctOneLineUnordered)
	}
	if checkCorrectness(correctOneLineOnlyHoles) != nil {
		t.Errorf("Should not have found an error on line with only holes %v", correctOneLineOnlyHoles)
	}
	if checkCorrectness(correctOneLineWithHoles) != nil {
		t.Errorf("Should not have found an error on line with holes %v", correctOneLineOnlyHoles)
	}

	expected := errors.New(fmt.Sprintf("Wrong length of line/column/square. Expected 9, got %d", len(incorrectOneLineTooShort)))
	err := checkCorrectness(incorrectOneLineTooShort)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported too short line error, expected %v, got %v", expected, err)
	}
	expected = errors.New(fmt.Sprintf("Duplicate 9 found in line/column/square %v", incorrectOneLineWithDuplicates))
	err = checkCorrectness(incorrectOneLineWithDuplicates)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New(fmt.Sprintf("Duplicate 4 found in line/column/square %v", incorrectOneLineWithDuplicatesAndHoles))
	err = checkCorrectness(incorrectOneLineWithDuplicatesAndHoles)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New("Wrong data present in grid: 19")
	err = checkCorrectness(incorrectOneLineWithNumberSup9)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New("Wrong data present in grid: -5")
	err = checkCorrectness(incorrectOneLineWithNumberInf0)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
}

func TestGenerateLinesFromEntries(t *testing.T) {
	grid := generateGridMock(correctGrid)

	for i, line := range correctGrid {
		if !reflect.DeepEqual(grid.getLine(i), line) {
			t.Errorf("Should have extracted the right line (%d). Expected %v, got %v", i+1, line, grid.getLine(i))
		}
	}
}

func TestGenerateColumnsFromEntries(t *testing.T) {
	grid := generateGridMock(correctGrid)

	for i, column := range correctGridAsColumns {
		if !reflect.DeepEqual(grid.getColumn(i), column) {
			t.Errorf("Should have extracted the right column (%d). Expected %v, got %v", i+1, column, grid.getColumn(i))
		}
	}
}

func TestGenerateGroupsFromEntries(t *testing.T) {
	grid := generateGridMock(correctGrid)

	for i, group := range correctGridAsGroups {
		if !reflect.DeepEqual(grid.getSquare(i), group) {
			t.Errorf("Should have extracted the right group (%d). Expected %v, got %v", i+1, group, grid.getSquare(i))
		}
	}
}

func TestGenerateGrid(t *testing.T) {
	result, err := generateGrid(correctGrid)
	if err != nil {
		t.Errorf("Should not have found error on correct grid\n%s", err)
	}

	for i, line := range correctGrid {
		if !reflect.DeepEqual(line, result.getLine(i)) {
			t.Errorf("Should have produced correct line at index %d: expected %v, got %v", i, line, result.getLine(i))
		}
	}
	for i, column := range correctGridAsColumns {
		if !reflect.DeepEqual(column, result.getColumn(i)) {
			t.Errorf("Should have produced correct column at index %d: expected %v, got %v", i, column, result.getColumn(i))
		}
	}
	for i, square := range correctGridAsGroups {
		if !reflect.DeepEqual(square, result.getSquare(i)) {
			t.Errorf("Should have produced correct square at index %d: expected %v, got %v", i, square, result.getSquare(i))
		}
	}
}

func TestFailuresOnGenerateGrid(t *testing.T) {

	_, err := generateGrid(incorrectGridWrongNumberOfLines)
	expected := errors.New(fmt.Sprintf("Wrong length of line/column/square. Expected 9, got %d", len(incorrectGridWrongNumberOfLines)))
	if err.Error() != expected.Error() {
		t.Errorf("Should have detected wrong number of line error on incorrect grid: expected %s, got %s", expected, err)
	}
	_, err = generateGrid(incorrectGridWrongNumberOfLines)
	expected = errors.New(fmt.Sprintf("Wrong length of line/column/square. Expected 9, got %d", len(incorrectGridWrongNumberOfColumns[0])))
	if err.Error() != expected.Error() {
		t.Errorf("Should have detected wrong number of line error on incorrect grid: expected %s, got %s", expected, err)
	}
	_, err = generateGrid(incorrectGridDuplicateInColumn)
	expected = errors.New("Duplicate 6 found in line/column/square")
	if strings.Contains(err.Error(), expected.Error()) {
		t.Errorf("Should have detected error on grid with duplicates in columns: expected %s, got %s", expected, err)
	}
	_, err = generateGrid(incorrectGridDuplicateInLine)
	expected = errors.New(fmt.Sprintf("Duplicate 8 found in line/column/square %v", incorrectGridDuplicateInLine[2]))
	if err.Error() != expected.Error() {
		t.Errorf("Should have detected error on grid with duplicates in line: expected %s, got %s", expected, err)
	}
}
