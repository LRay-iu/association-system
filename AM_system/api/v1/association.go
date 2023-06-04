package v1

import (
	"AM_system/model"
	"AM_system/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//注册社团
func AddAss(c *gin.Context){
	var data model.Association
	_ = c.ShouldBindJSON(&data)
	code1 := model.CheckAssId(data.AssID)
	code2 :=model.CheckAssName(data.AssName)
	if (code1 == errmsg.SUCCESS&&code2==errmsg.SUCCESS) {
		model.CreateAss(&data)
	}
	//HTTP中的code	//if code1 == errmsg.ERROR_STUID_USED {
	//	//	code1 = errmsg.ERROR_STUID_USED
	//	//}
	c.JSON(http.StatusOK, gin.H{
		"status1":  code1,
		"status2":code2,
		"data":    data,
		"message1": errmsg.GetErrMsg(code1),
		"message2": errmsg.GetErrMsg(code2),
	})
}
//查询单个社团
func GetAssSingle(c *gin.Context){
	assName := c.Query("ass_name")
	data := model.GetAssSingle(assName)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询社团列表
func GetAss(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//limit(-1)相当于不做limit这个操作,默认查询所有学生
	//if pageSize == 0 {
	//	pageSize = -1
	//}
	//if pageNum == 0 {
	//	pageNum = -1
	//}
	data := model.GetAss(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
//编辑社团
func EditAss(c *gin.Context){
	var data model.Association
	ass_id, _ := strconv.Atoi(c.Param("ass_id"))
	c.ShouldBindJSON(&data)
	code = model.CheckStu(ass_id)
	if code == errmsg.SUCCESS {
		c.Abort()
		code = errmsg.ERROR_ASS_NOT_EXIST
	} else {
		model.EditAss(ass_id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除社团
func DeleteAss(c *gin.Context){
	//将字符类型转换成整型
	ass_id, _ := strconv.Atoi(c.Param("ass_id"))
	code=model.DeleteAssSingle(ass_id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
