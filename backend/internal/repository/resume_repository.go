package repository

import (
	"backend/internal/gemini"
	"context"
	"database/sql"
	"time"
)

type CVProcedureSpecification struct {
	db *sql.DB
}

func NewWeldingProcedureSpecification(db *sql.DB) *CVProcedureSpecification {
	return &CVProcedureSpecification{db: db}
}

func (r *CVProcedureSpecification) Save(ctx context.Context, res *gemini.ResultHeader) error {
	query := `
			INSERT INTO headers (
			first_name,
			last_name,
			title
		) VALUES (?,?,?)
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query,
		res.FirstName,
		res.LastName,
		res.Title,
	)
	return err
}
