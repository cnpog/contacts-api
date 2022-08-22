package rest

import (
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/skills"
	"contacts-api/pkg/skillstocontact"

	"github.com/labstack/echo/v4"
)

func Handler(c contacts.Service, s skills.Service, sc skillstocontact.Service) *echo.Echo {
	router := echo.New()
	router.GET("/contacts", readContacts(c))
	router.GET("/contacts/:id", readContact(c))
	router.POST("/contacts", createContact(c))
	router.PUT("/contacts/:id", updateContact(c))
	router.DELETE("/contacts/:id", deleteContact(c))

	router.GET("/skills", readSkills(s))
	router.GET("/skills/:id", readSkill(s))
	router.POST("/skills", createSkill(s))
	router.PUT("/skills/:id", updateSkill(s))
	router.DELETE("/skills/:id", deleteSkill(s))

	router.GET("/contacts/:id/skills", readContactSkills(sc))
	router.POST("/contacts/:id/skills", addContactSkills(sc))
	router.PUT("/contacts/:id/skills", updateContactSkills(sc))
	router.DELETE("/contacts/:id/skills", deleteContactSkills(sc))

	return router
}
