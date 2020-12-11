//go:generate mockgen -package mock -destination=mock/port.go . PortUsecase

package handler

import (
	"context"

	"github.com/johnnywidth/9ty/client/entity"
)

// PortUsecase usecase interface
type PortUsecase interface {
	Get(ctx context.Context, key string) (*entity.PortData, error)
}
