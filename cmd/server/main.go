package main

import (
	"Colombo-Romina/cmd/server/routes"
)

// @title Certified Tech Developer
// @version 1.0
// @description API Clinica Odontologica

// @contact.name Romina Colombo
// @contact.email colombo.romina@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	router := routes.Routes()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
