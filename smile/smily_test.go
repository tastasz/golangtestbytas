package smile_test

import (
	"testing"

	"github.com/tasping/golangtestbytas/smile"
)

// Write your test here
func TestSmileCase1(t *testing.T) {
	input := []string{":)", ";(", ";}", ":-D"}
	expected := 2
	got := smile.CountSmile(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestSmileCase2(t *testing.T) {
	input := []string{";D", ":-(", ":-)", ";~)"}
	expected := 3
	got := smile.CountSmile(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}

func TestSmileCase3(t *testing.T) {
	input := []string{";]", ":[", ";*", ":$", ";-D"}
	expected := 1
	got := smile.CountSmile(input)
	if expected != got {
		t.Error("expected", expected, "but got", got)
	}
}
