package contacts

import (
	"errors"
	"strings"
)

var ErrorDuplicateEmail = errors.New("email already in use")
var ErrorDuplicateMobilePhoneNumber = errors.New("mobile number already in use")
var ErrorContactNotFound = errors.New("contact not found")
var ErrorIdMismatch = errors.New("id mismatch")
var ErrorInvalidContact = errors.New("invalid contact")

// Contacts Service with basic crud operations
type Service interface {
	ReadContacts() []Contact
	ReadContact(string) (Contact, error)
	CreateContact(Contact) (Contact, error)
	UpdateContact(string, Contact) (Contact, error)
	DeleteContact(string) error
}

// interface to the storage port
type Repository interface {
	GetAllContacts() []Contact
	GetContactById(string) (Contact, error)
	CreateContact(Contact) (Contact, error)
	UpdateContact(int, Contact) error
	DeleteContact(string) error
}

type service struct {
	r Repository
}

// NewService creates new contacts service with given storage port
func NewService(r Repository) Service {
	return &service{r}
}

// CreateContact checks if contact is valid and creates it
func (s *service) CreateContact(contact Contact) (Contact, error) {
	if contact.Firstname == "" ||
		contact.Lastname == "" ||
		contact.Email == "" || !strings.Contains(contact.Email, "@") ||
		contact.Address == "" ||
		contact.Fullname == "" ||
		contact.MobilePhoneNumber == "" {
		return Contact{}, ErrorInvalidContact
	}
	existingContacts := s.r.GetAllContacts()
	for _, ec := range existingContacts {
		if contact.Email == ec.Email {
			return Contact{}, ErrorDuplicateEmail
		}
		if contact.MobilePhoneNumber == ec.MobilePhoneNumber {
			return Contact{}, ErrorDuplicateMobilePhoneNumber
		}
	}
	c, err := s.r.CreateContact(contact)
	if err != nil {
		return Contact{}, err
	}
	return c, nil
}

// ReadContacts returns all contacts
func (s *service) ReadContacts() []Contact {
	return s.r.GetAllContacts()
}

// ReadContact returns contact with given id
func (s *service) ReadContact(id string) (Contact, error) {
	contacts := s.r.GetAllContacts()
	for _, c := range contacts {
		if c.Id == id {
			return c, nil
		}
	}
	return Contact{}, ErrorContactNotFound
}

// UpdateContact validates contact and updates it
func (s *service) UpdateContact(id string, contact Contact) (Contact, error) {
	if id != contact.Id {
		return Contact{}, ErrorIdMismatch
	}
	if contact.Firstname == "" ||
		contact.Lastname == "" ||
		contact.Email == "" || !strings.Contains(contact.Email, "@") ||
		contact.Address == "" ||
		contact.Fullname == "" ||
		contact.MobilePhoneNumber == "" {
		return Contact{}, ErrorInvalidContact
	}
	contacts := s.r.GetAllContacts()
	for i, c := range contacts {
		if c.Id == id {
			err := s.r.UpdateContact(i, contact)
			if err != nil {
				return s.r.GetContactById(id)
			}
			return contact, nil
		}
	}
	return Contact{}, ErrorContactNotFound
}

// DeleteContact deletes contact if id was found
func (s *service) DeleteContact(id string) error {
	contacts := s.r.GetAllContacts()
	for _, c := range contacts {
		if c.Id == id {
			return s.r.DeleteContact(id)
		}
	}
	return ErrorContactNotFound
}
