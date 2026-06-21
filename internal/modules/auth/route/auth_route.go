package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/handler"
)

func AuthRoute(rg *gin.RouterGroup, authHandler *handler.AuthHandler) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

}
