package carro

import (
	"testing"
	"time"

	"github.com/behh/locadora/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureCarro() *entity.Carro {
	return &entity.Carro{
		Placa:       "ABC1234",
		Modelo:      "Corsa",
		Cor:         "Preto",
		Renavan:     1234567,
		Hodometro:   0,
		DataCriacao: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	c := newFixtureCarro()
	_, err := m.CreateCarro(c.Placa, c.Modelo, c.Cor, c.Renavan, c.Hodometro)
	assert.Nil(t, err)
	assert.False(t, c.DataCriacao.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	c1 := newFixtureCarro()
	c2 := newFixtureCarro()
	c2.Placa = "QWE9876"

	cID, _ := m.CreateCarro(c1.Placa, c1.Modelo, c1.Cor, c1.Renavan, c1.Hodometro)
	_, _ = m.CreateCarro(c2.Placa, c2.Modelo, c2.Cor, c2.Renavan, c2.Hodometro)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchCarros("ABC")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "ABC1234", c[0].Placa)

		c, err = m.SearchCarros("valor aleatorio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListCarros()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetCarro(cID)
		assert.Nil(t, err)
		assert.Equal(t, c1.Placa, saved.Placa)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	c := newFixtureCarro()
	id, err := m.CreateCarro(c.Placa, c.Modelo, c.Cor, c.Renavan, c.Hodometro)
	assert.Nil(t, err)
	saved, _ := m.GetCarro(id)
	saved.Cor = "Azul"
	assert.Nil(t, m.UpdateCarro(saved))
	updated, err := m.GetCarro(id)
	assert.Nil(t, err)
	assert.Equal(t, "Azul", updated.Cor)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	c1 := newFixtureCarro()
	c2 := newFixtureCarro()
	c2ID, _ := m.CreateCarro(c2.Placa, c2.Modelo, c2.Cor, c2.Renavan, c2.Hodometro)

	err := m.DeleteCarro(c1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteCarro(c2ID)
	assert.Nil(t, err)
	_, err = m.GetCarro(c2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
