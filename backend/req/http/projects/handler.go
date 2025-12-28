package projects

import (
	"backend/internal/db"
	"errors"
)

func InsertProjectsTable(experienceID int64, form ProjectsForm) (int64, error) {
	query := `
		INSERT INTO projects (experience_id, name, description)
		VALUES (?,?,?)
	`
	res, err := db.DB.Exec(
		query,
		experienceID,
		form.Name,
		form.Description,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (p *ProjectsForm) Validate() error {
	if (p.Name) == "" {
		return errors.New("project name is required")
	}

	if (p.Description) == "" {
		return errors.New("describe your work on project please")
	}

	// if len(p.Description) < 45 {
	// 	return errors.New("description should have at least 45 characters")
	// }

	return nil
}
