package userUseCase

import (
	"encoding/json"

	"github.com/ggualbertosouza/go-rabbitMq/internal/application/inputs"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq"
	userRbMq "github.com/ggualbertosouza/go-rabbitMq/internal/infra/rabbitmq/user"
)

type UserUseCase struct {
	Logger       logger.Logger
	UserRabbitMq *userRbMq.UserRabbitMq
}

func New(log logger.Logger, rabbit *rabbitmq.RabbitMq) *UserUseCase {
	userrbmq := userRbMq.New(log, rabbit.Channel)

	return &UserUseCase{
		Logger:       log,
		UserRabbitMq: userrbmq,
	}
}

func (u *UserUseCase) CreateUser(email, password string) error {
	u.Logger.Info("creating user")

	user := userRbMq.UserCreatedMessage{
		Email:    email,
		Password: password,
	}

	u.Logger.Info("parsing user", logger.String("email", email))
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = u.UserRabbitMq.Publish(userRbMq.UsersExchanges, userRbMq.UserCreatedRK, body)
	if err != nil {
		return err
	}

	u.Logger.Info("user created successfully")
	return nil
}

func (u *UserUseCase) SendCreatedUserEmail(user inputs.UserSendEmail) error {
	u.Logger.Info("sending email to user", logger.String("email", user.Email))

	return nil
}
