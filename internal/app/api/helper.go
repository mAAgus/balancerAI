package api

import (
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLogger() {
	level, err := logrus.ParseLevel(a.config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	a.logger.SetLevel(level)
}
func (a *API) configureMiddleware() {
	a.app.Use(recover.New())
	a.app.Use(logger.New())
}
func (a *API) configureRouters() {
	api := a.app.Group("/api/v1")
	api.Post("/generate", a.handleGenerate)
	api.Get("/health", a.handleHealth)
	a.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}
