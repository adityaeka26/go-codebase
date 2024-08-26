package rest_handler

import (
	"github.com/adityaeka26/go-codebase/config"
	"github.com/adityaeka26/go-codebase/internal/dto"
	"github.com/adityaeka26/go-codebase/internal/middleware"
	"github.com/adityaeka26/go-codebase/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/helper"
	pkgValidator "github.com/adityaeka26/go-pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type restHandler struct {
	exampleUsecase usecase.ExampleUsecase
	validator      *pkgValidator.XValidator
}

func InitRestHandler(
	app *fiber.App,
	config *config.EnvConfig,
	middleware middleware.Middleware,
	exampleUsecase usecase.ExampleUsecase,
	validator *pkgValidator.XValidator,
) {
	handler := &restHandler{
		exampleUsecase: exampleUsecase,
		validator:      validator,
	}

	app.Get("/example/:id", handler.Example)
}

func (h *restHandler) Example(c *fiber.Ctx) error {
	req := &dto.ExampleRequest{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	resp, err := h.exampleUsecase.Example(c.Context(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, resp, "get example success")
}
