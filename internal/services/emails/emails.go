package emails

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/fentezi/api-rate/config"
	"github.com/fentezi/api-rate/internal/repositories"
	"gopkg.in/gomail.v2"
)

type EmailsService struct {
	db  repositories.Emails
	cfg *config.Config
}

func NewEmailsService(db repositories.Emails, cfg *config.Config) *EmailsService {
	return &EmailsService{
		db:  db,
		cfg: cfg,
	}
}

func (e *EmailsService) CreateEmail(email string) (int64, error) {
	err := e.db.GetEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			id, err := e.db.CreateEmail(email)
			if err != nil {
				return 0.0, err
			}
			return id, nil
		}
		return 0.0, err
	}

	return 0.0, fmt.Errorf("email already exists")
}

func (e *EmailsService) Mailing(rate float64) error {
	emails, err := e.db.GetAllEmails()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	user := e.cfg.EmailUser
	pass := e.cfg.EmailPass

	auth := gomail.NewDialer("smtp.gmail.com", 465, user, pass)
	auth.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	for _, email := range emails {
		wg.Add(1)
		go sendEmail(auth, user, email.Email, []byte(fmt.Sprintf("Rate: %f", rate)), &wg)
	}

	wg.Wait()

	return nil
}

func sendEmail(auth *gomail.Dialer, from string, to string, msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Rate")
	m.SetBody("text/plain", string(msg))

	if err := auth.DialAndSend(m); err != nil {
		log.Println(err)
	}

	log.Println("Email sent")
}
