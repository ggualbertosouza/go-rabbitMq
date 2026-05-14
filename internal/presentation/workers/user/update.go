package userWorkers

import (
	"github.com/ggualbertosouza/go-rabbitMq/internal/infra/config/logger"
	helpersWorkers "github.com/ggualbertosouza/go-rabbitMq/internal/presentation/workers/helpers"
)

type UpdateUserRequest struct {
	Email string `json:"email"`
}

type UpdateUserWorker struct {
	Logger logger.Logger
}

func NewUpdateUser(log logger.Logger) *UpdateUserWorker {
	return &UpdateUserWorker{Logger: log}
}

func (u *UpdateUserWorker) Execute(payload []byte) error {
	data, err := helpersWorkers.ParseMessage[UpdateUserRequest](payload)
	if err != nil {
		return err
	}

	u.Logger.Info("processing message", logger.String("email", data.Email))

	u.Logger.Info("user updated")

	return nil
}
