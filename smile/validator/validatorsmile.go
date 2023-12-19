package validator

import "github.com/tasping/golangtestbytas/smile/entity"

type smilefacevalidator struct {
	face     entity.Face
	criteria entity.FaceCriteria
}

func (s *smilefacevalidator) Init(input_face entity.Face) {
	s.face = input_face
	s.criteria = entity.FaceCriteria{
		Eye:   []string{":", ";"},
		Nose:  []string{"-", "~"},
		Mouse: []string{")", "D"},
	}
}

func (s *smilefacevalidator) Validate() bool {

	eye := checkcontains(s.criteria.Eye, s.face.Eye, false)
	nose := checkcontains(s.criteria.Nose, s.face.Nose, true)
	mouse := checkcontains(s.criteria.Mouse, s.face.Mouse, false)

	return eye && nose && mouse
}
