package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/johnnywidth/9ty/client/entity"
)

type callback func(ctx context.Context, e *entity.PortData) error

type PortJSON struct {
	reader           io.Reader
	portDataCallback callback
}

func NewPortJSON(
	reader io.Reader,
	portDataCallback callback,
) *PortJSON {
	return &PortJSON{
		reader:           reader,
		portDataCallback: portDataCallback,
	}
}

func (u *PortJSON) Parse(ctx context.Context) error {
	if u.portDataCallback == nil {
		return fmt.Errorf("callback not defined")
	}

	if u.reader == nil {
		return fmt.Errorf("reader not defined")
	}

	d := json.NewDecoder(u.reader)

	t, err := d.Token()
	if err != nil {
		return fmt.Errorf("release first token failed. %w", err)
	}

	for d.More() {
		ctxErr := ctx.Err()
		if ctxErr != nil {
			return ctxErr
		}

		t, err = d.Token()
		if err != nil {
			return fmt.Errorf("release key token failed. %w", err)
		}
		mapKey := t.(string)

		e := &entity.PortData{}
		err := d.Decode(e)
		if err != nil {
			return fmt.Errorf("decode port data failed. %w", err)
		}
		e.ID = mapKey

		err = u.portDataCallback(ctx, e)
		if err != nil {
			return fmt.Errorf("port callback failed. %w", err)
		}
	}

	t, err = d.Token()
	if err != nil {
		return fmt.Errorf("release last token failed. %w", err)
	}

	return nil
}
