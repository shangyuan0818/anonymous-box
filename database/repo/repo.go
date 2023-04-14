package repo

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("database.repo")
