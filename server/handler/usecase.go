package handler

import (
	"context"

	"github.com/johnnywidth/9ty/server/entity"
)

// PortUsecase port usecase
type PortUsecase interface {
	Create(ctx context.Context, key string, e *entity.PortData) error
	Get(ctx context.Context, key string) (*entity.PortData, error)
}
