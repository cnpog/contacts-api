package memory

import (
	"contacts-api/pkg/contacts"

	"github.com/google/uuid"
)

// GetContacts returns all contacts from storage
func (s *Storage) GetAllContacts() []contacts.Contact {
	var allContacts []contacts.Contact
	for i := range s.contacts {
		contact := contacts.Contact{
			Id:                s.contacts[i].Id,
			Firstname:         s.contacts[i].Firstname,
			Lastname:          s.contacts[i].Lastname,
			Fullname:          s.contacts[i].Fullname,
			Address:           s.contacts[i].Address,
			Email:             s.contacts[i].Email,
			MobilePhoneNumber: s.contacts[i].MobilePhoneNumber,
		}
		allContacts = append(allContacts, contact)
	}
	return allContacts
}

// GetContactById returns contact with given id
func (s *Storage) GetContactById(id string) (contacts.Contact, error) {
	var contact contacts.Contact
	for i := range s.contacts {
		if s.contacts[i].Id == id {
			contact = contacts.Contact{
				Id:                s.contacts[i].Id,
				Firstname:         s.contacts[i].Firstname,
				Lastname:          s.contacts[i].Lastname,
				Fullname:          s.contacts[i].Fullname,
				Address:           s.contacts[i].Address,
				Email:             s.contacts[i].Email,
				MobilePhoneNumber: s.contacts[i].MobilePhoneNumber,
			}
			return contact, nil
		}
	}
	return contact, contacts.ErrorContactNotFound
}

// CreateContact saves contact to storage
func (s *Storage) CreateContact(contact contacts.Contact) (contacts.Contact, error) {
	contact.Id = uuid.New().String()
	s.contacts = append(s.contacts, Contact{
		Id:                contact.Id,
		Firstname:         contact.Firstname,
		Lastname:          contact.Lastname,
		Fullname:          contact.Fullname,
		Address:           contact.Address,
		Email:             contact.Email,
		MobilePhoneNumber: contact.MobilePhoneNumber,
	})
	return contact, nil
}

// UpdateContact saves updated contact to storage
func (s *Storage) UpdateContact(index int, contact contacts.Contact) error {
	if index < 0 || index > len(s.contacts) {
		return contacts.ErrorContactNotFound
	}
	s.contacts[index] = Contact{
		Id:                contact.Id,
		Firstname:         contact.Firstname,
		Lastname:          contact.Lastname,
		Fullname:          contact.Fullname,
		Address:           contact.Address,
		Email:             contact.Email,
		MobilePhoneNumber: contact.MobilePhoneNumber,
	}
	return nil
}

// DeleteContact deletes contact from storage
func (s *Storage) DeleteContact(id string) error {
	for i := range s.contacts {
		if s.contacts[i].Id == id {
			s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
			return nil
		}
	}
	return contacts.ErrorContactNotFound
}
