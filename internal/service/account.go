package service

import (
	"ahyalfan/e_money_queue/domain"
	"ahyalfan/e_money_queue/internal/config"
	"context"
	"log"

	"github.com/hibiken/asynq"
)

type accountService struct {
	cnf *config.Config
}

func NewAccountService(cnf *config.Config) domain.AccountService {
	return &accountService{cnf: cnf}
}

// GenerateMutation implements domain.AccountService.
func (a *accountService) GenerateMutation() (string, func(context.Context, *asynq.Task) error) {
	return "generate:mutation", func(ctx context.Context, task *asynq.Task) error {
		log.Println("generate:mutation")
		return nil
	}
}
