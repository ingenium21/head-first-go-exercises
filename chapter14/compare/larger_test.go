package compare

import (
	"fmt"
	"testing"
)

//helper test string
func errorString(function string, a int, b int, want int, got int) string {
	return fmt.Sprintf("%s(%d, %d) Expected %d, got %d", function, a, b, want, got)
}
func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Errorf(errorString("Larger", 2, 1, want, got))
	}
}

func TestSecondLarger(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if got != want {
		t.Errorf(errorString("Larger", 4, 8, want, got))
	}
}
