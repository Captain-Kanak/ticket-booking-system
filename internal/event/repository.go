package event

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(event *Event) error
	GetAll() ([]*Event, error)
	GetByID(id uuid.UUID) (*Event, error)
	Update(id uuid.UUID, event *Event) (*Event, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(event *Event) error {
	tx := r.db.Create(event)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *repository) GetAll() ([]*Event, error) {
	var events []*Event

	return events, nil
}

func (r *repository) GetByID(id uuid.UUID) (*Event, error) {
	var event *Event

	return event, nil
}

func (r *repository) Update(id uuid.UUID, event *Event) (*Event, error) {

	return event, nil
}

func (r *repository) Delete(id uuid.UUID) error {

	return nil
}
