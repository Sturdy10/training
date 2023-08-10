package repository

import (
	"abc/model"
	"database/sql"
	"log"
)

type RepositoryPort interface {
	GetRep() ([]model.GetRequest, error)
}
type repositoryAdapter struct {
	db *sql.DB
}

func NewRorepositoryAdapter(db *sql.DB) repositoryAdapter {
	return repositoryAdapter{db: db}
}

func (r repositoryAdapter) GetRep() ([]model.GetRequest, error) {
	query := "SELECT id, name, type, detail, url FROM beer"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	defer rows.Close()
	var beers []model.GetRequest
	for rows.Next() {
		var beer model.GetRequest
		err = rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Detail, &beer.URL)
		if err != nil {
			log.Panicln(err)
			return nil, err
		}
		beers = append(beers, beer)
	}

	return beers, nil
}
