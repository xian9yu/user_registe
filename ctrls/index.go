package ctrls

import (
	"net/http"
	"strconv"
	"user_register/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "userList.html", gin.H{
		"Title": "用户列表",
		"user":  GetAllUser(),
	})
}

//GetAll 获取用户列表
func GetAllUser() []models.User {
	user := &models.User{}
	users := user.GetAll(0, 5)
	return users

}

//GetOne
func GetOneUser(c *gin.Context) {
	users := &models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := users.GetOne(int64(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "用户不存在", "/")
		return
	}
	c.HTML(http.StatusOK, "userInfo.html", gin.H{
		"user":  user,
		"email": user.Email,
	})
}
