package sudokusolver

import (
	"errors"
	"fmt"
)

var possibleValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

type grid struct {
	Entries []entry
}

type entry struct {
	value          int
	possibleValues []int
	indexLine      int
	indexColumn    int
}

func (g *grid) toString() string {
	s := ""
	for i := 0; i < 9; i++ {
		s += fmt.Sprintf("%v\n", g.getLine(i))
	}
	return s + "\n"
}

func (g *grid) getEntry(position int) entry {
	return g.Entries[position]
}

func newEntry(indexL, indexC, value int) entry {

	e := entry{
		value:          0,
		possibleValues: possibleValues,
		indexLine:      indexL,
		indexColumn:    indexC,
	}

	if value != 0 {
		e.value = value
		e.possibleValues = []int{value}
	}

	return e
}

func (g *grid) setEntryValue(position, value int) {
	e := g.Entries[position]
	e.value = value
	e.possibleValues = []int{value}
}

func (g *grid) removeEntryPossibleValues(position int, impossibleValues []int) {
	e := g.Entries[position]
	for _, impossible := range impossibleValues {
		removeValue(e.possibleValues, impossible)
	}
}

func (g *grid) getLine(index int) []int {

	values := make([]int, 0, 9)
	for _, entry := range g.Entries[9*index : 9*index+9] {
		values = append(values, entry.value)
	}

	return values
}

func (g *grid) getColumn(index int) []int {
	values := make([]int, 0, 9)

	for i, entry := range g.Entries {
		if i%9-index == 0 {
			values = append(values, entry.value)
		}
	}

	return values
}

func (g *grid) getSquare(index int) []int {
	selectedEntries := make([]entry, 0, 9)

	values := make([]int, 0, 9)

	offset := 0
	switch {
	case index > 2 && index <= 5:
		offset = 27 + (index-3)*3
	case index > 5:
		offset = 54 + (index-6)*3
	default:
		offset = index * 3
	}

	selectedEntries = append(selectedEntries, g.Entries[offset:offset+3]...)
	selectedEntries = append(selectedEntries, g.Entries[offset+9:offset+12]...)
	selectedEntries = append(selectedEntries, g.Entries[offset+18:offset+21]...)

	for _, e := range selectedEntries {
		values = append(values, e.value)
	}
	return values
}

func (g *grid) getEntryPossibleValues(position int) []int {
	return g.Entries[position].possibleValues
}

func generateEntries(input [][]int) []entry {

	entries := make([]entry, 0, 81)

	for indexLine, line := range input {
		for indexColumn, element := range line {
			entries = append(entries, newEntry(indexLine, indexColumn, element))
		}
	}

	return entries
}
func generateGrid(input [][]int) (grid, error) {

	g := grid{
		Entries: generateEntries(input),
	}

	var err error

	for i := 0; i < 9; i++ {
		err = checkCorrectness(g.getColumn(i))
		if err != nil {
			return grid{}, err
		}
		err = checkCorrectness(g.getLine(i))
		if err != nil {
			return grid{}, err
		}
		err = checkCorrectness(g.getSquare(i))
		if err != nil {
			return grid{}, err
		}
	}

	return g, nil
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
