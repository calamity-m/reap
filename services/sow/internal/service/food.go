package service

import "github.com/google/uuid"

type Record struct {
	Uuid        uuid.UUID
	UserUuid    uuid.UUID
	Name        string
	Description string
	KJ          float32
	Gram        float32
	ML          float32
}

func (r Record) ConvertCalories() float32 {
	return r.KJ / 4.184
}

func (r Record) ConvertPound() float32 {
	return r.Gram / 453.6
}

func (r Record) ConvertOunce() float32 {
	return r.Gram / 28.35
}
