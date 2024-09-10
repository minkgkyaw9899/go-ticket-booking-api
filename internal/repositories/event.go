package repositories

import (
	"context"
	"time"

	"github.com/minkgkyaw9899/go-ticket-booking-api/internal/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		Id:        "000001",
		Name:      "Event 1",
		Location:  "Yangon",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) error {
	return nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{db: db}
}
