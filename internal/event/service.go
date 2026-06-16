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
