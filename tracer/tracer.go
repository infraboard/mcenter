package tracer

import (
	"context"
	"os"

	"github.com/infraboard/mcenter/version"
	"github.com/infraboard/mcube/logger/zap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var Tracer oteltrace.Tracer

func InitTracer() error {
	ep := os.Getenv("JAEGER_ENDPINT")

	var (
		exporter sdktrace.SpanExporter
		err      error
	)
	if ep != "" {
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(ep)))
	} else {
		exporter, err = stdout.New(stdout.WithPrettyPrint())
	}
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))
	Tracer = otel.GetTracerProvider().Tracer(version.ServiceName, oteltrace.WithInstrumentationVersion("0.1"))
	return nil
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	resources, err := resource.New(
		context.Background(),
		resource.WithSchemaURL(semconv.SchemaURL), // shema url
		resource.WithAttributes( // specify resource attributes
			semconv.ServiceNameKey.String(version.ServiceName),
			semconv.ServiceVersionKey.String(version.Short())),
		resource.WithFromEnv(),   // pull attributes from OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables
		resource.WithProcess(),   // This option configures a set of Detectors that discover process information
		resource.WithOS(),        // This option configures a set of Detectors that discover OS information
		resource.WithContainer(), // This option configures a set of Detectors that discover container information
		resource.WithHost(),      // This option configures a set of Detectors that discover host information
	)
	if err != nil {
		zap.L().Errorf(err.Error())
	}

	return resources
}
