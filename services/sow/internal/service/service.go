package service

import (
	"context"
	"log/slog"

	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/calamity-m/reap/services/sow/internal/persistence"
	"github.com/google/uuid"
)

type FoodRecordService struct {
	store persistence.FoodStore
	log   *slog.Logger
}

func (s *FoodRecordService) Get(ctx context.Context, uuid uuid.UUID) (*sow.Record, error) {

	dummy := persistence.FoodRecordEntry{}

	record := createRecord(dummy)

	return record, nil
}

func (s *FoodRecordService) GetFiltered(ctx context.Context, userId uuid.UUID, record *sow.Record) ([]*sow.Record, error) {

	return nil, nil
}

func (s *FoodRecordService) Create(ctx context.Context, userId uuid.UUID, record *sow.Record) (*sow.Record, error) {

	return &sow.Record{}, nil
}

func (s *FoodRecordService) Delete(ctx context.Context, userId uuid.UUID, uuid uuid.UUID) error {

	return nil
}

func (s *FoodRecordService) Update(ctx context.Context, userId, uuid uuid.UUID, record *sow.Record) error {
	entry := createEntry(userId, record)

	s.log.DebugContext(ctx, "wip", slog.Any("entry", entry))

	return nil
}

func NewFoodRecorderService(logger *slog.Logger, store persistence.FoodStore) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger, store: store}, nil
}

func createRecord(entry persistence.FoodRecordEntry) *sow.Record {
	record := &sow.Record{
		UserId:      entry.UserId.String(),
		Name:        entry.Name,
		Description: entry.Description,
		Kj:          entry.KJ,
		Grams:       entry.Grams,
		Ml:          entry.ML,
		Calories:    kjToCals(entry.KJ),
		Oz:          gramsToOz(entry.Grams),
		FlOz:        mlToFLOz(entry.ML),
	}

	return record
}

func createEntry(userId uuid.UUID, record *sow.Record) persistence.FoodRecordEntry {

	entry := persistence.FoodRecordEntry{
		UserId:      userId,
		Name:        record.Name,
		Description: record.Description,
		KJ:          calsToKJ(record.Calories),
		ML:          flOzToML(record.FlOz),
		Grams:       ozToGrams(record.Oz),
		Created:     record.Time.AsTime(),
	}

	// Yucky imperial system. Premature optimization
	// here isn't worth it, so just on every single
	// create we'll blat over the imperial
	// with metric if provided
	if record.Kj != 0 {
		entry.KJ = record.Kj
	}
	if record.Grams != 0 {
		entry.Grams = record.Grams
	}
	if record.Ml != 0 {
		entry.ML = record.Ml
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
