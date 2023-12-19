package oddnumber

import "fmt"

func FindOddNumber(input []int) int {
	map_counting := make(map[int]int)

	for _, item := range input {
		map_counting[item]++
	}

	var result int
	var occur_value int
	for key, value := range map_counting {
		if value%2 != 0 {
			result = key
			occur_value = value
			break
		}
	}

	fmt.Printf("%v should return %d, because it occurs %d time (which is odd).", input, result, occur_value)

	return result
}
