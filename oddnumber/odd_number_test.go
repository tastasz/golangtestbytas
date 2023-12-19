package oddnumber_test

import (
	"testing"

	"github.com/tasping/golangtestbytas/oddnumber"
)

func TestOddCase1(t *testing.T) {

	input := []int{7}
	expected := 7

	got := oddnumber.FindOddNumber(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestOddCase2(t *testing.T) {

	input := []int{0}
	expected := 0

	got := oddnumber.FindOddNumber(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestOddCase3(t *testing.T) {

	input := []int{1, 1, 2}
	expected := 2

	got := oddnumber.FindOddNumber(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestOddCase4(t *testing.T) {

	input := []int{0, 1, 0, 1, 0, 1}
	expected := 0

	got := oddnumber.FindOddNumber(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestOddCase5(t *testing.T) {

	input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	expected := 4

	got := oddnumber.FindOddNumber(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}
