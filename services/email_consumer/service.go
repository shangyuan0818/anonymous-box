package email_consumer

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
)

var tracer = otel.Tracer(ServiceName)

const (
	ServiceName = "email-service-consumer"

	MQExchangeName = "email-exchange"
	MQQueueName    = "email-queue"

	MQKeyEmailSend = "email.send"
)

type EmailServiceConsumerImpl struct {
	fx.In
}

func NewEmailServiceConsumer(impl EmailServiceConsumerImpl) api.EmailServiceConsumer {
	return &impl
}
