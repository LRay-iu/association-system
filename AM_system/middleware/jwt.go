package middleware

import (
	"AM_system/utils"
	"AM_system/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	StuID       int    `json:"stu_id"`
	jwt.StandardClaims
}

//生成token
func SetToken(stu_id int) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		stu_id,
		jwt.StandardClaims{
			//令牌过期时间
			ExpiresAt: expireTime.Unix(),
			//令牌发布者
			Issuer: "AM_system",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims, int) {
	//func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error)
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims);setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取获取请求头中的Authorization字段的值，即JWT令牌。
		tokenHerder := c.Request.Header.Get("Authorization")
		//检查token是否为空
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//对令牌进行拆分，分隔符为空格。检查拆分后的结果是否包含两个部分，并且第一部分是否为"Bearer"
		//如果不符合条件，将code设置为errmsg.ERROR_TOKEN_TYPE_WRONG表示令牌格式错误，并中断请求处理。
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if (len(checkToken) != 2 && checkToken[0] != "Bearer") {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//验证token的有效性，将验证结果存储于key和Tcode中
		key, tcode := CheckToken(checkToken[1])
		if tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//检查令牌是否超时
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("usename", key.StuID)
	}
}
