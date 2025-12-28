package projecttechnologies

import (
	"backend/internal/db"
)

func InsertProjectTechnologiesTable(projectID int64, technologyID int64, form ProjectTechnologiesForm) error {
	query := `
		INSERT INTO technologies (project_id, technology_id)
		VALUES (?,?)
	`
	_, err := db.DB.Exec(
		query,
		projectID,
		technologyID,
	)
	return err
}
