package sudokusolver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type grid struct {
	Entries []entry
}

// Grid Generation

func generateBaseGrid() grid {
	entries := make([]entry, 0, 81)
	for i := 0; i < 81; i++ {
		entries = append(entries, newEntry(i, 0))
	}

	return grid{
		Entries: entries,
	}
}

func generateGridFromString(input string) (grid, error) {

	// input check
	if !strings.Contains(input, ",") {
		return grid{}, errors.New("Missing ',' separator between each value")
	}

	inputArray := strings.Split(input, ",")

	if len(inputArray) != 81 {
		return grid{}, errors.New(fmt.Sprintf("Wrong input, did not get right amount of input in %v", inputArray))
	}

	// we generate a base grid
	// is has only values to find and all possibilities on each entry
	g := generateBaseGrid()

	// we update the grid according to the input
	for i, v := range inputArray {

		value, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return grid{}, errors.New(fmt.Sprintf("Wrong input, could not convert element at index %d as a number : (%s)", i, v))
		}

		g.updateValue(i, value)
	}

	// check correctness of result
	if err := g.checkCorrectness(); err != nil {
		return grid{}, errors.New(fmt.Sprintf("Input grid seems wrong : %s", err))
	}

	return g, nil
}

func generateGrid(input [][]int) (grid, error) {

	if len(input) != 9 {
		return grid{}, errors.New(fmt.Sprintf("Wrong number of lines on input. Expected 9, got %d", len(input)))
	}
	for i, line := range input {
		if len(line) != 9 {
			return grid{}, errors.New(fmt.Sprintf("Wrong number of columns at line %d on input. Expected 9, got %d", i, len(line)))
		}
	}
	g := generateBaseGrid()

	for indexLine, line := range input {
		for indexColumn, value := range line {
			g.updateValue(indexLine*9+indexColumn, value)
		}
	}

	if err := g.checkCorrectness(); err != nil {
		return grid{}, err
	}

	return g, nil
}

// Grid representation
func (g *grid) String() string {
	s := ""
	for i := 0; i < 9; i++ {

		if i%3 == 0 {
			s += "|-----------------------------|\n"
		}

		for j, value := range g.getLine(i) {
			if j%3 == 0 {
				s += "|"
			}
			s += fmt.Sprintf(" %d ", value)
		}
		s += "|\n"
	}

	s += "|-----------------------------|"
	return s
}

// grid modification
func (g *grid) getEntry(position int) entry {
	return g.Entries[position]
}

func (g *grid) eGetLine(index int) []entry {
	return g.Entries[9*index : 9*index+9]
}

func (g *grid) eGetColumn(index int) []entry {
	values := make([]entry, 0, 9)

	for i, entry := range g.Entries {
		if i%9-index == 0 {
			values = append(values, entry)
		}
	}

	return values
}

func (g *grid) eGetSquare(index int) []entry {
	selectedEntries := make([]entry, 0, 9)

	values := make([]entry, 0, 9)

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
		values = append(values, e)
	}
	return values
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

	selectedEntries := g.eGetSquare(index)
	values := make([]int, 0, 9)

	for _, e := range selectedEntries {
		values = append(values, e.value)
	}
	return values
}

func (g *grid) checkCorrectness() error {

	var err error
	for i := 0; i < 9; i++ {
		err = checkCorrectness(g.getColumn(i))
		if err != nil {
			return err
		}
		err = checkCorrectness(g.getLine(i))
		if err != nil {
			return err
		}
		err = checkCorrectness(g.getSquare(i))
		if err != nil {
			return err
		}
	}

	return err
}

func (g *grid) updateValue(position, value int) {

	e := g.getEntry(position)

	e.setValue(value)
	removePossibility(g.getLineOfElement(e), value)
	removePossibility(g.getColumnOfElement(e), value)
	removePossibility(g.getSquareOfElement(e), value)

	g.Entries[position] = e
}

func removePossibility(entries []entry, possibility int) {
	for _, e := range entries {
		e.removePossibility(possibility)
	}
}

func (g *grid) getLineOfElement(e entry) []entry {
	return g.eGetLine(e.indexLine)
}

func (g *grid) getColumnOfElement(e entry) []entry {
	return g.eGetColumn(e.indexColumn)
}

func (g *grid) getSquareOfElement(e entry) []entry {
	return g.eGetSquare(e.indexSquare)
}
