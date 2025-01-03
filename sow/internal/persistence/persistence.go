package persistence

import (
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/calamity-m/reap/pkg/errs"
	"github.com/google/uuid"
)

type FoodRecordEntry struct {
	DbId        int
	Id          uuid.UUID
	UserId      uuid.UUID
	Name        string
	Description string
	KJ          float32
	Grams       float32
	ML          float32
	Created     time.Time
}

type FoodStore interface {
	// Create a food record entry
	CreateFood(record FoodRecordEntry) error
	// Retrieve a single food record based on the
	// record's uuid. Internal DB primary key is ignored
	// by this call.
	GetFood(uuid uuid.UUID) (FoodRecordEntry, error)
	// Provided FoodRecordEntry is treated as a filter, allowing
	// the caller to retrieve multiple food records at will.
	GetManyFood(filter FoodRecordEntry) ([]FoodRecordEntry, error)
	// Update the record in place
	UpdateFood(record FoodRecordEntry) error
	// Delete matching record
	DeleteFood(uuid uuid.UUID) error
}

type MemoryFoodStore struct {
	mux     sync.Mutex
	entries map[string]FoodRecordEntry
	log     *slog.Logger
}

// Create a food record entry
func (s *MemoryFoodStore) CreateFood(record FoodRecordEntry) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.entries[record.Id.String()] = record

	if s.log != nil {
		// Safety debug logging :)
		s.log.Debug("updated in memory store with a creation", slog.Any("entries", s.entries))
	}

	return nil
}

// Retrieve a single food record based on the
// record's uuid. Internal DB primary key is ignored
// by this call.
func (s *MemoryFoodStore) GetFood(uuid uuid.UUID) (FoodRecordEntry, error) {
	found, ok := s.entries[uuid.String()]

	if !ok {
		return FoodRecordEntry{}, errs.ErrNotFound
	}

	return found, nil
}

// Provided FoodRecordEntry is treated as a filter, allowing
// the caller to retrieve multiple food records at will.
func (s *MemoryFoodStore) GetManyFood(filter FoodRecordEntry) ([]FoodRecordEntry, error) {
	entries := make([]FoodRecordEntry, 0, 1)

	for _, val := range s.entries {
		// Skip non matching user ids
		if val.UserId != filter.UserId {
			continue
		}

		// Check KJ/Grams/ML only if they're non zero
		if filter.KJ != 0 && filter.KJ != val.KJ {
			continue
		}

		if filter.Grams != 0 && filter.Grams != val.Grams {
			continue
		}

		if filter.ML != 0 && filter.Grams != val.Grams {
			continue
		}

		if !strings.Contains(val.Name, filter.Name) {
			continue
		}

		if !strings.Contains(val.Description, filter.Description) {
			continue
		}

		entries = append(entries, val)
	}

	return entries, nil
}

// Update the record in place
func (s *MemoryFoodStore) UpdateFood(record FoodRecordEntry) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.entries[record.Id.String()] = record

	if s.log != nil {
		// Safety debug logging :)
		s.log.Debug("updated in memory store with update", slog.Any("entries", s.entries))
	}

	return nil
}

// Delete matching record
func (s *MemoryFoodStore) DeleteFood(uuid uuid.UUID) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.entries, uuid.String())

	return nil
}

func NewMemoryFoodStore() *MemoryFoodStore {
	entries := map[string]FoodRecordEntry{"01942487-8295-7800-95c1-44d5c0a67099": {}}
	return &MemoryFoodStore{entries: entries}
}
