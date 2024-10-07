package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"prom/pkg/jaeger"
	"time"

	"github.com/opentracing/opentracing-go/log"
)

func main() {
	j := jaeger.New(config)
	tracer, closer, err := j.NewTracer()
	if err != nil {
		fmt.Errorf("cannot create tracer: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	ctx := context.Background()

	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "parent-operation")
	defer parentSpan.Finish()

	firstOperation(ctx, tracer)

	secondOperation(ctx, tracer)
}

func firstOperation(ctx context.Context, tracer opentracing.Tracer) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "first-operation")
	defer span.Finish()

	span.LogFields(log.String("info", "First operation started"))

	time.Sleep(1 * time.Second)

	span.LogFields(log.String("info", "First operation finished"))
}

func secondOperation(ctx context.Context, tracer opentracing.Tracer) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "second-operation")
	defer span.Finish()

	span.LogFields(log.String("info", "Second operation started"))

	time.Sleep(1 * time.Second)

	span.LogFields(log.String("info", "Second operation finished"))
}
