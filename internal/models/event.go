package models

import (
	"context"
	"time"
)

type Event struct {
	Id        string
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, id string) (*Event, error)
	CreateOne(ctx context.Context, event *Event) error
	// UpdateOne(ctx context.Context, event *Event) error
	// DeleteOne(ctx context.Context, id string) error
}
