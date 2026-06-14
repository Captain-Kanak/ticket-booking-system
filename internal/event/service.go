package event

import "ticket-booking-system/internal/event/dto"

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

	res = &dto.Response{
		ID:               event.ID,
		Title:            event.Title,
		Description:      event.Description,
		Location:         event.Location,
		StartDate:        event.StartDate,
		TotalTickets:     event.TotalTickets,
		AvailableTickets: event.TotalTickets,
		Price:            event.Price,
		CreatedAt:        event.CreatedAt,
	}

	return res, nil
}
