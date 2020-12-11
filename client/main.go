package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/johnnywidth/9ty/api"

	"github.com/johnnywidth/9ty/client/handler"
	"github.com/johnnywidth/9ty/client/service"
	"github.com/johnnywidth/9ty/client/usecase"
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
	grpcConn, err := grpc.Dial(os.Getenv("CLIENT_PORT_DOMAIN_GRPC_HOST"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	portClient := api.NewPortDomainClient(grpcConn)

	portDomainService := service.NewPortDomain(portClient)

	portDataUsecase := usecase.NewPort(portDomainService)

	portHandler := handler.NewPort(portDataUsecase)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/port/{name}", portHandler.Get).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    os.Getenv("CLIENT_HTTP_HOST"),
		Handler: handlers.CORS()(router),
	}

	go func() {
		err := loadPortData(ctx, portDataUsecase)
		if err != nil {
			fmt.Println("load port data failed", err)
		}
	}()

	go func() {
		fmt.Println("Starting HTTP listener...")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Server has been stopped: %s\n", err)
		}
	}()

	return func() {
		_ = server.Shutdown(ctx)

		_ = grpcConn.Close()
	}, nil
}

func loadPortData(ctx context.Context, portDataUsecase *usecase.Port) error {
	f, err := os.Open(os.Getenv("CLIENT_PORT_DATA_JSON_FILE"))
	if err != nil {
		return fmt.Errorf("read file failed. %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("error during close file. %s\n", err)
		}
	}()

	r := bufio.NewReader(f)

	portParser := usecase.NewPortJSON(r, portDataUsecase.Create)
	return portParser.Parse(ctx)
}
