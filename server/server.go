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
	var n int64
	a, _ := strconv.Atoi(l)
	if a == 0 {
		n = 10
	} else {
		n = int64(a)
	}

	result, str := s.storage.LastStartime(n)

	return result, str, nil
}

func (s *Server) EventsByTime(start, end, event string) ([]model.Event, error) {
	startString := string("2021-12-06 " + start)
	startTime, err := time.Parse("2006-01-02 15:04:05", startString)
	if err != nil {
		return nil, fmt.Errorf("startTime parsing %w", err)
	}

	endString := string("2021-12-06 " + end)
	endTime, err := time.Parse("2006-01-02 15:04:05", endString)
	if err != nil {
		return nil, fmt.Errorf("startTime parsing %w", err)
	}

	res := s.storage.EventsTime(startTime, endTime, event)
	return res, nil

}
