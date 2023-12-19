package manipulate

func Manipulate(text string) []string {
	result := permutationFunc(text)
	result_unique := removeDuplicate(result)
	return result_unique
}

func permutationFunc(current_input string) []string {
	if len(current_input) == 1 {
		return []string{current_input}
	}
	result := []string{}
	for i, constant_char := range current_input {
		//remove char at "constant_char" from current string
		remaining := current_input[0:i] + current_input[i+1:]
		//recursive function
		subPermutations := permutationFunc(remaining)
		for _, permu_char := range subPermutations {
			result = append(result, string(constant_char)+permu_char)
		}
	}

	return result
}

func removeDuplicate(s []string) []string {
	map_string := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := map_string[str]; !ok {
			map_string[str] = true
			result = append(result, str)
		}
	}
	return result
}
