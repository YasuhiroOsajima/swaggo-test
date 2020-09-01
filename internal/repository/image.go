package repository

import (
	"github.com/jmoiron/sqlx"
)

type ImageRepository struct {
	db *sqlx.DB
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{db}
}

type Image struct {
	Uuid  string `db:"uuid"`
	Name  string `db:"name"`
	Owner string `db:"owner"`
}

type Imagelist []*Image

func (i ImageRepository) FindAll() (Imagelist, error) {
	var imagelist Imagelist

	rows, err := i.db.Queryx("SELECT * FROM images")
	if err != nil {
		return nil, err
	}

	var image Image
	for rows.Next() {
		err := rows.StructScan(&image)
		if err != nil {
			return nil, err
		}
		imagelist = append(imagelist, &image)
	}

	return imagelist, nil
}
