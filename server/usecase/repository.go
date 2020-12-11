package usecase

import (
	"context"

	"github.com/johnnywidth/9ty/server/entity"
)

// PortRepository port repository to create and get port data
type PortRepository interface {
	Create(ctx context.Context, key string, e *entity.PortData) error
	Get(ctx context.Context, key string) (*entity.PortData, error)
}
