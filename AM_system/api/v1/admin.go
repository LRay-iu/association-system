package v1

import (
	"AM_system/model"
	"AM_system/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加学生
func AddAdmin(c *gin.Context) {
	var data model.Admin
	_ = c.ShouldBindJSON(&data)
	code = model.CheckAdmin(data.AdminID)
	if code == errmsg.SUCCESS {
		model.CreateAdmin(&data)
	}
	if code == errmsg.ERROR_STUID_USED {
		code = errmsg.ERROR_STUID_USED
	}
	//HTTP中的code
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个学生
func GetAdminSingle(c *gin.Context) {
	admin_id, _ := strconv.Atoi(c.Query("admin_id"))
	data := model.GetAdminSingle(admin_id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		//"message":errmsg.GetErrMsg(code),
	})
}

//查询学生列表
func GetAdmin(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := model.GetAdmin(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status1":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑学生
func EditAdmin(c *gin.Context) {
	var data model.Admin
	admin_id, _ := strconv.Atoi(c.Param("admin_id"))
	c.ShouldBindJSON(&data)
	code = model.CheckAdmin(admin_id)
	if code == errmsg.SUCCESS {
		c.Abort()
		code = errmsg.ERROR_ADMIN_NOT_EXIST
	} else {
		model.EditAdmin(admin_id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除学生
func DeleteAdmin(c *gin.Context) {
	admin_id, _ := strconv.Atoi(c.Param("admin_id"))
	code = model.DeleteAdminSingle(admin_id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}