package Router

import (
	"TheLabSystem/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	g.POST("/auth/login", Controller.AuthController{}.Login)
}
