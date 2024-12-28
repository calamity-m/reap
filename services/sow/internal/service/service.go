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
	// Default used for storage of energy, if kj and calories are provided
	// KJ will always be selected
	KJ float32 `json:"kj,omitempty"`
	// Default used for storage of mass, if Gram and Oz are provided
	// Gram will always be selected
	Grams float32 `json:"grams,omitempty"`
	// Default used for storage of volume, if ML and FLOz are provided
	// ML will always be selected
	ML       float32   `json:"ml,omitempty"`
	Calories float32   `json:"calories,omitempty"`
	Oz       float32   `json:"oz,omitempty"`
	FLOz     float32   `json:"floz,omitempty"`
	Created  time.Time `json:"created"`
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

	dummy := persistence.FoodRecordEntry{}

	record := createRecord(dummy)

	return record, nil
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

	entry := createEntry(userId, fr)

	s.log.DebugContext(ctx, "wip", slog.Any("entry", entry))

	return nil
}

func NewFoodRecorderService(logger *slog.Logger) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger}, nil
}

func createRecord(entry persistence.FoodRecordEntry) FoodRecord {
	record := FoodRecord{
		Uuid:        entry.Uuid,
		Name:        entry.Name,
		Description: entry.Description,
		KJ:          entry.KJ,
		Grams:       entry.Grams,
		ML:          entry.ML,
		Calories:    kjToCals(entry.KJ),
		Oz:          gramsToOz(entry.Grams),
		FLOz:        mlToFLOz(entry.ML),
	}

	return record
}

func createEntry(userId uuid.UUID, fr FoodRecord) persistence.FoodRecordEntry {

	entry := persistence.FoodRecordEntry{
		UserId:      userId,
		Uuid:        fr.Uuid,
		Name:        fr.Name,
		Description: fr.Description,
		KJ:          calsToKJ(fr.Calories),
		ML:          flOzToML(fr.FLOz),
		Grams:       ozToGrams(fr.Oz),
		Created:     fr.Created,
	}

	// Yucky imperial system. Premature optimization
	// here isn't worth it, so just on every single
	// create we'll blat over the imperial
	// with metric if provided
	if fr.KJ != 0 {
		entry.KJ = fr.KJ
	}
	if fr.Grams != 0 {
		entry.Grams = fr.Grams
	}
	if fr.ML != 0 {
		entry.ML = fr.ML
	}

	return entry
}

func calsToKJ(cals float32) float32 {
	return cals * 4.184
}

func ozToGrams(oz float32) float32 {
	return oz * 28.35
}

func flOzToML(flOz float32) float32 {
	return flOz * 29.574
}

func kjToCals(kj float32) float32 {
	return kj / 4.184
}

func gramsToOz(grams float32) float32 {
	return grams / 28.35
}

func mlToFLOz(ml float32) float32 {
	return ml / 29.574
}
