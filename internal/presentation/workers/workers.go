package workers

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
	userRbMq "github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq/user"
	userWorkers "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/workers/user"
	"github.com/rabbitmq/amqp091-go"
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

	ch2, err := w.Rabbit.Consume(string(userRbMq.UserQueue2), "userConsumerQueue2")
	if err != nil {
		return err
	}

	createWorker := userWorkers.NewCreateUser(w.Logger)
	updateWorker := userWorkers.NewUpdateUser(w.Logger)

	go ConsumeWorker(
		w.Logger,
		ch,
		"create-user-worker",
		createWorker.Execute,
	)

	go ConsumeWorker(
		w.Logger,
		ch2,
		"update-user-worker",
		updateWorker.Execute,
	)

	select {}
}

func ConsumeWorker(
	log logger.Logger,
	channel <-chan amqp091.Delivery,
	workerName string,
	handler func([]byte) error,
) {
	for message := range channel {

		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Error("panic recovered")

					err := message.Nack(false, false)
					if err != nil {
						log.Error("failed to nack message")
					}
				}
			}()

			log.Info("message received",
				logger.String("Worker", workerName),
			)

			err := handler(message.Body)
			if err != nil {

				log.Error("worker execution failed",
					logger.Error(err),
				)

				err = message.Nack(false, false)
				if err != nil {
					log.Error("failed to nack message")
				}

				return
			}

			err = message.Ack(false)
			if err != nil {
				log.Error("failed to ack message",
					logger.Error(err),
				)

				return
			}

			log.Info("message acked")
		}()
	}
}
