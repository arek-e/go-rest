package main

import (
	"github.com/arek-e/lanexpense/domain"
	"github.com/arek-e/lanexpense/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&domain.Account{})
}
