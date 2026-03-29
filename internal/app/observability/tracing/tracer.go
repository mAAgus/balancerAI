package tracing

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitTracer(serviceName string, jaegerURL string) func(context.Context) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(jaegerURL),
	))
	if err != nil {
		log.Fatalf("failed to init jaeger exporter: %v", err)
	}
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		attribute.String("environment", "dev"),
	)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
