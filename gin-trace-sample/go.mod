module gin-trace-sample

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	go.opentelemetry.io/otel v1.11.0
	go.opentelemetry.io/otel/exporters/jaeger v1.11.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.0
	go.opentelemetry.io/otel/exporters/zipkin v1.11.0
	go.opentelemetry.io/otel/sdk v1.11.0
	go.opentelemetry.io/otel/trace v1.11.0
	google.golang.org/grpc v1.50.1
)
