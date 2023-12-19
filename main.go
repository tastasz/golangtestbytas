package main

import (
	"fmt"

	"github.com/tasping/golangtestbytas/manipulate"
	"github.com/tasping/golangtestbytas/oddnumber"
	"github.com/tasping/golangtestbytas/smile"
)

func main() {
	// Permutations Test
	str := "aabb"
	permutations := manipulate.Manipulate(str)
	fmt.Printf("Permutations : %#v\n", permutations)
	// Find Odd Number Test
	temp := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	oddnumber.FindOddNumber(temp)
	// Smile Count Test
	input := []string{":)", ";(", ";}", ":-D"}
	fmt.Println(smile.CountSmile(input))

}
