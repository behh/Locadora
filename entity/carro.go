package entity

import "time"

type Carro struct {
	ID              ID
	Placa           string
	Modelo          string
	Cor             string
	Renavan         int
	Hodometro       int
	DataCriacao     time.Time
	DataAtualizacao time.Time
}

func NewCarro(Placa string, Modelo string, Cor string, Renavan int, Hodometro int) (*Carro, error) {
	c := &Carro{
		ID:          NewID(),
		Placa:       Placa,
		Modelo:      Modelo,
		Cor:         Cor,
		Renavan:     Renavan,
		Hodometro:   Hodometro,
		DataCriacao: time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

func (c *Carro) Validate() error {
	if len(c.Placa) == 0 || len(c.Modelo) == 0 {
		return ErrInvalidEntity
	}
	return nil
}
