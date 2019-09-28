package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josephmathew1996/trendz-holidays/controllers"
)

// Creates a gin router with default middleware:
func SetUpRouter(router *gin.Engine) {
	router.POST("/login", controllers.SignUpPost)
}
