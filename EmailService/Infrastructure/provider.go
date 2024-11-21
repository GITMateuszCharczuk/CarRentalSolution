package infrastructure

import (
	"email-service/Domain/event"
	smtp "email-service/Infrastructure/MailHog"
	"email-service/Infrastructure/config"
	"email-service/Infrastructure/publisher"
	"email-service/Infrastructure/queue"
	"email-service/Infrastructure/receiver"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	EventPublisher event.EventPublisher
	EventReceiver  event.EventReceiver
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		config.WireSet,
		queue.WireSet,
		smtp.WireSet,
		receiver.WireSet,
		publisher.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}
