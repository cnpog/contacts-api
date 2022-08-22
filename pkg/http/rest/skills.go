package rest

import (
	"contacts-api/pkg/skills"

	"github.com/labstack/echo/v4"
)

// readSkills returns a HandlerFunc for GET /skills requests
func readSkills(c skills.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(200, c.ReadSkills())
	}
}

// readSkill returns a HandlerFunc for GET /skills/id requests
func readSkill(c skills.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := c.ReadSkill(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(404, err.Error())
		}
		return ctx.JSON(200, res)
	}
}

// createSkill returns a HandlerFunc for POST /skills requests
func createSkill(c skills.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var skill skills.Skill
		if err := ctx.Bind(&skill); err != nil {
			return ctx.JSON(400, err)
		}
		res, err := c.CreateSkill(skill)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(201, res)
	}
}

// updateSkill returns a HandlerFunc for PUT /skills/id requests
func updateSkill(c skills.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var skill skills.Skill
		if err := ctx.Bind(&skill); err != nil {
			return ctx.JSON(400, err)
		}
		res, err := c.UpdateSkill(ctx.Param("id"), skill)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(200, res)
	}
}

// deleteSkill returns a HandlerFunc for DELETE /skills/id requests
func deleteSkill(c skills.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := c.DeleteSkill(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(204, nil)
	}
}
