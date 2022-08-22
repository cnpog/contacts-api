package skills_test

import (
	"contacts-api/pkg/skills"
	"contacts-api/pkg/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the skills service. Trying to hit all possible responses
func TestSkillService(t *testing.T) {
	s := new(memory.Storage)
	skillService := skills.NewService(s)

	skill := skills.Skill{
		Name:  "Java",
		Level: "Expert",
	}
	sn, err := skillService.CreateSkill(skill)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(skillService.ReadSkills()), "Expected 1 skill")

	_, err = skillService.CreateSkill(skill)
	assert.EqualError(t, err, skills.ErrorDuplicateskill.Error())

	sn.Level = "Beginner"
	u, err := skillService.UpdateSkill(sn.Id, sn)
	assert.NoError(t, err)
	assert.Equal(t, "Beginner", u.Level, "Level should be beginner")

	cr, err := skillService.ReadSkill(sn.Id)
	assert.NoError(t, err)
	assert.Equal(t, sn, cr)

	_, err = skillService.ReadSkill("17ab03e4-b001-4027-85b2-4c102202e870")
	assert.EqualError(t, err, skills.ErrorSkillNotFound.Error())

	err = skillService.DeleteSkill("17ab03e4-b001-4027-85b2-4c102202e870")
	assert.EqualError(t, err, skills.ErrorSkillNotFound.Error())
	err = skillService.DeleteSkill(sn.Id)
	assert.NoError(t, err)
}
