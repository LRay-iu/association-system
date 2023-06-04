package routes

import (
	v1 "AM_system/api/v1"
	"AM_system/middleware"
	"AM_system/utils"
	"github.com/gin-gonic/gin"
)

//golang对于私有公用取决于变量的首字母是否是大写
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v1")
	//需要连接中间件token的部分
	auth.Use(middleware.JwtToken())
	{
		//Student模块的；路由接口

		auth.PUT("stu/:stu_id", v1.EditStu)
		auth.DELETE("stu/:stu_id", v1.DeleteStu)
		//社团模块的路由接口

		auth.PUT("ass/:ass_id", v1.EditAss)
		auth.DELETE("ass/:ass_id", v1.DeleteAss)
		//社团管理员的路由接口

		auth.PUT("admin/:admin_id", v1.EditAdmin)
		auth.DELETE("admin/:admin_id", v1.DeleteAdmin)
		//社团活动模块的路由接口
	}
	router := r.Group("api/v1")
	{
		router.POST("stu/add", v1.AddStu)
		router.POST("ass/add", v1.AddAss)
		router.POST("admin/add", v1.AddAdmin)
		router.GET("stu/ask", v1.GetStu)
		router.GET("stu/search", v1.GetStuSingle)
		router.GET("ass/ask", v1.GetAss)
		router.GET("ass/search", v1.GetAssSingle)
		router.GET("admin/ask", v1.GetAdmin)
		router.GET("admin/search", v1.GetAdminSingle)
		router.POST("login",v1.Login)
	}
	r.Run(utils.HttpPort)
}
