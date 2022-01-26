package models

type User struct {
	Uid         uint64 `json:"uid" gorm:"size:12;primaryKey;unique;notnull;comment:用户id"`
	Username    string `json:"username" gorm:"size:60;comment:用户名"`
	Mail        string `json:"mail" gorm:"size:60;notnull;comment:邮箱"`
	Password    string `json:"-"  gorm:"size:33;notnull;comment:用户登录密码"`
	Url         string `json:"url" gorm:"size:60;comment:网站url"`
	Group       string `json:"group" gorm:"size:33;notnull;comment:用户分组"`
	Status      int64  `json:"status" gorm:"notnull default(0)" `
	CreatedTime uint64 `json:"created_time" gorm:"autoCreateTime;notnull comment:user创建时间"`
	UpdatedTime uint64 `json:"updated_time" gorm:"autoUpdateTime;comment:上一次修改信息时间"`
	LastLogin   uint64 `json:"last_login" gorm:"comment:上一次登录时间"`
}

////用户注册
//func (user *User) Register() bool {
//	_, err := DB.InsertOne(user)
//	if err != nil {
//		log.Fatal("Failed to create user:", err)
//		return false
//	}
//	return true
//}
//
////检查邮箱是否已存在
//func (user *User) CheckEmailisExist(email string) bool {
//	//var err error
//	has, err := DB.Where("email=?", email).Exist(&User{Email: email})
//	if err != nil {
//		log.Fatal("检查用户存在失败 :", err)
//	}
//	//fmt.Println(has)
//	return has
//
//}
//
////CheckUserStatus
//func (user *User) AuthStatus(email string) (*User, error) {
//	var users []User
//	err := DB.Where("email=?", email).Cols("status").Find(&users)
//	if len(users) == 0 {
//		return nil, errors.New("not found data")
//	}
//	return &users[0], err
//
//}
//
////updateStatus
//func (user *User) UpdateStatus(email string) error {
//	//_, err := DB.Where("email=?", email).Cols("status").Update(&user)
//	_, err := DB.Exec("update user set status = ? where email = ?", 1, email)
//	return err
//}
//
////获取全部用户
//func (user *User) GetAll(start, count int64) []User {
//	var users []User
//	err := DB.Desc("id").Find(&users)
//	if err != nil {
//		log.Fatal("Failed to get all users:", err)
//	}
//	return users
//}
//
////get one by id
//func (user *User) GetOne(id int64) (*User, error) {
//	var users []User
//	err := DB.Where("id = ?", id).Cols("id", "email", "status").Find(&users)
//	if len(users) == 0 {
//		return nil, errors.New("not found page")
//	}
//	return &users[0], err
//}
