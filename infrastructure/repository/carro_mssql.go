package repository

import (
	"database/sql"
	"time"

	"github.com/behh/locadora/entity"
)

type CarroMSSQL struct {
	db *sql.DB
}

func NewCarroMSSQL(db *sql.DB) *CarroMSSQL {
	return &CarroMSSQL{
		db: db,
	}
}

func (r *CarroMSSQL) Create(e *entity.Carro) (entity.ID, error) {
	_, err := r.db.Exec(sqlCreateCarro, sql.Named("ID", e.ID), sql.Named("Placa", e.Placa), sql.Named("Modelo", e.Modelo),
		sql.Named("Cor", e.Cor), sql.Named("Renavan", e.Renavan), sql.Named("Hodometro", e.Hodometro),
		sql.Named("DataCriacao", e.DataCriacao.Format(LayoutTimeStampMSSQL)))
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

func (r *CarroMSSQL) Get(ID entity.ID) (*entity.Carro, error) {
	var c entity.Carro

	err := r.db.QueryRow(sqlGetCarroPorID, sql.Named("ID", ID)).Scan(&c.Placa, &c.Modelo, &c.Cor, &c.Renavan, &c.Hodometro, &c.DataCriacao)
	if err != nil {
		return nil, err
	}
	c.ID = ID
	return &c, nil
}

func (r *CarroMSSQL) Update(e *entity.Carro) error {
	e.DataAtualizacao = time.Now()
	_, err := r.db.Exec(sqlUpdateCarro, sql.Named("Placa", e.Placa), sql.Named("Modelo", e.Modelo),
		sql.Named("Cor", e.Cor), sql.Named("Renavan", e.Renavan), sql.Named("Hodometro", e.Hodometro),
		sql.Named("DataAtualizacao", e.DataAtualizacao.Format(LayoutTimeStampMSSQL)), sql.Named("ID", e.ID))
	if err != nil {
		return err
	}
	return nil
}

func (r *CarroMSSQL) Search(query string) ([]*entity.Carro, error) {
	var carros []*entity.Carro

	rows, err := r.db.Query(sqlSearchCarro, sql.Named("query", query))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Carro
		err := rows.Scan(&c.ID, &c.Placa, &c.Modelo, &c.Cor, &c.Renavan, &c.Hodometro, &c.DataCriacao)
		if err != nil {
			return nil, err
		}
		carros = append(carros, &c)
	}
	return carros, nil
}

func (r *CarroMSSQL) List() ([]*entity.Carro, error) {
	var carros []*entity.Carro

	rows, err := r.db.Query(sqlListCarro)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Carro
		err := rows.Scan(&c.ID, &c.Placa, &c.Modelo, &c.Cor, &c.Renavan, &c.Hodometro, &c.DataCriacao)
		if err != nil {
			return nil, err
		}
		carros = append(carros, &c)
	}
	return carros, nil
}

func (r *CarroMSSQL) Delete(ID entity.ID) error {
	_, err := r.db.Exec(sqlDeleteCarro, sql.Named("ID", ID))
	if err != nil {
		return err
	}
	return nil
}

const sqlCreateCarro = `
INSERT INTO Carros(
	ID, Placa, Modelo, Cor, Renavan, Hodometro, DataCriacao)
VALUES (
	@ID, @Placa, @Modelo, @Cor, @Renavan, @Hodometro, @DataCriacao)`

const sqlGetCarroPorID = `
SELECT
	Placa, 
	Modelo, 
	Cor, 
	Renavan, 
	Hodometro, 
	DataCriacao
FROM Carros
WHERE ID = @ID`

const sqlUpdateCarro = `
UPDATE Carros
SET Placa = @Placa,
	Modelo = @Modelo,
	Cor = @Cor,
	Renavan = @Renavan,
	Hodometro = @Hodometro,
	DataAtualizacao = @DataAtualizacao
WHERE ID = @ID`

const sqlSearchCarro = `
SELECT
	ID,
	Placa, 
	Modelo, 
	Cor, 
	Renavan, 
	Hodometro, 
	DataCriacao
FROM Carros
WHERE Placa like '%' +  @query + '%'`

const sqlListCarro = `
SELECT
	ID,
	Placa, 
	Modelo, 
	Cor, 
	Renavan, 
	Hodometro, 
	DataCriacao
FROM Carros`

const sqlDeleteCarro = `
DELETE FROM Carros WHERE ID = @ID`
