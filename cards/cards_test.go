package cards

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	//
	// test FillDesk
	desk := FillDesk(2)
	// check len desk
	if len(desk) != 38 {
		t.Error(fmt.Sprintf("Wrong cards count: %v", len(desk)))
	}
	// check jocker count
	joker_count := 0
	for _, value := range desk {
		if value.precedences == 15 {
			joker_count++
		}
	}
	if joker_count != 2 {
		t.Error(fmt.Sprintf("Wrong joker counts: %v", joker_count))
	}
	//
	// test Sorting
	var desk_reference = []Card{
		{0, 6}, {0, 7}, {0, 8}, {0, 9}, {0, 10}, {0, 11}, {0, 12}, {0, 13}, {0, 14},
		{1, 6}, {1, 7}, {1, 8}, {1, 9}, {1, 10}, {1, 11}, {1, 12}, {1, 13}, {1, 14},
		{2, 6}, {2, 7}, {2, 8}, {2, 9}, {2, 10}, {2, 11}, {2, 12}, {2, 13}, {2, 14},
		{3, 6}, {3, 7}, {3, 8}, {3, 9}, {3, 10}, {3, 11}, {3, 12}, {3, 13}, {3, 14},
		{4, 15}, {4, 15},
	}
	Sorting(&desk)
	if !reflect.DeepEqual(desk, desk_reference) {
		t.Error("Wrong desk sorting")
	}
}
