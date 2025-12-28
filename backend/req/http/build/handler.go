package build

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/req/http/contacts"
	"backend/req/http/experience"
	"backend/req/http/header"
	"backend/req/http/language"
	"backend/req/http/links"
	"backend/req/http/projects"
	"backend/req/http/projecttechnologies"
	"backend/req/http/responsibilities"
	"backend/req/http/technologies"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if rec := recover(); rec != nil {
			log.Println("panic:", rec)
			http.Error(w, "internal server error", 500)
		}
	}()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var form BuildForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println("DECODE ERROR:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := form.Language.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if form.Language != nil {
	resumeID, err := language.InsertResumesTable(*form.Language)
	if err != nil {
		log.Println(err)
		http.Error(w, "db error", 500)
		return
	}
	// }

	if err := form.Experience.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	experienceID, err := experience.InserExperienceTable(resumeID, *form.Experience)
	if err != nil {
		log.Println(err)
		http.Error(w, "db error", 500)
		return
	}

	if err := form.Projects.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if form.Projects != nil {
	projectID, err := projects.InsertProjectsTable(experienceID, *form.Projects)
	if err != nil {
		log.Println(err)
		http.Error(w, "db error", 500)
		return
	}
	// }

	if err := form.Header.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if form.Header != nil {
		if err := header.InsertHeadersTable(resumeID, *form.Header); err != nil {
			log.Println(err)
			http.Error(w, "db error", 500)
			return
		}
	}

	if err := form.Contacts.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if form.Contacts != nil {
		if err := contacts.InsertContactTable(resumeID, *form.Contacts); err != nil {
			log.Println("db error:", err)
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
	}
	// avoiding data repetition
	// exists, err := contacts.ContactsExists(form.Contacts.Phone, form.Contacts.Email, form.Contacts.Address)
	// if err != nil {
	// 	http.Error(w, "db error", 500)
	// 	return
	// }
	// if exists {
	// 	http.Error(w, "resume with this phone or email already exists", 400)
	// 	return
	// }

	if err := form.Responsibilities.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if form.Responsibilities != nil {
		if err := responsibilities.InsertResponsibilitiesTable(projectID, *form.Responsibilities); err != nil {
			log.Println("db error:", err)
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
	}

	if err := form.Technologies.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if form.Technologies != nil {
	technologyID, err := technologies.InsertTechnologiesTable(*form.Technologies)
	if err != nil {
		log.Println("db error:", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	// }

	if form.Projecttechnologies != nil {
		if err := projecttechnologies.InsertProjectTechnologiesTable(projectID, technologyID, *form.Projecttechnologies); err != nil {
			log.Println("db error:", err)
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
	}

	if err := form.Links.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if form.Links != nil {
		if err := links.InsertLinksTable(projectID, *form.Links); err != nil {
			log.Println("db error:", err)
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]any{
		"ok": true,
	})
}
