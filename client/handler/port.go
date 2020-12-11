package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/johnnywidth/9ty/client/entity"
)

// Port http server for port representation
type Port struct {
	portUsecase PortUsecase
}

// NewPort create new instance of Port with all dependencies
func NewPort(
	portUsecase PortUsecase,
) *Port {
	return &Port{
		portUsecase: portUsecase,
	}
}

// Get http GET request to retrieve Port data by port key
func (h *Port) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	routeVars := mux.Vars(r)
	name := routeVars["name"]

	e, err := h.portUsecase.Get(ctx, name)
	if errors.Is(err, entity.ErrNotFound) {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
