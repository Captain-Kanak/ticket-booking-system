package user

import (
	"errors"
	"gorm.io/gorm"
)

var errorAlreadyExists = errors.New("user already exists with this email")

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
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			return errorAlreadyExists
		}

		return tx.Error
	}

	return nil
}
