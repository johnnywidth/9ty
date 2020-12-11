package service

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/client/entity"
)

type PortDomain struct {
	client api.PortDomainClient
}

func NewPortDomain(
	client api.PortDomainClient,
) *PortDomain {
	return &PortDomain{
		client: client,
	}
}

func (s *PortDomain) Create(ctx context.Context, e *entity.PortData) error {
	_, err := s.client.Create(ctx, &api.PortRequest{
		Id:          e.ID,
		Name:        e.Name,
		City:        e.City,
		Country:     e.Country,
		Alias:       e.Alias,
		Regions:     e.Regions,
		Coordinates: e.Coordinates,
		Province:    e.Province,
		Timezone:    e.Timezone,
		Unlocs:      e.Unlocs,
		Code:        e.Code,
	})
	if err != nil {
		return fmt.Errorf("service create port data failed. %w", err)
	}

	return nil
}

func (s *PortDomain) GetByName(ctx context.Context, name string) (*entity.PortData, error) {
	e, err := s.client.Get(ctx, &api.GetRequest{Name: name})
	if err != nil {
		return nil, fmt.Errorf("service get port data failed. %w", err)
	}
	if e == nil || e.Name == "" {
		return nil, fmt.Errorf("service return empty port data. %w", entity.ErrNotFound)
	}

	return &entity.PortData{
		ID:          e.Id,
		Name:        e.Name,
		City:        e.City,
		Country:     e.Country,
		Alias:       e.Alias,
		Regions:     e.Regions,
		Coordinates: e.Coordinates,
		Province:    e.Province,
		Timezone:    e.Timezone,
		Unlocs:      e.Unlocs,
		Code:        e.Code,
	}, nil
}
