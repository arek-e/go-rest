package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/arek-e/lanexpense/services"
)

// NewRouter
func NewRouter(platform *services.AccountService) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)
	// router := gin.Default()

	h := AccountHandlers{Platform: platform}

	account := router.Group("/account")
	account.GET("/", h.GetAllAccounts)
	account.GET("/:id", h.GetAccount)
	account.POST("/", h.CreateAccount)
	account.PUT("/:id", h.UpdateAccount)
	account.DELETE("/:id", h.DeleteAccount)

	return router
}
