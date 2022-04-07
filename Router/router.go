package Router

import (
	"TheLabSystem/Controller/AuthController"
	"TheLabSystem/Controller/UserServiceController"
	"TheLabSystem/Controller/VerifyCodeController"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	// auth
	g.POST("/auth/login", AuthController.AuthController{}.Login)
	g.POST("/auth/logout", AuthController.AuthController{}.Logout)
	g.POST("/auth/whoAmI", AuthController.AuthController{}.WhoAmI)

	// user
	g.POST("/user/changeUserInfo", UserServiceController.UserServiceController{}.ChangeUserInfo)
	g.POST("/user/register", UserServiceController.UserServiceController{}.RegisterUser)

	// verify code
	g.POST("/verifyCode/addVerifyCode", VerifyCodeController.VerifyCodeController{}.AddVerifyCodeController)
	g.POST("/verifyCode/viewAllVerifyCode", VerifyCodeController.VerifyCodeController{}.ViewAllVerifyCode)
	g.POST("/verifyCode/deleteVerifyCode", VerifyCodeController.VerifyCodeController{}.DeleteVerifyCode)
}
