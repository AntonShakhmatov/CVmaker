package success

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"backend/internal/db"
	"backend/req/http/contacts"
	"backend/req/http/experience"
	"backend/req/http/header"
	"backend/req/http/links"
	"backend/req/http/projects"
	"backend/req/http/responsibilities"
	"backend/req/http/technologies"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "missing email", http.StatusBadRequest)
		return
	}

	cv, err := SelectBuildedCv(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "db error", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
		"cv": cv,
	})
}

func SelectBuildedCv(email string) (*BuildedCV, error) {
	query := `
		SELECT
			h.first_name,
			h.last_name,
			h.competence,
			h.title,
			c.phone,
			c.email,
			c.address_field,
			e.company,
			e.location_field,
			e.position,
			e.date_from,
			e.date_to,
			p.name,
			p.description,
			r.responsibilities_description,
			t.technologies_name,
			l.label
		FROM resumes res
		LEFT JOIN headers h ON h.resume_id = res.id
		LEFT JOIN contacts c ON c.resume_id = res.id
		LEFT JOIN experience e ON e.resume_id = res.id
		LEFT JOIN projects p ON p.experience_id = e.id
		LEFT JOIN responsibilities r ON r.project_id = p.id
		LEFT JOIN project_technologies pt ON pt.project_id = p.id
		LEFT JOIN technologies t ON t.id = pt.technology_id
		LEFT JOIN links l ON l.project_id = p.id
		WHERE c.email = ?
	`

	rows, err := db.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cv := BuildedCV{
		Header:           &header.HeaderForm{},
		Contacts:         &contacts.ContactsForm{},
		Experience:       &experience.ExperienceForm{},
		Projects:         &projects.ProjectsForm{},
		Responsibilities: &responsibilities.ResponsibilitiesForm{},
		Technologies:     &technologies.TechnologiesForm{},
		Links:            &links.LinksForm{},
	}

	var resp, tech, label sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&cv.Header.FirstName,
			&cv.Header.LastName,
			&cv.Header.Competence,
			&cv.Header.Title,

			&cv.Contacts.Phone,
			&cv.Contacts.Email,
			&cv.Contacts.Address,

			&cv.Experience.Company,
			&cv.Experience.Location_field,
			&cv.Experience.Position,
			&cv.Experience.Date_from,
			&cv.Experience.Date_to,

			&cv.Projects.Name,
			&cv.Projects.Description,

			&resp,
			&tech,
			&label,
		)
		if err != nil {
			return nil, err
		}
	}

	if resp.Valid {
		cv.Responsibilities.Description = resp.String
	}
	if tech.Valid {
		cv.Technologies.Name = tech.String
	}
	if label.Valid {
		cv.Links.Label = label.String
	}

	return &cv, rows.Err()
}
