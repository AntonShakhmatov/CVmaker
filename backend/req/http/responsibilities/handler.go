package responsibilities

import (
	"backend/internal/db"
	"errors"
)

func InsertResponsibilitiesTable(projectID int64, form ResponsibilitiesForm) error {
	query := `
		INSERT INTO responsibilities (project_id, responsibilities_description)
		VALUES (?,?)
	`
	_, err := db.DB.Exec(
		query,
		projectID,
		form.Description,
	)
	return err
}

func (r *ResponsibilitiesForm) Validate() error {
	if (r.Description) == "" {
		return errors.New("responsibilities description is required")
	}

	return nil
}
