package service

import (
	"context"
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
	// If we have an ID set on our create, exit early as that is an invalid attempt at creation
	if record.Id != "" {
		return nil, fmt.Errorf("id cannot be set: %w", errs.ErrInvalidRequest)
	}

	// We need to have some form of description, no matter what it is.
	if record.Description == "" {
		return nil, fmt.Errorf("description must be set: %w", errs.ErrInvalidRequest)
	}

	// Ensure we have a valid user id before updating
	userId, err := uuid.Parse(record.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("userid %s is not a valid uuid: %w", record.Id, errs.ErrInvalidRequest)
	}

	// Perform mapping of the record for storage layer, ignoring any of the uuids we have set
	// on the record.
	entry, err := MapRecordToEntryWithoutUuids(record)
	if err != nil {
		return nil, err
	}

	// Set uuids ourselveves
	entry.UserId = userId

	// Generate a UUID id
	if entry.Id == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
			s.log.ErrorContext(ctx, "failed to generate uuid for create", slog.Any("err", err))
			return nil, fmt.Errorf("failed to generate id: %w", errs.ErrInternal)
		}

		entry.Id = id
	}

	// Attempt to create it
	err = s.store.CreateFood(entry)
	if err != nil {
		return nil, err
	}

	// Grab the newly created entry
	created, err := s.Get(ctx, entry.Id)

	if err != nil {
		return nil, fmt.Errorf("failed grab record after it was created: %w", errs.ErrInternal)
	}

	return created, nil
}

func (s *FoodRecordService) Delete(ctx context.Context, id uuid.UUID) error {

	return nil
}

func (s *FoodRecordService) Update(ctx context.Context, record *sow.Record) error {
	return nil
}

func NewFoodRecorderService(logger *slog.Logger, store persistence.FoodStore) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger, store: store}, nil
}
