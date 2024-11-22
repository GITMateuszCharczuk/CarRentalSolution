package infrastructure

import (
	"email-service/Domain/event"
	"email-service/Infrastructure/config"
	fetcher "email-service/Infrastructure/data_fetcher"
	smtp "email-service/Infrastructure/email_sender"
	publisher "email-service/Infrastructure/event_publisher"
	receiver "email-service/Infrastructure/event_receiver"
	"email-service/Infrastructure/queue"

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
		fetcher.WireSet,
		receiver.WireSet,
		publisher.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}
