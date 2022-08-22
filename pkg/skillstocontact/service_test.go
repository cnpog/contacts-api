package skillstocontact_test

import (
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/skills"
	"contacts-api/pkg/skillstocontact"
	"contacts-api/pkg/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the contacts service. Trying to hit all possible responses
func TestSkillsToContactService(t *testing.T) {
	s := new(memory.Storage)
	contactService := contacts.NewService(s)
	skillService := skills.NewService(s)
	skillsToContactService := skillstocontact.NewService(s)
	contact := contacts.Contact{
		Firstname:         "John",
		Lastname:          "Doe",
		Fullname:          "John Doe",
		Address:           "123 Main St",
		Email:             "john@doe.com",
		MobilePhoneNumber: "1234567890",
	}
	skill := skills.Skill{
		Name:  "Java",
		Level: "Expert",
	}
	c, err := contactService.CreateContact(contact)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(contactService.ReadContacts()), "Expected 1 contact")
	skill, err = skillService.CreateSkill(skill)
	assert.Equal(t, 1, len(skillService.ReadSkills()), "Expected 1 skill")

	cs, err := skillsToContactService.AddSkillsToContact(c.Id, []string{skill.Id})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cs), "Expected 1 skill in contact")
	err = skillsToContactService.DeleteSkillsOfContact(c.Id)
	assert.NoError(t, err)
	cs, err = skillsToContactService.ReadSkillsFromContact(c.Id)
	assert.Equal(t, 0, len(cs), "Expected 0 skills in contact")

}
