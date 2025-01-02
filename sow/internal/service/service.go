package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/calamity-m/reap/pkg/errs"
	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/calamity-m/reap/sow/internal/persistence"
	"github.com/google/uuid"
)

type FoodRecordService struct {
	store persistence.FoodStore
	log   *slog.Logger
}

func (s *FoodRecordService) Get(ctx context.Context, id uuid.UUID) (*sow.Record, error) {

	entry, err := s.store.GetFood(id)

	if err != nil {
		return nil, err
	}

	found := MapEntryToRecord(entry)

	return found, nil
}

func (s *FoodRecordService) GetFiltered(ctx context.Context, record *sow.Record) ([]*sow.Record, error) {

	return nil, nil
}

func (s *FoodRecordService) Create(ctx context.Context, record *sow.Record) (*sow.Record, error) {

	// Perform mapping of the record for storage layer
	entry, err := MapRecordToEntry(record)
	if err != nil {
		return nil, err
	}

	// Generate a UUID id if not provided
	if entry.Uuid == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			s.log.ErrorContext(ctx, "failed to generate uuid for create", slog.Any("err", err))
			return nil, errors.Join(fmt.Errorf("failed to generate record id"), errs.ErrInternal)
		}

		entry.Uuid = id
	}

	// Attempt to create it
	err = s.store.CreateFood(entry)
	if err != nil {
		return nil, err
	}

	// Grab the newly created entry
	created, err := s.Get(ctx, entry.Uuid)

	if err != nil {
		return nil, fmt.Errorf("failed grab record after it was created: %w", errs.ErrInternal)
	}

	return created, nil
}

func (s *FoodRecordService) Delete(ctx context.Context, id uuid.UUID) error {

	return nil
}

func (s *FoodRecordService) Update(ctx context.Context, id uuid.UUID, record *sow.Record) error {
	entry, err := MapRecordToEntry(record)

	if err != nil {
		return err
	}

	s.log.DebugContext(ctx, "wip", slog.Any("entry", entry))

	return nil
}

func NewFoodRecorderService(logger *slog.Logger, store persistence.FoodStore) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger, store: store}, nil
}
