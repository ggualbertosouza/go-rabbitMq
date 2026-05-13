package workers

import (
	userUseCase "github.com/ggualbertosouza/go-rabbitMq/internal/application/useCases/user"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
	userRbMq "github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq/user"
	userWorkers "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/workers/user"
)

type Workers struct {
	Logger logger.Logger
	Rabbit *rabbitmq.RabbitMq
}

func New(log logger.Logger, rb *rabbitmq.RabbitMq) *Workers {
	return &Workers{
		Logger: log,
		Rabbit: rb,
	}
}

func (w *Workers) Init() error {
	ch, err := w.Rabbit.Consume(string(userRbMq.UserQueue1), "userConsumerQueue1")
	if err != nil {
		return err
	}

	userUc := userUseCase.New(w.Logger, w.Rabbit)
	userWk := userWorkers.New(w.Logger, ch, userUc)

	userWk.Init()
	select {}
}
