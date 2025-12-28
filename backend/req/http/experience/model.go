package experience

type ExperienceForm struct {
	Company        string `json:"company"`
	Location_field string `json:"location_field"`
	Position       string `json:"position"`
	Date_from      string `json:"date_from"`
	Date_to        string `json:"date_to"`
}
