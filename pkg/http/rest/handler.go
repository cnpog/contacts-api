package rest

import (
	"contacts-api/pkg/auth"
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/skills"
	"contacts-api/pkg/skillstocontact"

	"github.com/labstack/echo/v4"
)

func Handler(c contacts.Service, s skills.Service, sc skillstocontact.Service, a auth.Service) *echo.Echo {
	router := echo.New()
	openRouter := router.Group("")
	openRouter.GET("/skills", readSkills(s))
	openRouter.GET("/skills/:id", readSkill(s))
	openRouter.POST("/skills", createSkill(s))
	openRouter.PUT("/skills/:id", updateSkill(s))
	openRouter.DELETE("/skills/:id", deleteSkill(s))
	openRouter.GET("/contacts", readContacts(c))
	openRouter.GET("/contacts/:id", readContact(c))
	openRouter.GET("/contacts/:id/skills", readContactSkills(sc))
	openRouter.POST("/contacts", createContact(c))

	restrictedRouter := router.Group("/private/contacts/:id")

	restrictedRouter.Use(authMiddleware(a))
	restrictedRouter.PUT("", updateContact(c))
	restrictedRouter.DELETE("", deleteContact(c))
	restrictedRouter.POST("/skills", addContactSkills(sc))
	restrictedRouter.PUT("/skills", updateContactSkills(sc))
	restrictedRouter.DELETE("/skills", deleteContactSkills(sc))

	return router
}
