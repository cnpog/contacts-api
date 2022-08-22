package rest

import (
	"contacts-api/pkg/skillstocontact"

	"github.com/labstack/echo/v4"
)

// readContactSkills returns a HandlerFunc for GET /contacts/id/skills requests
func readContactSkills(sc skillstocontact.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		skills, err := sc.ReadSkillsFromContact(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(404, err.Error())
		}
		return ctx.JSON(200, skills)
	}
}

// addContactSkill returns a HandlerFunc for POST /contacts/id/skills requests
func addContactSkills(sc skillstocontact.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var skillIds []string
		if err := ctx.Bind(&skillIds); err != nil {
			return ctx.JSON(400, err)
		}
		res, err := sc.AddSkillsToContact(ctx.Param("id"), skillIds)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(201, res)
	}
}

// updateContactSkill returns a HandlerFunc for PUT /contacts/id/skills requests
func updateContactSkills(sc skillstocontact.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var skillIds []string
		if err := ctx.Bind(&skillIds); err != nil {
			return ctx.JSON(400, err.Error())
		}
		res, err := sc.UpdateSkillsOfContact(ctx.Param("id"), skillIds)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(200, res)
	}
}

// deleteContactSkill returns a HandlerFunc for DELETE /contacts/id/skills requests
func deleteContactSkills(sc skillstocontact.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := sc.DeleteSkillsOfContact(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(204, nil)
	}
}
