package validator

import "github.com/tasping/golangtestbytas/smile/entity"

type IValidator interface {
	Init(input_face entity.Face)
	Validate() bool
}

func Init(obj entity.Face) IValidator {
	temp := &smilefacevalidator{}
	temp.Init(obj)
	return temp
}

func checkcontains(condition []string, criteria string, optional bool) bool {
	if optional && criteria == "" {
		return true
	}
	for _, v := range condition {
		if v == criteria {
			return true
		}
	}
	return false
}
