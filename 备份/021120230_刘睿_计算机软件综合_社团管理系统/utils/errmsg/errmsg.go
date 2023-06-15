package errmsg

// golang声明常量的关键词
const (
	SUCCESS = 200
	ERROR   = 500
	//code = 1000...用户模块的错误
	ERROR_STUID_USED       = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_STU_NOT_EXIST    = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_ASSID_USED       = 1008
	ERROR_ASSNAME_USED     = 1009
	ERROR_ASS_NOT_EXIST    = 1010
	ERROR_ADMINID_USED     = 1011
	ERROR_ADMINNAME_USED   = 1012
	ERROR_ADMIN_NOT_EXIST  = 1013
	//code = 2000... 文章模块的错误

	//code = 3000... 分类模块的错误
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FATL",
	ERROR_STUID_USED:       "学生已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_STU_NOT_EXIST:    "学生不存在",
	ERROR_TOKEN_EXIST:      " TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
	ERROR_ASSID_USED:       "社团ID已被使用",
	ERROR_ASSNAME_USED:     "社团名已被使用",
	ERROR_ASS_NOT_EXIST:    "社团不存在",
	ERROR_ADMINID_USED:     "管理员ID已被使用",
	ERROR_ADMINNAME_USED:   "管理员姓名已经存在",
	ERROR_ADMIN_NOT_EXIST:  "管理员不存在",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
