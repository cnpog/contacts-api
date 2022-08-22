package contacts_test

import (
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the contacts service. Trying to hit all possible responses
func TestContactService(t *testing.T) {
	s := new(memory.Storage)
	contactService := contacts.NewService(s)

	contact := contacts.Contact{
		Firstname:         "John",
		Lastname:          "Doe",
		Fullname:          "John Doe",
		Address:           "123 Main St",
		Email:             "john@doe.com",
		MobilePhoneNumber: "1234567890",
	}
	c, err := contactService.CreateContact(contact)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(contactService.ReadContacts()), "Expected 1 contact")

	_, err = contactService.CreateContact(contact)
	assert.EqualError(t, err, contacts.ErrorDuplicateEmail.Error())
	contact.Email = "doe@john.com"
	_, err = contactService.CreateContact(contact)
	assert.EqualError(t, err, contacts.ErrorDuplicateMobilePhoneNumber.Error())

	c.Lastname = "Dough"
	u, err := contactService.UpdateContact(c.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, "Dough", u.Lastname, "Lastname should be Dough")

	cr, err := contactService.ReadContact(c.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, cr)

	_, err = contactService.ReadContact("17ab03e4-b001-4027-85b2-4c102202e870")
	assert.EqualError(t, err, contacts.ErrorContactNotFound.Error())

	err = contactService.DeleteContact("17ab03e4-b001-4027-85b2-4c102202e870")
	assert.EqualError(t, err, contacts.ErrorContactNotFound.Error())
	err = contactService.DeleteContact(c.Id)
	assert.NoError(t, err)
}
