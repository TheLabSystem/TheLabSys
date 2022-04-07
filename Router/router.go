package Router

import (
	"TheLabSystem/Controller/AuthController"
	"TheLabSystem/Controller/MentalListController"
	"TheLabSystem/Controller/UserServiceController"
	"TheLabSystem/Controller/MentalListController"
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
	g.POST("/user/findUserInfo", UserServiceController.UserServiceController{}.FindUserInfo)

	// verify code
	g.POST("/verifyCode/addVerifyCode", VerifyCodeController.VerifyCodeController{}.AddVerifyCodeController)
	g.POST("/verifyCode/viewAllVerifyCode", VerifyCodeController.VerifyCodeController{}.ViewAllVerifyCode)
	g.POST("/verifyCode/deleteVerifyCode", VerifyCodeController.VerifyCodeController{}.DeleteVerifyCode)

	// bill
	g.GET("/bill/getBill", BillController.BillController{}.GetBill)

	// mentorList service
	g.POST("/mentalList/addStudent", MentalListController.MentalListController{}.AddStudentController)
	g.POST("/mentalList/deleteStudent", MentalListController.MentalListController{}.DeleteStudentController)
}
