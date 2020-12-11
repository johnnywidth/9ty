package usecase

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/client/entity"
)

type Port struct {
	portDomainService PortDomainService
}

func NewPort(
	portDomainService PortDomainService,
) *Port {
	return &Port{
		portDomainService: portDomainService,
	}
}

func (u *Port) Create(ctx context.Context, e *entity.PortData) error {
	err := u.portDomainService.Create(ctx, e)
	if err != nil {
		return fmt.Errorf("create failed. %w", err)
	}

	return nil
}

func (u *Port) GetByName(ctx context.Context, name string) (*entity.PortData, error) {
	e, err := u.portDomainService.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("get port data failed. %w", err)
	}

	if e == nil {
		return nil, entity.ErrNotFound
	}

	return e, nil
}
