package handler

import (
	"context"
	"fmt"

	"github.com/johnnywidth/9ty/api/port"
)

type PortServer struct{}

func NewPortServer() *PortServer {
	return &PortServer{}
}

func (s *PortServer) Create(ctx context.Context, r *port.PortMessageRequest) (*port.Empty, error) {
	fmt.Println(ctx.Err(), r)

	return &port.Empty{}, nil
}

func (s *PortServer) GetByName(ctx context.Context, r *port.GetByNameRequest) (*port.PortMessageRequest, error) {
	fmt.Println(r)

	return &port.PortMessageRequest{}, nil
}
