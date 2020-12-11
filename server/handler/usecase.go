package handler

import (
	"context"

	"github.com/johnnywidth/9ty/server/entity"
)

type PortUsecase interface {
	Create(ctx context.Context, e *entity.PortData) error
	GetByName(ctx context.Context, name string) (*entity.PortData, error)
}
