package main

import (
	"go-code-generation/models"
	"go-code-generation/services"
)
func main() {
	// init config

	// start server
	services.Read(models.Postgres)
}
