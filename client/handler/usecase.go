package handler

import (
	"context"

	"github.com/johnnywidth/9ty/client/entity"
)

type PortDataUsecase interface {
	GetByName(ctx context.Context, name string) (*entity.PortData, error)
}
