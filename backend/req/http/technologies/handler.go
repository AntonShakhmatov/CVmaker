package technologies

import (
	"backend/internal/db"
	"errors"
)

func InsertTechnologiesTable(experienceID int64, form TechnologiesForm) (int64, error) {
	query := `
		INSERT INTO technologies (experience_id, technologies_name)
		VALUES (?,?)
	`
	res, err := db.DB.Exec(
		query,
		experienceID,
		form.Name,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (t *TechnologiesForm) Validate() error {
	if (t.Name) == "" {
		return errors.New("describe please your tech-stack")
	}

	return nil
}
