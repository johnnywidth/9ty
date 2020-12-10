package usecase

import (
	"context"

	"github.com/johnnywidth/9ty/client/entity"
)

type PortDomainService interface {
	Create(ctx context.Context, e *entity.PortData) error
	GetByName(ctx context.Context, name string) (*entity.PortData, error)
}
