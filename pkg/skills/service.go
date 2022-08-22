package skills

import (
	"errors"
)

var ErrorSkillNotFound = errors.New("skill not found")
var ErrorIdMismatch = errors.New("id mismatch")
var ErrorInvalidSkill = errors.New("invalid skill")
var ErrorDuplicateskill = errors.New("duplicate skill")

// Skills Service with basic crud operations
type Service interface {
	ReadSkills() []Skill
	ReadSkill(string) (Skill, error)
	CreateSkill(Skill) (Skill, error)
	UpdateSkill(string, Skill) (Skill, error)
	DeleteSkill(string) error
}

// interface to the storage port
type Repository interface {
	GetAllSkills() []Skill
	GetSkillById(string) (Skill, error)
	CreateSkill(Skill) (Skill, error)
	UpdateSkill(int, Skill) error
	DeleteSkill(string) error
}

type service struct {
	r Repository
}

// NewService creates new skills service with given storage port
func NewService(r Repository) Service {
	return &service{r}
}

// CreateSkill checks if skill is valid and creates it
func (s *service) CreateSkill(skill Skill) (Skill, error) {
	if skill.Name == "" ||
		skill.Level == "" {
		return Skill{}, ErrorInvalidSkill
	}
	existingSkills := s.r.GetAllSkills()
	for _, es := range existingSkills {
		if skill.Name == es.Name && skill.Level == es.Level {
			return Skill{}, ErrorDuplicateskill
		}
	}
	c, err := s.r.CreateSkill(skill)
	if err != nil {
		return Skill{}, err
	}
	return c, nil
}

// ReadSkills returns all skills
func (s *service) ReadSkills() []Skill {
	return s.r.GetAllSkills()
}

// ReadSkill returns skill with given id
func (s *service) ReadSkill(id string) (Skill, error) {
	skills := s.r.GetAllSkills()
	for _, c := range skills {
		if c.Id == id {
			return c, nil
		}
	}
	return Skill{}, ErrorSkillNotFound
}

// UpdateSkill validates skill and updates it
func (s *service) UpdateSkill(id string, skill Skill) (Skill, error) {
	if id != skill.Id {
		return Skill{}, ErrorIdMismatch
	}
	if skill.Name == "" ||
		skill.Level == "" {
		return Skill{}, ErrorInvalidSkill
	}
	skills := s.r.GetAllSkills()
	for i, c := range skills {
		if c.Id == id {
			err := s.r.UpdateSkill(i, skill)
			if err != nil {
				return s.r.GetSkillById(id)
			}
			return skill, nil
		}
	}
	return Skill{}, ErrorSkillNotFound
}

// DeleteSkill deletes skill if id was found
func (s *service) DeleteSkill(id string) error {
	skills := s.r.GetAllSkills()
	for _, c := range skills {
		if c.Id == id {
			return s.r.DeleteSkill(id)
		}
	}
	return ErrorSkillNotFound
}
