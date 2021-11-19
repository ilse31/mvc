package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `method=${method}, url=${uri}, status=${status} latency=${latency_human} ip=${remote_ip}` + "\n",
	}

	e.Use(middleware.LoggerWithConfig(logger))
}
