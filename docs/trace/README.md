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

代码中埋点
```go
_, span := tracer.Start(req.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", uid)))
defer span.End()
```

## 参考

+ [opentelemetry 官网](https://opentelemetry.io/)
+ [opentelemetry 文档](https://opentelemetry.io/docs/)
+ [opentelemetry SDK](https://opentelemetry.io/docs/instrumentation/)
+ [opentelemetry 中间件](https://opentelemetry.io/ecosystem/registry)
+ [go-restful 中间件](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/emicklei/go-restful)
+ [grpc 中间件](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/google.golang.org/grpc/otelgrpc)