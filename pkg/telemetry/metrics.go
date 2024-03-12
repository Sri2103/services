package telemetry

import (
	"encoding/json"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// implement metrics for the open telemetry

type OpenMetrics struct {
}

func newResource(service string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(service),
			semconv.ServiceVersion("0.1.0"),
		))
}

func newMetricConsoleProvider(service string) (*metric.MeterProvider, error) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	metricExporter, err := stdoutmetric.New(
		stdoutmetric.WithEncoder(enc),
		stdoutmetric.WithoutTimestamps(),
	)
	if err != nil {
		return nil, err
	}
	res, err := newResource(service)
	if err != nil {
		return nil, err
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)
	otel.SetMeterProvider(meterProvider)
	return meterProvider, nil
}

func prometheusMetricProvider() (*metric.MeterProvider, error) {
	exp, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exp),
	)
	otel.SetMeterProvider(provider)

	return provider, nil

}
