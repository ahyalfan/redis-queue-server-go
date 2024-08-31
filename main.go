package main

import (
	"ahyalfan/e_money_queue/internal/config"
	"ahyalfan/e_money_queue/internal/service"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	cnf := config.Get()

	// server config
	redisConnection := asynq.RedisClientOpt{
		Addr:     cnf.Redis.Addr,
		Password: cnf.Redis.Password,
	}

	email := service.NewEmailService(cnf)
	account := service.NewAccountService(cnf)

	// bikin server
	worker := asynq.NewServer(redisConnection, asynq.Config{
		Concurrency: 4, // kita contohkan 4 saja, rekomendasi sesuaikan dengan cpu
	})
	// server handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(email.SendEmailQueue())
	mux.HandleFunc(account.GenerateMutation())

	if err := worker.Run(mux); err != nil {
		log.Fatal(err.Error())
	}
}
