package api

import (
	"balancerAI/internal/app/config"
	"balancerAI/internal/app/models"
	"balancerAI/internal/app/observability/logging"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Generate(ctx context.Context, req models.GenerateRequest) (any, error)
}
type API struct {
	config  *config.Config
	logger  *logrus.Logger
	app     *fiber.App
	service Service
}

func New(config *config.Config, service Service) *API {
	logger := logging.NewLogger(config.Logger.Level)
	app := fiber.New()
	return &API{
		config:  config,
		logger:  logger,
		app:     app,
		service: service,
	}
}
func (a *API) Start() error {

	a.configureMiddleware()
	a.configureRouters()
	a.logger.WithFields(logrus.Fields{
		"port": a.config.Server.BindAddr,
	}).Info("starting api server")
	return a.app.Listen(a.config.Server.BindAddr)
}
