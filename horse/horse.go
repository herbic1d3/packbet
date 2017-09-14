package horse

import (
	"fmt"
)

type Field struct {
	Letter string
	Number int
}

var letters = map[int]string{
	1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h",
}

func number_by_letter(l string) int {
	for key, value := range letters {
		if value == l {
			return key
		}
	}

	return -1
}

func running(long int, short int, swap bool, result *[]Field) {
	var ltmp, stmp int

	for i := -2; i <= 2; i = i + 4 {
		ltmp = long + i
		if ltmp > 0 && ltmp < 9 {
			for j := -1; j <= 1; j = j + 2 {
				stmp = short + j

				if stmp > 0 && stmp < 9 {
					if swap {
						*result = append(*result, Field{letters[stmp], ltmp})
					} else {
						*result = append(*result, Field{letters[ltmp], stmp})
					}
				}
			}
		}
	}
}

// Slice with horse movens
func Movements(f *Field) []Field {
	var result []Field

	if f.Number > 8 || f.Number < 1 || number_by_letter(f.Letter) < 0 {
		panic(fmt.Sprintf("Wrong board position:  %v", *f))
	}

	running(number_by_letter(f.Letter), f.Number, false, &result)
	running(f.Number, number_by_letter(f.Letter), true, &result)

	return result
}
