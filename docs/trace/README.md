# Traces, Metrics, Logs

初始化Tracer
```go
var Tracer oteltrace.Tracer

func InitTracer() error {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))
	Tracer = otel.GetTracerProvider().Tracer(version.ServiceName, oteltrace.WithInstrumentationVersion("0.1"))
	return nil
}
```

http中间件: [go-restful 中间件](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/emicklei/go-restful)
```go
// create the Otel filter
filter := otelrestful.OTelFilter("my-service")
// use it
restful.DefaultContainer.Filter(filter)
```

grpc中间件: [grpc 中间件](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/google.golang.org/grpc/otelgrpc)
```go
s := grpc.NewServer(
    grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
    grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
)
```

数据库中间件: [mongodb](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo)
```go
addr := "mongodb://localhost:27017/?connect=direct"
opts := options.Client()
opts.Monitor = otelmongo.NewMonitor(
    otelmongo.WithTracerProvider(provider),
    otelmongo.WithCommandAttributeDisabled(tc.excludeCommand),
)
```

http client: [net/http](https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/net/http/httptrace/otelhttptrace/example/client/client.go)
```go
client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

ctx, span := tr.Start(ctx, "say hello", trace.WithAttributes(semconv.PeerService("ExampleService")))
defer span.End()

ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx))
req, _ := http.NewRequestWithContext(ctx, "GET", *url, nil)
```

代码中埋点
```go
_, span := tracer.Start(req.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", uid)))
defer span.End()
```

[jaeger exporter配置](https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters)
```go
import (
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	
	otel.SetTracerProvider(tp)
	return tp, nil
}
```

## 部署Jaeger

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.45
```

访问: http://localhost:16686 查看UI界面

## 参考

+ [opentelemetry 官网](https://opentelemetry.io/)
+ [opentelemetry 文档](https://opentelemetry.io/docs/)
+ [opentelemetry SDK](https://opentelemetry.io/docs/instrumentation/)
+ [opentelemetry 中间件](https://opentelemetry.io/ecosystem/registry)
+ [jaeger 官方文档](https://www.jaegertracing.io/docs/1.45/)