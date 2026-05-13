package userWorkers

import (
	"encoding/json"

	"github.com/ggualbertosouza/go-rabbitMq/internal/application/inputs"
	userUseCase "github.com/ggualbertosouza/go-rabbitMq/internal/application/useCases/user"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/rabbitmq/amqp091-go"
)

type UserWorker struct {
	Logger      logger.Logger
	Channel     <-chan amqp091.Delivery
	UserUseCase *userUseCase.UserUseCase
}

func New(
	log logger.Logger,
	channel <-chan amqp091.Delivery,
	userUc *userUseCase.UserUseCase,
) *UserWorker {
	return &UserWorker{
		Logger:      log,
		Channel:     channel,
		UserUseCase: userUc,
	}
}

func (u *UserWorker) Init() {
	log := u.Logger

	for message := range u.Channel {
		log.Info("message received")

		payload, err := u.parseMessage(message.Body)
		if err != nil {
			log.Error("failed to parse message", logger.Error(err))

			err = message.Nack(false, false)
			if err != nil {
				log.Error("failed to nack message", logger.Error(err))
			}

			continue
		}

		err = u.UserUseCase.SendCreatedUserEmail(payload)
		if err != nil {
			log.Error("failed to process message", logger.Error(err))

			err = message.Nack(false, true)
			if err != nil {
				log.Error("failed to nack message", logger.Error(err))
			}

			continue
		}

		err = message.Ack(false)
		if err != nil {
			log.Error("failed to ack message", logger.Error(err))
		}
	}
}

func (u *UserWorker) parseMessage(
	body []byte,
) (inputs.UserSendEmail, error) {
	var user inputs.UserSendEmail

	err := json.Unmarshal(body, &user)
	if err != nil {
		return inputs.UserSendEmail{}, err
	}

	return user, nil
}