package user

import (
	"ticket-booking-system/internal/auth"
	"ticket-booking-system/internal/user/dto"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
	jwt  auth.JWTService
}

func NewService(repo Repository, jwt auth.JWTService) *service {
	return &service{repo, jwt}
}

func (s *service) CreateUser(req *dto.CreateRequest) (res *dto.UserResponse, err error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	err = user.hashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}

	return response, nil
}

func (s *service) LoginUser(req *dto.LoginRequest) (res *dto.LoginResponse, err error) {
	user, err := s.repo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, gorm.ErrRecordNotFound
	}

	err = user.checkPassword(req.Password)

	if err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateToken(user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		},
	}

	return response, nil
}
