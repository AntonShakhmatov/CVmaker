package language

import (
	"backend/internal/db"
	"errors"
)

func InsertResumesTable(form LanguageForm) (int64, error) {
	query := `
		INSERT INTO resumes(language)
		VALUES(?)
	`

	res, err := db.DB.Exec(query, form.Language)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (l *LanguageForm) Validate() error {
	if (l.Language) == "" {
		return errors.New("choose the language of your CV, please")
	}

	return nil
}
