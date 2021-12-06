package joker

import (
	"program/model"
	"program/storage"
)

type Server struct {
	storage storage.Storage
}

func NewServer(storage storage.Storage) *Server {
	s := &Server{
		storage: storage,
	}
	// s.storage.Insert()
	return s
}

func (s *Server) Last() ([]model.Event, []string, error) {

	result, str := s.storage.LastStartime()

	return result, str, nil
}

func (s *Server) EventsByTime() []model.Event {

	res := s.storage.EventsTime()
	return res

}
