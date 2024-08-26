package grpc

import (
	"fmt"
	"net"

	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/handler/grpc_handler"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServeGRPC(
	config *config.EnvConfig,
	exampleUsecase usecase.ExampleUsecase,
) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", config.GrpcPort))
	if err != nil {
		return err
	}

	app := grpc.NewServer()

	grpc_handler.InitGrpcHandler(
		app,
		exampleUsecase,
	)
	reflection.Register(app)

	fmt.Println("running grpc on port", config.GrpcPort)
	err = app.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
