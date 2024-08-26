package usecase

import (
	"context"

	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/dto"
	"github.com/adityaeka26/go-codebase/internal/repository"
	"github.com/adityaeka26/go-pkg/elasticsearch"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	"go.uber.org/zap"
)

type exampleUsecase struct {
	logger            *logger.Logger
	config            *config.EnvConfig
	postgres          *postgres.Postgres
	elastic           *elasticsearch.Elasticsearch
	exampleRepository repository.ExampleRepository
}

func NewExampleUsecase(logger *logger.Logger, config *config.EnvConfig, postgres *postgres.Postgres, elastic *elasticsearch.Elasticsearch, exampleRepository repository.ExampleRepository) ExampleUsecase {
	return &exampleUsecase{
		logger:            logger,
		config:            config,
		postgres:          postgres,
		elastic:           elastic,
		exampleRepository: exampleRepository,
	}
}

func (u *exampleUsecase) Example(ctx context.Context, req dto.ExampleRequest) (*dto.ExampleResponse, error) {
	logger := u.logger.GetLog().With(zap.String("operationName", "exampleUsecase.Example"))

	// do all business logic here
	result, err := u.exampleRepository.Example(ctx, req.Id)
	if err != nil {
		logger.Error("get example fail", zap.Error(err))
		return nil, pkgError.InternalServerError("internal server error")
	}

	logger.Info("get example success")
	return &dto.ExampleResponse{
		Id:   result.Id,
		Name: result.Name,
	}, nil
}
