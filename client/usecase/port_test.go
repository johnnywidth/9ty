package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/johnnywidth/9ty/client/entity"
	"github.com/johnnywidth/9ty/client/usecase"
	"github.com/johnnywidth/9ty/client/usecase/mock"
)

func TestPortCreate(t *testing.T) {
	t.Run("success_create", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainService := mock.NewMockPortDomainService(mockCtrl)
		e := &entity.PortData{
			Name: "mocked_name",
		}

		mockedPortDomainService.EXPECT().
			Create(ctx, "mocked_key", e).
			Return(nil).
			Times(1)

		portUsecase := usecase.NewPort(mockedPortDomainService)
		err := portUsecase.Create(ctx, "mocked_key", e)
		assert.Nil(t, err)
	})

	t.Run("create_return_error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainService := mock.NewMockPortDomainService(mockCtrl)
		e := &entity.PortData{
			Name: "mocked_name",
		}
		mockedErr := errors.New("mocked_error")

		mockedPortDomainService.EXPECT().
			Create(ctx, "mocked_key", e).
			Return(mockedErr).
			Times(1)

		portUsecase := usecase.NewPort(mockedPortDomainService)
		err := portUsecase.Create(ctx, "mocked_key", e)
		assert.True(t, errors.Is(err, mockedErr))
	})
}

func TestPortGet(t *testing.T) {
	t.Run("get_succeeded", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainService := mock.NewMockPortDomainService(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}

		mockedPortDomainService.EXPECT().
			Get(ctx, mockedData.Name).
			Return(mockedData, nil).
			Times(1)

		portUsecase := usecase.NewPort(mockedPortDomainService)
		e, err := portUsecase.Get(ctx, mockedData.Name)
		assert.Nil(t, err)
		assert.Equal(t, mockedData, e)
	})

	t.Run("get_return_error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainService := mock.NewMockPortDomainService(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}
		mockedErr := errors.New("mocked_error")

		mockedPortDomainService.EXPECT().
			Get(ctx, mockedData.Name).
			Return(nil, mockedErr).
			Times(1)

		portUsecase := usecase.NewPort(mockedPortDomainService)
		e, err := portUsecase.Get(ctx, mockedData.Name)
		assert.True(t, errors.Is(err, mockedErr))
		assert.Nil(t, e)
	})

	t.Run("get_return_empty_data", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortDomainService := mock.NewMockPortDomainService(mockCtrl)
		mockedData := &entity.PortData{
			Name: "mocked_name",
		}

		mockedPortDomainService.EXPECT().
			Get(ctx, mockedData.Name).
			Return(nil, nil).
			Times(1)

		portUsecase := usecase.NewPort(mockedPortDomainService)
		e, err := portUsecase.Get(ctx, mockedData.Name)
		assert.EqualError(t, err, entity.ErrNotFound.Error())
		assert.Nil(t, e)
	})
}
