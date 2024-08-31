package service

import (
	"ahyalfan/e_money_queue/domain"
	"ahyalfan/e_money_queue/dto"
	"ahyalfan/e_money_queue/internal/config"
	"context"
	"encoding/json"
	"log"
	"net/smtp"

	"github.com/hibiken/asynq"
)

type emailService struct {
	cnf *config.Config
}

func NewEmailService(cnf *config.Config) domain.EmailService {
	return &emailService{cnf: cnf}
}

// Send implements domain.EmailService.
func (e *emailService) Send(to string, subject string, body string) error {
	// untuk identity kita kasih kosong karena biasanya smtp itu diisi kosongana
	auth := smtp.PlainAuth("", e.cnf.Mail.Username, e.cnf.Mail.Password, e.cnf.Mail.Host)

	// Mempersiapkan pesan email berbentu byte
	msg := []byte("From: simple-pay <" + e.cnf.Mail.Username + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body)
	// kirim email
	if err := smtp.SendMail(e.cnf.Mail.Host+":"+e.cnf.Mail.Port, auth, e.cnf.Mail.Username, []string{to}, msg); err != nil {
		log.Printf("[ERR] email send failed: %s", err.Error())
		return err
	}
	return nil
}

// SendEmailQueue implements domain.EmailService.
func (e *emailService) SendEmailQueue() (string, func(context.Context, *asynq.Task) error) {
	// disini kita return saja nama jobnya yg akan dijadikan pendaftaran di queue nanti, dan lakukan functionnya apa
	return "send:email", func(ctx context.Context, task *asynq.Task) error {
		var data dto.EmailSendReq
		_ = json.Unmarshal(task.Payload(), &data)

		log.Printf("execute send %s \n", data.To)
		return e.Send(data.To, data.Subject, data.Body) //kita bikin jika error akan retry send email
	}
}
