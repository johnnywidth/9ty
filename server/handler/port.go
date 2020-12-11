package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/server/entity"
)

type PortServer struct {
	portUsecase PortUsecase
}

func NewPortServer(
	portUsecase PortUsecase,
) *PortServer {
	return &PortServer{
		portUsecase: portUsecase,
	}
}

func (h *PortServer) Create(ctx context.Context, r *api.PortRequest) (*api.Empty, error) {
	err := h.portUsecase.Create(ctx, &entity.PortData{
		ID:          r.Id,
		Name:        r.Name,
		City:        r.City,
		Country:     r.Country,
		Alias:       r.Alias,
		Regions:     r.Regions,
		Coordinates: r.Coordinates,
		Province:    r.Province,
		Timezone:    r.Timezone,
		Unlocs:      r.Unlocs,
		Code:        r.Code,
	})

	if err != nil {
		err = fmt.Errorf("create port failed. %w", err)
		fmt.Println(err)
		return &api.Empty{}, err
	}

	return &api.Empty{}, nil
}

func (h *PortServer) Get(ctx context.Context, r *api.GetRequest) (*api.PortResponse, error) {
	e, err := h.portUsecase.GetByName(ctx, r.Name)
	if errors.Is(err, entity.ErrNotFound) {
		return &api.PortResponse{}, nil
	} else if err != nil {
		err = fmt.Errorf("get by name failed. %w", err)
		fmt.Println(err)
		return &api.PortResponse{}, err
	}

	return &api.PortResponse{
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
	}, nil
}
