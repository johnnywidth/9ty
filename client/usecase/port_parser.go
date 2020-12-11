package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/johnnywidth/9ty/client/entity"
)

type callback func(context.Context, string, *entity.PortData) error

// PortJSON json parser of port data
type PortJSON struct {
	reader           io.Reader
	portDataCallback callback
}

// NewPortJSON new instance of json parser with io reader and callback function with decoded data
func NewPortJSON(
	reader io.Reader,
	portDataCallback callback,
) *PortJSON {
	return &PortJSON{
		reader:           reader,
		portDataCallback: portDataCallback,
	}
}

// Parse json parser with decode port data entity by entity and call function with decoded data
func (u *PortJSON) Parse(ctx context.Context) error {
	if u.portDataCallback == nil {
		return fmt.Errorf("callback not defined")
	}

	if u.reader == nil {
		return fmt.Errorf("reader not defined")
	}

	d := json.NewDecoder(u.reader)

	_, err := d.Token()
	if err != nil {
		return fmt.Errorf("release first token failed. %w", err)
	}

	for d.More() {
		ctxErr := ctx.Err()
		if ctxErr != nil {
			return ctxErr
		}

		var t interface{}
		t, err = d.Token()
		if err != nil {
			return fmt.Errorf("release key token failed. %w", err)
		}
		mapKey, ok := t.(string)
		if !ok {
			return fmt.Errorf("cast key port data failed. %w", err)
		}

		e := &entity.PortData{}
		err = d.Decode(e)
		if err != nil {
			return fmt.Errorf("decode port data failed. %w", err)
		}

		err = u.portDataCallback(ctx, mapKey, e)
		if err != nil {
			return fmt.Errorf("port callback failed. %w", err)
		}
	}

	_, err = d.Token()
	if err != nil {
		return fmt.Errorf("release last token failed. %w", err)
	}

	return nil
}
