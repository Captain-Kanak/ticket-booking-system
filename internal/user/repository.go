package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo *repository) CreateUser(user *User) error {
	tx := repo.db.Create(user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
