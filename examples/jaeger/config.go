package main

import "prom/pkg/jaeger"

var config = jaeger.Config{
	Host:        "localhost:6831", // container: container-name:6831
	ServiceName: "go-monitoring-profiling",
	LogSpans:    false,
}
