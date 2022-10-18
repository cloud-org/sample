package middleware

import (
	"gin-trace-sample/trace"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// TracingHandler return a middleware that process the opentelemetry.
func TracingHandler(serviceName string) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		propagator := otel.GetTextMapPropagator()
		tracer := otel.Tracer(trace.TraceName) // 影响的是 trace 中的 otel.library.name 字段

		ctx := propagator.Extract(ginCtx.Request.Context(), propagation.HeaderCarrier(ginCtx.Request.Header))
		spanName := ginCtx.Request.URL.Path
		//log.Println("spanName is", spanName)
		spanCtx, span := tracer.Start(
			ctx,
			spanName,
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(
				serviceName, spanName, ginCtx.Request)...),
		)
		defer span.End()
		//log.Println(span.SpanContext().IsValid(), span.SpanContext().SpanID(), span.SpanContext().TraceID())

		// convenient for tracking error messages 注入返回 response key: Traceparent
		propagator.Inject(spanCtx, propagation.HeaderCarrier(ginCtx.Writer.Header()))

		// 注入 context
		ginCtx.Request = ginCtx.Request.WithContext(spanCtx)

		// 直接通过 Value(0) 无法获取到，需要通过 SpanContextFromContext 进行获取
		//spanValue := spanCtx.Value(0)
		//spanValue := oteltrace.SpanContextFromContext(spanCtx)
		//logx.Infof("span value is %+v", spanValue)
		//printContextInternals(spanCtx, false)

		ginCtx.Next()

		return
	}
}
