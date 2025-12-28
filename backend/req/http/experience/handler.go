package experience

import (
	"backend/internal/db"
	"errors"
)

func InserExperienceTable(resumeID int64, form ExperienceForm) (int64, error) {
	query := `
		INSERT INTO experience (resume_id, company, location_field, position, date_from, date_to)
		VALUES (?,?,?,?,?,?)
	`
	res, err := db.DB.Exec(
		query,
		resumeID,
		form.Company,
		form.Location_field,
		form.Position,
		form.Date_from,
		form.Date_to,
	)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (e *ExperienceForm) Validate() error {
	if e.Company == "" {
		return errors.New("company name is required")
	}

	if e.Location_field == "" {
		return errors.New("location of company is required")
	}

	if e.Position == "" {
		return errors.New("position in company is required")
	}

	return nil
}
