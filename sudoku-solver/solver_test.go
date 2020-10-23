package sudokusolver

import (
	"errors"
	"fmt"
	"reflect"
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
)

func TestOneLine(t *testing.T) {

	if checkLineCorrectness(correctOneLineOrdered) != nil {
		t.Errorf("Should not have found an error on correct line %v", correctOneLineOrdered)
	}
	if checkLineCorrectness(correctOneLineUnordered) != nil {
		t.Errorf("Should not have found an error on correct line %v", correctOneLineUnordered)
	}
	if checkLineCorrectness(correctOneLineOnlyHoles) != nil {
		t.Errorf("Should not have found an error on line with only holes %v", correctOneLineOnlyHoles)
	}
	if checkLineCorrectness(correctOneLineWithHoles) != nil {
		t.Errorf("Should not have found an error on line with holes %v", correctOneLineOnlyHoles)
	}

	expected := errors.New(fmt.Sprintf("Wrong length of line/column. Expected 9, got %d", len(incorrectOneLineTooShort)))
	err := checkLineCorrectness(incorrectOneLineTooShort)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported too short line error, expected %v, got %v", expected, err)
	}
	expected = errors.New(fmt.Sprintf("Duplicate 9 found in line %v", incorrectOneLineWithDuplicates))
	err = checkLineCorrectness(incorrectOneLineWithDuplicates)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New(fmt.Sprintf("Duplicate 4 found in line %v", incorrectOneLineWithDuplicatesAndHoles))
	err = checkLineCorrectness(incorrectOneLineWithDuplicatesAndHoles)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New("Wrong data present in grid: 19")
	err = checkLineCorrectness(incorrectOneLineWithNumberSup9)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
	expected = errors.New("Wrong data present in grid: -5")
	err = checkLineCorrectness(incorrectOneLineWithNumberInf0)
	if err.Error() != expected.Error() {
		t.Errorf("Should have reported duplicates in line error, expected %v, got %v", expected, err)
	}
}

func TestGenerateLines(t *testing.T) {
	result, err := generateLines(correctGrid)
	if err != nil {
		t.Errorf("Should not have found error on correct grid %s", err)
	}
	if result == nil {
		t.Errorf("Should not have returned nil on correct grid !")
	}

	for i, v := range result {
		if !reflect.DeepEqual(v, correctGrid[i]) {
			t.Errorf("Line %d is not what expected, expected\n%v\ngot\n%v", i, v, correctGrid[i])
		}
	}
}

func TestGenerateColumns(t *testing.T) {
	result, err := generateColumns(correctGrid)
	if err != nil {
		t.Errorf("Should not have found error on correct grid %s", err)
	}
	if result == nil {
		t.Errorf("Should not have returned nil on correct grid !")
	}

	for i, v := range result {
		if !reflect.DeepEqual(v, correctGridAsColumns[i]) {
			t.Errorf("Line %d is not what expected, expected\n%v\ngot\n%v", i, v, correctGridAsColumns[i])
		}
	}
}

func TestGenerateGroups(t *testing.T) {
	result, err := generateGroups(correctGrid)
	if err != nil {
		t.Errorf("Should not have found error on correct grid\n%s", err)
	}
	if result == nil {
		t.Errorf("Should not have returned nil on correct grid !")
	}

	for i, v := range result {
		if !reflect.DeepEqual(v, correctGridAsGroups[i]) {
			t.Errorf("Line %d is not what expected, expected\n%v\ngot\n%v", i, v, correctGridAsGroups[i])
		}
	}
}

func TestGenerateGrid(t *testing.T) {
	result, err := generateGrid(correctGrid)
	if err != nil {
		t.Errorf("Should not have found error on correct grid\n%s", err)
	}

	for i, v := range result.Lines {
		if !reflect.DeepEqual(v, correctGrid[i]) {
			t.Errorf("Lines produced are not the one expected:\n%v\ngot\n%v", Array2DToString(correctGrid), Map2DToString(result.Lines))
		}
	}
	for i, v := range result.Columns {
		if !reflect.DeepEqual(v, correctGridAsColumns[i]) {
			t.Errorf("Columns produced are not the one expected:\n%v\ngot\n%v", Array2DToString(correctGridAsColumns), Map2DToString(result.Columns))
		}
	}
	for i, v := range result.Groups {
		if !reflect.DeepEqual(v, correctGridAsGroups[i]) {
			t.Errorf("Groups produced are not the one expected:\n%v\ngot\n%v", Array2DToString(correctGridAsGroups), Map2DToString(result.Groups))
		}
	}
}
