package user

import (
	"errors"
	"gorm.io/gorm"
)

var ErrorAlreadyExists = errors.New("user already exists with this email")

type Repository interface {
	CreateUser(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) CreateUser(user *User) error {
	tx := repo.db.Create(user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExists
		}

		return tx.Error

		// return c.JSON(http.StatusInternalServerError, Response{
		// 	Success: false,
		// 	Message: "Failed to create user!",
		// })
	}

	return nil
}
