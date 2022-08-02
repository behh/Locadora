package presenter

import (
	"time"

	"github.com/behh/locadora/entity"
)

type Carro struct {
	ID          entity.ID `json:"id"`
	Placa       string    `json:"placa"`
	Modelo      string    `json:"modelo"`
	Cor         string    `json:"cor"`
	Renavan     int       `json:"renavan"`
	Hodometro   int       `json:"hodometro"`
	DataCriacao time.Time `json:"dataCadastro"`
}
