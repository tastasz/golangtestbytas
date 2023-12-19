package faceanalytic

import "github.com/tasping/golangtestbytas/smile/entity"

type faceFromText struct {
	face entity.Face
}

func (f *faceFromText) Init(face string) {
	if len(face) == 3 {
		f.face.Eye = face[:1]
		f.face.Nose = face[1:2]
		f.face.Mouse = face[2:]
	} else {
		f.face.Eye = face[:1]
		f.face.Mouse = face[1:2]
	}
}

func (f *faceFromText) GetFace() entity.Face {
	return f.face
}
