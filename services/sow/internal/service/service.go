package service

import "github.com/google/uuid"

type RecordService interface {
	Get(uuid.UUID) (Record, error)
	GetFiltered(Record) (Record, error)

	Create(Record) (Record, error)

	Delete(uuid.UUID) error
	DeleteFiltered(Record) error

	Update(uuid.UUID) error
	UpdateFiltered(Record) (Record, error)
}
