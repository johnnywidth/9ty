package service

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/client/entity"
)

// PortDomain port domain service with grpc client
type PortDomain struct {
	client api.PortDomainClient
}

// NewPortDomain create new instance of port domain service
func NewPortDomain(
	client api.PortDomainClient,
) *PortDomain {
	return &PortDomain{
		client: client,
	}
}

// Create send port data to Port Domain service
func (s *PortDomain) Create(ctx context.Context, key string, e *entity.PortData) error {
	_, err := s.client.Create(ctx, &api.PortRequest{
		Key:         key,
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

// Get get port data from Port Domain service for given port key
func (s *PortDomain) Get(ctx context.Context, key string) (*entity.PortData, error) {
	e, err := s.client.Get(ctx, &api.GetRequest{Key: key})
	if err != nil {
		return nil, fmt.Errorf("service get port data failed. %w", err)
	}
	if e == nil || e.Name == "" {
		return nil, fmt.Errorf("service return empty port data. %w", entity.ErrNotFound)
	}

	return &entity.PortData{
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
