package carro

import (
	"fmt"
	"strings"

	"github.com/behh/locadora/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Carro
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Carro{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Create(e *entity.Carro) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *inmem) Get(id entity.ID) (*entity.Carro, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

func (r *inmem) Update(e *entity.Carro) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *inmem) Search(query string) ([]*entity.Carro, error) {
	var d []*entity.Carro
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Placa), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

func (r *inmem) List() ([]*entity.Carro, error) {
	var d []*entity.Carro
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}
