package memory

// UserExists checks if a user exists
func (s *Storage) UserExists(email string) bool {
	for _, user := range s.contacts {
		if user.Email == email {
			return true
		}
	}
	return false
}

// GetUserIdByMail returns the user id of a user by email
func (s *Storage) GetUserIdByMail(email string) string {
	for _, user := range s.contacts {
		if user.Email == email {
			return user.Id
		}
	}
	return ""
}
