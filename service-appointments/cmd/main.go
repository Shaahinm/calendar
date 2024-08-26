package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Shaahinm/calendar/api"
	"github.com/Shaahinm/calendar/internal/db"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {
	ctx := context.Background()
	//exporter, err := otlptracegrpc.New(ctx)
	//handleErr(err, "Failed to create exporter")
	//println(err)
	// openTelemetryURL := attribute.KeyValue{
	// 	Key:   attribute.Key("opentelemetry.io/schemas"),
	// 	Value: attribute.StringValue("1.7.0"),
	// }

	// resource, err := resource.New(ctx,
	// 	resource.WithAttributes(
	// 		semconv.ServiceNameKey.String("service-appointments"),
	// 		openTelemetryURL,
	// 	),
	// )
	otelAgentAddr := "0.0.0.0:4317"
	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(otelAgentAddr))
	if err != nil {
		fmt.Println(err, "Failed to create the collector metric exporter")
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String("service-appointments"),
		),
	)
	//handleErr(err, "Failed to create resource")
	println(err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExp,
				sdkmetric.WithInterval(2*time.Second),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelAgentAddr))
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		fmt.Println(err, "Failed to create the collector trace exporter")
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	db.Up()
	api.Init()
	fmt.Println("Application started")
}
