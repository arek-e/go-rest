// The starting point of the backend
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/arek-e/lanexpense/handlers"
	"github.com/arek-e/lanexpense/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/accounts", handlers.AccountsIndex)
	r.GET("/accounts/:id", handlers.AccountsShow)
	r.PUT("/accounts/:id", handlers.AccountsUpdate)
	r.DELETE("/accounts/:id", handlers.AccountsDelete)
	r.POST("/accounts", handlers.AccountsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
