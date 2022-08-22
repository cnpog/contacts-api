package rest

import (
	"contacts-api/pkg/contacts"

	"github.com/labstack/echo/v4"
)

func Handler(c contacts.Service) *echo.Echo {
	router := echo.New()
	router.GET("/contacts", readContacts(c))
	router.GET("/contacts/:id", readContact(c))
	router.POST("/contacts", createContact(c))
	router.PUT("/contacts/:id", updateContact(c))
	router.DELETE("/contacts/:id", deleteContact(c))

	return router
}

// readContacts returns a HandlerFunc for GET /contacts requests
func readContacts(c contacts.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(200, c.ReadContacts())
	}
}

// readContact returns a HandlerFunc for GET /contacts/id requests
func readContact(c contacts.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := c.ReadContact(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(404, err.Error())
		}
		return ctx.JSON(200, res)
	}
}

// createContact returns a HandlerFunc for POST /contacts requests
func createContact(c contacts.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var contact contacts.Contact
		if err := ctx.Bind(&contact); err != nil {
			return ctx.JSON(400, err)
		}
		res, err := c.CreateContact(contact)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(201, res)
	}
}

// updateContact returns a HandlerFunc for PUT /contacts/id requests
func updateContact(c contacts.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var contact contacts.Contact
		if err := ctx.Bind(&contact); err != nil {
			return ctx.JSON(400, err)
		}
		res, err := c.UpdateContact(ctx.Param("id"), contact)
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(200, res)
	}
}

// deleteContact returns a HandlerFunc for DELETE /contacts/id requests
func deleteContact(c contacts.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := c.DeleteContact(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(400, err.Error())
		}
		return ctx.JSON(204, nil)
	}
}
