package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// 把每一步数据库操作封装成为函数
// 待logic层根据业务需求调用

// 加盐的字符串
const secret = "LiShun71"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (bool, error) {
	// 1. 编写数据库语句
	sqlStr := `select count(user_id) from user where username = ?`

	var count int
	// 2. 执行sql语句
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}

	if count > 0 {
		return count > 0, errors.New("用户已经存在")
	}

	return count > 0, nil
}

func QueryUserByUserName(name string) (*models.User, error) {
	sqlStr := `select * from user where username = ?`
	var user models.User
	if err := db.Get(&user, sqlStr, name); err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByID() {

}

// InsertUser InsetUser 先数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	password := EncryptPassword(user.Password)
	sqlStr := `insert into user (user_id, username, password,email) values (?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password, user.Email)
	// 执行SQL语句入库
	return
}

// EncryptPassword 使用 md5 简单的加个密
func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := EncryptPassword(user.Password)
	sqlStr := `select user_id, username, password from user where username = ?`
	var us models.User
	err = db.Get(&us, sqlStr, user.Username)
	if err != nil {
		return errors.New("用户名不存在")
	}

	// 判断密码是否相等
	if us.Password != oPassword {
		// 用户不存在
		return errors.New("密码错误")
	}
	return
}

func GetUserByID(id int64) (*models.User, error) {
	sqlStr := `select * from user_id,user_name where user_id = ?`
	user := new(models.User)
	err := db.Get(user, sqlStr, id)
	return user, err
}
