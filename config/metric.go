package config

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

func InitMetric(metricProvider string) {
	provider := metricsdk.NewMeterProvider(metricsdk.WithReader(changeMetricExporter(metricProvider)))
	provider.Meter(APP)
	otel.SetMeterProvider(provider)
}

func GetMetric() metric.MeterProvider {
	return otel.GetMeterProvider()
}

func changeMetricExporter(metricProvider string) metricsdk.Reader {
	switch metricProvider {
	case "prometheus":
		return getPrometheusExporter()
	}
	log.Fatal().Msgf("[config][metric][changeMetricExporter] - Metric exporter not defined")
	return nil
}

func getPrometheusExporter() *prometheus.Exporter {
	exp, err := prometheus.New()
	if err != nil {
		log.Fatal().Msgf("[config][metric][getPrometheusExporter] - Prometheus exporter create error: %v", err)
	}
	return exp
}
