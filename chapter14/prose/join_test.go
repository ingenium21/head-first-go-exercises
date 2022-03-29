package prose

import (
	"fmt"
	"testing"
)

//helper test function
func errorString(function string, list []string, want string, got string) string {
	return fmt.Sprintf("%s(%#v) Expected %v, got %v", function, list, want, got)
}

func TestNoElement(t *testing.T) {
	list := []string{}
	if JoinWithcommas(list) != "" {
		t.Error(errorString("JoinWithcommas", list, "", JoinWithcommas(list)))
	}
}
func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	if JoinWithcommas(list) != "apple" {
		t.Error(errorString("JoinWithcommas", list, "apple", JoinWithcommas(list)))
	}

}
func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithcommas(list) != "apple and orange" {
		t.Error(errorString("JoinWithcommas", list, "apple and orange", JoinWithcommas(list)))
	}
}
func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "banana"}
	if JoinWithcommas(list) != "apple, orange, and banana" {
		t.Error(errorString("JoinWithcommas", list, "apple, orange, and banana", JoinWithcommas(list)))
	}
}
func TestFourElements(t *testing.T) {
	list := []string{"apple", "orange", "banana", "pear"}
	if JoinWithcommas(list) != "apple, orange, banana, and pear" {
		t.Error(errorString("JoinWithcommas", list, "apple, orange, banana, and pear", JoinWithcommas(list)))
	}
}
