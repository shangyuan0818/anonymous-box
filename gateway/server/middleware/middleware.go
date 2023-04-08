package middleware

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("services.gateway.server.middleware")
