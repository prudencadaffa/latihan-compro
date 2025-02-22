package middleware

import (
	"latihan-compro/config"
	"latihan-compro/internal/adapter/handler/response"
	"latihan-compro/utils/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Middleware interface {
	CheckToken() echo.MiddlewareFunc
}

type Options struct {
	authJwt auth.JwtInterface
}

// CheckToken implements Middleware.
func (o *Options) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var errorResponse response.ErrorResponseDefault

			// Ambil header Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				errorResponse.Meta.Status = false
				errorResponse.Meta.Message = "Missing Authorization header"
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}

			// Validasi format header Authorization
			parts := strings.Split(authHeader, "Bearer ")
			if len(parts) != 2 {
				errorResponse.Meta.Status = false
				errorResponse.Meta.Message = "Invalid Authorization header format"
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}

			// Ambil token dari header
			tokenString := parts[1]

			// Verifikasi token
			claims, err := o.authJwt.VerifyAccessToken(tokenString)
			if err != nil {
				errorResponse.Meta.Status = false
				errorResponse.Meta.Message = "Invalid token"
				return c.JSON(http.StatusUnauthorized, errorResponse)
			}

			// Simpan claims ke context
			c.Set("user", claims)

			// Lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)

	return opt
}
