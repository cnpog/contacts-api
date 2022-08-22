package skillstocontact

import (
	"contacts-api/pkg/skills"
	"errors"
)

var ErrorInvalidListOfSkills = errors.New("Invalid list of skills")
var ErrorContactNotFound = errors.New("Contact not found")
var ErrorDuplicateSkill = errors.New("Duplicate skill")

// Skills To Contacts Service with basic crud operations
type Service interface {
	ReadSkillsFromContact(string) ([]skills.Skill, error)
	AddSkillsToContact(string, []string) ([]skills.Skill, error)
	UpdateSkillsOfContact(string, []string) ([]skills.Skill, error)
	DeleteSkillsOfContact(string) error
}

// interface to the storage port
type Repository interface {
	GetSkillsByContactId(string) ([]skills.Skill, error)
	AddSkillToContact(string, string) error
	DeleteSkillsOfContact(string) error
}

type service struct {
	r Repository
}

// NewService creates new skills to contacts service with given storage port
func NewService(r Repository) Service {
	return &service{r}
}

// AddSkillsToContact checks if skills are valid and adds them to the contact
func (s *service) AddSkillsToContact(contactid string, skills []string) ([]skills.Skill, error) {
	if len(skills) == 0 {
		return nil, ErrorInvalidListOfSkills
	}
	existingSkills, err := s.r.GetSkillsByContactId(contactid)
	if err != nil {
		return nil, ErrorContactNotFound
	}
	for _, es := range existingSkills {
		for _, sk := range skills {
			if es.Id == sk {
				return nil, ErrorDuplicateSkill
			}
		}
	}
	for _, sk := range skills {
		err := s.r.AddSkillToContact(contactid, sk)
		if err != nil {
			return nil, err
		}
	}
	return s.r.GetSkillsByContactId(contactid)
}

// ReadSkillsFromContact returns all skills of the contact
func (s *service) ReadSkillsFromContact(contactid string) ([]skills.Skill, error) {
	return s.r.GetSkillsByContactId(contactid)
}

// DeleteSkillsOfContact deletes all skills of the contact
func (s *service) DeleteSkillsOfContact(contactid string) error {
	return s.r.DeleteSkillsOfContact(contactid)
}

// UpdateSkillsOfContact just deletes all skills and adds a list of new ones for simplicity
func (s *service) UpdateSkillsOfContact(contactid string, skills []string) ([]skills.Skill, error) {
	if len(skills) == 0 {
		return nil, ErrorInvalidListOfSkills
	}
	err := s.r.DeleteSkillsOfContact(contactid)
	if err != nil {
		return nil, err
	}
	for _, sk := range skills {
		err := s.r.AddSkillToContact(contactid, sk)
		if err != nil {
			return nil, err
		}
	}
	return s.r.GetSkillsByContactId(contactid)
}
