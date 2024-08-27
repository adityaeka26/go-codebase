package rest_handler

import (
	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/dto"
	"github.com/adityaeka26/go-codebase/internal/middleware"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/helper/echo_wrapper"
	pkgValidator "github.com/adityaeka26/go-pkg/validator"
	"github.com/labstack/echo/v4"
)

type restHandler struct {
	exampleUsecase usecase.ExampleUsecase
	validator      *pkgValidator.XValidator
}

func InitRestHandler(
	app *echo.Echo,
	config *config.EnvConfig,
	middleware middleware.Middleware,
	exampleUsecase usecase.ExampleUsecase,
	validator *pkgValidator.XValidator,
) {
	handler := &restHandler{
		exampleUsecase: exampleUsecase,
		validator:      validator,
	}

	app.GET("/example/:id", handler.Example)
}

func (h *restHandler) Example(c echo.Context) error {
	req := &dto.ExampleRequest{}
	if err := c.Bind(req); err != nil {
		return echo_wrapper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return echo_wrapper.RespError(c, err)
	}

	resp, err := h.exampleUsecase.Example(c.Request().Context(), *req)
	if err != nil {
		return echo_wrapper.RespError(c, err)
	}
	return echo_wrapper.RespSuccess(c, resp, "get example success")
}
