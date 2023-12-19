package manipulate_test

import (
	"reflect"
	"testing"

	"github.com/tasping/golangtestbytas/manipulate"
)

func TestInput_a(t *testing.T) {
	input := "a"
	expected := []string{"a"}
	got := manipulate.Manipulate(input)
	if !compareSliceIgnoreOrder(expected, got) {
		t.Error("expected", expected, "but got", got)
	}
}

func TestInput_ab(t *testing.T) {
	input := "ab"
	expected := []string{"ab", "ba"}
	got := manipulate.Manipulate(input)
	if !compareSliceIgnoreOrder(expected, got) {
		t.Error("expected", expected, "but got", got)
	}
}

func TestInput_abc(t *testing.T) {
	input := "abc"
	expected := []string{"abc", "acb", "bac", "bca", "cba", "cab"}
	got := manipulate.Manipulate(input)
	if !compareSliceIgnoreOrder(expected, got) {
		t.Error("expected", expected, "but got", got)
	}
}
func TestInput_aabb(t *testing.T) {
	input := "aabb"
	expected := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}
	got := manipulate.Manipulate(input)
	if !compareSliceIgnoreOrder(expected, got) {
		t.Error("expected", expected, "but got", got)
	}
}

func compareSliceIgnoreOrder(slice1, slice2 []string) bool {
	// Create a map to count occurrences of each string in both slices
	counts1 := make(map[string]int)
	counts2 := make(map[string]int)

	for _, str := range slice1 {
		counts1[str]++
	}

	// Count occurrences in the second slice
	for _, str := range slice2 {
		counts2[str]++
	}

	// Use DeepEqual to compare the two maps
	return reflect.DeepEqual(counts1, counts2)

}
