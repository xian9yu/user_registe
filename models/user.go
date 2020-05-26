package models

import (
	"errors"
	"log"
)

type User struct {
	Id       int64  `xorm:"id pk unique autoincr notnull",json:"id"`
	Email    string `xorm:"unique notnull",json:"email"`
	Password string `xorm:"notnull",json:"password"`
	Status   int64  `xorm:"notnull default(0)",json:"status"`
}

//用户注册
func (user *User) Register() bool {
	_, err := DB.InsertOne(user)
	if err != nil {
		log.Fatal("Failed to create user:", err)
		return false
	}
	return true
}

//检查邮箱是否已存在
func (user *User) CheckEmailisExist(email string) bool {
	//var err error
	has, err := DB.Where("email=?", email).Exist(&User{Email: email})
	if err != nil {
		log.Fatal("检查用户存在失败 :", err)
	}
	//fmt.Println(has)
	return has

}

//CheckUserStatus
func (user *User) AuthStatus(email string) (*User, error) {
	var users []User
	err := DB.Where("email=?", email).Cols("status").Find(&users)
	if len(users) == 0 {
		return nil, errors.New("not found data")
	}
	return &users[0], err

}

//updateStatus
func (user *User) UpdateStatus(email string) error {
	//_, err := DB.Where("email=?", email).Cols("status").Update(&user)
	_, err := DB.Exec("update user set status = ? where email = ?", 1, email)
	return err
}

//获取全部用户
func (user *User) GetAll(start, count int64) []User {
	var users []User
	err := DB.Desc("id").Find(&users)
	if err != nil {
		log.Fatal("Failed to get all users:", err)
	}
	return users
}

//get one by id
func (user *User) GetOne(id int64) (*User, error) {
	var users []User
	err := DB.Where("id = ?", id).Cols("id", "email", "status").Find(&users)
	if len(users) == 0 {
		return nil, errors.New("not found page")
	}
	return &users[0], err
}
