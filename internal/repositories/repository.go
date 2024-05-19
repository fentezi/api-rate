package repositories

import (
	"database/sql"

	"github.com/fentezi/api-rate/internal/models"
	"github.com/fentezi/api-rate/internal/repositories/emails"
)

type Repository struct {
	Emails
}

type Emails interface {
	CreateEmail(email string) (int64, error)
	GetEmail(email string) error
	GetAllEmails() ([]models.Email, error)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Emails: emails.NewEmailsRepository(db),
	}
}
