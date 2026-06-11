package user

import "ticket-booking-system/internal/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
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
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
