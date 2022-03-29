package proseList

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

//helper test function
func errorString(function string, list []string, want string, got string) string {
	return fmt.Sprintf("%s(%#v) Expected %v, got %v", function, list, want, got)
}

func TestJoinWithCommas(t *testing.T) {
	test := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
		testData{list: []string{"apple", "orange", "banana", "pear"}, want: "apple, orange, banana, and pear"},
	}
	for _, test := range test {
		got := JoinWithcommasTwo(test.list)
		if got != test.want {
			t.Error(errorString("JoinWithCommas", test.list, test.want, got))
		}
	}
}
