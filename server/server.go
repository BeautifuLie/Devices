package server

import (
	"fmt"
	"program/model"
	"program/storage"
	"strconv"
	"time"
)

type Server struct {
	storage storage.Storage
}

func NewServer(storage storage.Storage) *Server {
	s := &Server{
		storage: storage,
	}
	// s.storage.Insert() //разкомментировать для вставки документов
	return s
}

func (s *Server) Last(l string) ([]model.Event, []string, error) {
	a, _ := strconv.Atoi(l)
	n := int64(a)
	result, str := s.storage.LastStartime(n)

	return result, str, nil
}

func (s *Server) EventsByTime(start, end string) []model.Event {
	startString := string("2021-12-06 " + start)
	startTime, err := time.Parse("2006-01-02 15:04:05", startString)
	if err != nil {
		fmt.Errorf("startTime parsing %w", err)
	}

	endString := string("2021-12-06 " + end)
	endTime, err := time.Parse("2006-01-02 15:04:05", endString)
	if err != nil {
		fmt.Errorf("startTime parsing %w", err)
	}

	res := s.storage.EventsTime(startTime, endTime)
	return res

}
