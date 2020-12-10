package service

import (
	"context"

	"github.com/johnnywidth/9ty/api/port"

	"github.com/johnnywidth/9ty/client/entity"
)

type PortDomain struct {
	client port.PortClient
}

func NewPortDomain(
	client port.PortClient,
) *PortDomain {
	return &PortDomain{
		client: client,
	}
}

func (s *PortDomain) Create(ctx context.Context, e *entity.PortData) error {
	_, err := s.client.Create(ctx, &port.PortMessageRequest{
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

	return err
}

func (s *PortDomain) GetByName(ctx context.Context, name string) (*entity.PortData, error) {
	e, err := s.client.GetByName(ctx, &port.GetByNameRequest{Name: name})
	if err != nil {
		return nil, err
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
