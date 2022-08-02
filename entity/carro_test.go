package entity_test

import (
	"testing"

	"github.com/behh/locadora/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewCarro(t *testing.T) {
	c, err := entity.NewCarro("ABC1234", "Corsa", "Preto", 1234567, 0)
	assert.Nil(t, err)
	assert.Equal(t, c.Placa, "ABC1234")
	assert.Equal(t, c.Renavan, 1234567)
	assert.NotNil(t, c.ID)
}

func TestCarroValidate(t *testing.T) {
	type test struct {
		Placa     string
		Modelo    string
		Cor       string
		Renavan   int
		Hodometro int
		want      error
	}

	tests := []test{
		{
			Placa:     "ABC1234",
			Modelo:    "Corsa",
			Cor:       "Preto",
			Renavan:   1234567,
			Hodometro: 0,
			want:      nil,
		},
		{
			Placa:     "",
			Modelo:    "Corsa",
			Cor:       "Preto",
			Renavan:   1234567,
			Hodometro: 0,
			want:      entity.ErrInvalidEntity,
		},
		{
			Placa:     "ABC1234",
			Modelo:    "",
			Cor:       "Preto",
			Renavan:   1234567,
			Hodometro: 0,
			want:      entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewCarro(tc.Placa, tc.Modelo, tc.Cor, tc.Renavan, tc.Hodometro)
		assert.Equal(t, err, tc.want)
	}

}
