package v1

import (
	"AM_system/model"
	"AM_system/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 查找用户是否存在
func StuExist(c *gin.Context) {

}

// 添加学生
func AddStu(c *gin.Context) {
	var data model.Student
	_ = c.ShouldBindJSON(&data)
	code = model.CheckStu(data.StuID)
	if code == errmsg.SUCCESS {
		model.CreateStu(&data)
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

// 查询单个学生
func GetStuSingle(c *gin.Context) {
	stu_id, _ := strconv.Atoi(c.Query("stu_id"))
	data := model.GetStuSingle(stu_id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		//"message":errmsg.GetErrMsg(code),
	})
}

// 查询学生列表
func GetStu(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := model.GetStu(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status1": code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑学生
func EditStu(c *gin.Context) {
	var data model.Student
	stu_id, _ := strconv.Atoi(c.Param("stu_id"))
	c.ShouldBindJSON(&data)
	code = model.CheckStu(stu_id)
	if code == errmsg.SUCCESS {
		c.Abort()
		code = errmsg.ERROR_STU_NOT_EXIST
	} else {
		model.EditStu(stu_id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除学生
func DeleteStu(c *gin.Context) {
	stu_id, _ := strconv.Atoi(c.Param("stu_id"))
	code = model.DeleteStuSingle(stu_id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
