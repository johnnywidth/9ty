package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/johnnywidth/9ty/client/entity"
)

type PortHandler struct {
	portDataUsecase PortDataUsecase
}

func NewPortHandler(
	portDataUsecase PortDataUsecase,
) *PortHandler {
	return &PortHandler{
		portDataUsecase: portDataUsecase,
	}
}

func (h *PortHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	routeVars := mux.Vars(r)
	name := routeVars["name"]

	e, err := h.portDataUsecase.GetByName(ctx, name)
	if errors.Is(err, entity.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
