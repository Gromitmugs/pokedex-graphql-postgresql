package graph

import "go.opentelemetry.io/otel/trace"

type Resolver struct {
	DB     Database
	Tracer trace.Tracer
}
