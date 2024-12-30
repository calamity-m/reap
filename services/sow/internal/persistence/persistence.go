package persistence

import (
	"time"

	"github.com/google/uuid"
)

type FoodRecordEntry struct {
	Id          int
	Uuid        uuid.UUID
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
	CreateFood(FoodRecordEntry) error
	// Retrieve a single food record based on the
	// record's uuid. Internal DB primary key is ignored
	// by this call.
	GetFood(uuid uuid.UUID) (FoodRecordEntry, error)
	// Provided FoodRecordEntry is treated as a filter, allowing
	// the caller to retrieve multiple food records at will.
	GetFoods(filter FoodRecordEntry) ([]FoodRecordEntry, error)
	// Update the record in place
	UpdateFood(uuid uuid.UUID) error
	// Delete matching record
	DeleteFood(uuid uuid.UUID) error
}

type MemoryFoodStore struct {
	entries map[string]FoodRecordEntry
}

// Create a food record entry
func (s *MemoryFoodStore) CreateFood(_ FoodRecordEntry) error {
	panic("not implemented") // TODO: Implement
}

// Retrieve a single food record based on the
// record's uuid. Internal DB primary key is ignored
// by this call.
func (s *MemoryFoodStore) GetFood(uuid uuid.UUID) (FoodRecordEntry, error) {
	panic("not implemented") // TODO: Implement
}

// Provided FoodRecordEntry is treated as a filter, allowing
// the caller to retrieve multiple food records at will.
func (s *MemoryFoodStore) GetFoods(filter FoodRecordEntry) ([]FoodRecordEntry, error) {
	panic("not implemented") // TODO: Implement
}

// Update the record in place
func (s *MemoryFoodStore) UpdateFood(uuid uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// Delete matching record
func (s *MemoryFoodStore) DeleteFood(uuid uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func NewMemoryFoodStore() *MemoryFoodStore {
	return &MemoryFoodStore{}
}
