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

	span, ctx := opentracing.StartSpanFromContext(ctx, "request-handler")
	defer span.Finish()

	span.SetBaggageItem("user_id", "12345")

	firstOp(ctx, tracer)

	secondOp(ctx, tracer)
}

func firstOp(ctx context.Context, tracer opentracing.Tracer) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "first-operation")
	defer span.Finish()

	userID := span.BaggageItem("user_id")
	span.LogFields(
		log.String("event", "processing"),
		log.String("user_id", userID),
	)

	fmt.Println("First Operation - User ID from Baggage:", userID)

	time.Sleep(1 * time.Second)
}

func secondOp(ctx context.Context, tracer opentracing.Tracer) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "second-operation")
	defer span.Finish()

	userID := span.BaggageItem("user_id")
	span.LogFields(
		log.String("event", "processing"),
		log.String("user_id", userID),
	)

	fmt.Println("Second Operation - User ID from Baggage:", userID)

	time.Sleep(1 * time.Second)
}
