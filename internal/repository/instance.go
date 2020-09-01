package repository

import (
	"github.com/jmoiron/sqlx"
)

type InstanceRepository struct {
	db *sqlx.DB
}

func NewInstanceRepository() *InstanceRepository {
	return &InstanceRepository{db}
}

type Instance struct {
	Uuid  string `db:"uuid"`
	Name  string `db:"name"`
	Owner string `db:"owner"`
}

type Instancelist []*Instance

func (i InstanceRepository) FindAll() (Instancelist, error) {
	var instancelist Instancelist

	rows, err := i.db.Queryx("SELECT * FROM instances")
	if err != nil {
		return nil, err
	}

	var instance Instance
	for rows.Next() {
		err := rows.StructScan(&instance)
		if err != nil {
			return nil, err
		}
		instancelist = append(instancelist, &instance)
	}

	return instancelist, nil
}
