package faceanalytic

import "github.com/tasping/golangtestbytas/smile/entity"

type Iface interface {
	Init(text string)
	GetFace() entity.Face
}

func Init(str string) Iface {
	temp := &faceFromText{}
	temp.Init(str)
	return temp
}
