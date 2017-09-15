package cards

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DESK_JOKER_COUNT = 2
	DESK_CARDS_COUNT = 4*9 + DESK_JOKER_COUNT
)

type card struct {
	sult, precedences, position int
}

var desk [DESK_CARDS_COUNT]card

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

func exists_value(arr *[]int, value int) bool {
	for _, val := range *arr {
		if val == value {
			return true
		}
	}
	*arr = append(*arr, value)
	return false
}

func randimize(args ...int) (result int) {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(args) > 1 {
		result = rand.Intn((args[0]+1)-args[1]) + args[1]
	} else {
		result = rand.Intn(args[0])
	}
	return result
}

func fill_desk() {
	var (
		use_position []int
		pos          int
	)
	// fill card position in desk
	cnt := 0
	for j := 0; j < len(precedences); j++ {
		for i := 0; i < len(sults)-1; i++ {
			// fmt.Println(use_position)
			// fmt.Println(desk)
			pos = randimize(DESK_CARDS_COUNT)
			for exists_value(&use_position, pos) {
				pos = randimize(DESK_CARDS_COUNT)
			}

			desk[cnt].sult = i
			desk[cnt].precedences = j + 6
			desk[cnt].position = pos

			cnt++
		}
	}
	// fill joker position in desk
	for i := 0; i < DESK_CARDS_COUNT; i++ {
		if !exists_value(&use_position, i) {
			desk[cnt].sult = len(sults) - 1
			desk[cnt].precedences = 15
			desk[cnt].position = i
			cnt++
		}
	}
}

func print_desk() {
	var line []card
	for i := 0; i < len(sults); i++ {
		line = line[:0]
		for j := 0; j < DESK_CARDS_COUNT; j++ {
			if desk[j].sult == i {
				line = append(line, desk[j])
			}
		}
		fmt.Println(line)
	}
}

func Sorting() {
	fill_desk()
	print_desk()

	joker_pos := DESK_CARDS_COUNT - DESK_JOKER_COUNT
	for idx := 0; idx < DESK_CARDS_COUNT; idx++ {
		if desk[idx].sult == len(sults)-1 {
			desk[idx].position = joker_pos
			joker_pos++
		} else {
			desk[idx].position = desk[idx].sult*len(precedences) + (desk[idx].precedences - 6)
		}
	}
	fmt.Println("...")
	print_desk()
}
