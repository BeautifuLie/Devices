package storage

import (
	"program/model"
	"time"
)

type Storage interface {
	LastStartime(n int64) ([]model.Event, []string)
	EventsTime(t1, t2 time.Time) []model.Event
	Insert()
	CloseClientDB()
}
