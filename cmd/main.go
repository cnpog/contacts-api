package main

import (
	"contacts-api/pkg/auth"
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/http/rest"
	"contacts-api/pkg/skills"
	"contacts-api/pkg/skillstocontact"
	"contacts-api/pkg/storage/memory"
	"log"
)

func main() {
	var contactService contacts.Service
	var skillService skills.Service
	var skillsToContactService skillstocontact.Service
	var authService auth.Service
	// init storage and services
	s := new(memory.Storage)
	contactService = contacts.NewService(s)
	skillService = skills.NewService(s)
	skillsToContactService = skillstocontact.NewService(s)
	authService = auth.NewService(s)

	// setup HTTP server
	router := rest.Handler(contactService, skillService, skillsToContactService, authService)
	log.Fatal(router.Start(":8080"))
}
