package telemetry

import (
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Telemetry struct {
	TelemetryTracer *trace.TracerProvider
	MetricsProvider *metric.MeterProvider
}

func NewTelemetry(service string) (*Telemetry, error) {
	t, err := NewOpenTelemetryTracer(service)
	if err != nil {
		return nil, err
	}

	m, err := prometheusMetricProvider()
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		TelemetryTracer: t,
		MetricsProvider: m,
	}, nil
}
