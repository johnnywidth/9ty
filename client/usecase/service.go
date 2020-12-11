//go:generate mockgen -package mock -destination=mock/port_domain.go . PortDomainService

package usecase

import (
	"context"

	"github.com/johnnywidth/9ty/client/entity"
)

// PortDomainService port domain service interface
type PortDomainService interface {
	Create(ctx context.Context, key string, e *entity.PortData) error
	Get(ctx context.Context, key string) (*entity.PortData, error)
}
