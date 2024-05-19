package services

import (
	"github.com/fentezi/api-rate/config"
	"github.com/fentezi/api-rate/internal/repositories"
	"github.com/fentezi/api-rate/internal/services/emails"
	"github.com/fentezi/api-rate/internal/services/rate"
)

type Rate interface {
	GetRate() (float64, error)
}

type Emails interface {
	CreateEmail(email string) (int64, error)
	Mailing(rate float64) error
}

type Service struct {
	Rate
	Emails
}

func NewService(db repositories.Repository, cfg *config.Config) *Service {
	return &Service{
		Rate:   rate.NewRateService(),
		Emails: emails.NewEmailsService(db.Emails, cfg),
	}
}
