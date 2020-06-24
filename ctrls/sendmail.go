package ctrls

import (
	"democ/models"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

//设置验证码并发送邮件
func setAndSend(email string) {
	var code string = RandomString()
	if models.Exists(email) == 1 {
		g := models.Get(email)
		gcode := g.(string)
		sendMail(gcode, email)
	} else {
		sendMail(code, email)
		models.Set(email, code)
	}
}

//发送邮件
func sendMail(code string, sjr string) {
	//126邮箱授权码
	//BDXZQHHCQUMTJPTC
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "etinifninfinite@gmail.com", "注册验证邮件") // 发件人
	m.SetHeader("To", m.FormatAddress(sjr, "111"))                  //收件人
	m.SetHeader("Subject", "网站验证激活邮件")                              // 主题
	body := "点击链接即可激活账号，1小时内有效，请尽快完成激活：" + "http://localhost:8080/auth?email=" + sjr + "&" + "code=" + code
	//"<a href='http://localhost:8080/auth/?email=" + sjr + "&" + "code=" + code + ">点击激活</a><br>"
	m.SetBody("text/html", body) // 正文

	d := gomail.NewDialer("smtp.126.com", 465, "etinifninfinite@gmail.com", "HHCZQTJPTQUMBDXC") // 发送邮件服务器、端口、发件人账号、发件人密码

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Printf("邮件发送失败: %v/n", err)
	}
}

//生成 6位随机验证码
//func getRand() string {
//	rand.Seed(time.Now().Unix())
//	randnums := strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
//		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
//		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
//	return randnums
//}

//生成并返回长度为 16位的随机字符串
// RandomString returns a random string with a fixed length
func RandomString(allowedChars ...[]rune) string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//伪随机种子
	rand.Seed(time.Now().Unix())
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, 16) //长度
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
