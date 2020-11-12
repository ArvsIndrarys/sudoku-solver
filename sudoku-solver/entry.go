package sudokusolver

type entry struct {
	value         int
	possibilities []int
	indexLine     int
	indexColumn   int
	indexSquare   int
}

func newEntry(position, value int) entry {

	e := entry{
		value:         0,
		possibilities: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		indexLine:     position / 9,
		indexColumn:   position % 9,
		indexSquare:   (position/27)*3 + (position/3)%3,
	}

	if value != 0 {
		e.value = value
		e.possibilities = []int{}
	}

	return e
}

func (e *entry) setValue(value int) {
	e.value = value
	e.possibilities = []int{}
}

func (e *entry) removePossibility(p int) {

	e.possibilities = removeValue(e.possibilities, p)

	// if only one possibility, set it
	if len(e.possibilities) == 1 {
		e.value = e.possibilities[0]
		e.possibilities = []int{}
	}
}
