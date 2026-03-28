package api

import (
	"balancerAI/internal/app/models"
	metrics "balancerAI/internal/app/observability/metrics"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

/*
func (a *API) handleGenerate(c *fiber.Ctx) error {

		start := time.Now()

		var req models.GenerateRequest

		if err := c.BodyParser(&req); err != nil {

			a.logger.WithError(err).Warn("invalid request body")

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request",
			})
		}

		logger := a.logger.WithFields(logrus.Fields{
			"type": req.Type,
		})

		logger.Info("generate request received")

		metrics.RequestTotal.WithLabelValues(string(req.Type)).Inc()

		result, err := a.service.Generate(c.UserContext(), req)

		duration := time.Since(start)

		metrics.RequestDuration.
			WithLabelValues(string(req.Type)).
			Observe(duration.Seconds())

		if err != nil {

			metrics.RequestErrors.WithLabelValues(string(req.Type)).Inc()

			logger.WithError(err).Error("generation failed")

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		logger.WithField("duration", duration.Seconds()).
			Info("generation finished")

		return c.JSON(models.GenerateResponse{
			Result: result,
		})
	}
*/
func (a *API) handleGenerate(c *fiber.Ctx) error {

	ctx := c.UserContext()
	tracer := otel.Tracer("ai-balancer")

	ctx, span := tracer.Start(ctx, "handleGenerate")
	defer span.End()

	start := time.Now()

	var req models.GenerateRequest
	if err := c.BodyParser(&req); err != nil {
		a.logger.WithError(err).Warn("invalid request body")
		span.RecordError(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	logger := a.logger.WithFields(logrus.Fields{
		"type": req.Type,
	})

	logger.Info("generate request received")
	metrics.RequestTotal.WithLabelValues(string(req.Type)).Inc()

	span.SetAttributes(attribute.String("request.type", string(req.Type)))

	result, err := a.service.Generate(ctx, req)

	duration := time.Since(start)
	metrics.RequestDuration.WithLabelValues(string(req.Type)).Observe(duration.Seconds())

	if err != nil {
		metrics.RequestErrors.WithLabelValues(string(req.Type)).Inc()
		logger.WithError(err).Error("generation failed")
		span.RecordError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.WithField("duration", duration.Seconds()).Info("generation finished")
	span.SetAttributes(attribute.Float64("request.duration", duration.Seconds()))

	return c.JSON(models.GenerateResponse{
		Result: result,
	})
}
func (a *API) handleHealth(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
