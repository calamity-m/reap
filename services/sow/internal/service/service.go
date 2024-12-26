package service

import (
	"time"

	"github.com/google/uuid"
)

type FoodRecord struct {
	Uuid        uuid.UUID `json:"id"`
	UserUuid    uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	KJ          float32   `json:"kj,omitempty"`
	Gram        float32   `json:"grams,omitempty"`
	ML          float32   `json:"ml,omitempty"`
	Created     time.Time `json:"created"`
}

func (r FoodRecord) ConvertCalories() float32 {
	return r.KJ / 4.184
}

func (r FoodRecord) ConvertPound() float32 {
	return r.Gram / 453.6
}

func (r FoodRecord) ConvertOunce() float32 {
	return r.Gram / 28.35
}

type FoodRecordService struct{}

func (s *FoodRecordService) Get(uuid.UUID) (FoodRecord, error) {
	return FoodRecord{}, nil
}

func (s *FoodRecordService) GetFiltered(FoodRecord) (FoodRecord, error) {
	return FoodRecord{}, nil
}

func (s *FoodRecordService) Create(FoodRecord) (FoodRecord, error) {
	return FoodRecord{}, nil
}

func (s *FoodRecordService) Delete(uuid.UUID) error {
	return nil
}

func (s *FoodRecordService) Update(uuid.UUID) error {
	return nil
}

func NewFoodRecorderService() (*FoodRecordService, error) {
	return &FoodRecordService{}, nil
}
