package usecase

import (
	"context"

	"github.com/johnnywidth/9ty/server/entity"
)

type PortRepository interface {
	Create(ctx context.Context, e *entity.PortData) error
	GetByName(ctx context.Context, name string) (*entity.PortData, error)
}
