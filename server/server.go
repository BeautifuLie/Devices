package server

import (
	"program/model"
	"program/storage"
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

func (s *Server) Last() ([]model.Event, []string, error) {

	result, str := s.storage.LastStartime()

	return result, str, nil
}

func (s *Server) EventsByTime() []model.Event {
	t1 := time.Date(2021, time.December, 6, 8, 1, 0, 0, time.Local)
	t2 := time.Date(2021, time.December, 6, 8, 2, 0, 0, time.Local)
	res := s.storage.EventsTime(t1, t2)
	return res

}
