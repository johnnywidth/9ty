package usecase

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/server/entity"
)

type Port struct {
	portRepository PortRepository
}

func NewPort(
	portRepository PortRepository,
) *Port {
	return &Port{
		portRepository: portRepository,
	}
}

func (u *Port) Create(ctx context.Context, e *entity.PortData) error {
	err := u.portRepository.Create(ctx, e)
	if err != nil {
		return fmt.Errorf("create port failed. %w", err)
	}

	return nil
}

func (u *Port) GetByName(ctx context.Context, name string) (*entity.PortData, error) {
	e, err := u.portRepository.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("get by name failed. %w", err)
	}
	if e == nil {
		return nil, fmt.Errorf("get port by name return empty data. %w", entity.ErrNotFound)
	}

	return e, nil
}
