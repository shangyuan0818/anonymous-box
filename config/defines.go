package config

import (
	"fmt"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("config")

const (
	Namespace = "anonymous-box-saas"
	Version   = "0.0.1"
)

var ServiceInstanceID = fmt.Sprintf("%s-%s-%s", Namespace, Version, uuid.NewString())
