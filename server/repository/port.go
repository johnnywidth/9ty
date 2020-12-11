package repository

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/server/entity"
)

// Port port repository with datastore
type Port struct {
	db DB
}

// NewPort new instance of port repository
func NewPort(
	db DB,
) *Port {
	return &Port{
		db: db,
	}
}

// Create set port data to datastore for given key
func (r *Port) Create(ctx context.Context, key string, e *entity.PortData) error {
	err := r.db.Set(key, *e)
	if err != nil {
		return fmt.Errorf("db set port data failed. %w", err)
	}

	return nil
}

// Get get port data for given key
func (r *Port) Get(ctx context.Context, key string) (*entity.PortData, error) {
	data, err := r.db.Get(key)
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
