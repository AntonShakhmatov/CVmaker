package contacts

import (
	"backend/internal/db"
	"errors"
	"strings"
)

func InsertContactTable(resumeID int64, form ContactsForm) error {
	query := `
			INSERT INTO contacts (resume_id, phone, email, address_field)
			VALUES (?,?,?,?)
		`
	_, err := db.DB.Exec(
		query,
		resumeID,
		form.Phone,
		form.Email,
		form.Address,
	)
	return err
}

func (c *ContactsForm) Validate() error {
	// Phone
	if c.Phone == "" {
		return errors.New("phone is required")
	}

	if len(c.Phone) < 9 || len(c.Phone) > 13 {
		return errors.New("phone number must be between 9 and 13 characters")
	}

	// Email
	if c.Email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(c.Email, "@") {
		return errors.New("invalid email format")
	}

	// Address
	if c.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

// avoiding data repetition
func ContactsExists(phone string, email string, address string) (bool, error) {
	var count int

	err := db.DB.QueryRow(`
		SELECT COUNT(*) FROM contacts
		WHERE phone = ? OR email = ? OR address = ?`,
		phone, email, address).Scan(&count)

	return count > 0, err
}
