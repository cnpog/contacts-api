package memory

import (
	"contacts-api/pkg/skills"
	"contacts-api/pkg/skillstocontact"
)

func (s *Storage) GetSkillsByContactId(contactid string) ([]skills.Skill, error) {
	var skills []skills.Skill
	for _, contact := range s.contacts {
		if contact.Id == contactid {
			for _, cs := range contact.Skills {
				skill, err := s.GetSkillById(cs)
				if err != nil {
					return nil, err
				}
				skills = append(skills, skill)
			}
			return skills, nil
		}
	}
	return skills, skillstocontact.ErrorContactNotFound
}
func (s *Storage) AddSkillToContact(contactid, skillid string) error {
	for i, contact := range s.contacts {
		if contact.Id == contactid {
			s.contacts[i].Skills = append(contact.Skills, skillid)
			return nil
		}
	}
	return skillstocontact.ErrorContactNotFound
}

func (s *Storage) DeleteSkillsOfContact(contactid string) error {
	for i, contact := range s.contacts {
		if contact.Id == contactid {
			s.contacts[i].Skills = []string{}
			return nil
		}
	}
	return skillstocontact.ErrorContactNotFound
}
