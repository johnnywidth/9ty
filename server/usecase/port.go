package usecase

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/server/entity"
)

// Port port usecase with port repository
type Port struct {
	portRepository PortRepository
}

// NewPort new instance of port usecase
func NewPort(
	portRepository PortRepository,
) *Port {
	return &Port{
		portRepository: portRepository,
	}
}

// Create create port data and store in the repository
func (u *Port) Create(ctx context.Context, key string, e *entity.PortData) error {
	err := u.portRepository.Create(ctx, key, e)
	if err != nil {
		return fmt.Errorf("create port failed. %w", err)
	}

	return nil
}

// Get get existed port data by key
func (u *Port) Get(ctx context.Context, key string) (*entity.PortData, error) {
	e, err := u.portRepository.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("get by name failed. %w", err)
	}
	if e == nil {
		return nil, fmt.Errorf("get port by name return empty data. %w", entity.ErrNotFound)
	}

	return e, nil
}
