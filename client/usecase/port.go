package usecase

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/client/entity"
)

// Port usecase for create and get port data
type Port struct {
	portDomainService PortDomainService
}

// NewPort new instance of port usecase with Port Domain service
func NewPort(
	portDomainService PortDomainService,
) *Port {
	return &Port{
		portDomainService: portDomainService,
	}
}

// Create send port data to Port Domain service
func (u *Port) Create(ctx context.Context, key string, e *entity.PortData) error {
	err := u.portDomainService.Create(ctx, key, e)
	if err != nil {
		return fmt.Errorf("create failed. %w", err)
	}

	return nil
}

// Get get port data by key from Port Domain service
func (u *Port) Get(ctx context.Context, name string) (*entity.PortData, error) {
	e, err := u.portDomainService.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("get port data failed. %w", err)
	}

	if e == nil {
		return nil, entity.ErrNotFound
	}

	return e, nil
}
