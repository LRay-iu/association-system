package model

import (
	"AM_system/utils/errmsg"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type Student struct {
	StuID       int    `gorm:"primaryKey;column:stu_id"`
	StuName     string `gorm:"column:stu_name"`
	Sex         string `gorm:"column:sex"`
	College     string `gorm:"column:college"`
	Major       string `gorm:"column:major"`
	StuPassword string `gorm:"column:stu_password"`
}

// 设置表名
func (Student) TableName() string {
	return "student"
}

//查找用户是否存在
func CheckStu (stu_id int)int {
	var stu Student
	db.Where("stu_id=?",stu_id).First(&stu)
	fmt.Println("test:",stu.StuID)
	if stu.StuName != ""{
		return errmsg.ERROR_STUID_USED
	}
	return errmsg.SUCCESS
}
//新增用户

func CreateStu (data *Student)int {
	data.before()
	result := db.Create(&data)
	if result.Error!=nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}
//查询单个学生
func GetStuSingle(stu_id int) Student{
	var students Student
	result := db.Where("stu_id=?",stu_id).First(&students)
	if result.Error!=nil&&result.Error!=gorm.ErrRecordNotFound{
		return Student{}
	}
	return students
}
//查询学生列表
func GetStu(pageSize int,pageNum int) []Student{
	var students []Student
	//Limit(pageSize)：设置查询的结果集大小，限制每页返回的学生数量为 pageSize。
	//Offset((pageNum-1)*pageSize)：根据页码和页面大小计算出要偏移的结果集的起始位置。假设当前页码是 pageNum，则 (pageNum-1)*pageSize 表示前面页码的学生数量。
	//Find(&students)：执行查询操作，并将查询结果存储在 students 中。
	if pageSize>0{
		db =db.Limit(pageSize).Offset((pageNum-1)*pageSize)
	}
	result := db.Find(&students)
	if result.Error !=nil&&result.Error!=gorm.ErrRecordNotFound{
		return nil
	}
	return students
}

//编辑学生
//使用map更新，只会更新有变化的部分
//使用struct更新，只会更新有变化且非零值的元素
func EditStu(stu_id int,data *Student) int {
	var student Student
	var maps = make(map[string]interface{})
	//maps["StuID"]=data.StuID
	maps["StuID"]=stu_id
	maps["StuName"] = data.StuName
	maps["StuPassword"] = data.StuPassword
	maps["College"] = data.College
	maps["Major"] = data.Major
	maps["Sex"] = data.Sex
	result :=db.Model(&student).Where("stu_id=?",stu_id).Updates(maps)
	if result.Error!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除学生
func DeleteStuSingle(stu_id int) int{
	var student Student
	result := db.Where("stu_id = ?", stu_id).Delete(&student)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_STU_NOT_EXIST
	}
	return errmsg.SUCCESS
}
//密码加密
func scryptPw(password string) string {
	const Keylen =10
	salt := make([]byte,8)
	salt =[]byte{12,23,4,12,32,45,47,23}
	//func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	HashPw,err :=scrypt.Key([]byte(password),salt,16384,8,1,Keylen)
	if err!=nil{
		log.Fatal(err)
	}
	//最终密码
	FPw :=base64.StdEncoding.EncodeToString(HashPw)
	return FPw
}
func (s *Student)before(){
	s.StuPassword=scryptPw(s.StuPassword)
}

//登录验证

func CheckLogin(stu_id int,stu_password string)int{
	var student Student
	db.Where("stu_id=?",stu_id).First(&student)
	if student.StuName==""{
		return errmsg.ERROR_STU_NOT_EXIST
	}
	if scryptPw(stu_password)!=student.StuPassword{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}