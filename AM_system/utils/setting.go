package utils

//把以下变量作为全局变量引用进来

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string


	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
)

func init() {
	file, err:= ini.Load("config/config.ini")
	if err !=nil{
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	//MustString作用：在[server]章节取不到值会取默认值debug
	HttpPort =file.Section("server").Key("HttpPort").MustString(":8000")
	JwtKey =file.Section("server").Key("JwtKey").MustString("89asd8")
}

func LoadData(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("denug")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("RayBOKI")
	DbPassword = file.Section("database").Key("DbPassword").MustString("10181024")
	DbName = file.Section("database").Key("DbName").MustString("AMsys")
}