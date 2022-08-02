package carro

import "github.com/behh/locadora/entity"

type Reader interface {
	Get(id entity.ID) (*entity.Carro, error)
	Search(query string) ([]*entity.Carro, error)
	List() ([]*entity.Carro, error)
}

type Writer interface {
	Create(c *entity.Carro) (entity.ID, error)
	Update(c *entity.Carro) error
	Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetCarro(id entity.ID) (*entity.Carro, error)
	SearchCarros(query string) ([]*entity.Carro, error)
	ListCarros() ([]*entity.Carro, error)
	CreateCarro(placa string, modelo string, cor string, renavan int, hodometro int) (entity.ID, error)
	UpdateCarro(c *entity.Carro) error
	DeleteCarro(id entity.ID) error
}
