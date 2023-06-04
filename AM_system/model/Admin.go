package model

import (
	"AM_system/utils/errmsg"
	"gorm.io/gorm"
)

type Admin struct {
	AdminID         int    `gorm:"primaryKey;column:admin_id"`
	AssociationID   int    `gorm:"column:association_id;foreignKey:AssociationID"`
	AdminName       string `gorm:"column:admin_name"`
	Password        string `gorm:"column:password"`
	Telephone       string `gorm:"column:telephone"`
}

// 设置表名
func (Admin) TableName() string {
	return "admin"
}

//查找用户是否存在
func CheckAdmin (admin_id int)int {
	var admin Admin
	db.Where("admin_id=?",admin_id).First(&admin)
	if admin.AdminName != ""{
		return errmsg.ERROR_ADMINID_USED
	}
	return errmsg.SUCCESS
}
//新增用户

func CreateAdmin (data *Admin)int {
	data.before()
	result := db.Create(&data)
	if result.Error!=nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}
//查询单个学生
func GetAdminSingle(admin_id int) Admin{
	var admin Admin
	result := db.Where("admin_id=?",admin_id).First(&admin)
	if result.Error!=nil&&result.Error!=gorm.ErrRecordNotFound{
		return Admin{}
	}
	return admin
}
//查询学生列表
func GetAdmin(pageSize int,pageNum int) []Admin{
	var admins []Admin
	//Limit(pageSize)：设置查询的结果集大小，限制每页返回的学生数量为 pageSize。
	//Offset((pageNum-1)*pageSize)：根据页码和页面大小计算出要偏移的结果集的起始位置。假设当前页码是 pageNum，则 (pageNum-1)*pageSize 表示前面页码的学生数量。
	//Find(&students)：执行查询操作，并将查询结果存储在 students 中。
	if pageSize>0{
		db =db.Limit(pageSize).Offset((pageNum-1)*pageSize)
	}
	result := db.Find(&admins)
	if result.Error !=nil&&result.Error!=gorm.ErrRecordNotFound{
		return nil
	}
	return admins
}

//编辑学生
//使用map更新，只会更新有变化的部分
//使用struct更新，只会更新有变化且非零值的元素
func EditAdmin(admin_id int,data *Admin) int {
	var admin Admin
	var maps = make(map[string]interface{})
	//maps["StuID"]=data.StuID
	maps["AdminID"]=admin_id
	maps["AdminName"] = data.AdminName
	maps["Password"] = data.Password
	maps["AssociationID"] = data.AssociationID
	maps["Telephone"] = data.Telephone
	result :=db.Model(&admin).Where("admin_id=?",admin_id).Updates(maps)
	if result.Error!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除学生
func DeleteAdminSingle(admin_id int) int{
	var admin Admin
	result := db.Where("admin_id = ?", admin_id).Delete(&admin)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ADMIN_NOT_EXIST
	}
	return errmsg.SUCCESS
}

func (s *Admin)before(){
	s.Password=scryptPw(s.Password)
}