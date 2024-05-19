package emails

import (
	"database/sql"

	"github.com/fentezi/api-rate/internal/models"
)

type EmailsRepository struct {
	db *sql.DB
}

func NewEmailsRepository(db *sql.DB) *EmailsRepository {
	return &EmailsRepository{
		db: db,
	}
}

func (e *EmailsRepository) CreateEmail(email string) (int64, error) {
	var id int64
	err := e.db.QueryRow("INSERT INTO emails (email) VALUES ($1) RETURNING id", email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (e *EmailsRepository) GetEmail(email string) error {
	var result models.Email
	err := e.db.QueryRow("SELECT email FROM emails WHERE email = $1", email).Scan(&result.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}

func (e *EmailsRepository) GetAllEmails() ([]models.Email, error) {
	rows, err := e.db.Query("SELECT email FROM emails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []models.Email

	for rows.Next() {
		var email models.Email
		if err := rows.Scan(&email.Email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}
