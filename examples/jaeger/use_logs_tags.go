package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/opentracing/opentracing-go/log"
	"prom/pkg/jaeger"
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

	span, ctx := opentracing.StartSpanFromContext(ctx, "db-operation")
	defer span.Finish()

	span.SetTag("db.type", "postgresql")
	span.SetTag("db.statement", "SELECT * FROM users WHERE id = ?")

	if err = simulateDBOperation(); err != nil {
		span.LogFields(
			log.String("event", "error"),
			log.String("message", err.Error()),
			log.String("error.kind", "DBConnectionError"),
		)

		span.SetTag("error", true)
	} else {
		span.LogFields(
			log.String("event", "db.query.success"),
			log.String("message", "Query executed successfully"),
		)
	}
}

func simulateDBOperation() error {
	time.Sleep(1 * time.Second)
	return errors.New("failed to connect to database")
}
