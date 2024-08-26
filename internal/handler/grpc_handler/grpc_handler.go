package grpc_handler

import (
	"context"

	"github.com/adityaeka26/go-codebase/internal/dto"
	pb "github.com/adityaeka26/go-codebase/internal/handler/grpc_handler/proto"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedExampleServiceServer
	exampleUsecase usecase.ExampleUsecase
}

func InitGrpcHandler(
	app *grpc.Server,
	exampleUsecase usecase.ExampleUsecase,
) {
	pb.RegisterExampleServiceServer(app, &grpcHandler{
		exampleUsecase: exampleUsecase,
	})
}

func (g *grpcHandler) GetAuthorDetail(ctx context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	response, err := g.exampleUsecase.Example(ctx, dto.ExampleRequest{
		Id: int(req.Id),
	})
	if err != nil {
		return nil, pkgError.GrpcError(err)
	}
	return &pb.ExampleResponse{
		Id:   int64(response.Id),
		Name: response.Name,
	}, nil
}
