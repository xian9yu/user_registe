package ctrls

import (
	"democ/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Check authLink
func CheckAuthLink(c *gin.Context) {
	user := new(models.User)
	auth := new(models.AuthRedis)
	// c.Query() 等同于 c.Request.URL.Query().Get()
	user.Email = c.Query("email")
	auth.Code = c.Query("code")

	val, err := user.AuthStatus(user.Email)
	if err != nil {
		log.Fatal("检查用户状态失败")
	}

	value := val.Status
	//a := models.Get(user.Email)
	if value == 1 {
		c.HTML(500, "authOK.html", gin.H{
			"Message": "不能重复激活",
			"GotoURL": "/login",
		})
	} else if value == 0 && models.Get(user.Email) == auth.Code {
		errs := user.UpdateStatus(user.Email)
		if errs != nil {
			log.Fatal("修改用户状态失败")
		}
		c.HTML(200, "authOK.html", gin.H{
			"Message": "邮箱验证成功",
			"GotoURL": "/login",
		})
	} else {
		c.HTML(500, "authOK.html", gin.H{
			"Message": "验证失败，请联系系统管理人员",
			"GotoURL": "/",
		})

	}
}

//用户注册
func Register(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", nil)

	} else if c.Request.Method == "POST" {
		user := new(models.User)
		user.Email = c.PostForm("email")
		user.Password = c.PostForm("password")
		//md5password := fmt.Sprintf("%x", md5.Sum([]byte(user.Email+user.Password)))
		if verifyEmail(user.Email) { //校验邮箱后缀
			if user.CheckEmailisExist(user.Email) {
				c.HTML(http.StatusOK, "register.html", gin.H{
					"Tip":     "提示：该邮箱已存在",
					"GotoURL": "/login",
				})
			} else {
				user.Register()
				setAndSend(user.Email)
				c.HTML(http.StatusOK, "msg.html", gin.H{
					"Status":  "success",
					"Message": "注册成功，请尽快登录邮箱点击链接激活账号",
					"GotoURL": "/",
				})
			}
		} else {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"Tip": "提示：输入有误或邮箱后缀不正确，请重新输入，仅支持以下后缀的邮箱：\"@163.com\",\"@126.com\",\"@qq.com\",\"@gmail.com\",\"@yahoo.com\",\"@hotmail.com\",\"@outlook.com\",\"@foxmail.com\",\"@sohu.com\",\"@tom.com\",\"@163.net\",\"@aliyun.com\"",
			})
		}

	}
}

//用户登录
func Login(c *gin.Context) {

}
