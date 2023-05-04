package tracing

import (
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	defaultProvider = "otel"
)

var (
	providers = make(map[string]func() (tracesdk.SpanExporter, error))
)

func Register(name string, provider func() (tracesdk.SpanExporter, error)) {
	providers[name] = provider
}

func Init() error {
	var provider = defaultProvider
	fn, ok := providers[provider]
	if !ok {
		return errors.New("tracing provider not found")
	}

	// 创建 SpanExporter
	spanExporter, err := fn()
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(

		// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
		// In a production application, use sdktrace.ProbabilitySampler with a desired probability.

		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(1)),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(""),
			// attribute.Int64("ID", id),
		)),
		sdktrace.WithBatcher(spanExporter),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagator())
	// todo graceful teardown
	return nil
}

func propagator() propagation.TextMapPropagator {
	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	return propagator
}
