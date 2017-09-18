package cards

import (
	// "fmt"
	"sort"
)

const (
	DESK_JOKER_COUNT = 2
)

type Card struct {
	sult, precedences int
}

var sults = [5]string{"heart", "diamond", "club", "spade", "joker"}

var precedences = map[int]string{
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "jack",
	12: "queen",
	13: "king",
	14: "ace",
}

type lessFunc func(p1, p2 *Card) bool

type multiSorter struct {
	changes []Card
	less    []lessFunc
}

func (ms *multiSorter) Sort(desk []Card) {
	ms.changes = desk
	sort.Sort(ms)
}

func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func FillDesk(joker_count int) []Card {
	var desk []Card
	// fill joker position in desk
	for i := 0; i < joker_count; i++ {
		desk = append(desk, Card{len(sults) - 1, 15})
	}
	// fill card position in desk
	for j := 0; j < len(precedences); j++ {
		for i := 0; i < len(sults)-1; i++ {
			desk = append(desk, Card{i, j + 6})
		}
	}
	return desk
}

func Sorting(desk *[]Card) {
	if len(*desk) == 0 {
		*desk = FillDesk(DESK_JOKER_COUNT)
	}

	sult := func(c1, c2 *Card) bool {
		return c1.sult < c2.sult
	}
	precedence := func(c1, c2 *Card) bool {
		return c1.precedences < c2.precedences
	}

	// OrderedBy(sult).Sort(desk)
	// fmt.Println("... Order by sult")
	// fmt.Println(desk)

	OrderedBy(sult, precedence).Sort(*desk)
}
