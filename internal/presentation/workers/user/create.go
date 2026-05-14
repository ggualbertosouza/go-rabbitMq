package userWorkers

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	helpersWorkers "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/workers/helpers"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserWorker struct {
	Logger logger.Logger
}

func NewCreateUser(log logger.Logger) *CreateUserWorker {
	return &CreateUserWorker{Logger: log}
}

func (u *CreateUserWorker) Execute(payload []byte) error {
	data, err := helpersWorkers.ParseMessage[CreateUserRequest](payload)
	if err != nil {
		return err
	}

	u.Logger.Info("processing message", logger.String("email", data.Email))

	u.Logger.Info("user created")

	return nil
}
