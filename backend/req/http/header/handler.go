package header

import (
	"backend/internal/db"
	"errors"
)

func InsertHeadersTable(resumeID int64, form HeaderForm) error {
	query := `
		INSERT INTO headers
		(resume_id, first_name, last_name, competence, title)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.DB.Exec(
		query,
		resumeID,
		form.FirstName,
		form.LastName,
		form.Competence,
		form.Title,
	)

	return err
}

func (h *HeaderForm) Validate() error {
	if (h.FirstName) == "" {
		return errors.New("your name is required")
	}

	if (h.LastName) == "" {
		return errors.New("your lastname is required")
	}

	if (h.Competence) == "" {
		return errors.New("your competence is required")
	}

	if (h.Title) == "" {
		return errors.New("your title is required")
	}

	return nil
}
