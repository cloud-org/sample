package main

import (
	"context"
	"database/sql"
	"gin-trace-sample/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var clientAttr = attribute.Key("gin.client")

func startSpan(ctx context.Context, method string) (context.Context, oteltrace.Span) {
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	start, span := tracer.Start(ctx,
		"httpclient",
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
	)
	span.SetAttributes(clientAttr.String(method))

	return start, span
}

func endSpan(span oteltrace.Span, err error) {
	defer span.End()

	if err == nil || err == sql.ErrNoRows {
		span.SetStatus(codes.Ok, "ok msg")
		return
	}

	span.SetStatus(codes.Error, err.Error())
	span.RecordError(err)
}
