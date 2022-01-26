package main

import (
	"log"
	"user_register/ctrls"
	"user_register/models"

	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	router.GET("/", ctrls.Index)
	router.GET("/user/:id", ctrls.GetOneUser)
	router.GET("/register", ctrls.Register)
	router.GET("/auth", ctrls.CheckAuthLink)
	router.POST("/register", ctrls.Register)
	//router.POST("/register/:on_click", ctrls.Register)
	router.POST("/login", ctrls.Login)

}

func main() {
	//初始化sql
	models.InitSQL()

	//在控制台打印 sql信息，便于调试
	models.DB.ShowSQL(true)

	//初始化router
	r := gin.Default()
	initRouter(r)

	//加载views目录下的所有html模板文件
	r.LoadHTMLGlob("./views/*")

	//使用gin自带的异常恢复中间件，避免出现异常时程序退出
	r.Use(gin.Recovery())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
