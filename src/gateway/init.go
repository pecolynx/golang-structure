package gateway

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/pecolynx/golang-structure/src/gateway")
