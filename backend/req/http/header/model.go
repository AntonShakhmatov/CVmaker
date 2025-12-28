package header

type HeaderForm struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Competence string `json:"competence"`
	Title      string `json:"title"`
}
