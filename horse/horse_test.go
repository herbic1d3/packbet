package horse

import (
	"fmt"
	"reflect"
	"testing"
)

type testvalues struct {
	value  Field
	result []Field
}

func TestMovements(t *testing.T) {
	var result []Field

	testparams := make([]testvalues, 3)
	testparams[0].value = Field{"h", 1}
	testparams[0].result = append(testparams[0].result, Field{"f", 2})
	testparams[0].result = append(testparams[0].result, Field{"g", 3})
	testparams[1].value = Field{"h", 8}
	testparams[1].result = append(testparams[1].result, Field{"f", 7})
	testparams[1].result = append(testparams[1].result, Field{"g", 6})
	testparams[2].value = Field{"a", 1}
	testparams[2].result = append(testparams[2].result, Field{"c", 2})
	testparams[2].result = append(testparams[2].result, Field{"b", 3})

	for _, param := range testparams {
		result = Movements(&param.value)
		if !reflect.DeepEqual(result, param.result) {
			t.Error(fmt.Sprintf("For field %v wrong result %v", param.value, result))
		}
	}
}
