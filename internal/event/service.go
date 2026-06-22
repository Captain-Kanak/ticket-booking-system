package event

import (
	"ticket-booking-system/internal/event/dto"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Create(req *dto.CreateRequest) (res *dto.Response, err error) {
	event := &Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartDate:        req.StartDate,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}

	err = s.repo.Create(event)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) GetAll() (res []*dto.Response, err error) {
	events, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, event := range events {
		res = append(res, event.ToResponse())
	}

	return res, nil
}

func (s *service) GetByID(id uuid.UUID) (res *dto.Response, err error) {
	event, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

// func (s *service) UpdateById(id uuid.UUID, req *dto.UpdateRequest) (res *dto.Response, err error) {
// 	event, err := s.repo.GetByID(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	event.Title = req.Title
// 	event.Description = req.Description
// 	event.Location = req.Location
// 	event.StartDate = req.StartDate
// 	event.Price = req.Price

// 	// err = s.repo.Update(id, event)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// }
