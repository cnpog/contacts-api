package main

import (
	"contacts-api/pkg/contacts"
	"contacts-api/pkg/http/rest"
	"contacts-api/pkg/storage/memory"
	"log"
)

func main() {
	var contactService contacts.Service
	// init storage and services
	s := new(memory.Storage)
	contactService = contacts.NewService(s)

	// setup HTTP server
	router := rest.Handler(contactService)
	log.Fatal(router.Start(":8080"))
}
