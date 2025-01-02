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
	s.log.DebugContext(ctx, "attempting retieval of specific record", slog.String("id", id.String()))

	// Get our record
	entry, err := s.store.GetFood(id)

	// Propogate any errors
	if err != nil {
		return nil, err
	}

	// Transform it into the domain sow struct
	found := MapEntryToRecord(entry)

	s.log.DebugContext(ctx, "successfully retrieved record", slog.Any("record", found))
	return found, nil
}

func (s *FoodRecordService) GetFiltered(ctx context.Context, record *sow.Record) ([]*sow.Record, error) {
	s.log.DebugContext(ctx, "attempting retieval of multiple records", slog.Any("filter", record))

	// Disallow filtering on a specific ID, as the single get can be used for that
	if record.Id != "" {
		return nil, fmt.Errorf("cannont do a filtered search through ID: %w", errs.ErrInvalidRequest)
	}

	// Grab out the user id and ensure it's a valid uuid
	userId, err := uuid.Parse(record.UserId)
	if err != nil {
		return nil, fmt.Errorf("user id is not a valid uuid: %w", errs.ErrInvalidRequest)
	}

	// Construct our filter
	filter := MapRecordToEntryWithoutUuids(record)
	filter.UserId = userId

	// Get the records
	found, err := s.store.GetManyFood(filter)
	if err != nil {
		return nil, err
	}

	// Map the found records into our sow domain struct
	records := make([]*sow.Record, len(found))
	for _, entry := range found {
		records = append(records, MapEntryToRecord(entry))
	}

	s.log.DebugContext(ctx, "successfully retrieved multiple records", slog.Any("records", records))
	return records, nil
}

func (s *FoodRecordService) Create(ctx context.Context, record *sow.Record) (*sow.Record, error) {
	s.log.DebugContext(ctx, "attempting creation of record", slog.Any("record", record))

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
	entry := MapRecordToEntryWithoutUuids(record)

	// Set uuids ourselveves
	entry.UserId = userId

	// Generate a UUID id
	if entry.Id == uuid.Nil {
		id, err := uuid.NewV7()
		if err != nil {
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

	s.log.DebugContext(ctx, "successfully created record", slog.Any("created", created))
	return created, nil
}

func (s *FoodRecordService) Delete(ctx context.Context, id uuid.UUID) error {
	s.log.DebugContext(ctx, "attempting deletion of record", slog.String("id", id.String()))

	err := s.store.DeleteFood(id)

	if err != nil {
		return err
	}

	s.log.DebugContext(ctx, "successfully deleted record", slog.String("id", id.String()))
	return nil
}

func (s *FoodRecordService) Update(ctx context.Context, record *sow.Record) error {
	s.log.DebugContext(ctx, "attempting creation of record", slog.Any("record", record))

	// Enforce a specific id has been set
	if record.Id == "" {
		return fmt.Errorf("id must be set: %w", errs.ErrInvalidRequest)
	}

	// We need to have some form of description, no matter what it is.
	if record.Description == "" {
		return fmt.Errorf("description must be set: %w", errs.ErrInvalidRequest)
	}

	// Perform mapping of the record for storage layer, ignoring any of the uuids we have set
	// on the record.
	entry, err := MapRecordToEntry(record)
	if err != nil {
		return err
	}

	// Attempt to update it
	err = s.store.UpdateFood(entry)
	if err != nil {
		return err
	}

	s.log.DebugContext(ctx, "successfully updated record", slog.Any("updated", record))
	return nil
}

func NewFoodRecorderService(logger *slog.Logger, store persistence.FoodStore) (*FoodRecordService, error) {
	return &FoodRecordService{log: logger, store: store}, nil
}
