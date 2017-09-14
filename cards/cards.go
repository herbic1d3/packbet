package cards

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MIN_CARD_RATING = 6
	MAX_CARD_RATING = 14
	CARDS_IN_SUIT   = 9
)

var sults = [4]string{"heart", "diamond", "club", "spade"}
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
var desk [4][9]int

func exists_value(arr []int, value int) bool {
	for _, val := range arr {
		if val == value {
			return true
		}
	}

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
		in_use_sult []int
		in_use_card []int
		sult        int
		card        int
	)
	// fill randomize sult
	for i := 0; i < len(sults); i++ {
		sult = randimize(len(sults))

		for exists_value(in_use_sult, sult) {
			sult = randimize(len(sults))
		}
		in_use_sult = append(in_use_sult, sult)

		// fill randomize cards
		in_use_card = in_use_card[:0]
		for j := 0; j < CARDS_IN_SUIT; j++ {
			card = randimize(MAX_CARD_RATING, MIN_CARD_RATING)

			for exists_value(in_use_card, card) {
				card = randimize(MAX_CARD_RATING, MIN_CARD_RATING)
			}

			in_use_card = append(in_use_card, card)
			desk[sult][j] = card
		}
	}
}

func Sorting() {
	var (
		left, right, i int
	)

	fill_desk()
	fmt.Println(desk)

	for sult := 0; sult < len(sults); sult++ {
		left = 0
		right = len(desk[sult]) - 1

		for left <= right {
			for i = left; i < right; i++ {
				if desk[sult][i] > desk[sult][i+1] {
					desk[sult][i], desk[sult][i+1] = desk[sult][i+1], desk[sult][i]
				}
			}

			right--

			for i = right; i > left; i-- {
				if desk[sult][i-1] > desk[sult][i] {
					desk[sult][i], desk[sult][i-1] = desk[sult][i-1], desk[sult][i]
				}
			}

			left++
		}
	}
	fmt.Println(desk)
}
