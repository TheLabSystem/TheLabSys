package Router

import (
	"TheLabSystem/Controller/AuthController"
	"TheLabSystem/Controller/BillController"
	"TheLabSystem/Controller/DeviceController"
	"TheLabSystem/Controller/MentalListController"
	"TheLabSystem/Controller/NoticeController"
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
	g.POST("/user/addMoney", UserServiceController.UserServiceController{}.AddMoney)

	// verify code
	g.POST("/verifyCode/addVerifyCode", VerifyCodeController.VerifyCodeController{}.AddVerifyCodeController)
	g.POST("/verifyCode/viewAllVerifyCode", VerifyCodeController.VerifyCodeController{}.ViewAllVerifyCode)
	g.POST("/verifyCode/deleteVerifyCode", VerifyCodeController.VerifyCodeController{}.DeleteVerifyCode)

	// notice
	g.POST("/notice/addNotice", NoticeController.NoticeController{}.AddNotice)
	g.GET("/notice/getNotice", NoticeController.NoticeController{}.GetNoticeList)
	g.POST("/notice/deleteNotice", NoticeController.NoticeController{}.DeleteNotice)
	g.POST("/notice/updateNotice", NoticeController.NoticeController{}.UpdateNotice)

	// bill
	g.GET("/bill/getBill", BillController.BillController{}.GetBill)
	g.POST("/bill/payBill", BillController.BillController{}.PayBill)

	// mentorList service
	g.POST("/mentalList/addStudent", MentalListController.MentalListController{}.AddStudentController)
	g.POST("/mentalList/deleteStudent", MentalListController.MentalListController{}.DeleteStudentController)
	g.GET("/mentalList/viewStudent", MentalListController.MentalListController{}.ViewStudentController)

	// device
	g.POST("/device/addDevice", DeviceController.DeviceController{}.AddDevice)
	g.GET("/device/getDeviceType", DeviceController.DeviceController{}.GetDeviceType)
	g.POST("/device/updateDevice", DeviceController.DeviceController{}.UpdateDevice)
	g.POST("/device/deleteDevice", DeviceController.DeviceController{}.DeleteDevice)
}
