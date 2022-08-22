package contacts

// Contact defines the properties of a contact
type Contact struct {
	Id                string `json:"id"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Fullname          string `json:"fullname"`
	Address           string `json:"address"`
	Email             string `json:"email"`
	MobilePhoneNumber string `json:"mobile_phone_number"`
}
