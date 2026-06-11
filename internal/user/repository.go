package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
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

func (repo *repository) GetUserByEmail(email string) (*User, error) {
	var user User

	tx := repo.db.Where(&User{Email: email}).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
