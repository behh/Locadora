package carro

import (
	"strings"
	"time"

	"github.com/behh/locadora/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateCarro(placa string, modelo string, cor string, renavan int, hodometro int) (entity.ID, error) {
	c, err := entity.NewCarro(placa, modelo, cor, renavan, hodometro)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(c)
}

func (s *Service) GetCarro(id entity.ID) (*entity.Carro, error) {
	c, err := s.repo.Get(id)
	if c == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) SearchCarros(query string) ([]*entity.Carro, error) {
	carros, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(carros) == 0 {
		return nil, entity.ErrNotFound
	}
	return carros, nil
}

func (s *Service) ListCarros() ([]*entity.Carro, error) {
	carros, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(carros) == 0 {
		return nil, entity.ErrNotFound
	}
	return carros, nil
}

func (s *Service) DeleteCarro(id entity.ID) error {
	_, err := s.GetCarro(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateCarro(e *entity.Carro) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.DataAtualizacao = time.Now()
	return s.repo.Update(e)
}
