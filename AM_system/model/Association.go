package model

import (
	"AM_system/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Association struct {
	AssID    int    `gorm:"primaryKey;column:association_id"`
	AssName  string `gorm:"column:association_name"`
	AssNote  string `gorm:"column:association_note"`
	AssScale int    `gorm:"column:association_scale;default:0"`
}

// 设置表名
func (Association) TableName() string {
	return "association"
}

//注册社团
func CreateAss(data *Association) int {
	result := db.Create(&data)
	if result.Error != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS

}

//查找社团是否存在
func CheckAssId(ass_id int) int {
	var ass Association
	db.Where("association_id=?", ass_id).First(&ass)
	fmt.Println("test:", ass.AssID)
	if ass.AssName != "" {
		return errmsg.ERROR_ASSID_USED
	}
	return errmsg.SUCCESS
}
func CheckAssName(ass_name string) int {
	var ass Association
	db.Where("association_name=?", ass_name).First(&ass)
	fmt.Println("test:", ass.AssName)
	if ass.AssName != "" {
		return errmsg.ERROR_ASSNAME_USED
	}
	return errmsg.SUCCESS
}
//查看所有社团
func GetAss(pageSize int,pageNum int)[]Association{
	var ass []Association
	if pageSize > 0 {
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := db.Find(&ass)
	if result.Error !=nil&&result.Error!=gorm.ErrRecordNotFound{
		fmt.Println(result.Error)
		return nil
	}
	return ass
}
//查看单个社团
func GetAssSingle(ass_name string) Association{
	var ass Association
	result := db.Where("association_name=?",ass_name).First(&ass)
	if result.Error!=nil&&result.Error!=gorm.ErrRecordNotFound{
		return Association{}
	}
	return ass
}
//修改社团信息
func EditAss(ass_id int,data *Association) int {
	var ass Association
	var maps = make(map[string]interface{})
	//maps["StuID"]=data.StuID
	maps["AssID"]=ass_id
	maps["AssName"] = data.AssName
	maps["AssNote"] = data.AssNote
	maps["AssScale"] = data.AssScale
	result :=db.Model(&ass).Where("association_id=?",ass_id).Updates(maps)
	if result.Error!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//删除社团
func DeleteAssSingle (ass_id int)int{
	var ass Association
	result := db.Where("association_id = ?", ass_id).Delete(&ass)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ASS_NOT_EXIST
	}
	return errmsg.SUCCESS
}