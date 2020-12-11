package usecase_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/johnnywidth/9ty/client/entity"
	"github.com/johnnywidth/9ty/client/usecase"
)

var mockedJSONData = `
	{ 
		"AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [], "regions": [], 
			"coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai", 
			"unlocs": [ "AEAJM" ], "code": "52000" }, 
		"AEAUH": { "name": "Abu Dhabi", "coordinates": [ 54.37, 24.47 ], "city": "Abu Dhabi", 
			"province": "Abu Z¸aby [Abu Dhabi]", "country": "United Arab Emirates", "alias": [], "regions": [], 
			"timezone": "Asia/Dubai", "unlocs": [ "AEAUH" ], "code": "52001" }
	}
`

func TestPortJSONParse(t *testing.T) {
	t.Run("parse_json_succeeded", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		expectedKey := "AEAJM"
		expectedData := &entity.PortData{
			Name:        "Ajman",
			City:        "Ajman",
			Country:     "United Arab Emirates",
			Alias:       []string{},
			Regions:     []string{},
			Coordinates: []float64{55.5136433, 25.4052165},
			Province:    "Ajman",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEAJM"},
			Code:        "52000",
		}

		ctx := context.Background()
		mockedReader := strings.NewReader(mockedJSONData)

		count := 0
		var firstPortDataKey string
		var firstPortData *entity.PortData
		mockedCallback := func(ctx context.Context, key string, e *entity.PortData) error {
			count++
			if count == 1 {
				firstPortDataKey = key
				firstPortData = e
			}
			return nil
		}

		u := usecase.NewPortJSON(mockedReader, mockedCallback)
		err := u.Parse(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expectedData, firstPortData)
		assert.Equal(t, expectedKey, firstPortDataKey)
		assert.Equal(t, 2, count)
	})

	t.Run("parse_not_json_failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedReader := strings.NewReader(`not_json`)

		mockedCallback := func(ctx context.Context, key string, e *entity.PortData) error {
			return nil
		}

		u := usecase.NewPortJSON(mockedReader, mockedCallback)
		err := u.Parse(ctx)
		assert.NotNil(t, err)
	})

	t.Run("parse_json_invalid", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedReader := strings.NewReader(`{"key": "value",}`)

		mockedCallback := func(ctx context.Context, key string, e *entity.PortData) error {
			return nil
		}

		u := usecase.NewPortJSON(mockedReader, mockedCallback)
		err := u.Parse(ctx)
		assert.NotNil(t, err)
	})

	t.Run("parse_json_callback_failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedReader := strings.NewReader(mockedJSONData)
		mockedErr := errors.New("mocked error")

		mockedCallback := func(ctx context.Context, key string, e *entity.PortData) error {
			return mockedErr
		}

		u := usecase.NewPortJSON(mockedReader, mockedCallback)
		err := u.Parse(ctx)
		assert.True(t, errors.Is(err, mockedErr))
	})

	t.Run("callback_undefined", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedReader := strings.NewReader(mockedJSONData)

		u := usecase.NewPortJSON(mockedReader, nil)
		err := u.Parse(ctx)
		assert.EqualError(t, err, "callback not defined")
	})

	t.Run("reader_undefined", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedCallback := func(ctx context.Context, key string, e *entity.PortData) error {
			return nil
		}

		u := usecase.NewPortJSON(nil, mockedCallback)
		err := u.Parse(ctx)
		assert.EqualError(t, err, "reader not defined")
	})
}
