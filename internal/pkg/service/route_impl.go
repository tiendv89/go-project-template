package service

import (
	"meta-aggregator/internal/pkg/repositories"
)

type RouteService struct {
	repository repositories.IRouteRepository
}

func NewRouteService(repository repositories.IRouteRepository) IRouteSvc {
	return &RouteService{
		repository,
	}
}

func (t *RouteService) FindRoute() {

}
