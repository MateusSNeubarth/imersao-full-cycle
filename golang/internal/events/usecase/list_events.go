package usecase

import "github.com/MateusSNeubarth/imersao-full-cycle/internal/events/domain"

type ListEventsOutputDTO struct {
	Events []EventsDTO `json:"events"`
}

type ListEventsUsecase struct {
	repo domain.EventRepository
}

func NewListEventUseCase(repo domain.EventRepository) *ListEventsUsecase {
	return &ListEventsUsecase{repo: repo}
}

func (uc *ListEventsUsecase) Execute() (*ListEventsOutputDTO, error) {
	events, err := uc.repo.ListEvents()
	if err != nil {
		return nil, err
	}

	eventDTOs := make([]EventsDTO, len(events))
	for i, event := range events {
		eventDTOs[i] = EventsDTO{
			ID:           event.ID,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01-02 15:04:05"),
			ImageURL:     event.ImageURL,
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerID:    event.PartnerID,
		}
	}

	return &ListEventsOutputDTO{Events: eventDTOs}, nil
}
