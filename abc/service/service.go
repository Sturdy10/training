package service

import (
	"abc/model"
	"abc/repository"
	"errors"
)

type ServicePort interface {
	GetSer() ([]model.GetRequestResponse, error)
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewserviceAdapter(r repository.RepositoryPort) serviceAdapter {
	return serviceAdapter{r: r}
}

func (s serviceAdapter) GetSer() ([]model.GetRequestResponse, error) {
	beers, err := s.r.GetRep()
	if err != nil {
		err = errors.New("can not connect to the server")
		return nil, err
	}
	var data []model.GetRequestResponse

	for _, beer := range beers {
		var a model.GetRequestResponse
		a.ID = beer.ID
		a.Name = beer.Name
		a.Type = beer.Type
		a.Detail = beer.Detail
		a.URL = beer.URL

		data = append(data, a)

	}

	return data, nil
}
