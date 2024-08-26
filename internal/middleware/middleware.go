package middleware

import (
	"github.com/adityaeka26/go-pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	logger *logger.Logger
}

type Middleware interface {
	ValidateToken(jwtPublicKey string) fiber.Handler
	ValidateRoles(roles []string) fiber.Handler
}

func NewMiddleware(logger *logger.Logger) Middleware {
	return &middleware{
		logger: logger,
	}
}
