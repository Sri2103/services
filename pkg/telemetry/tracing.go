package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// set up tracing exporter
func newTracerConsoleExport() (sdktrace.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func newTracingJagerExport(url string) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(context.Background(), otlptracehttp.WithEndpointURL(url))
}

// create a new TracerProvider with the console exporter
func NewOpenTelemetryTracer(service string) (*sdktrace.TracerProvider, error) {
	jaegerUrl := "http://localhost:4318" // default Jaeger endpoint for local development
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
		),
	)
	if err != nil {
		return nil, err
	}
	// exp, err := newTracerConsoleExport()
	exp, err := newTracingJagerExport(jaegerUrl)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithResource(r), sdktrace.WithBatcher(exp))
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil

}
