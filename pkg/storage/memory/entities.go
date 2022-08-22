package memory

// Contact is a storage Entity for a Contact
type Contact struct {
	Id                string
	Firstname         string
	Lastname          string
	Fullname          string
	Address           string
	Email             string
	MobilePhoneNumber string
	Skills            []string
}

// Skill is a storage Entity for a Skill
type Skill struct {
	Id    string
	Name  string
	Level string
}
