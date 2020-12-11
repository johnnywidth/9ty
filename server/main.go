package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/server/handler"
	"github.com/johnnywidth/9ty/server/persistance"
	"github.com/johnnywidth/9ty/server/repository"
	"github.com/johnnywidth/9ty/server/usecase"
)

type stopFunc func()

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	stop, err := loadApplicationServices(ctx)
	if err != nil {
		panic(err)
	}
	defer stop()
	defer cancel()

	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-s
}

func loadApplicationServices(ctx context.Context) (stopFunc, error) {
	lis, err := net.Listen("tcp", os.Getenv("PORT_DOMAIN_GRPC_HOST"))
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()

	db := persistance.NewKvDB()

	portRepository := repository.NewPort(db)

	portUsecase := usecase.NewPort(portRepository)

	api.RegisterPortDomainServer(grpcServer, handler.NewPortServer(portUsecase))

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	return func() {
		grpcServer.Stop()
	}, nil
}
