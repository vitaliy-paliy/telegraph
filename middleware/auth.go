package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("Authorization")
		// Remove Bearer prefix.
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate JWT.
		result, err := JwtValidate(token)
		if err != nil || !result.Valid {
			return c.JSON(http.StatusForbidden, "Invalid token.")
		}

		claims, _ := result.Claims.(*JWTAuthClaims)
		ctx := context.WithValue(c.Request().Context(), "auth", claims)
		c.SetRequest(c.Request().WithContext(ctx))

		next(c)
		return nil
	}
}

func GetToken(ctx context.Context) *JWTAuthClaims {
	raw, _ := ctx.Value("auth").(*JWTAuthClaims)
	return raw
}
