package cmd

import (
	"strings"

	"github.com/adityaeka26/go-codebase/cmd/grpc"
	"github.com/adityaeka26/go-codebase/cmd/rest"
	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/middleware"
	"github.com/adityaeka26/go-codebase/internal/repository"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	"github.com/adityaeka26/go-pkg/elasticsearch"
	pkgKafka "github.com/adityaeka26/go-pkg/kafka"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	"go.uber.org/zap"
)

func Execute() {
	logger := logger.NewLogger()

	config, err := config.Load(".env")
	if err != nil {
		logger.GetLog().Error("load config fail", zap.Error(err))
		panic(err)
	}

	postgres, err := postgres.NewPostgres(
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDb,
		config.PostgresSslEnabled,
	)
	if err != nil {
		logger.GetLog().Error("init postgres fail", zap.Error(err))
		panic(err)
	}

	_, err = pkgKafka.NewKafkaProducer(
		config.KafkaSasl,
		config.KafkaHosts,
		config.KafkaUsername,
		config.KafkaPassword,
	)
	if err != nil {
		logger.GetLog().Error("init kafka fail", zap.Error(err))
		panic(err)
	}

	elasticsearchAddresses := strings.Split(config.ElasticsearchHost, ",")
	elastic, err := elasticsearch.NewElasticsearch(config.ElasticsearchUsername, config.ElasticsearchPassword, elasticsearchAddresses)
	if err != nil {
		logger.GetLog().Error("init elasticsearch fail", zap.Error(err))
		panic(err)
	}

	exampleRepository, err := repository.NewExampleRepository(config)
	if err != nil {
		logger.GetLog().Error("init user repository fail", zap.Error(err))
		panic(err)
	}

	middleware := middleware.NewMiddleware(logger)

	exampleUsecase := usecase.NewExampleUsecase(logger, config, postgres, elastic, exampleRepository)

	go func() {
		err = grpc.ServeGRPC(config, exampleUsecase)
		if err != nil {
			logger.GetLog().Error("serve grpc fail", zap.Error(err))
			panic(err)
		}
	}()

	err = rest.ServeREST(logger, config, middleware, postgres, exampleUsecase)
	if err != nil {
		logger.GetLog().Error("serve rest fail", zap.Error(err))
		panic(err)
	}
}
