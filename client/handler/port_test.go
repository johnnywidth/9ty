package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/johnnywidth/9ty/client/entity"
	"github.com/johnnywidth/9ty/client/handler"
	"github.com/johnnywidth/9ty/client/handler/mock"
)

func TestPortGet(t *testing.T) {
	t.Run("get_succeeded", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortUsecase := mock.NewMockPortUsecase(mockCtrl)
		expectedData := &entity.PortData{
			Name: "mocked_name",
		}
		mockedURL := "/port/" + expectedData.Name

		mockedPortUsecase.EXPECT().
			Get(gomock.Any(), expectedData.Name).
			Return(expectedData, nil).
			Times(1)

		h := handler.NewPort(mockedPortUsecase)

		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/port/{name}", h.Get).Methods(http.MethodGet)

		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, mockedURL, nil)
		assert.Nil(t, err)

		req = req.WithContext(ctx)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("get_return_error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortUsecase := mock.NewMockPortUsecase(mockCtrl)
		expectedData := &entity.PortData{
			Name: "mocked_name",
		}
		mockedURL := "/port/" + expectedData.Name
		mockedError := errors.New("mocked error")

		mockedPortUsecase.EXPECT().
			Get(gomock.Any(), expectedData.Name).
			Return(nil, mockedError).
			Times(1)

		h := handler.NewPort(mockedPortUsecase)

		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/port/{name}", h.Get).Methods(http.MethodGet)

		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, mockedURL, nil)
		assert.Nil(t, err)

		req = req.WithContext(ctx)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("get_return_not_found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()
		mockedPortUsecase := mock.NewMockPortUsecase(mockCtrl)
		expectedData := &entity.PortData{
			Name: "mocked_name",
		}
		mockedURL := "/port/" + expectedData.Name

		mockedPortUsecase.EXPECT().
			Get(gomock.Any(), expectedData.Name).
			Return(nil, entity.ErrNotFound).
			Times(1)

		h := handler.NewPort(mockedPortUsecase)

		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/port/{name}", h.Get).Methods(http.MethodGet)

		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, mockedURL, nil)
		assert.Nil(t, err)

		req = req.WithContext(ctx)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
