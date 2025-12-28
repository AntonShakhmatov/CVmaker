package contacts

type ContactsForm struct {
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address_field"`
}
