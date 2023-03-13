package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func InitTrace(traceProvider string) {
	tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(changeTraceExporter(traceProvider)),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(APP),
		)),
	)
}

func GetTrace() trace.Tracer {
	return otel.GetTracerProvider().Tracer(APP)
}

func changeTraceExporter(traceProvider string) tracesdk.SpanExporter {
	switch traceProvider {
	case "jaeger":
		return getJaegerExporter()
	}
	log.Fatal().Msgf("[config][trace][changeTraceExporter] - Trace exporter not defined")
	return nil
}

func getJaegerExporter() *jaeger.Exporter {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(viper.GetString("TRACE_ENDPOINT"))),
	)
	if err != nil {
		log.Fatal().Msgf("[config][trace][getJaegerExporter] - Jaeger exporter create error: %v", err)
	}
	return exp
}
