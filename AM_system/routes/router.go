package routes

import (
	v1 "AM_system/api/v1"
	"AM_system/middleware"
	"AM_system/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// golang对于私有公用取决于变量的首字母是否是大写
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v1")
	go func() {
		for {
			backup()
			time.Sleep(5 * time.Minute) // 每5分钟执行一次备份
		}
	}()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	// 在 Gin 路由中使用 CORS 中间件
	r.Use(cors.New(config))
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
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
func backup() {
	// 获取当前时间作为备份文件名
	currentTime := time.Now().Format("20060102150405")
	backupDir := "D:/RayBOKI/Documents/copy"

	// 创建备份文件夹
	err := os.MkdirAll(filepath.FromSlash(backupDir), os.ModePerm)
	if err != nil {
		log.Println("创建备份文件夹失败:", err)
		return
	}

	sourceDir := "D:\\Program Files\\data\\MySQL\\MySQL Server 8.0\\Data\\amsys"
	// 执行备份操作，将目标文件或文件夹复制到备份文件夹中
	cmd := exec.Command("xcopy", sourceDir, filepath.Join(backupDir, currentTime), "/E", "/I", "/H", "/Y")
	err = cmd.Run()
	if err != nil {
		log.Println("备份失败:", err)
		return
	}

	log.Println("备份成功")
}
