package smile

import (
	"github.com/tasping/golangtestbytas/smile/faceanalytic"
	"github.com/tasping/golangtestbytas/smile/validator"
)

func CountSmile(input []string) int {
	var count_smiley int
	for _, value := range input {
		faceObj := faceanalytic.Init(value).GetFace()
		validator := validator.Init(faceObj).Validate()
		if validator {
			count_smiley++
		}
	}
	return count_smiley
}
