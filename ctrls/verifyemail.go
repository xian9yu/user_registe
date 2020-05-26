package ctrls

import "regexp"

//验证邮箱后缀是否正确
//func verifyEmailFormat(email string) bool {
//	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
//	reg, _ := regexp.Compile(pattern)
//	return reg.MatchString(email)
//}

//验证是否为系统指定邮箱后缀，但不检测邮箱body
//func verifyEmail(email string) bool {
//	arr := strings.Split(email, "@")
//
//	//仅允许
//	//"163.com"+"126.com"+"qq.com"+"gmail.com"+"yahoo.com"+"hotmail.com"+"outlook.com"+"foxmail.com"+"sohu.com"+"tom.com"+"163.net"+"aliyun.com"
//	//等12种邮箱后缀的用户注册账户
//	mapMailSuffix := map[string]bool{"163.com": true, "126.com": true, "qq.com": true, "gmail.com": true, "yahoo.com": true, "hotmail.com": true, "outlook.com": true, "foxmail.com": true, "sohu.com": true, "tom.com": true, "163.net": true, "aliyun.com": true}
//	if _, ok := mapMailSuffix[arr[1]]; !ok {
//		return false
//	}
//	return true
//}

//验证邮箱 body+后缀
func verifyEmail(email string) bool {
	if ok, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\b(gmail|qq|outlook|163|126|yahoo|hotmail|foxmail|soho|aliyun|yahoo|tom)$*\.\w+([-.]\w+)*$`, email); !ok {
		return false
	}
	return true
}
