package build

import (
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

type BuildForm struct {
	Header              *header.HeaderForm                           `json:"header"`
	Contacts            *contacts.ContactsForm                       `json:"contacts"`
	Experience          *experience.ExperienceForm                   `json:"experience"`
	Language            *language.LanguageForm                       `json:"language"`
	Links               *links.LinksForm                             `json:"links"`
	Projects            *projects.ProjectsForm                       `json:"projects"`
	Responsibilities    *responsibilities.ResponsibilitiesForm       `json:"responsibilities"`
	Technologies        *technologies.TechnologiesForm               `json:"technologies"`
	Projecttechnologies *projecttechnologies.ProjectTechnologiesForm `json:"projecttechnologies"`
}
