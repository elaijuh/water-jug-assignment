package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/elaijuh/water-jug/config"
	"github.com/elaijuh/water-jug/controller"
)

var cfg config.Properties

func init() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read : %v", err)
	}
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("x-auth-token") // 123456
		if token == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "No x-auth-token")
		}
		if token != "123456" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not authorized")
		}
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${user_agent} ` +
			`${status} ${error} ${latency_human}` + "\n",
	}))
	h := &controller.ProblemHandler{}
	e.POST("/problem", h.ResolveProblem, auth)

	e.Logger.Infof("Listening on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
