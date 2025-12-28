package links

import (
	"backend/internal/db"
	"errors"
)

func InsertLinksTable(projectID int64, form LinksForm) error {
	query := `
		INSERT INTO links (project_id,label)
		VALUE (?,?)
	`
	_, err := db.DB.Exec(
		query,
		projectID,
		form.Label,
	)

	return err
}

func (l *LinksForm) Validate() error {
	if (l.Label) == "" {
		return errors.New("link is required")
	}

	return nil
}
