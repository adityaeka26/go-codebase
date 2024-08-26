package middleware

import (
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/helper"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func (m *middleware) ValidateToken(jwtPublicKey string) fiber.Handler {
	logger := m.logger.GetLog().With(zap.String("operationName", "middleware.ValidateToken"))

	return func(c *fiber.Ctx) error {
		jwtPublicKey = strings.ReplaceAll(jwtPublicKey, "\\n", "\n")
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(jwtPublicKey))
		if err != nil {
			logger.Error("parse rsa public key fail", zap.Error(err))
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if len(token) <= 0 {
			logger.Warn("token not found or invalid")
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err != nil {
			logger.Error("parse token fail", zap.Error(err))
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}
		claims := parseToken.Claims.(jwt.MapClaims)

		c.Locals("id", claims["data"].(map[string]any)["id"])
		c.Locals("username", claims["data"].(map[string]any)["username"])
		c.Locals("role", claims["data"].(map[string]any)["role"])

		return c.Next()
	}
}

func (m *middleware) ValidateRoles(roles []string) fiber.Handler {
	logger := m.logger.GetLog().With(zap.String("operationName", "middleware.ValidateRoles"))

	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			logger.Warn("role not found")
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		if !helper.Contains(roles, role) {
			logger.Warn("role not allowed")
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		return c.Next()
	}
}
