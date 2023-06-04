package v1

import (
	"AM_system/middleware"
	"AM_system/model"
	"AM_system/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	var data model.Student
	c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.StuID,data.StuPassword)
	if code == errmsg.SUCCESS{
		token,code =middleware.SetToken(data.StuID)
	}
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}
