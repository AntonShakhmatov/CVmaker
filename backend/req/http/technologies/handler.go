package technologies

import (
	"backend/internal/db"
	"errors"
)

func InsertTechnologiesTable(form TechnologiesForm) (int64, error) {
	query := `
		INSERT INTO technologies (technologies_name)
		VALUES (?)
	`
	res, err := db.DB.Exec(
		query,
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
