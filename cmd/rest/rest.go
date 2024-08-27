package rest

import (
	"fmt"
	"time"

	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/handler/rest_handler"
	"github.com/adityaeka26/go-codebase/internal/middleware"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	graceful_shutdown "github.com/adityaeka26/go-pkg/echo_graceful_shutdown"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	pkgValidator "github.com/adityaeka26/go-pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ServeREST(
	logger *logger.Logger,
	config *config.EnvConfig,
	middleware middleware.Middleware,
	postgres *postgres.Postgres,
	exampleUsecase usecase.ExampleUsecase,
) error {
	gs := graceful_shutdown.GracefulShutdown{
		Timeout:        5 * time.Second,
		GracefulPeriod: time.Duration(config.GracefulPeriod) * time.Second,
	}

	app := echo.New()

	gs.Enable(app)
	gs.Register(logger, postgres)

	rest_handler.InitRestHandler(
		app,
		config,
		middleware,
		exampleUsecase,
		&pkgValidator.XValidator{
			Validator: &validator.Validate{},
		},
	)

	err := app.Start(fmt.Sprintf(":%s", config.RestPort))
	if err != nil {
		return err
	}

	gs.Cleanup()

	return nil
}
