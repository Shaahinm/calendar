package main

import (
	"context"
	"fmt"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/internal/db"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {
	ctx := context.Background()
	exporter, err := otlptracegrpc.New(ctx)
	//handleErr(err, "Failed to create exporter")
	println(err)
	openTelemetryURL := attribute.KeyValue{
		Key:   attribute.Key("opentelemetry.io/schemas"),
		Value: attribute.StringValue("1.7.0"),
	}

	resource, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("service-appointments"),
			openTelemetryURL,
		),
	)
	//handleErr(err, "Failed to create resource")
	println(err)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	db.Up()
	api.Init()
	fmt.Println("Application started")
}
