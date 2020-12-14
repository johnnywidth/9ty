package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/client/entity"
	"github.com/johnnywidth/9ty/client/service"
	"github.com/johnnywidth/9ty/client/service/mock"
)

func TestPortDomainCreate(t *testing.T) {
	t.Run("create_succeeded", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainClient := mock.NewMockPortDomainClient(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}

		mockedPortDomainClient.EXPECT().
			Create(ctx, &api.PortRequest{
				Key:  "mocked_key",
				Name: mockedData.Name,
			}).
			Return(&api.Empty{}, nil).
			Times(1)

		s := service.NewPortDomain(mockedPortDomainClient)
		err := s.Create(ctx, "mocked_key", mockedData)
		assert.Nil(t, err)
	})

	t.Run("create_return_error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainClient := mock.NewMockPortDomainClient(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}
		mockedError := errors.New("mocked error")

		mockedPortDomainClient.EXPECT().
			Create(ctx, &api.PortRequest{
				Key:  "mocked_key",
				Name: mockedData.Name,
			}).
			Return(&api.Empty{}, mockedError).
			Times(1)

		s := service.NewPortDomain(mockedPortDomainClient)
		err := s.Create(ctx, "mocked_key", mockedData)
		assert.True(t, errors.Is(err, mockedError))
	})
}

func TestPortDomainGet(t *testing.T) {
	t.Run("get_succeeded", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainClient := mock.NewMockPortDomainClient(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}

		mockedPortDomainClient.EXPECT().
			Get(ctx, &api.GetRequest{Key: "mocked_key"}).
			Return(&api.PortResponse{
				Name: mockedData.Name,
			}, nil).
			Times(1)

		s := service.NewPortDomain(mockedPortDomainClient)
		e, err := s.Get(ctx, "mocked_key")
		assert.Nil(t, err)
		assert.Equal(t, mockedData, e)
	})

	t.Run("get_return_error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainClient := mock.NewMockPortDomainClient(mockCtrl)
		mockedError := errors.New("mocked error")

		mockedPortDomainClient.EXPECT().
			Get(ctx, &api.GetRequest{Key: "mocked_key"}).
			Return(&api.PortResponse{}, mockedError).
			Times(1)

		s := service.NewPortDomain(mockedPortDomainClient)
		_, err := s.Get(ctx, "mocked_key")
		assert.True(t, errors.Is(err, mockedError))
	})

	t.Run("get_return_empty_data", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainClient := mock.NewMockPortDomainClient(mockCtrl)

		mockedPortDomainClient.EXPECT().
			Get(ctx, &api.GetRequest{Key: "mocked_key"}).
			Return(&api.PortResponse{}, nil).
			Times(1)

		s := service.NewPortDomain(mockedPortDomainClient)
		_, err := s.Get(ctx, "mocked_key")
		assert.True(t, errors.Is(err, entity.ErrNotFound))
	})
}
