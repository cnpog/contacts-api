package memory

import (
	"contacts-api/pkg/skills"

	"github.com/google/uuid"
)

// GetSkills returns all skills from storage
func (s *Storage) GetAllSkills() []skills.Skill {
	var allSkills []skills.Skill
	for i := range s.skills {
		skill := skills.Skill{
			Id:    s.skills[i].Id,
			Name:  s.skills[i].Name,
			Level: s.skills[i].Level,
		}
		allSkills = append(allSkills, skill)
	}
	return allSkills
}

// GetSkillById returns skill with given id
func (s *Storage) GetSkillById(id string) (skills.Skill, error) {
	var skill skills.Skill
	for i := range s.skills {
		if s.skills[i].Id == id {
			skill = skills.Skill{
				Id:    s.skills[i].Id,
				Name:  s.skills[i].Name,
				Level: s.skills[i].Level,
			}
			return skill, nil
		}
	}
	return skill, skills.ErrorSkillNotFound
}

// CreateSkill saves skill to storage
func (s *Storage) CreateSkill(skill skills.Skill) (skills.Skill, error) {
	skill.Id = uuid.New().String()
	s.skills = append(s.skills, Skill{
		Id:    skill.Id,
		Name:  skill.Name,
		Level: skill.Level,
	})
	return skill, nil
}

// UpdateSkill saves updated skill to storage
func (s *Storage) UpdateSkill(index int, skill skills.Skill) error {
	if index < 0 || index > len(s.skills) {
		return skills.ErrorSkillNotFound
	}
	s.skills[index] = Skill{
		Id:    skill.Id,
		Name:  skill.Name,
		Level: skill.Level,
	}
	return nil
}

// DeleteSkill deletes skill from storage
func (s *Storage) DeleteSkill(id string) error {
	for i := range s.skills {
		if s.skills[i].Id == id {
			s.skills = append(s.skills[:i], s.skills[i+1:]...)
			return nil
		}
	}
	return skills.ErrorSkillNotFound
}
