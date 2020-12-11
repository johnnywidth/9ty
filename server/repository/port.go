package repository

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/server/entity"
)

type Port struct {
	db DB
}

func NewPort(
	db DB,
) *Port {
	return &Port{
		db: db,
	}
}

func (r *Port) Create(ctx context.Context, e *entity.PortData) error {
	err := r.db.Set(e.ID, *e)
	if err != nil {
		return fmt.Errorf("db set port data failed. %w", err)
	}

	return nil
}

func (r *Port) GetByName(ctx context.Context, name string) (*entity.PortData, error) {
	data, err := r.db.Get(name)
	if err != nil {
		return nil, fmt.Errorf("db get port data failed. %w", err)
	}

	if data == nil {
		return nil, fmt.Errorf("get empty port data. %w", entity.ErrNotFound)
	}

	e, ok := data.(entity.PortData)
	if !ok {
		return nil, fmt.Errorf("cast tpye during get port data failed")
	}

	return &e, nil
}
