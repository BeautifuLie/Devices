package storage

import (
	"program/model"
)

type Storage interface {
	LastStartime() ([]model.Event, []string)
	EventsTime() []model.Event
	Insert()
	CloseClientDB()
}
