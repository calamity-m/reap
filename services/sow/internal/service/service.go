package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/calamity-m/reap/services/sow/internal/persistence"
	"github.com/google/uuid"
)

type FoodRecord struct {
	Uuid        uuid.UUID `json:"id"`
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

type FoodRecordService struct {
	log *slog.Logger
}

func (s *FoodRecordService) Get(ctx context.Context, userId uuid.UUID, uuid uuid.UUID) (FoodRecord, error) {
	s.log.DebugContext(
		ctx,
		"attempting to get food record",
		slog.String("userid", userId.String()),
		slog.String("uuid", uuid.String()),
	)

	return FoodRecord{}, nil
}

func (s *FoodRecordService) GetFiltered(ctx context.Context, userId uuid.UUID, fr FoodRecord) ([]FoodRecord, error) {
	s.log.DebugContext(
		ctx,
		"attempting to get food records based on a filter",
		slog.Any("filter", fr),
	)

	return []FoodRecord{}, nil
}

func (s *FoodRecordService) Create(ctx context.Context, userId uuid.UUID, fr FoodRecord) (FoodRecord, error) {
	s.log.DebugContext(
		ctx,
		"attempting to create food record",
		slog.Any("record", fr),
	)

	return FoodRecord{}, nil
}

func (s *FoodRecordService) Delete(ctx context.Context, userId uuid.UUID, uuid uuid.UUID) error {
	s.log.DebugContext(
		ctx,
		"attempting to get delete food record",
		slog.String("userid", userId.String()),
		slog.String("uuid", uuid.String()),
	)

	return nil
}

func (s *FoodRecordService) Update(ctx context.Context, userId, uuid uuid.UUID, fr FoodRecord) error {
	s.log.DebugContext(
		ctx,
		"attempting to update specific food record",
		slog.Any("record", fr),
		slog.String("uuid", uuid.String()),
	)

	return nil
}

func transformRecord(userId uuid.UUID, fr FoodRecord) persistence.FoodRecordEntry {
	return persistence.FoodRecordEntry{
		UserId:      userId,
		Uuid:        fr.Uuid,
		Name:        fr.Name,
		Description: fr.Description,
		KJ:          fr.KJ,
		ML:          fr.ML,
		Gram:        fr.Gram,
		Created:     fr.Created,
	}
}

func NewFoodRecorderService(logger *slog.Logger) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger}, nil
}
