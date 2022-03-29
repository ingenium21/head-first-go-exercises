package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1+2 did not equal 3, got: ", Add(1, 2))
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(8, 4) != 4 {
		t.Error("3-2 did not equal 1, got: ", Subtract(3, 2))
	}
}
